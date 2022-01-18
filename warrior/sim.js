import { RaidBuffs, StrengthOfEarthType } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as WarriorInputs from './inputs.js';
import * as Presets from './presets.js';
export class WarriorSimUI extends IndividualSimUI {
    constructor(parentElem, player) {
        super(parentElem, player, {
            cssClass: 'warrior-sim-ui',
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
            ],
            // Reference stat against which to calculate EP. I think all classes use either spell power or attack power.
            epReferenceStat: Stat.StatAttackPower,
            // Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
            displayStats: [
                Stat.StatStamina,
                Stat.StatStrength,
                Stat.StatAgility,
                Stat.StatAttackPower,
                Stat.StatExpertise,
                Stat.StatMeleeHit,
                Stat.StatMeleeCrit,
                Stat.StatMeleeHaste,
                Stat.StatArmorPenetration,
            ],
            defaults: {
                // Default equipped gear.
                gear: Presets.P1_FURY_PRESET.gear,
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
                consumes: Presets.DefaultFuryConsumes,
                // Default rotation settings.
                rotation: Presets.DefaultFuryRotation,
                // Default talents.
                talents: Presets.FuryTalents.data,
                // Default spec-specific settings.
                specOptions: Presets.DefaultFuryOptions,
                // Default raid/party buffs settings.
                raidBuffs: RaidBuffs.create({
                    giftOfTheWild: TristateEffect.TristateEffectImproved,
                }),
                partyBuffs: PartyBuffs.create({
                    drums: Drums.DrumsOfBattle,
                    bloodlust: 1,
                    strengthOfEarthTotem: StrengthOfEarthType.EnhancingAndCyclone,
                    windfuryTotemRank: 5,
                }),
                individualBuffs: IndividualBuffs.create({
                    blessingOfKings: true,
                    blessingOfMight: TristateEffect.TristateEffectImproved,
                    innervates: 1,
                }),
                debuffs: Debuffs.create({
                    bloodFrenzy: true,
                    curseOfRecklessness: true,
                    exposeArmor: TristateEffect.TristateEffectRegular
                }),
            },
            // IconInputs to include in the 'Self Buffs' section on the settings tab.
            selfBuffInputs: [
                // TODO: Move reck to cooldown tabs
                WarriorInputs.Recklessness,
                IconInputs.DrumsOfBattleConsume,
            ],
            // IconInputs to include in the 'Other Buffs' section on the settings tab.
            raidBuffInputs: [
                IconInputs.GiftOfTheWild,
            ],
            partyBuffInputs: [
                IconInputs.LeaderOfThePack,
                IconInputs.DrumsOfBattleBuff,
                IconInputs.Bloodlust,
                // IconInputs.WindfuryTotem,
                // IconInputs.StrengthOfEarthTotem,
                // IconInputs.UnleashedRage,
                IconInputs.FerociousInspiration,
                IconInputs.DraeneiRacialMelee,
                IconInputs.BraidedEterniumChain
            ],
            playerBuffInputs: [
                IconInputs.BlessingOfKings,
                IconInputs.BlessingOfMight,
            ],
            // IconInputs to include in the 'Debuffs' section on the settings tab.
            debuffInputs: [
                // IconInputs.ImprovedHuntersMark,
                IconInputs.BloodFrenzy,
                IconInputs.ImprovedSealOfTheCrusader,
                IconInputs.CurseOfRecklessness,
            ],
            // IconInputs to include in the 'Consumes' section on the settings tab.
            consumeInputs: [
                IconInputs.DefaultHastePotion,
                // IconInputs.DefaultMightyRagePotion,
                IconInputs.ElixirOfMajorAgility,
                // IconInputs.ElixirOfTheMongoose,
                IconInputs.ElixirOfDemonslaying,
                // IconInputs.AdamantiteStone,
                // IconInputs.ElementalStone,
                IconInputs.RoastedClefthoof,
                // IconInputs.SpicyHotTalbuk,
            ],
            // Inputs to include in the 'Rotation' section on the settings tab.
            rotationInputs: WarriorInputs.WarriorRotationConfig,
            // Inputs to include in the 'Other' section on the settings tab.
            otherInputs: {
                inputs: [],
            },
            encounterPicker: {
                // Whether to include 'Target Armor' in the 'Encounter' section of the settings tab.
                showTargetArmor: true,
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
                    Presets.ArmsSlamTalents,
                    Presets.ArmsDWTalents,
                    Presets.FuryTalents,
                ],
                // Preset gear configurations that the user can quickly select.
                gear: [
                    Presets.P1_FURY_PRESET,
                    Presets.P1_ARMSSLAM_PRESET,
                    Presets.P1_ARMSDW_PRESET,
                ],
            },
        });
    }
}
