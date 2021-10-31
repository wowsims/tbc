import { Stat } from '/tbc/core/proto/common.js';
import { statNames } from '/tbc/core/proto_utils/names.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Player } from '/tbc/core/player.js';

import { Component } from './component.js';
import { NumberPicker } from './number_picker.js';

export class CustomStatsPicker extends Component {
  readonly stats: Array<Stat>;
  readonly statPickers: Array<NumberPicker<Player<any>>>;

  constructor(parent: HTMLElement, player: Player<any>, stats: Array<Stat>) {
    super(parent, 'custom-stats-root');
    this.stats = stats;

    const label = document.createElement('span');
    label.classList.add('custom-stats-label');
    label.textContent = 'Custom Stats';
    this.rootElem.appendChild(label);

    this.statPickers = this.stats.map(stat => new NumberPicker(this.rootElem, player, {
      label: statNames[stat],
      changedEvent: (player: Player<any>) => player.customStatsChangeEmitter,
      getValue: (player: Player<any>) => player.getCustomStats().getStat(stat),
      setValue: (player: Player<any>, newValue: number) => {
        const customStats = player.getCustomStats().withStat(stat, newValue);
        player.setCustomStats(customStats);
      },
    }));
  }
}
