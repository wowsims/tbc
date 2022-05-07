import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { StrengthOfEarthType } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { BattleElixir } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { GuardianElixir } from '/tbc/core/proto/common.js';
import { Conjured } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as HunterInputs from './inputs.js';
import * as Presets from './presets.js';
export class HunterSimUI extends IndividualSimUI {
    constructor(parentElem, player) {
        super(parentElem, player, {
            cssClass: 'hunter-sim-ui',
            // List any known bugs / issues here and they'll be shown on the site.
            knownIssues: [],
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
                    blessingOfSalvation: true,
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
            ],
            playerBuffInputs: [
                IconInputs.BlessingOfKings,
                IconInputs.BlessingOfWisdom,
                IconInputs.BlessingOfMight,
                IconInputs.BlessingOfSalvation,
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
                IconInputs.GiftOfArthas,
            ],
            // Which options are selectable in the 'Consumes' section.
            consumeOptions: {
                potions: [
                    Potions.HastePotion,
                    Potions.SuperManaPotion,
                    Potions.FelManaPotion,
                ],
                conjured: [
                    Conjured.ConjuredDarkRune,
                    Conjured.ConjuredFlameCap,
                ],
                flasks: [
                    Flask.FlaskOfRelentlessAssault,
                ],
                battleElixirs: [
                    BattleElixir.ElixirOfDemonslaying,
                    BattleElixir.ElixirOfMajorAgility,
                    BattleElixir.ElixirOfTheMongoose,
                ],
                guardianElixirs: [
                    GuardianElixir.ElixirOfDraenicWisdom,
                    GuardianElixir.ElixirOfMajorMageblood,
                ],
                food: [
                    Food.FoodGrilledMudfish,
                    Food.FoodRavagerDog,
                    Food.FoodSpicyHotTalbuk,
                    Food.FoodRoastedClefthoof,
                ],
                alcohol: [],
                weaponImbues: [
                    WeaponImbue.WeaponImbueAdamantiteSharpeningStone,
                    WeaponImbue.WeaponImbueAdamantiteWeightstone,
                ],
                pet: [
                    IconInputs.KiblersBits,
                    IconInputs.PetScrollOfAgilityV,
                    IconInputs.PetScrollOfStrengthV,
                ],
                other: [
                    IconInputs.ScrollOfAgilityV,
                    IconInputs.ScrollOfStrengthV,
                ],
            },
            // Inputs to include in the 'Rotation' section on the settings tab.
            rotationInputs: HunterInputs.HunterRotationConfig,
            // Inputs to include in the 'Other' section on the settings tab.
            otherInputs: {
                inputs: [
                    HunterInputs.PetTypeInput,
                    HunterInputs.PetUptime,
                    HunterInputs.PetSingleAbility,
                    HunterInputs.LatencyMs,
                    OtherInputs.StartingPotion,
                    OtherInputs.NumStartingPotions,
                    OtherInputs.ExposeWeaknessUptime,
                    OtherInputs.ExposeWeaknessHunterAgility,
                    OtherInputs.SnapshotImprovedStrengthOfEarthTotem,
                    OtherInputs.SnapshotBsSolarianSapphire,
                    OtherInputs.SnapshotBsT2,
                    OtherInputs.InFrontOfTarget,
                ],
            },
            encounterPicker: {
                // Target stats to show for 'Simple' encounters.
                simpleTargetStats: [
                    Stat.StatArmor,
                ],
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
                    Presets.P4_BM_PRESET,
                    Presets.P5_BM_PRESET,
                    Presets.P1_SV_PRESET,
                    Presets.P2_SV_PRESET,
                    Presets.P3_SV_PRESET,
                    Presets.P4_SV_PRESET,
                    Presets.P5_SV_PRESET,
                ],
            },
        });
    }
}
