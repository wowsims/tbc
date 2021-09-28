import { raceNames } from '../api/names.js';
import { Buffs } from '../api/common.js';
import { Consumes } from '../api/common.js';
import { Encounter } from '../api/common.js';
import { EquipmentSpec } from '../api/common.js';
import { Stats } from '../api/stats.js';
import { specToEligibleRaces } from '../api/utils.js';
import { Actions } from '../components/actions.js';
import { CharacterStats } from '../components/character_stats.js';
import { CustomStatsPicker } from '../components/custom_stats_picker.js';
import { DetailedResults } from '../components/detailed_results.js';
import { EnumPicker } from '../components/enum_picker.js';
import { GearPicker } from '../components/gear_picker.js';
import { IconPicker } from '../components/icon_picker.js';
import { LogRunner } from '../components/log_runner.js';
import { NumberPicker } from '../components/number_picker.js';
import { Results } from '../components/results.js';
import { SavedDataManager } from '../components/saved_data_manager.js';
import { newTalentsPicker } from '../talents/factory.js';
import { SimUI } from '../sim_ui.js';
export class DefaultTheme extends SimUI {
    constructor(parentElem, config) {
        super(parentElem, config);
        this._config = config;
        this.parentElem.innerHTML = layoutHTML;
        const results = new Results(this.parentElem.getElementsByClassName('default-results')[0]);
        const detailedResults = new DetailedResults(this.parentElem.getElementsByClassName('detailed-results')[0]);
        const actions = new Actions(this.parentElem.getElementsByClassName('default-actions')[0], this.sim, config.epStats, config.epReferenceStat, results, detailedResults);
        const logRunner = new LogRunner(this.parentElem.getElementsByClassName('log-runner')[0], this.sim, results, detailedResults);
        const characterStats = new CharacterStats(this.parentElem.getElementsByClassName('default-stats')[0], config.displayStats, this.sim);
        const gearPicker = new GearPicker(this.parentElem.getElementsByClassName('gear-picker')[0], this.sim);
        const customStatsPicker = new CustomStatsPicker(this.parentElem.getElementsByClassName('custom-stats-picker')[0], this.sim, config.epStats);
        const talentsPicker = newTalentsPicker(config.spec, this.parentElem.getElementsByClassName('talents-picker')[0], this.sim);
        if (this._config.freezeTalents) {
            talentsPicker.freeze();
        }
        const settingsTab = document.getElementsByClassName('settings-inputs')[0];
        Object.keys(config.iconSections).forEach(sectionName => {
            const sectionConfig = config.iconSections[sectionName];
            const sectionElem = document.createElement('section');
            sectionElem.classList.add('settings-section', sectionName + '-section');
            sectionElem.innerHTML = `<label>${sectionName}</label>`;
            settingsTab.appendChild(sectionElem);
            const iconPicker = new IconPicker(sectionElem, sectionName + '-icon-picker', this.sim, sectionConfig, this);
        });
        Object.keys(config.otherSections).forEach(sectionName => {
            const sectionConfig = config.otherSections[sectionName];
            const sectionElem = document.createElement('section');
            sectionElem.classList.add('settings-section', sectionName + '-section');
            sectionElem.innerHTML = `<label>${sectionName}</label>`;
            settingsTab.appendChild(sectionElem);
            sectionConfig.forEach(inputConfig => {
                if (inputConfig.type == 'number') {
                    const picker = new NumberPicker(sectionElem, this.sim, inputConfig.config);
                }
                else if (inputConfig.type == 'enum') {
                    const picker = new EnumPicker(sectionElem, this.sim, inputConfig.config);
                }
            });
        });
        const races = specToEligibleRaces[this.sim.spec];
        const racePicker = new EnumPicker(this.parentElem.getElementsByClassName('race-picker')[0], this.sim, {
            names: races.map(race => raceNames[race]),
            values: races,
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
            new NumberPicker(encounterSectionElem, this.sim, {
                label: 'Target Armor',
                changedEvent: (sim) => sim.encounterChangeEmitter,
                getValue: (sim) => sim.getEncounter().targetArmor,
                setValue: (sim, newValue) => {
                    const encounter = sim.getEncounter();
                    encounter.targetArmor = newValue;
                    sim.setEncounter(encounter);
                },
            });
        }
        else {
        }
        if (config.showNumTargets) {
            new NumberPicker(encounterSectionElem, this.sim, {
                label: '# of Targets',
                changedEvent: (sim) => sim.encounterChangeEmitter,
                getValue: (sim) => sim.getEncounter().numTargets,
                setValue: (sim, newValue) => {
                    const encounter = sim.getEncounter();
                    encounter.numTargets = newValue;
                    sim.setEncounter(encounter);
                },
            });
        }
    }
    async init() {
        const savedGearManager = new SavedDataManager(this.parentElem.getElementsByClassName('saved-gear-manager')[0], this.sim, {
            label: 'Gear',
            getData: (sim) => {
                return {
                    gear: sim.getGear(),
                    customStats: sim.getCustomStats(),
                };
            },
            setData: (sim, newGearAndStats) => {
                sim.setGear(newGearAndStats.gear);
                sim.setCustomStats(newGearAndStats.customStats);
            },
            changeEmitters: [this.sim.gearChangeEmitter, this.sim.customStatsChangeEmitter],
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
            getData: (sim) => sim.getEncounter(),
            setData: (sim, newEncounter) => sim.setEncounter(newEncounter),
            changeEmitters: [this.sim.encounterChangeEmitter],
            equals: (a, b) => Encounter.equals(a, b),
            toJson: (a) => Encounter.toJson(a),
            fromJson: (obj) => Encounter.fromJson(obj),
        });
        const savedAgentManager = new SavedDataManager(this.parentElem.getElementsByClassName('saved-agent-manager')[0], this.sim, {
            label: 'Rotation',
            getData: (sim) => sim.getAgent(),
            setData: (sim, newAgent) => sim.setAgent(newAgent),
            changeEmitters: [this.sim.agentChangeEmitter],
            equals: (a, b) => this.sim.specTypeFunctions.agentEquals(a, b),
            toJson: (a) => this.sim.specTypeFunctions.agentToJson(a),
            fromJson: (obj) => this.sim.specTypeFunctions.agentFromJson(obj),
        });
        const savedSettingsManager = new SavedDataManager(this.parentElem.getElementsByClassName('saved-settings-manager')[0], this.sim, {
            label: 'Settings',
            getData: (sim) => {
                return {
                    buffs: sim.getBuffs(),
                    consumes: sim.getConsumes(),
                    race: sim.getRace(),
                };
            },
            setData: (sim, newSettings) => {
                sim.setBuffs(newSettings.buffs);
                sim.setConsumes(newSettings.consumes);
                sim.setRace(newSettings.race);
            },
            changeEmitters: [this.sim.buffsChangeEmitter, this.sim.consumesChangeEmitter, this.sim.raceChangeEmitter],
            equals: (a, b) => Buffs.equals(a.buffs, b.buffs) && Consumes.equals(a.consumes, b.consumes) && a.race == b.race,
            toJson: (a) => {
                return {
                    buffs: Buffs.toJson(a.buffs),
                    consumes: Consumes.toJson(a.consumes),
                    race: a.race,
                };
            },
            fromJson: (obj) => {
                return {
                    buffs: Buffs.fromJson(obj['buffs']),
                    consumes: Consumes.fromJson(obj['consumes']),
                    race: Number(obj['race']),
                };
            },
        });
        const savedTalentsManager = new SavedDataManager(this.parentElem.getElementsByClassName('saved-talents-manager')[0], this.sim, {
            label: 'Talents',
            getData: (sim) => sim.getTalentsString(),
            setData: (sim, newTalentsString) => sim.setTalentsString(newTalentsString),
            changeEmitters: [this.sim.talentsStringChangeEmitter],
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
      <li><a data-toggle="tab" href="#settings-tab">Settings</a></li>
      <li><a data-toggle="tab" href="#talents-tab">Talents</a></li>
      <li><a data-toggle="tab" href="#detailed-results-tab">Detailed Results</a></li>
      <li><a data-toggle="tab" href="#log-tab">Log</a></li>
      <li class="default-top-bar">
				<span class="share-link fa fa-link"></span
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
      <div id="settings-tab" class="tab-pane fade"">
        <div class="settings-tab">
          <div class="settings-inputs">
            <div class="settings-left-bar">
              <section class="settings-section encounter-section">
                <label>Encounter</label>
              </section>
              <section class="settings-section race-picker">
                <label>Race</label>
              </section>
            </div>
          </div>
          <div class="settings-bottom-bar">
            <div class="saved-encounter-manager">
            </div>
            <div class="saved-agent-manager">
            </div>
            <div class="saved-settings-manager">
            </div>
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
