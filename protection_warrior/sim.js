import { RaidBuffs, StrengthOfEarthType } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { BattleElixir } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { Conjured } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as ProtectionWarriorInputs from './inputs.js';
import * as Presets from './presets.js';
export class ProtectionWarriorSimUI extends IndividualSimUI {
    constructor(parentElem, player) {
        super(parentElem, player, {
            cssClass: 'protection-warrior-sim-ui',
            // List any known bugs / issues here and they'll be shown on the site.
            knownIssues: [],
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
                Stat.StatDefense,
                Stat.StatBlock,
                Stat.StatBlockValue,
                Stat.StatDodge,
                Stat.StatParry,
                Stat.StatResilience,
            ],
            // Reference stat against which to calculate EP. I think all classes use either spell power or attack power.
            epReferenceStat: Stat.StatAttackPower,
            // Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
            displayStats: [
                Stat.StatHealth,
                Stat.StatStamina,
                Stat.StatStrength,
                Stat.StatAgility,
                Stat.StatAttackPower,
                Stat.StatExpertise,
                Stat.StatMeleeHit,
                Stat.StatMeleeCrit,
                Stat.StatMeleeHaste,
                Stat.StatArmorPenetration,
                Stat.StatDefense,
                Stat.StatBlock,
                Stat.StatBlockValue,
                Stat.StatDodge,
                Stat.StatParry,
                Stat.StatResilience,
            ],
            defaults: {
                // Default equipped gear.
                gear: Presets.P1_BALANCED_PRESET.gear,
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
                talents: Presets.ImpaleProtTalents.data,
                // Default spec-specific settings.
                specOptions: Presets.DefaultOptions,
                // Default raid/party buffs settings.
                raidBuffs: RaidBuffs.create({
                    giftOfTheWild: TristateEffect.TristateEffectImproved,
                }),
                partyBuffs: PartyBuffs.create({
                    bloodlust: 1,
                    drums: Drums.DrumsOfBattle,
                    graceOfAirTotem: TristateEffect.TristateEffectImproved,
                    strengthOfEarthTotem: StrengthOfEarthType.EnhancingTotems,
                    windfuryTotemRank: 5,
                    windfuryTotemIwt: 2,
                    leaderOfThePack: TristateEffect.TristateEffectImproved,
                }),
                individualBuffs: IndividualBuffs.create({
                    blessingOfKings: true,
                    blessingOfMight: TristateEffect.TristateEffectImproved,
                    unleashedRage: true,
                }),
                debuffs: Debuffs.create({
                    mangle: true,
                    curseOfRecklessness: true,
                    faerieFire: TristateEffect.TristateEffectImproved,
                    improvedSealOfTheCrusader: true,
                    huntersMark: TristateEffect.TristateEffectImproved,
                    exposeWeaknessUptime: 0.95,
                    exposeWeaknessHunterAgility: 1200,
                }),
            },
            // IconInputs to include in the 'Self Buffs' section on the settings tab.
            selfBuffInputs: [],
            // IconInputs to include in the 'Other Buffs' section on the settings tab.
            raidBuffInputs: [
                IconInputs.GiftOfTheWild,
            ],
            partyBuffInputs: [
                IconInputs.DrumsOfBattleBuff,
                IconInputs.Bloodlust,
                IconInputs.StrengthOfEarthTotem,
                IconInputs.GraceOfAirTotem,
                IconInputs.WindfuryTotem,
                IconInputs.BattleShout,
                IconInputs.LeaderOfThePack,
                IconInputs.FerociousInspiration,
                IconInputs.TrueshotAura,
                IconInputs.SanctityAura,
                IconInputs.DraeneiRacialMelee,
                IconInputs.BraidedEterniumChain,
            ],
            playerBuffInputs: [
                IconInputs.BlessingOfKings,
                IconInputs.BlessingOfMight,
                IconInputs.BlessingOfSalvation,
                IconInputs.UnleashedRage,
            ],
            // IconInputs to include in the 'Debuffs' section on the settings tab.
            debuffInputs: [
                IconInputs.BloodFrenzy,
                IconInputs.Mangle,
                IconInputs.ImprovedSealOfTheCrusader,
                IconInputs.HuntersMark,
                IconInputs.FaerieFire,
                IconInputs.SunderArmor,
                IconInputs.ExposeArmor,
                IconInputs.CurseOfRecklessness,
                IconInputs.GiftOfArthas,
            ],
            // Which options are selectable in the 'Consumes' section.
            consumeOptions: {
                potions: [
                    Potions.HastePotion,
                ],
                conjured: [
                    Conjured.ConjuredFlameCap,
                ],
                flasks: [
                    Flask.FlaskOfRelentlessAssault,
                ],
                battleElixirs: [
                    BattleElixir.ElixirOfDemonslaying,
                    BattleElixir.ElixirOfMajorStrength,
                    BattleElixir.ElixirOfMajorAgility,
                    BattleElixir.ElixirOfTheMongoose,
                ],
                guardianElixirs: [],
                food: [
                    Food.FoodRoastedClefthoof,
                    Food.FoodGrilledMudfish,
                    Food.FoodSpicyHotTalbuk,
                    Food.FoodRavagerDog,
                ],
                alcohol: [],
                weaponImbues: [
                    WeaponImbue.WeaponImbueAdamantiteSharpeningStone,
                    WeaponImbue.WeaponImbueAdamantiteWeightstone,
                ],
                other: [
                    IconInputs.ScrollOfAgilityV,
                    IconInputs.ScrollOfStrengthV,
                ],
            },
            // Inputs to include in the 'Rotation' section on the settings tab.
            rotationInputs: ProtectionWarriorInputs.ProtectionWarriorRotationConfig,
            // Inputs to include in the 'Other' section on the settings tab.
            otherInputs: {
                inputs: [
                    ProtectionWarriorInputs.ShoutPicker,
                    ProtectionWarriorInputs.PrecastShout,
                    ProtectionWarriorInputs.PrecastShoutWithSapphire,
                    ProtectionWarriorInputs.PrecastShoutWithT2,
                    OtherInputs.ExposeWeaknessUptime,
                    OtherInputs.ExposeWeaknessHunterAgility,
                    OtherInputs.SnapshotImprovedStrengthOfEarthTotem,
                    OtherInputs.TankAssignment,
                    OtherInputs.InFrontOfTarget,
                ],
            },
            encounterPicker: {
                // Target stats to show for 'Simple' encounters.
                simpleTargetStats: [
                    Stat.StatArmor,
                ],
                // Whether to include 'Execute Duration (%)' in the 'Encounter' section of the settings tab.
                showExecuteProportion: true,
                // Whether to include 'Num Targets' in the 'Encounter' section of the settings tab.
                showNumTargets: true,
            },
            // If true, the talents on the talents tab will not be individually modifiable by the user.
            // Note that the use can still pick between preset talents, if there is more than 1.
            freezeTalents: false,
            presets: {
                // Preset talents that the user can quickly select.
                talents: [
                    Presets.ImpaleProtTalents,
                ],
                // Preset gear configurations that the user can quickly select.
                gear: [
                    Presets.P1_BALANCED_PRESET,
                    Presets.P4_BALANCED_PRESET,
                ],
            },
        });
    }
}
