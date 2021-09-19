import { Spec } from '../api/newapi';
import { Sim } from '../sim';
import { TypedEvent } from '../typed_event';

import { Component } from '../components/component';

declare var tippy: any;

export type SavedDataManagerConfig<SpecType extends Spec, T> = {
  label: string;
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
  private readonly _sim: Sim<SpecType>;
  private readonly _config: SavedDataManagerConfig<SpecType, T>;

  private readonly _userData: Array<SavedData<T>>;
  private readonly _presets: Array<SavedData<T>>;

  private readonly _savedDataDiv: HTMLElement;
  private readonly _saveInput: HTMLInputElement;

  constructor(parent: HTMLElement, sim: Sim<Spec>, config: SavedDataManagerConfig<SpecType, T>) {
    super(parent, 'saved-data-manager-root');
    this._sim = sim;
    this._config = config;

    this._userData = [];
    this._presets = [];

    this.rootElem.innerHTML = `
    <div class="saved-data-container">
    </div>
    <div class="saved-data-create-container">
      <input class="saved-data-save-input" type="text" placeholder="Label">
      <button class="saved-data-save-button">Save current ${config.label}</button>
    </div>
    `;

    this._savedDataDiv = this.rootElem.getElementsByClassName('saved-data-container')[0] as HTMLElement;

    this._saveInput = this.rootElem.getElementsByClassName('saved-data-save-input')[0] as HTMLInputElement;
    const saveButton = this.rootElem.getElementsByClassName('saved-data-save-button')[0] as HTMLButtonElement;
    saveButton.addEventListener('click', event => {
      const newName = this._saveInput.value;
      if (!newName) {
        alert(`Choose a label for your saved ${config.label}!`);
        return;
      }

      if (newName in this._presets) {
        alert(`${config.label} with name ${newName} already exists.`);
        return;
      }

      this.addSavedData(newName, config.getData(sim), false);
      this.saveUserData();
    });
  }

  addSavedData(newName: string, data: T, isPreset: boolean, tooltipInfo?: string) {
    const newData = this.makeSavedData(newName, data, isPreset, tooltipInfo);

    const dataArr = isPreset ? this._presets : this._userData;

    const oldIdx = dataArr.findIndex(data => data.name == newName);
    if (oldIdx == -1) {
      if (isPreset || this._presets.length == 0) {
        this._savedDataDiv.appendChild(newData.elem);
      } else {
        this._savedDataDiv.insertBefore(newData.elem, this._presets[0].elem);
      }
      dataArr.push(newData);
    } else {
      this._savedDataDiv.replaceChild(newData.elem, dataArr[oldIdx].elem);
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
      this._config.setData(this._sim, data);
      this._saveInput.value = dataName;
    });

    if (isPreset) {
      dataElem.classList.add('saved-data-preset');
    } else {
      const deleteButton = dataElem.getElementsByClassName('saved-data-set-delete')[0] as HTMLElement;
      deleteButton.addEventListener('click', event => {
        event.stopPropagation();
        const shouldDelete = confirm(`Delete saved ${this._config.label} '${dataName}'?`);
        if (!shouldDelete)
          return;

        const idx = this._userData.findIndex(data => data.name == dataName);
        this._userData[idx].elem.remove();
        this._userData.splice(idx, 1);
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
      if (this._config.equals(data, this._config.getData(this._sim))) {
        dataElem.classList.add('active');
      } else {
        dataElem.classList.remove('active');
      }
    };

    checkActive();
    this._config.changeEmitters.forEach(emitter => emitter.on(checkActive));

    return {
      name: dataName,
      data: data,
      elem: dataElem,
    };
  }

  // Save data to window.localStorage.
  private saveUserData() {
    const gearData: Record<string, Object> = {};
    this._userData.forEach(savedData => {
      gearData[savedData.name] = this._config.toJson(savedData.data);
    });

    window.localStorage.setItem(this._config.label, JSON.stringify(gearData));
  }

  // Load data from window.localStorage.
  loadUserData() {
    const dataStr = window.localStorage.getItem(this._config.label);
    if (!dataStr)
      return;

    const jsonData = JSON.parse(dataStr);
    for (let name in jsonData) {
      this.addSavedData(name, this._config.fromJson(jsonData[name]), false);
    }
  }
}
