import { Sim } from '../internal/sim.js';

export abstract class Theme {
  readonly parentElem: HTMLElement;
  readonly sim: Sim;

  constructor(parentElem: HTMLElement) {
    this.parentElem = parentElem;
    this.sim = new Sim(3);
  }
}
