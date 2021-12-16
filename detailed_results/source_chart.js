import { sum } from '/tbc/core/utils.js';
import { ResultComponent } from './result_component.js';
export class SourceChart extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'source-chart-root';
        super(config);
    }
    onSimResult(resultData) {
        const chartBounds = this.rootElem.getBoundingClientRect();
        this.rootElem.textContent = '';
        const chartCanvas = document.createElement("canvas");
        chartCanvas.height = chartBounds.height;
        chartCanvas.width = chartBounds.width;
        const colors = ['red', 'blue', 'lawngreen'];
        const actionMetrics = resultData.result.getActionMetrics(resultData.filter);
        const names = actionMetrics.map(am => am.name);
        const totalDmg = sum(actionMetrics.map(actionMetric => actionMetric.damage));
        const vals = actionMetrics.map(actionMetric => actionMetric.damage / totalDmg);
        const ctx = chartCanvas.getContext('2d');
        const chart = new Chart(ctx, {
            type: 'pie',
            data: {
                labels: names,
                datasets: [{
                        data: vals,
                        backgroundColor: colors,
                    }],
            },
            options: {
                plugins: {
                    legend: {
                        display: true,
                        position: 'right',
                    }
                },
            },
        });
        this.rootElem.appendChild(chartCanvas);
    }
}
