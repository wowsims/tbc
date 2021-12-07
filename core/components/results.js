import { Stat } from '/tbc/core/proto/common.js';
import { statNames } from '/tbc/core/proto_utils/names.js';
import { stDevToConf90 } from '/tbc/core/utils.js';
import { Component } from './component.js';
export class Results extends Component {
    constructor(parent, simUI) {
        super(parent, 'results-root');
        this.currentData = null;
        this.referenceData = null;
        this.simUI = simUI;
        this.rootElem.innerHTML = `
      <div class="results-pending">
        <div class="loader"></div>
      </div>
      <div class="results-sim">
				<div class="results-sim-dps">
				</div>
				<div class="results-sim-reference">
					<span class="results-sim-set-reference fa fa-bookmark"></span>
					<div class="results-sim-reference-bar">
						<span class="results-sim-reference-diff"></span>
						<span class="results-sim-reference-text"> vs. reference</span>
						<span class="results-sim-reference-swap fa fa-retweet"></span>
						<span class="results-sim-reference-delete fa fa-times"></span>
					</div>
				</div>
      </div>
      <div class="results-ep">
      </div>
    `;
        this.pendingElem = this.rootElem.getElementsByClassName('results-pending')[0];
        this.simElem = this.rootElem.getElementsByClassName('results-sim')[0];
        this.simDpsElem = this.rootElem.getElementsByClassName('results-sim-dps')[0];
        this.epElem = this.rootElem.getElementsByClassName('results-ep')[0];
        this.statsType = 'ep';
        this.hideAll();
        this.simReferenceElem = this.rootElem.getElementsByClassName('results-sim-reference')[0];
        this.simReferenceDiffElem = this.rootElem.getElementsByClassName('results-sim-reference-diff')[0];
        const simReferenceSetButton = this.rootElem.getElementsByClassName('results-sim-set-reference')[0];
        simReferenceSetButton.addEventListener('click', event => {
            this.referenceData = this.currentData;
            this.updateReference();
        });
        tippy(simReferenceSetButton, {
            'content': 'Use as reference',
            'allowHTML': true,
        });
        const simReferenceSwapButton = this.rootElem.getElementsByClassName('results-sim-reference-swap')[0];
        simReferenceSwapButton.addEventListener('click', event => {
            if (this.currentData && this.referenceData) {
                const tmpData = this.currentData;
                this.currentData = this.referenceData;
                this.referenceData = tmpData;
                this.simUI.fromJson(this.currentData.settings);
                this.setSimResult(this.currentData.request, this.currentData.result);
                this.updateReference();
            }
        });
        tippy(simReferenceSwapButton, {
            'content': 'Swap reference with current',
            'allowHTML': true,
        });
        const simReferenceDeleteButton = this.rootElem.getElementsByClassName('results-sim-reference-delete')[0];
        simReferenceDeleteButton.addEventListener('click', event => {
            this.referenceData = null;
            this.updateReference();
        });
        tippy(simReferenceDeleteButton, {
            'content': 'Remove reference',
            'allowHTML': true,
        });
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
    setSimResult(request, result) {
        this.currentData = {
            request: request,
            result: result,
            settings: this.simUI.toJson(),
        };
        this.hideAll();
        this.simElem.style.display = 'initial';
        const dpsMetrics = result.raidMetrics.dps;
        this.simDpsElem.innerHTML = `
      <span class="results-sim-dps-avg">${dpsMetrics.avg.toFixed(2)}</span>
      <span class="results-sim-dps-stdev">${dpsMetrics.stdev.toFixed(2)}</span>
    `;
        this.updateReference();
    }
    setStatWeights(request, result, epStats) {
        const iterations = request.simOptions.iterations;
        if (request.epReferenceStat == Stat.StatSpellPower) {
            result.epValues.forEach((value, index) => {
                if (index == Stat.StatArcaneSpellPower ||
                    index == Stat.StatFireSpellPower ||
                    index == Stat.StatFrostSpellPower ||
                    index == Stat.StatHolySpellPower ||
                    index == Stat.StatNatureSpellPower ||
                    index == Stat.StatShadowSpellPower) {
                    if (value > result.epValues[request.epReferenceStat]) {
                        const diff = value - result.epValues[request.epReferenceStat];
                        result.epValues[index] = result.epValues[request.epReferenceStat];
                        result.epValuesStdev[index] -= diff;
                        const wdiff = result.weights[index] - result.weights[request.epReferenceStat];
                        result.weights[index] = result.weights[request.epReferenceStat];
                        result.weightsStdev[index] -= wdiff;
                    }
                }
            });
        }
        this.hideAll();
        this.epElem.style.display = 'initial';
        this.epElem.innerHTML = `
		<select class="results-ep-type-select">
			<option value="ep">EP</option>
			<option value="weight">DPS</option>
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
        const setType = (type) => {
            if (type == 'ep') {
                this.epElem.classList.remove('stats-type-dps');
                this.epElem.classList.add('stats-type-ep');
            }
            else {
                this.epElem.classList.add('stats-type-dps');
                this.epElem.classList.remove('stats-type-ep');
            }
            this.statsType = type;
        };
        const selectElem = this.epElem.getElementsByClassName('results-ep-type-select')[0];
        selectElem.addEventListener('input', event => {
            setType(selectElem.value);
        });
        selectElem.value = this.statsType;
        setType(this.statsType);
    }
    updateReference() {
        if (!this.referenceData || !this.currentData) {
            this.simReferenceElem.classList.remove('has-reference');
            return;
        }
        this.simReferenceElem.classList.add('has-reference');
        const currentDpsMetrics = this.currentData.result.raidMetrics.dps;
        const referenceDpsMetrics = this.referenceData.result.raidMetrics.dps;
        const delta = currentDpsMetrics.avg - referenceDpsMetrics.avg;
        const deltaStr = delta.toFixed(2);
        if (delta >= 0) {
            this.simReferenceDiffElem.textContent = '+' + deltaStr;
            this.simReferenceDiffElem.classList.remove('negative');
            this.simReferenceDiffElem.classList.add('positive');
        }
        else {
            this.simReferenceDiffElem.textContent = '' + deltaStr;
            this.simReferenceDiffElem.classList.remove('positive');
            this.simReferenceDiffElem.classList.add('negative');
        }
    }
}
