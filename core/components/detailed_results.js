import { REPO_NAME } from '/tbc/core/constants/other.js';
import { Component } from './component.js';
export class DetailedResults extends Component {
    constructor(parent, simUI, simResultsManager) {
        super(parent, 'detailed-results-manager-root');
        this.simUI = simUI;
        this.tabWindow = null;
        this.latestResult = null;
        this.simUI.sim.showThreatMetricsChangeEmitter.on(() => this.updateSettings());
        const computedStyles = window.getComputedStyle(this.rootElem);
        const url = new URL(`${window.location.protocol}//${window.location.host}/${REPO_NAME}/detailed_results/index.html`);
        url.searchParams.append('mainTextColor', computedStyles.getPropertyValue('--main-text-color').trim());
        url.searchParams.append('themeColorPrimary', computedStyles.getPropertyValue('--theme-color-primary').trim());
        url.searchParams.append('themeColorBackground', computedStyles.getPropertyValue('--theme-color-background').trim());
        url.searchParams.append('themeColorBackgroundRaw', computedStyles.getPropertyValue('--theme-color-background-raw').trim());
        url.searchParams.append('themeBackgroundImage', computedStyles.getPropertyValue('--theme-background-image').trim());
        url.searchParams.append('themeBackgroundOpacity', computedStyles.getPropertyValue('--theme-background-opacity').trim());
        if (simUI.isIndividualSim()) {
            url.searchParams.append('isIndividualSim', '');
        }
        this.rootElem.innerHTML = `
		<div class="detailed-results-controls-div">
			<button class="detailed-results-new-tab-button sim-button">VIEW IN SEPARATE TAB</button>
		</div>
		<iframe class="detailed-results-iframe" src="${url.href}" allowtransparency="true"></iframe>
		`;
        this.iframeElem = this.rootElem.getElementsByClassName('detailed-results-iframe')[0];
        const newTabButton = this.rootElem.getElementsByClassName('detailed-results-new-tab-button')[0];
        newTabButton.addEventListener('click', event => {
            if (this.tabWindow == null || this.tabWindow.closed) {
                this.tabWindow = window.open(url.href, 'Detailed Results');
                this.tabWindow.addEventListener('load', event => {
                    if (this.latestResult) {
                        this.updateSettings();
                        this.setSimResult(this.latestResult);
                    }
                });
            }
            else {
                this.tabWindow.focus();
            }
        });
        simResultsManager.currentChangeEmitter.on(() => {
            const cur = simResultsManager.getCurrentData();
            if (cur) {
                this.setSimResult(cur.simResult);
            }
        });
    }
    // TODO: Decide whether to continue using this or just remove it.
    //setPending() {
    //	this.latestResult = null;
    //	this.iframeElem.contentWindow!.postMessage(null, '*');
    //	if (this.tabWindow) {
    //		this.tabWindow.postMessage(null, '*');
    //	}
    //}
    setSimResult(simResult) {
        this.latestResult = simResult;
        this.postMessage(simResult.toJson());
    }
    updateSettings() {
        if (this.simUI.sim.getShowThreatMetrics()) {
            this.postMessage('showThreatMetrics');
        }
        else {
            this.postMessage('hideThreatMetrics');
        }
    }
    postMessage(data) {
        this.iframeElem.contentWindow.postMessage(data, '*');
        if (this.tabWindow) {
            this.tabWindow.postMessage(data, '*');
        }
    }
}
