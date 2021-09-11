import { Spec } from '../api/newapi';
import { Stat } from '../api/newapi';
import { Actions } from '../components/actions.js';
import { CharacterStats } from '../components/character_stats.js';
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
      Stat.intellect,
      Stat.spell_power,
      Stat.nature_spell_power,
      Stat.spell_hit,
      Stat.spell_crit,
      Stat.spell_haste,
      Stat.mp5,
    ];
    const epReferenceStat = Stat.spell_power;

    const displayStats = [
      Stat.stamina,
      Stat.intellect,
      Stat.spell_power,
      Stat.spell_hit,
      Stat.spell_crit,
      Stat.spell_haste,
      Stat.mp5,
    ];

    const results = new Results();
    results.appendTo(this.parentElem.getElementsByClassName('default-results')[0]);

    const actions = new Actions(this.sim, results, epStats, epReferenceStat);
    actions.appendTo(this.parentElem.getElementsByClassName('default-actions')[0]);

    const characterStats = new CharacterStats(displayStats);
    characterStats.appendTo(this.parentElem.getElementsByClassName('default-stats')[0]);

    const racePicker = new RacePicker(this.sim);
    racePicker.appendTo(this.parentElem.getElementsByClassName('race-picker')[0]);
  }
}
