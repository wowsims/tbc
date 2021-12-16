import { ResultComponent } from './result_component.js';
export class DpsResult extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'dps-result';
        super(config);
    }
    onSimResult(resultData) {
        const damageMetrics = resultData.result.getDamageMetrics(resultData.filter);
        this.rootElem.innerHTML = `
      <span class="results-sim-dps-avg">${damageMetrics.avg.toFixed(2)}</span>
      <span class="results-sim-dps-stdev">${damageMetrics.stdev.toFixed(2)}</span>
    `;
    }
}
