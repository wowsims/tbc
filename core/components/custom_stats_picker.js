import { statNames } from '/tbc/core/proto_utils/names.js';
import { Component } from './component.js';
import { NumberPicker } from './number_picker.js';
export class CustomStatsPicker extends Component {
    constructor(parent, sim, stats) {
        super(parent, 'custom-stats-root');
        this.stats = stats;
        const label = document.createElement('span');
        label.classList.add('custom-stats-label');
        label.textContent = 'Custom Stats';
        this.rootElem.appendChild(label);
        this.statPickers = this.stats.map(stat => new NumberPicker(this.rootElem, sim, {
            label: statNames[stat],
            changedEvent: (sim) => sim.customStatsChangeEmitter,
            getValue: (sim) => sim.getCustomStats().getStat(stat),
            setValue: (sim, newValue) => {
                const customStats = sim.getCustomStats().withStat(stat, newValue);
                sim.setCustomStats(customStats);
            },
        }));
    }
}
