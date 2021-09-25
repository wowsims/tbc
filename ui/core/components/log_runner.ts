import { IndividualSimRequest } from '../api/api';
import { Stat } from '../api/common';
import { StatWeightsRequest } from '../api/api';
import { Sim } from '../sim';

import { Component } from './component';
import { DetailedResults } from './detailed_results';
import { Results } from './results';

export class LogRunner extends Component {
  constructor(parent: HTMLElement, sim: Sim<any>, results: Results, detailedResults: DetailedResults) {
    super(parent, 'log-runner-root');

    const simButton = document.createElement('button');
    simButton.classList.add('log-runner-button');
    simButton.textContent = 'Sim once with logs';
    this.rootElem.appendChild(simButton);

		const logsDiv = document.createElement('div');
		logsDiv.classList.add('log-runner-logs');
		this.rootElem.appendChild(logsDiv);

    simButton.addEventListener('click', async () => {
      const simRequest = sim.makeCurrentIndividualSimRequest(1, true);

      results.setPending();
      detailedResults.setPending();
      const result = await sim.individualSim(simRequest);
      results.setSimResult(result);
      detailedResults.setSimResult(result);

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
