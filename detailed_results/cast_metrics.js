import { setWowheadHref } from '/tbc/core/resources.js';
import { parseActionMetrics } from './metrics_helpers.js';
import { ResultComponent } from './result_component.js';
export class CastMetrics extends ResultComponent {
    constructor(config) {
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
        this.tableElem = this.rootElem.getElementsByClassName('cast-metrics-table')[0];
        this.bodyElem = this.rootElem.getElementsByClassName('cast-metrics-table-body')[0];
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
    onSimResult(request, result) {
        this.bodyElem.textContent = '';
        const iterations = request.iterations;
        const duration = request.encounter?.duration || 1;
        parseActionMetrics(result.actionMetrics).then(actionMetrics => {
            actionMetrics.forEach(actionMetric => {
                const rowElem = document.createElement('tr');
                this.bodyElem.appendChild(rowElem);
                const nameCellElem = document.createElement('td');
                rowElem.appendChild(nameCellElem);
                nameCellElem.innerHTML = `
				<a class="cast-metrics-action-icon"></a>
				<span class="cast-metrics-action-name">${actionMetric.name}</span>
				`;
                const iconElem = nameCellElem.getElementsByClassName('cast-metrics-action-icon')[0];
                iconElem.style.backgroundImage = `url('${actionMetric.iconUrl}')`;
                if (!('otherId' in actionMetric.actionId)) {
                    setWowheadHref(iconElem, actionMetric.actionId);
                }
                const addCell = (value) => {
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
