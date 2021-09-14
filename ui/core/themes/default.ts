import { Spec } from '../api/newapi';
import { Stat } from '../api/newapi';
import { Actions } from '../components/actions.js';
import { CharacterStats } from '../components/character_stats';
import { GearPicker } from '../components/gear_picker';
import { IconInput } from '../components/icon_picker';
import { IconPicker } from '../components/icon_picker';
import { RacePicker } from '../components/race_picker';
import { Results } from '../components/results';
import { newTalentsPicker } from '../talents/factory';

import { Theme } from './theme.js';

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
        <div class="character-picker">
          <div class="gear-picker">
          </div>
        </div>
      </div>
      <div id="settings-tab" class="tab-pane fade"">
        <section class="settings-section race-picker">
          <label>Race</label>
        </section>
      </div>
      <div id="talents-tab" class="tab-pane fade"">
        <div class="talents-picker">
        </div>
      </div>
    </div>
  </section>
</div>
`;

export class DefaultTheme extends Theme {
  constructor(parentElem: HTMLElement, spec: Spec, iconPickers: Record<string, Array<IconInput>>) {
    super(parentElem, spec)

    this.parentElem.innerHTML = layoutHTML;

    const epStats = [
      Stat.StatIntellect,
      Stat.StatSpellPower,
      Stat.StatNatureSpellPower,
      Stat.StatSpellHit,
      Stat.StatSpellCrit,
      Stat.StatSpellHaste,
      Stat.StatMP5,
    ];
    const epReferenceStat = Stat.StatSpellPower;

    const displayStats = [
      Stat.StatStamina,
      Stat.StatIntellect,
      Stat.StatSpellPower,
      Stat.StatSpellHit,
      Stat.StatSpellCrit,
      Stat.StatSpellHaste,
      Stat.StatMP5,
    ];

    const results = new Results(this.parentElem.getElementsByClassName('default-results')[0] as HTMLElement);
    const actions = new Actions(this.parentElem.getElementsByClassName('default-actions')[0] as HTMLElement, this.sim, results, epStats, epReferenceStat);

    const characterStats = new CharacterStats(this.parentElem.getElementsByClassName('default-stats')[0] as HTMLElement, displayStats);

    const gearPicker = new GearPicker(this.parentElem.getElementsByClassName('gear-picker')[0] as HTMLElement, this.sim);
    const racePicker = new RacePicker(this.parentElem.getElementsByClassName('race-picker')[0] as HTMLElement, this.sim);
    const talentsPicker = newTalentsPicker(spec, this.parentElem.getElementsByClassName('talents-picker')[0] as HTMLElement, this.sim);

    const settingsTab = document.getElementById('settings-tab') as HTMLElement;
    Object.keys(iconPickers).forEach(pickerName => {
      const sectionElem = document.createElement('section');
      sectionElem.classList.add('settings-section', pickerName + '-section');
      sectionElem.innerHTML = `<label>${pickerName}</label>`;
      settingsTab.appendChild(sectionElem);

      const iconPicker = new IconPicker(sectionElem, pickerName + '-icon-picker', this.sim, iconPickers[pickerName], this);
    });
  }
}
