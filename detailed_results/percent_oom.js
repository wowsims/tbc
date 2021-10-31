import { ResultComponent } from './result_component.js';
export class PercentOom extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'percent-oom';
        super(config);
    }
    onSimResult(request, result) {
        const percentOom = result.numOom / request.simOptions.iterations;
        this.rootElem.innerHTML = `
      <span class="percent-oom-value">${Math.round(percentOom * 100)}%</span>
      <span class="percent-oom-label">of simulations went OOM</span>
    `;
        const dangerLevel = percentOom < 0.05 ? 'safe' : (percentOom < 0.25 ? 'warning' : 'danger');
        this.rootElem.classList.remove('safe', 'warning', 'danger');
        this.rootElem.classList.add(dangerLevel);
    }
}
