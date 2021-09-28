import { GemColor } from './api/common.js';
import { ItemQuality } from './api/common.js';
const gemSocketCssClasses = {
    [GemColor.GemColorBlue]: 'socket-color-blue',
    [GemColor.GemColorMeta]: 'socket-color-meta',
    [GemColor.GemColorRed]: 'socket-color-red',
    [GemColor.GemColorYellow]: 'socket-color-yellow',
};
export function setGemSocketCssClass(elem, color) {
    Object.values(gemSocketCssClasses).forEach(cssClass => elem.classList.remove(cssClass));
    if (gemSocketCssClasses[color]) {
        elem.classList.add(gemSocketCssClasses[color]);
        return;
    }
    throw new Error('No css class for gem socket color: ' + color);
}
const itemQualityCssClasses = {
    [ItemQuality.ItemQualityJunk]: 'item-quality-junk',
    [ItemQuality.ItemQualityCommon]: 'item-quality-common',
    [ItemQuality.ItemQualityUncommon]: 'item-quality-uncommon',
    [ItemQuality.ItemQualityRare]: 'item-quality-rare',
    [ItemQuality.ItemQualityEpic]: 'item-quality-epic',
    [ItemQuality.ItemQualityLegendary]: 'item-quality-legendary',
};
export function setItemQualityCssClass(elem, quality) {
    Object.values(itemQualityCssClasses).forEach(cssClass => elem.classList.remove(cssClass));
    if (quality) {
        elem.classList.add(itemQualityCssClasses[quality]);
    }
}
