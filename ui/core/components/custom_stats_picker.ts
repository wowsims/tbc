import { Stat } from '../api/newapi';
import { statNames } from '../api/names';
import { Stats } from '../api/stats';
import { Sim } from '../sim';

import { Component } from './component';
import { NumberPicker } from './number_picker';

export class CustomStatsPicker extends Component {
  readonly stats: Array<Stat>;
  readonly statPickers: Array<NumberPicker>;

  constructor(parent: HTMLElement, sim: Sim<any>, stats: Array<Stat>) {
    super(parent, 'custom-stats-root');
    this.stats = stats;

    const label = document.createElement('span');
    label.classList.add('custom-stats-label');
    label.textContent = 'Custom Stats';
    this.rootElem.appendChild(label);

    this.statPickers = this.stats.map(stat => new NumberPicker(this.rootElem, sim, {
      label: statNames[stat],
      changedEvent: (sim: Sim<any>) => sim.customStatsChangeEmitter,
      getValue: (sim: Sim<any>) => sim.getCustomStats().getStat(stat),
      setValue: (sim: Sim<any>, newValue: number) => {
        const customStats = sim.getCustomStats().withStat(stat, newValue);
        sim.setCustomStats(customStats);
      },
    }));
  }
}
