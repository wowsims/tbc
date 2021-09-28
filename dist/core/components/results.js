import { statNames } from '../api/names.js';
import { Component } from './component.js';
export class Results extends Component {
    constructor(parent) {
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
        this.pendingElem = this.rootElem.getElementsByClassName('results-pending')[0];
        this.simElem = this.rootElem.getElementsByClassName('results-sim')[0];
        this.epElem = this.rootElem.getElementsByClassName('results-ep')[0];
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
    setSimResult(result) {
        this.hideAll();
        this.simElem.style.display = 'initial';
        this.simElem.innerHTML = `
      <span class="results-sim-dps-avg">${result.dpsAvg.toFixed(2)}</span>
      <span class="results-sim-dps-stdev">${result.dpsStdev.toFixed(2)}</span>
    `;
    }
    setStatWeights(result, epStats) {
        this.hideAll();
        this.epElem.style.display = 'initial';
        this.epElem.innerHTML = '<table class="results-ep-table">'
            + epStats.map(stat => `<tr>
            <td>${statNames[stat]}:</td>
            <td>${result.weights[stat].toFixed(2)}</td>
            <td>${result.weightsStdev[stat].toFixed(2)}</td>
            <td>${result.epValues[stat].toFixed(2)}</td>
            <td>${result.epValuesStdev[stat].toFixed(2)}</td>
            </tr>`)
                .join('')
            + '</table';
    }
}
