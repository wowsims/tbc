import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as MageInputs from './inputs.js';
import * as Presets from './presets.js';
export class MageSimUI extends IndividualSimUI {
    constructor(parentElem, player) {
        super(parentElem, player, {
            cssClass: 'mage-sim-ui',
            // List any known bugs / issues here and they'll be shown on the site.
            knownIssues: [],
            // All stats for which EP should be calculated.
            epStats: [
                Stat.StatIntellect,
                Stat.StatSpirit,
                Stat.StatSpellPower,
                Stat.StatArcaneSpellPower,
                Stat.StatFireSpellPower,
                Stat.StatFrostSpellPower,
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
                Stat.StatArcaneSpellPower,
                Stat.StatFireSpellPower,
                Stat.StatFrostSpellPower,
                Stat.StatSpellHit,
                Stat.StatSpellCrit,
                Stat.StatSpellHaste,
                Stat.StatMP5,
            ],
            defaults: {
                // Default equipped gear.
                gear: Presets.P1_ARCANE_PRESET.gear,
                // Default EP weights for sorting gear in the gear picker.
                epWeights: Stats.fromMap({
                    [Stat.StatIntellect]: 1.29,
                    [Stat.StatSpirit]: 0.89,
                    [Stat.StatSpellPower]: 1,
                    [Stat.StatArcaneSpellPower]: 0.78,
                    [Stat.StatFireSpellPower]: 0,
                    [Stat.StatFrostSpellPower]: 0.21,
                    [Stat.StatSpellHit]: 0.5,
                    [Stat.StatSpellCrit]: 0.77,
                    [Stat.StatSpellHaste]: 0.84,
                    [Stat.StatMP5]: 0.61,
                }),
                // Default consumes settings.
                consumes: Presets.DefaultArcaneConsumes,
                // Default rotation settings.
                rotation: Presets.DefaultArcaneRotation,
                // Default talents.
                talents: Presets.ArcaneTalents.data,
                // Default spec-specific settings.
                specOptions: Presets.DefaultArcaneOptions,
                // Default raid/party buffs settings.
                raidBuffs: RaidBuffs.create({
                    giftOfTheWild: TristateEffect.TristateEffectImproved,
                }),
                partyBuffs: PartyBuffs.create({
                    drums: Drums.DrumsOfBattle,
                    bloodlust: 1,
                    manaSpringTotem: TristateEffect.TristateEffectImproved,
                    manaTideTotems: 1,
                    wrathOfAirTotem: TristateEffect.TristateEffectRegular,
                }),
                individualBuffs: IndividualBuffs.create({
                    blessingOfKings: true,
                    blessingOfWisdom: TristateEffect.TristateEffectImproved,
                    innervates: 1,
                }),
                debuffs: Debuffs.create({
                    judgementOfWisdom: true,
                    misery: true,
                    curseOfElements: TristateEffect.TristateEffectRegular,
                }),
            },
            // IconInputs to include in the 'Self Buffs' section on the settings tab.
            selfBuffInputs: [
                MageInputs.MageArmor,
                MageInputs.MoltenArmor,
                IconInputs.DrumsOfBattleConsume,
                IconInputs.DrumsOfRestorationConsume,
            ],
            // IconInputs to include in the 'Other Buffs' section on the settings tab.
            raidBuffInputs: [
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
                IconInputs.ManaTideTotem,
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
                IconInputs.Innervate,
                IconInputs.PowerInfusion,
            ],
            // IconInputs to include in the 'Debuffs' section on the settings tab.
            debuffInputs: [
                IconInputs.JudgementOfWisdom,
                IconInputs.ImprovedSealOfTheCrusader,
                IconInputs.CurseOfElements,
                IconInputs.Misery,
                IconInputs.ImprovedScorch,
                IconInputs.WintersChill,
            ],
            // IconInputs to include in the 'Consumes' section on the settings tab.
            consumeInputs: [
                IconInputs.DefaultSuperManaPotion,
                IconInputs.DefaultDestructionPotion,
                MageInputs.ManaEmerald,
                IconInputs.DefaultDarkRune,
                IconInputs.DefaultFlameCap,
                IconInputs.FlaskOfBlindingLight,
                IconInputs.FlaskOfPureDeath,
                IconInputs.FlaskOfSupremePower,
                IconInputs.AdeptsElixir,
                IconInputs.ElixirOfMajorFirePower,
                IconInputs.ElixirOfMajorFrostPower,
                IconInputs.ElixirOfMajorMageblood,
                IconInputs.ElixirOfDraenicWisdom,
                IconInputs.MainHandBrilliantWizardOil,
                IconInputs.MainHandSuperiorWizardOil,
                IconInputs.BlackenedBasilisk,
                IconInputs.SkullfishSoup,
                IconInputs.ScrollOfSpiritV,
                IconInputs.KreegsStoutBeatdown,
            ],
            // Inputs to include in the 'Rotation' section on the settings tab.
            rotationInputs: MageInputs.MageRotationConfig,
            // Inputs to include in the 'Other' section on the settings tab.
            otherInputs: {
                inputs: [
                    MageInputs.EvocationTicks,
                    OtherInputs.ShadowPriestDPS,
                    OtherInputs.StartingPotion,
                    OtherInputs.NumStartingPotions,
                    OtherInputs.SnapshotImprovedWrathOfAirTotem,
                ],
            },
            encounterPicker: {
                // Whether to include 'Target Armor' in the 'Encounter' section of the settings tab.
                showTargetArmor: false,
                // Whether to include 'Execute Duration (%)' in the 'Encounter' section of the settings tab.
                showExecuteProportion: true,
                // Whether to include 'Num Targets' in the 'Encounter' section of the settings tab.
                showNumTargets: false,
            },
            // If true, the talents on the talents tab will not be individually modifiable by the user.
            // Note that the use can still pick between preset talents, if there is more than 1.
            freezeTalents: false,
            presets: {
                // Preset talents that the user can quickly select.
                talents: [
                    Presets.ArcaneTalents,
                    Presets.FireTalents,
                    Presets.FrostTalents,
                    Presets.DeepFrostTalents,
                ],
                // Preset gear configurations that the user can quickly select.
                gear: [
                    Presets.P1_ARCANE_PRESET,
                    Presets.P2_ARCANE_PRESET,
                    Presets.P3_ARCANE_PRESET,
                    Presets.P1_FIRE_PRESET,
                    Presets.P2_FIRE_PRESET,
                    Presets.P3_FIRE_PRESET,
                    Presets.P1_FROST_PRESET,
                    Presets.P2_FROST_PRESET,
                    Presets.P3_FROST_PRESET,
                ],
            },
        });
    }
}
