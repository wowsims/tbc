import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';

import { ResultComponent, ResultComponentConfig } from './result_component.js';

export class DpsResult extends ResultComponent {
  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'dps-result';
    super(config);
  }

	onSimResult(request: RaidSimRequest, result: RaidSimResult) {
    this.rootElem.innerHTML = `
      <span class="results-sim-dps-avg">${result.raidMetrics!.dps!.avg.toFixed(2)}</span>
      <span class="results-sim-dps-stdev">${result.raidMetrics!.dps!.stdev.toFixed(2)}</span>
    `;
	}
}
