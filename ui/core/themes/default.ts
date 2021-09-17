import { Gear } from '../gear';
import { Sim } from '../sim';
import { EquipmentSpec } from '../api/newapi';
import { Spec } from '../api/newapi';
import { Stat } from '../api/newapi';
import { Actions } from '../components/actions';
import { CharacterStats } from '../components/character_stats';
import { CustomStatsPicker } from '../components/custom_stats_picker';
import { GearPicker } from '../components/gear_picker';
import { IconInput } from '../components/icon_picker';
import { IconPicker } from '../components/icon_picker';
import { NumberPicker } from '../components/number_picker';
import { RacePicker } from '../components/race_picker';
import { Results } from '../components/results';
import { SavedDataManager } from '../components/saved_data_manager';
import { newTalentsPicker } from '../talents/factory';

import { Theme, ThemeConfig } from './theme';

export interface DefaultThemeConfig extends ThemeConfig {
  displayStats: Array<Stat>;
  iconSections: Record<string, Array<IconInput>>;
  showTargetArmor: boolean;
  showNumTargets: boolean;
  presets: {
    gear: Record<string, EquipmentSpec>;
  },
}

export class DefaultTheme extends Theme {
  private readonly _config: DefaultThemeConfig;

  constructor(parentElem: HTMLElement, config: DefaultThemeConfig) {
    super(parentElem, config)
    this._config = config;
    this.parentElem.innerHTML = layoutHTML;

    const results = new Results(this.parentElem.getElementsByClassName('default-results')[0] as HTMLElement);
    const actions = new Actions(this.parentElem.getElementsByClassName('default-actions')[0] as HTMLElement, this.sim, results, config.epStats, config.epReferenceStat);

    const characterStats = new CharacterStats(this.parentElem.getElementsByClassName('default-stats')[0] as HTMLElement, config.displayStats);

    const gearPicker = new GearPicker(this.parentElem.getElementsByClassName('gear-picker')[0] as HTMLElement, this.sim);
    const customStatsPicker = new CustomStatsPicker(this.parentElem.getElementsByClassName('custom-stats-picker')[0] as HTMLElement, this.sim, config.epStats);

    const racePicker = new RacePicker(this.parentElem.getElementsByClassName('race-picker')[0] as HTMLElement, this.sim);
    const talentsPicker = newTalentsPicker(config.spec, this.parentElem.getElementsByClassName('talents-picker')[0] as HTMLElement, this.sim);

    const settingsTab = document.getElementsByClassName('settings-tab')[0] as HTMLElement;
    Object.keys(config.iconSections).forEach(pickerName => {
      const section = config.iconSections[pickerName];

      const sectionElem = document.createElement('section');
      sectionElem.classList.add('settings-section', pickerName + '-section');
      sectionElem.innerHTML = `<label>${pickerName}</label>`;
      settingsTab.appendChild(sectionElem);

      const iconPicker = new IconPicker(sectionElem, pickerName + '-icon-picker', this.sim, section, this);
    });


    const encounterSectionElem = document.createElement('section');
    encounterSectionElem.classList.add('settings-section', 'encounter-section');
    encounterSectionElem.innerHTML = `<label>Encounter</label>`;
    settingsTab.appendChild(encounterSectionElem);

    new NumberPicker(encounterSectionElem, this.sim, {
      label: 'Duration',
      changedEvent: (sim: Sim) => sim.encounterChangeEmitter,
      getValue: (sim: Sim) => sim.getEncounter().duration,
      setValue: (sim: Sim, newValue: number) => {
        const encounter = sim.getEncounter();
        encounter.duration = newValue;
        sim.setEncounter(encounter);
      },
    });

    if (config.showTargetArmor) {
      new NumberPicker(encounterSectionElem, this.sim, {
        label: 'Target Armor',
        changedEvent: (sim: Sim) => sim.encounterChangeEmitter,
        getValue: (sim: Sim) => sim.getEncounter().targetArmor,
        setValue: (sim: Sim, newValue: number) => {
          const encounter = sim.getEncounter();
          encounter.targetArmor = newValue;
          sim.setEncounter(encounter);
        },
      });
    } else {
    }

    if (config.showNumTargets) {
      new NumberPicker(encounterSectionElem, this.sim, {
        label: '# of Targets',
        changedEvent: (sim: Sim) => sim.encounterChangeEmitter,
        getValue: (sim: Sim) => sim.getEncounter().numTargets,
        setValue: (sim: Sim, newValue: number) => {
          const encounter = sim.getEncounter();
          encounter.numTargets = newValue;
          sim.setEncounter(encounter);
        },
      });
    }
  }

  async init(): Promise<void> {
    const savedGearManager = new SavedDataManager<Gear>(this.parentElem.getElementsByClassName('saved-gear-manager')[0] as HTMLElement, this.sim, {
      label: 'Gear',
      presets: {},
      getData: (sim: Sim) => sim.getGear(),
      setData: (sim: Sim, newGear: Gear) => sim.setGear(newGear),
      changeEmitter: this.sim.gearChangeEmitter,
      equals: (a: Gear, b: Gear) => a.equals(b),
    });

    await super.init();

    Object.keys(this._config.presets.gear).forEach(setName => {
      const gear = this.sim.lookupEquipmentSpec(this._config.presets.gear[setName]);
      savedGearManager.addSavedData(setName, gear, true);
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
    </ul>
    <div class="tab-content">
      <div id="gear-tab" class="tab-pane fade in active">
        <div class="gear-tab">
          <div class="left-gear-panel">
            <div class="gear-picker">
            </div>
            <div class="saved-gear-manager">
            </div>
          </div>
          <div class="right-gear-panel">
            <div class="custom-stats-picker">
            </div>
          </div>
        </div>
      </div>
      <div id="settings-tab" class="tab-pane fade"">
        <div class="settings-tab">
          <section class="settings-section race-picker">
            <label>Race</label>
          </section>
        </div>
      </div>
      <div id="talents-tab" class="tab-pane fade"">
        <div class="talents-tab">
          <div class="talents-picker">
          </div>
        </div>
      </div>
    </div>
  </section>
</div>
`;
