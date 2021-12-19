import { BonusStatsPicker } from '/tbc/core/components/bonus_stats_picker.js';
import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { CharacterStats } from '/tbc/core/components/character_stats.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { DetailedResults } from '/tbc/core/components/detailed_results.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EncounterPicker } from '/tbc/core/components/encounter_picker.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { GearPicker } from '/tbc/core/components/gear_picker.js';
import { IconPicker } from '/tbc/core/components/icon_picker.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { LogRunner } from '/tbc/core/components/log_runner.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { SavedDataManager } from '/tbc/core/components/saved_data_manager.js';
import { SimUI } from './sim_ui.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { addRaidSimAction } from '/tbc/core/components/raid_sim_action.js';
import { addStatWeightsAction } from '/tbc/core/components/stat_weights_action.js';
import { equalsOrBothNull } from '/tbc/core/utils.js';
import { newTalentsPicker } from '/tbc/core/talents/factory.js';
import { raceNames } from '/tbc/core/proto_utils/names.js';
import { specNames } from '/tbc/core/proto_utils/utils.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
const SAVED_GEAR_STORAGE_KEY = '__savedGear__';
const SAVED_ROTATION_STORAGE_KEY = '__savedRotation__';
const SAVED_SETTINGS_STORAGE_KEY = '__savedSettings__';
const SAVED_TALENTS_STORAGE_KEY = '__savedTalents__';
;
class IndividualSimIconPicker extends IconPicker {
    constructor(parent, modObj, input, simUI) {
        super(parent, modObj, input);
        if (input.exclusivityTags) {
            simUI.registerExclusiveEffect({
                tags: input.exclusivityTags,
                changedEvent: this.changeEmitter,
                isActive: () => Boolean(this.getInputValue()),
                deactivate: () => this.setInputValue(0),
            });
        }
    }
}
// Extended shared UI for all individual player sims.
export class IndividualSimUI extends SimUI {
    constructor(parentElem, player, config) {
        let title = 'TBC ' + specNames[player.spec] + ' Sim';
        if (config.releaseStatus == 'Alpha') {
            title += ' Alpha';
        }
        else if (config.releaseStatus == 'Beta') {
            title += ' Beta';
        }
        super(parentElem, player.sim, {
            title: title,
            knownIssues: config.knownIssues,
        });
        this.rootElem.classList.add('individual-sim-ui', config.cssClass);
        this.player = player;
        this.individualConfig = config;
        this.isWithinRaidSim = this.rootElem.closest('.within-raid-sim') != null;
        this.raidSimResultsManager = null;
        this.exclusivityMap = {
            'Battle Elixir': [],
            'Drums': [],
            'Food': [],
            'Alchohol': [],
            'Guardian Elixir': [],
            'Potion': [],
            'Rune': [],
            'Weapon Imbue': [],
        };
        if (!this.isWithinRaidSim) {
            // This needs to go before all the UI components so that gear loading is the
            // first callback invoked from waitForInit().
            this.sim.waitForInit().then(() => {
                this.loadSettings();
            });
        }
        this.player.setEpWeights(this.individualConfig.defaults.epWeights);
        this.addSidebarComponents();
        this.addTopbarComponents();
        this.addGearTab();
        this.addSettingsTab();
        this.addTalentsTab();
        if (!this.isWithinRaidSim) {
            this.addDetailedResultsTab();
            this.addLogTab();
        }
    }
    loadSettings() {
        let loadedSettings = false;
        let hash = window.location.hash;
        if (hash.length > 1) {
            // Remove leading '#'
            hash = hash.substring(1);
            try {
                let jsonData;
                if (new URLSearchParams(window.location.search).has('uncompressed')) {
                    const jsonStr = atob(hash);
                    jsonData = JSON.parse(jsonStr);
                }
                else {
                    const binary = atob(hash);
                    const bytes = new Uint8Array(binary.length);
                    for (let i = 0; i < bytes.length; i++) {
                        bytes[i] = binary.charCodeAt(i);
                    }
                    const jsonStr = pako.inflate(bytes, { to: 'string' });
                    jsonData = JSON.parse(jsonStr);
                }
                this.sim.fromJson(jsonData, this.player.spec);
                loadedSettings = true;
            }
            catch (e) {
                console.warn('Failed to parse settings from window hash: ' + e);
            }
        }
        window.location.hash = '';
        const savedSettings = window.localStorage.getItem(this.getSettingsStorageKey());
        if (!loadedSettings && savedSettings != null) {
            try {
                this.sim.fromJson(JSON.parse(savedSettings), this.player.spec);
                loadedSettings = true;
            }
            catch (e) {
                console.warn('Failed to parse saved settings: ' + e);
            }
        }
        if (!loadedSettings) {
            this.applyDefaults();
        }
        this.player.setName('Player');
        // This needs to go last so it doesn't re-store things as they are initialized.
        this.changeEmitter.on(() => {
            const jsonStr = JSON.stringify(this.sim.toJson());
            window.localStorage.setItem(this.getSettingsStorageKey(), jsonStr);
        });
    }
    addSidebarComponents() {
        this.raidSimResultsManager = addRaidSimAction(this);
        addStatWeightsAction(this, this.individualConfig.epStats, this.individualConfig.epReferenceStat);
        const characterStats = new CharacterStats(this.rootElem.getElementsByClassName('sim-sidebar-footer')[0], this.player, this.individualConfig.displayStats, this.individualConfig.modifyDisplayStats);
    }
    addTopbarComponents() {
        Array.from(document.getElementsByClassName('share-link')).forEach(element => {
            tippy(element, {
                'content': 'Shareable link',
                'allowHTML': true,
            });
            element.addEventListener('click', event => {
                const jsonStr = JSON.stringify(this.sim.toJson());
                const val = pako.deflate(jsonStr, { to: 'string' });
                const encoded = btoa(String.fromCharCode(...val));
                const linkUrl = new URL(window.location.href);
                linkUrl.hash = encoded;
                if (navigator.clipboard == undefined) {
                    alert(linkUrl.toString());
                }
                else {
                    navigator.clipboard.writeText(linkUrl.toString());
                    alert('Current settings copied to clipboard!');
                }
            });
        });
    }
    addGearTab() {
        this.addTab('Gear', 'gear-tab', `
			<div class="gear-tab-columns">
				<div class="left-gear-panel">
					<div class="gear-picker">
					</div>
				</div>
				<div class="right-gear-panel">
					<div class="bonus-stats-picker">
					</div>
					<div class="saved-gear-manager">
					</div>
				</div>
			</div>
		`);
        const gearPicker = new GearPicker(this.rootElem.getElementsByClassName('gear-picker')[0], this.player);
        const bonusStatsPicker = new BonusStatsPicker(this.rootElem.getElementsByClassName('bonus-stats-picker')[0], this.player, this.individualConfig.epStats);
        const savedGearManager = new SavedDataManager(this.rootElem.getElementsByClassName('saved-gear-manager')[0], this.player, {
            label: 'Gear',
            storageKey: this.getSavedGearStorageKey(),
            getData: (player) => {
                return {
                    gear: player.getGear(),
                    bonusStats: player.getBonusStats(),
                };
            },
            setData: (player, newGearAndStats) => {
                player.setGear(newGearAndStats.gear);
                if (newGearAndStats.bonusStats) {
                    player.setBonusStats(newGearAndStats.bonusStats);
                }
            },
            changeEmitters: [this.player.changeEmitter],
            equals: (a, b) => a.gear.equals(b.gear) && equalsOrBothNull(a.bonusStats, b.bonusStats, (a, b) => a.equals(b)),
            toJson: (a) => {
                return {
                    gear: EquipmentSpec.toJson(a.gear.asSpec()),
                    bonusStats: a.bonusStats?.toJson(),
                };
            },
            fromJson: (obj) => {
                return {
                    gear: this.sim.lookupEquipmentSpec(EquipmentSpec.fromJson(obj['gear'])),
                    bonusStats: Stats.fromJson(obj['bonusStats']),
                };
            },
        });
        this.sim.waitForInit().then(() => {
            savedGearManager.loadUserData();
            this.individualConfig.presets.gear.forEach(presetGear => {
                savedGearManager.addSavedData({
                    name: presetGear.name,
                    tooltip: presetGear.tooltip,
                    isPreset: true,
                    data: {
                        gear: this.sim.lookupEquipmentSpec(presetGear.gear),
                        bonusStats: new Stats(),
                    },
                    enableWhen: presetGear.enableWhen,
                });
            });
        });
    }
    addSettingsTab() {
        this.addTab('Settings', 'settings-tab', `
			<div class="settings-inputs">
				<div class="settings-section-container">
					<section class="settings-section encounter-section within-raid-sim-hide">
						<label>Encounter</label>
					</section>
					<section class="settings-section race-section">
						<label>Race</label>
					</section>
					<section class="settings-section rotation-section">
						<label>Rotation</label>
					</section>
				</div>
				<div class="settings-section-container">
					<section class="settings-section self-buffs-section">
						<label>Self Buffs</label>
					</section>
				</div>
				<div class="settings-section-container within-raid-sim-hide">
					<section class="settings-section buffs-section">
						<label>Other Buffs</label>
					</section>
				</div>
				<div class="settings-section-container">
					<section class="settings-section consumes-section">
						<label>Consumes</label>
					</section>
				</div>
				<div class="settings-section-container within-raid-sim-hide">
					<section class="settings-section debuffs-section">
						<label>Debuffs</label>
					</section>
				</div>
				<div class="settings-section-container">
					<section class="settings-section other-settings-section">
						<label>Other</label>
					</section>
				</div>
			</div>
			<div class="settings-bottom-bar">
				<div class="saved-encounter-manager within-raid-sim-hide">
				</div>
				<div class="saved-rotation-manager">
				</div>
				<div class="saved-settings-manager within-raid-sim-hide">
				</div>
			</div>
		`);
        const settingsTab = this.rootElem.getElementsByClassName('settings-inputs')[0];
        const configureIconSection = (sectionElem, iconPickers, tooltip) => {
            if (tooltip) {
                tippy(sectionElem, {
                    'content': tooltip,
                    'allowHTML': true,
                });
            }
            if (iconPickers.length == 0) {
                sectionElem.style.display = 'none';
            }
        };
        const selfBuffsSection = this.rootElem.getElementsByClassName('self-buffs-section')[0];
        configureIconSection(selfBuffsSection, this.individualConfig.selfBuffInputs.map(iconInput => new IndividualSimIconPicker(selfBuffsSection, this.player, iconInput, this)), Tooltips.SELF_BUFFS_SECTION);
        const buffsSection = this.rootElem.getElementsByClassName('buffs-section')[0];
        configureIconSection(buffsSection, [
            this.individualConfig.raidBuffInputs.map(iconInput => new IndividualSimIconPicker(buffsSection, this.sim.raid, iconInput, this)),
            this.individualConfig.playerBuffInputs.map(iconInput => new IndividualSimIconPicker(buffsSection, this.player, iconInput, this)),
            this.individualConfig.partyBuffInputs.map(iconInput => new IndividualSimIconPicker(buffsSection, this.player.getParty(), iconInput, this)),
        ].flat(), Tooltips.OTHER_BUFFS_SECTION);
        const debuffsSection = this.rootElem.getElementsByClassName('debuffs-section')[0];
        configureIconSection(debuffsSection, this.individualConfig.debuffInputs.map(iconInput => new IndividualSimIconPicker(debuffsSection, this.sim.encounter.primaryTarget, iconInput, this)));
        const consumesSection = this.rootElem.getElementsByClassName('consumes-section')[0];
        configureIconSection(consumesSection, this.individualConfig.consumeInputs.map(iconInput => new IndividualSimIconPicker(consumesSection, this.player, iconInput, this)));
        const configureInputSection = (sectionElem, sectionConfig) => {
            if (sectionConfig.tooltip) {
                tippy(sectionElem, {
                    'content': sectionConfig.tooltip,
                    'allowHTML': true,
                });
            }
            sectionConfig.inputs.forEach(inputConfig => {
                if (inputConfig.type == 'number') {
                    const picker = new NumberPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
                }
                else if (inputConfig.type == 'boolean') {
                    const picker = new BooleanPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
                }
                else if (inputConfig.type == 'enum') {
                    const picker = new EnumPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
                }
            });
        };
        configureInputSection(this.rootElem.getElementsByClassName('rotation-section')[0], this.individualConfig.rotationInputs);
        if (this.individualConfig.otherInputs?.inputs.length) {
            configureInputSection(this.rootElem.getElementsByClassName('other-settings-section')[0], this.individualConfig.otherInputs);
        }
        const savedRotationManager = new SavedDataManager(this.rootElem.getElementsByClassName('saved-rotation-manager')[0], this.player, {
            label: 'Rotation',
            storageKey: this.getSavedRotationStorageKey(),
            getData: (player) => player.getRotation(),
            setData: (player, newRotation) => player.setRotation(newRotation),
            changeEmitters: [this.player.rotationChangeEmitter],
            equals: (a, b) => this.player.specTypeFunctions.rotationEquals(a, b),
            toJson: (a) => this.player.specTypeFunctions.rotationToJson(a),
            fromJson: (obj) => this.player.specTypeFunctions.rotationFromJson(obj),
        });
        const makeInputSection = (sectionName, sectionConfig) => {
            const sectionCssPrefix = sectionName.replace(/\s+/g, '');
            const sectionElem = document.createElement('section');
            sectionElem.classList.add('settings-section', sectionCssPrefix + '-section');
            sectionElem.innerHTML = `<label>${sectionName}</label>`;
            settingsTab.appendChild(sectionElem);
            configureInputSection(sectionElem, sectionConfig);
        };
        for (const [sectionName, sectionConfig] of Object.entries(this.individualConfig.additionalSections || {})) {
            makeInputSection(sectionName, sectionConfig);
        }
        ;
        const races = specToEligibleRaces[this.player.spec];
        const racePicker = new EnumPicker(this.rootElem.getElementsByClassName('race-section')[0], this.player, {
            values: races.map(race => {
                return {
                    name: raceNames[race],
                    value: race,
                };
            }),
            changedEvent: sim => sim.raceChangeEmitter,
            getValue: sim => sim.getRace(),
            setValue: (sim, newValue) => sim.setRace(newValue),
        });
        const encounterSectionElem = settingsTab.getElementsByClassName('encounter-section')[0];
        new EncounterPicker(encounterSectionElem, this.sim.encounter, this.individualConfig.encounterPicker);
        const savedEncounterManager = new SavedDataManager(this.rootElem.getElementsByClassName('saved-encounter-manager')[0], this.sim.encounter, {
            label: 'Encounter',
            storageKey: this.getSavedEncounterStorageKey(),
            getData: (encounter) => encounter.toProto(),
            setData: (encounter, newEncounter) => encounter.fromProto(newEncounter),
            changeEmitters: [this.sim.encounter.changeEmitter],
            equals: (a, b) => EncounterProto.equals(a, b),
            toJson: (a) => EncounterProto.toJson(a),
            fromJson: (obj) => EncounterProto.fromJson(obj),
        });
        // Init Muuri layout only when settings tab is clicked, because it needs the elements
        // to be shown so it can calculate sizes.
        let muuriInit = false;
        this.rootElem.getElementsByClassName('settings-tab-tab')[0].addEventListener('click', event => {
            if (muuriInit) {
                return;
            }
            muuriInit = true;
            setTimeout(() => {
                new Muuri('.settings-inputs');
            }, 200); // Magic amount of time before Muuri init seems to work
        });
        const savedSettingsManager = new SavedDataManager(this.rootElem.getElementsByClassName('saved-settings-manager')[0], this, {
            label: 'Settings',
            storageKey: this.getSavedSettingsStorageKey(),
            getData: (simUI) => {
                return {
                    raidBuffs: simUI.sim.raid.getBuffs(),
                    partyBuffs: simUI.player.getParty().getBuffs(),
                    individualBuffs: simUI.player.getBuffs(),
                    consumes: simUI.player.getConsumes(),
                    race: simUI.player.getRace(),
                };
            },
            setData: (simUI, newSettings) => {
                simUI.sim.raid.setBuffs(newSettings.raidBuffs);
                simUI.player.getParty().setBuffs(newSettings.partyBuffs);
                simUI.player.setBuffs(newSettings.individualBuffs);
                simUI.player.setConsumes(newSettings.consumes);
                simUI.player.setRace(newSettings.race);
            },
            changeEmitters: [
                this.sim.raid.buffsChangeEmitter,
                this.player.getParty().buffsChangeEmitter,
                this.player.buffsChangeEmitter,
                this.player.consumesChangeEmitter,
                this.player.raceChangeEmitter,
            ],
            equals: (a, b) => RaidBuffs.equals(a.raidBuffs, b.raidBuffs)
                && PartyBuffs.equals(a.partyBuffs, b.partyBuffs)
                && IndividualBuffs.equals(a.individualBuffs, b.individualBuffs)
                && Consumes.equals(a.consumes, b.consumes)
                && a.race == b.race,
            toJson: (a) => {
                return {
                    raidBuffs: RaidBuffs.toJson(a.raidBuffs),
                    partyBuffs: PartyBuffs.toJson(a.partyBuffs),
                    individualBuffs: IndividualBuffs.toJson(a.individualBuffs),
                    consumes: Consumes.toJson(a.consumes),
                    race: a.race,
                };
            },
            fromJson: (obj) => {
                return {
                    raidBuffs: RaidBuffs.fromJson(obj['raidBuffs']),
                    partyBuffs: PartyBuffs.fromJson(obj['partyBuffs']),
                    individualBuffs: IndividualBuffs.fromJson(obj['individualBuffs']),
                    consumes: Consumes.fromJson(obj['consumes']),
                    race: Number(obj['race']),
                };
            },
        });
        this.sim.waitForInit().then(() => {
            savedEncounterManager.loadUserData();
            savedSettingsManager.loadUserData();
        });
    }
    addTalentsTab() {
        this.addTab('Talents', 'talents-tab', `
			<div class="talents-picker">
			</div>
			<div class="saved-talents-manager">
			</div>
		`);
        const talentsPicker = newTalentsPicker(this.player.spec, this.rootElem.getElementsByClassName('talents-picker')[0], this.player);
        const savedTalentsManager = new SavedDataManager(this.rootElem.getElementsByClassName('saved-talents-manager')[0], this.player, {
            label: 'Talents',
            storageKey: this.getSavedTalentsStorageKey(),
            getData: (player) => player.getTalentsString(),
            setData: (player, newTalentsString) => player.setTalentsString(newTalentsString),
            changeEmitters: [this.player.talentsStringChangeEmitter],
            equals: (a, b) => a == b,
            toJson: (a) => a,
            fromJson: (obj) => obj,
        });
        // Add a url parameter to help people trapped in the wrong talents   ;)
        const freezeTalents = this.individualConfig.freezeTalents && !(new URLSearchParams(window.location.search).has('unlockTalents'));
        if (freezeTalents) {
            savedTalentsManager.freeze();
            talentsPicker.freeze();
        }
        this.sim.waitForInit().then(() => {
            savedTalentsManager.loadUserData();
            this.individualConfig.presets.talents.forEach(config => {
                config.isPreset = true;
                savedTalentsManager.addSavedData(config);
            });
        });
    }
    addDetailedResultsTab() {
        this.addTab('Detailed Results', 'detailed-results-tab', `
			<div class="detailed-results">
			</div>
		`);
        const detailedResults = new DetailedResults(this.rootElem.getElementsByClassName('detailed-results')[0], this, this.raidSimResultsManager);
    }
    addLogTab() {
        this.addTab('Log', 'log-tab', `
			<div class="log-runner">
			</div>
		`);
        const logRunner = new LogRunner(this.rootElem.getElementsByClassName('log-runner')[0], this);
    }
    applyDefaults() {
        this.player.setGear(this.sim.lookupEquipmentSpec(this.individualConfig.defaults.gear));
        this.player.setConsumes(this.individualConfig.defaults.consumes);
        this.player.setRotation(this.individualConfig.defaults.rotation);
        this.player.setTalentsString(this.individualConfig.defaults.talents);
        this.player.setSpecOptions(this.individualConfig.defaults.specOptions);
        this.player.setBuffs(this.individualConfig.defaults.individualBuffs);
        this.player.getParty().setBuffs(this.individualConfig.defaults.partyBuffs);
        this.player.getRaid().setBuffs(this.individualConfig.defaults.raidBuffs);
        this.sim.encounter.primaryTarget.setDebuffs(this.individualConfig.defaults.debuffs);
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
}
