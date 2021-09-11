import { Spec } from '../api/newapi';
import { Stat } from '../api/newapi';
import { Actions } from '../components/actions.js';
import { CharacterStats } from '../components/character_stats.js';
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
  </div>
  <div class="default-main">
  </div>
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
  }
}
