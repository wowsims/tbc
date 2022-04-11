import { Exporter } from '/tbc/core/components/exporters.js';
import { Importer } from '/tbc/core/components/importers.js';
import { RaidSimSettings } from '/tbc/core/proto/ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

import { RaidSimUI } from './raid_sim_ui.js';

declare var $: any;
declare var tippy: any;

export function newRaidImporters(simUI: RaidSimUI): HTMLElement {
	const importSettings = document.createElement('div');
	importSettings.classList.add('import-settings', 'sim-dropdown-menu');
	importSettings.innerHTML = `
		<span id="importMenuLink" class="dropdown-toggle fas fa-file-import" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="importMenuLink">
		</div>
	`;
	const linkElem = importSettings.getElementsByClassName('dropdown-toggle')[0] as HTMLElement;
	tippy(linkElem, {
		'content': 'Import',
		'allowHTML': true,
	});

	const menuElem = importSettings.getElementsByClassName('dropdown-menu')[0] as HTMLElement;
	const addMenuItem = (label: string, onClick: () => void) => {
		const itemElem = document.createElement('span');
		itemElem.classList.add('dropdown-item');
		itemElem.textContent = label;
		itemElem.addEventListener('click', onClick);
		menuElem.appendChild(itemElem);
	};

	addMenuItem('Json', () => new RaidJsonImporter(menuElem, simUI));

	return importSettings;
}

export function newRaidExporters(simUI: RaidSimUI): HTMLElement {
	const exportSettings = document.createElement('div');
	exportSettings.classList.add('export-settings', 'sim-dropdown-menu');
	exportSettings.innerHTML = `
		<span id="exportMenuLink" class="dropdown-toggle fas fa-file-export" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="exportMenuLink">
		</div>
	`;
	const linkElem = exportSettings.getElementsByClassName('dropdown-toggle')[0] as HTMLElement;
	tippy(linkElem, {
		'content': 'Export',
		'allowHTML': true,
	});

	const menuElem = exportSettings.getElementsByClassName('dropdown-menu')[0] as HTMLElement;
	const addMenuItem = (label: string, onClick: () => void) => {
		const itemElem = document.createElement('span');
		itemElem.classList.add('dropdown-item');
		itemElem.textContent = label;
		itemElem.addEventListener('click', onClick);
		menuElem.appendChild(itemElem);
	};

	addMenuItem('Json', () => new RaidJsonExporter(menuElem, simUI));

	return exportSettings;
}

class RaidJsonImporter extends Importer {
	private readonly simUI: RaidSimUI;
	constructor(parent: HTMLElement, simUI: RaidSimUI) {
		super(parent, 'JSON Import');
		this.simUI = simUI;

		this.descriptionElem.innerHTML = `
			<p>
				Import settings from a JSON text file, which can be created using the JSON Export feature of this site.
			</p>
			<p>
				To import, paste the JSON text below and click, 'Import'.
			</p>
		`;
	}

	onImport(data: string) {
		const settings = RaidSimSettings.fromJsonString(data);
		this.simUI.fromProto(TypedEvent.nextEventID(), settings);
		this.close();
	}
}

class RaidJsonExporter extends Exporter {
	private readonly simUI: RaidSimUI;

	constructor(parent: HTMLElement, simUI: RaidSimUI) {
		super(parent, 'JSON Export', true);
		this.simUI = simUI;
		this.init();
	}

	getData(): string {
		return JSON.stringify(RaidSimSettings.toJson(this.simUI.toProto()), null, 2);
	}
}
