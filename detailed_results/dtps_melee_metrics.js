import { ActionMetrics } from '/tbc/core/proto_utils/sim_result.js';
import { ColumnSortType, MetricsTable } from './metrics_table.js';
export class DtpsMeleeMetricsTable extends MetricsTable {
    constructor(config) {
        config.rootCssClass = 'dtps-melee-metrics-root';
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
                name: 'Avg Cast',
                tooltip: 'Damage / Casts',
                getValue: (metric) => metric.avgCast,
                getDisplayString: (metric) => metric.avgCast.toFixed(1),
            },
            {
                name: 'Avg Hit',
                tooltip: 'Damage / (Hits + Crits + Glances + Blocks)',
                getValue: (metric) => metric.avgHit,
                getDisplayString: (metric) => metric.avgHit.toFixed(1),
            },
            {
                name: 'Casts',
                tooltip: 'Casts',
                getValue: (metric) => metric.casts,
                getDisplayString: (metric) => metric.casts.toFixed(1),
            },
            {
                name: 'Hits',
                tooltip: 'Hits + Crits + Glances + Blocks',
                getValue: (metric) => metric.landedHits,
                getDisplayString: (metric) => metric.landedHits.toFixed(1),
            },
            {
                name: 'Miss %',
                tooltip: 'Misses / Swings',
                getValue: (metric) => metric.missPercent,
                getDisplayString: (metric) => metric.missPercent.toFixed(2) + '%',
            },
            {
                name: 'Dodge %',
                tooltip: 'Dodges / Swings',
                getValue: (metric) => metric.dodgePercent,
                getDisplayString: (metric) => metric.dodgePercent.toFixed(2) + '%',
            },
            {
                name: 'Parry %',
                tooltip: 'Parries / Swings',
                getValue: (metric) => metric.parryPercent,
                getDisplayString: (metric) => metric.parryPercent.toFixed(2) + '%',
            },
            {
                name: 'Block %',
                tooltip: 'Blocks / Swings',
                getValue: (metric) => metric.blockPercent,
                getDisplayString: (metric) => metric.blockPercent.toFixed(2) + '%',
            },
            {
                name: 'Crit %',
                tooltip: 'Crits / Swings',
                getValue: (metric) => metric.critPercent,
                getDisplayString: (metric) => metric.critPercent.toFixed(2) + '%',
            },
            {
                name: 'Crush %',
                tooltip: 'Crushes / Swings',
                getValue: (metric) => metric.crushPercent,
                getDisplayString: (metric) => metric.crushPercent.toFixed(2) + '%',
            },
        ]);
    }
    getGroupedMetrics(resultData) {
        const players = resultData.result.getPlayers(resultData.filter);
        if (players.length != 1) {
            return [];
        }
        const player = players[0];
        const targets = resultData.result.getTargets(resultData.filter);
        const targetActions = targets.map(target => target.getMeleeActions().map(action => action.forTarget(player.index)));
        return targetActions;
    }
    mergeMetrics(metrics) {
        // TODO: Use NPC ID here instead of pet ID.
        return ActionMetrics.merge(metrics, true, metrics[0].unit?.petActionId || undefined);
    }
    shouldCollapse(metric) {
        return false;
    }
}
