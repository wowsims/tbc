import { Stat } from '../api/newapi';
import { StatNames } from '../api/names';
import { Sim } from '../sim';

import { Component } from './component';
import { NumberPicker } from './number_picker';

export class CustomStatsPicker extends Component {
  readonly stats: Array<Stat>;
  readonly statPickers: Array<NumberPicker>;

  constructor(parent: HTMLElement, sim: Sim, stats: Array<Stat>) {
    super(parent, 'custom-stats-root');
    this.stats = stats;

    const label = document.createElement('span');
    label.classList.add('custom-stats-label');
    label.textContent = 'Custom Stats';
    this.rootElem.appendChild(label);

    this.statPickers = this.stats.map(stat => new NumberPicker(this.rootElem, sim, {
      label: StatNames[stat],
      changedEvent: (sim: Sim) => sim.customStatsChangeEmitter,
      getValue: (sim: Sim) => sim.customStats[stat],
      setValue: (sim: Sim, newValue: number) => {
        const customStats = sim.customStats;
        customStats[stat] = newValue;
        sim.customStats = customStats;
      },
    }));
  }
}
