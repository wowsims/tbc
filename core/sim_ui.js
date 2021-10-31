import { makeIndividualSimRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';
import { Player } from './player.js';
import { Sim } from './sim.js';
import { Target } from './target.js';
import { TypedEvent } from './typed_event.js';
const CURRENT_SETTINGS_STORAGE_KEY = '__currentSettings__';
const SAVED_GEAR_STORAGE_KEY = '__savedGear__';
const SAVED_ENCOUNTER_STORAGE_KEY = '__savedEncounter__';
const SAVED_ROTATION_STORAGE_KEY = '__savedRotation__';
const SAVED_SETTINGS_STORAGE_KEY = '__savedSettings__';
const SAVED_TALENTS_STORAGE_KEY = '__savedTalents__';
// Core UI module.
export class SimUI {
    constructor(parentElem, config) {
        // Emits when anything from sim, player, or target changes.
        this.changeEmitter = new TypedEvent();
        this.parentElem = parentElem;
        this.sim = new Sim(config.sim);
        this.player = new Player(config.player, this.sim);
        this.target = new Target(config.target, this.sim);
        this.simUiConfig = config;
        [
            this.sim.changeEmitter,
            this.player.changeEmitter,
            this.target.changeEmitter,
        ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
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
    // Returns JSON representing all the current values.
    toJson() {
        return {
            'sim': this.sim.toJson(),
            'player': this.player.toJson(),
            'target': this.target.toJson(),
        };
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(obj) {
        if (obj['sim']) {
            this.sim.fromJson(obj['sim']);
        }
        if (obj['player']) {
            this.player.fromJson(obj['player']);
        }
        if (obj['target']) {
            this.target.fromJson(obj['target']);
        }
    }
    async init() {
        await this.sim.init(this.player.spec);
        let loadedSettings = false;
        let hash = window.location.hash;
        if (hash.length > 1) {
            // Remove leading '#'
            hash = hash.substring(1);
            try {
                const jsonStr = atob(hash);
                this.fromJson(JSON.parse(jsonStr));
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
                this.fromJson(JSON.parse(savedSettings));
                loadedSettings = true;
            }
            catch (e) {
                console.warn('Failed to parse saved settings: ' + e);
            }
        }
        if (!loadedSettings) {
            this.player.setGear(this.sim.lookupEquipmentSpec(this.player.defaultGear));
        }
        this.changeEmitter.on(() => {
            const jsonStr = JSON.stringify(this.toJson());
            window.localStorage.setItem(this.getStorageKey(CURRENT_SETTINGS_STORAGE_KEY), jsonStr);
        });
        Array.from(document.getElementsByClassName('share-link')).forEach(element => {
            tippy(element, {
                'content': 'Shareable link',
                'allowHTML': true,
            });
            element.addEventListener('click', event => {
                const linkUrl = new URL(window.location.href);
                const jsonStr = JSON.stringify(this.toJson());
                const encoded = btoa(jsonStr);
                linkUrl.hash = encoded;
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
        return specToLocalStorageKey[this.player.spec] + keyPart;
    }
    makeCurrentIndividualSimRequest(iterations, debug) {
        const encounter = this.sim.getEncounter();
        const numTargets = Math.max(1, this.sim.getNumTargets());
        for (let i = 0; i < numTargets; i++) {
            encounter.targets.push(this.target.toProto());
        }
        return makeIndividualSimRequest(this.sim.getBuffs(), this.player.getConsumes(), this.player.getCustomStats(), encounter, this.player.getGear(), this.player.getRace(), this.player.getRotation(), this.player.getTalents(), this.player.getSpecOptions(), iterations, debug);
    }
}
