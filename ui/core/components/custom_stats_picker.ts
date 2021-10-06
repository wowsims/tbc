import { Stat } from '/tbc/core/proto/common.js';
import { statNames } from '/tbc/core/api/names.js';
import { Stats } from '/tbc/core/api/stats.js';
import { Sim } from '/tbc/core/sim.js';

import { Component } from './component.js';
import { NumberPicker } from './number_picker.js';

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
