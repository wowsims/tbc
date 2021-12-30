import { Class } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Enchant } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Gem } from '/tbc/core/proto/common.js';
import { GemColor } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { ItemQuality } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { ItemType } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { PlayerStats } from '/tbc/core/proto/api.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { StatWeightsResult } from '/tbc/core/proto/api.js';

import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import {
	gemEligibleForSocket,
	gemMatchesSocket,
} from '/tbc/core/proto_utils/gems.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';

import {
	Faction,
	SpecRotation,
	SpecTalents,
	SpecTypeFunctions,
	SpecOptions,
	canEquipItem,
	classColors,
	getEligibleEnchantSlots,
	getEligibleItemSlots,
	getTalentTreeIcon,
	getMetaGemEffectEP,
	raceToFaction,
	specToClass,
	specToEligibleRaces,
	specTypeFunctions,
	withSpecProto,
} from '/tbc/core/proto_utils/utils.js';

import { Listener } from './typed_event.js';
import { EventID, TypedEvent } from './typed_event.js';
import { Party, MAX_PARTY_SIZE } from './party.js';
import { Raid } from './raid.js';
import { Sim } from './sim.js';
import { sum } from './utils.js';
import { wait } from './utils.js';
import { WorkerPool } from './worker_pool.js';

// Manages all the gear / consumes / other settings for a single Player.
export class Player<SpecType extends Spec> {
	readonly sim: Sim;
	private party: Party | null;
	private raid: Raid | null;

  readonly spec: Spec;
	private name: string = '';
  private buffs: IndividualBuffs = IndividualBuffs.create();
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
	private currentStats: PlayerStats = PlayerStats.create();

  readonly nameChangeEmitter = new TypedEvent<void>('PlayerName');
  readonly buffsChangeEmitter = new TypedEvent<void>('PlayerBuffs');
  readonly consumesChangeEmitter = new TypedEvent<void>('PlayerConsumes');
  readonly bonusStatsChangeEmitter = new TypedEvent<void>('PlayerBonusStats');
  readonly gearChangeEmitter = new TypedEvent<void>('PlayerGear');
  readonly raceChangeEmitter = new TypedEvent<void>('PlayerRace');
  readonly rotationChangeEmitter = new TypedEvent<void>('PlayerRotation');
  readonly talentsChangeEmitter = new TypedEvent<void>('PlayerTalents');
  // Talents dont have all fields so we need this.
  readonly talentsStringChangeEmitter = new TypedEvent<void>('PlayerTalentsString');
  readonly specOptionsChangeEmitter = new TypedEvent<void>('PlayerSpecOptions');

  readonly currentStatsEmitter = new TypedEvent<void>('PlayerCurrentStats');

  // Emits when any of the above emitters emit.
  readonly changeEmitter: TypedEvent<void>;

  constructor(spec: Spec, sim: Sim) {
		this.sim = sim;
		this.party = null;
		this.raid = null;

    this.spec = spec;
    this.race = specToEligibleRaces[this.spec][0];
    this.specTypeFunctions = specTypeFunctions[this.spec] as SpecTypeFunctions<SpecType>;
		this.rotation = this.specTypeFunctions.rotationCreate();
    this.talents = this.specTypeFunctions.talentsCreate();
		this.specOptions = this.specTypeFunctions.optionsCreate();

		this.changeEmitter = TypedEvent.onAny([
      this.nameChangeEmitter,
      this.buffsChangeEmitter,
      this.consumesChangeEmitter,
      this.bonusStatsChangeEmitter,
      this.gearChangeEmitter,
      this.raceChangeEmitter,
      this.rotationChangeEmitter,
      this.talentsChangeEmitter,
      this.talentsStringChangeEmitter,
      this.specOptionsChangeEmitter,
		], 'PlayerChange');
  }

	getSpecIcon(): string {
		return getTalentTreeIcon(this.spec, this.getTalentsString());
	}

	getClass(): Class {
		return specToClass[this.spec];
	}

	getClassColor(): string {
		return classColors[this.getClass()];
	}

	getParty(): Party | null {
		return this.party;
	}

	getRaid(): Raid | null {
		return this.raid;
	}

	// Returns this player's index within its party [0-4].
	getPartyIndex(): number {
		if (this.party == null) {
			throw new Error('Can\'t get party index for player without a party!');
		}

		return this.party.getPlayers().indexOf(this);
	}

	// Returns this player's index within its raid [0-24].
	getRaidIndex(): number {
		if (this.party == null) {
			throw new Error('Can\'t get raid index for player without a party!');
		}

		return this.party.getIndex() * MAX_PARTY_SIZE + this.getPartyIndex();
	}

