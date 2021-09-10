import { Component } from './component.js';

export class CharacterStats extends Component {
  readonly rootElem: HTMLDivElement;
  readonly stats: Array<string>;
  readonly valueElems: Array<HTMLTableCellElement>;

  constructor() {
    super();

    this.stats = [
      'Stamina',
      'Intellect',
      'Spirit',
      'Spell Power',
      'Nature Power',
      'Hit',
      'Crit',
      'Haste',
    ];

    this.rootElem = document.createElement('div');
    this.rootElem.classList.add('character-stats-root');

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
      label.textContent = stat;
      row.appendChild(label);

      const value = document.createElement('td');
      value.classList.add('character-stats-table-value');
      row.appendChild(value);
      value.textContent = String(Math.floor(Math.random() * 100));
      this.valueElems.push(value);
    });
  }

  getRootElement() {
    return this.rootElem;
  }
}
