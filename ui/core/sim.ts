import { Class } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Enchant } from '/tbc/core/proto/common.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Gem } from '/tbc/core/proto/common.js';
import { GemColor } from '/tbc/core/proto/common.js';
import { ItemQuality } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { ItemType } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { ComputeStatsRequest, ComputeStatsResult } from '/tbc/core/proto/api.js';
import { GearListRequest, GearListResult } from '/tbc/core/proto/api.js';
import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { SimOptions } from '/tbc/core/proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from '/tbc/core/proto/api.js';

import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { SimResult } from '/tbc/core/proto_utils/sim_result.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { gemEligibleForSocket } from '/tbc/core/proto_utils/gems.js';
import { gemMatchesSocket } from '/tbc/core/proto_utils/gems.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { SpecTalents } from '/tbc/core/proto_utils/utils.js';
import { SpecTypeFunctions } from '/tbc/core/proto_utils/utils.js';
import { specTypeFunctions } from '/tbc/core/proto_utils/utils.js';
import { SpecOptions } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { getEligibleItemSlots } from '/tbc/core/proto_utils/utils.js';
import { getEligibleEnchantSlots } from '/tbc/core/proto_utils/utils.js';

import { Encounter } from './encounter.js';
import { Player } from './player.js';
import { Raid } from './raid.js';
import { Listener } from './typed_event.js';
import { EventID, TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
import { wait } from './utils.js';
import { WorkerPool } from './worker_pool.js';

import * as OtherConstants from '/tbc/core/constants/other.js';

export type RaidSimData = {
	request: RaidSimRequest,
	result: RaidSimResult,
};

export type StatWeightsData = {
	request: StatWeightsRequest,
	result: StatWeightsResult,
};

// Core Sim module which deals only with api types, no UI-related stuff.
export class Sim {
	private readonly workerPool: WorkerPool;

	private iterations: number = 3000;
  private phase: number = OtherConstants.CURRENT_PHASE;

  readonly raid: Raid;
  readonly encounter: Encounter;

  // Database
  private items: Record<number, Item> = {};
  private enchants: Record<number, Enchant> = {};
  private gems: Record<number, Gem> = {};

  readonly iterationsChangeEmitter = new TypedEvent<void>();
  readonly phaseChangeEmitter = new TypedEvent<void>();

  // Emits when any of the above emitters emit.
  readonly changeEmitter: TypedEvent<void>;

	// Fires when a raid sim API call completes.
  readonly simResultEmitter = new TypedEvent<SimResult>();

	private readonly _initPromise: Promise<void>;

	// These callbacks are needed so we can apply BuffBot modifications automatically before sending requests.
	private modifyRaidProto: ((raidProto: RaidProto) => void) = () => {};
	private modifyEncounterProto: ((encounterProto: EncounterProto) => void) = () => {};

  constructor() {
		this.workerPool = new WorkerPool(3);

    this._initPromise = this.workerPool.getGearList(GearListRequest.create()).then(result => {
			result.items.forEach(item => this.items[item.id] = item);
			result.enchants.forEach(enchant => this.enchants[enchant.id] = enchant);
			result.gems.forEach(gem => this.gems[gem.id] = gem);
		});

		this.raid = new Raid(this);
    this.encounter = new Encounter(this);

		
		this.changeEmitter = TypedEvent.onAny([
      this.iterationsChangeEmitter,
      this.phaseChangeEmitter,
			this.raid.changeEmitter,
			this.encounter.changeEmitter,
		]);

		this.raid.changeEmitter.on(eventID => this.updateCharacterStats(eventID));
  }

  waitForInit(): Promise<void> {
		return this._initPromise;
  }

	setModifyRaidProto(newModFn: (raidProto: RaidProto) => void) {
		this.modifyRaidProto = newModFn;
	}
	getModifiedRaidProto(): RaidProto {
		const raidProto = this.raid.toProto();
		this.modifyRaidProto(raidProto);

		// Remove any inactive meta gems, since the backend doesn't have its own validation.
		raidProto.parties.forEach(party => {
			party.players.forEach(player => {
				if (!player.equipment) {
					return;
				}

				const gear = this.lookupEquipmentSpec(player.equipment);
				if (gear.hasInactiveMetaGem()) {
					player.equipment = gear.withoutMetaGem().asSpec();
				}
			});
		});

		return raidProto;
	}

	setModifyEncounterProto(newModFn: (encounterProto: EncounterProto) => void) {
		this.modifyEncounterProto = newModFn;
	}
	getModifiedEncounterProto(): EncounterProto {
		const encounterProto = this.encounter.toProto();
		this.modifyEncounterProto(encounterProto);
		return encounterProto;
	}

  private makeRaidSimRequest(debug: boolean): RaidSimRequest {
		return RaidSimRequest.create({
			raid: this.getModifiedRaidProto(),
			encounter: this.getModifiedEncounterProto(),
			simOptions: SimOptions.create({
				iterations: debug ? 1 : this.getIterations(),
				debugFirstIteration: true,
			}),
		});
  }

  async runRaidSim(eventID: EventID): Promise<SimResult> {
		if (this.raid.isEmpty()) {
			throw new Error('Raid is empty! Try adding some players first.');
		} else if (this.encounter.getNumTargets() < 1) {
			throw new Error('Encounter has no targets! Try adding some targets first.');
		}

		await this.waitForInit();

		const request = this.makeRaidSimRequest(false);
		const result = await this.workerPool.raidSim(request);

		const simResult = await SimResult.makeNew(request, result);
		this.simResultEmitter.emit(eventID, simResult);
		return simResult;
	}

  async runRaidSimWithLogs(eventID: EventID): Promise<SimResult> {
		if (this.raid.isEmpty()) {
			throw new Error('Raid is empty! Try adding some players first.');
		} else if (this.encounter.getNumTargets() < 1) {
			throw new Error('Encounter has no targets! Try adding some targets first.');
		}

		await this.waitForInit();

		const request = this.makeRaidSimRequest(true);
		const result = await this.workerPool.raidSim(request);

		const simResult = await SimResult.makeNew(request, result);
		this.simResultEmitter.emit(eventID, simResult);
		return simResult;
	}

	// This should be invoked internally whenever stats might have changed.
	private async updateCharacterStats(eventID: EventID) {
		await this.waitForInit();

		// Capture the current players so we avoid issues if something changes while
		// request is in-flight.
		const players = this.raid.getPlayers();

		const result = await this.workerPool.computeStats(ComputeStatsRequest.create({
			raid: this.getModifiedRaidProto(),
		}));

		TypedEvent.freezeAllAndDo(() => {
			result.raidStats!.parties
					.forEach((partyStats, partyIndex) =>
							partyStats.players.forEach((playerStats, playerIndex) =>
									players[partyIndex*5 + playerIndex]?.setCurrentStats(eventID, playerStats)));
		});
	}

  async statWeights(player: Player<any>, epStats: Array<Stat>, epReferenceStat: Stat): Promise<StatWeightsResult> {
		if (this.raid.isEmpty()) {
			throw new Error('Raid is empty! Try adding some players first.');
		} else if (this.encounter.getNumTargets() < 1) {
			throw new Error('Encounter has no targets! Try adding some targets first.');
		}

		await this.waitForInit();

		if (player.getParty() == null) {
			console.warn('Trying to get stat weights without a party!');
			return StatWeightsResult.create();
		} else {
			const request = StatWeightsRequest.create({
				player: player.toProto(),
				raidBuffs: this.raid.getBuffs(),
				partyBuffs: player.getParty()!.getBuffs(),
				encounter: this.encounter.toProto(),
				simOptions: SimOptions.create({
					iterations: this.getIterations(),
					debug: false,
				}),

				statsToWeigh: epStats,
				epReferenceStat: epReferenceStat,
			});

			return await this.workerPool.statWeights(request);
		}
	}

	getItems(slot: ItemSlot | undefined): Array<Item> {
		let items = Object.values(this.items);
		if (slot != undefined) {
			items = items.filter(item => getEligibleItemSlots(item).includes(slot));
		}
		return items;
	}

	getEnchants(slot: ItemSlot | undefined): Array<Enchant> {
		let enchants = Object.values(this.enchants);
		if (slot != undefined) {
			enchants = enchants.filter(enchant => getEligibleEnchantSlots(enchant).includes(slot));
		}
		return enchants;
	}

  getGems(socketColor: GemColor | undefined): Array<Gem> {
    let gems = Object.values(this.gems);
		if (socketColor) {
			gems = gems.filter(gem => gemEligibleForSocket(gem, socketColor));
		}
		return gems;
  }

	getMatchingGems(socketColor: GemColor): Array<Gem> {
    return Object.values(this.gems).filter(gem => gemMatchesSocket(gem, socketColor));
	}
  
  getPhase(): number {
    return this.phase;
  }
  setPhase(eventID: EventID, newPhase: number) {
    if (newPhase != this.phase) {
      this.phase = newPhase;
      this.phaseChangeEmitter.emit(eventID);
    }
  }
  
  getIterations(): number {
    return this.iterations;
  }
  setIterations(eventID: EventID, newIterations: number) {
    if (newIterations != this.iterations) {
      this.iterations = newIterations;
      this.iterationsChangeEmitter.emit(eventID);
    }
  }

  lookupItemSpec(itemSpec: ItemSpec): EquippedItem | null {
    const item = this.items[itemSpec.id];
    if (!item)
      return null;

    const enchant = this.enchants[itemSpec.enchant] || null;
    const gems = itemSpec.gems.map(gemId => this.gems[gemId] || null);

    return new EquippedItem(item, enchant, gems);
  }

  lookupEquipmentSpec(equipSpec: EquipmentSpec): Gear {
    // EquipmentSpec is supposed to be indexed by slot, but here we assume
    // it isn't just in case.
    const gearMap: Partial<Record<ItemSlot, EquippedItem | null>> = {};

    equipSpec.items.forEach(itemSpec => {
      const item = this.lookupItemSpec(itemSpec);
      if (!item)
        return;

      const itemSlots = getEligibleItemSlots(item.item);

      const assignedSlot = itemSlots.find(slot => !gearMap[slot]);
      if (assignedSlot == null)
        throw new Error('No slots left to equip ' + Item.toJsonString(item.item));

      gearMap[assignedSlot] = item;
    });

    return new Gear(gearMap);
  }

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'raid': this.raid.toJson(),
      'encounter': this.encounter.toJson(),
    };
	}

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(eventID: EventID, obj: any, spec?: Spec) {
		TypedEvent.freezeAllAndDo(() => {
			// For legacy format. Do not remove this until 2022/01/05 (1 month).
			if (obj['sim']) {
				if (!obj['raid']) {
					obj['raid'] = {
						'parties': [
							{
								'players': [
									{
										'spec': spec,
										'player': obj['player'],
									},
								],
								'buffs': obj['sim']['partyBuffs'],
							},
						],
						'buffs': obj['sim']['raidBuffs'],
					};
					obj['raid']['parties'][0]['players'][0]['player']['buffs'] = obj['sim']['individualBuffs'];
				}
			}

			if (obj['raid']) {
				this.raid.fromJson(eventID, obj['raid']);
			}

			if (obj['encounter']) {
				this.encounter.fromJson(eventID, obj['encounter']);
			}
		});
  }
}