	// This should only ever be called from party.
	setParty(newParty: Party | null) {
		if (newParty == null) {
			this.party = null;
			this.raid = null;
		} else {
			this.party = newParty;
			this.raid = newParty.raid;
		}
	}

	getOtherPartyMembers(): Array<Player<any>> {
		if (this.party == null) {
			return [];
		}

		return this.party.getPlayers().filter(player => player != null && player != this) as Array<Player<any>>;
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

  async computeStatWeights(epStats: Array<Stat>, epReferenceStat: Stat): Promise<StatWeightsResult> {
		const result = await this.sim.statWeights(this, epStats, epReferenceStat);
		this.epWeights = new Stats(result.epValues);
		return result;
	}

	getCurrentStats(): PlayerStats {
		return PlayerStats.clone(this.currentStats);
	}

	setCurrentStats(eventID: EventID, newStats: PlayerStats) {
		this.currentStats = newStats;
		this.currentStatsEmitter.emit(eventID);
	}
  
  getName(): string {
    return this.name;
  }
  setName(eventID: EventID, newName: string) {
    if (newName != this.name) {
      this.name = newName;
      this.nameChangeEmitter.emit(eventID);
    }
  }

	getLabel(): string {
		if (this.party) {
			return `${this.name} (#${this.getRaidIndex() + 1})`;
		} else {
			return this.name;
		}
	}
  
  getRace(): Race {
    return this.race;
  }
  setRace(eventID: EventID, newRace: Race) {
    if (newRace != this.race) {
      this.race = newRace;
      this.raceChangeEmitter.emit(eventID);
    }
  }

	getFaction(): Faction {
		return raceToFaction[this.getRace()];
	}

  getBuffs(): IndividualBuffs {
    // Make a defensive copy
    return IndividualBuffs.clone(this.buffs);
  }

  setBuffs(eventID: EventID, newBuffs: IndividualBuffs) {
    if (IndividualBuffs.equals(this.buffs, newBuffs))
      return;

    // Make a defensive copy
    this.buffs = IndividualBuffs.clone(newBuffs);
    this.buffsChangeEmitter.emit(eventID);
  }

  getConsumes(): Consumes {
    // Make a defensive copy
    return Consumes.clone(this.consumes);
  }

  setConsumes(eventID: EventID, newConsumes: Consumes) {
    if (Consumes.equals(this.consumes, newConsumes))
      return;

    // Make a defensive copy
    this.consumes = Consumes.clone(newConsumes);
    this.consumesChangeEmitter.emit(eventID);
  }

  equipItem(eventID: EventID, slot: ItemSlot, newItem: EquippedItem | null) {
    const newGear = this.gear.withEquippedItem(slot, newItem);
    if (newGear.equals(this.gear))
      return;

    this.gear = newGear;
    this.gearChangeEmitter.emit(eventID);
  }

  getEquippedItem(slot: ItemSlot): EquippedItem | null {
    return this.gear.getEquippedItem(slot);
  }

  getGear(): Gear {
    return this.gear;
  }

  setGear(eventID: EventID, newGear: Gear) {
    if (newGear.equals(this.gear))
      return;

    this.gear = newGear;
    this.gearChangeEmitter.emit(eventID);
  }

  getBonusStats(): Stats {
    return this.bonusStats;
  }

  setBonusStats(eventID: EventID, newBonusStats: Stats) {
    if (newBonusStats.equals(this.bonusStats))
      return;

    this.bonusStats = newBonusStats;
    this.bonusStatsChangeEmitter.emit(eventID);
  }

  getRotation(): SpecRotation<SpecType> {
    return this.specTypeFunctions.rotationCopy(this.rotation);
  }

  setRotation(eventID: EventID, newRotation: SpecRotation<SpecType>) {
    if (this.specTypeFunctions.rotationEquals(newRotation, this.rotation))
      return;

    this.rotation = this.specTypeFunctions.rotationCopy(newRotation);
    this.rotationChangeEmitter.emit(eventID);
  }

  getTalents(): SpecTalents<SpecType> {
    return this.specTypeFunctions.talentsCopy(this.talents);
  }

  setTalents(eventID: EventID, newTalents: SpecTalents<SpecType>) {
    if (this.specTypeFunctions.talentsEquals(newTalents, this.talents))
      return;

    this.talents = this.specTypeFunctions.talentsCopy(newTalents);
    this.talentsChangeEmitter.emit(eventID);
  }

  getTalentsString(): string {
    return this.talentsString;
  }

  setTalentsString(eventID: EventID, newTalentsString: string) {
    if (newTalentsString == this.talentsString)
      return;

    this.talentsString = newTalentsString;
    this.talentsStringChangeEmitter.emit(eventID);
  }

	getTalentTreeIcon(): string {
		return getTalentTreeIcon(this.spec, this.getTalentsString());
	}

  getSpecOptions(): SpecOptions<SpecType> {
    return this.specTypeFunctions.optionsCopy(this.specOptions);
  }

  setSpecOptions(eventID: EventID, newSpecOptions: SpecOptions<SpecType>) {
    if (this.specTypeFunctions.optionsEquals(newSpecOptions, this.specOptions))
      return;

    this.specOptions = this.specTypeFunctions.optionsCopy(newSpecOptions);
    this.specOptionsChangeEmitter.emit(eventID);
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
				this.spec,
				PlayerProto.create({
					name: this.getName(),
					race: this.getRace(),
					class: this.getClass(),
					equipment: this.getGear().asSpec(),
					consumes: this.getConsumes(),
					bonusStats: this.getBonusStats().asArray(),
					buffs: this.getBuffs(),
					talentsString: this.getTalentsString(),
				}),
				this.getRotation(),
				this.getTalents(),
				this.getSpecOptions());
	}

