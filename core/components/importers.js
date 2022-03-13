import { Class } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { IndividualSimSettings } from '/tbc/core/proto/ui.js';
import { classNames, nameToClass, nameToRace } from '/tbc/core/proto_utils/names.js';
import { talentSpellIdsToTalentString } from '/tbc/core/talents/factory.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Popup } from './popup.js';
export function newIndividualImporters(simUI) {
    const importSettings = document.createElement('div');
    importSettings.classList.add('import-settings', 'sim-dropdown-menu');
    importSettings.innerHTML = `
		<span id="importMenuLink" class="dropdown-toggle fas fa-file-import" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="importMenuLink">
		</div>
	`;
    const linkElem = importSettings.getElementsByClassName('dropdown-toggle')[0];
    tippy(linkElem, {
        'content': 'Import',
        'allowHTML': true,
    });
    const menuElem = importSettings.getElementsByClassName('dropdown-menu')[0];
    const addMenuItem = (label, onClick) => {
        const itemElem = document.createElement('span');
        itemElem.classList.add('dropdown-item');
        itemElem.textContent = label;
        itemElem.addEventListener('click', onClick);
        menuElem.appendChild(itemElem);
    };
    addMenuItem('Json', () => new IndividualJsonImporter(menuElem, simUI));
    addMenuItem('70U', () => new Individual70UImporter(menuElem, simUI));
    addMenuItem('Addon', () => new IndividualAddonImporter(menuElem, simUI));
    return importSettings;
}
class Importer extends Popup {
    constructor(parent, title) {
        super(parent);
        const uploadInputId = 'upload-input-' + title.toLowerCase().replaceAll(' ', '-');
        this.rootElem.classList.add('importer');
        this.rootElem.innerHTML = `
			<span class="importer-title">${title}</span>
			<div class="import-description">
			</div>
			<div class="import-content">
				<textarea class="importer-textarea"></textarea>
			</div>
			<div class="actions-row">
				<label for="${uploadInputId}" class="importer-button sim-button upload-button">UPLOAD FROM FILE</label>
				<input type="file" id="${uploadInputId}" class="importer-upload-input" hidden>
				<button class="importer-button sim-button import-button">IMPORT</button>
			</div>
		`;
        this.addCloseButton();
        this.textElem = this.rootElem.getElementsByClassName('importer-textarea')[0];
        this.descriptionElem = this.rootElem.getElementsByClassName('import-description')[0];
        const uploadInput = this.rootElem.getElementsByClassName('importer-upload-input')[0];
        uploadInput.addEventListener('change', async (event) => {
            const data = await event.target.files[0].text();
            this.textElem.textContent = data;
        });
        const importButton = this.rootElem.getElementsByClassName('import-button')[0];
        importButton.addEventListener('click', event => {
            try {
                this.onImport(this.textElem.value || '');
            }
            catch (error) {
                alert('Import error: ' + error);
            }
        });
    }
    finishIndividualImport(simUI, charClass, race, equipmentSpec, talentsStr) {
        const playerClass = simUI.player.getClass();
        if (charClass != playerClass) {
            throw new Error(`Wrong Class! Expected ${classNames[playerClass]} but found ${classNames[charClass]}!`);
        }
        equipmentSpec.items.forEach(item => {
            if (item.enchant) {
                const dbEnchant = simUI.sim.getEnchantFlexible(item.enchant);
                if (dbEnchant) {
                    item.enchant = dbEnchant.id;
                }
                else {
                    item.enchant = 0;
                }
            }
        });
        const gear = simUI.sim.lookupEquipmentSpec(equipmentSpec);
        const expectedEnchantIds = equipmentSpec.items.map(item => item.enchant);
        const foundEnchantIds = gear.asSpec().items.map(item => item.enchant);
        const missingEnchants = expectedEnchantIds.filter(expectedId => expectedId != 0 && !foundEnchantIds.includes(expectedId));
        const expectedItemIds = equipmentSpec.items.map(item => item.id);
        const foundItemIds = gear.asSpec().items.map(item => item.id);
        const missingItems = expectedItemIds.filter(expectedId => !foundItemIds.includes(expectedId));
        // Now update settings using the parsed values.
        const eventID = TypedEvent.nextEventID();
        TypedEvent.freezeAllAndDo(() => {
            simUI.player.setRace(eventID, race);
            simUI.player.setGear(eventID, gear);
            if (talentsStr) {
                simUI.player.setTalentsString(eventID, talentsStr);
            }
        });
        this.close();
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
class IndividualJsonImporter extends Importer {
    constructor(parent, simUI) {
        super(parent, 'JSON Import');
        this.simUI = simUI;
        this.descriptionElem.innerHTML = `
			<p>
				Import settings from a JSON text file, which can be created using the JSON Export feature of this site.</a>.
			</p>
			<p>
				To import, paste the JSON text below and click, 'Import'.
			</p>
		`;
    }
    onImport(data) {
        const proto = IndividualSimSettings.fromJsonString(data);
        this.simUI.fromProto(TypedEvent.nextEventID(), proto);
        this.close();
    }
}
class Individual70UImporter extends Importer {
    constructor(parent, simUI) {
        super(parent, '70 Upgrades Import');
        this.simUI = simUI;
        this.descriptionElem.innerHTML = `
			<p>
				Import settings from <a href="https://seventyupgrades.com" target="_blank">Seventy Upgrades</a>.
			</p>
			<p>
				This feature imports gear, race, and (optionally) talents. It does NOT import buffs, debuffs, consumes, rotation, or custom stats.
			</p>
			<p>
				To import, paste the output from the site's export option below and click, 'Import'.
			</p>
		`;
    }
    onImport(data) {
        const importJson = JSON.parse(data);
        // Parse all the settings.
        const charClass = nameToClass(importJson?.character?.gameClass || '');
        if (charClass == Class.ClassUnknown) {
            throw new Error('Could not parse Class!');
        }
        const race = nameToRace(importJson?.character?.race || '');
        if (race == Race.RaceUnknown) {
            throw new Error('Could not parse Race!');
        }
        let talentsStr = '';
        if (importJson?.talents.length > 0) {
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
        this.finishIndividualImport(this.simUI, charClass, race, equipmentSpec, talentsStr);
    }
}
class IndividualAddonImporter extends Importer {
    constructor(parent, simUI) {
        super(parent, 'Addon Import');
        this.simUI = simUI;
        this.descriptionElem.innerHTML = `
			<p>
				Import settings from the <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank">WoWSims Importer In-Game Addon</a>.
			</p>
			<p>
				This feature imports gear, race, and talents. It does NOT import buffs, debuffs, consumes, rotation, or custom stats.
			</p>
			<p>
				To import, paste the output from the addon below and click, 'Import'.
			</p>
		`;
    }
    onImport(data) {
        const importJson = JSON.parse(data);
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
        this.finishIndividualImport(this.simUI, charClass, race, equipmentSpec, talentsStr);
    }
}
