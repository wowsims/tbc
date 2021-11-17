import { IndividualSimRequest } from '/tbc/core/proto/api.js';
import { Stat } from '/tbc/core/proto/common.js';
import { StatWeightsRequest } from '/tbc/core/proto/api.js';
import { SimUI } from '/tbc/core/sim_ui.js';

import { Component } from './component.js';
import { DetailedResults } from './detailed_results.js';
import { Results } from './results.js';

export class Actions extends Component {
  constructor(parent: HTMLElement, simUI: SimUI<any>, epStats: Array<Stat>, epReferenceStat: Stat, results: Results, detailedResults: DetailedResults) {
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
      <input class="iterations-input" type="number" min="1" value="3000" step="1000">
    `;
    this.rootElem.appendChild(iterationsDiv);
    const iterationsInput = iterationsDiv.getElementsByClassName('iterations-input')[0] as HTMLInputElement;

    simButton.addEventListener('click', async () => {
      const iterations = parseInt(iterationsInput.value);
      const simRequest = simUI.makeCurrentIndividualSimRequest(iterations, false);

      results.setPending();
      detailedResults.setPending();
      const result = await simUI.sim.individualSim(simRequest);
      results.setSimResult(simRequest, result);
      detailedResults.setSimResult(simRequest, result);
    });

    statWeightsButton.addEventListener('click', async () => {
      const iterations = parseInt(iterationsInput.value);
      const simRequest = simUI.makeCurrentIndividualSimRequest(iterations, false);

      const statWeightsRequest = StatWeightsRequest.create({
        options: simRequest,
        statsToWeigh: epStats,
        epReferenceStat: epReferenceStat,
      });

      results.setPending();
      const result = await simUI.player.statWeights(statWeightsRequest);
      results.setStatWeights(statWeightsRequest, result, epStats);
    });
  }
}
