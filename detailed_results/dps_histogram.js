import { ResultComponent } from './result_component.js';
export class DpsHistogram extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'dps-histogram-root';
        super(config);
    }
    onSimResult(request, result) {
        const chartBounds = this.rootElem.getBoundingClientRect();
        this.rootElem.textContent = '';
        const chartCanvas = document.createElement("canvas");
        chartCanvas.height = chartBounds.height;
        chartCanvas.width = chartBounds.width;
        const min = result.dpsAvg - result.dpsStdev;
        const max = result.dpsAvg + result.dpsStdev;
        const vals = [];
        const colors = [];
        const labels = Object.keys(result.dpsHist);
        labels.forEach((k, i) => {
            vals.push(result.dpsHist[Number(k)]);
            const val = parseInt(k);
            if (val > min && val < max) {
                colors.push('#1E87F0');
            }
            else {
                colors.push('#FF6961');
            }
        });
        const ctx = chartCanvas.getContext('2d');
        const chart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: labels,
                datasets: [{
                        data: vals,
                        backgroundColor: colors,
                    }],
            },
            options: {
                plugins: {
                    title: {
                        display: true,
                        text: 'DPS Histogram',
                    },
                    legend: {
                        display: false,
                        labels: {},
                    }
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        ticks: {
                            display: false
                        },
                    },
                },
            },
        });
        this.rootElem.appendChild(chartCanvas);
    }
}
