import { Stat } from '../api/newapi';
import { StatNames } from '../api/names';
import { Component } from './component.js';

export class CharacterStats extends Component {
  readonly stats: Array<Stat>;
  readonly valueElems: Array<HTMLTableCellElement>;

  constructor(parent: HTMLElement, stats: Array<Stat>) {
    super(parent, 'character-stats-root');
    this.stats = stats;

    const table = document.createElement('table');
    table.classList.add('character-stats-table');
    this.rootElem.appendChild(table);

    this.valueElems = [];
    this.stats.forEach(stat => {
      const row = document.createElement('tr');
      row.classList.add('character-stats-table-row');
      table.appendChild(row);

      const label = document.createElement('td');
      label.classList.add('character-stats-table-label');
      label.textContent = StatNames[stat];
      row.appendChild(label);

      const value = document.createElement('td');
      value.classList.add('character-stats-table-value');
      row.appendChild(value);
      value.textContent = String(Math.floor(Math.random() * 100));
      this.valueElems.push(value);
    });
  }
}
