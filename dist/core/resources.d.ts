import { GemColor } from './api/common.js';
import { ItemSlot } from './api/common.js';
export declare function getEmptySlotIconUrl(slot: ItemSlot): string;
export declare type IconId = {
    itemId: number;
};
export declare type SpellId = {
    spellId: number;
};
export declare type ItemOrSpellId = IconId | SpellId;
export declare function getIconUrl(id: ItemOrSpellId): Promise<string>;
export declare function setWowheadHref(elem: HTMLAnchorElement, id: ItemOrSpellId): void;
export declare function getEmptyGemSocketIconUrl(color: GemColor): string;
