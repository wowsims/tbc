import { Component } from '/tbc/core/components/component.js';
import { sum } from '/tbc/core/utils.js';
const sliceColors = [
    '#05878a',
    '#074e67',
    '#5a175d',
    '#67074e',
    '#dd9933',
    '#c9c1e7',
    '#bdd5ef',
    '#c7e3d0',
    '#e7e6ce',
    '#f2d8cc',
    '#e9ccce',
];
export class SourceChart extends Component {
    constructor(parentElem, allActionMetrics) {
        const chartCanvas = document.createElement("canvas");
        super(parentElem, 'source-chart-root', chartCanvas);
        chartCanvas.style.height = '400px';
        chartCanvas.style.width = '600px';
        chartCanvas.height = 400;
        chartCanvas.width = 600;
        const actionMetrics = allActionMetrics.filter(actionMetric => actionMetric.damage > 0);
        const names = actionMetrics.map(am => am.name);
        const totalDmg = sum(actionMetrics.map(actionMetric => actionMetric.damage));
        const vals = actionMetrics.map(actionMetric => actionMetric.damage / totalDmg);
        const bgColors = sliceColors.slice(0, actionMetrics.length);
        const ctx = chartCanvas.getContext('2d');
        const chart = new Chart(ctx, {
            type: 'pie',
            data: {
                labels: names,
                datasets: [{
                        data: vals,
                        backgroundColor: bgColors,
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
    }
}
