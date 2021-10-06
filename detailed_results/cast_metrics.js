import { sum } from '/tbc/core/utils.js';
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
        const totalDmg = sum(Object.values(result.casts).map(castMetrics => castMetrics.dmgs[0]));
        for (const [id, castMetrics] of Object.entries(result.casts)) {
            const rowElem = document.createElement('tr');
            this.bodyElem.appendChild(rowElem);
            const addCell = (value) => {
                const cellElem = document.createElement('td');
                cellElem.textContent = String(value);
                rowElem.appendChild(cellElem);
            };
            addCell(id); // Name
            addCell((castMetrics.dmgs[0] / iterations / duration).toFixed(1)); // DPS
            addCell((castMetrics.casts[0] / iterations).toFixed(1)); // Casts
            addCell((castMetrics.dmgs[0] / castMetrics.casts[0]).toFixed(1)); // Avg Cast
            addCell(((castMetrics.casts[0] - castMetrics.misses[0]) / iterations).toFixed(1)); // Hits
            addCell((castMetrics.dmgs[0] / (castMetrics.casts[0] - castMetrics.misses[0])).toFixed(1)); // Avg Hit
            addCell(((castMetrics.crits[0] / castMetrics.casts[0]) * 100).toFixed(2) + ' %'); // Crit %
            addCell(((castMetrics.misses[0] / castMetrics.casts[0]) * 100).toFixed(2) + ' %'); // Miss %
        }
        $(this.tableElem).trigger('update');
    }
}
