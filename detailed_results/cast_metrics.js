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
        $(this.tableElem).tablesorter({ sortList: [1, 0] });
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
				<img class="cast-metrics-action-icon" src="${actionMetric.iconUrl}"></img>
				<span class="cast-metrics-action-name">${actionMetric.name}</span>
				`;
                const addCell = (value) => {
                    const cellElem = document.createElement('td');
                    cellElem.textContent = String(value);
                    rowElem.appendChild(cellElem);
                    return cellElem;
                };
                addCell((actionMetric.totalDmg / iterations / duration).toFixed(1)); // DPS
                addCell((actionMetric.casts / iterations).toFixed(1)); // Casts
                addCell((actionMetric.totalDmg / actionMetric.casts).toFixed(1)); // Avg Cast
                addCell(((actionMetric.casts - actionMetric.misses) / iterations).toFixed(1)); // Hits
                addCell((actionMetric.totalDmg / (actionMetric.casts - actionMetric.misses)).toFixed(1)); // Avg Hit
                addCell(((actionMetric.crits / actionMetric.casts) * 100).toFixed(2) + ' %'); // Crit %
                addCell(((actionMetric.misses / actionMetric.casts) * 100).toFixed(2) + ' %'); // Miss %
            });
            $(this.tableElem).trigger('update');
        });
    }
}
