import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { SimResult } from '/tbc/core/proto_utils/sim_result.js';
import { repoName } from '/tbc/core/resources.js';
import { SimUI } from '/tbc/core/sim_ui.js';

import { Component } from './component.js';

export class DetailedResults extends Component {
	private readonly iframeElem: HTMLIFrameElement;
	private tabWindow: Window | null;
	private latestResult: SimResult | null;

  constructor(parent: HTMLElement, simUI: SimUI) {
    super(parent, 'detailed-results-manager-root');
		this.tabWindow = null;
		this.latestResult = null;

		const computedStyles = window.getComputedStyle(this.rootElem);

		const url = new URL(`${window.location.protocol}//${window.location.host}/${repoName}/detailed_results/index.html`);
		url.searchParams.append('mainBgColor', computedStyles.getPropertyValue('--main-bg-color').trim());
		url.searchParams.append('mainTextColor', computedStyles.getPropertyValue('--main-text-color').trim());

		this.rootElem.innerHTML = `
		<div class="detailed-results-controls-div">
			<button class="detailed-results-new-tab-button">View in separate tab</button>
		</div>
		<iframe class="detailed-results-iframe" src="${url.href}"></iframe>
		`;

		this.iframeElem = this.rootElem.getElementsByClassName('detailed-results-iframe')[0] as HTMLIFrameElement;

		const newTabButton = this.rootElem.getElementsByClassName('detailed-results-new-tab-button')[0] as HTMLButtonElement;
		newTabButton.addEventListener('click', event => {
			if (this.tabWindow == null || this.tabWindow.closed) {
				this.tabWindow = window.open(url.href, 'Detailed Results');
				this.tabWindow!.addEventListener('load', event => {
					if (this.latestResult) {
						this.setSimResult(this.latestResult);
					}
				});
			} else {
				this.tabWindow.focus();
			}
		});

		simUI.sim.simResultEmitter.on(simResult => {
			this.setSimResult(simResult);
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

  private setSimResult(simResult: SimResult) {
		this.latestResult = simResult;
		const serialized = simResult.toJson();
		this.iframeElem.contentWindow!.postMessage(serialized, '*');
		if (this.tabWindow) {
			this.tabWindow.postMessage(serialized, '*');
		}
  }
}
