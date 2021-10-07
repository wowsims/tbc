import { Buffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { DefaultTheme } from '/tbc/core/themes/default.js';
import { ElementalShaman_Agent as ElementalShamanAgent, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';
import { ElementalShaman_Agent_AgentType as AgentType } from '/tbc/core/proto/shaman.js';
import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as Gems from '/tbc/core/constants/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
import * as ShamanInputs from './inputs.js';
import * as Presets from './presets.js';
const theme = new DefaultTheme(document.body, {
    releaseStatus: 'Beta',
    knownIssues: [
        'Can only use 1 type of potion, cannot use 1 Destruction Potion and then Super Mana Potions after that.',
        'Detailed results labels are numbers instead of names, and LO procs are missing.',
    ],
    spec: Spec.SpecElementalShaman,
    epStats: [
        Stat.StatIntellect,
        Stat.StatSpellPower,
        Stat.StatNatureSpellPower,
        Stat.StatSpellHit,
        Stat.StatSpellCrit,
        Stat.StatSpellHaste,
        Stat.StatMP5,
    ],
    epReferenceStat: Stat.StatSpellPower,
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
    iconSections: {
        'Self Buffs': {
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
        'Other Buffs': {
            tooltip: Tooltips.OTHER_BUFFS_SECTION,
            icons: [
                IconInputs.DrumsOfBattleBuff,
                IconInputs.DrumsOfRestorationBuff,
                IconInputs.ArcaneBrilliance,
                IconInputs.DivineSpirit,
                IconInputs.BlessingOfKings,
                IconInputs.BlessingOfWisdom,
                IconInputs.GiftOfTheWild,
                IconInputs.MoonkinAura,
                IconInputs.Bloodlust,
                IconInputs.WrathOfAirTotem,
                IconInputs.TotemOfWrath,
                IconInputs.ManaSpringTotem,
                IconInputs.ManaTideTotem,
                IconInputs.EyeOfTheNight,
                IconInputs.ChainOfTheTwilightOwl,
                IconInputs.JadePendantOfBlasting,
                IconInputs.AtieshWarlock,
                IconInputs.AtieshMage,
            ],
        },
        'Debuffs': {
            icons: [
                IconInputs.JudgementOfWisdom,
                IconInputs.ImprovedSealOfTheCrusader,
                IconInputs.Misery,
            ],
        },
        'Consumes': {
            icons: [
                IconInputs.SuperManaPotion,
                IconInputs.DestructionPotion,
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
    },
    otherSections: {
        'Rotation': ShamanInputs.ElementalShamanRotationConfig,
        'Other': {
            inputs: [
                OtherInputs.ShadowPriestDPS,
            ],
        },
    },
    freezeTalents: true,
    showTargetArmor: false,
    showNumTargets: true,
    defaults: {
        phase: 2,
        gear: Presets.PRERAID_GEAR,
        epWeights: Stats.fromMap({
            [Stat.StatIntellect]: 0.33,
            [Stat.StatSpellPower]: 1,
            [Stat.StatNatureSpellPower]: 1,
            [Stat.StatSpellCrit]: 0.78,
            [Stat.StatSpellHaste]: 1.25,
            [Stat.StatMP5]: 0.08,
        }),
        encounter: Encounter.create({
            duration: 300,
            numTargets: 1,
        }),
        buffs: Buffs.create({
            bloodlust: 0,
            arcaneBrilliance: true,
            divineSpirit: TristateEffect.TristateEffectImproved,
            blessingOfKings: true,
            blessingOfWisdom: 2,
            giftOfTheWild: TristateEffect.TristateEffectImproved,
            judgementOfWisdom: true,
            misery: true,
        }),
        consumes: Consumes.create({
            drumsOfBattle: true,
            superManaPotion: true,
        }),
        agent: ElementalShamanAgent.create({
            type: AgentType.Adaptive,
        }),
        talents: Presets.StandardTalents,
        specOptions: ElementalShamanOptions.create({
            waterShield: true,
            bloodlust: true,
            totemOfWrath: true,
            manaSpringTotem: true,
            wrathOfAirTotem: true,
        }),
    },
    presets: {
        talents: [
            {
                name: 'Standard',
                talents: Presets.StandardTalents,
            },
        ],
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
        encounters: [],
    },
    metaGemEffectEP: (gem, sim) => {
        if (gem.id == Gems.CHAOTIC_SKYFIRE_DIAMOND) {
            const finalStats = new Stats(sim.getCurrentStats().finalStats);
            return (((finalStats.getStat(Stat.StatSpellPower) * 0.795) + 603) * 2 * (finalStats.getStat(Stat.StatSpellCrit) / 2208) * 0.045) / 0.795;
        }
        return 0;
    },
});
theme.init();
