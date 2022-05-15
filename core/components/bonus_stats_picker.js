import { statNames, statOrder } from '/tbc/core/proto_utils/names.js';
import { Component } from './component.js';
import { NumberPicker } from './number_picker.js';
export class BonusStatsPicker extends Component {
    constructor(parent, player, stats) {
        super(parent, 'bonus-stats-root');
        this.stats = stats;
        const label = document.createElement('span');
        label.classList.add('bonus-stats-label');
        label.textContent = 'Bonus Stats';
        tippy(label, {
            'content': 'Extra stats to add on top of gear, buffs, etc.',
            'allowHTML': true,
        });
        this.rootElem.appendChild(label);
        this.statPickers = statOrder.filter(stat => this.stats.includes(stat)).map(stat => new NumberPicker(this.rootElem, player, {
            label: statNames[stat],
            changedEvent: (player) => player.bonusStatsChangeEmitter,
            getValue: (player) => player.getBonusStats().getStat(stat),
            setValue: (eventID, player, newValue) => {
                const bonusStats = player.getBonusStats().withStat(stat, newValue);
                player.setBonusStats(eventID, bonusStats);
            },
        }));
        player.bonusStatsChangeEmitter.on(() => {
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
