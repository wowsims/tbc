import { Spec } from '../api/newapi';
import { Stat } from '../api/newapi';
import { Actions } from '../components/actions.js';
import { CharacterStats } from '../components/character_stats.js';
import { GearPicker } from '../components/gear_picker.js';
import { RacePicker } from '../components/race_picker.js';
import { Results } from '../components/results.js';

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
    </ul>
    <div class="tab-content">
      <div id="gear-tab" class="tab-pane fade in active">
        <div class="race-picker">
        </div>
        <div class="gear-picker">
        </div>
      </div>
      <div id="settings-tab" class="tab-pane fade">
      </div>
    </div>
  </section>
</div>
`;

export class DefaultTheme extends Theme {
  constructor(parentElem: HTMLElement, spec: Spec) {
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

    const racePicker = new RacePicker(this.parentElem.getElementsByClassName('race-picker')[0] as HTMLElement, this.sim);
    const gearPicker = new GearPicker(this.parentElem.getElementsByClassName('gear-picker')[0] as HTMLElement, this.sim);
  }
}
