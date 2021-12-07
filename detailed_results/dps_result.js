import { ResultComponent } from './result_component.js';
export class DpsResult extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'dps-result';
        super(config);
    }
    onSimResult(request, result) {
        this.rootElem.innerHTML = `
      <span class="results-sim-dps-avg">${result.raidMetrics.dps.avg.toFixed(2)}</span>
      <span class="results-sim-dps-stdev">${result.raidMetrics.dps.stdev.toFixed(2)}</span>
    `;
    }
}
