import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';
import { maxIndex } from '/tbc/core/utils.js';
import { sum } from '/tbc/core/utils.js';

import { ResultComponent, ResultComponentConfig, SimResultData } from './result_component.js';
import { ResultsFilter } from './results_filter.js';
import { SourceChart } from './source_chart.js';

declare var $: any;
declare var tippy: any;

export class PlayerDamageMetrics extends ResultComponent {
	private readonly tableElem: HTMLTableSectionElement;
	private readonly bodyElem: HTMLTableSectionElement;
	private readonly resultsFilter: ResultsFilter;

  constructor(config: ResultComponentConfig, resultsFilter: ResultsFilter) {
		config.rootCssClass = 'player-damage-metrics-root';
    super(config);
		this.resultsFilter = resultsFilter;

		this.rootElem.innerHTML = `
		<table class="metrics-table tablesorter">
			<thead class="metrics-table-header">
				<tr class="metrics-table-header-row">
					<th class="metrics-table-header-cell"><span>Player</span></th>
					<th class="metrics-table-header-cell amount-header-cell"><span>Amount</span></th>
					<th class="metrics-table-header-cell"><span>DPS</span></th>
				</tr>
			</thead>
			<tbody class="metrics-table-body">
			</tbody>
		</table>
		`;

		this.tableElem = this.rootElem.getElementsByClassName('metrics-table')[0] as HTMLTableSectionElement;
		this.bodyElem = this.rootElem.getElementsByClassName('metrics-table-body')[0] as HTMLTableSectionElement;

		const headerElems = Array.from(this.tableElem.querySelectorAll('th'));

		// Amount
		tippy(headerElems[1], {
			'content': 'Player Damage / Raid Damage',
			'allowHTML': true,
		});

		// DPS
		tippy(headerElems[2], {
			'content': 'Damage / Encounter Duration',
			'allowHTML': true,
		});

		$(this.tableElem).tablesorter({ sortList: [[2, 1]] });
	}

	onSimResult(resultData: SimResultData) {
		this.bodyElem.textContent = '';

		const raidDps = resultData.result.raidMetrics.dps.avg;
		const players = resultData.result.getPlayers(resultData.filter);
		if (players.length == 0) {
			return;
		}

		const maxDpsIndex = maxIndex(players.map(player => player.dps.avg))!;
		const maxDps = players[maxDpsIndex].dps.avg;

		players.forEach(player => {
			const rowElem = document.createElement('tr');
			rowElem.classList.add('player-damage-row');
			this.bodyElem.appendChild(rowElem);
			rowElem.addEventListener('click', event => {
				this.resultsFilter.setPlayer(resultData.eventID, player.raidIndex);
			});

			let chart: HTMLElement | null = null;
			const makeChart = () => {
				const chartContainer = document.createElement('div');
				rowElem.appendChild(chartContainer);
				const sourceChart = new SourceChart(chartContainer, player.actions);
				return chartContainer;
			};

			tippy(rowElem, {
				content: 'Loading...',
				placement: 'bottom',
				onShow(instance: any) {
					if (!chart) {
						chart = makeChart();
						instance.setContent(chart);
					}
				},
			});

			const nameCellElem = document.createElement('td');
			rowElem.appendChild(nameCellElem);
			nameCellElem.innerHTML = `
			<img class="metrics-action-icon" src="${player.iconUrl}"></img>
			<span class="metrics-action-name" style="color:${player.classColor}">${player.label}</span>
			`;

			const amountCellElem = document.createElement('td');
			amountCellElem.classList.add('amount-cell');
			rowElem.appendChild(amountCellElem);
			amountCellElem.innerHTML = `
				<div class="player-damage-percent">
					<span>${(player.dps.avg / raidDps * 100).toFixed(2)}%</span>
				</div>
				<div class="player-damage-bar-container">
					<div class="player-damage-bar" style="background-color:${player.classColor}; width:${player.dps.avg / maxDps * 100}%"></div>
				</div>
				<div class="player-damage-total">
					<span>${(player.totalDamage / 1000).toFixed(1)}k</span>
				</div>
			`;

			const dpsCellElem = document.createElement('td');
			rowElem.appendChild(dpsCellElem);
			dpsCellElem.textContent = player.dps.avg.toFixed(1);
		});

		$(this.tableElem).trigger('update');
	}
}
