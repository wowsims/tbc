import { statNames } from '/tbc/core/proto_utils/names.js';
import { Component } from './component.js';
import { NumberPicker } from './number_picker.js';
export class CustomStatsPicker extends Component {
    constructor(parent, player, stats) {
        super(parent, 'custom-stats-root');
        this.stats = stats;
        const label = document.createElement('span');
        label.classList.add('custom-stats-label');
        label.textContent = 'Custom Stats';
        this.rootElem.appendChild(label);
        this.statPickers = this.stats.map(stat => new NumberPicker(this.rootElem, player, {
            label: statNames[stat],
            changedEvent: (player) => player.customStatsChangeEmitter,
            getValue: (player) => player.getCustomStats().getStat(stat),
            setValue: (player, newValue) => {
                const customStats = player.getCustomStats().withStat(stat, newValue);
                player.setCustomStats(customStats);
            },
        }));
    }
}
