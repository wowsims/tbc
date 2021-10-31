import { Buffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { DefaultTheme } from '/tbc/core/themes/default.js';
import { ElementalShaman_Rotation as ElementalShamanRotation, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';
import { ElementalShaman_Rotation_RotationType as RotationType } from '/tbc/core/proto/shaman.js';
import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as Gems from '/tbc/core/constants/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
import * as ShamanInputs from './inputs.js';
import * as Presets from './presets.js';
const theme = new DefaultTheme(document.body, {
    // Can be 'Alpha', 'Beta', or 'Live'. Just adds a postfix to the generated title.
    releaseStatus: 'Beta',
    // List any known bugs / issues here and they'll be shown on the site.
    knownIssues: [],
    player: {
        spec: Spec.SpecElementalShaman,
        // All stats for which EP should be calculated.
        epStats: [
            Stat.StatIntellect,
            Stat.StatSpellPower,
            Stat.StatNatureSpellPower,
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
            Stat.StatSpellPower,
            Stat.StatNatureSpellPower,
            Stat.StatSpellHit,
            Stat.StatSpellCrit,
            Stat.StatSpellHaste,
            Stat.StatMP5,
        ],
        defaults: {
            // Default equipped gear.
            gear: Presets.PRERAID_GEAR,
            // Default EP weights for sorting gear in the gear picker.
            epWeights: Stats.fromMap({
                [Stat.StatIntellect]: 0.33,
                [Stat.StatSpellPower]: 1,
                [Stat.StatNatureSpellPower]: 1,
                [Stat.StatSpellCrit]: 0.78,
                [Stat.StatSpellHaste]: 1.25,
                [Stat.StatMP5]: 0.08,
            }),
            // Default consumes settings.
            consumes: Consumes.create({
                drums: Drums.DrumsOfBattle,
                defaultPotion: Potions.SuperManaPotion,
            }),
            // Default rotation settings.
            rotation: ElementalShamanRotation.create({
                type: RotationType.Adaptive,
            }),
            // Default talents.
            talents: Presets.StandardTalents,
            // Default spec-specific settings.
            specOptions: ElementalShamanOptions.create({
                waterShield: true,
                bloodlust: true,
                totemOfWrath: true,
                manaSpringTotem: true,
                wrathOfAirTotem: true,
            }),
        },
        // Custom function for determining the EP value of meta gem effects.
        // Default meta effect EP value is 0, so just handle the ones relevant to your spec.
        metaGemEffectEP: (gem, player) => {
            if (gem.id == Gems.CHAOTIC_SKYFIRE_DIAMOND) {
                const finalStats = new Stats(player.getCurrentStats().finalStats);
                return (((finalStats.getStat(Stat.StatSpellPower) * 0.795) + 603) * 2 * (finalStats.getStat(Stat.StatSpellCrit) / 2208) * 0.045) / 0.795;
            }
            return 0;
        },
    },
    sim: {
        defaults: {
            // TBC Release Phase, i.e. Black Temple is phase 3.
            phase: 2,
            // Default encounter settings.
            encounter: Encounter.create({
                duration: 300,
            }),
            // Default raid/party buffs settings.
            buffs: Buffs.create({
                bloodlust: 0,
                arcaneBrilliance: true,
                divineSpirit: TristateEffect.TristateEffectImproved,
                blessingOfKings: true,
                blessingOfWisdom: 2,
                giftOfTheWild: TristateEffect.TristateEffectImproved,
            }),
        },
    },
    target: {
        defaults: {
            armor: 0,
            debuffs: Debuffs.create({
                judgementOfWisdom: true,
                misery: true,
            }),
        },
    },
    // IconInputs to include in the 'Self Buffs' section on the settings tab.
    selfBuffInputs: {
        tooltip: Tooltips.SELF_BUFFS_SECTION,
        icons: [
            ShamanInputs.IconWaterShield,
            ShamanInputs.IconBloodlust,
            ShamanInputs.IconWrathOfAirTotem,
            ShamanInputs.IconTotemOfWrath,
            ShamanInputs.IconManaSpringTotem,
            IconInputs.DrumsOfBattleConsume,
            IconInputs.DrumsOfRestorationConsume,
        ],
    },
    // IconInputs to include in the 'Other Buffs' section on the settings tab.
    buffInputs: {
        tooltip: Tooltips.OTHER_BUFFS_SECTION,
        icons: [
            IconInputs.ArcaneBrilliance,
            IconInputs.DivineSpirit,
            IconInputs.BlessingOfKings,
            IconInputs.BlessingOfWisdom,
            IconInputs.GiftOfTheWild,
            IconInputs.MoonkinAura,
            IconInputs.DrumsOfBattleBuff,
            IconInputs.DrumsOfRestorationBuff,
            IconInputs.Bloodlust,
            IconInputs.WrathOfAirTotem,
            IconInputs.TotemOfWrath,
            IconInputs.ManaSpringTotem,
            IconInputs.EyeOfTheNight,
            IconInputs.ChainOfTheTwilightOwl,
            IconInputs.JadePendantOfBlasting,
            IconInputs.AtieshWarlock,
            IconInputs.AtieshMage,
        ],
    },
    // IconInputs to include in the 'Debuffs' section on the settings tab.
    debuffInputs: {
        icons: [
            IconInputs.JudgementOfWisdom,
            IconInputs.ImprovedSealOfTheCrusader,
            IconInputs.Misery,
        ],
    },
    // IconInputs to include in the 'Consumes' section on the settings tab.
    consumeInputs: {
        icons: [
            IconInputs.DefaultSuperManaPotion,
            IconInputs.DefaultDestructionPotion,
            IconInputs.DarkRune,
            IconInputs.FlaskOfBlindingLight,
            IconInputs.FlaskOfSupremePower,
            IconInputs.AdeptsElixir,
            IconInputs.ElixirOfMajorMageblood,
            IconInputs.ElixirOfDraenicWisdom,
            IconInputs.BrilliantWizardOil,
            IconInputs.SuperiorWizardOil,
            IconInputs.BlackenedBasilisk,
            IconInputs.SkullfishSoup,
        ],
    },
    // Inputs to include in the 'Rotation' section on the settings tab.
    rotationInputs: ShamanInputs.ElementalShamanRotationConfig,
    // Inputs to include in the 'Other' section on the settings tab.
    otherInputs: {
        inputs: [
            OtherInputs.ShadowPriestDPS,
            OtherInputs.StartingPotion,
            OtherInputs.NumStartingPotions,
        ],
    },
    // If true, the talents on the talents tab will not be individually modifiable by the user.
    // Note that the use can still pick between preset talents, if there is more than 1.
    freezeTalents: true,
    // Whether to include 'Target Armor' in the 'Encounter' section of the settings tab.
    showTargetArmor: false,
    // Whether to include 'Num Targets' in the 'Encounter' section of the settings tab.
    showNumTargets: true,
    presets: {
        // Preset talents that the user can quickly select.
        talents: [
            {
                name: 'Standard',
                talents: Presets.StandardTalents,
            },
        ],
        // Preset gear configurations that the user can quickly select.
        gear: [
            {
                name: 'P1 BIS',
                tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
                equipment: Presets.P1_BIS,
            },
            {
                name: 'P2 BIS',
                tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
                equipment: Presets.P2_BIS,
            },
        ],
        // Preset encounter settings that the user can quickly select.
        encounters: [],
    },
});
theme.init();
