import { GemColor } from './api/newapi';
import { Item } from './api/newapi';
import { ItemQuality } from './api/newapi';
import { ItemSlot } from './api/newapi';

const emptySlotIcons: Record<ItemSlot, string> = {
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
export function GetEmptySlotIconUrl(slot: ItemSlot): string {
  return emptySlotIcons[slot];
}

async function GetIconUrl(id: number, tooltipPostfix: string, cache: Map<number, string>): Promise<string> {
  if (cache.has(id)) {
    return cache.get(id) as string;
  }

  return fetch(`https://tbc.wowhead.com/tooltip/${tooltipPostfix}/${id}`)
  .then(response => response.json())
  .then(info => {
    const url = "https://wow.zamimg.com/images/wow/icons/large/" + info['icon'] + ".jpg";
    cache.set(id, url);
    return url;
  });
}

const itemToIconCache = new Map<number, string>();
export async function GetItemIconUrl(id: number): Promise<string> {
  return await GetIconUrl(id, 'item', itemToIconCache);
}

const spellToIconCache = new Map<number, string>();
export async function GetSpellIconUrl(id: number): Promise<string> {
  return await GetIconUrl(id, 'spell', spellToIconCache);
}

const emptyGemSocketIcons: Partial<Record<GemColor, string>> = {
  [GemColor.GemColorBlue]: 'https://wow.zamimg.com/images/icons/socket-blue.gif',
  [GemColor.GemColorMeta]: 'https://wow.zamimg.com/images/icons/socket-meta.gif',
  [GemColor.GemColorRed]: 'https://wow.zamimg.com/images/icons/socket-red.gif',
  [GemColor.GemColorYellow]: 'https://wow.zamimg.com/images/icons/socket-yellow.gif',
};
export function GetEmptyGemSocketIconUrl(color: GemColor): string {
  if (emptyGemSocketIcons[color])
    return emptyGemSocketIcons[color] as string;

  throw new Error('No empty socket url for gem socket color: ' + color);
}
