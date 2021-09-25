import { IndividualSimResult } from '../api/api';
import { Component } from './component';

export class DetailedResults extends Component {
	private readonly iframeElem: HTMLIFrameElement;
	private tabWindow: Window | null;
	private latestResult: IndividualSimResult | null;

  constructor(parent: HTMLElement) {
    super(parent, 'detailed-results-manager-root');
		this.tabWindow = null;
		this.latestResult = null;

		const computedStyles = window.getComputedStyle(document.body);

		const url = new URL(`${window.location.protocol}//${window.location.host}/detailed_results/index.html`);
		url.searchParams.append('mainBgColor', computedStyles.getPropertyValue('--main-bg-color').trim());
		url.searchParams.append('mainTextColor', computedStyles.getPropertyValue('--main-text-color').trim());

		const cssFilePath = '/elemental_shaman/index.css';
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
	}

	setPending() {
		this.latestResult = null;
		this.iframeElem.contentWindow!.postMessage(null, '*');
		if (this.tabWindow) {
			this.tabWindow.postMessage(null, '*');
		}
	}

  setSimResult(result: IndividualSimResult) {
		this.latestResult = result;
		this.iframeElem.contentWindow!.postMessage(result, '*');
		if (this.tabWindow) {
			this.tabWindow.postMessage(result, '*');
		}
  }
}
