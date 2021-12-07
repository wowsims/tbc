import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { setWowheadHref } from '/tbc/core/resources.js';
import { sum } from '/tbc/core/utils.js';

import { parseActionMetrics } from './metrics_helpers.js';
import { ResultComponent, ResultComponentConfig } from './result_component.js';

declare var $: any;
declare var tippy: any;

export class SpellMetrics extends ResultComponent {
	private readonly tableElem: HTMLTableSectionElement;
	private readonly bodyElem: HTMLTableSectionElement;

  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'cast-metrics-root';
    super(config);

		this.rootElem.innerHTML = `
		<table class="cast-metrics-table tablesorter">
			<thead class="cast-metrics-table-header">
				<tr class="cast-metrics-table-header-row">
					<th class="cast-metrics-table-header-cell"><span>Name</span></th>
					<th class="cast-metrics-table-header-cell"><span>DPS</span></th>
					<th class="cast-metrics-table-header-cell"><span>Casts</span></th>
					<th class="cast-metrics-table-header-cell"><span>Avg Cast</span></th>
					<th class="cast-metrics-table-header-cell"><span>Hits</span></th>
					<th class="cast-metrics-table-header-cell"><span>Avg Hit</span></th>
					<th class="cast-metrics-table-header-cell"><span>Crit %</span></th>
					<th class="cast-metrics-table-header-cell"><span>Miss %</span></th>
				</tr>
			</thead>
			<tbody class="cast-metrics-table-body">
			</tbody>
		</table>
		`;

		this.tableElem = this.rootElem.getElementsByClassName('cast-metrics-table')[0] as HTMLTableSectionElement;
		this.bodyElem = this.rootElem.getElementsByClassName('cast-metrics-table-body')[0] as HTMLTableSectionElement;

		const headerElems = Array.from(this.tableElem.querySelectorAll('th'));

		// DPS
		tippy(headerElems[1], {
			'content': 'Damage / Encounter Duration',
			'allowHTML': true,
		});

		// Casts
		tippy(headerElems[2], {
			'content': 'Casts',
			'allowHTML': true,
		});

		// Avg Cast
		tippy(headerElems[3], {
			'content': 'Damage / Casts',
			'allowHTML': true,
		});

		// Hits
		tippy(headerElems[4], {
			'content': 'Hits',
			'allowHTML': true,
		});

		// Avg Hit
		tippy(headerElems[5], {
			'content': 'Damage / Hits',
			'allowHTML': true,
		});

		// Crit %
		tippy(headerElems[6], {
			'content': 'Crits / Hits',
			'allowHTML': true,
		});

		// Miss %
		tippy(headerElems[7], {
			'content': 'Misses / (Hits + Misses)',
			'allowHTML': true,
		});

		$(this.tableElem).tablesorter({ sortList: [[1, 1]] });
	}

	onSimResult(request: RaidSimRequest, result: RaidSimResult) {
		this.bodyElem.textContent = '';

		const iterations = request.simOptions!.iterations;
		const duration = request.encounter?.duration || 1;

		parseActionMetrics(result.raidMetrics!.parties[0].players[0].actions).then(actionMetrics => {
			actionMetrics.filter(e => e.hits + e.misses != 0).forEach(actionMetric => {

				const rowElem = document.createElement('tr');
				this.bodyElem.appendChild(rowElem);

				const nameCellElem = document.createElement('td');
				rowElem.appendChild(nameCellElem);
				nameCellElem.innerHTML = `
				<a class="cast-metrics-action-icon"></a>
				<span class="cast-metrics-action-name">${actionMetric.name}</span>
				`;

				const iconElem = nameCellElem.getElementsByClassName('cast-metrics-action-icon')[0] as HTMLAnchorElement;
				iconElem.style.backgroundImage = `url('${actionMetric.iconUrl}')`;
				if (!('otherId' in actionMetric.actionId.id)) {
					setWowheadHref(iconElem, actionMetric.actionId.id);
				}

				const addCell = (value: string | number): HTMLElement => {
					const cellElem = document.createElement('td');
					cellElem.textContent = String(value);
					rowElem.appendChild(cellElem);
					return cellElem;
				};

				addCell((actionMetric.totalDmg / iterations / duration).toFixed(1)); // DPS
				addCell((actionMetric.casts / iterations).toFixed(1)); // Casts
				addCell((actionMetric.totalDmg / actionMetric.casts).toFixed(1)); // Avg Cast
				addCell((actionMetric.hits / iterations).toFixed(1)); // Hits
				addCell((actionMetric.totalDmg / actionMetric.hits).toFixed(1)); // Avg Hit
				addCell(((actionMetric.crits / actionMetric.hits) * 100).toFixed(2) + ' %'); // Crit %
				addCell(((actionMetric.misses / (actionMetric.hits + actionMetric.misses)) * 100).toFixed(2) + ' %'); // Miss %
			});

			$(this.tableElem).trigger('update');
		});
	}
}

