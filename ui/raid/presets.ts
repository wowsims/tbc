import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { Class } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { SpecOptions } from '/tbc/core/proto_utils/utils.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { specIconsLarge } from '/tbc/core/proto_utils/utils.js';
import { specNames } from '/tbc/core/proto_utils/utils.js';
import { talentTreeIcons } from '/tbc/core/proto_utils/utils.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { BuffBot } from './buff_bot.js';

import * as BalanceDruidPresets from '/tbc/balance_druid/presets.js';
import * as ElementalShamanPresets from '/tbc/elemental_shaman/presets.js';
import * as EnhancementShamanPresets from '/tbc/enhancement_shaman/presets.js';
import * as HunterPresets from '/tbc/hunter/presets.js';
import * as MagePresets from '/tbc/mage/presets.js';
import * as ShadowPriestPresets from '/tbc/shadow_priest/presets.js';

import { BalanceDruidSimUI } from '/tbc/balance_druid/sim.js';
import { EnhancementShamanSimUI } from '/tbc/enhancement_shaman/sim.js';
import { ElementalShamanSimUI } from '/tbc/elemental_shaman/sim.js';
import { HunterSimUI } from '/tbc/hunter/sim.js';
import { MageSimUI } from '/tbc/mage/sim.js';
import { ShadowPriestSimUI } from '/tbc/shadow_priest/sim.js';

export const specSimFactories: Partial<Record<Spec, (parentElem: HTMLElement, player: Player<any>) => IndividualSimUI<any>>> = {
	[Spec.SpecBalanceDruid]: (parentElem: HTMLElement, player: Player<any>) => new BalanceDruidSimUI(parentElem, player),
	[Spec.SpecElementalShaman]: (parentElem: HTMLElement, player: Player<any>) => new ElementalShamanSimUI(parentElem, player),
	[Spec.SpecEnhancementShaman]: (parentElem: HTMLElement, player: Player<any>) => new EnhancementShamanSimUI(parentElem, player),
	[Spec.SpecHunter]: (parentElem: HTMLElement, player: Player<any>) => new HunterSimUI(parentElem, player),
	[Spec.SpecMage]: (parentElem: HTMLElement, player: Player<any>) => new MageSimUI(parentElem, player),
	[Spec.SpecShadowPriest]: (parentElem: HTMLElement, player: Player<any>) => new ShadowPriestSimUI(parentElem, player),
};

// Configuration necessary for creating new players.
export interface PresetSpecSettings<SpecType extends Spec> {
	spec: Spec,
	rotation: SpecRotation<SpecType>,
	talents: string,
	specOptions: SpecOptions<SpecType>,
	consumes: Consumes,

	defaultName: string,
	defaultFactionRaces: Record<Faction, Race>,
	defaultGear: Record<Faction, Record<number, EquipmentSpec>>,

	tooltip: string,
	iconUrl: string,
}

// Configuration necessary for creating new BuffBots.
export interface BuffBotSettings {
	// The value of this field must never change, to preserve local storage data.
	buffBotId: string,

	// Set this to true to remove a buff bot option after launching a real sim.
	// This will allow users with saved settings to properly load the buffbot but
	// also remove the buffbot as an option from the UI.
	deprecated?: boolean,

	spec: Spec,
	name: string,
	tooltip: string,
	iconUrl: string,

	// Callback to apply buffs from this buff bot.
	modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => void,
	modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => void,
}

