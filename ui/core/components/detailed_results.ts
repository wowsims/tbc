import { IndividualSimResult } from '../api/api';
import { Component } from './component';

export class DetailedResults extends Component {
	private readonly iframeElem: HTMLIFrameElement;

  constructor(parent: HTMLElement) {
    super(parent, 'detailed-results-manager-root');

		this.rootElem.innerHTML = `
		<iframe class="detailed-results-iframe" src="/detailed_results"></iframe>
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
