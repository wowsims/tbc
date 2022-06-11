import { BooleanPicker } from "/tbc/core/components/boolean_picker.js";
import { DetailedResults } from "/tbc/core/components/detailed_results.js";
import { EncounterPicker } from "/tbc/core/components/encounter_picker.js";
import { LogRunner } from "/tbc/core/components/log_runner.js";
import { addRaidSimAction } from "/tbc/core/components/raid_sim_action.js";
import { SavedDataManager } from "/tbc/core/components/saved_data_manager.js";
import { SettingsMenu } from "/tbc/core/components/settings_menu.js";
import * as Tooltips from "/tbc/core/constants/tooltips.js";
import { Raid as RaidProto } from "/tbc/core/proto/api.js";
import { Class, Encounter as EncounterProto, Stat, TristateEffect } from "/tbc/core/proto/common.js";
import { Blessings } from "/tbc/core/proto/paladin.js";
import { BlessingsAssignments, RaidSimSettings, SavedEncounter, SavedRaid } from "/tbc/core/proto/ui.js";
import { playerToSpec } from "/tbc/core/proto_utils/utils.js";
import { Sim } from "/tbc/core/sim.js";
import { SimUI } from "/tbc/core/sim_ui.js";
import { TypedEvent } from "/tbc/core/typed_event.js";
import { AssignmentsPicker } from "./assignments_picker.js";
import { BlessingsPicker } from "./blessings_picker.js";
import { newRaidExporters, newRaidImporters } from "./import_export.js";
import { implementedSpecs } from "./presets.js";
import { RaidPicker } from "./raid_picker.js";
import { TanksPicker } from "./tanks_picker.js";
const extraKnownIssues = [
    'We\'re still missing implementations for many specs. If you\'d like to help us out, check out our <a href="https://github.com/wowsims/tbc">Github project</a> or <a href="https://discord.gg/jJMPr9JWwx">join our discord</a>!',
];
export class RaidSimUI extends SimUI {
    constructor(parentElem, config) {
        super(parentElem, new Sim(), {
            spec: null,
            knownIssues: (config.knownIssues || []).concat(extraKnownIssues),
        });
        this.raidSimResultsManager = null;
        this.raidPicker = null;
        this.blessingsPicker = null;
        // Emits when the raid comp changes. Includes changes to buff bots.
        this.compChangeEmitter = new TypedEvent();
        this.changeEmitter = new TypedEvent();
        this.referenceChangeEmitter = new TypedEvent();
        this.rootElem.classList.add('raid-sim-ui');
        this.config = config;
        this.settingsMuuri = null;
        this.sim.raid.compChangeEmitter.on(eventID => this.compChangeEmitter.emit(eventID));
        this.sim.setModifyRaidProto(raidProto => this.modifyRaidProto(raidProto));
        [
            this.compChangeEmitter,
            this.sim.changeEmitter,
        ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
        this.sim.waitForInit().then(() => this.loadSettings());
        this.addSidebarComponents();
        this.addTopbarComponents();
        this.addRaidTab();
        this.addSettingsTab();
        this.addDetailedResultsTab();
        this.addLogTab();
        this.changeEmitter.on(() => this.recomputeSettingsLayout());
    }
    loadSettings() {
        const initEventID = TypedEvent.nextEventID();
        TypedEvent.freezeAllAndDo(() => {
            let loadedSettings = false;
            const savedSettings = window.localStorage.getItem(this.getSettingsStorageKey());
            if (savedSettings != null) {
                try {
                    const settings = RaidSimSettings.fromJsonString(savedSettings);
                    this.fromProto(initEventID, settings);
                    loadedSettings = true;
                }
                catch (e) {
                    console.warn('Failed to parse saved settings: ' + e);
                }
            }
            if (!loadedSettings) {
                this.applyDefaults(initEventID);
            }
            // This needs to go last so it doesn't re-store things as they are initialized.
            this.changeEmitter.on(eventID => {
                const jsonStr = RaidSimSettings.toJsonString(this.toProto());
                window.localStorage.setItem(this.getSettingsStorageKey(), jsonStr);
            });
        });
    }
    addSidebarComponents() {
        this.raidSimResultsManager = addRaidSimAction(this);
        this.raidSimResultsManager.changeEmitter.on(eventID => this.referenceChangeEmitter.emit(eventID));
    }
    addTopbarComponents() {
        this.addToolbarItem(newRaidImporters(this));
        this.addToolbarItem(newRaidExporters(this));
        const settingsMenu = document.createElement('span');
        settingsMenu.classList.add('fas', 'fa-cog');
        tippy(settingsMenu, {
            'content': 'Settings',
            'allowHTML': true,
        });
        settingsMenu.addEventListener('click', event => {
            new SettingsMenu(this.rootElem, this);
        });
        this.addToolbarItem(settingsMenu);
    }
    addRaidTab() {
        this.addTab('RAID', 'raid-tab', `
			<div class="raid-picker">
			</div>
			<div class="saved-raids-div">
				<div class="saved-raids-manager">
				</div>
			</div>
		`);
        this.raidPicker = new RaidPicker(this.rootElem.getElementsByClassName('raid-picker')[0], this);
        const savedRaidManager = new SavedDataManager(this.rootElem.getElementsByClassName('saved-raids-manager')[0], this, {
            label: 'Raid',
            storageKey: this.getSavedRaidStorageKey(),
            getData: (raidSimUI) => SavedRaid.create({
                raid: this.sim.raid.toProto(),
                buffBots: this.getBuffBots().map(b => b.toProto()),
                blessings: this.blessingsPicker.getAssignments(),
            }),
            setData: (eventID, raidSimUI, newRaid) => {
                TypedEvent.freezeAllAndDo(() => {
                    this.sim.raid.fromProto(eventID, newRaid.raid || RaidProto.create());
                    this.raidPicker.setBuffBots(eventID, newRaid.buffBots);
                    this.blessingsPicker.setAssignments(eventID, newRaid.blessings || BlessingsAssignments.create());
                });
            },
            changeEmitters: [this.changeEmitter],
            equals: (a, b) => {
                return SavedRaid.equals(a, b);
            },
            toJson: (a) => SavedRaid.toJson(a),
            fromJson: (obj) => SavedRaid.fromJson(obj),
        });
        this.sim.waitForInit().then(() => {
            savedRaidManager.loadUserData();
        });
    }
    addSettingsTab() {
        this.addTab('SETTINGS', 'raid-settings-tab', `
			<div class="raid-settings-sections">
				<div class="settings-section-container raid-settings-section-container">
					<fieldset class="settings-section raid-encounter-section">
						<legend>Encounter</legend>
					</fieldset>
				</div>
				<div class="settings-section-container blessings-section-container">
					<fieldset class="settings-section blessings-section">
						<legend>Blessings</legend>
					</fieldset>
				</div>
				<div class="settings-section-container assignments-section-container">
				</div>
				<div class="settings-section-container tanks-section-container">
				</div>
				<div class="settings-section-container raid-settings-section-container">
					<fieldset class="settings-section other-options-section">
						<legend>Other Options</legend>
					</fieldset>
				</div>
			</div>
			<div class="settings-bottom-bar">
				<div class="saved-encounter-manager">
				</div>
			</div>
		`);
        const encounterSectionElem = this.rootElem.getElementsByClassName('raid-encounter-section')[0];
        new EncounterPicker(encounterSectionElem, this.sim.encounter, {
            simpleTargetStats: [
                Stat.StatArmor,
            ],
            showExecuteProportion: true,
        });
        const savedEncounterManager = new SavedDataManager(this.rootElem.getElementsByClassName('saved-encounter-manager')[0], this.sim.encounter, {
            label: 'Encounter',
            storageKey: this.getSavedEncounterStorageKey(),
            getData: (encounter) => SavedEncounter.create({ encounter: encounter.toProto() }),
            setData: (eventID, encounter, newEncounter) => encounter.fromProto(eventID, newEncounter.encounter),
            changeEmitters: [this.sim.encounter.changeEmitter],
            equals: (a, b) => SavedEncounter.equals(a, b),
            toJson: (a) => SavedEncounter.toJson(a),
            fromJson: (obj) => SavedEncounter.fromJson(obj),
        });
        this.sim.waitForInit().then(() => {
            savedEncounterManager.loadUserData();
        });
        const blessingsSection = this.rootElem.getElementsByClassName('blessings-section')[0];
        this.blessingsPicker = new BlessingsPicker(blessingsSection, this);
        this.blessingsPicker.changeEmitter.on(eventID => this.changeEmitter.emit(eventID));
        tippy(blessingsSection, {
            content: Tooltips.BLESSINGS_SECTION,
            allowHTML: true,
            placement: 'left',
        });
        const assignmentsPicker = new AssignmentsPicker(this.rootElem.getElementsByClassName('assignments-section-container')[0], this);
        const tanksPicker = new TanksPicker(this.rootElem.getElementsByClassName('tanks-section-container')[0], this);
        const otherOptionsSectionElem = this.rootElem.getElementsByClassName('other-options-section')[0];
        new BooleanPicker(otherOptionsSectionElem, this.sim.raid, {
            label: 'Stagger Stormstrikes',
            labelTooltip: 'When there are multiple Enhancement Shaman in the raid, causes them to coordinate their Stormstrike casts for optimal SS charge usage.',
            changedEvent: (raid) => raid.staggerStormstrikesChangeEmitter,
            getValue: (raid) => raid.getStaggerStormstrikes(),
            setValue: (eventID, raid, newValue) => {
                raid.setStaggerStormstrikes(eventID, newValue);
            },
        });
        // Init Muuri layout only when settings tab is clicked, because it needs the elements
        // to be shown so it can calculate sizes.
        this.rootElem.getElementsByClassName('raid-settings-tab-tab')[0].addEventListener('click', event => {
            if (this.settingsMuuri == null) {
                setTimeout(() => {
                    this.settingsMuuri = new Muuri('.raid-settings-sections');
                }, 200); // Magic amount of time before Muuri init seems to work
            }
            setTimeout(() => {
                this.recomputeSettingsLayout();
            }, 200);
        });
    }
    addDetailedResultsTab() {
        this.addTab('DETAILED RESULTS', 'detailed-results-tab', `
			<div class="detailed-results">
			</div>
		`);
        const detailedResults = new DetailedResults(this.rootElem.getElementsByClassName('detailed-results')[0], this, this.raidSimResultsManager);
    }
    addLogTab() {
        this.addTab('LOG', 'log-tab', `
			<div class="log-runner">
			</div>
		`);
        const logRunner = new LogRunner(this.rootElem.getElementsByClassName('log-runner')[0], this);
    }
    recomputeSettingsLayout() {
        if (this.settingsMuuri) {
            //this.settingsMuuri.refreshItems();
        }
        window.dispatchEvent(new Event('resize'));
    }
    modifyRaidProto(raidProto) {
        // Invoke all the buff bot callbacks.
        this.getBuffBots().forEach(buffBot => {
            const partyProto = raidProto.parties[buffBot.getPartyIndex()];
            if (!partyProto) {
                throw new Error('No party proto for party index: ' + buffBot.getPartyIndex());
            }
            buffBot.settings.modifyRaidProto(buffBot, raidProto, partyProto);
        });
        // Apply blessings.
        const numPaladins = this.getClassCount(Class.ClassPaladin);
        const blessingsAssignments = this.blessingsPicker.getAssignments();
        implementedSpecs.forEach(spec => {
            const playerProtos = raidProto.parties
                .map(party => party.players.filter(player => player.class != Class.ClassUnknown && playerToSpec(player) == spec))
                .flat();
            blessingsAssignments.paladins.forEach((paladin, i) => {
                if (i >= numPaladins) {
                    return;
                }
                if (paladin.blessings[spec] == Blessings.BlessingOfKings) {
                    playerProtos.forEach(playerProto => playerProto.buffs.blessingOfKings = true);
                }
                else if (paladin.blessings[spec] == Blessings.BlessingOfMight) {
                    playerProtos.forEach(playerProto => playerProto.buffs.blessingOfMight = TristateEffect.TristateEffectImproved);
                }
                else if (paladin.blessings[spec] == Blessings.BlessingOfWisdom) {
                    playerProtos.forEach(playerProto => playerProto.buffs.blessingOfWisdom = TristateEffect.TristateEffectImproved);
                }
                else if (paladin.blessings[spec] == Blessings.BlessingOfSalvation) {
                    playerProtos.forEach(playerProto => playerProto.buffs.blessingOfSalvation = true);
                }
                else if (paladin.blessings[spec] == Blessings.BlessingOfSanctuary) {
                    playerProtos.forEach(playerProto => playerProto.buffs.blessingOfSanctuary = true);
                }
            });
        });
    }
    getCurrentData() {
        if (this.raidSimResultsManager) {
            return this.raidSimResultsManager.getCurrentData();
        }
        else {
            return null;
        }
    }
    getReferenceData() {
        if (this.raidSimResultsManager) {
            return this.raidSimResultsManager.getReferenceData();
        }
        else {
            return null;
        }
    }
    getClassCount(playerClass) {
        return this.sim.raid.getClassCount(playerClass)
            + this.getBuffBots()
                .filter(buffBot => buffBot.getClass() == playerClass).length;
    }
    getBuffBots() {
        return this.raidPicker.getBuffBots();
    }
    setBuffBots(eventID, buffBotProtos) {
        this.raidPicker.setBuffBots(eventID, buffBotProtos);
    }
    clearBuffBots(eventID) {
        this.raidPicker.setBuffBots(eventID, []);
    }
    getPlayersAndBuffBots() {
        const players = this.sim.raid.getPlayers();
        const buffBots = this.getBuffBots();
        const playersAndBuffBots = players.slice();
        buffBots.forEach(buffBot => {
            playersAndBuffBots[buffBot.getRaidIndex()] = buffBot;
        });
        return playersAndBuffBots;
    }
    applyDefaults(eventID) {
        TypedEvent.freezeAllAndDo(() => {
            this.sim.raid.fromProto(eventID, RaidProto.create());
            this.sim.encounter.applyDefaults(eventID);
            this.sim.applyDefaults(eventID, true);
        });
    }
    toProto() {
        return RaidSimSettings.create({
            settings: this.sim.toProto(),
            raid: this.sim.raid.toProto(true),
            buffBots: this.getBuffBots().map(b => b.toProto()),
            blessings: this.blessingsPicker.getAssignments(),
            encounter: this.sim.encounter.toProto(),
        });
    }
    fromProto(eventID, settings) {
        TypedEvent.freezeAllAndDo(() => {
            if (settings.settings) {
                this.sim.fromProto(eventID, settings.settings);
            }
            this.sim.raid.fromProto(eventID, settings.raid || RaidProto.create());
            this.sim.encounter.fromProto(eventID, settings.encounter || EncounterProto.create());
            this.raidPicker.setBuffBots(eventID, settings.buffBots);
            this.blessingsPicker.setAssignments(eventID, settings.blessings || BlessingsAssignments.create());
        });
    }
    clearRaid(eventID) {
        this.sim.raid.clear(eventID);
        this.clearBuffBots(eventID);
    }
    // Returns the actual key to use for local storage, based on the given key part and the site context.
    getStorageKey(keyPart) {
        return '__raid__' + keyPart;
    }
    getSavedRaidStorageKey() {
        return this.getStorageKey('__savedRaid__');
    }
}