export const playerPresets: Array<PresetSpecSettings<any>> = [
	{
		spec: Spec.SpecBalanceDruid,
		rotation: BalanceDruidPresets.DefaultRotation,
		talents: BalanceDruidPresets.StandardTalents.data,
		specOptions: BalanceDruidPresets.DefaultOptions,
		consumes: BalanceDruidPresets.DefaultConsumes,
		defaultName: 'Balance Druid',
		defaultFactionRaces: {
			[Faction.Unknown]: Race.RaceUnknown,
			[Faction.Alliance]: Race.RaceNightElf,
			[Faction.Horde]: Race.RaceTauren,
		},
		defaultGear: {
			[Faction.Unknown]: {},
			[Faction.Alliance]: {
				1: BalanceDruidPresets.P1_ALLIANCE_PRESET.gear,
				2: BalanceDruidPresets.P2_ALLIANCE_PRESET.gear,
				3: BalanceDruidPresets.P3_PRESET.gear,
			},
			[Faction.Horde]: {
				1: BalanceDruidPresets.P1_HORDE_PRESET.gear,
				2: BalanceDruidPresets.P2_HORDE_PRESET.gear,
				3: BalanceDruidPresets.P3_PRESET.gear,
			},
		},
		tooltip: specNames[Spec.SpecBalanceDruid],
		iconUrl: specIconsLarge[Spec.SpecBalanceDruid],
	},
	{
		spec: Spec.SpecHunter,
		rotation: HunterPresets.DefaultRotation,
		talents: HunterPresets.BeastMasteryTalents.data,
		specOptions: HunterPresets.DefaultOptions,
		consumes: HunterPresets.DefaultConsumes,
		defaultName: 'BM Hunter',
		defaultFactionRaces: {
			[Faction.Unknown]: Race.RaceUnknown,
			[Faction.Alliance]: Race.RaceNightElf,
			[Faction.Horde]: Race.RaceOrc,
		},
		defaultGear: {
			[Faction.Unknown]: {},
			[Faction.Alliance]: {
				1: HunterPresets.P1_BM_PRESET.gear,
				2: HunterPresets.P2_BM_PRESET.gear,
				3: HunterPresets.P3_BM_PRESET.gear,
			},
			[Faction.Horde]: {
				1: HunterPresets.P1_BM_PRESET.gear,
				2: HunterPresets.P2_BM_PRESET.gear,
				3: HunterPresets.P3_BM_PRESET.gear,
			},
		},
		tooltip: 'BM Hunter',
		iconUrl: talentTreeIcons[Class.ClassHunter][0],
	},
	{
		spec: Spec.SpecHunter,
		rotation: HunterPresets.DefaultRotation,
		talents: HunterPresets.SurvivalTalents.data,
		specOptions: HunterPresets.DefaultOptions,
		consumes: HunterPresets.DefaultConsumes,
		defaultName: 'SV Hunter',
		defaultFactionRaces: {
			[Faction.Unknown]: Race.RaceUnknown,
			[Faction.Alliance]: Race.RaceNightElf,
			[Faction.Horde]: Race.RaceOrc,
		},
		defaultGear: {
			[Faction.Unknown]: {},
			[Faction.Alliance]: {
				1: HunterPresets.P1_SV_PRESET.gear,
				2: HunterPresets.P2_SV_PRESET.gear,
				3: HunterPresets.P3_SV_PRESET.gear,
			},
			[Faction.Horde]: {
				1: HunterPresets.P1_SV_PRESET.gear,
				2: HunterPresets.P2_SV_PRESET.gear,
				3: HunterPresets.P3_SV_PRESET.gear,
			},
		},
		tooltip: 'SV Hunter',
		iconUrl: talentTreeIcons[Class.ClassHunter][2],
	},
	{
		spec: Spec.SpecMage,
		rotation: MagePresets.DefaultArcaneRotation,
		talents: MagePresets.ArcaneTalents.data,
		specOptions: MagePresets.DefaultArcaneOptions,
		consumes: MagePresets.DefaultArcaneConsumes,
		defaultName: 'Arcane Mage',
		defaultFactionRaces: {
			[Faction.Unknown]: Race.RaceUnknown,
			[Faction.Alliance]: Race.RaceGnome,
			[Faction.Horde]: Race.RaceTroll10,
		},
		defaultGear: {
			[Faction.Unknown]: {},
			[Faction.Alliance]: {
				1: MagePresets.P1_ARCANE_PRESET.gear,
				2: MagePresets.P2_ARCANE_PRESET.gear,
				3: MagePresets.P3_ARCANE_PRESET.gear,
			},
			[Faction.Horde]: {
				1: MagePresets.P1_ARCANE_PRESET.gear,
				2: MagePresets.P2_ARCANE_PRESET.gear,
				3: MagePresets.P3_ARCANE_PRESET.gear,
			},
		},
		tooltip: 'Arcane Mage',
		iconUrl: talentTreeIcons[Class.ClassMage][0],
	},
	{
		spec: Spec.SpecMage,
		rotation: MagePresets.DefaultFireRotation,
		talents: MagePresets.FireTalents.data,
		specOptions: MagePresets.DefaultFireOptions,
		consumes: MagePresets.DefaultFireConsumes,
		defaultName: 'Fire Mage',
		defaultFactionRaces: {
			[Faction.Unknown]: Race.RaceUnknown,
			[Faction.Alliance]: Race.RaceGnome,
			[Faction.Horde]: Race.RaceTroll10,
		},
		defaultGear: {
			[Faction.Unknown]: {},
			[Faction.Alliance]: {
				1: MagePresets.P1_FIRE_PRESET.gear,
				2: MagePresets.P2_FIRE_PRESET.gear,
				3: MagePresets.P3_FIRE_PRESET.gear,
			},
			[Faction.Horde]: {
				1: MagePresets.P1_FIRE_PRESET.gear,
				2: MagePresets.P2_FIRE_PRESET.gear,
				3: MagePresets.P3_FIRE_PRESET.gear,
			},
		},
		tooltip: 'Fire Mage',
		iconUrl: talentTreeIcons[Class.ClassMage][1],
	},
	{
		spec: Spec.SpecMage,
		rotation: MagePresets.DefaultFrostRotation,
		talents: MagePresets.DeepFrostTalents.data,
		specOptions: MagePresets.DefaultFrostOptions,
		consumes: MagePresets.DefaultFrostConsumes,
		defaultName: 'Frost Mage',
		defaultFactionRaces: {
			[Faction.Unknown]: Race.RaceUnknown,
			[Faction.Alliance]: Race.RaceGnome,
			[Faction.Horde]: Race.RaceTroll10,
		},
		defaultGear: {
			[Faction.Unknown]: {},
			[Faction.Alliance]: {
				1: MagePresets.P1_FROST_PRESET.gear,
				2: MagePresets.P2_FROST_PRESET.gear,
				3: MagePresets.P3_FROST_PRESET.gear,
			},
			[Faction.Horde]: {
				1: MagePresets.P1_FROST_PRESET.gear,
				2: MagePresets.P2_FROST_PRESET.gear,
				3: MagePresets.P3_FROST_PRESET.gear,
			},
		},
		tooltip: 'Frost Mage',
		iconUrl: talentTreeIcons[Class.ClassMage][2],
	},
	{
		spec: Spec.SpecElementalShaman,
		rotation: ElementalShamanPresets.DefaultRotation,
		talents: ElementalShamanPresets.StandardTalents.data,
		specOptions: ElementalShamanPresets.DefaultOptions,
		consumes: ElementalShamanPresets.DefaultConsumes,
		defaultName: 'Ele Shaman',
		defaultFactionRaces: {
			[Faction.Unknown]: Race.RaceUnknown,
			[Faction.Alliance]: Race.RaceDraenei,
			[Faction.Horde]: Race.RaceOrc,
		},
		defaultGear: {
			[Faction.Unknown]: {},
			[Faction.Alliance]: {
				1: ElementalShamanPresets.P1_PRESET.gear,
				2: ElementalShamanPresets.P2_PRESET.gear,
				3: ElementalShamanPresets.P3_PRESET.gear,
			},
			[Faction.Horde]: {
				1: ElementalShamanPresets.P1_PRESET.gear,
				2: ElementalShamanPresets.P2_PRESET.gear,
				3: ElementalShamanPresets.P3_PRESET.gear,
			},
		},
		tooltip: specNames[Spec.SpecElementalShaman],
		iconUrl: specIconsLarge[Spec.SpecElementalShaman],
	},
	{
		spec: Spec.SpecEnhancementShaman,
		rotation: EnhancementShamanPresets.DefaultRotation,
		talents: EnhancementShamanPresets.StandardTalents.data,
		specOptions: EnhancementShamanPresets.DefaultOptions,
		consumes: EnhancementShamanPresets.DefaultConsumes,
		defaultName: 'Enh Shaman',
		defaultFactionRaces: {
			[Faction.Unknown]: Race.RaceUnknown,
			[Faction.Alliance]: Race.RaceDraenei,
			[Faction.Horde]: Race.RaceOrc,
		},
		defaultGear: {
			[Faction.Unknown]: {},
			[Faction.Alliance]: {
				1: EnhancementShamanPresets.P1_PRESET.gear,
				2: EnhancementShamanPresets.P2_PRESET.gear,
				3: EnhancementShamanPresets.P3_PRESET.gear,
			},
			[Faction.Horde]: {
				1: EnhancementShamanPresets.P1_PRESET.gear,
				2: EnhancementShamanPresets.P2_PRESET.gear,
				3: EnhancementShamanPresets.P3_PRESET.gear,
			},
		},
		tooltip: specNames[Spec.SpecEnhancementShaman],
		iconUrl: specIconsLarge[Spec.SpecEnhancementShaman],
	},
	{
		spec: Spec.SpecShadowPriest,
		rotation: ShadowPriestPresets.DefaultRotation,
		talents: ShadowPriestPresets.StandardTalents.data,
		specOptions: ShadowPriestPresets.DefaultOptions,
		consumes: ShadowPriestPresets.DefaultConsumes,
		defaultName: 'Shadow Priest',
		defaultFactionRaces: {
			[Faction.Unknown]: Race.RaceUnknown,
			[Faction.Alliance]: Race.RaceDwarf,
			[Faction.Horde]: Race.RaceUndead,
		},
		defaultGear: {
			[Faction.Unknown]: {},
			[Faction.Alliance]: {
				1: ShadowPriestPresets.P1_PRESET.gear,
				2: ShadowPriestPresets.P2_PRESET.gear,
				3: ShadowPriestPresets.P3_PRESET.gear,
			},
			[Faction.Horde]: {
				1: ShadowPriestPresets.P1_PRESET.gear,
				2: ShadowPriestPresets.P2_PRESET.gear,
				3: ShadowPriestPresets.P3_PRESET.gear,
			},
		},
		tooltip: specNames[Spec.SpecShadowPriest],
		iconUrl: specIconsLarge[Spec.SpecShadowPriest],
	},
];

