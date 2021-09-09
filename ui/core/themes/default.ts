import { CharacterStats } from '../components/character_stats.js';
import { Theme } from './theme.js';

const layoutHTML = `
<div class="default-root">
  <section class="default-sidebar">
    <div class="default-stats">
    </div>
    <div class="default-results">
    </div>
    <div class="default-actions">
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

    const characterStats = new CharacterStats();
    characterStats.appendTo(this.parentElem.getElementsByClassName('default-stats')[0]);
  }
}
