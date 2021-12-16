import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';

import { ResultComponent, ResultComponentConfig, SimResultData } from './result_component.js';

const ALL_PLAYERS = -1;
const ALL_TARGETS = -1;

export class ResultsFilter extends ResultComponent {
	private currentPlayer: number;
	private currentTarget: number;

	readonly changeEmitter: TypedEvent<void>;

  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'results-filter-root';
    super(config);
		this.currentPlayer = ALL_PLAYERS;
		this.currentTarget = ALL_TARGETS;
		this.changeEmitter = new TypedEvent<void>();
  }

	getFilter(): SimResultFilter {
		return {
			player: this.currentPlayer == ALL_PLAYERS ? null : this.currentPlayer,
			target: this.currentTarget == ALL_TARGETS ? null : this.currentTarget,
		};
	}

	onSimResult(resultData: SimResultData) {
    this.rootElem.innerHTML = '';

		if (this.currentPlayer != ALL_PLAYERS && !resultData.result.getPlayerWithRaidIndex(this.currentPlayer)) {
			this.currentPlayer = ALL_PLAYERS;
			this.changeEmitter.emit();
		}

		const players = resultData.result.getPlayers();
		const playerPicker = new EnumPicker<ResultsFilter>(this.rootElem, this, {
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
			changedEvent: (resultsFilter: ResultsFilter) => resultsFilter.changeEmitter,
			getValue: (resultsFilter: ResultsFilter) => resultsFilter.currentPlayer,
			setValue: (resultsFilter: ResultsFilter, newValue: number) => {
				resultsFilter.currentPlayer = newValue;
				this.changeEmitter.emit();
			},
		});
	}
}
