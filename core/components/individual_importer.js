import { Class } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { classNames, nameToClass, nameToRace } from '/tbc/core/proto_utils/names.js';
import { talentSpellIdsToTalentString } from '/tbc/core/talents/factory.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { CloseButton } from './close_button.js';
import { Component } from './component.js';
export class IndividualImporter extends Component {
    constructor(parent, simUI) {
        super(parent, 'individual-importer');
        this.simUI = simUI;
        this.rootElem.id = 'individualImporter';
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
        this.importButton = this.rootElem.getElementsByClassName('import-button')[0];
        const computedStyles = window.getComputedStyle(parent);
        this.rootElem.style.setProperty('--main-text-color', computedStyles.getPropertyValue('--main-text-color').trim());
        this.rootElem.style.setProperty('--theme-color-primary', computedStyles.getPropertyValue('--theme-color-primary').trim());
        this.rootElem.style.setProperty('--theme-color-background', computedStyles.getPropertyValue('--theme-color-background').trim());
        this.rootElem.style.setProperty('--theme-color-background-raw', computedStyles.getPropertyValue('--theme-color-background-raw').trim());
        new CloseButton(this.rootElem, () => {
            $('#individualImporter').bPopup().close();
            this.rootElem.remove();
        });
        $('#individualImporter').bPopup({
            closeClass: 'item-picker-close',
            onClose: () => {
                this.rootElem.remove();
            },
        });
        this.setup70UpgradesImport();
        this.setupAddonImport();
    }
    setup70UpgradesImport() {
        const seventyUpgradesContent = document.getElementById('70upgradesTab');
        const seventyUpgradesInput = this.rootElem.getElementsByClassName('seventy-upgrades-input')[0];
        this.importButton.addEventListener('click', event => {
            if (!seventyUpgradesContent.classList.contains('active')) {
                return;
            }
            try {
                const importJson = JSON.parse(seventyUpgradesInput.value);
                // Parse all the settings.
                const charClass = nameToClass(importJson.character.gameClass || '');
                if (charClass == Class.ClassUnknown) {
                    throw new Error('Could not parse Class!');
                }
                const race = nameToRace(importJson.character.race || '');
                if (race == Race.RaceUnknown) {
                    throw new Error('Could not parse Race!');
                }
                let talentsStr = '';
                if (importJson.talents && importJson.talents.length > 0) {
                    const talentIds = importJson.talents.map(talentJson => talentJson.spellId);
                    talentsStr = talentSpellIdsToTalentString(charClass, talentIds);
                }
                let equipmentSpec = EquipmentSpec.create();
                importJson.items.forEach(itemJson => {
                    let itemSpec = ItemSpec.create();
                    itemSpec.id = itemJson.id;
                    if (itemJson.enchant?.id) {
                        itemSpec.enchant = itemJson.enchant.id;
                    }
                    if (itemJson.gems) {
                        itemSpec.gems = itemJson.gems.filter(gemJson => gemJson?.id).map(gemJson => gemJson.id);
                    }
                    equipmentSpec.items.push(itemSpec);
                });
                const gear = this.simUI.sim.lookupEquipmentSpec(equipmentSpec);
                this.finishImport(charClass, race, equipmentSpec, talentsStr);
            }
            catch (error) {
                alert('Error importing from seventyUpgrades: ' + error);
            }
        });
    }
    setupAddonImport() {
        const addonContent = document.getElementById('addonTab');
        const addonInput = this.rootElem.getElementsByClassName('addon-input')[0];
        this.importButton.addEventListener('click', event => {
            if (!addonContent.classList.contains('active')) {
                return;
            }
            try {
                const importJson = JSON.parse(addonInput.value);
                // Parse all the settings.
                const charClass = nameToClass(importJson['class'] || '');
                if (charClass == Class.ClassUnknown) {
                    throw new Error('Could not parse Class!');
                }
                const race = nameToRace(importJson['race'] || '');
                if (race == Race.RaceUnknown) {
                    throw new Error('Could not parse Race!');
                }
                const talentsStr = importJson['talents'] || '';
                const equipmentSpec = EquipmentSpec.fromJson(importJson['gear']);
                this.finishImport(charClass, race, equipmentSpec, talentsStr);
            }
            catch (error) {
                alert('Error importing from addon: ' + error);
            }
        });
    }
    finishImport(charClass, race, equipmentSpec, talentsStr) {
        const playerClass = this.simUI.player.getClass();
        if (charClass != playerClass) {
            throw new Error(`Wrong Class! Expected ${classNames[playerClass]} but found ${classNames[charClass]}!`);
        }
        equipmentSpec.items.forEach(item => {
            if (item.enchant) {
                const dbEnchant = this.simUI.sim.getEnchantFlexible(item.enchant);
                if (dbEnchant) {
                    item.enchant = dbEnchant.id;
                }
                else {
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
        $('#individualImporter').bPopup().close();
        if (missingItems.length == 0 && missingEnchants.length == 0) {
            alert('Import successful!');
        }
        else {
            alert('Import successful, but the following IDs were not found in the sim database:' +
                (missingItems.length == 0 ? '' : '\n\nItems: ' + missingItems.join(', ')) +
                (missingEnchants.length == 0 ? '' : '\n\nEnchants: ' + missingEnchants.join(', ')));
        }
    }
}
