import { IndividualSimResult } from '../api/api.js';
import { StatWeightsRequest, StatWeightsResult } from '../api/api.js';
import { Stat } from '../api/common.js';
import { statNames } from '../api/names.js';
import { stDevToConf90 } from '../utils.js';
import { Component } from './component.js';

export class Results extends Component {
  readonly pendingElem: HTMLDivElement;
  readonly simElem: HTMLDivElement;
  readonly epElem: HTMLDivElement;

  constructor(parent: HTMLElement) {
    super(parent, 'results-root');

    this.rootElem.innerHTML = `
      <div class="results-pending">
        <div class="loader"></div>
      </div>
      <div class="results-sim">
      </div>
      <div class="results-ep">
      </div>
    `;

    this.pendingElem = this.rootElem.getElementsByClassName('results-pending')[0] as HTMLDivElement;
    this.simElem = this.rootElem.getElementsByClassName('results-sim')[0] as HTMLDivElement;
    this.epElem = this.rootElem.getElementsByClassName('results-ep')[0] as HTMLDivElement;
    this.hideAll();
  }

  hideAll() {
    this.pendingElem.style.display = 'none';
    this.simElem.style.display = 'none';
    this.epElem.style.display = 'none';
  }

  setPending() {
    this.hideAll();
    this.pendingElem.style.display = 'initial';
  }

  setSimResult(result: IndividualSimResult) {
    this.hideAll();
    this.simElem.style.display = 'initial';
    this.simElem.innerHTML = `
      <span class="results-sim-dps-avg">${result.dpsAvg.toFixed(2)}</span>
      <span class="results-sim-dps-stdev">${result.dpsStdev.toFixed(2)}</span>
    `;
  }

  setStatWeights(request: StatWeightsRequest, result: StatWeightsResult, epStats: Array<Stat>) {
		const iterations = request.options!.iterations;

    this.hideAll();
    this.epElem.style.display = 'initial';
    this.epElem.innerHTML = '<table class="results-ep-table">'
        + epStats.map(stat => `<tr>
            <td>${statNames[stat]}:</td>
            <td>${result.weights[stat].toFixed(2)}</td>
            <td>${stDevToConf90(result.weightsStdev[stat], iterations).toFixed(2)}</td>
            <td>${result.epValues[stat].toFixed(2)}</td>
            <td>${stDevToConf90(result.epValuesStdev[stat], iterations).toFixed(2)}</td>
            </tr>`)
        .join('')
        + '</table';
  }
}
