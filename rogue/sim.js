import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { StrengthOfEarthType } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as RogueInputs from './inputs.js';
import * as Presets from './presets.js';
export class RogueSimUI extends IndividualSimUI {
    constructor(parentElem, player) {
        super(parentElem, player, {
            cssClass: 'rogue-sim-ui',
            // List any known bugs / issues here and they'll be shown on the site.
            knownIssues: [],
            // All stats for which EP should be calculated.
            epStats: [
                Stat.StatAgility,
                Stat.StatStrength,
                Stat.StatAttackPower,
                Stat.StatMeleeHit,
                Stat.StatMeleeCrit,
                Stat.StatMeleeHaste,
                Stat.StatArmorPenetration,
                Stat.StatExpertise,
            ],
            // Reference stat against which to calculate EP.
            epReferenceStat: Stat.StatAttackPower,
            // Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
            displayStats: [
                Stat.StatStamina,
                Stat.StatAgility,
                Stat.StatStrength,
                Stat.StatAttackPower,
                Stat.StatMeleeHit,
                Stat.StatMeleeCrit,
                Stat.StatMeleeHaste,
                Stat.StatArmorPenetration,
                Stat.StatExpertise,
            ],
            defaults: {
                // Default equipped gear.
                gear: Presets.P1_PRESET.gear,
                // Default EP weights for sorting gear in the gear picker.
                epWeights: Stats.fromMap({
                    [Stat.StatAgility]: 2.5,
                    [Stat.StatStrength]: 1,
                    [Stat.StatAttackPower]: 1,
                    [Stat.StatMeleeHit]: 1,
                    [Stat.StatMeleeCrit]: 1,
                    [Stat.StatMeleeHaste]: 1.4,
                    [Stat.StatArmorPenetration]: 0.4,
                    [Stat.StatExpertise]: 3,
                }),
                // Default consumes settings.
                consumes: Presets.DefaultConsumes,
                // Default rotation settings.
                rotation: Presets.DefaultRotation,
                // Default talents.
                talents: Presets.CombatTalents.data,
                // Default spec-specific settings.
                specOptions: Presets.DefaultOptions,
                // Default raid/party buffs settings.
                raidBuffs: RaidBuffs.create({
                    giftOfTheWild: TristateEffect.TristateEffectImproved,
                }),
                partyBuffs: PartyBuffs.create({
                    bloodlust: 1,
                    graceOfAirTotem: TristateEffect.TristateEffectImproved,
                    strengthOfEarthTotem: StrengthOfEarthType.EnhancingTotems,
                    windfuryTotemRank: 5,
                    battleShout: TristateEffect.TristateEffectImproved,
                    leaderOfThePack: TristateEffect.TristateEffectImproved,
                }),
                individualBuffs: IndividualBuffs.create({
                    blessingOfKings: true,
                    blessingOfMight: TristateEffect.TristateEffectImproved,
                }),
                debuffs: Debuffs.create({
                    sunderArmor: true,
                    curseOfRecklessness: true,
                    faerieFire: TristateEffect.TristateEffectImproved,
                    improvedSealOfTheCrusader: true,
                }),
            },
            // IconInputs to include in the 'Self Buffs' section on the settings tab.
            selfBuffInputs: [
                IconInputs.DrumsOfBattleConsume,
                IconInputs.BattleChicken,
            ],
            // IconInputs to include in the 'Other Buffs' section on the settings tab.
            raidBuffInputs: [
                IconInputs.GiftOfTheWild,
            ],
            partyBuffInputs: [
                IconInputs.DrumsOfBattleBuff,
                IconInputs.Bloodlust,
                IconInputs.GraceOfAirTotem,
                IconInputs.WindfuryTotem,
                IconInputs.StrengthOfEarthTotem,
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
                IconInputs.BlessingOfMight,
                IconInputs.UnleashedRage,
            ],
            // IconInputs to include in the 'Debuffs' section on the settings tab.
            debuffInputs: [
                IconInputs.BloodFrenzy,
                IconInputs.ImprovedSealOfTheCrusader,
                IconInputs.HuntersMark,
                IconInputs.FaerieFire,
                IconInputs.SunderArmor,
                IconInputs.ExposeArmor,
                IconInputs.CurseOfRecklessness,
                IconInputs.Misery,
            ],
            // IconInputs to include in the 'Consumes' section on the settings tab.
            consumeInputs: [
                IconInputs.MainHandAdamantiteSharpeningStone,
                IconInputs.MainHandAdamantiteWeightstone,
                IconInputs.OffHandAdamantiteSharpeningStone,
                IconInputs.OffHandAdamantiteWeightstone,
                IconInputs.DefaultHastePotion,
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
            rotationInputs: RogueInputs.RogueRotationConfig,
            // Inputs to include in the 'Other' section on the settings tab.
            otherInputs: {
                inputs: [
                    OtherInputs.StartingPotion,
                    OtherInputs.NumStartingPotions,
                    OtherInputs.ExposeWeaknessUptime,
                    OtherInputs.ExposeWeaknessHunterAgility,
                    OtherInputs.SnapshotImprovedStrengthOfEarthTotem,
                    OtherInputs.SnapshotBsSolarianSapphire,
                    OtherInputs.SnapshotBsT2,
                ],
            },
            additionalIconSections: {},
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
                    Presets.CombatTalents,
                    Presets.CombatMaceTalents,
                ],
                // Preset gear configurations that the user can quickly select.
                gear: [
                    Presets.P1_PRESET,
                    Presets.P2_PRESET,
                    Presets.P3_PRESET,
                ],
            },
        });
    }
}
