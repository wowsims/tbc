import { IndividualSimResult } from '../core/api/api';
import { Component } from '../core/components/component';
import { TypedEvent } from '../core/typed_event';

declare var Chart: any;

export class DpsHistogram extends Component {
  constructor(parent: HTMLElement, resultsEmitter: TypedEvent<IndividualSimResult | null>) {
    super(parent, 'dps-histogram-root');

		resultsEmitter.on(simResult => {
			if (!simResult)
				return;

			const bounds = this.rootElem.getBoundingClientRect();
			const chartCanvas = this.createDpsChartFromSimResult(simResult, bounds);

			this.rootElem.textContent = '';
			this.rootElem.appendChild(chartCanvas);
		});
	}

	private createDpsChartFromSimResult(simResult: IndividualSimResult, chartBounds: DOMRect): HTMLCanvasElement {
		const chartCanvas = document.createElement("canvas");
		chartCanvas.height = chartBounds.height;
		chartCanvas.width = chartBounds.width;


		const min = simResult.dpsAvg - simResult.dpsStdev;
		const max = simResult.dpsAvg + simResult.dpsStdev;
		const vals: Array<number> = [];
		const colors: Array<string> = [];

		const labels = Object.keys(simResult.dpsHist);
		labels.forEach((k, i) => {
			vals.push(simResult.dpsHist[Number(k)]);
			const val = parseInt(k);
			if (val > min && val < max) {
				colors.push('#1E87F0');
			} else {
				colors.push('#FF6961');
			}
		});

		const ctx = chartCanvas.getContext('2d');
		const chart = new Chart(ctx, {
			type: 'bar',
			data: {
				labels: labels,
				datasets: [{
					data: vals,
					backgroundColor: colors,
				}],
			},
			options: {
				plugins: {
					title: {
						display: true,
						text: 'DPS Histogram',
					},
					legend: {
						display: false,
						labels: {},
					}
				},
				scales: {
					y: {
						beginAtZero: true,
						ticks: {
							display: false
						},
					},
				},
			},
		});
		return chartCanvas;
	}
}
