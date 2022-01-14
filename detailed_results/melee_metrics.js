import { ActionMetrics } from '/tbc/core/proto_utils/sim_result.js';
import { ResultComponent } from './result_component.js';
export class MeleeMetrics extends ResultComponent {
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
					<th class="metrics-table-header-cell"><span>Dodge %</span></th>
					<th class="metrics-table-header-cell"><span>Glance %</span></th>
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
            'content': 'Crits / Swings',
            'allowHTML': true,
        });
        // Miss %
        tippy(headerElems[7], {
            'content': 'Misses / Swings',
            'allowHTML': true,
        });
        // Dodge %
        tippy(headerElems[8], {
            'content': 'Dodges / Swings',
            'allowHTML': true,
        });
        // Glance %
        tippy(headerElems[9], {
            'content': 'Glances / Swings',
            'allowHTML': true,
        });
        $(this.tableElem).tablesorter({ sortList: [[1, 1]] });
    }
    onSimResult(resultData) {
        this.bodyElem.textContent = '';
        const addRow = (meleeMetric, isChildMetric) => {
            const rowElem = document.createElement('tr');
            if (isChildMetric) {
                rowElem.classList.add('child-metric');
            }
            this.bodyElem.appendChild(rowElem);
            const nameCellElem = document.createElement('td');
            rowElem.appendChild(nameCellElem);
            nameCellElem.innerHTML = `
			<a class="metrics-action-icon"></a>
			<span class="metrics-action-name">${meleeMetric.name}</span>
			<span class="expand-toggle fa fa-caret-down"></span>
			<span class="expand-toggle fa fa-caret-right"></span>
			`;
            const iconElem = nameCellElem.getElementsByClassName('metrics-action-icon')[0];
            meleeMetric.actionId.setBackgroundAndHref(iconElem);
            const addCell = (value) => {
                const cellElem = document.createElement('td');
                cellElem.textContent = String(value);
                rowElem.appendChild(cellElem);
                return cellElem;
            };
            addCell(meleeMetric.dps.toFixed(1)); // DPS
            addCell(meleeMetric.casts.toFixed(1)); // Casts
            addCell(meleeMetric.avgCast.toFixed(1)); // Avg Cast
            addCell(meleeMetric.hits.toFixed(1)); // Hits
            addCell(meleeMetric.avgHit.toFixed(1)); // Avg Hit
            addCell(meleeMetric.critPercent.toFixed(2) + ' %'); // Crit %
            addCell(meleeMetric.missPercent.toFixed(2) + ' %'); // Miss %
            addCell(meleeMetric.dodgePercent.toFixed(2) + ' %'); // Dodge %
            addCell(meleeMetric.glancePercent.toFixed(2) + ' %'); // Glance %
            return rowElem;
        };
        const meleeMetrics = resultData.result.getMeleeMetrics(resultData.filter);
        const meleeGroups = ActionMetrics.groupById(meleeMetrics);
        if (meleeMetrics.length == 0) {
            this.rootElem.classList.add('empty');
        }
        else {
            this.rootElem.classList.remove('empty');
        }
        meleeGroups.forEach(meleeGroup => {
            if (meleeGroup.length == 1) {
                addRow(meleeGroup[0], false);
                return;
            }
            const mergedMetrics = ActionMetrics.merge(meleeGroup, true);
            const parentRow = addRow(mergedMetrics, false);
            const childRows = meleeGroup.map(meleeMetric => addRow(meleeMetric, true));
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