	fromProto(eventID: EventID, proto: PlayerProto) {
		TypedEvent.freezeAllAndDo(() => {
			this.setName(eventID, proto.name);
			this.setRace(eventID, proto.race);
			this.setGear(eventID, proto.equipment ? this.sim.lookupEquipmentSpec(proto.equipment) : new Gear({}));
			this.setConsumes(eventID, proto.consumes || Consumes.create());
			this.setBonusStats(eventID, new Stats(proto.bonusStats));
			this.setBuffs(eventID, proto.buffs || IndividualBuffs.create());
			this.setTalentsString(eventID, proto.talentsString);
			this.setRotation(eventID, this.specTypeFunctions.rotationFromPlayer(proto));
			this.setTalents(eventID, this.specTypeFunctions.talentsFromPlayer(proto));
			this.setSpecOptions(eventID, this.specTypeFunctions.optionsFromPlayer(proto));
		});
	}

	// TODO: Remove to/from json functions and use proto versions instead. This will require
	// some way to store all talents in the proto.
  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'name': this.name,
      'buffs': IndividualBuffs.toJson(this.buffs),
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
  fromJson(eventID: EventID, obj: any) {
		TypedEvent.freezeAllAndDo(() => {
			try {
				if (obj['name']) {
					this.setName(eventID, obj['name']);
				}
			} catch (e) {
				console.warn('Failed to parse name: ' + e);
			}

			try {
				this.setBuffs(eventID, IndividualBuffs.fromJson(obj['buffs']));
			} catch (e) {
				console.warn('Failed to parse player buffs: ' + e);
			}

			try {
				this.setConsumes(eventID, Consumes.fromJson(obj['consumes']));
			} catch (e) {
				console.warn('Failed to parse consumes: ' + e);
			}

			// For legacy format. Do not remove this until 2022/01/02 (1 month).
			if (obj['customStats']) {
				obj['bonusStats'] = obj['customStats'];
			}

			try {
				this.setBonusStats(eventID, Stats.fromJson(obj['bonusStats']));
			} catch (e) {
				console.warn('Failed to parse bonus stats: ' + e);
			}

			try {
				this.setGear(eventID, this.sim.lookupEquipmentSpec(EquipmentSpec.fromJson(obj['gear'])));
			} catch (e) {
				console.warn('Failed to parse gear: ' + e);
			}

			try {
				this.setRace(eventID, obj['race']);
			} catch (e) {
				console.warn('Failed to parse race: ' + e);
			}

			try {
				this.setRotation(eventID, this.specTypeFunctions.rotationFromJson(obj['rotation']));
			} catch (e) {
				console.warn('Failed to parse rotation: ' + e);
			}

			try {
				this.setTalentsString(eventID, obj['talents']);
			} catch (e) {
				console.warn('Failed to parse talents: ' + e);
			}

			try {
				this.setSpecOptions(eventID, this.specTypeFunctions.optionsFromJson(obj['specOptions']));
			} catch (e) {
				console.warn('Failed to parse spec options: ' + e);
			}
		});
  }

	clone(eventID: EventID): Player<SpecType> {
		const newPlayer = new Player<SpecType>(this.spec, this.sim);
		newPlayer.fromProto(eventID, this.toProto());
		return newPlayer;
	}
}
