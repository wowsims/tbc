import { Spec } from '../core/api/newapi';
import { DefaultTheme } from '../core/themes/default';
import * as IconInputs from '../core/components/icon_inputs';

const theme = new DefaultTheme(
    document.body,
    Spec.ElementalShaman,
    {
      'Buffs': [
        IconInputs.ArcaneBrilliance,
        IconInputs.BlessingOfKings,
        IconInputs.BlessingOfWisdom,
        IconInputs.Bloodlust,
        IconInputs.GiftOfTheWild,
      ],
    });
theme.init();
