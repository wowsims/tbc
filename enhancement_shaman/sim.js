import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as ShamanInputs from './inputs.js';
import * as Presets from './presets.js';
export class EnhancementShamanSimUI extends IndividualSimUI {
    constructor(parentElem, player) {
        super(parentElem, player, {
            cssClass: 'enhancement-shaman-sim-ui',
            // List any known bugs / issues here and they'll be shown on the site.
            knownIssues: [],
            // All stats for which EP should be calculated.
            epStats: [
                Stat.StatIntellect,
                Stat.StatAgility,
                Stat.StatStrength,
                Stat.StatSpellPower,
                Stat.StatAttackPower,
                Stat.StatMeleeHit,
                Stat.StatMeleeCrit,
                Stat.StatMeleeHaste,
                Stat.StatArmorPenetration,
                Stat.StatExpertise,
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
                Stat.StatNatureSpellPower,
                Stat.StatSpellHit,
                Stat.StatSpellCrit,
                Stat.StatSpellHaste,
            ],
            defaults: {
                // Default equipped gear.
                gear: Presets.P1_PRESET.gear,
                // Default EP weights for sorting gear in the gear picker.
                epWeights: Stats.fromMap({
                    [Stat.StatIntellect]: 0.078,
                    [Stat.StatAgility]: 1.317,
                    [Stat.StatStrength]: 2.2,
                    [Stat.StatSpellPower]: 0.433,
                    [Stat.StatNatureSpellPower]: 0.216,
                    [Stat.StatAttackPower]: 1.0,
                    [Stat.StatMeleeHit]: 1.665,
                    [Stat.StatMeleeCrit]: 1.357,
                    [Stat.StatMeleeHaste]: 1.944,
                    [Stat.StatArmorPenetration]: 0.283,
                    [Stat.StatExpertise]: 2.871,
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
                partyBuffs: PartyBuffs.create({}),
                individualBuffs: IndividualBuffs.create({
                    blessingOfKings: true,
                    blessingOfWisdom: 2,
                    blessingOfMight: 2,
                }),
                debuffs: Debuffs.create({
                    faerieFire: 2,
                    improvedSealOfTheCrusader: true,
                    sunderArmor: true,
                    curseOfRecklessness: true,
                }),
            },
            // IconInputs to include in the 'Self Buffs' section on the settings tab.
            selfBuffInputs: [
                ShamanInputs.IconWaterShield,
                ShamanInputs.MainHandImbue,
                ShamanInputs.OffHandImbue,
                ShamanInputs.IconBloodlust,
                IconInputs.DrumsOfBattleConsume,
                IconInputs.DrumsOfRestorationConsume,
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
                IconInputs.FaerieFire,
                IconInputs.SunderArmor,
                IconInputs.ExposeArmor,
                IconInputs.CurseOfRecklessness,
                IconInputs.CurseOfElements,
                IconInputs.Misery,
                IconInputs.ImprovedScorch,
                IconInputs.WintersChill,
            ],
            // IconInputs to include in the 'Consumes' section on the settings tab.
            consumeInputs: [
                IconInputs.DefaultHastePotion,
                IconInputs.DefaultDarkRune,
                IconInputs.FlaskOfRelentlessAssault,
                IconInputs.ElixirOfDemonslaying,
                IconInputs.ElixirOfMajorAgility,
                IconInputs.RoastedClefthoof,
                IconInputs.ScrollOfAgilityV,
                IconInputs.ScrollOfStrengthV,
            ],
            // Inputs to include in the 'Rotation' section on the settings tab.
            rotationInputs: ShamanInputs.EnhancementShamanRotationConfig,
            // Inputs to include in the 'Other' section on the settings tab.
            otherInputs: {
                inputs: [
                    ShamanInputs.DelayOffhandSwings,
                    OtherInputs.StartingPotion,
                    OtherInputs.NumStartingPotions,
                    OtherInputs.ExposeWeaknessUptime,
                    OtherInputs.ExposeWeaknessHunterAgility,
                ],
            },
            customSections: [
                ShamanInputs.TotemsSection,
            ],
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
                    Presets.StandardTalents,
                    Presets.RestoSubspecTalents,
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
