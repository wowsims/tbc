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
        player.customStatsChangeEmitter.on(() => {
            this.statPickers.forEach(statPicker => {
                if (statPicker.getInputValue() > 0) {
                    statPicker.rootElem.classList.remove('negative');
                    statPicker.rootElem.classList.add('positive');
                }
                else if (statPicker.getInputValue() < 0) {
                    statPicker.rootElem.classList.remove('positive');
                    statPicker.rootElem.classList.add('negative');
                }
                else {
                    statPicker.rootElem.classList.remove('negative');
                    statPicker.rootElem.classList.remove('positive');
                }
            });
        });
    }
}
