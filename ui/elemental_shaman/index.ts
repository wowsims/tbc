import { Spec } from '../core/api/newapi';
import { DefaultTheme } from '../core/themes/default';
import * as IconInputs from '../core/components/icon_inputs';

const theme = new DefaultTheme(
    document.body,
    Spec.ElementalShaman,
    {
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
    });
theme.init();