export const implementedSpecs: Array<Spec> = [...new Set(playerPresets.map(preset => preset.spec))];

export const buffBotPresets: Array<BuffBotSettings> = [
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'Resto Druid',
		spec: Spec.SpecBalanceDruid,
		name: 'Resto Druid',
		tooltip: 'Resto Druid: Adds Improved Gift of the Wild, and an Innervate.',
		iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_nature_healingtouch.jpg',
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
			raidProto.buffs!.giftOfTheWild = TristateEffect.TristateEffectImproved;

			const innervateIndex = buffBot.getInnervateAssignment().targetIndex;
			if (innervateIndex != NO_TARGET) {
				const partyIndex = Math.floor(innervateIndex / 5);
				const playerIndex = innervateIndex % 5;
				const playerProto = raidProto.parties[partyIndex].players[playerIndex];
				if (playerProto.buffs) {
					playerProto.buffs.innervates++;
				}
			}
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
		},
	},
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'Mage',
		deprecated: true,
		spec: Spec.SpecMage,
		name: 'Mage',
		tooltip: 'Mage: Adds Arcane Brilliance.',
		iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_arcaneintellect.jpg',
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
			raidProto.buffs!.arcaneBrilliance = true;
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
		},
	},
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'Paladin',
		spec: Spec.SpecRetributionPaladin,
		name: 'Holy Paladin',
		tooltip: 'Holy Paladin: Adds a set of blessings.',
		iconUrl: talentTreeIcons[Class.ClassPaladin][0],
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
			// Do nothing, blessings are handled elswhere.
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
		},
	},
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'JoW Paladin',
		spec: Spec.SpecRetributionPaladin,
		name: 'JoW Paladin',
		tooltip: 'JoW Paladin: Adds a set of blessings and Judgement of Wisdom.',
		iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_righteousnessaura.jpg',
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
			// Do nothing, blessings are handled elswhere.
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
			encounterProto.targets[0].debuffs!.judgementOfWisdom = true;
		},
	},
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'JoC Paladin',
		spec: Spec.SpecRetributionPaladin,
		name: 'JoC Paladin',
		tooltip: 'JoC Paladin: Adds a set of blessings and Improved Judgement of the Crusader (+3% crit).',
		iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_holysmite.jpg',
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
			// Do nothing, blessings are handled elswhere.
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
			encounterProto.targets[0].debuffs!.improvedSealOfTheCrusader = true;
		},
	},
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'Holy Priest',
		spec: Spec.SpecShadowPriest,
		name: 'Holy Priest',
		tooltip: 'Holy Priest: Doesn\'t contribute to DPS, just fills a raid slot.',
		iconUrl: talentTreeIcons[Class.ClassPriest][1],
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
		},
	},
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'Divine Spirit Priest',
		spec: Spec.SpecShadowPriest,
		name: 'Disc Priest',
		tooltip: 'Disc Priest: Adds Improved Divine Spirit and a Power Infusion.',
		iconUrl: 'https://wow.zamimg.com/images/wow/icons/medium/spell_holy_powerinfusion.jpg',
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
			raidProto.buffs!.divineSpirit = TristateEffect.TristateEffectImproved;

			const powerInfusionIndex = buffBot.getPowerInfusionAssignment().targetIndex;
			if (powerInfusionIndex != NO_TARGET) {
				const partyIndex = Math.floor(powerInfusionIndex / 5);
				const playerIndex = powerInfusionIndex % 5;
				const playerProto = raidProto.parties[partyIndex].players[playerIndex];
				if (playerProto.buffs) {
					playerProto.buffs.powerInfusions++;
				}
			}
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
		},
	},
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'Resto Shaman',
		spec: Spec.SpecElementalShaman,
		name: 'Resto Shaman',
		tooltip: 'Resto Shaman: Adds Bloodlust, Mana Spring Totem, Wrath of Air Totem, Mana Tide Totem, and Drums of Battle.',
		iconUrl: talentTreeIcons[Class.ClassShaman][2],
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
			partyProto.buffs!.bloodlust++;
			partyProto.buffs!.manaSpringTotem = TristateEffect.TristateEffectImproved;
			partyProto.buffs!.wrathOfAirTotem = Math.max(partyProto.buffs!.wrathOfAirTotem, TristateEffect.TristateEffectRegular);
			partyProto.buffs!.manaTideTotems++;
			partyProto.buffs!.drums = Drums.DrumsOfBattle;
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
		},
	},
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'CoE Warlock',
		spec: Spec.SpecWarlock,
		name: 'CoE Warlock',
		tooltip: 'CoE Warlock: Adds Curse of Elements (regular). Also adds +20% uptime to ISB.',
		iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_shadow_chilltouch.jpg',
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
			const debuffs = encounterProto.targets[0].debuffs!;
			debuffs.curseOfElements = Math.max(debuffs.curseOfElements, TristateEffect.TristateEffectRegular);
			debuffs.isbUptime = Math.min(1.0, debuffs.isbUptime + 0.2);
		},
	},
	{
		// The value of this field must never change, to preserve local storage data.
		buffBotId: 'Malediction Warlock',
		spec: Spec.SpecWarlock,
		name: 'Aff Warlock',
		tooltip: 'Afflication Warlock: Adds Curse of Elements (improved). Also adds +20% uptime to ISB.',
		iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_shadow_curseofachimonde.jpg',
		modifyRaidProto: (buffBot: BuffBot, raidProto: RaidProto, partyProto: PartyProto) => {
		},
		modifyEncounterProto: (buffBot: BuffBot, encounterProto: EncounterProto) => {
			const debuffs = encounterProto.targets[0].debuffs!;
			debuffs.curseOfElements = TristateEffect.TristateEffectImproved;
			debuffs.isbUptime = Math.min(1.0, debuffs.isbUptime + 0.2);
		},
	},
];