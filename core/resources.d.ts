import { GemColor } from './proto/common.js';
import { Item } from './proto/common.js';
import { ItemSlot } from './proto/common.js';
export declare const repoName = "tbc";
export declare const specDirectory: string;
export declare function getEmptySlotIconUrl(slot: ItemSlot): string;
export declare type ItemId = {
    itemId: number;
};
export declare type SpellId = {
    spellId: number;
};
export declare type OtherId = {
    otherId: number;
};
export declare type ItemOrSpellId = ItemId | SpellId;
export declare type ActionId = ItemId | SpellId | OtherId;
export declare function getTooltipData(id: ItemOrSpellId): Promise<any>;
export declare function getIconUrl(id: ActionId): Promise<string>;
export declare function getItemIconUrl(item: Item): Promise<string>;
export declare function getName(id: ItemOrSpellId | ActionId): Promise<string>;
export declare function setWowheadHref(elem: HTMLAnchorElement, id: ItemOrSpellId): void;
export declare function setWowheadItemHref(elem: HTMLAnchorElement, item: Item): void;
export declare function getEmptyGemSocketIconUrl(color: GemColor): string;
