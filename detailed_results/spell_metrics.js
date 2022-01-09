import { ActionMetrics } from '/tbc/core/proto_utils/sim_result.js';
import { ResultComponent } from './result_component.js';
export class SpellMetrics extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'cast-metrics-root';
        super(config);
        this.rootElem.innerHTML = `
		<table class="metrics-table tablesorter">
			<thead class="metrics-table-header">
				<tr class="metrics-table-header-row">
					<th class="metrics-table-header-cell"><span>Name</span></th>
					<th class="metrics-table-header-cell"><span>DPS</span></th>
					<th class="metrics-table-header-cell"><span>Casts</span></th>
					<th class="metrics-table-header-cell"><span>Avg Cast</span></th>
					<th class="metrics-table-header-cell"><span>Hits</span></th>
					<th class="metrics-table-header-cell"><span>Avg Hit</span></th>
					<th class="metrics-table-header-cell"><span>Crit %</span></th>
					<th class="metrics-table-header-cell"><span>Miss %</span></th>
				</tr>
			</thead>
			<tbody class="metrics-table-body">
			</tbody>
		</table>
		`;
        this.tableElem = this.rootElem.getElementsByClassName('metrics-table')[0];
        this.bodyElem = this.rootElem.getElementsByClassName('metrics-table-body')[0];
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
    onSimResult(resultData) {
        this.bodyElem.textContent = '';
        const addRow = (spellMetric, isChildMetric) => {
            const rowElem = document.createElement('tr');
            if (isChildMetric) {
                rowElem.classList.add('child-metric');
            }
            this.bodyElem.appendChild(rowElem);
            const nameCellElem = document.createElement('td');
            rowElem.appendChild(nameCellElem);
            nameCellElem.innerHTML = `
			<a class="metrics-action-icon"></a>
			<span class="metrics-action-name">${spellMetric.name}</span>
			<span class="expand-toggle fa fa-caret-down"></span>
			<span class="expand-toggle fa fa-caret-right"></span>
			`;
            const iconElem = nameCellElem.getElementsByClassName('metrics-action-icon')[0];
            spellMetric.actionId.setBackgroundAndHref(iconElem);
            const addCell = (value) => {
                const cellElem = document.createElement('td');
                cellElem.textContent = String(value);
                rowElem.appendChild(cellElem);
                return cellElem;
            };
            addCell(spellMetric.dps.toFixed(1)); // DPS
            addCell(spellMetric.casts.toFixed(1)); // Casts
            addCell(spellMetric.avgCast.toFixed(1)); // Avg Cast
            addCell(spellMetric.hits.toFixed(1)); // Hits
            addCell(spellMetric.avgHit.toFixed(1)); // Avg Hit
            addCell(spellMetric.critPercent.toFixed(2) + ' %'); // Crit %
            addCell(spellMetric.missPercent.toFixed(2) + ' %'); // Miss %
            return rowElem;
        };
        const spellMetrics = resultData.result.getSpellMetrics(resultData.filter);
        const spellGroups = ActionMetrics.groupById(spellMetrics);
        spellGroups.forEach(spellGroup => {
            if (spellGroup.length == 1) {
                addRow(spellGroup[0], false);
                return;
            }
            const mergedMetrics = ActionMetrics.merge(spellGroup, true);
            const parentRow = addRow(mergedMetrics, false);
            const childRows = spellGroup.map(spellMetric => addRow(spellMetric, true));
            const defaultDisplay = childRows[0].style.display;
            let expand = true;
            parentRow.classList.add('parent-metric', 'expand');
            parentRow.addEventListener('click', event => {
                expand = !expand;
                const newDisplayValue = expand ? defaultDisplay : 'none';
                childRows.forEach(row => row.style.display = newDisplayValue);
                if (expand) {
                    parentRow.classList.add('expand');
                }
                else {
                    parentRow.classList.remove('expand');
                }
            });
        });
        $(this.tableElem).trigger('update');
    }
}
