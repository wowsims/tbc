import { ResultComponent } from './result_component.js';
export class PercentOom extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'percent-oom';
        super(config);
    }
    onSimResult(resultData) {
        const players = resultData.result.getPlayers(resultData.filter);
        if (players.length == 1) {
            const player = players[0];
            const secondsOOM = player.secondsOomAvg;
            const percentOOM = secondsOOM / resultData.result.encounterMetrics.durationSeconds;
            this.rootElem.innerHTML = `
				<span class="percent-oom-value">${secondsOOM.toFixed(1)}</span>
				<span class="percent-oom-label">seconds spent OOM on average</span>
			`;
            const dangerLevel = percentOOM < 0.01 ? 'safe' : (percentOOM < 0.05 ? 'warning' : 'danger');
            this.rootElem.classList.remove('safe', 'warning', 'danger');
            this.rootElem.classList.add(dangerLevel);
            this.rootElem.style.display = 'initial';
        }
        else {
            this.rootElem.style.display = 'none';
        }
    }
}
