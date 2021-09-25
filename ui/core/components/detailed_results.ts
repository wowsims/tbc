import { IndividualSimResult } from '../api/api';
import { Component } from './component';

export class DetailedResults extends Component {
	private readonly iframeElem: HTMLIFrameElement;

  constructor(parent: HTMLElement) {
    super(parent, 'detailed-results-manager-root');

		const computedStyles = window.getComputedStyle(document.body);

		const url = new URL(`https://${window.location.host}/detailed_result`);
		url.searchParams.append('mainBgColor', computedStyles.getPropertyValue('--main-bg-color').trim());
		url.searchParams.append('mainTextColor', computedStyles.getPropertyValue('--main-text-color').trim());

		const cssFilePath = '/elemental_shaman/index.css';
		this.rootElem.innerHTML = `
		<iframe class="detailed-results-iframe" src="/detailed_results?${url.href}"></iframe>
		`;

		this.iframeElem = this.rootElem.getElementsByClassName('detailed-results-iframe')[0] as HTMLIFrameElement;
	}

	setPending() {
		this.iframeElem.contentWindow!.postMessage(null, '*');
	}

  setSimResult(result: IndividualSimResult) {
		this.iframeElem.contentWindow!.postMessage(result, '*');
  }
}
