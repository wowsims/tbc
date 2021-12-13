export function addRaidSimAction(simUI) {
    simUI.addAction('DPS', 'dps-action', async () => {
        simUI.setResultsPending();
        try {
            const result = await simUI.sim.runRaidSim();
        }
        catch (e) {
            simUI.hideAllResults();
            alert(e);
        }
    });
    const resultsManager = new RaidSimResultsManager(simUI);
    simUI.sim.raidSimEmitter.on(data => {
        resultsManager.setSimResult(data.request, data.result);
    });
}
class RaidSimResultsManager {
    constructor(simUI) {
        this.currentData = null;
        this.referenceData = null;
        this.simUI = simUI;
    }
    setSimResult(request, result) {
        this.currentData = {
            request: request,
            result: result,
            settings: this.simUI.sim.toJson(),
        };
        const dpsMetrics = result.raidMetrics.dps;
        this.simUI.setResultsContent(`
      <div class="results-sim">
				<div class="results-sim-dps">
					<span class="results-sim-dps-avg">${dpsMetrics.avg.toFixed(2)}</span>
					<span class="results-sim-dps-stdev">${dpsMetrics.stdev.toFixed(2)}</span>
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
    `);
        const simReferenceElem = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference')[0];
        const simReferenceDiffElem = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference-diff')[0];
        const simReferenceSetButton = this.simUI.resultsContentElem.getElementsByClassName('results-sim-set-reference')[0];
        simReferenceSetButton.addEventListener('click', event => {
            this.referenceData = this.currentData;
            this.updateReference();
        });
        tippy(simReferenceSetButton, {
            'content': 'Use as reference',
            'allowHTML': true,
        });
        const simReferenceSwapButton = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference-swap')[0];
        simReferenceSwapButton.addEventListener('click', event => {
            if (this.currentData && this.referenceData) {
                const tmpData = this.currentData;
                this.currentData = this.referenceData;
                this.referenceData = tmpData;
                this.simUI.sim.fromJson(this.currentData.settings);
                this.setSimResult(this.currentData.request, this.currentData.result);
                this.updateReference();
            }
        });
        tippy(simReferenceSwapButton, {
            'content': 'Swap reference with current',
            'allowHTML': true,
        });
        const simReferenceDeleteButton = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference-delete')[0];
        simReferenceDeleteButton.addEventListener('click', event => {
            this.referenceData = null;
            this.updateReference();
        });
        tippy(simReferenceDeleteButton, {
            'content': 'Remove reference',
            'allowHTML': true,
        });
        this.updateReference();
    }
    updateReference() {
        const simReferenceElem = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference')[0];
        const simReferenceDiffElem = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference-diff')[0];
        if (!this.referenceData || !this.currentData) {
            simReferenceElem.classList.remove('has-reference');
            return;
        }
        simReferenceElem.classList.add('has-reference');
        const currentDpsMetrics = this.currentData.result.raidMetrics.dps;
        const referenceDpsMetrics = this.referenceData.result.raidMetrics.dps;
        const delta = currentDpsMetrics.avg - referenceDpsMetrics.avg;
        const deltaStr = delta.toFixed(2);
        if (delta >= 0) {
            simReferenceDiffElem.textContent = '+' + deltaStr;
            simReferenceDiffElem.classList.remove('negative');
            simReferenceDiffElem.classList.add('positive');
        }
        else {
            simReferenceDiffElem.textContent = '' + deltaStr;
            simReferenceDiffElem.classList.remove('positive');
            simReferenceDiffElem.classList.add('negative');
        }
    }
}
