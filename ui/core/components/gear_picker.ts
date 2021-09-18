import { EquippedItem } from '../api/equipped_item';
import { Enchant } from '../api/newapi';
import { Item } from '../api/newapi';
import { ItemQuality } from '../api/newapi';
import { ItemSlot } from '../api/newapi';
import { enchantDescriptions } from '../api/names';
import { slotNames } from '../api/names';
import { gemEligibleForSocket } from '../api/utils';
import { getEligibleEnchantSlots } from '../api/utils';
import { getEligibleItemSlots } from '../api/utils';
import { getEmptyGemSocketIconUrl } from '../resources';
import { getEmptySlotIconUrl } from '../resources';
import { getIconUrl } from '../resources';
import { setWowheadHref } from '../resources';
import { setGemSocketCssClass } from '../css_utils';
import { setItemQualityCssClass } from '../css_utils';
import { Sim } from '../sim';
import { getEnumValues } from '../utils';

import { Component } from './component';

export class GearPicker extends Component {
  // ItemSlot is used as the index
  readonly itemPickers: Array<ItemPicker>;

  constructor(parent: HTMLElement, sim: Sim<any>) {
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

  private readonly sim: Sim<any>;
  private readonly iconElem: HTMLAnchorElement;
  private readonly nameElem: HTMLElement;
  private readonly enchantElem: HTMLElement;
  private readonly socketsContainerElem: HTMLElement;

  // All items and enchants that are eligible for this slot
  private _items: Array<Item> = [];
  private _enchants: Array<Enchant> = [];

  private _equippedItem: EquippedItem | null = null;
  

  constructor(parent: HTMLElement, sim: Sim<any>, slot: ItemSlot, selectorModal: SelectorModal) {
    super(parent, 'item-picker-root');
    this.slot = slot;
    this.sim = sim;

    this.rootElem.innerHTML = `
      <a class="item-picker-icon" target="_blank" data-toggle="modal" data-target="#selectorModal">
        <div class="item-picker-sockets-container">
        </div>
      </a>
      <div class="item-picker-labels-container">
        <span class="item-picker-name"></span>
        <span class="item-picker-enchant"></span>
      </div>
    `;

    this.iconElem = this.rootElem.getElementsByClassName('item-picker-icon')[0] as HTMLAnchorElement;
    this.nameElem = this.rootElem.getElementsByClassName('item-picker-name')[0] as HTMLElement;
    this.enchantElem = this.rootElem.getElementsByClassName('item-picker-enchant')[0] as HTMLElement;
    this.socketsContainerElem = this.rootElem.getElementsByClassName('item-picker-sockets-container')[0] as HTMLElement;

    this.item = null;
    sim.gearListEmitter.on(gearListResult => {
      this._items = gearListResult.items.filter(item => getEligibleItemSlots(item).includes(this.slot));
      this._enchants = gearListResult.enchants.filter(enchant => getEligibleEnchantSlots(enchant).includes(this.slot));

      this.iconElem.addEventListener('click', event => {
        selectorModal.setData(this.slot, this._equippedItem, this._items, this._enchants);
      });
    });
    sim.gearChangeEmitter.on(() => {
      this.item = sim.getEquippedItem(slot);
    });
  }

  set item(newItem: EquippedItem | null) {
    // Clear everything first
    this.iconElem.style.backgroundImage = `url('${getEmptySlotIconUrl(this.slot)}')`;
    this.iconElem.removeAttribute('data-wowhead');
    this.iconElem.removeAttribute('href');

    this.nameElem.textContent = slotNames[this.slot];
    setItemQualityCssClass(this.nameElem, null);

    this.enchantElem.textContent = '';
    //this.enchantElem.removeAttribute('data-wowhead');
    //this.enchantElem.removeAttribute('href');
    this.socketsContainerElem.innerHTML = '';

    if (newItem != null) {
      this.nameElem.textContent = newItem.item.name;
      setItemQualityCssClass(this.nameElem, newItem.item.quality);

      this.sim.setWowheadData(newItem, this.iconElem);
      setWowheadHref(this.iconElem, {itemId: newItem.item.id});
      getIconUrl({itemId: newItem.item.id}).then(url => {
        this.iconElem.style.backgroundImage = `url('${url}')`;
      });

      if (newItem.enchant) {
        this.enchantElem.textContent = enchantDescriptions.get(newItem.enchant.id) || newItem.enchant.name;
        //this.enchantElem.setAttribute('href', 'https://tbc.wowhead.com/item=' + newItem.enchant.id);
      }

      newItem.item.gemSockets.forEach((socketColor, gemIdx) => {
        const gemIconElem = document.createElement('div');
        gemIconElem.classList.add('item-picker-gem-icon');
        setGemSocketCssClass(gemIconElem, socketColor);
        if (newItem.gems[gemIdx] == null) {
          gemIconElem.style.backgroundImage = `url('${getEmptyGemSocketIconUrl(socketColor)}')`;
        } else {
          getIconUrl({itemId: newItem.gems[gemIdx]!.id}).then(url => {
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
  private readonly sim: Sim<any>;
  private readonly tabsElem: HTMLElement;
  private readonly contentElem: HTMLElement;
  private readonly closeButton: HTMLButtonElement;

  constructor(parent: HTMLElement, sim: Sim<any>) {
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
            onEquip: item => {
              const equippedItem = this.sim.getEquippedItem(slot);
              if (equippedItem) {
                this.sim.equipItem(slot, equippedItem.withItem(item));
              } else {
                this.sim.equipItem(slot, new EquippedItem(item));
              }
            },
          };
        },
        () => {
          this.sim.equipItem(slot, null);
        });

    this.addTab(
        'Enchants',
        slot,
        equippedItem,
        eligibleEnchants,
        equippedItem => equippedItem?.enchant?.id || 0,
        enchant => {
          return {
            id: enchant.id,
            name: enchant.name,
            quality: enchant.quality,
            onEquip: enchant => {
              const equippedItem = this.sim.getEquippedItem(slot);
              if (equippedItem)
                this.sim.equipItem(slot, equippedItem.withEnchant(enchant));
            },
          };
        },
        () => {
          const equippedItem = this.sim.getEquippedItem(slot);
          if (equippedItem)
            this.sim.equipItem(slot, equippedItem.withEnchant(null));
        });

    this.addGemTabs(slot, equippedItem);
  }

  private addGemTabs(slot: ItemSlot, equippedItem: EquippedItem | null) {
    equippedItem?.item.gemSockets.forEach((socketColor, socketIdx) => {
      this.addTab(
          'Gem ' + (socketIdx + 1),
          slot,
          equippedItem,
          this.sim.getGems().filter(gem => gemEligibleForSocket(gem, socketColor)),
          equippedItem => equippedItem?.gems[socketIdx]?.id || 0,
          gem => {
            return {
              id: gem.id,
              name: gem.name,
              quality: gem.quality,
              onEquip: gem => {
                const equippedItem = this.sim.getEquippedItem(slot);
                if (equippedItem)
                  this.sim.equipItem(slot, equippedItem.withGem(gem, socketIdx));
              },
            };
          },
          () => {
            const equippedItem = this.sim.getEquippedItem(slot);
            if (equippedItem)
              this.sim.equipItem(slot, equippedItem.withGem(null, socketIdx));
          });
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
          onEquip: (item: T) => void,
        },
        onRemove: () => void) {
    if (items.length == 0) {
      return;
    }

    const tabElem = document.createElement('li');
    this.tabsElem.insertBefore(tabElem, this.tabsElem.lastChild);
    const tabContentId = (label + '-tab').split(' ').join('');
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
      listItemElem.dataset.name = itemData.name;

      listItemElem.innerHTML = `
        <a class="selector-modal-list-item-icon"></a>
        <a class="selector-modal-list-item-name">${itemData.name}</a>
      `;
      setWowheadHref(listItemElem.children[0] as HTMLAnchorElement, {itemId: itemData.id});
      setWowheadHref(listItemElem.children[1] as HTMLAnchorElement, {itemId: itemData.id});

      const iconElem = listItemElem.getElementsByClassName('selector-modal-list-item-icon')[0] as HTMLImageElement;
      getIconUrl({itemId: itemData.id}).then(url => {
        iconElem.style.backgroundImage = `url('${url}')`;
      });

      const nameElem = listItemElem.getElementsByClassName('selector-modal-list-item-name')[0] as HTMLImageElement;
      setItemQualityCssClass(nameElem, itemData.quality);

      const onclick = (event: Event) => {
        event.preventDefault();
        itemData.onEquip(item);

        // If the item changes, the gem slots might change, so remove and recreate the gem tabs
        if (Item.is(item)) {
          this.removeTabs('Gem');
          this.addGemTabs(slot, this.sim.getEquippedItem(slot));
        }
      };
      nameElem.addEventListener('click', onclick);
      iconElem.addEventListener('click', onclick);

      return listItemElem;
    });

    const removeButton = tabContent.getElementsByClassName('selector-modal-remove-button')[0] as HTMLButtonElement;
    removeButton.addEventListener('click', event => {
      listItemElems.forEach(elem => elem.classList.remove('active'));
      onRemove();
    });

    this.sim.gearChangeEmitter.on(() => {
      const newEquippedItem = this.sim.getEquippedItem(slot);
      listItemElems.forEach(elem => {
        elem.classList.remove('active');
        if (parseInt(elem.dataset.id!) == equippedToIdFn(newEquippedItem)) {
          elem.classList.add('active');
        }
      });
    });

    const searchInput = tabContent.getElementsByClassName('selector-modal-search')[0] as HTMLInputElement;
    searchInput.addEventListener('input', event => {
      listItemElems.forEach(elem => {
        if (elem.dataset.name!.toLowerCase().includes(searchInput.value.toLowerCase())) {
          elem.style.display = 'flex';
        } else {
          elem.style.display = 'none';
        }
      });
    });
  }

  private removeTabs(labelSubstring: string) {
    const tabElems = Array.prototype.slice.call(this.tabsElem.getElementsByClassName('selector-modal-item-tab')).filter(tab => tab.textContent.includes(labelSubstring));
    const contentElems = tabElems
        .map(tabElem => document.getElementById(tabElem.href.substring(1)))
        .filter(tabElem => Boolean(tabElem));

    tabElems.forEach(elem => elem.remove());
    contentElems.forEach(elem => elem!.remove());
  }
}
