import { IndividualSimResult } from '/tbc/core/proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from '/tbc/core/proto/api.js';
import { Stat } from '/tbc/core/proto/common.js';
import { statNames } from '/tbc/core/proto_utils/names.js';
import { stDevToConf90 } from '/tbc/core/utils.js';
import { Component } from './component.js';

export class Results extends Component {
  readonly pendingElem: HTMLDivElement;
  readonly simElem: HTMLDivElement;
  readonly epElem: HTMLDivElement;
	private statsType: string;

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
		this.statsType = 'ep';
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
    this.epElem.innerHTML = `
		<select class="results-ep-type-select">
			<option value="ep">EP</option>
			<option value="weight">DPS / Stat</option>
		</select>
		<table class="results-ep-table">
		` + epStats.map(stat => `<tr>
					<td>${statNames[stat]}:</td>
					<td>${result.weights[stat].toFixed(2)}</td>
					<td>${stDevToConf90(result.weightsStdev[stat], iterations).toFixed(2)}</td>
					<td>${result.epValues[stat].toFixed(2)}</td>
					<td>${stDevToConf90(result.epValuesStdev[stat], iterations).toFixed(2)}</td>
					</tr>`)
			.join('')
			+ '</table';

		const setType = (type: string) => {
			if (type == 'ep') {
				this.epElem.classList.remove('stats-type-dps');
				this.epElem.classList.add('stats-type-ep');
			} else {
				this.epElem.classList.add('stats-type-dps');
				this.epElem.classList.remove('stats-type-ep');
			}

			this.statsType = type;
		};

		const selectElem = this.epElem.getElementsByClassName('results-ep-type-select')[0] as HTMLSelectElement;
		selectElem.addEventListener('input', event => {
			setType(selectElem.value);
		});
		selectElem.value = this.statsType;
		setType(this.statsType);
  }
}
