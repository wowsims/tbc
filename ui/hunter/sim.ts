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
import { StrengthOfEarthType } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js'
import { WeaponImbue } from '/tbc/core/proto/common.js'
import { Player } from '/tbc/core/player.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Sim } from '/tbc/core/sim.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

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
				'This sim is newly released and there are likely still bugs. Take the DPS values with a grain of salt and let us know if you spot any issues!',
			],
			warnings: [
				(simUI: IndividualSimUI<Spec.SpecHunter>) => {
					return {
						updateOn: TypedEvent.onAny([simUI.player.rotationChangeEmitter, simUI.player.consumesChangeEmitter, simUI.player.getParty()!.buffsChangeEmitter]),
						shouldDisplay: () => {
							const rotation = simUI.player.getRotation();
							const isMeleeWeaving = rotation.meleeWeave && rotation.percentWeaved > 0;

							return !isMeleeWeaving &&
									simUI.player.getConsumes().mainHandImbue != WeaponImbue.WeaponImbueUnknown &&
									(simUI.player.getParty() != null && simUI.player.getParty()!.getBuffs().windfuryTotemRank > 0);
						},
						getContent: () => 'Melee weaving is off but Windfury Totem is on, so your main hand imbue is being ignored without any benefit.',
					};
				},
			],

			// All stats for which EP should be calculated.
			epStats: [
				Stat.StatIntellect,
				Stat.StatAgility,
				Stat.StatStrength,
				Stat.StatAttackPower,
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
				Stat.StatAgility,
				Stat.StatStrength,
				Stat.StatIntellect,
				Stat.StatAttackPower,
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
					[Stat.StatIntellect]: 0.01,
					[Stat.StatAgility]: 2.5,
					[Stat.StatStrength]: 0.15,
					[Stat.StatAttackPower]: 0.15,
					[Stat.StatRangedAttackPower]: 1.0,
					[Stat.StatMeleeHit]: 0.3,
					[Stat.StatMeleeCrit]: 2.3,
					[Stat.StatMeleeHaste]: 1.97,
					[Stat.StatArmorPenetration]: 0.4,
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
					giftOfTheWild: TristateEffect.TristateEffectImproved,
				}),
				partyBuffs: PartyBuffs.create({
					bloodlust: 1,
					graceOfAirTotem: TristateEffect.TristateEffectImproved,
					strengthOfEarthTotem: StrengthOfEarthType.EnhancingAndCyclone,
					windfuryTotemRank: 5,
					battleShout: TristateEffect.TristateEffectImproved,
					leaderOfThePack: TristateEffect.TristateEffectImproved,
				}),
				individualBuffs: IndividualBuffs.create({
					blessingOfKings: true,
					blessingOfWisdom: 2,
					blessingOfMight: 2,
				}),
				debuffs: Debuffs.create({
					sunderArmor: true,
					curseOfRecklessness: true,
					faerieFire: TristateEffect.TristateEffectImproved,
					improvedSealOfTheCrusader: true,
					judgementOfWisdom: true,
					curseOfElements: TristateEffect.TristateEffectRegular,
				}),
			},

			// IconInputs to include in the 'Self Buffs' section on the settings tab.
			selfBuffInputs: [
				HunterInputs.Quiver,
				HunterInputs.WeaponAmmo,
				IconInputs.DrumsOfBattleConsume,
				IconInputs.BattleChicken,
			],
			// IconInputs to include in the 'Other Buffs' section on the settings tab.
			raidBuffInputs: [
				IconInputs.ArcaneBrilliance,
				IconInputs.DivineSpirit,
				IconInputs.GiftOfTheWild,
			],
			partyBuffInputs: [
				IconInputs.DrumsOfBattleBuff,
				IconInputs.Bloodlust,
				IconInputs.GraceOfAirTotem,
				IconInputs.WindfuryTotem,
				IconInputs.StrengthOfEarthTotem,
				IconInputs.ManaSpringTotem,
				IconInputs.BattleShout,
				IconInputs.DraeneiRacialMelee,
				IconInputs.LeaderOfThePack,
				IconInputs.FerociousInspiration,
				IconInputs.TrueshotAura,
				IconInputs.SanctityAura,
				IconInputs.BraidedEterniumChain,
				IconInputs.BattleChickens,
			],
			playerBuffInputs: [
				IconInputs.BlessingOfKings,
				IconInputs.BlessingOfWisdom,
				IconInputs.BlessingOfMight,
				IconInputs.UnleashedRage,
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
				IconInputs.Misery,
			],
			// IconInputs to include in the 'Consumes' section on the settings tab.
			consumeInputs: [
				IconInputs.MainHandAdamantiteSharpeningStone,
				IconInputs.MainHandAdamantiteWeightstone,
				IconInputs.OffHandAdamantiteSharpeningStone,
				IconInputs.OffHandAdamantiteWeightstone,
				IconInputs.DefaultHastePotion,
				IconInputs.DefaultSuperManaPotion,
				IconInputs.DefaultDarkRune,
				IconInputs.FlaskOfRelentlessAssault,
				IconInputs.ElixirOfDemonslaying,
				IconInputs.ElixirOfMajorAgility,
				IconInputs.ElixirOfTheMongoose,
				IconInputs.ElixirOfDraenicWisdom,
				IconInputs.ElixirOfMajorMageblood,
				IconInputs.RavagerDog,
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
					HunterInputs.PetUptime,
					HunterInputs.LatencyMs,
					OtherInputs.StartingPotion,
					OtherInputs.NumStartingPotions,
					OtherInputs.ExposeWeaknessUptime,
					OtherInputs.ExposeWeaknessHunterAgility,
					OtherInputs.SnapshotImprovedStrengthOfEarthTotem,
					OtherInputs.SnapshotBsSolarianSapphire,
					OtherInputs.SnapshotBsT2,
				],
			},
			additionalIconSections: {
				'Pet Buffs': [
					IconInputs.KiblersBits,
					IconInputs.PetScrollOfAgilityV,
					IconInputs.PetScrollOfStrengthV,
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
					Presets.MarksmanTalents,
					Presets.SurvivalTalents,
				],
				// Preset gear configurations that the user can quickly select.
				gear: [
					Presets.P1_BM_PRESET,
					Presets.P2_BM_PRESET,
					Presets.P3_BM_PRESET,
					Presets.P1_SV_PRESET,
					Presets.P2_SV_PRESET,
					Presets.P3_SV_PRESET,
				],
			},
		});
	}
}
