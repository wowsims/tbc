import { GemColor } from './api/common.js';
import { ItemSlot } from './api/common.js';
const emptySlotIcons = {
    [ItemSlot.ItemSlotHead]: 'https://cdn.seventyupgrades.com/item-slots/Head.jpg',
    [ItemSlot.ItemSlotNeck]: 'https://cdn.seventyupgrades.com/item-slots/Neck.jpg',
    [ItemSlot.ItemSlotShoulder]: 'https://cdn.seventyupgrades.com/item-slots/Shoulders.jpg',
    [ItemSlot.ItemSlotBack]: 'https://cdn.seventyupgrades.com/item-slots/Back.jpg',
    [ItemSlot.ItemSlotChest]: 'https://cdn.seventyupgrades.com/item-slots/Chest.jpg',
    [ItemSlot.ItemSlotWrist]: 'https://cdn.seventyupgrades.com/item-slots/Wrists.jpg',
    [ItemSlot.ItemSlotHands]: 'https://cdn.seventyupgrades.com/item-slots/Hands.jpg',
    [ItemSlot.ItemSlotWaist]: 'https://cdn.seventyupgrades.com/item-slots/Waist.jpg',
    [ItemSlot.ItemSlotLegs]: 'https://cdn.seventyupgrades.com/item-slots/Legs.jpg',
    [ItemSlot.ItemSlotFeet]: 'https://cdn.seventyupgrades.com/item-slots/Feet.jpg',
    [ItemSlot.ItemSlotFinger1]: 'https://cdn.seventyupgrades.com/item-slots/Finger.jpg',
    [ItemSlot.ItemSlotFinger2]: 'https://cdn.seventyupgrades.com/item-slots/Finger.jpg',
    [ItemSlot.ItemSlotTrinket1]: 'https://cdn.seventyupgrades.com/item-slots/Trinket.jpg',
    [ItemSlot.ItemSlotTrinket2]: 'https://cdn.seventyupgrades.com/item-slots/Trinket.jpg',
    [ItemSlot.ItemSlotMainHand]: 'https://cdn.seventyupgrades.com/item-slots/MainHand.jpg',
    [ItemSlot.ItemSlotOffHand]: 'https://cdn.seventyupgrades.com/item-slots/OffHand.jpg',
    [ItemSlot.ItemSlotRanged]: 'https://cdn.seventyupgrades.com/item-slots/Ranged.jpg',
};
export function getEmptySlotIconUrl(slot) {
    return emptySlotIcons[slot];
}
// Some items/spells have weird icons, so use this to show a different icon instead.
const idOverrides = {};
idOverrides[JSON.stringify({ spellId: 37212 })] = { itemId: 29035 }; // Improved Wrath of Air Totem
async function getIconUrlHelper(id, tooltipPostfix, cache) {
    if (cache.has(id)) {
        return cache.get(id);
    }
    return fetch(`https://tbc.wowhead.com/tooltip/${tooltipPostfix}/${id}`)
        .then(response => response.json())
        .then(info => {
        const url = "https://wow.zamimg.com/images/wow/icons/large/" + info['icon'] + ".jpg";
        cache.set(id, url);
        return url;
    });
}
const itemToIconCache = new Map();
const spellToIconCache = new Map();
export async function getIconUrl(id) {
    const idString = JSON.stringify(id);
    if (idOverrides[idString])
        id = idOverrides[idString];
    if ('itemId' in id) {
        return await getIconUrlHelper(id.itemId, 'item', itemToIconCache);
    }
    else {
        return await getIconUrlHelper(id.spellId, 'spell', spellToIconCache);
    }
}
export function setWowheadHref(elem, id) {
    if ('itemId' in id) {
        elem.href = 'https://tbc.wowhead.com/item=' + id.itemId;
    }
    else {
        elem.href = 'https://tbc.wowhead.com/spell=' + id.spellId;
    }
}
const emptyGemSocketIcons = {
    [GemColor.GemColorBlue]: 'https://wow.zamimg.com/images/icons/socket-blue.gif',
    [GemColor.GemColorMeta]: 'https://wow.zamimg.com/images/icons/socket-meta.gif',
    [GemColor.GemColorRed]: 'https://wow.zamimg.com/images/icons/socket-red.gif',
    [GemColor.GemColorYellow]: 'https://wow.zamimg.com/images/icons/socket-yellow.gif',
};
export function getEmptyGemSocketIconUrl(color) {
    if (emptyGemSocketIcons[color])
        return emptyGemSocketIcons[color];
    throw new Error('No empty socket url for gem socket color: ' + color);
}
