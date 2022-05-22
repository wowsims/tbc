import { Component } from '/tbc/core/components/component.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { ResultsViewer } from '/tbc/core/components/results_viewer.js';
import { Title } from '/tbc/core/components/title.js';
import { TypedEvent } from './typed_event.js';
const noticeText = 'We are looking for help migrating our sims to Wrath of the Lich King. If you\'d like to participate in a fun side project working with an open-source community please <a href="https://discord.gg/jJMPr9JWwx" target="_blank">join our discord!</a>';
// Shared UI for all individual sims and the raid sim.
export class SimUI extends Component {
    constructor(parentElem, sim, config) {
        super(parentElem, 'sim-ui');
        this.sim = sim;
        this.rootElem.innerHTML = simHTML;
        this.isWithinRaidSim = this.rootElem.closest('.within-raid-sim') != null;
        if (!this.isWithinRaidSim) {
            this.rootElem.classList.add('not-within-raid-sim');
        }
        this.changeEmitter = TypedEvent.onAny([
            this.sim.changeEmitter,
        ], 'SimUIChange');
        const updateShowThreatMetrics = () => {
            if (this.sim.getShowThreatMetrics()) {
                this.rootElem.classList.remove('hide-threat-metrics');
            }
            else {
                this.rootElem.classList.add('hide-threat-metrics');
            }
        };
        updateShowThreatMetrics();
        this.sim.showThreatMetricsChangeEmitter.on(updateShowThreatMetrics);
        const updateShowExperimental = () => {
            if (this.sim.getShowExperimental()) {
                this.rootElem.classList.remove('hide-experimental');
            }
            else {
                this.rootElem.classList.add('hide-experimental');
            }
        };
        updateShowExperimental();
        this.sim.showExperimentalChangeEmitter.on(updateShowExperimental);
        const noticesElem = document.getElementsByClassName('notices')[0];
        if (noticeText) {
            tippy(noticesElem, {
                content: noticeText,
                allowHTML: true,
                interactive: true,
            });
        }
        else {
            noticesElem.remove();
        }
        this.warnings = [];
        const warningsElem = document.getElementsByClassName('warnings')[0];
        this.warningsTippy = tippy(warningsElem, {
            content: '',
            allowHTML: true,
        });
        this.updateWarnings();
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
        const resultsViewerElem = this.rootElem.getElementsByClassName('sim-sidebar-results')[0];
        this.resultsViewer = new ResultsViewer(resultsViewerElem);
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
        const reportBug = document.createElement('span');
        reportBug.classList.add('report-bug', 'fa', 'fa-bug');
        tippy(reportBug, {
            'content': 'Report a bug / request a feature',
            'allowHTML': true,
        });
        reportBug.addEventListener('click', event => {
            window.open('https://github.com/wowsims/tbc/issues/new/choose', '_blank');
        });
        this.addToolbarItem(reportBug);
        if (!this.isWithinRaidSim) {
            window.addEventListener('message', async (event) => {
                if (event.data == 'runOnce') {
                    this.runSimOnce();
                }
            });
        }
        const downloadBinary = document.createElement('span');
        // downloadBinary.src = "/tbc/assets/gauge.svg"
        downloadBinary.classList.add('downbin');
        downloadBinary.addEventListener('click', event => {
            window.open('https://github.com/wowsims/tbc/releases', '_blank');
        });
        if (document.location.href.includes("localhost")) {
            fetch(document.location.protocol + "//" + document.location.host + "/version").then(resp => {
                resp.json()
                    .then((versionInfo) => {
                    if (versionInfo.outdated == 2) {
                        tippy(downloadBinary, {
                            'content': 'Newer version of simulator available for download',
                            'allowHTML': true,
                        });
                        downloadBinary.classList.add('downbinalert');
                        this.addToolbarItem(downloadBinary);
                    }
                })
                    .catch(error => {
                    console.warn('No version info found!');
                });
            });
        }
        else {
            tippy(downloadBinary, {
                'content': 'Download simulator for faster simulating',
                'allowHTML': true,
            });
            downloadBinary.classList.add('downbinnorm');
            this.addToolbarItem(downloadBinary);
        }
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
    updateWarnings() {
        const activeWarnings = this.warnings.filter(warning => warning.shouldDisplay());
        const warningsElem = document.getElementsByClassName('warnings')[0];
        if (activeWarnings.length == 0) {
            warningsElem.style.display = 'none';
        }
        else {
            warningsElem.style.display = 'initial';
            this.warningsTippy.setContent(`
				<ul class="known-issues-tooltip">
					${activeWarnings.map(warning => '<li>' + warning.getContent() + '</li>').join('')}
				</ul>`);
        }
    }
    addWarning(warning) {
        this.warnings.push(warning);
        warning.updateOn.on(() => this.updateWarnings());
        this.updateWarnings();
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
    async runSim(onProgress) {
        this.resultsViewer.setPending();
        try {
            const result = await this.sim.runRaidSim(TypedEvent.nextEventID(), onProgress);
        }
        catch (e) {
            this.resultsViewer.hideAll();
            alert(e);
        }
    }
    async runSimOnce() {
        this.resultsViewer.setPending();
        try {
            const result = await this.sim.runRaidSimWithLogs(TypedEvent.nextEventID());
        }
        catch (e) {
            this.resultsViewer.hideAll();
            alert(e);
        }
    }
}
const simHTML = `
<div class="sim-root">
  <section class="sim-sidebar">
    <div class="sim-sidebar-title"></div>
    <div class="sim-sidebar-actions within-raid-sim-hide"></div>
    <div class="sim-sidebar-results within-raid-sim-hide"></div>
    <div class="sim-sidebar-footer"></div>
  </section>
  <section class="sim-main">
		<div class="sim-toolbar">
			<ul class="sim-tabs nav nav-tabs">
				<li class="sim-top-bar">
					<span class="notices fas fa-exclamation-circle"></span>
					<span class="warnings fa fa-exclamation-triangle"></span>
					<div class="known-issues">Known Issues</div>
				</li>
			</ul>
    </div>
    <div class="tab-content">
    </div>
  </section>
</div>
`;
