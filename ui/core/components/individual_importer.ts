import { Class } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Player } from '/tbc/core/player.js';
import { classNames, nameToClass, nameToRace } from '/tbc/core/proto_utils/names.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

import { CloseButton } from './close_button.js';
import { Component } from './component.js';

declare var $: any;

export class IndividualImporter<SpecType extends Spec> extends Component {
	private readonly simUI: IndividualSimUI<SpecType>;
  private readonly importButton: HTMLElement;

  constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>) {
    super(parent, 'individual-importer');
		this.simUI = simUI;

		this.rootElem.id = 'individualImporter';
		this.rootElem.innerHTML = `
			<ul class="nav nav-tabs individual-importer-tabs">
				<li class="individual-importer-tab active"><a data-toggle="tab" href="#addonTab">Addon</a></li>
			</ul>
			<div class="tab-content individual-importer-contents">
				<div id="addonTab" class="tab-pane individual-importer-tab-content active in">
					<div class="individual-importer-description">
						<p>
							Import settings from the <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank">WoWSims Importer In-Game Addon</a>.
						</p>
						<p>
							This feature imports gear, race, and talents. It does NOT import buffs, debuffs, consumes, rotation, or custom stats.
						</p>
						<p>
							To import, paste the output from the addon below and click, 'Import'.
						</p>
					</div>
					<textarea class="individual-importer-textarea addon-input"></textarea>
				</div>
				<div class="actions-row">
					<button class="individual-importer-action-button sim-button import-button">IMPORT</button>
				</div>
			</div>
		`;

    this.importButton = this.rootElem.getElementsByClassName('import-button')[0] as HTMLElement;

		const computedStyles = window.getComputedStyle(parent);
		this.rootElem.style.setProperty('--main-text-color', computedStyles.getPropertyValue('--main-text-color').trim());
		this.rootElem.style.setProperty('--theme-color-primary', computedStyles.getPropertyValue('--theme-color-primary').trim());
		this.rootElem.style.setProperty('--theme-color-background', computedStyles.getPropertyValue('--theme-color-background').trim());
		this.rootElem.style.setProperty('--theme-color-background-raw', computedStyles.getPropertyValue('--theme-color-background-raw').trim());

		new CloseButton(this.rootElem, () => {
			$('#individualImporter').bPopup().close();
		});

		$('#individualImporter').bPopup({
			closeClass: 'item-picker-close',
			onClose: () => {
				this.rootElem.remove();
			},
		});

		this.setupAddonImport();
  }

	private setupAddonImport() {
    const addonInput = this.rootElem.getElementsByClassName('addon-input')[0] as HTMLTextAreaElement;
		this.importButton.addEventListener('click', event => {
			try {
				const importJson = JSON.parse(addonInput.value);

				// Parse all the settings.
				const charClass = nameToClass((importJson['class'] as string) || '');
				if (charClass == Class.ClassUnknown) {
					throw new Error('Could not parse Class!');
				}
				const playerClass = this.simUI.player.getClass();
				if (charClass != playerClass) {
					throw new Error(`Wrong Class! Expected ${classNames[playerClass]} but found ${classNames[charClass]}!`);
				}

				const race = nameToRace((importJson['race'] as string) || '');
				if (race == Race.RaceUnknown) {
					throw new Error('Could not parse Race!');
				}

				const talentsStr = (importJson['talents'] as string) || '';

				const equipmentSpec = EquipmentSpec.fromJson(importJson['gear']);
				const gear = this.simUI.sim.lookupEquipmentSpec(equipmentSpec);

				// Now update settings using the parsed values.
				const eventID = TypedEvent.nextEventID();
				TypedEvent.freezeAllAndDo(() => {
					this.simUI.player.setRace(eventID, race);
					this.simUI.player.setTalentsString(eventID, talentsStr);
					this.simUI.player.setGear(eventID, gear);
				});

				$('#individualImporter').bPopup().close();
			} catch (error) {
				alert('Error importing from addon: ' + error);
			}
		});
	}
}
