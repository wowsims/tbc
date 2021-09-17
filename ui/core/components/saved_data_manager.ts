import { Sim } from '../sim';
import { TypedEvent } from '../typed_event';

import { Component } from '../components/component';

export type SavedDataManagerConfig<T> = {
  label: string;
  presets: Record<string, T>;

  changeEmitter: TypedEvent<any>;
  equals: (a: T, b: T) => boolean;
  getData: (sim: Sim) => T;
  setData: (sim: Sim, data: T) => void;
};

type SavedData<T> = {
  name: string;
  data: T;
  elem: HTMLElement;
};

export class SavedDataManager<T> extends Component {
  private readonly _sim: Sim;
  private readonly _config: SavedDataManagerConfig<T>;

  private readonly _userData: Array<SavedData<T>>;
  private readonly _presets: Array<SavedData<T>>;

  private readonly _userDataDiv: HTMLElement;
  private readonly _presetsDiv: HTMLElement;

  constructor(parent: HTMLElement, sim: Sim, config: SavedDataManagerConfig<T>) {
    super(parent, 'saved-data-manager-root');
    this._sim = sim;
    this._config = config;

    this._userData = [];
    this._presets = [];

    this.rootElem.innerHTML = `
    <div class="saved-data-user">
    </div>
    <div class="saved-data-presets">
    </div>
    <div class="saved-data-create-container">
      <input class="saved-data-save-input" type="text" placeholder="Label">
      <button class="saved-data-save-button">Save current ${config.label}</button>
    </div>
    `;

    this._userDataDiv = this.rootElem.getElementsByClassName('saved-data-user')[0] as HTMLElement;
    this._presetsDiv = this.rootElem.getElementsByClassName('saved-data-presets')[0] as HTMLElement;

    const saveInput = this.rootElem.getElementsByClassName('saved-data-save-input')[0] as HTMLInputElement;
    const saveButton = this.rootElem.getElementsByClassName('saved-data-save-button')[0] as HTMLButtonElement;
    saveButton.addEventListener('click', event => {
      const newName = saveInput.value;
      if (!newName) {
        alert(`Choose a label for your saved ${config.label}!`);
        return;
      }

      if (newName in this._presets) {
        alert(`${config.label} with name ${newName} already exists.`);
        return;
      }

      this.addSavedData(newName, config.getData(sim), false);
    });

    for (let presetName in config.presets) {
      this.addSavedData(presetName, config.presets[presetName], true);
    }
  }

  addSavedData(newName: string, data: T, isPreset: boolean) {
    const newData = this.makeSavedData(newName, data, isPreset);

    const dataArr = isPreset ? this._presets : this._userData;
    const containerDiv = isPreset ? this._presetsDiv : this._userDataDiv;

    const oldIdx = dataArr.findIndex(data => data.name == newName);
    if (oldIdx == -1) {
      containerDiv.appendChild(newData.elem);
      dataArr.push(newData);
    } else {
      containerDiv.replaceChild(newData.elem, dataArr[oldIdx].elem);
      dataArr[oldIdx] = newData;
    }
  }

  private makeSavedData(dataName: string, data: T, isPreset: boolean): SavedData<T> {
    const dataElem = document.createElement('div');
    dataElem.classList.add('saved-data-set-chip');
    dataElem.innerHTML = `
    <span class="saved-data-set-name">${dataName}</span>
    <span class="saved-data-set-delete fa fa-times"></span>
    `;

    dataElem.addEventListener('click', event => {
      this._config.setData(this._sim, data);
    });

    if (isPreset) {
      dataElem.classList.add('saved-data-preset');
    } else {
      const deleteButton = dataElem.getElementsByClassName('saved-data-set-delete')[0] as HTMLElement;
      deleteButton.addEventListener('click', event => {
        const idx = this._userData.findIndex(data => data.name == dataName);
        this._userData[idx].elem.remove();
        this._userData.splice(idx, 1);
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
    this._config.changeEmitter.on(checkActive);

    return {
      name: dataName,
      data: data,
      elem: dataElem,
    };
  }
}
