import { Component } from './component.js';
import { Sim } from '../sim.js';

export class Actions extends Component {
  readonly rootElem: HTMLDivElement;

  constructor(sim: Sim) {
    super();

    this.rootElem = document.createElement('div');
    this.rootElem.classList.add('actions-root');

    const simButton = document.createElement('button');
    simButton.classList.add('actions-button');
    simButton.textContent = 'DPS';
    this.rootElem.appendChild(simButton);

    const statWeightsButton = document.createElement('button');
    statWeightsButton.classList.add('actions-button');
    statWeightsButton.textContent = 'Calculate EP';
    this.rootElem.appendChild(statWeightsButton);

    const iterationsDiv = document.createElement('div');
    iterationsDiv.classList.add('iterations-div');
    iterationsDiv.innerHTML = `
      <span class="iterations-label">Iterations</span>
      <input class="iterations-input" type="number" min="1" value="1000" step="1000">
    `;
    this.rootElem.appendChild(iterationsDiv);
    const iterationsInput = document.createElement('input');
  }

  getRootElement() {
    return this.rootElem;
  }
}
