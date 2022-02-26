import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';
import { sum } from '/tbc/core/utils.js';

import { ResultComponent, ResultComponentConfig, SimResultData } from './result_component.js';

declare var $: any;
declare var tippy: any;

// For the no-damage casts
export class ResourceMetrics extends ResultComponent {
	private readonly tableElem: HTMLTableSectionElement;
	private readonly bodyElem: HTMLTableSectionElement;

  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'resource-metrics-root';
    super(config);

		this.rootElem.innerHTML = `
		<table class="metrics-table tablesorter">
			<thead class="metrics-table-header">
				<tr class="metrics-table-header-row">
					<th class="metrics-table-header-cell"><span>Name</span></th>
					<th class="metrics-table-header-cell"><span>Casts</span></th>
					<th class="metrics-table-header-cell"><span>Gain</span></th>
					<th class="metrics-table-header-cell"><span>Gain / s</span></th>
					<th class="metrics-table-header-cell"><span>Avg Gain</span></th>
					<th class="metrics-table-header-cell"><span>Avg Actual Gain</span></th>
				</tr>
			</thead>
			<tbody class="metrics-table-body">
			</tbody>
		</table>
		`;

		this.tableElem = this.rootElem.getElementsByClassName('metrics-table')[0] as HTMLTableSectionElement;
		this.bodyElem = this.rootElem.getElementsByClassName('metrics-table-body')[0] as HTMLTableSectionElement;

		const headerElems = Array.from(this.tableElem.querySelectorAll('th'));

		// Casts
		tippy(headerElems[1], {
			'content': 'Casts',
			'allowHTML': true,
		});

		// GPS
		tippy(headerElems[2], {
			'content': 'Gain',
			'allowHTML': true,
		});

		// GPS
		tippy(headerElems[2], {
			'content': 'Gain / Second',
			'allowHTML': true,
		});

		// Avg Gain
		tippy(headerElems[3], {
			'content': 'Gain / Event',
			'allowHTML': true,
		});

		// Avg Actual Gain
		tippy(headerElems[4], {
			'content': 'Actual Gain / Event',
			'allowHTML': true,
		});

		$(this.tableElem).tablesorter({
			sortList: [[2, 1]],
			cssChildRow: 'child-metric',
		});
	}

	onSimResult(resultData: SimResultData) {
		this.bodyElem.textContent = '';

		const resourceMetrics = resultData.result.getResourceMetrics(resultData.filter);
		resourceMetrics.forEach(resourceMetric => {
			const rowElem = document.createElement('tr');
			this.bodyElem.appendChild(rowElem);

			const nameCellElem = document.createElement('td');
			rowElem.appendChild(nameCellElem);
			nameCellElem.innerHTML = `
			<a class="metrics-action-icon"></a>
			<span class="metrics-action-name">${resourceMetric.name}</span>
			`;

			const iconElem = nameCellElem.getElementsByClassName('metrics-action-icon')[0] as HTMLAnchorElement;
			resourceMetric.actionId.setBackgroundAndHref(iconElem);

			const addCell = (value: string | number): HTMLElement => {
				const cellElem = document.createElement('td');
				cellElem.textContent = String(value);
				rowElem.appendChild(cellElem);
				return cellElem;
			};

			addCell(resourceMetric.events.toFixed(1)); // Casts
			addCell(resourceMetric.gain.toFixed(1)); // GPS
			addCell(resourceMetric.gainPerSecond.toFixed(1)); // GPS
			addCell(resourceMetric.avgGain.toFixed(1)); // Avg Gain
			addCell(resourceMetric.avgActualGain.toFixed(1)); // Avg Actual Gain
		});

		$(this.tableElem).trigger('update');
	}
}
