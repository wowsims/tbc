import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';

import { ResultComponent, ResultComponentConfig, SimResultData } from './result_component.js';

export class DpsResult extends ResultComponent {
  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'dps-result';
    super(config);
  }

	onSimResult(resultData: SimResultData) {
		const damageMetrics = resultData.result.getDamageMetrics(resultData.filter);

    this.rootElem.innerHTML = `
      <span class="results-sim-dps-avg">${damageMetrics.avg.toFixed(2)}</span>
      <span class="results-sim-dps-stdev">${damageMetrics.stdev.toFixed(2)}</span>
    `;
	}
}
