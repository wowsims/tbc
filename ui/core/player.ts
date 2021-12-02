import { Class } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Enchant } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
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
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { PlayerOptions as PlayerOptionsProto } from '/tbc/core/proto/api.js';
import { ComputeStatsRequest, ComputeStatsResult } from '/tbc/core/proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from '/tbc/core/proto/api.js';

import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';

import {
	Faction,
	SpecRotation,
	SpecTalents,
	SpecTypeFunctions,
	SpecOptions,
	canEquipItem,
	getEligibleEnchantSlots,
	gemEligibleForSocket,
	getEligibleItemSlots,
	getMetaGemEffectEP,
	gemMatchesSocket,
	raceToFaction,
	specToClass,
	specToEligibleRaces,
	specTypeFunctions,
	withSpecProto,
} from '/tbc/core/proto_utils/utils.js';

import { Listener } from './typed_event.js';
import { TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';
import { sum } from './utils.js';
import { wait } from './utils.js';
import { WorkerPool } from './worker_pool.js';

// Manages all the gear / consumes / other settings for a single Player.
export class Player<SpecType extends Spec> {
	readonly sim: Sim;

  readonly spec: Spec;
  private consumes: Consumes = Consumes.create();
  private bonusStats: Stats = new Stats();
  private gear: Gear = new Gear({});
  private race: Race;
  private rotation: SpecRotation<SpecType>;
  private talents: SpecTalents<SpecType>;
  private talentsString: string = '';
  private specOptions: SpecOptions<SpecType>;

  readonly specTypeFunctions: SpecTypeFunctions<SpecType>;

	private epWeights: Stats = new Stats();
	private currentStats: ComputeStatsResult;

  readonly consumesChangeEmitter = new TypedEvent<void>();
  readonly bonusStatsChangeEmitter = new TypedEvent<void>();
  readonly gearChangeEmitter = new TypedEvent<void>();
  readonly raceChangeEmitter = new TypedEvent<void>();
  readonly rotationChangeEmitter = new TypedEvent<void>();
  readonly talentsChangeEmitter = new TypedEvent<void>();
  // Talents dont have all fields so we need this.
  readonly talentsStringChangeEmitter = new TypedEvent<void>();
  readonly specOptionsChangeEmitter = new TypedEvent<void>();

  readonly currentStatsEmitter = new TypedEvent<void>();

  // Emits when any of the above emitters emit.
  readonly changeEmitter = new TypedEvent<void>();

  constructor(spec: Spec, sim: Sim) {
		this.sim = sim;

    this.spec = spec;
    this.race = specToEligibleRaces[this.spec][0];
    this.specTypeFunctions = specTypeFunctions[this.spec] as SpecTypeFunctions<SpecType>;
		this.rotation = this.specTypeFunctions.rotationCreate();
    this.talents = this.specTypeFunctions.talentsCreate();
		this.specOptions = this.specTypeFunctions.optionsCreate();

    [
      this.consumesChangeEmitter,
      this.bonusStatsChangeEmitter,
      this.gearChangeEmitter,
      this.raceChangeEmitter,
      this.rotationChangeEmitter,
      this.talentsChangeEmitter,
      this.talentsStringChangeEmitter,
      this.specOptionsChangeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));

		this.currentStats = ComputeStatsResult.create();
		this.sim.changeEmitter.on(() => {
			this.updateCharacterStats();
		});
		this.changeEmitter.on(() => {
			this.updateCharacterStats();
		});
  }

	// Returns all items that this player can wear in the given slot.
	getItems(slot: ItemSlot | undefined): Array<Item> {
		return this.sim.getItems(slot).filter(item => canEquipItem(item, this.spec));
	}

	// Returns all enchants that this player can wear in the given slot.
	getEnchants(slot: ItemSlot | undefined): Array<Enchant> {
		return this.sim.getEnchants(slot);
	}

	// Returns all gems that this player can wear of the given color.
  getGems(socketColor: GemColor | undefined): Array<Gem> {
		return this.sim.getGems(socketColor);
  }

	getEpWeights(): Stats {
		return this.epWeights;
	}

	setEpWeights(newEpWeights: Stats) {
		this.epWeights = newEpWeights;
	}

  async statWeights(request: StatWeightsRequest): Promise<StatWeightsResult> {
		const result = await this.sim.statWeights(request);
		this.epWeights = new Stats(result.epValues);
		return result;
	}

	// This should be invoked internally whenever stats might have changed.
	private async updateCharacterStats() {
		// Sometimes a ui change triggers other changes, so waiting a bit makes sure
		// we get all of them.
		await wait(10);


		const computeStatsResult = await this.sim.computeStats(ComputeStatsRequest.create({
			player: this.toProto(),
			raidBuffs: this.sim.getRaidBuffs(),
			partyBuffs: this.sim.getPartyBuffs(),
			individualBuffs: this.sim.getIndividualBuffs(),
		}));

		this.currentStats = computeStatsResult;
		this.currentStatsEmitter.emit();
	}

	getCurrentStats(): ComputeStatsResult {
		return ComputeStatsResult.clone(this.currentStats);
	}
  
  getRace(): Race {
    return this.race;
  }
  setRace(newRace: Race) {
    if (newRace != this.race) {
      this.race = newRace;
      this.raceChangeEmitter.emit();
    }
  }

	getFaction(): Faction {
		return raceToFaction[this.getRace()];
	}

  getConsumes(): Consumes {
    // Make a defensive copy
    return Consumes.clone(this.consumes);
  }

  setConsumes(newConsumes: Consumes) {
    if (Consumes.equals(this.consumes, newConsumes))
      return;

    // Make a defensive copy
    this.consumes = Consumes.clone(newConsumes);
    this.consumesChangeEmitter.emit();
  }

  equipItem(slot: ItemSlot, newItem: EquippedItem | null) {
    const newGear = this.gear.withEquippedItem(slot, newItem);
    if (newGear.equals(this.gear))
      return;

    this.gear = newGear;
    this.gearChangeEmitter.emit();
  }

  getEquippedItem(slot: ItemSlot): EquippedItem | null {
    return this.gear.getEquippedItem(slot);
  }

  getGear(): Gear {
    return this.gear;
  }

  setGear(newGear: Gear) {
    if (newGear.equals(this.gear))
      return;

    this.gear = newGear;
    this.gearChangeEmitter.emit();
  }

  getBonusStats(): Stats {
    return this.bonusStats;
  }

  setBonusStats(newBonusStats: Stats) {
    if (newBonusStats.equals(this.bonusStats))
      return;

    this.bonusStats = newBonusStats;
    this.bonusStatsChangeEmitter.emit();
  }

  getRotation(): SpecRotation<SpecType> {
    return this.specTypeFunctions.rotationCopy(this.rotation);
  }

  setRotation(newRotation: SpecRotation<SpecType>) {
    if (this.specTypeFunctions.rotationEquals(newRotation, this.rotation))
      return;

    this.rotation = this.specTypeFunctions.rotationCopy(newRotation);
    this.rotationChangeEmitter.emit();
  }

  getTalents(): SpecTalents<SpecType> {
    return this.specTypeFunctions.talentsCopy(this.talents);
  }

  setTalents(newTalents: SpecTalents<SpecType>) {
    if (this.specTypeFunctions.talentsEquals(newTalents, this.talents))
      return;

    this.talents = this.specTypeFunctions.talentsCopy(newTalents);
    this.talentsChangeEmitter.emit();
  }

  getTalentsString(): string {
    return this.talentsString;
  }

  setTalentsString(newTalentsString: string) {
    if (newTalentsString == this.talentsString)
      return;

    this.talentsString = newTalentsString;
    this.talentsStringChangeEmitter.emit();
  }

  getSpecOptions(): SpecOptions<SpecType> {
    return this.specTypeFunctions.optionsCopy(this.specOptions);
  }

  setSpecOptions(newSpecOptions: SpecOptions<SpecType>) {
    if (this.specTypeFunctions.optionsEquals(newSpecOptions, this.specOptions))
      return;

    this.specOptions = this.specTypeFunctions.optionsCopy(newSpecOptions);
    this.specOptionsChangeEmitter.emit();
  }

	computeGemEP(gem: Gem): number {
		const epFromStats = new Stats(gem.stats).computeEP(this.epWeights);
		const epFromEffect = getMetaGemEffectEP(this.spec, gem, new Stats(this.currentStats.finalStats));
		return epFromStats + epFromEffect;
	}

	computeEnchantEP(enchant: Enchant): number {
		return new Stats(enchant.stats).computeEP(this.epWeights);
	}

	computeItemEP(item: Item): number {
		if (item == null)
			return 0;

		let ep = new Stats(item.stats).computeEP(this.epWeights);

		const slot = getEligibleItemSlots(item)[0];
		const enchants = this.sim.getEnchants(slot);
		if (enchants.length > 0) {
			ep += Math.max(...enchants.map(enchant => this.computeEnchantEP(enchant)));
		}

		// Compare whether its better to match sockets + get socket bonus, or just use best gems.
		const bestGemEPNotMatchingSockets = sum(item.gemSockets.map(socketColor => {
			const gems = this.sim.getGems(socketColor).filter(gem => !gem.unique && gem.phase <= this.sim.getPhase());
			if (gems.length > 0) {
				return Math.max(...gems.map(gem => this.computeGemEP(gem)));
			} else {
				return 0;
			}
		}));

		const bestGemEPMatchingSockets = sum(item.gemSockets.map(socketColor => {
			const gems = this.sim.getGems(socketColor).filter(gem => !gem.unique && gem.phase <= this.sim.getPhase() && gemMatchesSocket(gem, socketColor));
			if (gems.length > 0) {
				return Math.max(...gems.map(gem => this.computeGemEP(gem)));
			} else {
				return 0;
			}
		})) + new Stats(item.socketBonus).computeEP(this.epWeights);

		ep += Math.max(bestGemEPMatchingSockets, bestGemEPNotMatchingSockets);

		return ep;
	}

  setWowheadData(equippedItem: EquippedItem, elem: HTMLElement) {
    let parts = [];
    if (equippedItem.gems.length > 0) {
      parts.push('gems=' + equippedItem.gems.map(gem => gem ? gem.id : 0).join(':'));
    }
    if (equippedItem.enchant != null) {
      parts.push('ench=' + equippedItem.enchant.effectId);
    }
    parts.push('pcs=' + this.gear.asArray().filter(ei => ei != null).map(ei => ei!.item.id).join(':'));

    elem.setAttribute('data-wowhead', parts.join('&'));
  }

	toProto(): PlayerProto {
    return withSpecProto(
				PlayerProto.create({
					race: this.getRace(),
					class: specToClass[this.spec],
					equipment: this.getGear().asSpec(),
					consumes: this.getConsumes(),
					bonusStats: this.getBonusStats().asArray(),
				}),
				this.getRotation(),
				this.getTalents(),
				this.getSpecOptions());
	}

	// TODO: Remove to/from json functions and use proto versions instead. This will require
	// some way to store all talents in the proto.
  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'consumes': Consumes.toJson(this.consumes),
      'bonusStats': this.bonusStats.toJson(),
      'gear': EquipmentSpec.toJson(this.gear.asSpec()),
      'race': this.race,
      'rotation': this.specTypeFunctions.rotationToJson(this.rotation),
      'talents': this.talentsString,
      'specOptions': this.specTypeFunctions.optionsToJson(this.specOptions),
    };
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		try {
			this.setConsumes(Consumes.fromJson(obj['consumes']));
		} catch (e) {
			console.warn('Failed to parse consumes: ' + e);
		}

		try {
			this.setBonusStats(Stats.fromJson(obj['bonusStats']));
		} catch (e) {
			console.warn('Failed to parse bonus stats: ' + e);
		}

		try {
			this.setGear(this.sim.lookupEquipmentSpec(EquipmentSpec.fromJson(obj['gear'])));
		} catch (e) {
			console.warn('Failed to parse gear: ' + e);
		}

		try {
			this.setRace(obj['race']);
		} catch (e) {
			console.warn('Failed to parse race: ' + e);
		}

		try {
			this.setRotation(this.specTypeFunctions.rotationFromJson(obj['rotation']));
		} catch (e) {
			console.warn('Failed to parse rotation: ' + e);
		}

		try {
			this.setTalentsString(obj['talents']);
		} catch (e) {
			console.warn('Failed to parse talents: ' + e);
		}

		try {
			this.setSpecOptions(this.specTypeFunctions.optionsFromJson(obj['specOptions']));
		} catch (e) {
			console.warn('Failed to parse spec options: ' + e);
		}
  }
}
