import { Enchant } from '../api/newapi';
import { Item } from '../api/newapi';
import { ItemQuality } from '../api/newapi';
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

    const selectorModal = new SelectorModal(document.body, sim);

    const leftItemPickers = [
      ItemSlot.ItemSlotHead,
      ItemSlot.ItemSlotNeck,
      ItemSlot.ItemSlotShoulder,
      ItemSlot.ItemSlotBack,
      ItemSlot.ItemSlotChest,
      ItemSlot.ItemSlotWrist,
      ItemSlot.ItemSlotMainHand,
      ItemSlot.ItemSlotOffHand,
    ].map(slot => new ItemPicker(leftSide, sim, slot, selectorModal));

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
    ].map(slot => new ItemPicker(rightSide, sim, slot, selectorModal));

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
  

  constructor(parent: HTMLElement, sim: Sim, slot: ItemSlot, selectorModal: SelectorModal) {
    super(parent, 'item-picker-root');
    this.slot = slot;

    this.rootElem.innerHTML = `
      <a class="item-picker-icon" target="_blank" data-toggle="modal" data-target="#selectorModal">
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

      this.iconElem.addEventListener('click', event => {
        selectorModal.setData(this.slot, this._equippedItem, this._items, this._enchants);
      });
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

class SelectorModal extends Component {
  private readonly sim: Sim;
  private readonly tabsElem: HTMLElement;
  private readonly contentElem: HTMLElement;
  private readonly closeButton: HTMLButtonElement;

  constructor(parent: HTMLElement, sim: Sim) {
    super(parent, 'selector-model-root');
    this.sim = sim;

    this.rootElem.innerHTML = `
    <div class="modal fade selector-modal" id="selectorModal" tabindex="-1" role="dialog" aria-labelledby="selectorModalTitle" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-body">
            <ul class="nav nav-tabs selector-modal-tabs">
            </ul>
            <div class="tab-content selector-modal-tab-content">
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary selector-modal-close-button" data-dismiss="modal">Close</button>
          </div>
        </div>
      </div>
    </div>
    `;

    this.tabsElem = this.rootElem.getElementsByClassName('selector-modal-tabs')[0] as HTMLElement;
    this.contentElem = this.rootElem.getElementsByClassName('selector-modal-tab-content')[0] as HTMLElement;
    this.closeButton = this.rootElem.getElementsByClassName('selector-modal-close-button')[0] as HTMLButtonElement;
  }

  setData(slot: ItemSlot, equippedItem: EquippedItem | null, eligibleItems: Array<Item>, eligibleEnchants: Array<Enchant>) {
    this.tabsElem.innerHTML = '';
    this.contentElem.innerHTML = '';
    this.tabsElem.innerHTML = `
    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
      <span aria-hidden="true">&times;</span>
    </button>
    `;

    this.addTab(
        'Items',
        slot,
        equippedItem,
        eligibleItems,
        equippedItem => equippedItem?.item.id || 0,
        item => {
          return {
            id: item.id,
            name: item.name,
            quality: item.quality,
            onclick: item => {
              this.sim.equipItem(slot, new EquippedItem(item));
            },
          };
        });
  }

  /**
   * Adds one of the tabs for the item selector menu.
   *
   * T is expected to be Item, Enchant, or Gem. Tab menus for all 3 looks extremely
   * similar so this function uses extra functions to do it generically.
   */
  private addTab<T>(
        label: string,
        slot: ItemSlot,
        equippedItem: EquippedItem | null,
        items: Array<T>,
        equippedToIdFn: (equippedItem: EquippedItem | null) => number,
        getItemData: (item: T) => {
          id: number,
          name: string,
          quality: ItemQuality,
          onclick: (item: T) => void,
        }) {
    const tabElem = document.createElement('li');
    this.tabsElem.insertBefore(tabElem, this.tabsElem.lastChild);
    const tabContentId = label + '-tab';
    tabElem.innerHTML = `<a class="selector-modal-item-tab" data-toggle="tab" href="#${tabContentId}">${label}</a>`;

    const tabContent = document.createElement('div');
    tabContent.id = tabContentId;
    tabContent.classList.add('tab-pane', 'fade', 'selector-modal-tab-content');
    this.contentElem.appendChild(tabContent);
    tabContent.innerHTML = `
    <div class="selector-modal-tab-content-header">
      <button class="selector-modal-remove-button">Remove</button>
      <input class="selector-modal-search" type="text" placeholder="Search...">
    </div>
    <ul class="selector-modal-list"></ul>
    `;

    if (label == 'Items') {
      tabElem.classList.add('active', 'in');
      tabContent.classList.add('active', 'in');
    }

    const searchInput = tabContent.getElementsByClassName('selector-modal-search')[0] as HTMLInputElement;
    const listElem = tabContent.getElementsByClassName('selector-modal-list')[0] as HTMLElement;

    const listItemElems = items.map(item => {
      const itemData = getItemData(item);

      const listItemElem = document.createElement('li');
      listItemElem.classList.add('selector-modal-list-item');
      if (itemData.id == equippedToIdFn(equippedItem)) {
        listItemElem.classList.add('active');
      }
      listElem.appendChild(listItemElem);

      listItemElem.dataset.id = String(itemData.id);
      listItemElem.innerHTML = `
        <a class="selector-modal-list-item-icon" href="https://tbc.wowhead.com/item=${itemData.id}"></a>
        <a class="selector-modal-list-item-name" href="https://tbc.wowhead.com/item=${itemData.id}">${itemData.name}</a>
      `;

      const iconElem = tabContent.getElementsByClassName('selector-modal-list-item-icon')[0] as HTMLImageElement;
      GetItemIconUrl(itemData.id).then(url => {
        iconElem.style.backgroundImage = `url('${url}')`;
      });

      const nameElem = tabContent.getElementsByClassName('selector-modal-list-item-name')[0] as HTMLImageElement;
      SetItemQualityCssClass(nameElem, itemData.quality);

      const onclick = (event: Event) => {
        event.preventDefault();
        itemData.onclick(item);
      };
      nameElem.addEventListener('click', onclick);
      iconElem.addEventListener('click', onclick);

      return listItemElem;
    });

    const removeButton = tabContent.getElementsByClassName('selector-modal-remove-button')[0] as HTMLButtonElement;
    removeButton.addEventListener('click', event => {
      listItemElems.forEach(elem => elem.classList.remove('active'));
      this.sim.unequipItem(slot);
    });

    this.sim.gearChangeEmitter.on(newGear => {
      const newEquippedItem = newGear[slot];
      listItemElems.forEach(elem => {
        elem.classList.remove('active');
        if (parseInt(elem.dataset.id!) == equippedToIdFn(newEquippedItem)) {
          elem.classList.add('active');
        }
      });
    });
  }
}
