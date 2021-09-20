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
