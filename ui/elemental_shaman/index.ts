import { GemColor } from '../core/api/newapi';
import { Item } from '../core/api/newapi';
import { ItemQuality } from '../core/api/newapi';
import { ItemSlot } from '../core/api/newapi';
import { ItemType } from '../core/api/newapi';
import { Spec } from '../core/api/newapi';
import { EquippedItem } from '../core/equipped_item';
import { DefaultTheme } from '../core/themes/default';

const theme = new DefaultTheme(document.body, Spec.ElementalShaman)
theme.init();
