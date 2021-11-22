import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';

import { ResultComponent, ResultComponentConfig } from './result_component.js';

export class DpsResult extends ResultComponent {
  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'dps-result';
    super(config);
  }

	onSimResult(request: IndividualSimRequest, result: IndividualSimResult) {
    this.rootElem.innerHTML = `
      <span class="results-sim-dps-avg">${result.playerMetrics!.dpsAvg.toFixed(2)}</span>
      <span class="results-sim-dps-stdev">${result.playerMetrics!.dpsStdev.toFixed(2)}</span>
    `;
	}
}
