import { Buffs } from '../core/api/common.js';
import { Consumes } from '../core/api/common.js';
import { Encounter } from '../core/api/common.js';
import { EquipmentSpec } from '../core/api/common.js';
import { ItemSpec } from '../core/api/common.js';
import { Spec } from '../core/api/common.js';
import { Stat } from '../core/api/common.js';
import { TristateEffect } from '../core/api/common.js';
import { Stats } from '../core/api/stats.js';
import { DefaultTheme } from '../core/themes/default.js';
import { ElementalShaman_ElementalShamanAgent as ElementalShamanAgent, ElementalShaman_ElementalShamanOptions as ElementalShamanOptions } from '../core/api/shaman.js';
import { ElementalShaman_ElementalShamanAgent_AgentType as AgentType } from '../core/api/shaman.js';
import * as IconInputs from '../core/components/icon_inputs.js';
import * as OtherInputs from '../core/components/other_inputs.js';
import * as Enchants from '../core/constants/enchants.js';
import * as Gems from '../core/constants/gems.js';
import * as Tooltips from '../core/constants/tooltips.js';
const IconInputWaterShield = {
    id: { spellId: 33736 },
    states: 2,
    changedEvent: (sim) => sim.specOptionsChangeEmitter,
    getValue: (sim) => sim.getSpecOptions().waterShield,
    setBooleanValue: (sim, newValue) => {
        const newOptions = sim.getSpecOptions();
        newOptions.waterShield = newValue;
        sim.setSpecOptions(newOptions);
    },
};
const StandardTalentsString = '55030105100213351051--05105301005';
const ElementalShamanRotationConfig = [
    {
        type: 'enum',
        cssClass: 'rotation-enum-picker',
        config: {
            names: ['Adaptive', 'CL On Clearcast', 'Fixed LB+CL'],
            values: [AgentType.Adaptive, AgentType.CLOnClearcast, AgentType.FixedLBCL],
            changedEvent: (sim) => sim.agentChangeEmitter,
            getValue: (sim) => sim.getAgent().type,
            setValue: (sim, newValue) => {
                const newAgent = sim.getAgent();
                newAgent.type = newValue;
                sim.setAgent(newAgent);
            },
        },
    },
];
const theme = new DefaultTheme(document.body, {
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
        'Buffs': [
            IconInputWaterShield,
            IconInputs.WrathOfAirTotem,
            IconInputs.TotemOfWrath,
            IconInputs.ManaSpringTotem,
            IconInputs.ManaTideTotem,
            IconInputs.Bloodlust,
            IconInputs.DrumsOfBattle,
            IconInputs.DrumsOfRestoration,
            IconInputs.ArcaneBrilliance,
            IconInputs.DivineSpirit,
            IconInputs.BlessingOfKings,
            IconInputs.BlessingOfWisdom,
            IconInputs.GiftOfTheWild,
            IconInputs.MoonkinAura,
            IconInputs.EyeOfTheNight,
            IconInputs.ChainOfTheTwilightOwl,
            IconInputs.JadePendantOfBlasting,
            IconInputs.AtieshWarlock,
            IconInputs.AtieshMage,
        ],
        'Debuffs': [
            IconInputs.JudgementOfWisdom,
            IconInputs.ImprovedSealOfTheCrusader,
            IconInputs.Misery,
        ],
        'Consumes': [
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
    otherSections: {
        'Rotation': ElementalShamanRotationConfig,
        'Other': [
            OtherInputs.ShadowPriestDPS,
        ],
    },
    freezeTalents: true,
    showTargetArmor: false,
    showNumTargets: true,
    defaults: {
        phase: 2,
        epWeights: Stats.fromMap({
            [Stat.StatIntellect]: 0.41,
            [Stat.StatSpellPower]: 1,
            [Stat.StatNatureSpellPower]: 1,
            [Stat.StatSpellCrit]: 0.88,
            [Stat.StatSpellHaste]: 1.21,
            [Stat.StatMP5]: 0.37,
        }),
        encounter: Encounter.create({
            duration: 300,
            numTargets: 1,
        }),
        buffs: Buffs.create({
            bloodlust: 1,
            arcaneBrilliance: true,
            divineSpirit: TristateEffect.TristateEffectImproved,
            blessingOfKings: true,
            blessingOfWisdom: 2,
            giftOfTheWild: TristateEffect.TristateEffectImproved,
            judgementOfWisdom: true,
            misery: true,
            wrathOfAirTotem: TristateEffect.TristateEffectRegular,
            totemOfWrath: 1,
            manaSpringTotem: TristateEffect.TristateEffectRegular,
        }),
        consumes: Consumes.create({
            drumsOfBattle: true,
            superManaPotion: true,
        }),
        agent: ElementalShamanAgent.create({
            type: AgentType.Adaptive,
        }),
        talents: StandardTalentsString,
        specOptions: ElementalShamanOptions.create({
            waterShield: true,
        }),
    },
    presets: {
        gear: [
            {
                name: 'P1 BIS',
                tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
                equipment: EquipmentSpec.create({
                    items: [
                        ItemSpec.create({
                            id: 29035,
                            enchant: Enchants.GLYPH_OF_POWER,
                            gems: [
                                Gems.CHAOTIC_SKYFIRE_DIAMOND,
                                Gems.POTENT_NOBLE_TOPAZ,
                            ],
                        }),
                        ItemSpec.create({
                            id: 28762, // Adornment of Stolen Souls
                        }),
                        ItemSpec.create({
                            id: 29037,
                            enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                            gems: [
                                Gems.POTENT_NOBLE_TOPAZ,
                                Gems.POTENT_NOBLE_TOPAZ,
                            ],
                        }),
                        ItemSpec.create({
                            id: 28797, // Brute Cloak of the Ogre-Magi
                        }),
                        ItemSpec.create({
                            id: 29519,
                            enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                            gems: [
                                Gems.RUNED_LIVING_RUBY,
                                Gems.RUNED_LIVING_RUBY,
                                Gems.RUNED_LIVING_RUBY,
                            ],
                        }),
                        ItemSpec.create({
                            id: 29521,
                            enchant: Enchants.WRIST_SPELLPOWER,
                            gems: [
                                Gems.POTENT_NOBLE_TOPAZ,
                            ],
                        }),
                        ItemSpec.create({
                            id: 28780,
                            enchant: Enchants.GLOVES_SPELLPOWER,
                            gems: [
                                Gems.POTENT_NOBLE_TOPAZ,
                                Gems.GLOWING_NIGHTSEYE,
                            ],
                        }),
                        ItemSpec.create({
                            id: 29520,
                            gems: [
                                Gems.GLOWING_NIGHTSEYE,
                                Gems.POTENT_NOBLE_TOPAZ,
                            ],
                        }),
                        ItemSpec.create({
                            id: 24262,
                            enchant: Enchants.RUNIC_SPELLTHREAD,
                            gems: [
                                Gems.RUNED_LIVING_RUBY,
                                Gems.RUNED_LIVING_RUBY,
                                Gems.RUNED_LIVING_RUBY,
                            ],
                        }),
                        ItemSpec.create({
                            id: 28517,
                            gems: [
                                Gems.RUNED_LIVING_RUBY,
                                Gems.RUNED_LIVING_RUBY,
                            ],
                        }),
                        ItemSpec.create({
                            id: 30667,
                            enchant: Enchants.RING_SPELLPOWER,
                        }),
                        ItemSpec.create({
                            id: 28753,
                            enchant: Enchants.RING_SPELLPOWER,
                        }),
                        ItemSpec.create({
                            id: 29370, // Icon of the Silver Crescent
                        }),
                        ItemSpec.create({
                            id: 28785, // Lightning Capacitor
                        }),
                        ItemSpec.create({
                            id: 28770,
                            enchant: Enchants.WEAPON_SPELLPOWER,
                        }),
                        ItemSpec.create({
                            id: 29273, // Khadgar's Knapsack
                        }),
                        ItemSpec.create({
                            id: 28248, // Totem of the Void
                        }),
                    ],
                }),
            },
            {
                name: 'P2 BIS',
                tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
                equipment: EquipmentSpec.create({
                    items: [
                        ItemSpec.create({
                            id: 29035,
                            enchant: Enchants.GLYPH_OF_POWER,
                            gems: [
                                Gems.CHAOTIC_SKYFIRE_DIAMOND,
                                Gems.POTENT_NOBLE_TOPAZ,
                            ],
                        }),
                        ItemSpec.create({
                            id: 30015, // The Sun King's Talisman
                        }),
                        ItemSpec.create({
                            id: 29037,
                            enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                            gems: [
                                Gems.POTENT_NOBLE_TOPAZ,
                                Gems.POTENT_NOBLE_TOPAZ,
                            ],
                        }),
                        ItemSpec.create({
                            id: 28797, // Brute Cloak of the Ogre-Magi
                        }),
                        ItemSpec.create({
                            id: 30169,
                            enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                            gems: [
                                Gems.RUNED_LIVING_RUBY,
                                Gems.RUNED_LIVING_RUBY,
                                Gems.RUNED_LIVING_RUBY,
                            ],
                        }),
                        ItemSpec.create({
                            id: 29918,
                            enchant: Enchants.WRIST_SPELLPOWER,
                        }),
                        ItemSpec.create({
                            id: 28780,
                            enchant: Enchants.GLOVES_SPELLPOWER,
                            gems: [
                                Gems.POTENT_NOBLE_TOPAZ,
                                Gems.GLOWING_NIGHTSEYE,
                            ],
                        }),
                        ItemSpec.create({
                            id: 30038,
                            gems: [
                                Gems.GLOWING_NIGHTSEYE,
                                Gems.POTENT_NOBLE_TOPAZ,
                            ],
                        }),
                        ItemSpec.create({
                            id: 30172,
                            enchant: Enchants.RUNIC_SPELLTHREAD,
                            gems: [
                                Gems.POTENT_NOBLE_TOPAZ,
                            ],
                        }),
                        ItemSpec.create({
                            id: 30067, // Velvet Boots of the Guardian
                        }),
                        ItemSpec.create({
                            id: 30667,
                            enchant: Enchants.RING_SPELLPOWER,
                        }),
                        ItemSpec.create({
                            id: 30109,
                            enchant: Enchants.RING_SPELLPOWER,
                        }),
                        ItemSpec.create({
                            id: 29370, // Icon of the Silver Crescent
                        }),
                        ItemSpec.create({
                            id: 28785, // Lightning Capacitor
                        }),
                        ItemSpec.create({
                            id: 29988,
                            enchant: Enchants.WEAPON_SPELLPOWER,
                        }),
                        ItemSpec.create({
                            id: 28248, // Totem of the Void
                        }),
                    ],
                }),
            },
        ],
        encounters: [],
        talents: [
            {
                name: 'Standard',
                talents: StandardTalentsString,
            },
        ],
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
