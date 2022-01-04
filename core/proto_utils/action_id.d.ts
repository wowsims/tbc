import { ActionID as ActionIdProto } from '/tbc/core/proto/common.js';
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
export declare type RawActionId = ItemId | SpellId | OtherId;
export declare type ActionId = {
    id: RawActionId;
    tag?: number;
};
export declare function sameActionId(id1: ActionId, id2: ActionId): boolean;
export declare function actionIdToString(id: ActionId): string;
export declare function actionIdToProto(actionId: ActionId): ActionIdProto;
export declare function protoToActionId(protoId: ActionIdProto): ActionId;
