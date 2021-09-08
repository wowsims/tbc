import { Component } from './component.js';

export class CharacterStats extends Component {
  readonly rootElem: HTMLDivElement;

  constructor() {
    super();
    this.rootElem = document.createElement('div');
    this.rootElem.classList.add('character-stats-root');
  }

  getRootElement() {
    return this.rootElem;
  }
}
