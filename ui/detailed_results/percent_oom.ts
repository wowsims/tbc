import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';

import { ResultComponent, ResultComponentConfig, SimResultData } from './result_component.js';

export class PercentOom extends ResultComponent {
  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'percent-oom';
    super(config);
  }

	onSimResult(resultData: SimResultData) {
		const players = resultData.result.getPlayers(resultData.filter);

		if (players.length == 1) {
			const percentOom = players[0].oomPercent;

			this.rootElem.innerHTML = `
				<span class="percent-oom-value">${Math.round(percentOom)}%</span>
				<span class="percent-oom-label">of simulations went OOM</span>
			`;

			const dangerLevel = percentOom < 5 ? 'safe' : (percentOom < 25 ? 'warning' : 'danger');
			this.rootElem.classList.remove('safe', 'warning', 'danger');
			this.rootElem.classList.add(dangerLevel);
			this.rootElem.style.display = 'initial';
		} else {
			this.rootElem.style.display = 'none';
		}
	}
}
