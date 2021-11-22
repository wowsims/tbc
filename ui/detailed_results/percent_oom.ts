import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';

import { ResultComponent, ResultComponentConfig } from './result_component.js';

export class PercentOom extends ResultComponent {
  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'percent-oom';
    super(config);
  }

	onSimResult(request: IndividualSimRequest, result: IndividualSimResult) {
		const percentOom = result.playerMetrics!.numOom / request.simOptions!.iterations;

    this.rootElem.innerHTML = `
      <span class="percent-oom-value">${Math.round(percentOom * 100)}%</span>
      <span class="percent-oom-label">of simulations went OOM</span>
    `;

		const dangerLevel = percentOom < 0.05 ? 'safe' : (percentOom < 0.25 ? 'warning' : 'danger');
		this.rootElem.classList.remove('safe', 'warning', 'danger');
		this.rootElem.classList.add(dangerLevel);
	}
}
