import { Component } from './component.js';

export class Actions extends Component {
  readonly rootElem: HTMLDivElement;
  readonly simButton: HTMLElement;
  readonly statWeightsButton: HTMLElement;

  constructor() {
    super();

    this.rootElem = document.createElement('div');
    this.rootElem.classList.add('actions-root');

    this.simButton = document.createElement('button');
    this.simButton.classList.add('actions-button');
    this.simButton.textContent = 'DPS';
    this.rootElem.appendChild(this.simButton);

    this.statWeightsButton = document.createElement('button');
    this.statWeightsButton.classList.add('actions-button');
    this.statWeightsButton.textContent = 'Calculate EP';
    this.rootElem.appendChild(this.statWeightsButton);
  }

  getRootElement() {
    return this.rootElem;
  }
}
