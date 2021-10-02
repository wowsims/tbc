import { IndividualSimRequest, IndividualSimResult } from '../core/api/api.js';

import { ResultComponent, ResultComponentConfig } from './result_component.js';

export class PercentOom extends ResultComponent {
  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'percent-oom';
    super(config);
  }

	onSimResult(request: IndividualSimRequest, result: IndividualSimResult) {
		const percentOom = Math.round(result.numOom / request.iterations);

    this.rootElem.innerHTML = `
      <span class="percent-oom-value">${percentOom}%</span>
      <span class="percent-oom-label">of simulations went OOM</span>
    `;

		const dangerLevel = percentOom < 0.05 ? 'safe' : (percentOom < 0.25 ? 'warning' : 'danger');
		this.rootElem.classList.remove('safe', 'warning', 'danger');
		this.rootElem.classList.add(dangerLevel);
	}
}
