import { ResultComponent } from './result_component.js';
export class Timeline extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'timeline-root';
        super(config);
        this.rootElem.innerHTML = `
		<div class="timeline-disclaimer">
			<span class="timeline-warning fa fa-exclamation-triangle"></span>
			<span class="timeline-warning-description">Timeline only includes data from 1 sim iteration.</span>
			<div class="timeline-run-again-button sim-button">SIM 1 ITERATION</span>
		</div>
		<div class="timeline-plot">
		</div>
		`;
        const runAgainButton = this.rootElem.getElementsByClassName('timeline-run-again-button')[0];
        runAgainButton.addEventListener('click', event => {
            (window.opener || window.parent).postMessage('runOnce', '*');
        });
        this.plotElem = this.rootElem.getElementsByClassName('timeline-plot')[0];
        this.plot = new ApexCharts(this.plotElem, {
            chart: {
                type: 'line',
                foreColor: 'white',
                animations: {
                    enabled: false,
                },
            },
            series: [],
            xaxis: {
                title: {
                    text: 'Time (s)',
                },
            },
            yaxis: {},
            noData: {
                text: 'Waiting for data...',
            },
        });
        this.plot.render();
    }
    onSimResult(resultData) {
        const players = resultData.result.getPlayers(resultData.filter);
        if (players.length != 1) {
            this.plotElem.textContent = '';
            return;
        }
        const player = players[0];
        const duration = resultData.result.request.encounter.duration || 1;
        let logsToShow = player.manaChangedLogs;
        if (logsToShow.length == 0) {
            return;
        }
        const maxMana = logsToShow[0].manaBefore;
        // Remove events that happen at the same time.
        let curTime = -1;
        logsToShow = logsToShow.filter(log => {
            if (log.timestamp == curTime) {
                return false;
            }
            curTime = log.timestamp;
            return true;
        });
        // Reduce to ~100 logs.
        const desiredNumLogs = 100;
        if (logsToShow.length / desiredNumLogs >= 2) {
            const reductionFactor = Math.floor(logsToShow.length / desiredNumLogs);
            logsToShow = logsToShow.filter((log, i) => i % reductionFactor == 0);
        }
        const data = logsToShow.map(log => log.manaAfter);
        this.plot.updateOptions({
            series: [
                {
                    name: 'Mana',
                    data: data,
                },
            ],
            xaxis: {
                min: 0,
                max: duration,
                tickAmount: 10,
                categories: logsToShow.map(log => log.timestamp),
                labels: {
                    show: true,
                    formatter: (val) => val,
                },
            },
            yaxis: {
                min: 0,
                max: maxMana,
                tickAmount: 10,
                title: {
                    text: 'Mana',
                },
                labels: {
                    formatter: (val) => {
                        const v = parseFloat(val);
                        return `${v.toFixed(0)} (${(v / maxMana * 100).toFixed(0)}%)`;
                    },
                },
            },
            dataLabels: {
                formatter: (val) => {
                    const v = parseFloat(val);
                    return `${v.toFixed(0)} (${(v / maxMana * 100).toFixed(0)}%)`;
                },
            },
        });
        this.plot.zoomX(0, duration);
    }
    // Returns the time intervals to use for the chart.
    getTimeIntervals(duration) {
        const candidateWindows = [1, 5, 10, 30, 60];
        const candidateWindow = 30;
        const intervals = [];
        let cur = 0;
        while (cur < duration) {
            intervals.push(cur);
            cur += candidateWindow;
        }
        intervals.push(duration);
        return intervals;
    }
    render() {
        setTimeout(() => this.plot.render(), 300);
    }
}
