import { AuraMetrics } from '/tbc/core/proto_utils/sim_result.js';
import { ColumnSortType, MetricsTable } from './metrics_table.js';
export class AuraMetricsTable extends MetricsTable {
    constructor(config, useDebuffs) {
        if (useDebuffs) {
            config.rootCssClass = 'debuff-metrics-root';
        }
        else {
            config.rootCssClass = 'buff-metrics-root';
        }
        super(config, [
            MetricsTable.nameCellConfig((metric) => {
                return {
                    name: metric.name,
                    actionId: metric.actionId,
                };
            }),
            {
                name: 'Uptime',
                tooltip: 'Uptime / Encounter Duration',
                sort: ColumnSortType.Descending,
                getValue: (metric) => metric.uptimePercent,
                getDisplayString: (metric) => metric.uptimePercent.toFixed(2) + '%',
            },
        ]);
        this.useDebuffs = useDebuffs;
    }
    getGroupedMetrics(resultData) {
        if (this.useDebuffs) {
            return AuraMetrics.groupById(resultData.result.getDebuffMetrics(resultData.filter));
        }
        else {
            const players = resultData.result.getPlayers(resultData.filter);
            if (players.length != 1) {
                return [];
            }
            const player = players[0];
            const auras = player.auras;
            const actionGroups = AuraMetrics.groupById(auras);
            const petGroups = player.pets.map(pet => pet.auras);
            return actionGroups.concat(petGroups);
        }
    }
    mergeMetrics(metrics) {
        return AuraMetrics.merge(metrics, true, metrics[0].unit?.petActionId || undefined);
    }
    shouldCollapse(metric) {
        return !metric.unit?.isPet;
    }
}
