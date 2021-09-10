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
  constructor(parentElem: HTMLElement) {
    super(parentElem)

    this.parentElem.innerHTML = layoutHTML;

    const actions = new Actions();
    actions.appendTo(this.parentElem.getElementsByClassName('default-actions')[0]);

    const characterStats = new CharacterStats([
      Stat.stamina,
      Stat.intellect,
    ]);
    characterStats.appendTo(this.parentElem.getElementsByClassName('default-stats')[0]);
  }
}
