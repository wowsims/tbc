import { Buffs } from '../core/api/common.js';
import { Class } from '../core/api/common.js';
import { Consumes } from '../core/api/common.js';
import { Encounter } from '../core/api/common.js';
import { ItemSlot } from '../core/api/common.js';
import { Spec } from '../core/api/common.js';
import { Stat } from '../core/api/common.js';
import { TristateEffect } from '../core/api/common.js'
import { Stats } from '../core/api/stats.js';
import { Sim } from '../core/sim.js';
import { DefaultTheme } from '../core/themes/default.js';

import { ElementalShaman, ElementalShaman_Agent as ElementalShamanAgent, ShamanTalents as ShamanTalents, ElementalShaman_Options as ElementalShamanOptions } from '../core/api/shaman.js';
import { ElementalShaman_Agent_AgentType as AgentType } from '../core/api/shaman.js';

import * as IconInputs from '../core/components/icon_inputs.js';
import * as OtherInputs from '../core/components/other_inputs.js';
import * as Gems from '../core/constants/gems.js';
import * as Tooltips from '../core/constants/tooltips.js';

import * as Presets from './presets.js';


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
    talents: Presets.StandardTalents,
    specOptions: ElementalShamanOptions.create({
      waterShield: true,
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
    encounters: [
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
