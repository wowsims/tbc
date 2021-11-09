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
import { ComputeStatsRequest, ComputeStatsResult } from '/tbc/core/proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from '/tbc/core/proto/api.js';

import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { makeComputeStatsRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { SpecTalents } from '/tbc/core/proto_utils/utils.js';
import { SpecTypeFunctions } from '/tbc/core/proto_utils/utils.js';
import { specTypeFunctions } from '/tbc/core/proto_utils/utils.js';
import { SpecOptions } from '/tbc/core/proto_utils/utils.js';
import { canEquipItem } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { getEligibleItemSlots } from '/tbc/core/proto_utils/utils.js';
import { getEligibleEnchantSlots } from '/tbc/core/proto_utils/utils.js';
import { gemEligibleForSocket } from '/tbc/core/proto_utils/utils.js';
import { gemMatchesSocket } from '/tbc/core/proto_utils/utils.js';

import { Listener } from './typed_event.js';
import { TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';
import { sum } from './utils.js';
import { wait } from './utils.js';
import { WorkerPool } from './worker_pool.js';

export interface PlayerConfig<SpecType extends Spec> {
  spec: Spec;
  epStats: Array<Stat>;
  epReferenceStat: Stat;
  displayStats: Array<Stat>;
  defaults: {
		gear: EquipmentSpec,
		epWeights: Stats,
    consumes: Consumes,
    rotation: SpecRotation<SpecType>,
    talents: string,
    specOptions: SpecOptions<SpecType>,
  },
	metaGemEffectEP?: ((gem: Gem, player: Player<SpecType>) => number),
}

// Manages all the gear / consumes / other settings for a single Player.
export class Player<SpecType extends Spec> {
  readonly spec: Spec;

  readonly phaseChangeEmitter = new TypedEvent<void>();
  readonly consumesChangeEmitter = new TypedEvent<void>();
  readonly customStatsChangeEmitter = new TypedEvent<void>();
  readonly gearChangeEmitter = new TypedEvent<void>();
  readonly raceChangeEmitter = new TypedEvent<void>();
  readonly rotationChangeEmitter = new TypedEvent<void>();
  readonly talentsChangeEmitter = new TypedEvent<void>();
  // Talents dont have all fields so we need this
  readonly talentsStringChangeEmitter = new TypedEvent<void>();
  readonly specOptionsChangeEmitter = new TypedEvent<void>();

  // Emits when any of the above emitters emit.
  readonly changeEmitter = new TypedEvent<void>();

  readonly characterStatsEmitter = new TypedEvent<void>();
	private _currentStats: ComputeStatsResult;

  // Current values
  private _consumes: Consumes;
  private _customStats: Stats;
  private _gear: Gear;
  private _race: Race;
  private _rotation: SpecRotation<SpecType>;
  private _talents: SpecTalents<SpecType>;
  private _talentsString: string;
  private _specOptions: SpecOptions<SpecType>;
	private _epWeights: Stats;

  readonly specTypeFunctions: SpecTypeFunctions<SpecType>;
	private readonly _metaGemEffectEP: (gem: Gem, player: Player<SpecType>) => number;

	readonly defaultGear: EquipmentSpec;

	readonly sim: Sim;

  constructor(config: PlayerConfig<SpecType>, sim: Sim) {
		this.sim = sim;

    this.spec = config.spec;
    this._race = specToEligibleRaces[this.spec][0];

    this.specTypeFunctions = specTypeFunctions[this.spec] as SpecTypeFunctions<SpecType>;
		this._metaGemEffectEP = config.metaGemEffectEP || (() => 0);

    this._consumes = config.defaults.consumes;
    this._customStats = new Stats();
    this._gear = new Gear({});
    this._rotation = config.defaults.rotation;
    this._talents = this.specTypeFunctions.talentsCreate();
    this._talentsString = config.defaults.talents;
		this._epWeights = config.defaults.epWeights;
    this._specOptions = config.defaults.specOptions;
		this.defaultGear = config.defaults.gear;

    [
      this.consumesChangeEmitter,
      this.customStatsChangeEmitter,
      this.gearChangeEmitter,
      this.raceChangeEmitter,
      this.rotationChangeEmitter,
      this.talentsChangeEmitter,
      this.talentsStringChangeEmitter,
      this.specOptionsChangeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));

		this._currentStats = ComputeStatsResult.create();
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

  async statWeights(request: StatWeightsRequest): Promise<StatWeightsResult> {
		const result = await this.sim.statWeights(request);
		this._epWeights = new Stats(result.epValues);
		return result;
	}

	// This should be invoked internally whenever stats might have changed.
	private async updateCharacterStats() {
		// Sometimes a ui change triggers other changes, so waiting a bit makes sure
		// we get all of them.
		await wait(10);

		const computeStatsResult = await this.sim.computeStats(makeComputeStatsRequest(
      this.sim.getRaidBuffs(),
      this.sim.getPartyBuffs(),
      this._consumes,
      this._customStats,
      this.sim.getEncounter(),
      this._gear,
      this._race,
      this._rotation,
      this._talents,
      this._specOptions));

		this._currentStats = computeStatsResult;
		this.characterStatsEmitter.emit();
	}

	getCurrentStats(): ComputeStatsResult {
		return ComputeStatsResult.clone(this._currentStats);
	}
  
  getRace(): Race {
    return this._race;
  }
  setRace(newRace: Race) {
    if (newRace != this._race) {
      this._race = newRace;
      this.raceChangeEmitter.emit();
    }
  }

  getConsumes(): Consumes {
    // Make a defensive copy
    return Consumes.clone(this._consumes);
  }

  setConsumes(newConsumes: Consumes) {
    if (Consumes.equals(this._consumes, newConsumes))
      return;

    // Make a defensive copy
    this._consumes = Consumes.clone(newConsumes);
    this.consumesChangeEmitter.emit();
  }

  equipItem(slot: ItemSlot, newItem: EquippedItem | null) {
    const newGear = this._gear.withEquippedItem(slot, newItem);
    if (newGear.equals(this._gear))
      return;

    this._gear = newGear;
    this.gearChangeEmitter.emit();
  }

  getEquippedItem(slot: ItemSlot): EquippedItem | null {
    return this._gear.getEquippedItem(slot);
  }

  getGear(): Gear {
    return this._gear;
  }

  setGear(newGear: Gear) {
    if (newGear.equals(this._gear))
      return;

    this._gear = newGear;
    this.gearChangeEmitter.emit();
  }

  getCustomStats(): Stats {
    return this._customStats;
  }

  setCustomStats(newCustomStats: Stats) {
    if (newCustomStats.equals(this._customStats))
      return;

    this._customStats = newCustomStats;
    this.customStatsChangeEmitter.emit();
  }

  getRotation(): SpecRotation<SpecType> {
    return this.specTypeFunctions.rotationCopy(this._rotation);
  }

  setRotation(newRotation: SpecRotation<SpecType>) {
    if (this.specTypeFunctions.rotationEquals(newRotation, this._rotation))
      return;

    this._rotation = this.specTypeFunctions.rotationCopy(newRotation);
    this.rotationChangeEmitter.emit();
  }

  getTalents(): SpecTalents<SpecType> {
    return this.specTypeFunctions.talentsCopy(this._talents);
  }

  setTalents(newTalents: SpecTalents<SpecType>) {
    if (this.specTypeFunctions.talentsEquals(newTalents, this._talents))
      return;

    this._talents = this.specTypeFunctions.talentsCopy(newTalents);
    this.talentsChangeEmitter.emit();
  }

  getTalentsString(): string {
    return this._talentsString;
  }

  setTalentsString(newTalentsString: string) {
    if (newTalentsString == this._talentsString)
      return;

    this._talentsString = newTalentsString;
    this.talentsStringChangeEmitter.emit();
  }

  getSpecOptions(): SpecOptions<SpecType> {
    return this.specTypeFunctions.optionsCopy(this._specOptions);
  }

  setSpecOptions(newSpecOptions: SpecOptions<SpecType>) {
    if (this.specTypeFunctions.optionsEquals(newSpecOptions, this._specOptions))
      return;

    this._specOptions = this.specTypeFunctions.optionsCopy(newSpecOptions);
    this.specOptionsChangeEmitter.emit();
  }

	computeGemEP(gem: Gem): number {
		const epFromStats = new Stats(gem.stats).computeEP(this._epWeights);
		const epFromEffect = this._metaGemEffectEP(gem, this);
		return epFromStats + epFromEffect;
	}

	computeEnchantEP(enchant: Enchant): number {
		return new Stats(enchant.stats).computeEP(this._epWeights);
	}

	computeItemEP(item: Item): number {
		if (item == null)
			return 0;

		let ep = new Stats(item.stats).computeEP(this._epWeights);

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
		})) + new Stats(item.socketBonus).computeEP(this._epWeights);

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
    parts.push('pcs=' + this._gear.asArray().filter(ei => ei != null).map(ei => ei!.item.id).join(':'));

    elem.setAttribute('data-wowhead', parts.join('&'));
  }

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'consumes': Consumes.toJson(this._consumes),
      'customStats': this._customStats.toJson(),
      'gear': EquipmentSpec.toJson(this._gear.asSpec()),
      'race': this._race,
      'rotation': this.specTypeFunctions.rotationToJson(this._rotation),
      'talents': this._talentsString,
      'specOptions': this.specTypeFunctions.optionsToJson(this._specOptions),
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
			this.setCustomStats(Stats.fromJson(obj['customStats']));
		} catch (e) {
			console.warn('Failed to parse custom stats: ' + e);
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
