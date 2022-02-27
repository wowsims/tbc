import { ResourceType } from '/tbc/core/proto/api.js';
import { resourceNames } from '/tbc/core/proto_utils/names.js';
import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { sum } from '/tbc/core/utils.js';

import { ResultComponent, ResultComponentConfig, SimResultData } from './result_component.js';

declare var $: any;
declare var tippy: any;

export class ResourceMetrics extends ResultComponent {
  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'resource-metrics-root';
    super(config);

		const resourceTypes = (getEnumValues(ResourceType) as Array<ResourceType>).filter(val => val != ResourceType.ResourceTypeNone);
		resourceTypes.forEach(resourceType => {
			const childConfig = config;
			childConfig.parent = this.rootElem;
			new ResourceMetricsTable(childConfig, resourceType);
		});
	}

	onSimResult(resultData: SimResultData) {
	}
}

export class ResourceMetricsTable extends ResultComponent {
	private readonly tableElem: HTMLTableSectionElement;
	private readonly bodyElem: HTMLTableSectionElement;
	readonly resourceType: ResourceType;

  constructor(config: ResultComponentConfig, resourceType: ResourceType) {
		config.rootCssClass = 'resource-metrics-table-root';
    super(config);
		this.resourceType = resourceType;

		this.rootElem.innerHTML = `
		<div class="resource-metrics-table-container">
			<span class="resource-metrics-table-title">${resourceNames[this.resourceType]}</span>
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
		</div>
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
		tippy(headerElems[3], {
			'content': 'Gain / Second',
			'allowHTML': true,
		});

		// Avg Gain
		tippy(headerElems[4], {
			'content': 'Gain / Event',
			'allowHTML': true,
		});

		// Avg Actual Gain
		tippy(headerElems[5], {
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

		const resourceMetrics = resultData.result.getResourceMetrics(resultData.filter, this.resourceType);
		if (resourceMetrics.length == 0) {
			this.rootElem.classList.add('hide');
		} else {
			this.rootElem.classList.remove('hide');
		}
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
