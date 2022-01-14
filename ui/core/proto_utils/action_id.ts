import { ActionID as ActionIdProto } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { OtherAction } from '/tbc/core/proto/common.js';
import { getWowheadItemId } from '/tbc/core/proto_utils/equipped_item.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';

// Uniquely identifies a specific item / spell / thing in WoW. This object is immutable.
export class ActionId {
	readonly itemId: number;
	readonly spellId: number;
	readonly otherId: OtherAction;
	readonly tag: number;

	readonly baseName: string; // The name without any tag additions.
	readonly name: string;
	readonly iconUrl: string;

	private constructor(itemId: number, spellId: number, otherId: OtherAction, tag: number, baseName: string, name: string, iconUrl: string) {
		this.itemId = itemId;
		this.spellId = spellId;
		this.otherId = otherId;
		this.tag = tag;

		switch (otherId) {
			case OtherAction.OtherActionNone:
				break;
			case OtherAction.OtherActionWait:
				baseName = 'Wait';
				break;
			case OtherAction.OtherActionManaRegen:
				baseName = 'Regen';
				// Tag is number of milliseconds worth of regen.
				if (tag) {
					name = (tag/1000).toFixed(3) + 's ' + baseName;
				}
				break;
			case OtherAction.OtherActionAttack:
				name = 'Attack';
				iconUrl = 'https://wow.zamimg.com/images/wow/icons/large/inv_sword_04.jpg';
				if (tag == 1) {
					name += ' (Main Hand)';
				} else if (tag == 2) {
					name += ' (Off Hand)';
				}
				break;
		}
		this.baseName = baseName;
		this.name = name;
		this.iconUrl = iconUrl;
	}

	equals(other: ActionId): boolean {
		return this.equalsIgnoringTag(other) && this.tag == other.tag;
	}

	equalsIgnoringTag(other: ActionId): boolean {
		return (
				this.itemId == other.itemId
				&& this.spellId == other.spellId
				&& this.otherId == other.otherId);
	}

	setBackground(elem: HTMLElement) {
		if (this.iconUrl) {
			elem.style.backgroundImage = `url('${this.iconUrl}')`;
		}
	}

	setWowheadHref(elem: HTMLAnchorElement) {
		if (this.itemId) {
			elem.href = 'https://tbc.wowhead.com/item=' + this.itemId;
		} else if (this.spellId) {
			elem.href = 'https://tbc.wowhead.com/spell=' + this.spellId;
		}
	}

	setBackgroundAndHref(elem: HTMLAnchorElement) {
		this.setBackground(elem);
		this.setWowheadHref(elem);
	}

	async fillAndSet(elem: HTMLAnchorElement, setHref: boolean, setBackground: boolean): Promise<ActionId> {
		const filled = await this.fill();
		if (setHref) {
			filled.setWowheadHref(elem);
		}
		if (setBackground) {
			filled.setBackground(elem);
		}
		return filled;
	}

	// Returns an ActionId with the name and iconUrl fields filled.
	// playerIndex is the optional index of the player to whom this ID corresponds.
	async fill(playerIndex?: number): Promise<ActionId> {
		if (this.name || this.iconUrl) {
			return this;
		}

		if (this.otherId) {
			return this;
		}

		const tooltipData = await ActionId.getTooltipData(this);

		const baseName = tooltipData['name'];
		let name = baseName;
		switch (baseName) {
			case 'Arcane Blast':
				if (this.tag == 1) {
					name += ' (No Stacks)';
				} else if (this.tag == 2) {
					name += ` (1 Stack)`;
				} else if (this.tag > 2) {
					name += ` (${this.tag - 1} Stacks)`;
				}
				break;
			case 'Fireball':
			case 'Pyroblast':
				if (this.tag) name += ' (DoT)';
				break;
			case 'Mind Flay':
				if (this.tag == 1) {
					name += ' (1 Tick)';
				} else if (this.tag == 2) {
					name += ' (2 Tick)';
				} else if (this.tag == 3) {
					name += ' (3 Tick)';
				}
				break;
			case 'Lightning Bolt':
				if (this.tag) name += ' (LO)';
				break;
			// For targetted buffs, tag is the source player's raid index or -1 if none.
			case 'Bloodlust':
			case 'Innervate':
			case 'Mana Tide Totem':
			case 'Power Infusion':
				if (this.tag != NO_TARGET) {
					if (this.tag === playerIndex) {
						name += ` (self)`;
					} else {
						name += ` (from #${this.tag+1})`;
					}
				}
				break;
			case 'Darkmoon Card: Crusade':
				if (this.tag == 1) {
					name += ' (Melee)';
				} else if (this.tag == 2) {
					name += ' (Spell)';
				}
				break;
			default:
				if (this.tag == 10) {
					name += ' (Auto)';
				} else if (this.tag == 11) {
					name += ' (Offhand Auto)';
				} else if (this.tag) {
					name += ' (??)';
				}
				break;
		}

		const idString = this.toProtoString();
		const iconOverrideId = idOverrides[idString] || null;

		let iconUrl = "https://wow.zamimg.com/images/wow/icons/large/" + tooltipData['icon'] + ".jpg";
		if (iconOverrideId) {
			const overrideTooltipData = await ActionId.getTooltipData(iconOverrideId);
			iconUrl = "https://wow.zamimg.com/images/wow/icons/large/" + overrideTooltipData['icon'] + ".jpg";
		}

		return new ActionId(this.itemId, this.spellId, this.otherId, this.tag, baseName, name, iconUrl);
	}

