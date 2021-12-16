import { setWowheadHref } from '/tbc/core/resources.js';
import { ResultComponent } from './result_component.js';
export class AuraMetrics extends ResultComponent {
    constructor(config, useDebuffs) {
        if (useDebuffs) {
            config.rootCssClass = 'debuff-aura-metrics-root';
        }
        else {
            config.rootCssClass = 'buff-aura-metrics-root';
        }
        super(config);
        this.useDebuffs = useDebuffs;
        this.rootElem.innerHTML = `
		<table class="metrics-table tablesorter">
			<thead class="metrics-table-header">
				<tr class="metrics-table-header-row">
					<th class="metrics-table-header-cell"><span>Name</span></th>
					<th class="metrics-table-header-cell"><span>Uptime</span></th>
				</tr>
			</thead>
			<tbody class="metrics-table-body">
			</tbody>
		</table>
		`;
        this.tableElem = this.rootElem.getElementsByClassName('metrics-table')[0];
        this.bodyElem = this.rootElem.getElementsByClassName('metrics-table-body')[0];
        const headerElems = Array.from(this.tableElem.querySelectorAll('th'));
        // Uptime
        tippy(headerElems[1], {
            'content': 'Uptime / Encounter Duration',
            'allowHTML': true,
        });
        $(this.tableElem).tablesorter({ sortList: [[1, 1]] });
    }
    onSimResult(resultData) {
        this.bodyElem.textContent = '';
        const auraMetrics = this.useDebuffs
            ? resultData.result.getDebuffMetrics(resultData.filter)
            : resultData.result.getBuffMetrics(resultData.filter);
        auraMetrics.forEach(auraMetric => {
            const rowElem = document.createElement('tr');
            this.bodyElem.appendChild(rowElem);
            const nameCellElem = document.createElement('td');
            rowElem.appendChild(nameCellElem);
            nameCellElem.innerHTML = `
			<a class="metrics-action-icon"></a>
			<span class="metrics-action-name">${auraMetric.name}</span>
			`;
            const iconElem = nameCellElem.getElementsByClassName('metrics-action-icon')[0];
            iconElem.style.backgroundImage = `url('${auraMetric.iconUrl}')`;
            if (!('otherId' in auraMetric.actionId.id)) {
                setWowheadHref(iconElem, auraMetric.actionId.id);
            }
            const addCell = (value) => {
                const cellElem = document.createElement('td');
                cellElem.textContent = String(value);
                rowElem.appendChild(cellElem);
                return cellElem;
            };
            addCell(auraMetric.uptimePercent.toFixed(2) + '%'); // Uptime
        });
        $(this.tableElem).trigger('update');
    }
}
