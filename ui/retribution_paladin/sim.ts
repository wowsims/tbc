import { RaidBuffs, StrengthOfEarthType } from '/tbc/core/proto/common.js';
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

import { RetributionPaladin, RetributionPaladin_Rotation as RetributionPaladinRotation, PaladinTalents as PaladinTalents, RetributionPaladin_Options as RetributionPaladinOptions } from '/tbc/core/proto/paladin.js';

import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';

import * as RetributionPaladinInputs from './inputs.js';
import * as Presets from './presets.js';

export class RetributionPaladinSimUI extends IndividualSimUI<Spec.SpecRetributionPaladin> {
  constructor(parentElem: HTMLElement, player: Player<Spec.SpecRetributionPaladin>) {
		super(parentElem, player, {
			cssClass: 'retribution-paladin-sim-ui',
			// List any known bugs / issues here and they'll be shown on the site.
			knownIssues: [
				"Basically everything is WIP"
			],

			// All stats for which EP should be calculated.
			epStats: [
				Stat.StatStrength,
				Stat.StatAgility,
				Stat.StatAttackPower,
				Stat.StatExpertise,
				Stat.StatMeleeHit,
				Stat.StatMeleeCrit,
				Stat.StatMeleeHaste,
				Stat.StatArmorPenetration,
			],
			// Reference stat against which to calculate EP. I think all classes use either spell power or attack power.
			epReferenceStat: Stat.StatAttackPower,
			// Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
			displayStats: [
				Stat.StatStamina,
				Stat.StatStrength,
				Stat.StatAgility,
				Stat.StatIntellect,
				Stat.StatAttackPower,
				Stat.StatMeleeHit,
				Stat.StatMeleeCrit,
				Stat.StatMeleeHaste,
				Stat.StatExpertise,
				Stat.StatArmorPenetration,
				Stat.StatSpellPower,
				Stat.StatHolySpellPower,
				Stat.StatSpellHit,
				Stat.StatSpellCrit,
				Stat.StatSpellHaste,
			],

			defaults: {
				// Default equipped gear.
				gear: Presets.P3_PRESET.gear,
				// Default EP weights for sorting gear in the gear picker.
				epWeights: Stats.fromMap({
					[Stat.StatStrength]: 2.5,
					[Stat.StatAgility]: 1.75,
					[Stat.StatAttackPower]: 1,
					[Stat.StatExpertise]: 3.75,
					[Stat.StatMeleeHit]: 1.5,
					[Stat.StatMeleeCrit]: 2.5,
					[Stat.StatMeleeHaste]: 3,
					[Stat.StatArmorPenetration]: 0.5,
				}),
				// Default consumes settings.
				consumes: Presets.DefaultConsumes,
				// Default rotation settings.
				rotation: Presets.DefaultRotation,
				// Default talents.
				talents: Presets.RetKingsPaladinTalents.data,
				// Default spec-specific settings.
				specOptions: Presets.DefaultOptions,
				// Default raid/party buffs settings.
				raidBuffs: RaidBuffs.create({
				}),
				partyBuffs: PartyBuffs.create({
				}),
				individualBuffs: IndividualBuffs.create({
					blessingOfKings: true,
					blessingOfMight: TristateEffect.TristateEffectImproved,
				}),
				debuffs: Debuffs.create({
				}),
			},

			// IconInputs to include in the 'Self Buffs' section on the settings tab.
			selfBuffInputs: [
				IconInputs.DrumsOfBattleConsume,
				IconInputs.BattleChicken,
			],
			// IconInputs to include in the 'Other Buffs' section on the settings tab.
			raidBuffInputs: [
				IconInputs.ArcaneBrilliance,
				IconInputs.GiftOfTheWild,
			],
			partyBuffInputs: [
				IconInputs.DrumsOfBattleBuff,
				IconInputs.Bloodlust,
				IconInputs.ManaSpringTotem,
				IconInputs.WrathOfAirTotem,
				IconInputs.WindfuryTotem,
				IconInputs.TotemOfWrath,
				IconInputs.BattleShout,
				IconInputs.DraeneiRacialMelee,
				IconInputs.LeaderOfThePack,
				IconInputs.MoonkinAura,
				IconInputs.FerociousInspiration,
				IconInputs.TrueshotAura,
				IconInputs.SanctityAura,
				IconInputs.BattleChickens,
				IconInputs.BraidedEterniumChain,
				IconInputs.EyeOfTheNight,
				IconInputs.ChainOfTheTwilightOwl,
			],
			playerBuffInputs: [
				IconInputs.BlessingOfKings,
				IconInputs.BlessingOfWisdom,
				IconInputs.BlessingOfMight,
			],
			// IconInputs to include in the 'Debuffs' section on the settings tab.
			debuffInputs: [
				IconInputs.BloodFrenzy,
				IconInputs.JudgementOfWisdom,
				IconInputs.HuntersMark,
				IconInputs.FaerieFire,
				IconInputs.SunderArmor,
				IconInputs.ExposeArmor,
				IconInputs.CurseOfRecklessness,
				IconInputs.Misery,
			],
			// IconInputs to include in the 'Consumes' section on the settings tab.
			consumeInputs: [
				IconInputs.DefaultHastePotion,
				IconInputs.DefaultSuperManaPotion,
				IconInputs.DefaultDarkRune,
				IconInputs.DefaultFlameCap,
				IconInputs.FlaskOfRelentlessAssault,
				IconInputs.ElixirOfDemonslaying,
				IconInputs.ElixirOfMajorAgility,
				IconInputs.ElixirOfMajorStrength,
				IconInputs.ElixirOfTheMongoose,
				IconInputs.ElixirOfMajorMageblood,
				IconInputs.ElixirOfDraenicWisdom,
				IconInputs.MainHandAdamantiteSharpeningStone,
				IconInputs.MainHandElementalSharpeningStone,
				IconInputs.MainHandSuperiorWizardOil,
				IconInputs.MainHandBrilliantWizardOil,
				IconInputs.RoastedClefthoof,
				IconInputs.SpicyHotTalbuk,
				IconInputs.GrilledMudfish,
				IconInputs.BlackenedBasilisk,
			],
			// Inputs to include in the 'Rotation' section on the settings tab.
			rotationInputs: RetributionPaladinInputs.RetributionPaladinRotationConfig,
			// Inputs to include in the 'Other' section on the settings tab.
			otherInputs: {
				inputs: [
				],
			},
			encounterPicker: {
				// Whether to include 'Target Armor' in the 'Encounter' section of the settings tab.
				showTargetArmor: true,
				// Whether to include 'Execute Duration (%)' in the 'Encounter' section of the settings tab.
				showExecuteProportion: false,
				// Whether to include 'Num Targets' in the 'Encounter' section of the settings tab.
				showNumTargets: true,
			},

			// If true, the talents on the talents tab will not be individually modifiable by the user.
			// Note that the use can still pick between preset talents, if there is more than 1.
			freezeTalents: false,

			presets: {
				// Preset talents that the user can quickly select.
				talents: [
					Presets.RetKingsPaladinTalents,
					Presets.RetNoKingsPaladinTalents,
				],
				// Preset gear configurations that the user can quickly select.
				gear: [
					Presets.PRE_RAID_PRESET, 
					Presets.P1_PRESET,
					Presets.P2_PRESET,
					Presets.P3_PRESET,
					Presets.P4_PRESET,
					Presets.P5_PRESET,
				],
			},
		});
	}
}