	toString(): string {
		return this.toStringIgnoringTag() + (this.tag ? ('-' + this.tag) : '');
	}

	toStringIgnoringTag(): string {
		if (this.itemId) {
			return 'item-' + this.itemId;
		} else if (this.spellId) {
			return 'spell-' + this.spellId;
		} else if (this.otherId) {
			return 'other-' + this.otherId;
		} else {
			throw new Error('Empty action id!');
		}
	}

	toProto(): ActionIdProto {
		const protoId = ActionIdProto.create({
			tag: this.tag,
		});

		if (this.itemId) {
			protoId.rawId = {
				oneofKind: 'itemId',
				itemId: this.itemId,
			};
		} else if (this.spellId) {
			protoId.rawId = {
				oneofKind: 'spellId',
				spellId: this.spellId,
			};
		} else if (this.otherId) {
			protoId.rawId = {
				oneofKind: 'otherId',
				otherId: this.otherId,
			};
		}

		return protoId;
	}

	toProtoString(): string {
		return ActionIdProto.toJsonString(this.toProto());
	}

	withoutTag(): ActionId {
		return new ActionId(this.itemId, this.spellId, this.otherId, 0, this.baseName, this.baseName, this.iconUrl);
	}

	static fromEmpty(): ActionId {
		return new ActionId(0, 0, OtherAction.OtherActionNone, 0, '', '', '');
	}

	static fromItemId(itemId: number, tag?: number): ActionId {
		return new ActionId(itemId, 0, OtherAction.OtherActionNone, tag || 0, '', '', '');
	}

	static fromSpellId(spellId: number, tag?: number): ActionId {
		return new ActionId(0, spellId, OtherAction.OtherActionNone, tag || 0, '', '', '');
	}

	static fromOtherId(otherId: OtherAction, tag?: number): ActionId {
		return new ActionId(0, 0, otherId, tag || 0, '', '', '');
	}

	static fromItem(item: Item): ActionId {
		return ActionId.fromItemId(getWowheadItemId(item));
	}

	static fromProto(protoId: ActionIdProto): ActionId {
		if (protoId.rawId.oneofKind == 'spellId') {
			return ActionId.fromSpellId(protoId.rawId.spellId, protoId.tag);
		} else if (protoId.rawId.oneofKind == 'itemId') {
			return ActionId.fromItemId(protoId.rawId.itemId, protoId.tag);
		} else if (protoId.rawId.oneofKind == 'otherId') {
			return ActionId.fromOtherId(protoId.rawId.otherId, protoId.tag);
		} else {
			return ActionId.fromEmpty();
		}
	}

	static fromLogString(str: string): ActionId {
		const match = str.match(/{((SpellID)|(ItemID)|(OtherID)): (\d+)(, Tag: (-?\d+))?}/);
		if (match) {
			const idType = match[1];
			const id = parseInt(match[5]);
			return new ActionId(
					idType == 'ItemID' ? id : 0,
					idType == 'SpellID' ? id : 0,
					idType == 'OtherID' ? id : 0,
					match[7] ? parseInt(match[7]) : 0,
					'', '', '');
		} else {
			console.warn('Failed to parse action id from log: ' + str);
			return ActionId.fromEmpty();
		}
	}

	private static async getTooltipDataHelper(id: number, tooltipPostfix: string, cache: Map<number, Promise<any>>): Promise<any> {
		if (!cache.has(id)) {
			cache.set(id,
					fetch(`https://tbc.wowhead.com/tooltip/${tooltipPostfix}/${id}`)
					.then(response => response.json()));
		}

		return cache.get(id) as Promise<any>;
	}

	private static async getTooltipData(actionId: ActionId): Promise<any> {
		if (actionId.itemId) {
			return await ActionId.getTooltipDataHelper(actionId.itemId, 'item', itemToTooltipDataCache);
		} else {
			return await ActionId.getTooltipDataHelper(actionId.spellId, 'spell', spellToTooltipDataCache);
		}
	}
}

const itemToTooltipDataCache = new Map<number, Promise<any>>();
const spellToTooltipDataCache = new Map<number, Promise<any>>();

// Some items/spells have weird icons, so use this to show a different icon instead.
const idOverrides: Record<string, ActionId> = {};
idOverrides[ActionId.fromSpellId(37212).toProtoString()] = ActionId.fromItemId(29035); // Improved Wrath of Air Totem
idOverrides[ActionId.fromSpellId(37447).toProtoString()] = ActionId.fromItemId(30720); // Serpent-Coil Braid
idOverrides[ActionId.fromSpellId(37443).toProtoString()] = ActionId.fromItemId(30196); // Robes of Tirisfal (4pc bonus)
