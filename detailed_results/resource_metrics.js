import { ResourceMetrics } from '/tbc/core/proto_utils/sim_result.js';
import { ResourceType } from '/tbc/core/proto/api.js';
import { resourceNames } from '/tbc/core/proto_utils/names.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { ColumnSortType, MetricsTable } from './metrics_table.js';
import { ResultComponent } from './result_component.js';
export class ResourceMetricsTable extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'resource-metrics-root';
        super(config);
        const resourceTypes = getEnumValues(ResourceType).filter(val => val != ResourceType.ResourceTypeNone);
        resourceTypes.forEach(resourceType => {
            const containerElem = document.createElement('div');
            containerElem.classList.add('resource-metrics-table-container', 'hide');
            containerElem.innerHTML = `<span class="resource-metrics-table-title">${resourceNames[resourceType]}</span>`;
            this.rootElem.appendChild(containerElem);
            const childConfig = config;
            childConfig.parent = containerElem;
            const table = new TypedResourceMetricsTable(childConfig, resourceType);
            table.onUpdate.on(() => {
                if (table.rootElem.classList.contains('hide')) {
                    containerElem.classList.add('hide');
                }
                else {
                    containerElem.classList.remove('hide');
                }
            });
        });
    }
    onSimResult(resultData) {
    }
}
export class TypedResourceMetricsTable extends MetricsTable {
    constructor(config, resourceType) {
        config.rootCssClass = 'resource-metrics-table-root';
        super(config, [
            MetricsTable.nameCellConfig((metric) => {
                return {
                    name: metric.name,
                    actionId: metric.actionId,
                };
            }),
            {
                name: 'Casts',
                tooltip: 'Casts',
                getValue: (metric) => metric.events,
                getDisplayString: (metric) => metric.events.toFixed(1),
            },
            {
                name: 'Gain',
                tooltip: 'Gain',
                sort: ColumnSortType.Descending,
                getValue: (metric) => metric.gain,
                getDisplayString: (metric) => metric.gain.toFixed(1),
            },
            {
                name: 'Gain / s',
                tooltip: 'Gain / Second',
                getValue: (metric) => metric.gainPerSecond,
                getDisplayString: (metric) => metric.gainPerSecond.toFixed(1),
            },
            {
                name: 'Avg Gain',
                tooltip: 'Gain / Event',
                getValue: (metric) => metric.avgGain,
                getDisplayString: (metric) => metric.avgGain.toFixed(1),
            },
            {
                name: 'Wasted Gain',
                tooltip: 'Gain that was wasted because of resource cap.',
                getValue: (metric) => metric.wastedGain,
                getDisplayString: (metric) => metric.wastedGain.toFixed(1),
            },
        ]);
        this.resourceType = resourceType;
    }
    getGroupedMetrics(resultData) {
        const players = resultData.result.getPlayers(resultData.filter);
        if (players.length != 1) {
            return [];
        }
        const player = players[0];
        const resources = player.getResourceMetrics(this.resourceType);
        const resourceGroups = ResourceMetrics.groupById(resources);
        return resourceGroups;
    }
    mergeMetrics(metrics) {
        return ResourceMetrics.merge(metrics, true, metrics[0].unit?.petActionId || undefined);
    }
}
