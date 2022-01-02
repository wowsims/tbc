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
                gear: Presets.PRERAID_GEAR.gear,
                // Default EP weights for sorting gear in the gear picker.
                epWeights: Stats.fromMap({
                    [Stat.StatIntellect]: 1.02,
                    [Stat.StatAgility]: 2.23,
                    [Stat.StatStrength]: 2.04,
                    [Stat.StatSpellPower]: 0.45,
                    [Stat.StatAttackPower]: 1.0,
                    [Stat.StatMeleeHit]: 1.0,
                    [Stat.StatMeleeCrit]: 2.86,
                    [Stat.StatMeleeHaste]: 0.62,
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
                ShamanInputs.IconBloodlust,
                // TODO: add totem icons
                // ShamanInputs.IconManaSpringTotem,
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
            ],
            playerBuffInputs: [
                IconInputs.BlessingOfKings,
                IconInputs.BlessingOfWisdom,
                IconInputs.BlessingOfMight,
            ],
            // IconInputs to include in the 'Debuffs' section on the settings tab.
            debuffInputs: [
                IconInputs.ImprovedSealOfTheCrusader,
                IconInputs.FaerieFire,
                IconInputs.SunderArmor,
                IconInputs.CurseOfRecklessness,
                IconInputs.ExposeArmor,
            ],
            // IconInputs to include in the 'Consumes' section on the settings tab.
            consumeInputs: [
                IconInputs.DefaultSuperManaPotion,
                IconInputs.DefaultHastePotion,
                IconInputs.DefaultDarkRune,
                IconInputs.FlaskOfRelentlessAssault,
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
                    OtherInputs.StartingPotion,
                    OtherInputs.NumStartingPotions,
                ],
            },
            customSections: [
                ShamanInputs.TotemsSection,
            ],
            encounterPicker: {
                // Whether to include 'Target Armor' in the 'Encounter' section of the settings tab.
                showTargetArmor: false,
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
                ],
                // Preset gear configurations that the user can quickly select.
                gear: [
                    Presets.PRERAID_GEAR,
                ],
            },
        });
    }
}
