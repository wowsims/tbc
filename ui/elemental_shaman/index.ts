import { Buffs } from '../core/api/newapi';
import { Consumes } from '../core/api/newapi';
import { Encounter } from '../core/api/newapi';
import { Spec } from '../core/api/newapi';
import { Stat } from '../core/api/newapi';
import { DefaultTheme } from '../core/themes/default';
import * as IconInputs from '../core/components/icon_inputs';

const theme = new DefaultTheme(document.body, {
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
      IconInputs.Bloodlust,
      IconInputs.DrumsOfBattle,
      IconInputs.DrumsOfRestoration,
      IconInputs.ArcaneBrilliance,
      IconInputs.ImprovedDivineSpirit,
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
      improvedDivineSpirit: true,
      blessingOfKings: true,
      blessingOfWisdom: 2,
      giftOfTheWild: true,

      judgementOfWisdom: true,
      misery: true,
    }),
    consumes: Consumes.create({
      drumsOfBattle: true,
      superManaPotion: true,
    }),
  },
});
theme.init();
