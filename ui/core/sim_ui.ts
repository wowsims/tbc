import { Component } from '/tbc/core/components/component.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { SimOptions } from '/tbc/core/proto/api.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';

import { Sim } from './sim.js';
import { Target } from './target.js';
import { TypedEvent } from './typed_event.js';

declare var tippy: any;
declare var pako: any;

export interface SimUIConfig {
	title: string,
	knownIssues?: Array<string>;
}

// Shared UI for all individual sims and the raid sim.
export abstract class SimUI extends Component {
  readonly sim: Sim;

  // Emits when anything from the sim, raid, or encounter changes.
  readonly changeEmitter = new TypedEvent<void>();

	readonly resultsPendingElem: HTMLElement;
	readonly resultsContentElem: HTMLElement;

  constructor(parentElem: HTMLElement, sim: Sim, config: SimUIConfig) {
		super(parentElem, 'sim-ui');
    this.sim = sim;
    this.rootElem.innerHTML = simHTML;

    [
      this.sim.changeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));

		Array.from(document.getElementsByClassName('known-issues')).forEach(element => {
			if (config.knownIssues?.length) {
				(element as HTMLElement).style.display = 'initial';
			} else {
				return;
			}

			
			tippy(element, {
				'content': `
				<ul class="known-issues-tooltip">
					${config.knownIssues.map(issue => '<li>' + issue + '</li>').join('')}
				</ul>
				`,
				'allowHTML': true,
			});
		});

		this.resultsPendingElem = this.rootElem.getElementsByClassName('results-pending')[0] as HTMLElement;
		this.resultsContentElem = this.rootElem.getElementsByClassName('results-content')[0] as HTMLElement;
		this.hideAllResults();

		const titleElem = this.rootElem.getElementsByClassName('sim-sidebar-title')[0];
		titleElem.textContent = config.title;

		const simActionsContainer = this.rootElem.getElementsByClassName('sim-sidebar-actions')[0] as HTMLElement;
		const iterationsPicker = new NumberPicker(simActionsContainer, this.sim, {
			label: 'Iterations',
			extraCssClasses: [
				'iterations-picker',
				'within-raid-sim-hide',
			],
			changedEvent: (sim: Sim) => sim.iterationsChangeEmitter,
			getValue: (sim: Sim) => sim.getIterations(),
			setValue: (sim: Sim, newValue: number) => {
				sim.setIterations(newValue);
			},
		});
  }

	addAction(name: string, cssClass: string, actFn: () => void) {
		const simActionsContainer = this.rootElem.getElementsByClassName('sim-sidebar-actions')[0] as HTMLElement;
		const iterationsPicker = this.rootElem.getElementsByClassName('iterations-picker')[0] as HTMLElement;

    const button = document.createElement('button');
    button.classList.add('sim-sidebar-actions-button', cssClass);
    button.textContent = name;
    button.addEventListener('click', actFn);
    simActionsContainer.insertBefore(button, iterationsPicker);
	}

	addTab(title: string, cssClass: string, innerHTML: string) {
		const simTabsContainer = this.rootElem.getElementsByClassName('sim-tabs')[0] as HTMLElement;
		const simTabContentsContainer = this.rootElem.getElementsByClassName('tab-content')[0] as HTMLElement;
		const topBar = simTabsContainer.getElementsByClassName('sim-top-bar')[0] as HTMLElement;

		const contentId = cssClass.replace(/\s+/g, '-') + '-tab';
		const isFirstTab = simTabsContainer.children.length == 1;

		const newTab = document.createElement('li');
		newTab.innerHTML = `<a data-toggle="tab" href="#${contentId}">${title}</a>`;
		newTab.classList.add(cssClass + '-tab');

		const newContent = document.createElement('div');
		newContent.id = contentId;
		newContent.classList.add(cssClass, 'tab-pane', 'fade');
		newContent.innerHTML = innerHTML;

		if (isFirstTab) {
			newTab.classList.add('active');
			newContent.classList.add('active', 'in');
		}

		simTabsContainer.insertBefore(newTab, topBar);
		simTabContentsContainer.appendChild(newContent);
	}

	hideAllResults() {
		this.resultsContentElem.style.display = 'none';
    this.resultsPendingElem.style.display = 'none';
	}

  setResultsPending() {
		this.resultsContentElem.style.display = 'none';
    this.resultsPendingElem.style.display = 'initial';
  }

	setResultsContent(innerHTML: string) {
		this.resultsContentElem.innerHTML = innerHTML;
		this.resultsContentElem.style.display = 'initial';
    this.resultsPendingElem.style.display = 'none';
	}

	// Returns a key suitable for the browser's localStorage feature.
	abstract getStorageKey(postfix: string): string;

	getSettingsStorageKey(): string {
		return this.getStorageKey('__currentSettings__');
	}

	getSavedEncounterStorageKey(): string {
		return this.getStorageKey('__savedEncounter__');
	}
}

const simHTML = `
<div class="sim-root">
  <section class="sim-sidebar">
    <div class="sim-sidebar-title"></div>
    <div class="sim-sidebar-actions within-raid-sim-hide"></div>
    <div class="sim-sidebar-results within-raid-sim-hide">
      <div class="results-pending">
        <div class="loader"></div>
      </div>
      <div class="results-content">
      </div>
		</div>
    <div class="sim-sidebar-footer"></div>
  </section>
  <section class="sim-main">
    <ul class="sim-tabs nav nav-tabs">
      <li class="sim-top-bar">
				<div class="known-issues">Known Issues</div>
				<span class="share-link fa fa-link within-raid-sim-hide"></span>
			</li>
    </ul>
    <div class="tab-content">
    </div>
  </section>
</div>
`;
