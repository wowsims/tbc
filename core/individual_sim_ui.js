import { BonusStatsPicker } from '/tbc/core/components/bonus_stats_picker.js';
import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { CharacterStats } from '/tbc/core/components/character_stats.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Cooldowns } from '/tbc/core/proto/common.js';
import { CooldownsPicker } from '/tbc/core/components/cooldowns_picker.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { DetailedResults } from '/tbc/core/components/detailed_results.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EncounterPicker } from '/tbc/core/components/encounter_picker.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { TypedEvent } from './typed_event.js';
import { GearPicker } from '/tbc/core/components/gear_picker.js';
import { IconEnumPicker } from '/tbc/core/components/icon_enum_picker.js';
import { IconPicker } from '/tbc/core/components/icon_picker.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { IndividualSimSettings } from '/tbc/core/proto/ui.js';
import { LogRunner } from '/tbc/core/components/log_runner.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { SavedDataManager } from '/tbc/core/components/saved_data_manager.js';
import { SavedEncounter } from '/tbc/core/proto/ui.js';
import { SavedGearSet } from '/tbc/core/proto/ui.js';
import { SavedSettings } from '/tbc/core/proto/ui.js';
import { SavedTalents } from '/tbc/core/proto/ui.js';
import { SettingsMenu } from '/tbc/core/components/settings_menu.js';
import { ShattrathFaction } from '/tbc/core/proto/common.js';
import { SimUI } from './sim_ui.js';
import { nameToShattFaction } from '/tbc/core/proto_utils/utils.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { addRaidSimAction } from '/tbc/core/components/raid_sim_action.js';
import { addStatWeightsAction } from '/tbc/core/components/stat_weights_action.js';
import { getMetaGemConditionDescription } from '/tbc/core/proto_utils/gems.js';
import { isDualWieldSpec } from '/tbc/core/proto_utils/utils.js';
import { launchedSpecs } from '/tbc/core/launched_sims.js';
import { newIndividualExporters } from '/tbc/core/components/exporters.js';
import { newIndividualImporters } from '/tbc/core/components/importers.js';
import { newTalentsPicker } from '/tbc/core/talents/factory.js';
import { raceNames } from '/tbc/core/proto_utils/names.js';
import { isTankSpec } from '/tbc/core/proto_utils/utils.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';
import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
const SAVED_GEAR_STORAGE_KEY = '__savedGear__';
const SAVED_ROTATION_STORAGE_KEY = '__savedRotation__';
const SAVED_SETTINGS_STORAGE_KEY = '__savedSettings__';
const SAVED_TALENTS_STORAGE_KEY = '__savedTalents__';
class IndividualSimIconPicker {
    constructor(parent, modObj, input, simUI) {
        let picker = null;
        if ('states' in input) {
            picker = new IconPicker(parent, modObj, input);
        }
        else {
            picker = new IconEnumPicker(parent, modObj, input);
        }
        if (input.exclusivityTags) {
            simUI.registerExclusiveEffect({
                tags: input.exclusivityTags,
                changedEvent: picker.changeEmitter,
                isActive: () => Boolean(picker.getInputValue()),
                deactivate: (eventID) => picker.setValue(eventID, (typeof picker.getInputValue() == 'number') ? 0 : false),
            });
        }
    }
}
// Extended shared UI for all individual player sims.
export class IndividualSimUI extends SimUI {
    constructor(parentElem, player, config) {
        super(parentElem, player.sim, {
            spec: player.spec,
            knownIssues: config.knownIssues,
        });
        this.rootElem.classList.add('individual-sim-ui', config.cssClass);
        this.player = player;
        this.individualConfig = config;
        this.raidSimResultsManager = null;
        this.settingsMuuri = null;
        this.prevEpIterations = 0;
        this.prevEpSimResult = null;
        if (!launchedSpecs.includes(this.player.spec)) {
            this.addWarning({
                updateOn: new TypedEvent(),
                shouldDisplay: () => true,
                getContent: () => 'This sim is still under development.',
            });
        }
        this.addWarning({
            updateOn: this.player.gearChangeEmitter,
            shouldDisplay: () => this.player.getGear().hasInactiveMetaGem(),
            getContent: () => {
                const metaGem = this.player.getGear().getMetaGem();
                return `Meta gem disabled (${metaGem.name}): ${getMetaGemConditionDescription(metaGem)}`;
            },
        });
        (config.warnings || []).forEach(warning => this.addWarning(warning(this)));
        this.exclusivityMap = {
            'Battle Elixir': [],
            'Drums': [],
            'Food': [],
            'Pet Food': [],
            'Alchohol': [],
            'Guardian Elixir': [],
            'Potion': [],
            'Conjured': [],
            'Spirit': [],
            'MH Weapon Imbue': [],
            'OH Weapon Imbue': [],
        };
        if (!this.isWithinRaidSim) {
            // This needs to go before all the UI components so that gear loading is the
            // first callback invoked from waitForInit().
            this.sim.waitForInit().then(() => this.loadSettings());
        }
        this.addSidebarComponents();
        this.addTopbarComponents();
        this.addGearTab();
        this.addSettingsTab();
        this.addTalentsTab();
        if (!this.isWithinRaidSim) {
            this.addDetailedResultsTab();
            this.addLogTab();
        }
        this.player.changeEmitter.on(() => this.recomputeSettingsLayout());
    }
    loadSettings() {
        const initEventID = TypedEvent.nextEventID();
        TypedEvent.freezeAllAndDo(() => {
            let loadedSettings = false;
            let hash = window.location.hash;
            if (hash.length > 1) {
                // Remove leading '#'
                hash = hash.substring(1);
                try {
                    const binary = atob(hash);
                    const bytes = new Uint8Array(binary.length);
                    for (let i = 0; i < bytes.length; i++) {
                        bytes[i] = binary.charCodeAt(i);
                    }
                    const settingsBytes = pako.inflate(bytes);
                    const settings = IndividualSimSettings.fromBinary(settingsBytes);
                    this.fromProto(initEventID, settings);
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
                    const settings = IndividualSimSettings.fromJsonString(savedSettings);
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
            this.player.setName(initEventID, 'Player');
            // This needs to go last so it doesn't re-store things as they are initialized.
            this.changeEmitter.on(eventID => {
                const jsonStr = IndividualSimSettings.toJsonString(this.toProto());
                window.localStorage.setItem(this.getSettingsStorageKey(), jsonStr);
            });
        });
    }
    addSidebarComponents() {
        this.raidSimResultsManager = addRaidSimAction(this);
        addStatWeightsAction(this, this.individualConfig.epStats, this.individualConfig.epReferenceStat);
        const characterStats = new CharacterStats(this.rootElem.getElementsByClassName('sim-sidebar-footer')[0], this.player, this.individualConfig.displayStats, this.individualConfig.modifyDisplayStats, this.individualConfig.statBreakdowns);
    }
    addTopbarComponents() {
        this.addToolbarItem(newIndividualImporters(this));
        this.addToolbarItem(newIndividualExporters(this));
        const optionsMenu = document.createElement('span');
        optionsMenu.classList.add('fas', 'fa-cog');
        tippy(optionsMenu, {
            'content': 'Options',
            'allowHTML': true,
        });
        optionsMenu.addEventListener('click', event => {
            new SettingsMenu(this.rootElem, this);
        });
        this.addToolbarItem(optionsMenu);
    }
    addGearTab() {
        this.addTab('GEAR', 'gear-tab', `
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
                return SavedGearSet.create({
                    gear: player.getGear().asSpec(),
                    bonusStats: player.getBonusStats().asArray(),
                });
            },
            setData: (eventID, player, newSavedGear) => {
                TypedEvent.freezeAllAndDo(() => {
                    player.setGear(eventID, this.sim.lookupEquipmentSpec(newSavedGear.gear || EquipmentSpec.create()));
                    player.setBonusStats(eventID, new Stats(newSavedGear.bonusStats || []));
                });
            },
            changeEmitters: [this.player.changeEmitter],
            equals: (a, b) => SavedGearSet.equals(a, b),
            toJson: (a) => SavedGearSet.toJson(a),
            fromJson: (obj) => SavedGearSet.fromJson(obj),
        });
        this.sim.waitForInit().then(() => {
            savedGearManager.loadUserData();
            this.individualConfig.presets.gear.forEach(presetGear => {
                savedGearManager.addSavedData({
                    name: presetGear.name,
                    tooltip: presetGear.tooltip,
                    isPreset: true,
                    data: SavedGearSet.create({
                        // Convert to gear and back so order is always the same.
                        gear: this.sim.lookupEquipmentSpec(presetGear.gear).asSpec(),
                        bonusStats: new Stats().asArray(),
                    }),
                    enableWhen: presetGear.enableWhen,
                });
            });
        });
    }
    addSettingsTab() {
        this.addTab('SETTINGS', 'settings-tab', `
			<div class="settings-inputs">
				<div class="settings-section-container">
					<fieldset class="settings-section encounter-section within-raid-sim-hide">
						<legend>Encounter</legend>
					</fieldset>
					<fieldset class="settings-section race-section">
						<legend>Player</legend>
					</fieldset>
					<fieldset class="settings-section rotation-section">
						<legend>Rotation</legend>
					</fieldset>
				</div>
				<div class="settings-section-container custom-sections-container">
				</div>
				<div class="settings-section-container">
					<fieldset class="settings-section self-buffs-section">
						<legend>Self Buffs</legend>
					</fieldset>
				</div>
				<div class="settings-section-container within-raid-sim-hide">
					<fieldset class="settings-section buffs-section">
						<legend>Other Buffs</legend>
					</fieldset>
				</div>
				<div class="settings-section-container">
					<fieldset class="settings-section new-consumes-section">
						<legend>Consumes</legend>
						<div class="consumes-row">
							<span>Potions</span>
							<div class="consumes-row-inputs">
								<div class="consumes-potions"></div>
								<div class="consumes-conjured"></div>
							</div>
						</div>
						<div class="consumes-row">
							<span>Elixirs</span>
							<div class="consumes-row-inputs">
								<div class="consumes-flasks"></div>
								<span>OR</span>
								<div class="consumes-battle-elixirs"></div>
								<div class="consumes-guardian-elixirs"></div>
							</div>
						</div>
						<div class="consumes-row">
							<span>Imbues</span>
							<div class="consumes-row-inputs">
								<div class="consumes-imbue-mh"></div>
								<div class="consumes-imbue-oh"></div>
							</div>
						</div>
						<div class="consumes-row">
							<span>Food</span>
							<div class="consumes-row-inputs">
								<div class="consumes-food"></div>
								<div class="consumes-alcohol"></div>
							</div>
						</div>
						<div class="consumes-row">
							<span>Trade</span>
							<div class="consumes-row-inputs consumes-trade">
							</div>
						</div>
						<div class="consumes-row consumes-row-pet">
							<span>Pet</span>
							<div class="consumes-row-inputs consumes-pet">
							</div>
						</div>
						<div class="consumes-row">
							<div class="consumes-row-inputs consumes-other">
							</div>
						</div>
					</fieldset>
				</div>
				<div class="settings-section-container cooldowns-section-container">
					<fieldset class="settings-section cooldowns-section">
						<legend>Cooldowns</legend>
						<div class="cooldowns-section-content">
						</div>
					</fieldset>
				</div>
				<div class="settings-section-container within-raid-sim-hide">
					<fieldset class="settings-section debuffs-section">
						<legend>Debuffs</legend>
					</fieldset>
				</div>
				<div class="settings-section-container">
					<fieldset class="settings-section other-settings-section">
						<legend>Other</legend>
					</fieldset>
				</div>
			</div>
			<div class="settings-bottom-bar">
				<div class="saved-encounter-manager within-raid-sim-hide">
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
        configureIconSection(debuffsSection, this.individualConfig.debuffInputs.map(iconInput => new IndividualSimIconPicker(debuffsSection, this.sim.raid, iconInput, this)), Tooltips.DEBUFFS_SECTION);
        if (this.individualConfig.consumeOptions?.potions.length) {
            const elem = this.rootElem.getElementsByClassName('consumes-potions')[0];
            new IconEnumPicker(elem, this.player, IconInputs.makePotionsInput(this.individualConfig.consumeOptions.potions));
        }
        if (this.individualConfig.consumeOptions?.conjured.length) {
            const elem = this.rootElem.getElementsByClassName('consumes-conjured')[0];
            new IconEnumPicker(elem, this.player, IconInputs.makeConjuredInput(this.individualConfig.consumeOptions.conjured));
        }
        if (this.individualConfig.consumeOptions?.flasks.length) {
            const elem = this.rootElem.getElementsByClassName('consumes-flasks')[0];
            new IconEnumPicker(elem, this.player, IconInputs.makeFlasksInput(this.individualConfig.consumeOptions.flasks));
        }
        if (this.individualConfig.consumeOptions?.battleElixirs.length) {
            const elem = this.rootElem.getElementsByClassName('consumes-battle-elixirs')[0];
            new IconEnumPicker(elem, this.player, IconInputs.makeBattleElixirsInput(this.individualConfig.consumeOptions.battleElixirs));
        }
        if (this.individualConfig.consumeOptions?.guardianElixirs.length) {
            const elem = this.rootElem.getElementsByClassName('consumes-guardian-elixirs')[0];
            new IconEnumPicker(elem, this.player, IconInputs.makeGuardianElixirsInput(this.individualConfig.consumeOptions.guardianElixirs));
        }
        if (this.individualConfig.consumeOptions?.food.length) {
            const elem = this.rootElem.getElementsByClassName('consumes-food')[0];
            new IconEnumPicker(elem, this.player, IconInputs.makeFoodInput(this.individualConfig.consumeOptions.food));
        }
        if (this.individualConfig.consumeOptions?.alcohol.length) {
            const elem = this.rootElem.getElementsByClassName('consumes-alcohol')[0];
            new IconEnumPicker(elem, this.player, IconInputs.makeAlcoholInput(this.individualConfig.consumeOptions.alcohol));
        }
        if (this.individualConfig.consumeOptions?.weaponImbues.length) {
            const mhImbueElem = this.rootElem.getElementsByClassName('consumes-imbue-mh')[0];
            const ohImbueElem = this.rootElem.getElementsByClassName('consumes-imbue-oh')[0];
            new IconEnumPicker(mhImbueElem, this.player, IconInputs.makeWeaponImbueInput(true, this.individualConfig.consumeOptions.weaponImbues));
            if (isDualWieldSpec(this.player.spec)) {
                new IconEnumPicker(ohImbueElem, this.player, IconInputs.makeWeaponImbueInput(false, this.individualConfig.consumeOptions.weaponImbues));
            }
        }
        const tradeConsumesElem = this.rootElem.getElementsByClassName('consumes-trade')[0];
        new IndividualSimIconPicker(tradeConsumesElem, this.player, IconInputs.DrumsInput, this);
        new IndividualSimIconPicker(tradeConsumesElem, this.player, IconInputs.SuperSapper, this);
        new IndividualSimIconPicker(tradeConsumesElem, this.player, IconInputs.GoblinSapper, this);
        new IndividualSimIconPicker(tradeConsumesElem, this.player, IconInputs.FillerExplosiveInput, this);
        if (this.individualConfig.consumeOptions?.pet?.length) {
            const petConsumesElem = this.rootElem.getElementsByClassName('consumes-pet')[0];
            this.individualConfig.consumeOptions.pet.map(iconInput => new IndividualSimIconPicker(petConsumesElem, this.player, iconInput, this));
        }
        else {
            const petRowElem = this.rootElem.getElementsByClassName('consumes-row-pet')[0];
            petRowElem.style.display = 'none';
        }
        if (this.individualConfig.consumeOptions?.other?.length) {
            const containerElem = this.rootElem.getElementsByClassName('consumes-other')[0];
            this.individualConfig.consumeOptions.other.map(iconInput => new IndividualSimIconPicker(containerElem, this.player, iconInput, this));
        }
        const configureInputSection = (sectionElem, sectionConfig) => {
            if (sectionConfig.tooltip) {
                tippy(sectionElem, {
                    'content': sectionConfig.tooltip,
                    'allowHTML': true,
                });
            }
            sectionConfig.inputs.forEach(inputConfig => {
                if (inputConfig.type == 'number') {
                    new NumberPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
                }
                else if (inputConfig.type == 'boolean') {
                    new BooleanPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
                }
                else if (inputConfig.type == 'enum') {
                    new EnumPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
                }
                else if (inputConfig.type == 'iconEnum') {
                    new IconEnumPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
                }
            });
        };
        configureInputSection(this.rootElem.getElementsByClassName('rotation-section')[0], this.individualConfig.rotationInputs);
        if (this.individualConfig.otherInputs?.inputs.length) {
            configureInputSection(this.rootElem.getElementsByClassName('other-settings-section')[0], this.individualConfig.otherInputs);
        }
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
            setValue: (eventID, sim, newValue) => sim.setRace(eventID, newValue),
        });
        const shattFactionPicker = new EnumPicker(this.rootElem.getElementsByClassName('race-section')[0], this.player, {
            values: ["Scryer", "Aldor"].map(faction => {
                return {
                    name: faction,
                    value: nameToShattFaction[faction],
                };
            }),
            changedEvent: sim => sim.gearChangeEmitter,
            getValue: sim => sim.getShattFaction(),
            setValue: (eventID, sim, newValue) => sim.setShattFaction(eventID, newValue),
            showWhen: (player) => this.player.getEquippedItem(ItemSlot.ItemSlotNeck)?.item.id == 34678 || this.player.getEquippedItem(ItemSlot.ItemSlotNeck)?.item.id == 34679,
        });
        const encounterSectionElem = settingsTab.getElementsByClassName('encounter-section')[0];
        new EncounterPicker(encounterSectionElem, this.sim.encounter, this.individualConfig.encounterPicker);
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
        const cooldownSectionElem = settingsTab.getElementsByClassName('cooldowns-section')[0];
        const cooldownContentElem = settingsTab.getElementsByClassName('cooldowns-section-content')[0];
        new CooldownsPicker(cooldownContentElem, this.player);
        tippy(cooldownSectionElem, {
            content: Tooltips.COOLDOWNS_SECTION,
            allowHTML: true,
            placement: 'left',
        });
        // Init Muuri layout only when settings tab is clicked, because it needs the elements
        // to be shown so it can calculate sizes.
        this.rootElem.getElementsByClassName('settings-tab-tab')[0].addEventListener('click', event => {
            if (this.settingsMuuri == null) {
                setTimeout(() => {
                    this.settingsMuuri = new Muuri('.settings-inputs');
                }, 200); // Magic amount of time before Muuri init seems to work
            }
            setTimeout(() => {
                this.recomputeSettingsLayout();
            }, 200);
        });
        const savedSettingsManager = new SavedDataManager(this.rootElem.getElementsByClassName('saved-settings-manager')[0], this, {
            label: 'Settings',
            storageKey: this.getSavedSettingsStorageKey(),
            getData: (simUI) => {
                return SavedSettings.create({
                    raidBuffs: simUI.sim.raid.getBuffs(),
                    partyBuffs: simUI.player.getParty()?.getBuffs() || PartyBuffs.create(),
                    playerBuffs: simUI.player.getBuffs(),
                    debuffs: simUI.sim.raid.getDebuffs(),
                    consumes: simUI.player.getConsumes(),
                    race: simUI.player.getRace(),
                    cooldowns: simUI.player.getCooldowns(),
                });
            },
            setData: (eventID, simUI, newSettings) => {
                TypedEvent.freezeAllAndDo(() => {
                    simUI.sim.raid.setBuffs(eventID, newSettings.raidBuffs || RaidBuffs.create());
                    simUI.sim.raid.setDebuffs(eventID, newSettings.debuffs || Debuffs.create());
                    const party = simUI.player.getParty();
                    if (party) {
                        party.setBuffs(eventID, newSettings.partyBuffs || PartyBuffs.create());
                    }
                    simUI.player.setBuffs(eventID, newSettings.playerBuffs || IndividualBuffs.create());
                    simUI.player.setConsumes(eventID, newSettings.consumes || Consumes.create());
                    simUI.player.setRace(eventID, newSettings.race);
                    simUI.player.setCooldowns(eventID, newSettings.cooldowns || Cooldowns.create());
                });
            },
            changeEmitters: [
                this.sim.raid.buffsChangeEmitter,
                this.sim.raid.debuffsChangeEmitter,
                this.player.getParty().buffsChangeEmitter,
                this.player.buffsChangeEmitter,
                this.player.consumesChangeEmitter,
                this.player.raceChangeEmitter,
                this.player.cooldownsChangeEmitter,
            ],
            equals: (a, b) => SavedSettings.equals(a, b),
            toJson: (a) => SavedSettings.toJson(a),
            fromJson: (obj) => SavedSettings.fromJson(obj),
        });
        const customSectionsContainer = this.rootElem.getElementsByClassName('custom-sections-container')[0];
        let anyCustomSections = false;
        for (const [sectionName, sectionConfig] of Object.entries(this.individualConfig.additionalSections || {})) {
            const sectionCssPrefix = sectionName.replace(/\s+/g, '');
            const sectionElem = document.createElement('fieldset');
            sectionElem.classList.add('settings-section', sectionCssPrefix + '-section');
            sectionElem.innerHTML = `<legend>${sectionName}</legend>`;
            customSectionsContainer.appendChild(sectionElem);
            configureInputSection(sectionElem, sectionConfig);
            anyCustomSections = true;
        }
        ;
        for (const [sectionName, sectionConfig] of Object.entries(this.individualConfig.additionalIconSections || {})) {
            const sectionCssPrefix = sectionName.replace(/\s+/g, '');
            const sectionElem = document.createElement('fieldset');
            sectionElem.classList.add('settings-section', sectionCssPrefix + '-section');
            sectionElem.innerHTML = `<legend>${sectionName}</legend>`;
            customSectionsContainer.appendChild(sectionElem);
            configureIconSection(sectionElem, sectionConfig.map(iconInput => new IndividualSimIconPicker(sectionElem, this.player, iconInput, this)));
            anyCustomSections = true;
        }
        ;
        (this.individualConfig.customSections || []).forEach(customSection => {
            const sectionElem = document.createElement('fieldset');
            customSectionsContainer.appendChild(sectionElem);
            const sectionName = customSection(this, sectionElem);
            const sectionCssPrefix = sectionName.replace(/\s+/g, '');
            sectionElem.classList.add('settings-section', sectionCssPrefix + '-section');
            const labelElem = document.createElement('legend');
            labelElem.textContent = sectionName;
            sectionElem.prepend(labelElem);
            anyCustomSections = true;
        });
        if (!anyCustomSections) {
            customSectionsContainer.remove();
        }
        this.sim.waitForInit().then(() => {
            savedEncounterManager.loadUserData();
            savedSettingsManager.loadUserData();
        });
        Array.from(this.rootElem.getElementsByClassName('settings-section-container')).forEach((container, i) => {
            container.style.zIndex = String(1000 - i);
        });
    }
    addTalentsTab() {
        this.addTab('TALENTS', 'talents-tab', `
			<div class="talents-picker">
			</div>
			<div class="saved-talents-manager">
			</div>
		`);
        const talentsPicker = newTalentsPicker(this.rootElem.getElementsByClassName('talents-picker')[0], this.player);
        const savedTalentsManager = new SavedDataManager(this.rootElem.getElementsByClassName('saved-talents-manager')[0], this.player, {
            label: 'Talents',
            storageKey: this.getSavedTalentsStorageKey(),
            getData: (player) => SavedTalents.create({
                talentsString: player.getTalentsString(),
            }),
            setData: (eventID, player, newTalents) => player.setTalentsString(eventID, newTalents.talentsString),
            changeEmitters: [this.player.talentsChangeEmitter],
            equals: (a, b) => SavedTalents.equals(a, b),
            toJson: (a) => SavedTalents.toJson(a),
            fromJson: (obj) => SavedTalents.fromJson(obj),
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
                savedTalentsManager.addSavedData({
                    name: config.name,
                    isPreset: true,
                    data: SavedTalents.create({
                        talentsString: config.data,
                    }),
                });
            });
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
    applyDefaults(eventID) {
        TypedEvent.freezeAllAndDo(() => {
            const tankSpec = isTankSpec(this.player.spec);
            this.player.setRace(eventID, specToEligibleRaces[this.player.spec][0]);
            this.player.setShattFaction(eventID, ShattrathFaction.ShattrathFactionAldor);
            this.player.setGear(eventID, this.sim.lookupEquipmentSpec(this.individualConfig.defaults.gear));
            this.player.setBonusStats(eventID, new Stats());
            this.player.setConsumes(eventID, this.individualConfig.defaults.consumes);
            this.player.setRotation(eventID, this.individualConfig.defaults.rotation);
            this.player.setTalentsString(eventID, this.individualConfig.defaults.talents);
            this.player.setSpecOptions(eventID, this.individualConfig.defaults.specOptions);
            this.player.setBuffs(eventID, this.individualConfig.defaults.individualBuffs);
            this.player.setCooldowns(eventID, Cooldowns.create());
            this.player.getParty().setBuffs(eventID, this.individualConfig.defaults.partyBuffs);
            this.player.getRaid().setBuffs(eventID, this.individualConfig.defaults.raidBuffs);
            this.player.setEpWeights(eventID, this.individualConfig.defaults.epWeights);
            this.player.setInFrontOfTarget(eventID, tankSpec);
            if (!this.isWithinRaidSim) {
                this.sim.encounter.applyDefaults(eventID);
                this.sim.raid.setDebuffs(eventID, this.individualConfig.defaults.debuffs);
                this.sim.applyDefaults(eventID, tankSpec);
                if (tankSpec) {
                    this.sim.raid.setTanks(eventID, [this.player.makeRaidTarget()]);
                }
                else {
                    this.sim.raid.setTanks(eventID, []);
                }
            }
        });
    }
    registerExclusiveEffect(effect) {
        effect.tags.forEach(tag => {
            this.exclusivityMap[tag].push(effect);
            effect.changedEvent.on(eventID => {
                if (!effect.isActive())
                    return;
                // TODO: Mark the parent somehow so we can track this for undo/redo.
                const newEventID = TypedEvent.nextEventID();
                TypedEvent.freezeAllAndDo(() => {
                    this.exclusivityMap[tag].forEach(otherEffect => {
                        if (otherEffect == effect || !otherEffect.isActive())
                            return;
                        otherEffect.deactivate(newEventID);
                    });
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
    recomputeSettingsLayout() {
        if (this.settingsMuuri) {
            //this.settingsMuuri.refreshItems();
        }
        window.dispatchEvent(new Event('resize'));
    }
    // Returns the actual key to use for local storage, based on the given key part and the site context.
    getStorageKey(keyPart) {
        // Local storage is shared by all sites under the same domain, so we need to use
        // different keys for each spec site.
        return specToLocalStorageKey[this.player.spec] + keyPart;
    }
    toProto() {
        return IndividualSimSettings.create({
            settings: this.sim.toProto(),
            player: this.player.toProto(true),
            raidBuffs: this.sim.raid.getBuffs(),
            debuffs: this.sim.raid.getDebuffs(),
            tanks: this.sim.raid.getTanks(),
            partyBuffs: this.player.getParty()?.getBuffs() || PartyBuffs.create(),
            encounter: this.sim.encounter.toProto(),
            epWeights: this.player.getEpWeights().asArray(),
        });
    }
    fromProto(eventID, settings) {
        TypedEvent.freezeAllAndDo(() => {
            // TODO: Deprecate this
            if (settings.encounter.targets[0] && settings.encounter.targets[0].debuffs) {
                settings.debuffs = settings.encounter.targets[0].debuffs;
                settings.encounter.targets[0].debuffs = undefined;
            }
            if (!settings.player) {
                return;
            }
            if (settings.settings) {
                this.sim.fromProto(eventID, settings.settings);
            }
            this.player.fromProto(eventID, settings.player);
            if (settings.epWeights?.length > 0) {
                this.player.setEpWeights(eventID, new Stats(settings.epWeights));
            }
            else {
                this.player.setEpWeights(eventID, this.individualConfig.defaults.epWeights);
            }
            this.sim.raid.setBuffs(eventID, settings.raidBuffs || RaidBuffs.create());
            this.sim.raid.setDebuffs(eventID, settings.debuffs || Debuffs.create());
            this.sim.raid.setTanks(eventID, settings.tanks || []);
            const party = this.player.getParty();
            if (party) {
                party.setBuffs(eventID, settings.partyBuffs || PartyBuffs.create());
            }
            this.sim.encounter.fromProto(eventID, settings.encounter || EncounterProto.create());
        });
    }
}
