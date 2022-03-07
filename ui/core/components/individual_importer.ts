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

export class IndividualImporter<SpecType extends Spec> extends Popup {
	private readonly simUI: IndividualSimUI<SpecType>;
  private readonly importButton: HTMLElement;

  constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>) {
    super(parent);
		this.simUI = simUI;

		this.rootElem.classList.add('individual-importer');
		this.rootElem.innerHTML = `
			<ul class="nav nav-tabs individual-importer-tabs">
				<li class="individual-importer-tab active"><a data-toggle="tab" href="#70upgradesTab">70 Upgrades</a></li>
				<li class="individual-importer-tab"><a data-toggle="tab" href="#addonTab">Addon</a></li>
			</ul>
			<div class="tab-content individual-importer-contents">
				<div id="70upgradesTab" class="tab-pane individual-importer-tab-content fade active in">
					<div class="individual-importer-description">
						<p>
							Import settings from <a href="https://seventyupgrades.com" target="_blank">Seventy Upgrades</a>.
						</p>
						<p>
							This feature imports gear, race, and (optionally) talents. It does NOT import buffs, debuffs, consumes, rotation, or custom stats.
						</p>
						<p>
							To import, paste the output from the site's export option below and click, 'Import'.
						</p>
					</div>
					<textarea class="individual-importer-textarea seventy-upgrades-input"></textarea>
				</div>
				<div id="addonTab" class="tab-pane individual-importer-tab-content fade">
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

		this.addCloseButton();
    this.importButton = this.rootElem.getElementsByClassName('import-button')[0] as HTMLElement;
		this.setup70UpgradesImport();
		this.setupAddonImport();
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

	private setupAddonImport() {
		const addonContent = document.getElementById('addonTab') as HTMLElement;
    const addonInput = this.rootElem.getElementsByClassName('addon-input')[0] as HTMLTextAreaElement;
		this.importButton.addEventListener('click', event => {
			if (!addonContent.classList.contains('active')) {
				return
			}

			try {
				const importJson = JSON.parse(addonInput.value);

				// Parse all the settings.
				const charClass = nameToClass((importJson['class'] as string) || '');
				if (charClass == Class.ClassUnknown) {
					throw new Error('Could not parse Class!');
				}

				const race = nameToRace((importJson['race'] as string) || '');
				if (race == Race.RaceUnknown) {
					throw new Error('Could not parse Race!');
				}

				const talentsStr = (importJson['talents'] as string) || '';
				const equipmentSpec = EquipmentSpec.fromJson(importJson['gear']);

				this.finishImport(charClass, race, equipmentSpec, talentsStr);
			} catch (error) {
				alert('Error importing from addon: ' + error);
			}
		});
	}

	private finishImport(charClass: Class, race: Race, equipmentSpec: EquipmentSpec, talentsStr: string) {
		const playerClass = this.simUI.player.getClass();
		if (charClass != playerClass) {
			throw new Error(`Wrong Class! Expected ${classNames[playerClass]} but found ${classNames[charClass]}!`);
		}

		equipmentSpec.items.forEach(item => {
			if (item.enchant) {
				const dbEnchant = this.simUI.sim.getEnchantFlexible(item.enchant);
				if (dbEnchant) {
					item.enchant = dbEnchant.id;
				} else {
					item.enchant = 0;
				}
			}
		});

		const gear = this.simUI.sim.lookupEquipmentSpec(equipmentSpec);

		const expectedEnchantIds = equipmentSpec.items.map(item => item.enchant);
		const foundEnchantIds = gear.asSpec().items.map(item => item.enchant);
		const missingEnchants = expectedEnchantIds.filter(expectedId => expectedId != 0 && !foundEnchantIds.includes(expectedId));

		const expectedItemIds = equipmentSpec.items.map(item => item.id);
		const foundItemIds = gear.asSpec().items.map(item => item.id);
		const missingItems = expectedItemIds.filter(expectedId => !foundItemIds.includes(expectedId));

		// Now update settings using the parsed values.
		const eventID = TypedEvent.nextEventID();
		TypedEvent.freezeAllAndDo(() => {
			this.simUI.player.setRace(eventID, race);
			this.simUI.player.setGear(eventID, gear);
			if (talentsStr) {
				this.simUI.player.setTalentsString(eventID, talentsStr);
			}
		});

		this.close();

		if (missingItems.length == 0 && missingEnchants.length == 0) {
			alert('Import successful!');
		} else {
			alert('Import successful, but the following IDs were not found in the sim database:' + 
					(missingItems.length == 0 ? '' : '\n\nItems: ' + missingItems.join(', ')) +
					(missingEnchants.length == 0 ? '' : '\n\nEnchants: ' + missingEnchants.join(', ')));
		}
	}
}
