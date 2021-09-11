import { IndividualSimRequest } from '../api/newapi';
import { Stat } from '../api/newapi';
import { StatWeightsRequest } from '../api/newapi';
import { Sim } from '../sim.js';

import { Component } from './component.js';
import { Results } from './results.js';

export class Actions extends Component {
  readonly rootElem: HTMLDivElement;

  constructor(sim: Sim, results: Results, epStats: Array<Stat>, epReferenceStat: Stat) {
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
      <span class="iterations-label">Iterations:</span>
      <input class="iterations-input" type="number" min="1" value="1000" step="1000">
    `;
    this.rootElem.appendChild(iterationsDiv);
    const iterationsInput = iterationsDiv.getElementsByClassName('iterations-input')[0] as HTMLInputElement;

    simButton.addEventListener('click', async () => {
      const request = sim.createSimRequest();
      request.iterations = parseInt(iterationsInput.value);

      results.setPending();
      const result = await sim.individualSim(request);
      results.setSimResult(result);
    });

    statWeightsButton.addEventListener('click', async () => {
      const simRequest = sim.createSimRequest();
      simRequest.iterations = parseInt(iterationsInput.value);

      const statWeightsRequest = StatWeightsRequest.create({
        options: simRequest,
        statsToWeigh: epStats,
        epReferenceStat: epReferenceStat,
      });

      results.setPending();
      const result = await sim.statWeights(statWeightsRequest);
      results.setStatWeights(result, epStats);
    });
  }

  getRootElement() {
    return this.rootElem;
  }
}
