import { Encounter } from '/tbc/core/encounter.js';
import { Raid } from '/tbc/core/raid.js';
import { Sim } from '/tbc/core/sim.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { DetailedResults } from '/tbc/core/components/detailed_results.js';
import { LogRunner } from '/tbc/core/components/log_runner.js';
import { SavedDataConfig } from '/tbc/core/components/saved_data_manager.js';
import { SavedDataManager } from '/tbc/core/components/saved_data_manager.js';

import { RaidPicker, PresetSpecSettings } from './raid_picker.js';

declare var tippy: any;

export interface RaidSimConfig {
	knownIssues?: Array<string>,
	presets: Array<PresetSpecSettings<any>>,
}

export class RaidSimUI extends SimUI{
  private readonly config: RaidSimConfig;

  constructor(parentElem: HTMLElement, config: RaidSimConfig) {
		super(parentElem, new Sim(), {
			title: 'TBC RaidSim',
			knownIssues: config.knownIssues,
		});
		this.rootElem.classList.add('raid-sim-ui');

    this.config = config;

		this.addSidebarComponents();
		this.addTopbarComponents();
		this.addRaidTab();
  }

	private addSidebarComponents() {
		//addRaidSimAction(this);
		//addStatWeightsAction(this, this.individualConfig.epStats, this.individualConfig.epReferenceStat);
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

	// Returns the actual key to use for local storage, based on the given key part and the site context.
	getStorageKey(keyPart: string): string {
		return '__raid__' + keyPart;
	}
}
