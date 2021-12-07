import { RaidSimRequest } from '/tbc/core/proto/api.js';
import { SimOptions } from '/tbc/core/proto/api.js';
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
      const simRequest = simUI.makeRaidSimRequest(iterations, false);

      results.setPending();
      detailedResults.setPending();
      const result = await simUI.sim.raidSim(simRequest);
      results.setSimResult(simRequest, result);
      detailedResults.setSimResult(simRequest, result);
    });

    statWeightsButton.addEventListener('click', async () => {
      const iterations = parseInt(iterationsInput.value);
      const statWeightsRequest = this.makeStatWeightsRequest(simUI, iterations, false, epStats, epReferenceStat);

      results.setPending();
      const result = await simUI.player.statWeights(statWeightsRequest);
      results.setStatWeights(statWeightsRequest, result, epStats);
    });
  }

  private makeStatWeightsRequest(simUI: SimUI<any>, iterations: number, debug: boolean, epStats: Array<Stat>, epReferenceStat: Stat): StatWeightsRequest {
		return StatWeightsRequest.create({
			player: simUI.player.toProto(),
			raidBuffs: simUI.raid.getBuffs(),
			partyBuffs: simUI.party.getBuffs(),
			encounter: simUI.encounter.toProto(),
			simOptions: SimOptions.create({
				iterations: iterations,
				debug: debug,
			}),

			statsToWeigh: epStats,
			epReferenceStat: epReferenceStat,
		});
  }
}
