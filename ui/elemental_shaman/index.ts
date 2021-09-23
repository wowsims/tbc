import { Buffs } from '../core/api/common';
import { Class } from '../core/api/common';
import { Consumes } from '../core/api/common';
import { Encounter } from '../core/api/common';
import { EquipmentSpec } from '../core/api/common';
import { ItemSlot } from '../core/api/common';
import { ItemSpec } from '../core/api/common';
import { Spec } from '../core/api/common';
import { Stat } from '../core/api/common';
import { TristateEffect } from '../core/api/common'
import { Sim } from '../core/sim';
import { DefaultTheme } from '../core/themes/default';

import { ElementalShaman, ElementalShaman_ElementalShamanAgent as ElementalShamanAgent, ShamanTalents as ShamanTalents, ElementalShaman_ElementalShamanOptions as ElementalShamanOptions } from '../core/api/shaman';
import { ElementalShaman_ElementalShamanAgent_AgentType as AgentType } from '../core/api/shaman';

import * as IconInputs from '../core/components/icon_inputs';
import * as OtherInputs from '../core/components/other_inputs';
import * as Enchants from '../core/constants/enchants';
import * as Gems from '../core/constants/gems';
import * as Tooltips from '../core/constants/tooltips';


const IconInputWaterShield = {
  id: { spellId: 33736 },
  states: 2,
  changedEvent: (sim: Sim<Spec.SpecElementalShaman>) => sim.specOptionsChangeEmitter,
  getValue: (sim: Sim<Spec.SpecElementalShaman>) => sim.getSpecOptions().waterShield,
  setBooleanValue: (sim: Sim<Spec.SpecElementalShaman>, newValue: boolean) => {
    const newOptions = sim.getSpecOptions();
    newOptions.waterShield = newValue;
    sim.setSpecOptions(newOptions);
  },
};

const StandardTalentsString = '55030105100213351051--05105301005';

const ElementalShamanRotationConfig = [
  {
    type: 'enum' as const,
    cssClass: 'rotation-enum-picker',
    config: {
      names: ['Adaptive', 'CL On Clearcast', 'Fixed LB+CL'],
      values: [AgentType.Adaptive, AgentType.CLOnClearcast, AgentType.FixedLBCL],
      changedEvent: (sim: Sim<Spec.SpecElementalShaman>) => sim.agentChangeEmitter,
      getValue: (sim: Sim<Spec.SpecElementalShaman>) => sim.getAgent().type,
      setValue: (sim: Sim<Spec.SpecElementalShaman>, newValue: number) => {
        const newAgent = sim.getAgent();
        newAgent.type = newValue;
        sim.setAgent(newAgent);
      },
    },
  },
];

const theme = new DefaultTheme<Spec.SpecElementalShaman>(document.body, {
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
  showTargetArmor: false,
  showNumTargets: true,
  defaults: {
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
              id: 29035, // Cyclone Faceguard
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
              id: 29037, // Cyclone Shoulderguards
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
              id: 29519, // Netherstrike Breastplate
              enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
              gems: [
                Gems.RUNED_LIVING_RUBY,
                Gems.RUNED_LIVING_RUBY,
                Gems.RUNED_LIVING_RUBY,
              ],
            }),
            ItemSpec.create({
              id: 29521, // Netherstrike Bracers
              enchant: Enchants.WRIST_SPELLPOWER,
              gems: [
                Gems.POTENT_NOBLE_TOPAZ,
              ],
            }),
            ItemSpec.create({
              id: 28780, // Soul-Eaters's Handwraps
              enchant: Enchants.GLOVES_SPELLPOWER,
              gems: [
                Gems.POTENT_NOBLE_TOPAZ,
                Gems.GLOWING_NIGHTSEYE,
              ],
            }),
            ItemSpec.create({
              id: 29520, // Netherstrike Belt
              gems: [
                Gems.GLOWING_NIGHTSEYE,
                Gems.POTENT_NOBLE_TOPAZ,
              ],
            }),
            ItemSpec.create({
              id: 24262, // Spellstrike Pants
              enchant: Enchants.RUNIC_SPELLTHREAD,
              gems: [
                Gems.RUNED_LIVING_RUBY,
                Gems.RUNED_LIVING_RUBY,
                Gems.RUNED_LIVING_RUBY,
              ],
            }),
            ItemSpec.create({
              id: 28517, // Boots of Foretelling
              gems: [
                Gems.RUNED_LIVING_RUBY,
                Gems.RUNED_LIVING_RUBY,
              ],
            }),
            ItemSpec.create({
              id: 30667, // Ring of Unrelenting Storms
              enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
              id: 28753, // Ring of Recurrence
              enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
              id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
              id: 28785, // Lightning Capacitor
            }),
            ItemSpec.create({
              id: 28770, // Nathrezim Mindblade
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
    ],
    encounters: [
    ],
    talents: [
      {
        name: 'Standard',
        talents: StandardTalentsString,
      },
    ],
  },
});
theme.init();
