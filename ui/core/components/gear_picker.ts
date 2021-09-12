import { Enchant } from '../api/newapi';
import { Item } from '../api/newapi';
import { ItemSlot } from '../api/newapi';
import { SlotNames } from '../api/names';
import { GetEligibleEnchantSlots } from '../api/utils';
import { GetEligibleItemSlots } from '../api/utils';
import { EquippedItem } from '../equipped_item';
import { GetEmptyGemSocketIconUrl } from '../resources';
import { GetEmptySlotIconUrl } from '../resources';
import { GetItemIconUrl } from '../resources';
import { SetGemSocketCssClass } from '../css_utils';
import { SetItemQualityCssClass } from '../css_utils';
import { Sim } from '../sim';
import { getEnumValues } from '../utils';

import { Component } from './component';

export class GearPicker extends Component {
  // ItemSlot is used as the index
  readonly itemPickers: Array<ItemPicker>;

  constructor(parent: HTMLElement, sim: Sim) {
    super(parent, 'gear-picker-root');

    const leftSide = document.createElement('div');
    leftSide.classList.add('gear-picker-left');
    this.rootElem.appendChild(leftSide);

    const rightSide = document.createElement('div');
    rightSide.classList.add('gear-picker-right');
    this.rootElem.appendChild(rightSide);

    const leftItemPickers = [
      ItemSlot.ItemSlotHead,
      ItemSlot.ItemSlotNeck,
      ItemSlot.ItemSlotShoulder,
      ItemSlot.ItemSlotBack,
      ItemSlot.ItemSlotChest,
      ItemSlot.ItemSlotWrist,
      ItemSlot.ItemSlotMainHand,
      ItemSlot.ItemSlotOffHand,
    ].map(slot => new ItemPicker(leftSide, sim, slot));

    const rightItemPickers = [
      ItemSlot.ItemSlotHands,
      ItemSlot.ItemSlotWaist,
      ItemSlot.ItemSlotLegs,
      ItemSlot.ItemSlotFeet,
      ItemSlot.ItemSlotFinger1,
      ItemSlot.ItemSlotFinger2,
      ItemSlot.ItemSlotTrinket1,
      ItemSlot.ItemSlotTrinket2,
      ItemSlot.ItemSlotRanged,
    ].map(slot => new ItemPicker(rightSide, sim, slot));

    this.itemPickers = leftItemPickers.concat(rightItemPickers).sort((a, b) => a.slot - b.slot);
  }
}

class ItemPicker extends Component {
  readonly slot: ItemSlot;

  private readonly iconElem: HTMLElement;
  private readonly nameElem: HTMLElement;
  private readonly enchantElem: HTMLElement;
  private readonly socketsContainerElem: HTMLElement;

  // All items and enchants that are eligible for this slot
  private _items: Array<Item> = [];
  private _enchants: Array<Enchant> = [];

  private _equippedItem: EquippedItem | null = null;
  

  constructor(parent: HTMLElement, sim: Sim, slot: ItemSlot) {
    super(parent, 'item-picker-root');
    this.slot = slot;

    this.rootElem.innerHTML = `
      <a class="item-picker-icon" target="_blank">
        <div class="item-picker-sockets-container">
        </div>
      </a>
      <div class="item-picker-labels-container">
        <span class="item-picker-name"></span>
        <a class="item-picker-enchant" target="_blank"></a>
      </div>
    `;

    this.iconElem = this.rootElem.getElementsByClassName('item-picker-icon')[0] as HTMLElement;
    this.nameElem = this.rootElem.getElementsByClassName('item-picker-name')[0] as HTMLElement;
    this.enchantElem = this.rootElem.getElementsByClassName('item-picker-enchant')[0] as HTMLElement;
    this.socketsContainerElem = this.rootElem.getElementsByClassName('item-picker-sockets-container')[0] as HTMLElement;

    this.item = null;
    sim.gearListEmitter.on(gearListResult => {
      this._items = gearListResult.items.filter(item => GetEligibleItemSlots(item).includes(this.slot));
      this._enchants = gearListResult.enchants.filter(enchant => GetEligibleEnchantSlots(enchant).includes(this.slot));
    });
    sim.gearChangeEmitter.on(newGear => {
      if (newGear[this.slot]?.item.id != this._equippedItem?.item.id) {
        this.item = newGear[this.slot];
      }
    });
  }

  set item(newItem: EquippedItem | null) {
    if (newItem == null) {
      this.iconElem.style.backgroundImage = `url('${GetEmptySlotIconUrl(this.slot)}')`;
      this.iconElem.removeAttribute('data-wowhead');
      this.iconElem.removeAttribute('href');

      this.nameElem.textContent = SlotNames[this.slot];
      SetItemQualityCssClass(this.nameElem, null);

      this.enchantElem.textContent = '';
      this.socketsContainerElem.innerHTML = '';
    } else {
      this.nameElem.textContent = newItem.item.name;
      SetItemQualityCssClass(this.nameElem, newItem.item.quality);

      newItem.setWowheadData(this.iconElem);
      this.iconElem.setAttribute('href', 'https://tbc.wowhead.com/item=' + newItem.item.id);
      GetItemIconUrl(newItem.item.id).then(url => {
        this.iconElem.style.backgroundImage = `url('${url}')`;
      });

      this.socketsContainerElem.innerHTML = '';
      newItem.item.gemSockets.forEach((socketColor, gemIdx) => {
        const gemIconElem = document.createElement('div');
        gemIconElem.classList.add('item-picker-gem-icon');
        SetGemSocketCssClass(gemIconElem, socketColor);
        if (newItem.gems[gemIdx] == null) {
          gemIconElem.style.backgroundImage = `url('${GetEmptyGemSocketIconUrl(socketColor)}')`;
        } else {
          GetItemIconUrl(newItem.gems[gemIdx]!.id).then(url => {
            gemIconElem.style.backgroundImage = `url('${url}')`;
          });
        }
        this.socketsContainerElem.appendChild(gemIconElem);
      });
    }
    this._equippedItem = newItem;
  }
}
