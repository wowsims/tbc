import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';
import { specIconsLarge } from '/tbc/core/proto_utils/utils.js';
import { specNames } from '/tbc/core/proto_utils/utils.js';

import * as BalanceDruidPresets from '/tbc/balance_druid/presets.js';
import * as ElementalShamanPresets from '/tbc/elemental_shaman/presets.js';
import * as ShadowPriestPresets from '/tbc/shadow_priest/presets.js';

import { RaidSimUI } from './raid_sim_ui.js';

const ui = new RaidSimUI(document.body, {
	presets: [
		{
			spec: Spec.SpecBalanceDruid,
			rotation: BalanceDruidPresets.DefaultRotation,
			talents: BalanceDruidPresets.StandardTalents.data,
			specOptions: BalanceDruidPresets.DefaultOptions,
			consumes: BalanceDruidPresets.DefaultConsumes,
			defaultName: specNames[Spec.SpecBalanceDruid],
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceNightElf,
				[Faction.Horde]: Race.RaceTauren,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: BalanceDruidPresets.P1_ALLIANCE_BIS.gear,
					2: BalanceDruidPresets.P2_ALLIANCE_BIS.gear,
				},
				[Faction.Horde]: {
					1: BalanceDruidPresets.P1_HORDE_BIS.gear,
					2: BalanceDruidPresets.P2_HORDE_BIS.gear,
				},
			},
			tooltip: specNames[Spec.SpecBalanceDruid],
			iconUrl: specIconsLarge[Spec.SpecBalanceDruid],
		},
		{
			spec: Spec.SpecElementalShaman,
			rotation: ElementalShamanPresets.DefaultRotation,
			talents: ElementalShamanPresets.StandardTalents.data,
			specOptions: ElementalShamanPresets.DefaultOptions,
			consumes: ElementalShamanPresets.DefaultConsumes,
			defaultName: specNames[Spec.SpecElementalShaman],
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceDraenei,
				[Faction.Horde]: Race.RaceOrc,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: ElementalShamanPresets.P1_BIS.gear,
					2: ElementalShamanPresets.P2_BIS.gear,
				},
				[Faction.Horde]: {
					1: ElementalShamanPresets.P1_BIS.gear,
					2: ElementalShamanPresets.P2_BIS.gear,
				},
			},
			tooltip: specNames[Spec.SpecElementalShaman],
			iconUrl: specIconsLarge[Spec.SpecElementalShaman],
		},
		{
			spec: Spec.SpecShadowPriest,
			rotation: ShadowPriestPresets.DefaultRotation,
			talents: ShadowPriestPresets.StandardTalents.data,
			specOptions: ShadowPriestPresets.DefaultOptions,
			consumes: ShadowPriestPresets.DefaultConsumes,
			defaultName: specNames[Spec.SpecShadowPriest],
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceDwarf,
				[Faction.Horde]: Race.RaceUndead,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: ShadowPriestPresets.P1_BIS.gear,
					2: ShadowPriestPresets.P2_BIS.gear,
				},
				[Faction.Horde]: {
					1: ShadowPriestPresets.P1_BIS.gear,
					2: ShadowPriestPresets.P2_BIS.gear,
				},
			},
			tooltip: specNames[Spec.SpecShadowPriest],
			iconUrl: specIconsLarge[Spec.SpecShadowPriest],
		},
	],
	buffBots: [
		{
			// The value of this field must never change, to preserve local storage data.
			buffBotId: 'Mage',
			spec: Spec.SpecMage,
			name: 'Mage',
			tooltip: 'Adds Arcane Brilliance.',
			iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_arcaneintellect.jpg',
			modifyRaidProto: (raidProto: RaidProto, partyProto: PartyProto) => {
				raidProto.buffs!.arcaneBrilliance = true;
			},
			modifyEncounterProto: (encounterProto: EncounterProto) => {
			},
		},
		{
			// The value of this field must never change, to preserve local storage data.
			buffBotId: 'Paladin',
			spec: Spec.SpecRetributionPaladin,
			name: 'Paladin',
			tooltip: 'Adds a set of blessings.',
			iconUrl: specIconsLarge[Spec.SpecRetributionPaladin],
			modifyRaidProto: (raidProto: RaidProto, partyProto: PartyProto) => {
				// Do nothing, blessings are handled elswhere.
			},
			modifyEncounterProto: (encounterProto: EncounterProto) => {
			},
		},
		{
			// The value of this field must never change, to preserve local storage data.
			buffBotId: 'JoW Paladin',
			spec: Spec.SpecRetributionPaladin,
			name: 'JoW Paladin',
			tooltip: 'Adds a set of blessings and Judgement of Wisdom.',
			iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_righteousnessaura.jpg',
			modifyRaidProto: (raidProto: RaidProto, partyProto: PartyProto) => {
				// Do nothing, blessings are handled elswhere.
			},
			modifyEncounterProto: (encounterProto: EncounterProto) => {
				encounterProto.targets[0].debuffs!.judgementOfWisdom = true;
			},
		},
		{
			// The value of this field must never change, to preserve local storage data.
			buffBotId: 'JoC Paladin',
			spec: Spec.SpecRetributionPaladin,
			name: 'JoC Paladin',
			tooltip: 'Adds a set of blessings and Improved Judgement of the Crusader (+3% crit).',
			iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_holysmite.jpg',
			modifyRaidProto: (raidProto: RaidProto, partyProto: PartyProto) => {
				// Do nothing, blessings are handled elswhere.
			},
			modifyEncounterProto: (encounterProto: EncounterProto) => {
				encounterProto.targets[0].debuffs!.improvedSealOfTheCrusader = true;
			},
		},
		{
			// The value of this field must never change, to preserve local storage data.
			buffBotId: 'Divine Spirit Priest',
			spec: Spec.SpecShadowPriest,
			name: 'Holy Priest',
			tooltip: 'Adds Improved Divine Spirit',
			iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_divinespirit.jpg',
			modifyRaidProto: (raidProto: RaidProto, partyProto: PartyProto) => {
				raidProto.buffs!.divineSpirit = TristateEffect.TristateEffectImproved;
			},
			modifyEncounterProto: (encounterProto: EncounterProto) => {
			},
		},
		{
			// The value of this field must never change, to preserve local storage data.
			buffBotId: 'CoE Warlock',
			spec: Spec.SpecWarlock,
			name: 'CoE Warlock',
			tooltip: 'Adds Curse of Elements (regular). Also adds +20% uptime to ISB.',
			iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_shadow_chilltouch.jpg',
			modifyRaidProto: (raidProto: RaidProto, partyProto: PartyProto) => {
			},
			modifyEncounterProto: (encounterProto: EncounterProto) => {
				const debuffs = encounterProto.targets[0].debuffs!;
				debuffs.curseOfElements = Math.max(debuffs.curseOfElements, TristateEffect.TristateEffectRegular);
				debuffs.isbUptime = Math.min(1.0, debuffs.isbUptime + 0.2);
			},
		},
		{
			// The value of this field must never change, to preserve local storage data.
			buffBotId: 'Malediction Warlock',
			spec: Spec.SpecWarlock,
			name: 'Malediction Warlock',
			tooltip: 'Adds Curse of Elements (improved). Also adds +20% uptime to ISB.',
			iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_shadow_curseofachimonde.jpg',
			modifyRaidProto: (raidProto: RaidProto, partyProto: PartyProto) => {
			},
			modifyEncounterProto: (encounterProto: EncounterProto) => {
				const debuffs = encounterProto.targets[0].debuffs!;
				debuffs.curseOfElements = TristateEffect.TristateEffectImproved;
				debuffs.isbUptime = Math.min(1.0, debuffs.isbUptime + 0.2);
			},
		},
	],
});
