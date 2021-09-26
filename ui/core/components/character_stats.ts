import { Stat } from '../api/common.js';
import { statNames } from '../api/names.js';
import { Stats } from '../api/stats.js';
import { Sim } from '../sim.js';

import { Component } from './component.js';

export class CharacterStats extends Component {
  readonly stats: Array<Stat>;
  readonly valueElems: Array<HTMLTableCellElement>;

  constructor(parent: HTMLElement, stats: Array<Stat>, sim: Sim<any>) {
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
      label.textContent = statNames[stat];
      row.appendChild(label);

      const value = document.createElement('td');
      value.classList.add('character-stats-table-value');
      row.appendChild(value);
      this.valueElems.push(value);
    });

		this.updateStats(new Stats());
		sim.characterStatsEmitter.on(computeStatsResult => {
			this.updateStats(new Stats(computeStatsResult.finalStats));
		});
  }

	private updateStats(newStats: Stats) {
		this.stats.forEach((stat, idx) => {
			this.valueElems[idx].textContent = String(Math.round(newStats.getStat(stat)));
		});
	}
}
