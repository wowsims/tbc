import { RaidSimRequest } from '/tbc/core/proto/api.js';
import { Stat } from '/tbc/core/proto/common.js';
import { StatWeightsRequest } from '/tbc/core/proto/api.js';
import { SimUI } from '/tbc/core/sim_ui.js';

import { Component } from './component.js';
import { DetailedResults } from './detailed_results.js';
import { Results } from './results.js';

export class LogRunner extends Component {
  constructor(parent: HTMLElement, simUI: SimUI<any>, results: Results, detailedResults: DetailedResults) {
    super(parent, 'log-runner-root');

		const controlBar = document.createElement('div');
		controlBar.classList.add('log-runner-control-bar');
		this.rootElem.appendChild(controlBar);

    const simButton = document.createElement('button');
    simButton.classList.add('log-runner-button');
    simButton.textContent = 'Sim once with logs';
    controlBar.appendChild(simButton);

		const logsDiv = document.createElement('div');
		logsDiv.classList.add('log-runner-logs');
		this.rootElem.appendChild(logsDiv);

    simButton.addEventListener('click', async () => {
      const simRequest = simUI.makeRaidSimRequest(1, true);

      results.setPending();
      detailedResults.setPending();
      const result = await simUI.sim.raidSim(simRequest);
      results.setSimResult(simRequest, result);
      detailedResults.setSimResult(simRequest, result);

			const lines = result.logs.split('\n');
			logsDiv.textContent = '';
			lines.forEach(line => {
				const lineElem = document.createElement('span');
				lineElem.textContent = line;
				logsDiv.appendChild(lineElem);
			});
    });
  }
}
