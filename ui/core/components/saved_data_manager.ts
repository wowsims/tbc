import { Spec } from '/tbc/core/proto/common.js';
import { Sim } from '/tbc/core/sim.js';
import { TypedEvent } from '/tbc/core/typed_event.js';

import { Component } from '/tbc/core/components/component.js';

declare var tippy: any;

export type SavedDataManagerConfig<SpecType extends Spec, T> = {
  label: string;
  storageKey: string;
  changeEmitters: Array<TypedEvent<any>>,
  equals: (a: T, b: T) => boolean;
  getData: (sim: Sim<SpecType>) => T;
  setData: (sim: Sim<SpecType>, data: T) => void;
  toJson: (a: T) => any;
  fromJson: (obj: any) => T;
};

type SavedData<T> = {
  name: string;
  data: T;
  elem: HTMLElement;
};

export class SavedDataManager<SpecType extends Spec, T> extends Component {
  private readonly sim: Sim<SpecType>;
  private readonly config: SavedDataManagerConfig<SpecType, T>;

  private readonly userData: Array<SavedData<T>>;
  private readonly presets: Array<SavedData<T>>;

  private readonly savedDataDiv: HTMLElement;
  private readonly saveInput: HTMLInputElement;
	private frozen: boolean;

  constructor(parent: HTMLElement, sim: Sim<Spec>, config: SavedDataManagerConfig<SpecType, T>) {
    super(parent, 'saved-data-manager-root');
    this.sim = sim;
    this.config = config;

    this.userData = [];
    this.presets = [];
		this.frozen = false;

    this.rootElem.innerHTML = `
    <div class="saved-data-container">
    </div>
    <div class="saved-data-create-container">
      <input class="saved-data-save-input" type="text" placeholder="Label">
      <button class="saved-data-save-button">Save current ${config.label}</button>
    </div>
    `;

    this.savedDataDiv = this.rootElem.getElementsByClassName('saved-data-container')[0] as HTMLElement;

    this.saveInput = this.rootElem.getElementsByClassName('saved-data-save-input')[0] as HTMLInputElement;
    const saveButton = this.rootElem.getElementsByClassName('saved-data-save-button')[0] as HTMLButtonElement;
    saveButton.addEventListener('click', event => {
			if (this.frozen)
				return;

      const newName = this.saveInput.value;
      if (!newName) {
        alert(`Choose a label for your saved ${config.label}!`);
        return;
      }

      if (newName in this.presets) {
        alert(`${config.label} with name ${newName} already exists.`);
        return;
      }

      this.addSavedData(newName, config.getData(this.sim), false);
      this.saveUserData();
    });
  }

  addSavedData(newName: string, data: T, isPreset: boolean, tooltipInfo?: string) {
    const newData = this.makeSavedData(newName, data, isPreset, tooltipInfo);

    const dataArr = isPreset ? this.presets : this.userData;

    const oldIdx = dataArr.findIndex(data => data.name == newName);
    if (oldIdx == -1) {
      if (isPreset || this.presets.length == 0) {
        this.savedDataDiv.appendChild(newData.elem);
      } else {
        this.savedDataDiv.insertBefore(newData.elem, this.presets[0].elem);
      }
      dataArr.push(newData);
    } else {
      this.savedDataDiv.replaceChild(newData.elem, dataArr[oldIdx].elem);
      dataArr[oldIdx] = newData;
    }
  }

  private makeSavedData(dataName: string, data: T, isPreset: boolean, tooltipInfo?: string): SavedData<T> {
    const dataElem = document.createElement('div');
    dataElem.classList.add('saved-data-set-chip');
    dataElem.innerHTML = `
    <span class="saved-data-set-name">${dataName}</span>
    <span class="saved-data-set-tooltip fa fa-info-circle"></span>
    <span class="saved-data-set-delete fa fa-times"></span>
    `;

    dataElem.addEventListener('click', event => {
      this.config.setData(this.sim, data);
      this.saveInput.value = dataName;
    });

    if (isPreset) {
      dataElem.classList.add('saved-data-preset');
    } else {
      const deleteButton = dataElem.getElementsByClassName('saved-data-set-delete')[0] as HTMLElement;
      deleteButton.addEventListener('click', event => {
        event.stopPropagation();
        const shouldDelete = confirm(`Delete saved ${this.config.label} '${dataName}'?`);
        if (!shouldDelete)
          return;

        const idx = this.userData.findIndex(data => data.name == dataName);
        this.userData[idx].elem.remove();
        this.userData.splice(idx, 1);
        this.saveUserData();
      });
    }

    if (tooltipInfo) {
      dataElem.classList.add('saved-data-has-tooltip');
      tippy(dataElem.getElementsByClassName('saved-data-set-tooltip')[0], {
        'content': tooltipInfo,
        'allowHTML': true,
      });
    }

    const checkActive = () => {
      if (this.config.equals(data, this.config.getData(this.sim))) {
        dataElem.classList.add('active');
      } else {
        dataElem.classList.remove('active');
      }
    };

    checkActive();
    this.config.changeEmitters.forEach(emitter => emitter.on(checkActive));

    return {
      name: dataName,
      data: data,
      elem: dataElem,
    };
  }

  // Save data to window.localStorage.
  private saveUserData() {
    const gearData: Record<string, Object> = {};
    this.userData.forEach(savedData => {
      gearData[savedData.name] = this.config.toJson(savedData.data);
    });

    window.localStorage.setItem(this.config.storageKey, JSON.stringify(gearData));
  }

  // Load data from window.localStorage.
  loadUserData() {
    const dataStr = window.localStorage.getItem(this.config.storageKey);
    if (!dataStr)
      return;

    const jsonData = JSON.parse(dataStr);
    for (let name in jsonData) {
      this.addSavedData(name, this.config.fromJson(jsonData[name]), false);
    }
  }

	// Prevent user input from creating / deleting saved data.
	freeze() {
		this.frozen = true;
		this.rootElem.classList.add('frozen');
	}
}
