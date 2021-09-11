import { Sim } from '../sim.js';

import { Component } from './component.js';

export class GearPicker extends Component {
  readonly rootElem: HTMLDivElement;

  constructor(sim: Sim) {
    super();

    this.rootElem = document.createElement('div');
    this.rootElem.classList.add('gear-picker-root');
  }

  getRootElement() {
    return this.rootElem;
  }
}
