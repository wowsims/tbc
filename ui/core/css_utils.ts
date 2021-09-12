import { GemColor } from './api/newapi';
import { ItemQuality } from './api/newapi';

const gemSocketCssClasses: Partial<Record<GemColor, string>> = {
  [GemColor.GemColorBlue]: 'socket-color-blue',
  [GemColor.GemColorMeta]: 'socket-color-meta',
  [GemColor.GemColorRed]: 'socket-color-red',
  [GemColor.GemColorYellow]: 'socket-color-yellow',
};
export function SetGemSocketCssClass(elem: HTMLElement, color: GemColor) {
  Object.values(gemSocketCssClasses).forEach(cssClass => elem.classList.remove(cssClass));

  if (gemSocketCssClasses[color]) {
    elem.classList.add(gemSocketCssClasses[color] as string);
    return;
  }

  throw new Error('No css class for gem socket color: ' + color);
}

const itemQualityCssClasses: Record<ItemQuality, string> = {
  [ItemQuality.ItemQualityJunk]: 'item-quality-junk',
  [ItemQuality.ItemQualityUncommon]: 'item-quality-uncommon',
  [ItemQuality.ItemQualityRare]: 'item-quality-rare',
  [ItemQuality.ItemQualityEpic]: 'item-quality-epic',
  [ItemQuality.ItemQualityLegendary]: 'item-quality-legendary',
};
export function SetItemQualityCssClass(elem: HTMLElement, quality: ItemQuality | null) {
  Object.values(itemQualityCssClasses).forEach(cssClass => elem.classList.remove(cssClass));

  if (quality) {
    elem.classList.add(itemQualityCssClasses[quality]);
  }
}
