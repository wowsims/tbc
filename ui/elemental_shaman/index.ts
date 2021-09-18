import { Buffs } from '../core/api/newapi';
import { Class } from '../core/api/newapi';
import { Consumes } from '../core/api/newapi';
import { Encounter } from '../core/api/newapi';
import { EquipmentSpec } from '../core/api/newapi';
import { ItemSlot } from '../core/api/newapi';
import { ItemSpec } from '../core/api/newapi';
import { Spec } from '../core/api/newapi';
import { Stat } from '../core/api/newapi';
import { TristateEffect } from '../core/api/newapi'
import { Sim } from '../core/sim';
import { DefaultTheme } from '../core/themes/default';

import { Shaman, Shaman_ShamanAgent as ShamanAgent, Shaman_ShamanTalents as ShamanTalents, Shaman_ShamanOptions as ShamanOptions } from '../core/api/newapi';

import * as IconInputs from '../core/components/icon_inputs';
import * as Enchants from '../core/constants/enchants';
import * as Gems from '../core/constants/gems';
import * as Tooltips from '../core/constants/tooltips';


const theme = new DefaultTheme<Class.ClassShaman>(document.body, {
  spec: Spec.ElementalShaman,
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
    Stat.StatSpellHit,
    Stat.StatSpellCrit,
    Stat.StatSpellHaste,
    Stat.StatMP5,
  ],
  iconSections: {
    'Buffs': [
      {
        id: { spellId: 33736 },
        states: 2,
        changedEvent: (sim: Sim<Class.ClassShaman>) => sim.classOptionsChangeEmitter,
        getValue: (sim: Sim<Class.ClassShaman>) => sim.getClassOptions().waterShield,
        setBooleanValue: (sim: Sim<Class.ClassShaman>, newValue: boolean) => {
          const newOptions = sim.getClassOptions();
          newOptions.waterShield = newValue;
          sim.setClassOptions(newOptions);
        },
      },
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
    classOptions: ShamanOptions.create({
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
        talents: '55030105100213351051--05105301005',
      },
    ],
  },
});
theme.init();
