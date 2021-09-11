import { Sim } from '../sim.js';

import { Component } from './component.js';

export class GearPicker extends Component {
  constructor(parent: HTMLElement, sim: Sim) {
    super(parent, 'gear-picker-root');
  }
}

class ItemPicker extends Component {
  constructor(parent: HTMLElement) {
    super(parent, 'item-picker-root');
  }
}
