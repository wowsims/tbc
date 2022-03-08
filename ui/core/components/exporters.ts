import { Class } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Player } from '/tbc/core/player.js';
import { classNames, nameToClass, nameToRace } from '/tbc/core/proto_utils/names.js';
import { talentSpellIdsToTalentString } from '/tbc/core/talents/factory.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

import { Popup } from './popup.js';

declare var $: any;

export function newIndividualExporters<SpecType extends Spec>(simUI: IndividualSimUI<SpecType>): HTMLElement {
	const exportSettings = document.createElement('div');
	exportSettings.classList.add('export-settings', 'sim-dropdown-menu');
	exportSettings.innerHTML = `
		<span id="exportMenuLink" class="dropdown-toggle fas fa-file-export" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="exportMenuLink">
			<span class="dropdown-item">Link</span>
			<span class="dropdown-item">JSON</span>
		</div>
	`;
	tippy(exportSettings, {
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

	addMenuItem('Link', () => {
		const protoBytes = IndividualSimSettings.toBinary(simUI.toProto());
		const deflated = pako.deflate(protoBytes, { to: 'string' });
		const encoded = btoa(String.fromCharCode(...deflated));

		const linkUrl = new URL(window.location.href);
		linkUrl.hash = encoded;
		
		if (navigator.clipboard == undefined) {
			alert(linkUrl.toString());
		} else {
			navigator.clipboard.writeText(linkUrl.toString());
			alert('Current settings copied to clipboard!');
		}
	});

	return exportSettings;
}

export class JsonExporter<SpecType extends Spec> extends Popup {
	private readonly simUI: IndividualSimUI<SpecType>;
  private readonly importButton: HTMLElement;

  constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>) {
    super(parent);
		this.simUI = simUI;

		this.rootElem.classList.add('individual-importer');
		this.rootElem.innerHTML = `
			<div id="70upgradesTab" class="tab-pane individual-importer-tab-content fade active in">
				<textarea class="exporter-textarea"></textarea>
			</div>
			<div class="actions-row">
				<button class="individual-importer-action-button sim-button clipboard-button">COPY TO CLIPBOARD</button>
				<button class="individual-importer-action-button sim-button download-button">DOWNLOAD</button>
			</div>
		`;

		this.addCloseButton();
    this.importButton = this.rootElem.getElementsByClassName('import-button')[0] as HTMLElement;
		this.setup70UpgradesImport();
  }

	private setup70UpgradesImport() {
		const seventyUpgradesContent = document.getElementById('70upgradesTab') as HTMLElement;
    const seventyUpgradesInput = this.rootElem.getElementsByClassName('seventy-upgrades-input')[0] as HTMLTextAreaElement;
		this.importButton.addEventListener('click', event => {
			if (!seventyUpgradesContent.classList.contains('active')) {
				return
			}

			try {
				const importJson = JSON.parse(seventyUpgradesInput.value);

				// Parse all the settings.
				const charClass = nameToClass((importJson.character.gameClass as string) || '');
				if (charClass == Class.ClassUnknown) {
					throw new Error('Could not parse Class!');
				}

				const race = nameToRace((importJson.character.race as string) || '');
				if (race == Race.RaceUnknown) {
					throw new Error('Could not parse Race!');
				}

				let talentsStr = '';
				if (importJson.talents && importJson.talents.length > 0) {
					const talentIds = (importJson.talents as Array<any>).map(talentJson => talentJson.spellId);
					talentsStr = talentSpellIdsToTalentString(charClass, talentIds);
				}

				let equipmentSpec = EquipmentSpec.create();
				(importJson.items as Array<any>).forEach(itemJson => {
					let itemSpec = ItemSpec.create();
					itemSpec.id = itemJson.id;
					if (itemJson.enchant?.id) {
						itemSpec.enchant = itemJson.enchant.id;
					}
					if (itemJson.gems) {
						itemSpec.gems = (itemJson.gems as Array<any>).filter(gemJson => gemJson?.id).map(gemJson => gemJson.id);
					}
					equipmentSpec.items.push(itemSpec);
				});

				const gear = this.simUI.sim.lookupEquipmentSpec(equipmentSpec);

				this.finishImport(charClass, race, equipmentSpec, talentsStr);
			} catch (error) {
				alert('Error importing from seventyUpgrades: ' + error);
			}
		});
	}
}
