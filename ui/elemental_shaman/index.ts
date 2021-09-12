import { GemColor } from '../core/api/newapi';
import { Item } from '../core/api/newapi';
import { ItemQuality } from '../core/api/newapi';
import { ItemSlot } from '../core/api/newapi';
import { ItemType } from '../core/api/newapi';
import { Spec } from '../core/api/newapi';
import { EquippedItem } from '../core/equipped_item';
import { DefaultTheme } from '../core/themes/default';

const theme = new DefaultTheme(document.body, Spec.ElementalShaman);
theme.init().then(() => {
  theme.sim.equipItem(ItemSlot.ItemSlotHead, new EquippedItem(Item.create({
    id: 32235,
    type: ItemType.ItemTypeHead,
    name: 'Cursed Vision of Sargeras',
    quality: ItemQuality.ItemQualityEpic,
    gemSockets: [GemColor.GemColorMeta, GemColor.GemColorYellow, GemColor.GemColorRed],
  })));
});
