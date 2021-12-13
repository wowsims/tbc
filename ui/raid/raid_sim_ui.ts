import { Encounter } from '/tbc/core/encounter.js';
import { Raid } from '/tbc/core/raid.js';
import { Sim } from '/tbc/core/sim.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { DetailedResults } from '/tbc/core/components/detailed_results.js';
import { EncounterPicker, EncounterPickerConfig } from '/tbc/core/components/encounter_picker.js';
import { LogRunner } from '/tbc/core/components/log_runner.js';
import { SavedDataConfig } from '/tbc/core/components/saved_data_manager.js';
import { SavedDataManager } from '/tbc/core/components/saved_data_manager.js';
import { addRaidSimAction } from '/tbc/core/components/raid_sim_action.js';

import { BlessingsPicker } from './blessings_picker.js';
import { RaidPicker, PresetSpecSettings } from './raid_picker.js';

declare var tippy: any;

export interface RaidSimConfig {
	knownIssues?: Array<string>,
	presets: Array<PresetSpecSettings<any>>,
}

export class RaidSimUI extends SimUI{
  private readonly config: RaidSimConfig;
	private readonly implementedSpecs: Array<Spec>;

  constructor(parentElem: HTMLElement, config: RaidSimConfig) {
		super(parentElem, new Sim(), {
			title: 'TBC RaidSim',
			knownIssues: config.knownIssues,
		});
		this.rootElem.classList.add('raid-sim-ui');

    this.config = config;
		
		this.implementedSpecs = [...new Set(config.presets.map(preset => preset.spec))];

		this.addSidebarComponents();
		this.addTopbarComponents();
		this.addRaidTab();
		this.addSettingsTab();
		this.addDetailedResultsTab();
		this.addLogTab();
  }

	private addSidebarComponents() {
		addRaidSimAction(this);
	}

	private addTopbarComponents() {
		//Array.from(document.getElementsByClassName('share-link')).forEach(element => {
		//	tippy(element, {
		//		'content': 'Shareable link',
		//		'allowHTML': true,
		//	});

		//	element.addEventListener('click', event => {
		//		const jsonStr = JSON.stringify(this.sim.toJson());
    //    const val = pako.deflate(jsonStr, { to: 'string' });
    //    const encoded = btoa(String.fromCharCode(...val));
		//		
    //    const linkUrl = new URL(window.location.href);
    //    linkUrl.hash = encoded;
    //    if (navigator.clipboard == undefined) {
    //      alert(linkUrl.toString());
    //    } else {
    //      navigator.clipboard.writeText(linkUrl.toString());
    //      alert('Current settings copied to clipboard!');
    //    }
		//	});
		//});
	}

	private addRaidTab() {
		this.addTab('Raid', 'raid-tab', `
			<div class="raid-picker">
			</div>
			<div class="saved-raids-div">
				<div class="saved-raids-manager">
				</div>
			</div>
		`);

		const raidPicker = new RaidPicker(this.rootElem.getElementsByClassName('raid-picker')[0] as HTMLElement, this.sim.raid, this.config.presets);
	}

	private addSettingsTab() {
		this.addTab('Settings', 'raid-settings-tab', `
			<div class="raid-settings-sections">
				<div class="raid-settings-section-container">
					<section class="settings-section raid-encounter-section">
						<label>Encounter</label>
					</section>
				</div>
				<div class="blessings-section-container">
					<section class="settings-section blessings-section">
						<label>Blessings</label>
					</section>
				</div>
			</div>
			<div class="settings-bottom-bar">
				<div class="saved-encounter-manager">
				</div>
			</div>
		`);

    const encounterSectionElem = this.rootElem.getElementsByClassName('raid-encounter-section')[0] as HTMLElement;
		new EncounterPicker(encounterSectionElem, this.sim.encounter, {
			showTargetArmor: true,
			showNumTargets: true,
		});
    const savedEncounterManager = new SavedDataManager<Encounter, EncounterProto>(this.rootElem.getElementsByClassName('saved-encounter-manager')[0] as HTMLElement, this.sim.encounter, {
      label: 'Encounter',
			storageKey: this.getSavedEncounterStorageKey(),
      getData: (encounter: Encounter) => encounter.toProto(),
      setData: (encounter: Encounter, newEncounter: EncounterProto) => encounter.fromProto(newEncounter),
      changeEmitters: [this.sim.encounter.changeEmitter],
      equals: (a: EncounterProto, b: EncounterProto) => EncounterProto.equals(a, b),
      toJson: (a: EncounterProto) => EncounterProto.toJson(a),
      fromJson: (obj: any) => EncounterProto.fromJson(obj),
    });

		const blessingsPicker = new BlessingsPicker(this.rootElem.getElementsByClassName('blessings-section')[0] as HTMLElement, this, this.implementedSpecs);

		this.sim.waitForInit().then(() => {
			savedEncounterManager.loadUserData();
		});
	}

	private addDetailedResultsTab() {
		this.addTab('Detailed Results', 'detailed-results-tab', `
			<div class="detailed-results">
			</div>
		`);

    const detailedResults = new DetailedResults(this.rootElem.getElementsByClassName('detailed-results')[0] as HTMLElement, this);
	}

	private addLogTab() {
		this.addTab('Log', 'log-tab', `
			<div class="log-runner">
			</div>
		`);

    const logRunner = new LogRunner(this.rootElem.getElementsByClassName('log-runner')[0] as HTMLElement, this);
	}

	// Returns the actual key to use for local storage, based on the given key part and the site context.
	getStorageKey(keyPart: string): string {
		return '__raid__' + keyPart;
	}
}
