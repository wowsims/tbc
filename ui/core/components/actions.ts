import { IndividualSimRequest } from '../proto/api.js';
import { Stat } from '../proto/common.js';
import { StatWeightsRequest } from '../proto/api.js';
import { Sim } from '../sim.js';

import { Component } from './component.js';
import { DetailedResults } from './detailed_results.js';
import { Results } from './results.js';

export class Actions extends Component {
  constructor(parent: HTMLElement, sim: Sim<any>, epStats: Array<Stat>, epReferenceStat: Stat, results: Results, detailedResults: DetailedResults) {
    super(parent, 'actions-root');

    const simButton = document.createElement('button');
    simButton.classList.add('actions-button');
    simButton.textContent = 'DPS';
    this.rootElem.appendChild(simButton);

    const statWeightsButton = document.createElement('button');
    statWeightsButton.classList.add('actions-button');
    statWeightsButton.textContent = 'EP Weights';
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
      const iterations = parseInt(iterationsInput.value);
      const simRequest = sim.makeCurrentIndividualSimRequest(iterations, false);

      results.setPending();
      detailedResults.setPending();
      const result = await sim.individualSim(simRequest);
      results.setSimResult(result);
      detailedResults.setSimResult(simRequest, result);
    });

    statWeightsButton.addEventListener('click', async () => {
      const iterations = parseInt(iterationsInput.value);
      const simRequest = sim.makeCurrentIndividualSimRequest(iterations, false);

      const statWeightsRequest = StatWeightsRequest.create({
        options: simRequest,
        statsToWeigh: epStats,
        epReferenceStat: epReferenceStat,
      });

      results.setPending();
      const result = await sim.statWeights(statWeightsRequest);
      results.setStatWeights(statWeightsRequest, result, epStats);
    });
  }
}
