import { TypedEvent } from '/tbc/core/typed_event.js';
import { ResultComponent } from './result_component.js';
export var ColumnSortType;
(function (ColumnSortType) {
    ColumnSortType[ColumnSortType["None"] = 0] = "None";
    ColumnSortType[ColumnSortType["Ascending"] = 1] = "Ascending";
    ColumnSortType[ColumnSortType["Descending"] = 2] = "Descending";
})(ColumnSortType || (ColumnSortType = {}));
;
export class MetricsTable extends ResultComponent {
    constructor(config, columnConfigs) {
        super(config);
        this.onUpdate = new TypedEvent('MetricsTableUpdate');
        this.columnConfigs = columnConfigs;
        this.rootElem.innerHTML = `
		<table class="metrics-table tablesorter">
			<thead class="metrics-table-header">
				<tr class="metrics-table-header-row"></tr>
			</thead>
			<tbody class="metrics-table-body">
			</tbody>
		</table>
		`;
        this.tableElem = this.rootElem.getElementsByClassName('metrics-table')[0];
        this.bodyElem = this.rootElem.getElementsByClassName('metrics-table-body')[0];
        const headerRowElem = this.rootElem.getElementsByClassName('metrics-table-header-row')[0];
        this.columnConfigs.forEach(columnConfig => {
            const headerCell = document.createElement('th');
            headerCell.classList.add('metrics-table-header-cell');
            if (columnConfig.headerCellClass) {
                headerCell.classList.add(columnConfig.headerCellClass);
            }
            headerCell.innerHTML = `<span>${columnConfig.name}</span>`;
            if (columnConfig.tooltip) {
                tippy(headerCell, {
                    'content': columnConfig.tooltip,
                    'allowHTML': true,
                });
            }
            headerRowElem.appendChild(headerCell);
        });
        const sortList = this.columnConfigs
            .map((config, i) => [i, config.sort == ColumnSortType.Ascending ? 0 : 1])
            .filter(sortData => this.columnConfigs[sortData[0]].sort);
        $(this.tableElem).tablesorter({
            sortList: sortList,
            cssChildRow: 'child-metric',
        });
    }
    sortMetrics(metrics) {
        this.columnConfigs.filter(config => config.sort).forEach(config => {
            if (!config.getValue) {
                throw new Error('Can\' apply group sorting without getValue');
            }
            if (config.sort == ColumnSortType.Ascending) {
                metrics.sort((a, b) => config.getValue(a) - config.getValue(b));
            }
            else {
                metrics.sort((a, b) => config.getValue(b) - config.getValue(a));
            }
        });
    }
    addRow(metric) {
        const rowElem = document.createElement('tr');
        this.bodyElem.appendChild(rowElem);
        this.columnConfigs.forEach(columnConfig => {
            const cellElem = document.createElement('td');
            if (columnConfig.fillCell) {
                columnConfig.fillCell(metric, cellElem, rowElem);
            }
            else if (columnConfig.getDisplayString) {
                cellElem.textContent = columnConfig.getDisplayString(metric);
            }
            else {
                throw new Error('Metrics column config does not provide content function: ' + columnConfig.name);
            }
            rowElem.appendChild(cellElem);
        });
        this.customizeRowElem(metric, rowElem);
        return rowElem;
    }
    addGroup(metrics) {
        if (metrics.length == 0) {
            return;
        }
        if (metrics.length == 1 && this.shouldCollapse(metrics[0])) {
            this.addRow(metrics[0]);
            return;
        }
        // Manually sort because tablesorter doesn't let us apply sorting to child rows.
        this.sortMetrics(metrics);
        const mergedMetrics = this.mergeMetrics(metrics);
        const parentRow = this.addRow(mergedMetrics);
        const childRows = metrics.map(metric => this.addRow(metric));
        childRows.forEach(childRow => childRow.classList.add('child-metric'));
        let expand = true;
        parentRow.classList.add('parent-metric', 'expand');
        parentRow.addEventListener('click', event => {
            expand = !expand;
            if (expand) {
                childRows.forEach(row => row.classList.remove('hide'));
                parentRow.classList.add('expand');
            }
            else {
                childRows.forEach(row => row.classList.add('hide'));
                parentRow.classList.remove('expand');
            }
        });
    }
    onSimResult(resultData) {
        this.bodyElem.textContent = '';
        const groupedMetrics = this.getGroupedMetrics(resultData).filter(group => group.length > 0);
        if (groupedMetrics.length == 0) {
            this.rootElem.classList.add('hide');
            this.onUpdate.emit(resultData.eventID);
            return;
        }
        else {
            this.rootElem.classList.remove('hide');
        }
        groupedMetrics.forEach(group => this.addGroup(group));
        $(this.tableElem).trigger('update');
        this.onUpdate.emit(resultData.eventID);
    }
    // Whether a single-element group should have its parent row removed.
    // Override this to add custom behavior.
    shouldCollapse(metric) {
        return true;
    }
    // Override this to customize rowElem after it has been populated.
    customizeRowElem(metric, rowElem) { }
    // Override this to provide custom merge behavior.
    mergeMetrics(metrics) {
        return metrics[0];
    }
    static nameCellConfig(getData) {
        return {
            name: 'Name',
            fillCell: (metric, cellElem, rowElem) => {
                const data = getData(metric);
                cellElem.innerHTML = `
				<a class="metrics-action-icon"></a>
				<span class="metrics-action-name">${data.name}</span>
				<span class="expand-toggle fa fa-caret-down"></span>
				<span class="expand-toggle fa fa-caret-right"></span>
				`;
                const iconElem = cellElem.getElementsByClassName('metrics-action-icon')[0];
                data.actionId.setBackgroundAndHref(iconElem);
            },
        };
    }
    static playerNameCellConfig() {
        return {
            name: 'Name',
            fillCell: (player, cellElem, rowElem) => {
                cellElem.innerHTML = `
				<img class="metrics-action-icon" src="${player.iconUrl}"></img>
				<span class="metrics-action-name" style="color:${player.classColor}">${player.label}</span>
				`;
            },
        };
    }
}
