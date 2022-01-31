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
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js'
import { Player } from '/tbc/core/player.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Sim } from '/tbc/core/sim.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';

import { Hunter, Hunter_Rotation as HunterRotation, Hunter_Options as HunterOptions } from '/tbc/core/proto/hunter.js';

import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';

import * as HunterInputs from './inputs.js';
import * as Presets from './presets.js';

export class HunterSimUI extends IndividualSimUI<Spec.SpecHunter> {
  constructor(parentElem: HTMLElement, player: Player<Spec.SpecHunter>) {
		super(parentElem, player, {
			cssClass: 'hunter-sim-ui',
			// List any known bugs / issues here and they'll be shown on the site.
			knownIssues: [
			],

			// All stats for which EP should be calculated.
			epStats: [
				Stat.StatIntellect,
				Stat.StatAgility,
				Stat.StatStrength,
				Stat.StatRangedAttackPower,
				Stat.StatMeleeHit,
				Stat.StatMeleeCrit,
				Stat.StatMeleeHaste,
				Stat.StatArmorPenetration,
			],
			// Reference stat against which to calculate EP.
			epReferenceStat: Stat.StatRangedAttackPower,
			// Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
			displayStats: [
				Stat.StatStamina,
				Stat.StatStrength,
				Stat.StatAgility,
				Stat.StatIntellect,
				Stat.StatRangedAttackPower,
				Stat.StatMeleeHit,
				Stat.StatMeleeCrit,
				Stat.StatMeleeHaste,
				Stat.StatArmorPenetration,
			],

			defaults: {
				// Default equipped gear.
				gear: Presets.P1_BM_PRESET.gear,
				// Default EP weights for sorting gear in the gear picker.
				epWeights: Stats.fromMap({
					[Stat.StatIntellect]: 0.078,
					[Stat.StatAgility]: 1.317,
					[Stat.StatStrength]: 2.2,
					[Stat.StatRangedAttackPower]: 1.0,
					[Stat.StatMeleeHit]: 1.665,
					[Stat.StatMeleeCrit]: 1.357,
					[Stat.StatMeleeHaste]: 1.944,
					[Stat.StatArmorPenetration]: 0.283,
				}),
				// Default consumes settings.
				consumes: Presets.DefaultConsumes,
				// Default rotation settings.
				rotation: Presets.DefaultRotation,
				// Default talents.
				talents: Presets.BeastMasteryTalents.data,
				// Default spec-specific settings.
				specOptions: Presets.DefaultOptions,
				// Default raid/party buffs settings.
				raidBuffs: RaidBuffs.create({
					arcaneBrilliance: true,
					divineSpirit: TristateEffect.TristateEffectImproved,
					giftOfTheWild: TristateEffect.TristateEffectImproved,
				}),
				partyBuffs: PartyBuffs.create({
				}),
				individualBuffs: IndividualBuffs.create({
					blessingOfKings: true,
					blessingOfWisdom: 2,
					blessingOfMight: 2,
				}),
				debuffs: Debuffs.create({
					faerieFire: TristateEffect.TristateEffectImproved,
					judgementOfWisdom: true,
					improvedSealOfTheCrusader: true,
					sunderArmor: true,
					curseOfRecklessness: true,
					curseOfElements: TristateEffect.TristateEffectRegular,
				}),
			},

			// IconInputs to include in the 'Self Buffs' section on the settings tab.
			selfBuffInputs: [
				HunterInputs.Quiver,
				HunterInputs.WeaponAmmo,
				IconInputs.DrumsOfBattleConsume,
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
				IconInputs.BattleShout,
				IconInputs.DraeneiRacialMelee,
				IconInputs.LeaderOfThePack,
				IconInputs.FerociousInspiration,
				IconInputs.TrueshotAura,
				IconInputs.SanctityAura,
				IconInputs.BraidedEterniumChain,
			],
			playerBuffInputs: [
				IconInputs.BlessingOfKings,
				IconInputs.BlessingOfWisdom,
				IconInputs.BlessingOfMight,
			],
			// IconInputs to include in the 'Debuffs' section on the settings tab.
			debuffInputs: [
				IconInputs.BloodFrenzy,
				IconInputs.ImprovedSealOfTheCrusader,
				IconInputs.JudgementOfWisdom,
				IconInputs.HuntersMark,
				IconInputs.FaerieFire,
				IconInputs.SunderArmor,
				IconInputs.ExposeArmor,
				IconInputs.CurseOfRecklessness,
				IconInputs.CurseOfElements,
			],
			// IconInputs to include in the 'Consumes' section on the settings tab.
			consumeInputs: [
				IconInputs.DefaultHastePotion,
				IconInputs.DefaultSuperManaPotion,
				IconInputs.DefaultDarkRune,
				IconInputs.FlaskOfRelentlessAssault,
				IconInputs.ElixirOfDemonslaying,
				IconInputs.ElixirOfMajorStrength,
				IconInputs.ElixirOfMajorAgility,
				IconInputs.ElixirOfTheMongoose,
				IconInputs.ElixirOfDraenicWisdom,
				IconInputs.ElixirOfMajorMageblood,
				IconInputs.RoastedClefthoof,
				IconInputs.GrilledMudfish,
				IconInputs.SpicyHotTalbuk,
				IconInputs.ScrollOfAgilityV,
				IconInputs.ScrollOfStrengthV,
			],
			// Inputs to include in the 'Rotation' section on the settings tab.
			rotationInputs: HunterInputs.HunterRotationConfig,
			// Inputs to include in the 'Other' section on the settings tab.
			otherInputs: {
				inputs: [
					HunterInputs.PetTypeInput,
					HunterInputs.LatencyMs,
					OtherInputs.StartingPotion,
					OtherInputs.NumStartingPotions,
					OtherInputs.ExposeWeaknessUptime,
					OtherInputs.ExposeWeaknessHunterAgility,
					OtherInputs.SnapshotBsSolarianSapphire,
					OtherInputs.SnapshotBsT2,
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
					Presets.BeastMasteryTalents,
				],
				// Preset gear configurations that the user can quickly select.
				gear: [
					Presets.P1_BM_PRESET,
				],
			},
		});
	}
}
