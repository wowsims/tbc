import { ActionMetrics } from '/tbc/core/proto_utils/sim_result.js';
import { ColumnSortType, MetricsTable } from './metrics_table.js';
export class MeleeMetricsTable extends MetricsTable {
    constructor(config) {
        config.rootCssClass = 'melee-metrics-root';
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
                name: 'TPS',
                tooltip: 'Threat / Encounter Duration',
                columnClass: 'threat-metrics',
                getValue: (metric) => metric.tps,
                getDisplayString: (metric) => metric.tps.toFixed(1),
            },
            {
                name: 'Avg Cast',
                tooltip: 'Threat / Casts',
                columnClass: 'threat-metrics',
                getValue: (metric) => metric.avgCastThreat,
                getDisplayString: (metric) => metric.avgCastThreat.toFixed(1),
            },
            {
                name: 'Avg Hit',
                tooltip: 'Threat / (Hits + Crits + Glances + Blocks)',
                columnClass: 'threat-metrics',
                getValue: (metric) => metric.avgHitThreat,
                getDisplayString: (metric) => metric.avgHitThreat.toFixed(1),
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
                name: 'Crit %',
                tooltip: 'Crits / Swings',
                getValue: (metric) => metric.critPercent,
                getDisplayString: (metric) => metric.critPercent.toFixed(2) + '%',
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
                name: 'Glance %',
                tooltip: 'Glances / Swings',
                getValue: (metric) => metric.glancePercent,
                getDisplayString: (metric) => metric.glancePercent.toFixed(2) + '%',
            },
        ]);
    }
    getGroupedMetrics(resultData) {
        const players = resultData.result.getPlayers(resultData.filter);
        if (players.length != 1) {
            return [];
        }
        const player = players[0];
        const actions = player.getMeleeActions();
        const actionGroups = ActionMetrics.groupById(actions);
        const petGroups = player.pets.map(pet => pet.getMeleeActions());
        return actionGroups.concat(petGroups);
    }
    mergeMetrics(metrics) {
        return ActionMetrics.merge(metrics, true, metrics[0].player?.petActionId || undefined);
    }
    shouldCollapse(metric) {
        return !metric.player?.isPet;
    }
}
