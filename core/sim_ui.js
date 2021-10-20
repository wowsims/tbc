import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';
import { Sim } from './sim.js';
const CURRENT_SETTINGS_STORAGE_KEY = '__currentSettings__';
const SAVED_GEAR_STORAGE_KEY = '__savedGear__';
const SAVED_ENCOUNTER_STORAGE_KEY = '__savedEncounter__';
const SAVED_ROTATION_STORAGE_KEY = '__savedRotation__';
const SAVED_SETTINGS_STORAGE_KEY = '__savedSettings__';
const SAVED_TALENTS_STORAGE_KEY = '__savedTalents__';
// Core UI module.
export class SimUI {
    constructor(parentElem, config) {
        this.parentElem = parentElem;
        this.sim = new Sim(config);
        this.simUiConfig = config;
        this.exclusivityMap = {
            'Battle Elixir': [],
            'Drums': [],
            'Food': [],
            'Guardian Elixir': [],
            'Potion': [],
            'Rune': [],
            'Weapon Imbue': [],
        };
        Array.from(document.getElementsByClassName('known-issues')).forEach(element => {
            if (this.simUiConfig.knownIssues?.length) {
                element.style.display = 'initial';
            }
            else {
                return;
            }
            tippy(element, {
                'content': `
				<ul class="known-issues-tooltip">
					${this.simUiConfig.knownIssues.map(issue => '<li>' + issue + '</li>').join('')}
				</ul>
				`,
                'allowHTML': true,
            });
        });
    }
    async init() {
        await this.sim.init();
        let loadedSettings = false;
        let hash = window.location.hash;
        if (hash.length > 1) {
            // Remove leading '#'
            hash = hash.substring(1);
            try {
                const simJsonStr = atob(hash);
                this.sim.fromJson(JSON.parse(simJsonStr));
                loadedSettings = true;
            }
            catch (e) {
                console.warn('Failed to parse settings from window hash: ' + e);
            }
        }
        window.location.hash = '';
        const savedSettings = window.localStorage.getItem(this.getStorageKey(CURRENT_SETTINGS_STORAGE_KEY));
        if (!loadedSettings && savedSettings != null) {
            try {
                this.sim.fromJson(JSON.parse(savedSettings));
                loadedSettings = true;
            }
            catch (e) {
                console.warn('Failed to parse saved settings: ' + e);
            }
        }
        if (!loadedSettings) {
            this.sim.setGear(this.sim.lookupEquipmentSpec(this.simUiConfig.defaults.gear));
        }
        this.sim.changeEmitter.on(() => {
            const simJsonStr = JSON.stringify(this.sim.toJson());
            window.localStorage.setItem(this.getStorageKey(CURRENT_SETTINGS_STORAGE_KEY), simJsonStr);
        });
        Array.from(document.getElementsByClassName('share-link')).forEach(element => {
            tippy(element, {
                'content': 'Shareable link',
                'allowHTML': true,
            });
            element.addEventListener('click', event => {
                const linkUrl = new URL(window.location.href);
                const simJsonStr = JSON.stringify(this.sim.toJson());
                const simEncoded = btoa(simJsonStr);
                linkUrl.hash = simEncoded;
                navigator.clipboard.writeText(linkUrl.toString());
                alert('Current settings copied to clipboard!');
            });
        });
    }
    registerExclusiveEffect(effect) {
        effect.tags.forEach(tag => {
            this.exclusivityMap[tag].push(effect);
            effect.changedEvent.on(() => {
                if (!effect.isActive())
                    return;
                this.exclusivityMap[tag].forEach(otherEffect => {
                    if (otherEffect == effect || !otherEffect.isActive())
                        return;
                    otherEffect.deactivate();
                });
            });
        });
    }
    getSavedGearStorageKey() {
        return this.getStorageKey(SAVED_GEAR_STORAGE_KEY);
    }
    getSavedEncounterStorageKey() {
        return this.getStorageKey(SAVED_ENCOUNTER_STORAGE_KEY);
    }
    getSavedRotationStorageKey() {
        return this.getStorageKey(SAVED_ROTATION_STORAGE_KEY);
    }
    getSavedSettingsStorageKey() {
        return this.getStorageKey(SAVED_SETTINGS_STORAGE_KEY);
    }
    getSavedTalentsStorageKey() {
        return this.getStorageKey(SAVED_TALENTS_STORAGE_KEY);
    }
    // Returns the actual key to use for local storage, based on the given key part and the site context.
    getStorageKey(keyPart) {
        // Local storage is shared by all sites under the same domain, so we need to use
        // different keys for each spec site.
        return specToLocalStorageKey[this.simUiConfig.spec] + keyPart;
    }
}
