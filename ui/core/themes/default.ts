import { CharacterStats } from '../components/character_stats.js';
import { Theme } from './theme.js';

export class DefaultTheme extends Theme {
  constructor(parentElem: HTMLElement) {
    super(parentElem)

    const characterStats = new CharacterStats();
    this.parentElem.appendChild(characterStats.getRootElement());
  }
}
