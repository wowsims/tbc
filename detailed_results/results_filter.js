import { TypedEvent } from '/tbc/core/typed_event.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { ResultComponent } from './result_component.js';
const ALL_PLAYERS = -1;
const ALL_TARGETS = -1;
export class ResultsFilter extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'results-filter-root';
        super(config);
        this.currentPlayer = ALL_PLAYERS;
        this.currentTarget = ALL_TARGETS;
        this.changeEmitter = new TypedEvent();
    }
    getFilter() {
        return {
            player: this.currentPlayer == ALL_PLAYERS ? null : this.currentPlayer,
            target: this.currentTarget == ALL_TARGETS ? null : this.currentTarget,
        };
    }
    onSimResult(resultData) {
        this.rootElem.innerHTML = '';
        if (this.currentPlayer != ALL_PLAYERS && !resultData.result.getPlayerWithRaidIndex(this.currentPlayer)) {
            this.currentPlayer = ALL_PLAYERS;
            this.changeEmitter.emit();
        }
        const players = resultData.result.getPlayers();
        const playerPicker = new EnumPicker(this.rootElem, this, {
            extraCssClasses: [
                'player-filter',
            ],
            values: [
                { name: 'All Players', value: ALL_PLAYERS },
            ].concat(players.map(player => {
                return {
                    name: `${player.name} (#${player.raidIndex + 1})`,
                    value: player.raidIndex,
                };
            })),
            changedEvent: (resultsFilter) => resultsFilter.changeEmitter,
            getValue: (resultsFilter) => resultsFilter.currentPlayer,
            setValue: (resultsFilter, newValue) => {
                resultsFilter.currentPlayer = newValue;
                this.changeEmitter.emit();
            },
        });
    }
}
