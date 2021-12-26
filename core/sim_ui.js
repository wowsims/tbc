import { Component } from '/tbc/core/components/component.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { Title } from '/tbc/core/components/title.js';
import { TypedEvent } from './typed_event.js';
// Shared UI for all individual sims and the raid sim.
export class SimUI extends Component {
    constructor(parentElem, sim, config) {
        super(parentElem, 'sim-ui');
        // Emits when anything from the sim, raid, or encounter changes.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        this.rootElem.innerHTML = simHTML;
        [
            this.sim.changeEmitter,
        ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
        if (config.knownIssues && config.knownIssues.length) {
            const knownIssuesContainer = document.getElementsByClassName('known-issues')[0];
            knownIssuesContainer.style.display = 'initial';
            tippy(knownIssuesContainer, {
                content: `
				<ul class="known-issues-tooltip">
					${config.knownIssues.map(issue => '<li>' + issue + '</li>').join('')}
				</ul>
				`,
                allowHTML: true,
                interactive: true,
            });
        }
        this.resultsPendingElem = this.rootElem.getElementsByClassName('results-pending')[0];
        this.resultsContentElem = this.rootElem.getElementsByClassName('results-content')[0];
        this.hideAllResults();
        const titleElem = this.rootElem.getElementsByClassName('sim-sidebar-title')[0];
        const title = new Title(titleElem, config.spec);
        const simActionsContainer = this.rootElem.getElementsByClassName('sim-sidebar-actions')[0];
        const iterationsPicker = new NumberPicker(simActionsContainer, this.sim, {
            label: 'Iterations',
            extraCssClasses: [
                'iterations-picker',
                'within-raid-sim-hide',
            ],
            changedEvent: (sim) => sim.iterationsChangeEmitter,
            getValue: (sim) => sim.getIterations(),
            setValue: (eventID, sim, newValue) => {
                sim.setIterations(eventID, newValue);
            },
        });
    }
    addAction(name, cssClass, actFn) {
        const simActionsContainer = this.rootElem.getElementsByClassName('sim-sidebar-actions')[0];
        const iterationsPicker = this.rootElem.getElementsByClassName('iterations-picker')[0];
        const button = document.createElement('button');
        button.classList.add('sim-sidebar-actions-button', cssClass);
        button.textContent = name;
        button.addEventListener('click', actFn);
        simActionsContainer.insertBefore(button, iterationsPicker);
    }
    addTab(title, cssClass, innerHTML) {
        const simTabsContainer = this.rootElem.getElementsByClassName('sim-tabs')[0];
        const simTabContentsContainer = this.rootElem.getElementsByClassName('tab-content')[0];
        const topBar = simTabsContainer.getElementsByClassName('sim-top-bar')[0];
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
    addToolbarItem(elem) {
        const topBar = this.rootElem.getElementsByClassName('sim-top-bar')[0];
        elem.classList.add('sim-top-bar-item');
        topBar.appendChild(elem);
    }
    hideAllResults() {
        this.resultsContentElem.style.display = 'none';
        this.resultsPendingElem.style.display = 'none';
    }
    setResultsPending() {
        this.resultsContentElem.style.display = 'none';
        this.resultsPendingElem.style.display = 'initial';
    }
    setResultsContent(innerHTML) {
        this.resultsContentElem.innerHTML = innerHTML;
        this.resultsContentElem.style.display = 'initial';
        this.resultsPendingElem.style.display = 'none';
    }
    getSettingsStorageKey() {
        return this.getStorageKey('__currentSettings__');
    }
    getSavedEncounterStorageKey() {
        // By skipping the call to this.getStorageKey(), saved encounters will be
        // shared across all sims.
        return 'sharedData__savedEncounter__';
    }
    isIndividualSim() {
        return this.rootElem.classList.contains('individual-sim-ui');
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
		<div class="sim-toolbar">
			<ul class="sim-tabs nav nav-tabs">
				<li class="sim-top-bar">
					<div class="known-issues">Known Issues</div>
				</li>
			</ul>
    </div>
    <div class="tab-content">
    </div>
  </section>
</div>
`;
