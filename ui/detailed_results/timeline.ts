import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';
import { maxIndex, sum } from '/tbc/core/utils.js';

import {
	DamageDealtLog,
	DpsLog,
	ManaChangedLog,
	SimLog,
} from '/tbc/core/proto_utils/logs_parser.js';

import { ResultComponent, ResultComponentConfig, SimResultData } from './result_component.js';

declare var $: any;
declare var tippy: any;
declare var ApexCharts: any;

const dpsColor = '#ed5653';
const manaColor = '#2E93fA';

export class Timeline extends ResultComponent {
	private readonly plotElem: HTMLElement;

	private readonly plot: any;

	private resultData: SimResultData | null;

  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'timeline-root';
    super(config);
		this.resultData = null;

		this.rootElem.innerHTML = `
		<div class="timeline-disclaimer">
			<span class="timeline-warning fa fa-exclamation-triangle"></span>
			<span class="timeline-warning-description">Timeline only includes data from 1 sim iteration.</span>
			<div class="timeline-run-again-button sim-button">SIM 1 ITERATION</span>
		</div>
		<div class="timeline-plot">
		</div>
		`;

		const runAgainButton = this.rootElem.getElementsByClassName('timeline-run-again-button')[0] as HTMLElement;
		runAgainButton.addEventListener('click', event => {
			(window.opener || window.parent)!.postMessage('runOnce', '*');
		});

		this.plotElem = this.rootElem.getElementsByClassName('timeline-plot')[0] as HTMLElement;
		this.plot = new ApexCharts(this.plotElem, {
			chart: {
				foreColor: 'white',
				animations: {
					enabled: false,
				},
				height: '100%',
				events: {
					zoomed: (charContext: any) => {
						//this.updatePlot();
					},
					scrolled: (charContext: any) => {
						//this.updatePlot();
					},
				},
			},
			series: [], // Set dynamically
			xaxis: {
				title: {
					text: 'Time (s)',
				},
			},
			yaxis: {
			},
			noData: {
				text: 'Waiting for data...',
			},
		});
		this.plot.render();
	}

	onSimResult(resultData: SimResultData) {
		this.resultData = resultData;

		this.updatePlot();

		// Doesn't work if this is called before updatePlot().
		const duration = this.resultData!.result.request.encounter!.duration || 1;
		this.plot.zoomX(0, duration);
	}

	private updatePlot() {
		const players = this.resultData!.result.getPlayers(this.resultData!.filter);
		if (players.length != 1) {
			this.plotElem.textContent = '';
			return;
		}
		const player = players[0];

		const duration = this.resultData!.result.request.encounter!.duration || 1;

		let manaLogsToShow = player.manaChangedLogs;
		let dpsLogsToShow = player.dpsLogs;
		if (manaLogsToShow.length == 0) {
			return;
		}
		const maxMana = manaLogsToShow[0].manaBefore;

		manaLogsToShow = SimLog.filterDuplicateTimestamps(manaLogsToShow);
		dpsLogsToShow = SimLog.filterDuplicateTimestamps(dpsLogsToShow);

		// Reduce to ~100 logs.
		const desiredNumLogs = 100;
		if (manaLogsToShow.length / desiredNumLogs >= 2) {
			const reductionFactor = Math.floor(manaLogsToShow.length / desiredNumLogs);
			manaLogsToShow = manaLogsToShow.filter((log, i) => i % reductionFactor == 0);
		}

		if (dpsLogsToShow.length / desiredNumLogs >= 2) {
			const reductionFactor = Math.floor(dpsLogsToShow.length / desiredNumLogs);
			dpsLogsToShow = dpsLogsToShow.filter((log, i) => i % reductionFactor == 0);
		}

		const maxDps = dpsLogsToShow[maxIndex(dpsLogsToShow.map(l => l.dps))!].dps;
		const dpsAxisMax = (Math.floor(maxDps / 100) + 1) * 100;

		this.plot.updateOptions({
			colors: [
				dpsColor,
				manaColor,
			],
			series: [
				{
					name: 'DPS',
					type: 'line',
					data: dpsLogsToShow.map(log => {
						return {
							x: log.timestamp,
							y: log.dps,
						};
					}),
				},
				{
					name: 'Mana',
					type: 'line',
					data: manaLogsToShow.map(log => {
						return {
							x: log.timestamp,
							y: log.manaAfter,
						};
					}),
				},
			],
			xaxis: {
				min: 0,
				max: duration,
				tickAmount: 10,
				categories: manaLogsToShow.map(log => log.timestamp),
				labels: {
					show: true,
					formatter: (val: string) => val,
				},
			},
			yaxis: [
				{
					color: dpsColor,
					seriesName: 'DPS',
					min: 0,
					max: dpsAxisMax,
					tickAmount: 10,
					decimalsInFloat: 0,
					title: {
						text: 'DPS',
						style: {
							color: dpsColor,
						},
					},
					axisBorder: {
						show: true,
						color: dpsColor,
					},
					axisTicks: {
						color: dpsColor,
					},
					labels: {
						style: {
							colors: [dpsColor],
						},
					},
				},
				{
					seriesName: 'Mana',
					opposite: true, // Appear on right side
					min: 0,
					max: maxMana,
					tickAmount: 10,
					title: {
						text: 'Mana',
						style: {
							color: manaColor,
						},
					},
					axisBorder: {
						show: true,
						color: manaColor,
					},
					axisTicks: {
						color: manaColor,
					},
					labels: {
						style: {
							colors: [manaColor],
						},
						formatter: (val: string) => {
							const v = parseFloat(val);
							return `${v.toFixed(0)} (${(v/maxMana*100).toFixed(0)}%)`;
						},
					},
				},
			],
			dataLabels: {
			},
		});
	}

	render() {
		setTimeout(() => this.plot.render(), 300);
	}
}
