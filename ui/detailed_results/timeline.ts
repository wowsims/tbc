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
	private readonly dpsResourcesPlotElem: HTMLElement;
	private readonly cooldownsPlotElem: HTMLElement;
	private dpsResourcesPlot: any;
	private cooldownsPlot: any;

	private resultData: SimResultData | null;

  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'timeline-root';
    super(config);
		this.resultData = null;

		this.rootElem.innerHTML = `
		<div class="timeline-disclaimer">
			<span class="timeline-warning fa fa-exclamation-triangle"></span>
			<span class="timeline-warning-description">Timeline data visualizes only 1 sim iteration.</span>
			<div class="timeline-run-again-button sim-button">SIM 1 ITERATION</span>
		</div>
		<div class="timeline-plots-container">
			<div class="timeline-plot dps-resources-plot"></div>
			<div class="timeline-plot cooldowns-plot"></div>
		</div>
		`;

		const runAgainButton = this.rootElem.getElementsByClassName('timeline-run-again-button')[0] as HTMLElement;
		runAgainButton.addEventListener('click', event => {
			(window.opener || window.parent)!.postMessage('runOnce', '*');
		});

		this.dpsResourcesPlotElem = this.rootElem.getElementsByClassName('dps-resources-plot')[0] as HTMLElement;

		this.cooldownsPlotElem = this.rootElem.getElementsByClassName('cooldowns-plot')[0] as HTMLElement;
		this.cooldownsPlot = new ApexCharts(this.cooldownsPlotElem, {
			chart: {
				type: 'rangeBar',
				foreColor: 'white',
				id: 'cooldowns',
				animations: {
					enabled: false,
				},
				height: '50%',
			},
			series: [], // Set dynamically
			noData: {
				text: 'Waiting for data...',
			},
		});

		//this.dpsResourcesPlot.render();
		//this.cooldownsPlot.render();
	}

	onSimResult(resultData: SimResultData) {
		this.resultData = resultData;

		this.updatePlot();
	}

	private updatePlot() {
		const players = this.resultData!.result.getPlayers(this.resultData!.filter);
		if (players.length != 1) {
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

		const maxDps = dpsLogsToShow[maxIndex(dpsLogsToShow.map(l => l.dps))!].dps;
		const dpsAxisMax = (Math.floor(maxDps / 100) + 1) * 100;

		this.dpsResourcesPlot = new ApexCharts(this.dpsResourcesPlotElem, {
		//	chart: {
		//		type: 'line',
		//		foreColor: 'white',
		//		id: 'dpsResources',
		//		animations: {
		//			enabled: false,
		//		},
		//		height: '50%',
		//	},
		//	series: [], // Set dynamically
		//	xaxis: {
		//		title: {
		//			text: 'Time (s)',
		//		},
		//		type: 'datetime',
		//	},
		//	yaxis: {
		//	},
		//	noData: {
		//		text: 'Waiting for data...',
		//	},
		//	stroke: {
		//		width: 2,
		//		curve: 'straight',
		//	},
		//});
		//this.dpsResourcesPlot.updateOptions({
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
							x: this.toDatetime(log.timestamp),
							y: log.dps,
						};
					}),
				},
				{
					name: 'Mana',
					type: 'line',
					data: manaLogsToShow.map(log => {
						return {
							x: this.toDatetime(log.timestamp),
							y: log.manaAfter,
						};
					}),
				},
			],
			xaxis: {
				min: this.toDatetime(0).getTime(),
				max: this.toDatetime(duration).getTime(),
				type: 'datetime',
				tickAmount: 10,
				decimalsInFloat: 1,
				labels: {
					show: true,
					formatter: (defaultValue: string, timestamp: number) => {
						return (timestamp/1000).toFixed(1);
					},
				},
				title: {
					text: 'Time (s)',
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
						minWidth: 30,
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
						minWidth: 30,
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
			tooltip: {
				enabled: true,
			},
			chart: {
				events: {
					beforeResetZoom: () => {
						return {
							xaxis: {
								min: this.toDatetime(0),
								max: this.toDatetime(duration),
							},
						};
					},
				},
				type: 'line',
				foreColor: 'white',
				id: 'dpsResources',
				animations: {
					enabled: false,
				},
				height: '50%',
			},
			annotations: {
				position: 'back',
				xaxis: [
					{
						x: this.toDatetime(20).getTime(),
						x2: this.toDatetime(60).getTime(),
						fillColor: '#B3F7CA',
					},
					{
						x: this.toDatetime(40).getTime(),
						x2: this.toDatetime(80).getTime(),
						fillColor: '#ff742e',
					},
				],
				points: [
					{
						x: this.toDatetime(60).getTime(),
						y: 0,
						image: {
							path: 'https://wow.zamimg.com/images/wow/icons/medium/spell_nature_starfall.jpg',
							width: 20,
							height: 20,
						},
					},
					{
						x: this.toDatetime(80).getTime(),
						y: 0,
						image: {
							path: 'https://wow.zamimg.com/images/wow/icons/medium/spell_nature_starfall.jpg',
							width: 20,
							height: 20,
						},
					},
					{
						x: this.toDatetime(80).getTime(),
						y: 0,
						image: {
							path: 'https://wow.zamimg.com/images/wow/icons/medium/ability_racial_bearform.jpg',
							width: 20,
							height: 20,
							offsetY: -25,
						},
					},
				],
			},
			stroke: {
				width: 2,
				curve: 'straight',
			},
		});
		this.dpsResourcesPlot.render();

		//this.cooldownsPlot.updateOptions({
		//	series: [
		//		{
		//			name: 'Lightning Bolt',
		//			data: [
		//				{
		//					x: 'GCD',
		//					y: [0, 40],
		//				},
		//				{
		//					x: 'GCD',
		//					y: [60, 100],
		//				},
		//			],
		//		},
		//		{
		//			name: 'Chain Lightning',
		//			data: [
		//				{
		//					x: 'GCD',
		//					y: [0, 40],
		//				},
		//				{
		//					x: 'GCD',
		//					y: [60, 100],
		//				},
		//			],
		//		},
		//		{
		//			name: 'Bloodlust',
		//			data: [
		//				{
		//					x: 'Cooldowns',
		//					y: [0, 40],
		//				},
		//				{
		//					x: 'Cooldowns',
		//					y: [60, 100],
		//				},
		//			],
		//		},
		//		{
		//			name: 'Innervate',
		//			data: [
		//				{
		//					x: 'Cooldowns',
		//					y: [30, 70],
		//				},
		//				{
		//					x: 'Cooldowns',
		//					y: [150, 200],
		//				},
		//			],
		//		},
		//	],
		//	xaxis: {
		//		min: this.toDatetime(0),
		//		max: this.toDatetime(duration),
		//		tickAmount: 10,
		//		decimalsInFloat: 1,
		//		labels: {
		//			show: true,
		//		},
		//	},
		//	yaxis: {
		//		title: {
		//			text: 'Cooldowns',
		//		},
		//		labels: {
		//			minWidth: 30,
		//		},
		//	},
		//	plotOptions: {
		//		bar: {
		//			horizontal: true,
		//			barHeight: '80%',
		//		},
		//	},
		//	stroke: {
		//		width: 1,
		//	},
		//	fill: {
		//		type: 'solid',
		//		opacity: 0.6,
		//	},
		//	tooltip: {
		//		enabled: true,
		//	},
		//	chart: {
		//		events: {
		//			beforeResetZoom: () => {
		//				return {
		//					xaxis: {
		//						min: this.toDatetime(0),
		//						max: this.toDatetime(duration),
		//					},
		//				};
		//			},
		//		},
		//	},
		//});
	}

	render() {
		setTimeout(() => {
			//this.dpsResourcesPlot = new ApexCharts(this.dpsResourcesPlotElem, {
			//	chart: {
			//		type: 'line',
			//		foreColor: 'white',
			//		id: 'dpsResources',
			//		animations: {
			//			enabled: false,
			//		},
			//		height: '50%',
			//	},
			//	series: [], // Set dynamically
			//	xaxis: {
			//		title: {
			//			text: 'Time (s)',
			//		},
			//		type: 'datetime',
			//	},
			//	yaxis: {
			//	},
			//	noData: {
			//		text: 'Waiting for data...',
			//	},
			//	stroke: {
			//		width: 2,
			//		curve: 'straight',
			//	},
			//});
			//this.dpsResourcesPlot.render();
			this.cooldownsPlot.render();
		}, 300);
	}

	private toDatetime(timestamp: number): Date {
		//return timestamp;
		return new Date(timestamp * 1000);
		//return timestamp * 1000;
	}
}