// For the no-damage casts
export class CastMetrics extends ResultComponent {
	private readonly tableElem: HTMLTableSectionElement;
	private readonly bodyElem: HTMLTableSectionElement;

  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'other-cast-metrics-root';
    super(config);

		this.rootElem.innerHTML = `
		<table class="cast-metrics-table tablesorter">
			<thead class="cast-metrics-table-header">
				<tr class="cast-metrics-table-header-row">
					<th class="cast-metrics-table-header-cell"><span>Name</span></th>
					<th class="cast-metrics-table-header-cell"><span>Casts</span></th>
					<th class="cast-metrics-table-header-cell"><span>CPM</span></th>
				</tr>
			</thead>
			<tbody class="cast-metrics-table-body">
			</tbody>
		</table>
		`;

		this.tableElem = this.rootElem.getElementsByClassName('cast-metrics-table')[0] as HTMLTableSectionElement;
		this.bodyElem = this.rootElem.getElementsByClassName('cast-metrics-table-body')[0] as HTMLTableSectionElement;

		const headerElems = Array.from(this.tableElem.querySelectorAll('th'));

		// Casts
		tippy(headerElems[1], {
			'content': 'Casts',
			'allowHTML': true,
		});

		// CPM
		tippy(headerElems[2], {
			'content': 'Casts / (Encounter Duration / 60 Seconds)',
			'allowHTML': true,
		});

		$(this.tableElem).tablesorter({ sortList: [[1, 1]] });
	}

	onSimResult(request: RaidSimRequest, result: RaidSimResult) {
		this.bodyElem.textContent = '';

		const iterations = request.simOptions!.iterations;
		const duration = request.encounter?.duration || 1;

		parseActionMetrics(result.raidMetrics!.parties[0].players[0].actions).then(actionMetrics => {
			actionMetrics.forEach(actionMetric => {
				const rowElem = document.createElement('tr');
				this.bodyElem.appendChild(rowElem);

				const nameCellElem = document.createElement('td');
				rowElem.appendChild(nameCellElem);
				nameCellElem.innerHTML = `
				<a class="cast-metrics-action-icon"></a>
				<span class="cast-metrics-action-name">${actionMetric.name}</span>
				`;

				const iconElem = nameCellElem.getElementsByClassName('cast-metrics-action-icon')[0] as HTMLAnchorElement;
				iconElem.style.backgroundImage = `url('${actionMetric.iconUrl}')`;
				if (!('otherId' in actionMetric.actionId.id)) {
					setWowheadHref(iconElem, actionMetric.actionId.id);
				}

				const addCell = (value: string | number): HTMLElement => {
					const cellElem = document.createElement('td');
					cellElem.textContent = String(value);
					rowElem.appendChild(cellElem);
					return cellElem;
				};

				addCell((actionMetric.casts / iterations).toFixed(1)); // Casts
				addCell((actionMetric.casts / iterations / (duration / 60)).toFixed(1)); // CPM
			});

			$(this.tableElem).trigger('update');
		});
	}
}
