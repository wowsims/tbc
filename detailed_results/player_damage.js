import { maxIndex } from '/tbc/core/utils.js';
import { ColumnSortType, MetricsTable } from './metrics_table.js';
import { SourceChart } from './source_chart.js';
export class PlayerDamageMetricsTable extends MetricsTable {
    constructor(config, resultsFilter) {
        config.rootCssClass = 'player-damage-metrics-root';
        super(config, [
            MetricsTable.playerNameCellConfig(),
            {
                name: 'Amount',
                tooltip: 'Player Damage / Raid Damage',
                headerCellClass: 'amount-header-cell',
                fillCell: (player, cellElem, rowElem) => {
                    cellElem.classList.add('amount-cell');
                    let chart = null;
                    const makeChart = () => {
                        const chartContainer = document.createElement('div');
                        rowElem.appendChild(chartContainer);
                        const sourceChart = new SourceChart(chartContainer, player.actions);
                        return chartContainer;
                    };
                    tippy(rowElem, {
                        content: 'Loading...',
                        placement: 'bottom',
                        onShow(instance) {
                            if (!chart) {
                                chart = makeChart();
                                instance.setContent(chart);
                            }
                        },
                    });
                    cellElem.innerHTML = `
						<div class="player-damage-percent">
							<span>${(player.dps.avg / this.raidDps * 100).toFixed(2)}%</span>
						</div>
						<div class="player-damage-bar-container">
							<div class="player-damage-bar" style="background-color:${player.classColor}; width:${player.dps.avg / this.maxDps * 100}%"></div>
						</div>
						<div class="player-damage-total">
							<span>${(player.totalDamage / 1000).toFixed(1)}k</span>
						</div>
					`;
                },
            },
            {
                name: 'DPS',
                tooltip: 'Damage / Encounter Duration',
                sort: ColumnSortType.Descending,
                getValue: (metric) => metric.dps.avg,
                getDisplayString: (metric) => metric.dps.avg.toFixed(1),
            },
        ]);
        this.resultsFilter = resultsFilter;
        this.raidDps = 0;
        this.maxDps = 0;
    }
    customizeRowElem(player, rowElem) {
        rowElem.classList.add('player-damage-row');
        rowElem.addEventListener('click', event => {
            this.resultsFilter.setPlayer(this.getLastSimResult().eventID, player.index);
        });
    }
    getGroupedMetrics(resultData) {
        const players = resultData.result.getPlayers(resultData.filter);
        this.raidDps = resultData.result.raidMetrics.dps.avg;
        const maxDpsIndex = maxIndex(players.map(player => player.dps.avg));
        this.maxDps = players[maxDpsIndex].dps.avg;
        return players.map(player => [player]);
    }
}
