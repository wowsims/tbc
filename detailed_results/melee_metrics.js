import { AbilityMetrics } from './ability_metrics.js';
export class MeleeMetrics extends AbilityMetrics {
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
        const tableElem = this.rootElem.getElementsByClassName('metrics-table')[0];
        const headerElems = Array.from(tableElem.querySelectorAll('th'));
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
            'content': 'Hits + Crits + Glances + Blocks',
            'allowHTML': true,
        });
        // Avg Hit
        tippy(headerElems[5], {
            'content': 'Damage / (Hits + Crits + Glances + Blocks)',
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
        $(tableElem).tablesorter({ sortList: [[1, 1]] });
    }
    getPlayerActions(player) {
        return player.getMeleeActions();
    }
    addRowCells(action, addCell) {
        addCell(action.dps.toFixed(1)); // DPS
        addCell(action.casts.toFixed(1)); // Casts
        addCell(action.avgCast.toFixed(1)); // Avg Cast
        addCell(action.landedHits.toFixed(1)); // Hits
        addCell(action.avgHit.toFixed(1)); // Avg Hit
        addCell(action.critPercent.toFixed(2) + ' %'); // Crit %
        addCell(action.missPercent.toFixed(2) + ' %'); // Miss %
        addCell(action.dodgePercent.toFixed(2) + ' %'); // Dodge %
        addCell(action.glancePercent.toFixed(2) + ' %'); // Glance %
    }
}
