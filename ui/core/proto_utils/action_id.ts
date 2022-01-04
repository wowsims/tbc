import { ActionID as ActionIdProto } from '/tbc/core/proto/common.js';

export type ItemId = {
  itemId: number;
};
export type SpellId = {
  spellId: number;
};
export type OtherId = {
  otherId: number;
};
export type ItemOrSpellId = ItemId | SpellId;
export type RawActionId = ItemId | SpellId | OtherId;
export type ActionId = {
	id: RawActionId,
	tag?: number,
};

export function sameActionId(id1: ActionId, id2: ActionId): boolean {
	return ((('itemId' in id1.id && 'itemId' in id2.id && id1.id.itemId == id2.id.itemId)
					|| ('spellId' in id1.id && 'spellId' in id2.id && id1.id.spellId == id2.id.spellId)
					|| ('otherId' in id1.id && 'otherId' in id2.id && id1.id.otherId == id2.id.otherId))
					&& id1.tag == id2.tag);
}

export function actionIdToString(id: ActionId): string {
	let tagStr = id.tag ? ('-' + id.tag) : '';

	if ('itemId' in id.id) {
		return 'item-' + id.id.itemId + tagStr;
	} else if ('spellId' in id.id) {
		return 'spell-' + id.id.spellId + tagStr;
	} else if ('otherId' in id.id) {
		return 'other-' + id.id.otherId + tagStr;
	} else {
		throw new Error('Invalid Action Id: ' + JSON.stringify(id));
	}
}

export function actionIdToProto(actionId: ActionId): ActionIdProto {
	const protoId = ActionIdProto.create({
		tag: actionId.tag,
	});

	if ('itemId' in actionId.id) {
		protoId.rawId = {
			oneofKind: 'itemId',
			itemId: actionId.id.itemId,
		};
	} else if ('spellId' in actionId.id) {
		protoId.rawId = {
			oneofKind: 'spellId',
			spellId: actionId.id.spellId,
		};
	} else if ('otherId' in actionId.id) {
		protoId.rawId = {
			oneofKind: 'otherId',
			otherId: actionId.id.otherId,
		};
	}

	return protoId;
}

export function protoToActionId(protoId: ActionIdProto): ActionId {
	if (protoId.rawId.oneofKind == 'spellId') {
		return {
			id: {
				spellId: protoId.rawId.spellId,
			},
			tag: protoId.tag,
		};
	} else if (protoId.rawId.oneofKind == 'itemId') {
		return {
			id: {
				itemId: protoId.rawId.itemId,
			},
			tag: protoId.tag,
		};
	} else if (protoId.rawId.oneofKind == 'otherId') {
		return {
			id: {
				otherId: protoId.rawId.otherId,
			},
			tag: protoId.tag,
		};
	} else {
		return {
			id: {
				otherId: 0,
			},
		};
	}
}
