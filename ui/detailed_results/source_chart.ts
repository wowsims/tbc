import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';
import { sum } from '/tbc/core/utils.js';

import { ResultComponent, ResultComponentConfig } from './result_component.js';

declare var Chart: any;

export class SourceChart extends ResultComponent {
  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'source-chart-root';
    super(config);
	}

	onSimResult(request: IndividualSimRequest, result: IndividualSimResult) {
		const chartBounds = this.rootElem.getBoundingClientRect();

		this.rootElem.textContent = '';
		const chartCanvas = document.createElement("canvas");
		chartCanvas.height = chartBounds.height;
		chartCanvas.width = chartBounds.width;

		const colors: Array<string> = ['red', 'blue', 'lawngreen'];

		const labels = Object.keys(result.casts);
		const totalDmg = sum(Object.values(result.casts).map(castMetrics => castMetrics.dmgs[0]));
		const vals = labels.map(label => result.casts[parseInt(label)].dmgs[0] / totalDmg);

		const ctx = chartCanvas.getContext('2d');
		const chart = new Chart(ctx, {
			type: 'pie',
			data: {
				labels: labels,
				datasets: [{
					data: vals,
					backgroundColor: colors,
				}],
			},
			options: {
				plugins: {
					legend: {
						display: true,
						position: 'right',
					}
				},
			},
		});
		this.rootElem.appendChild(chartCanvas);
	}
}
