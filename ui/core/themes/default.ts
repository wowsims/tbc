import { Spec } from '../api/newapi';
import { Stat } from '../api/newapi';
import { Actions } from '../components/actions.js';
import { CharacterStats } from '../components/character_stats.js';
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

    const stats = [
      Stat.stamina,
      Stat.intellect,
      Stat.spell_power,
      Stat.spell_hit,
      Stat.spell_crit,
      Stat.spell_haste,
      Stat.mp5,
    ];

    const actions = new Actions(this.sim);
    actions.appendTo(this.parentElem.getElementsByClassName('default-actions')[0]);

    const characterStats = new CharacterStats(stats);
    characterStats.appendTo(this.parentElem.getElementsByClassName('default-stats')[0]);
  }
}
