import { Actions } from '/tbc/core/components/actions.js';
import { CharacterStats } from '/tbc/core/components/character_stats.js';
import { CustomStatsPicker } from '/tbc/core/components/custom_stats_picker.js';
import { DetailedResults } from '/tbc/core/components/detailed_results.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { GearPicker } from '/tbc/core/components/gear_picker.js';
import { IconPicker } from '/tbc/core/components/icon_picker.js';
import { LogRunner } from '/tbc/core/components/log_runner.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { MobTypePickerConfig } from '/tbc/core/components/other_inputs.js';
import { Results } from '/tbc/core/components/results.js';
import { SavedDataManager } from '/tbc/core/components/saved_data_manager.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { raceNames } from '/tbc/core/proto_utils/names.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { newTalentsPicker } from '/tbc/core/talents/factory.js';
import { SimUI } from '/tbc/core/sim_ui.js';
export class DefaultTheme extends SimUI {
    constructor(parentElem, config) {
        super(parentElem, config);
        this._config = config;
        this.parentElem.innerHTML = layoutHTML;
        const titleElem = this.parentElem.getElementsByClassName('default-title')[0];
        if (config.releaseStatus == 'Alpha') {
            titleElem.textContent += ' Alpha';
        }
        else if (config.releaseStatus == 'Beta') {
            titleElem.textContent += ' Beta';
        }
        const results = new Results(this.parentElem.getElementsByClassName('default-results')[0], this);
        const detailedResults = new DetailedResults(this.parentElem.getElementsByClassName('detailed-results')[0]);
        const actions = new Actions(this.parentElem.getElementsByClassName('default-actions')[0], this, config.player.epStats, config.player.epReferenceStat, results, detailedResults);
        const logRunner = new LogRunner(this.parentElem.getElementsByClassName('log-runner')[0], this, results, detailedResults);
        const characterStats = new CharacterStats(this.parentElem.getElementsByClassName('default-stats')[0], config.player.displayStats, this.player);
        const gearPicker = new GearPicker(this.parentElem.getElementsByClassName('gear-picker')[0], this.player);
        const customStatsPicker = new CustomStatsPicker(this.parentElem.getElementsByClassName('custom-stats-picker')[0], this.player, config.player.epStats);
        const talentsPicker = newTalentsPicker(this.player.spec, this.parentElem.getElementsByClassName('talents-picker')[0], this.player);
        if (this._config.freezeTalents) {
            talentsPicker.freeze();
        }
        const settingsTab = document.getElementsByClassName('settings-inputs')[0];
        const configureIconSection = (sectionElem, sectionConfig, modObject) => {
            if (sectionConfig.tooltip) {
                tippy(sectionElem, {
                    'content': sectionConfig.tooltip,
                    'allowHTML': true,
                });
            }
            const iconPicker = new IconPicker(sectionElem, modObject, sectionConfig.icons, this);
        };
        configureIconSection(this.parentElem.getElementsByClassName('self-buffs-section')[0], config.selfBuffInputs, this.player);
        configureIconSection(this.parentElem.getElementsByClassName('buffs-section')[0], config.buffInputs, this.sim);
        configureIconSection(this.parentElem.getElementsByClassName('debuffs-section')[0], config.debuffInputs, this.target);
        configureIconSection(this.parentElem.getElementsByClassName('consumes-section')[0], config.consumeInputs, this.player);
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
                else if (inputConfig.type == 'enum') {
                    const picker = new EnumPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
                }
            });
        };
        configureInputSection(this.parentElem.getElementsByClassName('rotation-section')[0], config.rotationInputs);
        if (config.otherInputs?.inputs.length) {
            configureInputSection(this.parentElem.getElementsByClassName('other-settings-section')[0], config.otherInputs);
        }
        const makeInputSection = (sectionName, sectionConfig) => {
            const sectionCssPrefix = sectionName.replace(/\s+/g, '');
            const sectionElem = document.createElement('section');
            sectionElem.classList.add('settings-section', sectionCssPrefix + '-section');
            sectionElem.innerHTML = `<label>${sectionName}</label>`;
            settingsTab.appendChild(sectionElem);
            configureInputSection(sectionElem, sectionConfig);
        };
        for (const [sectionName, sectionConfig] of Object.entries(config.additionalSections || {})) {
            makeInputSection(sectionName, sectionConfig);
        }
        ;
        const races = specToEligibleRaces[this.player.spec];
        const racePicker = new EnumPicker(this.parentElem.getElementsByClassName('race-section')[0], this.player, {
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
        new NumberPicker(encounterSectionElem, this.sim, {
            label: 'Duration',
            changedEvent: (sim) => sim.encounterChangeEmitter,
            getValue: (sim) => sim.getEncounter().duration,
            setValue: (sim, newValue) => {
                const encounter = sim.getEncounter();
                encounter.duration = newValue;
                sim.setEncounter(encounter);
            },
        });
        if (config.showTargetArmor) {
            new NumberPicker(encounterSectionElem, this.target, {
                label: 'Target Armor',
                changedEvent: (target) => target.armorChangeEmitter,
                getValue: (target) => target.getArmor(),
                setValue: (target, newValue) => {
                    target.setArmor(newValue);
                },
            });
        }
        new EnumPicker(encounterSectionElem, this.target, MobTypePickerConfig);
        if (config.showNumTargets) {
            new NumberPicker(encounterSectionElem, this.sim, {
                label: '# of Targets',
                changedEvent: (sim) => sim.numTargetsChangeEmitter,
                getValue: (sim) => sim.getNumTargets(),
                setValue: (sim, newValue) => {
                    sim.setNumTargets(newValue);
                },
            });
        }
        // Init Muuri layout only when settings tab is clicked, because it needs the elements
        // to be shown so it can calculate sizes.
        let muuriInit = false;
        document.getElementById('settings-tab-toggle').addEventListener('click', event => {
            if (muuriInit) {
                return;
            }
            muuriInit = true;
            setTimeout(() => {
                new Muuri('.settings-inputs');
            }, 200); // Magic amount of time before Muuri init seems to work
        });
    }
    async init() {
        const savedGearManager = new SavedDataManager(this.parentElem.getElementsByClassName('saved-gear-manager')[0], this.player, {
            label: 'Gear',
            storageKey: this.getSavedGearStorageKey(),
            getData: (player) => {
                return {
                    gear: player.getGear(),
                    customStats: player.getCustomStats(),
                };
            },
            setData: (player, newGearAndStats) => {
                player.setGear(newGearAndStats.gear);
                player.setCustomStats(newGearAndStats.customStats);
            },
            changeEmitters: [this.player.gearChangeEmitter, this.player.customStatsChangeEmitter],
            equals: (a, b) => a.gear.equals(b.gear) && a.customStats.equals(b.customStats),
            toJson: (a) => {
                return {
                    gear: EquipmentSpec.toJson(a.gear.asSpec()),
                    customStats: a.customStats.toJson(),
                };
            },
            fromJson: (obj) => {
                return {
                    gear: this.sim.lookupEquipmentSpec(EquipmentSpec.fromJson(obj['gear'])),
                    customStats: Stats.fromJson(obj['customStats']),
                };
            },
        });
        const savedEncounterManager = new SavedDataManager(this.parentElem.getElementsByClassName('saved-encounter-manager')[0], this.sim, {
            label: 'Encounter',
            storageKey: this.getSavedEncounterStorageKey(),
            getData: (sim) => sim.getEncounter(),
            setData: (sim, newEncounter) => sim.setEncounter(newEncounter),
            changeEmitters: [this.sim.encounterChangeEmitter],
            equals: (a, b) => Encounter.equals(a, b),
            toJson: (a) => Encounter.toJson(a),
            fromJson: (obj) => Encounter.fromJson(obj),
        });
        const savedRotationManager = new SavedDataManager(this.parentElem.getElementsByClassName('saved-rotation-manager')[0], this.player, {
            label: 'Rotation',
            storageKey: this.getSavedRotationStorageKey(),
            getData: (player) => player.getRotation(),
            setData: (player, newRotation) => player.setRotation(newRotation),
            changeEmitters: [this.player.rotationChangeEmitter],
            equals: (a, b) => this.player.specTypeFunctions.rotationEquals(a, b),
            toJson: (a) => this.player.specTypeFunctions.rotationToJson(a),
            fromJson: (obj) => this.player.specTypeFunctions.rotationFromJson(obj),
        });
        const savedSettingsManager = new SavedDataManager(this.parentElem.getElementsByClassName('saved-settings-manager')[0], this, {
            label: 'Settings',
            storageKey: this.getSavedSettingsStorageKey(),
            getData: (simUI) => {
                return {
                    raidBuffs: simUI.sim.getRaidBuffs(),
                    partyBuffs: simUI.sim.getPartyBuffs(),
                    individualBuffs: simUI.sim.getIndividualBuffs(),
                    consumes: simUI.player.getConsumes(),
                    race: simUI.player.getRace(),
                };
            },
            setData: (simUI, newSettings) => {
                simUI.sim.setRaidBuffs(newSettings.raidBuffs);
                simUI.sim.setPartyBuffs(newSettings.partyBuffs);
                simUI.sim.setIndividualBuffs(newSettings.individualBuffs);
                simUI.player.setConsumes(newSettings.consumes);
                simUI.player.setRace(newSettings.race);
            },
            changeEmitters: [
                this.sim.raidBuffsChangeEmitter,
                this.sim.partyBuffsChangeEmitter,
                this.sim.individualBuffsChangeEmitter,
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
        const savedTalentsManager = new SavedDataManager(this.parentElem.getElementsByClassName('saved-talents-manager')[0], this.player, {
            label: 'Talents',
            storageKey: this.getSavedTalentsStorageKey(),
            getData: (player) => player.getTalentsString(),
            setData: (player, newTalentsString) => player.setTalentsString(newTalentsString),
            changeEmitters: [this.player.talentsStringChangeEmitter],
            equals: (a, b) => a == b,
            toJson: (a) => a,
            fromJson: (obj) => obj,
        });
        if (this._config.freezeTalents) {
            savedTalentsManager.freeze();
        }
        await super.init();
        savedGearManager.loadUserData();
        this._config.presets.gear.forEach(gearConfig => {
            const gear = this.sim.lookupEquipmentSpec(gearConfig.equipment);
            savedGearManager.addSavedData(gearConfig.name, { gear: gear, customStats: new Stats(), }, true, gearConfig.tooltip);
        });
        savedEncounterManager.loadUserData();
        this._config.presets.encounters.forEach(encounterConfig => {
            savedEncounterManager.addSavedData(encounterConfig.name, encounterConfig.encounter, true, encounterConfig.tooltip);
        });
        savedSettingsManager.loadUserData();
        savedTalentsManager.loadUserData();
        this._config.presets.talents.forEach(talentsConfig => {
            savedTalentsManager.addSavedData(talentsConfig.name, talentsConfig.talents, true, talentsConfig.tooltip);
        });
    }
}
const layoutHTML = `
<div class="default-root">
  <section class="default-sidebar">
    <div class="default-title">
      TBC Elemental Shaman Sim
    </div>
    <div class="default-actions">
    </div>
    <div class="default-results">
    </div>
    <div class="default-stats">
    </div>
  </section>
  <section class="default-main">
    <ul class="nav nav-tabs">
      <li class="active"><a data-toggle="tab" href="#gear-tab">Gear</a></li>
      <li><a id="settings-tab-toggle" data-toggle="tab" href="#settings-tab">Settings</a></li>
      <li><a data-toggle="tab" href="#talents-tab">Talents</a></li>
      <li><a data-toggle="tab" href="#detailed-results-tab">Detailed Results</a></li>
      <li><a data-toggle="tab" href="#log-tab">Log</a></li>
      <li class="default-top-bar">
				<div class="known-issues">Known Issues</div>
				<span class="share-link fa fa-link"></span>
			</li>
    </ul>
    <div class="tab-content">
      <div id="gear-tab" class="tab-pane fade in active">
        <div class="gear-tab">
          <div class="left-gear-panel">
            <div class="gear-picker">
            </div>
          </div>
          <div class="right-gear-panel">
            <div class="custom-stats-picker">
            </div>
            <div class="saved-gear-manager">
            </div>
          </div>
        </div>
      </div>
      <div id="settings-tab" class="settings-tab tab-pane fade"">
        <div class="settings-inputs">
          <div class="settings-section-container">
            <section class="settings-section encounter-section">
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
          <div class="settings-section-container">
            <section class="settings-section buffs-section">
              <label>Other Buffs</label>
            </section>
          </div>
          <div class="settings-section-container">
            <section class="settings-section consumes-section">
              <label>Consumes</label>
            </section>
          </div>
          <div class="settings-section-container">
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
          <div class="saved-encounter-manager">
          </div>
          <div class="saved-rotation-manager">
          </div>
          <div class="saved-settings-manager">
          </div>
        </div>
      </div>
      <div id="talents-tab" class="tab-pane fade"">
        <div class="talents-tab">
          <div class="talents-picker">
          </div>
          <div class="saved-talents-manager">
          </div>
        </div>
      </div>
      <div id="detailed-results-tab" class="tab-pane fade">
				<div class="detailed-results">
				</div>
      </div>
      <div id="log-tab" class="tab-pane fade">
				<div class="log-runner">
				</div>
      </div>
    </div>
  </section>
</div>
`;
