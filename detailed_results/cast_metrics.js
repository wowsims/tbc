import { setWowheadHref } from '/tbc/core/resources.js';
import { ResultComponent } from './result_component.js';
// For the no-damage casts
export class CastMetrics extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'other-cast-metrics-root';
        super(config);
        this.rootElem.innerHTML = `
		<table class="metrics-table tablesorter">
			<thead class="metrics-table-header">
				<tr class="metrics-table-header-row">
					<th class="metrics-table-header-cell"><span>Name</span></th>
					<th class="metrics-table-header-cell"><span>Casts</span></th>
					<th class="metrics-table-header-cell"><span>CPM</span></th>
				</tr>
			</thead>
			<tbody class="metrics-table-body">
			</tbody>
		</table>
		`;
        this.tableElem = this.rootElem.getElementsByClassName('metrics-table')[0];
        this.bodyElem = this.rootElem.getElementsByClassName('metrics-table-body')[0];
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
    onSimResult(resultData) {
        this.bodyElem.textContent = '';
        const actionMetrics = resultData.result.getActionMetrics(resultData.filter);
        actionMetrics.forEach(actionMetric => {
            const rowElem = document.createElement('tr');
            this.bodyElem.appendChild(rowElem);
            const nameCellElem = document.createElement('td');
            rowElem.appendChild(nameCellElem);
            nameCellElem.innerHTML = `
			<a class="metrics-action-icon"></a>
			<span class="metrics-action-name">${actionMetric.name}</span>
			`;
            const iconElem = nameCellElem.getElementsByClassName('metrics-action-icon')[0];
            iconElem.style.backgroundImage = `url('${actionMetric.iconUrl}')`;
            if (!('otherId' in actionMetric.actionId.id)) {
                setWowheadHref(iconElem, actionMetric.actionId.id);
            }
            const addCell = (value) => {
                const cellElem = document.createElement('td');
                cellElem.textContent = String(value);
                rowElem.appendChild(cellElem);
                return cellElem;
            };
            addCell(actionMetric.casts.toFixed(1)); // Casts
            addCell(actionMetric.castsPerMinute.toFixed(1)); // CPM
        });
        $(this.tableElem).trigger('update');
    }
}
