import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js'
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';

import { ShadowPriest, ShadowPriest_Rotation as Rotation, ShadowPriest_Options as Options, ShadowPriest_Rotation, ShadowPriest_Rotation_RotationType } from '/tbc/core/proto/priest.js';

import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as Gems from '/tbc/core/constants/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';

import * as ShadowPriestInputs from './inputs.js';
import * as Presets from './presets.js';

export class ShadowPriestSimUI extends IndividualSimUI<Spec.SpecShadowPriest> {
  constructor(parentElem: HTMLElement, player: Player<Spec.SpecShadowPriest>) {
		super(parentElem, player, {
			// Can be 'Alpha', 'Beta', or 'Live'. Just adds a postfix to the generated title.
			releaseStatus: 'Alpha',
			cssClass: 'shadow-priest-sim-ui',
			// List any known bugs / issues here and they'll be shown on the site.
			knownIssues: [
			],

			// All stats for which EP should be calculated.
			epStats: [
				Stat.StatIntellect,
				Stat.StatSpirit,
				Stat.StatSpellPower,
				Stat.StatShadowSpellPower,
				Stat.StatSpellHit,
				Stat.StatSpellCrit,
				Stat.StatSpellHaste,
				Stat.StatMP5,
			],
			// Reference stat against which to calculate EP. I think all classes use either spell power or attack power.
			epReferenceStat: Stat.StatSpellPower,
			// Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
			displayStats: [
				Stat.StatStamina,
				Stat.StatIntellect,
				Stat.StatSpirit,
				Stat.StatSpellPower,
				Stat.StatShadowSpellPower,
				Stat.StatSpellHit,
				Stat.StatSpellCrit,
				Stat.StatSpellHaste,
				Stat.StatMP5,
			],

			defaults: {
				// Default equipped gear.
				gear: Presets.P3_BIS.gear,
				// Default EP weights for sorting gear in the gear picker.
				epWeights: Stats.fromMap({
					[Stat.StatIntellect]: 0.05,
					[Stat.StatSpirit]: 0.11,
					[Stat.StatSpellPower]: 1,
					[Stat.StatShadowSpellPower]: 1,
					[Stat.StatSpellCrit]: 0.163,
					[Stat.StatSpellHaste]: 1.0, // tricky because SP is tricky
					[Stat.StatMP5]: 0.00,
				}),
				// Default consumes settings.
				consumes: Presets.DefaultConsumes,
				// Default rotation settings.
				rotation: Presets.DefaultRotation,
				// Default talents.
				talents: Presets.StandardTalents.data,
				// Default spec-specific settings.
				specOptions: Presets.DefaultOptions,
				// Default raid/party buffs settings.
				raidBuffs: RaidBuffs.create({
					arcaneBrilliance: true,
					divineSpirit: TristateEffect.TristateEffectImproved,
					giftOfTheWild: TristateEffect.TristateEffectImproved,
				}),
				partyBuffs: PartyBuffs.create({
					bloodlust: 1,
					manaSpringTotem: TristateEffect.TristateEffectRegular,
					totemOfWrath: 1,
					wrathOfAirTotem: TristateEffect.TristateEffectRegular,
				}),
				individualBuffs: IndividualBuffs.create({
					blessingOfKings: true,
					blessingOfWisdom: 2,
				}),
				debuffs: Debuffs.create({
					judgementOfWisdom: true,
					misery: true,
					curseOfElements: TristateEffect.TristateEffectRegular,
				}),
			},

			// IconInputs to include in the 'Self Buffs' section on the settings tab.
			selfBuffInputs: [
				IconInputs.DrumsOfBattleConsume,
				IconInputs.DrumsOfRestorationConsume,
			],
			// IconInputs to include in the 'Other Buffs' section on the settings tab.
			raidBuffInputs: [
				IconInputs.ArcaneBrilliance,
				IconInputs.DivineSpirit,
				IconInputs.GiftOfTheWild,
			],
			partyBuffInputs: [
				IconInputs.MoonkinAura,
				IconInputs.DrumsOfBattleBuff,
				IconInputs.DrumsOfRestorationBuff,
				IconInputs.Bloodlust,
				IconInputs.WrathOfAirTotem,
				IconInputs.TotemOfWrath,
				IconInputs.ManaSpringTotem,
				IconInputs.DraeneiRacialCaster,
				IconInputs.EyeOfTheNight,
				IconInputs.ChainOfTheTwilightOwl,
				IconInputs.JadePendantOfBlasting,
				IconInputs.AtieshWarlock,
				IconInputs.AtieshMage,
			],
			playerBuffInputs: [
				IconInputs.BlessingOfKings,
				IconInputs.BlessingOfWisdom,
			],
			// IconInputs to include in the 'Debuffs' section on the settings tab.
			debuffInputs: [
				IconInputs.JudgementOfWisdom,
				IconInputs.ImprovedSealOfTheCrusader,
				IconInputs.CurseOfElements,
			],
			// IconInputs to include in the 'Consumes' section on the settings tab.
			consumeInputs: [
				IconInputs.DefaultSuperManaPotion,
				IconInputs.DefaultDestructionPotion,
				IconInputs.DarkRune,
				IconInputs.FlaskOfPureDeath,
				IconInputs.FlaskOfSupremePower,
				IconInputs.AdeptsElixir,
				IconInputs.ElixirOfMajorMageblood,
				IconInputs.ElixirOfDraenicWisdom,
				IconInputs.BrilliantWizardOil,
				IconInputs.SuperiorWizardOil,
				IconInputs.BlackenedBasilisk,
				IconInputs.SkullfishSoup,
				IconInputs.KreegsStoutBeatdown,
			],
			// Inputs to include in the 'Rotation' section on the settings tab.
			rotationInputs: ShadowPriestInputs.ShadowPriestRotationConfig,
			// Inputs to include in the 'Other' section on the settings tab.
			otherInputs: {
				inputs: [
					OtherInputs.ISBUptime,
					OtherInputs.ShadowPriestDPS,
					OtherInputs.StartingPotion,
					OtherInputs.NumStartingPotions,
				],
			},
			encounterPicker: {
				// Whether to include 'Target Armor' in the 'Encounter' section of the settings tab.
				showTargetArmor: false,
				// Whether to include 'Num Targets' in the 'Encounter' section of the settings tab.
				showNumTargets: true,
			},

			// If true, the talents on the talents tab will not be individually modifiable by the user.
			// Note that the use can still pick between preset talents, if there is more than 1.
			freezeTalents: false,

			presets: {
				// Preset talents that the user can quickly select.
				talents: [
					Presets.StandardTalents,
				],
				// Preset gear configurations that the user can quickly select.
				gear: [
					Presets.P1_BIS,
					Presets.P2_BIS,
					Presets.P3_BIS,
				],
			},
		});
	}
}
