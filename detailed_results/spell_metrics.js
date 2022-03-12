import { ActionMetrics } from '/tbc/core/proto_utils/sim_result.js';
import { ColumnSortType, MetricsTable } from './metrics_table.js';
export class SpellMetricsTable extends MetricsTable {
    constructor(config) {
        config.rootCssClass = 'spell-metrics-root';
        super(config, [
            MetricsTable.nameCellConfig((metric) => {
                return {
                    name: metric.name,
                    actionId: metric.actionId,
                };
            }),
            {
                name: 'DPS',
                tooltip: 'Damage / Encounter Duration',
                sort: ColumnSortType.Descending,
                getValue: (metric) => metric.dps,
                getDisplayString: (metric) => metric.dps.toFixed(1),
            },
            {
                name: 'Casts',
                tooltip: 'Casts',
                getValue: (metric) => metric.casts,
                getDisplayString: (metric) => metric.casts.toFixed(1),
            },
            {
                name: 'Avg Cast',
                tooltip: 'Damage / Casts',
                getValue: (metric) => metric.avgCast,
                getDisplayString: (metric) => metric.avgCast.toFixed(1),
            },
            {
                name: 'Hits',
                tooltip: 'Hits',
                getValue: (metric) => metric.hits,
                getDisplayString: (metric) => metric.hits.toFixed(1),
            },
            {
                name: 'Avg Hit',
                tooltip: 'Damage / Hits',
                getValue: (metric) => metric.avgHit,
                getDisplayString: (metric) => metric.avgHit.toFixed(1),
            },
            {
                name: 'Crit %',
                tooltip: 'Crits / Hits',
                getValue: (metric) => metric.critPercent,
                getDisplayString: (metric) => metric.critPercent.toFixed(2) + '%',
            },
            {
                name: 'Miss %',
                tooltip: 'Misses / (Hits + Misses)',
                getValue: (metric) => metric.missPercent,
                getDisplayString: (metric) => metric.missPercent.toFixed(2) + '%',
            },
        ]);
    }
    getGroupedMetrics(resultData) {
        const players = resultData.result.getPlayers(resultData.filter);
        if (players.length != 1) {
            return [];
        }
        const player = players[0];
        const actions = player.getSpellActions();
        const actionGroups = ActionMetrics.groupById(actions);
        const petGroups = player.pets.map(pet => pet.getSpellActions());
        return actionGroups.concat(petGroups);
    }
    mergeMetrics(metrics) {
        return ActionMetrics.merge(metrics, true, metrics[0].player?.petActionId || undefined);
    }
    shouldCollapse(metric) {
        return !metric.player?.isPet;
    }
}
