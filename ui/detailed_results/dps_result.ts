import { IndividualSimRequest, IndividualSimResult } from '../core/proto/api.js';

import { ResultComponent, ResultComponentConfig } from './result_component.js';

export class DpsResult extends ResultComponent {
  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'dps-result';
    super(config);
  }

	onSimResult(request: IndividualSimRequest, result: IndividualSimResult) {
    this.rootElem.innerHTML = `
      <span class="results-sim-dps-avg">${result.dpsAvg.toFixed(2)}</span>
      <span class="results-sim-dps-stdev">${result.dpsStdev.toFixed(2)}</span>
    `;
	}
}
