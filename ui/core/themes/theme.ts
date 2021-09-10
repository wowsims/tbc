import { Sim } from '../sim.js';
import { Spec } from '../api/newapi';

export abstract class Theme {
  readonly parentElem: HTMLElement;
  readonly sim: Sim;

  constructor(parentElem: HTMLElement, spec: Spec) {
    this.parentElem = parentElem;
    this.sim = new Sim(spec);
  }
}
