import { StatWeightsRequest, StatWeightsResult, StatWeightValues, ProgressMetrics } from '/tbc/core/proto/api.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { statNames, statOrder } from '/tbc/core/proto_utils/names.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { Player } from '/tbc/core/player.js';
import { stDevToConf90 } from '/tbc/core/utils.js';
import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { getEnumValues } from '/tbc/core/utils.js';

import { Popup } from './popup.js';

declare var tippy: any;

export function addStatWeightsAction(simUI: IndividualSimUI<any>, epStats: Array<Stat>, epReferenceStat: Stat) {
	simUI.addAction('EP WEIGHTS', 'ep-weights-action', () => {
		new EpWeightsMenu(simUI, epStats, epReferenceStat);
	});
}

class EpWeightsMenu extends Popup {
	private readonly simUI: IndividualSimUI<any>;
	private readonly tableContainer: HTMLElement;
	private readonly tableBody: HTMLElement;
	private readonly tableHeader: HTMLElement;

	private statsType: string;
	private epStats: Array<Stat>;
	private epReferenceStat: Stat;

	constructor(simUI: IndividualSimUI<any>, epStats: Array<Stat>, epReferenceStat: Stat) {
		super(simUI.rootElem);
		this.simUI = simUI;
		this.statsType = 'ep';
		this.epStats = epStats;
		this.epReferenceStat = epReferenceStat;

		this.rootElem.classList.add('ep-weights-menu');
		this.rootElem.innerHTML = `
			<div class="ep-weights-actions">
				<button class="sim-button calc-weights">CALCULATE EP</button>
				<div class="ep-weights-options">
					<select class="ep-type-select">
						<option value="ep">EP</option>
						<option value="weight">DPS</option>
					</select>
				</div>
				<div class="show-all-stats-container">
				</div>
			</div>
			<div class="ep-weights-table">
				<table class="results-ep-table">
					<tbody id="ep-tbody">
						<tr>
							<th>Stat</th>
							<th class="type-weight"><span>DPS Weight</span><span class="col-action fa fa-copy"></span></th>
							<th class="type-ep"><span>DPS EP</span><span class="col-action fa fa-copy"></span></th>
							<th class="threat-metrics type-weight"><span>TPS Weight</span><span class="col-action fa fa-copy"></span></th>
							<th class="threat-metrics type-ep"><span>TPS EP</span><span class="col-action fa fa-copy"></span></th>
							<th class="threat-metrics type-weight"><span>DTPS Weight</span><span class="col-action fa fa-copy"></span></th>
							<th class="threat-metrics type-ep"><span>DTPS EP</span><span class="col-action fa fa-copy"></span></th>
							<th><span>Current EP</span><span class="col-action fa fa-recycle"></span></th>
						</tr>
					</tbody>
				</table>
			</div>
		`;

		this.tableContainer = this.rootElem.getElementsByClassName('ep-weights-table')[0] as HTMLElement;
		this.tableBody = this.rootElem.querySelector('#ep-tbody') as HTMLElement;
		this.tableHeader = this.rootElem.querySelector('#ep-tbody > tr') as HTMLElement;

		const updateType = () => {
			if (this.statsType == 'ep') {
				this.tableContainer.classList.remove('stats-type-weight');
				this.tableContainer.classList.add('stats-type-ep');
			} else {
				this.tableContainer.classList.add('stats-type-weight');
				this.tableContainer.classList.remove('stats-type-ep');
			}
		};

		const selectElem = this.rootElem.getElementsByClassName('ep-type-select')[0] as HTMLSelectElement;
		selectElem.addEventListener('input', event => {
			this.statsType = selectElem.value;
			updateType();
		});
		selectElem.value = this.statsType;
		updateType();

		const calcButton = this.rootElem.getElementsByClassName('calc-weights')[0] as HTMLElement;
		calcButton.addEventListener('click', async event => {
			this.simUI.setResultsPending();
			const iterations = this.simUI.sim.getIterations();
			const result = await this.simUI.player.computeStatWeights(TypedEvent.nextEventID(), this.epStats, this.epReferenceStat, (progress: ProgressMetrics) => {
				this.setSimProgress(progress);
			});
			this.simUI.prevEpIterations = iterations;
			this.simUI.prevEpSimResult = result;
			this.preprocessResults(result);
			this.updateTable(iterations, result);
		});

		const colActionButtons = Array.from(this.rootElem.getElementsByClassName('col-action')) as Array<HTMLSelectElement>;
		const makeUpdateWeights = (button: HTMLElement, tooltip: string, weightsFunc: () => Array<number>) => {
			tippy(button, {
				'content': tooltip,
				'allowHTML': true,
			});
			button.addEventListener('click', event => {
				this.simUI.player.setEpWeights(TypedEvent.nextEventID(), new Stats(weightsFunc()));
			});
		};
		makeUpdateWeights(colActionButtons[0], 'Copy to Current EP', () => this.getPrevSimResult().dps!.weights);
		makeUpdateWeights(colActionButtons[1], 'Copy to Current EP', () => this.getPrevSimResult().dps!.epValues);
		makeUpdateWeights(colActionButtons[2], 'Copy to Current EP', () => this.getPrevSimResult().tps!.weights);
		makeUpdateWeights(colActionButtons[3], 'Copy to Current EP', () => this.getPrevSimResult().tps!.epValues);
		makeUpdateWeights(colActionButtons[4], 'Copy to Current EP', () => this.getPrevSimResult().dtps!.weights);
		makeUpdateWeights(colActionButtons[5], 'Copy to Current EP', () => this.getPrevSimResult().dtps!.epValues);
		makeUpdateWeights(colActionButtons[6], 'Restore Default EP', () => this.simUI.individualConfig.defaults.epWeights.asArray());

		const showAllStatsContainer = this.rootElem.getElementsByClassName('show-all-stats-container')[0] as HTMLElement;
		new BooleanPicker(showAllStatsContainer, this, {
			label: 'Show All Stats',
			changedEvent: () => new TypedEvent(),
			getValue: () => this.tableContainer.classList.contains('show-all-stats'),
			setValue: (eventID: EventID, menu: EpWeightsMenu, newValue: boolean) => {
				if (newValue) {
					this.tableContainer.classList.add('show-all-stats');
				} else {
					this.tableContainer.classList.remove('show-all-stats');
				}
			},
		});

		this.updateTable(this.simUI.prevEpIterations || 1, this.getPrevSimResult());

		this.addCloseButton();
	}

	setSimProgress(progress: ProgressMetrics) {
		this.simUI.setResultsContent(`
  <div class="results-sim">
  			<div class=""> ${progress.completedSims} / ${progress.totalSims}<br>simulations complete</div>
  			<div class="">
				${progress.completedIterations} / ${progress.totalIterations}<br>iterations complete
			</div>
  </div>
`);
	}

	private preprocessResults(result: StatWeightsResult) {
		// Values for a school's power should never exceed the value for regular spell power.
		result.dps!.epValues.forEach((value, index) => {
			if (index == Stat.StatArcaneSpellPower ||
				index == Stat.StatFireSpellPower ||
				index == Stat.StatFrostSpellPower ||
				index == Stat.StatHolySpellPower ||
				index == Stat.StatNatureSpellPower ||
				index == Stat.StatShadowSpellPower) {
				if (value > result.dps!.epValues[Stat.StatSpellPower]) {
					const diff = value - result.dps!.epValues[Stat.StatSpellPower];
					result.dps!.epValues[index] = result.dps!.epValues[Stat.StatSpellPower];
					result.dps!.epValuesStdev[index] -= diff;
					const wdiff = result.dps!.weights[index] - result.dps!.weights[Stat.StatSpellPower];
					result.dps!.weights[index] = result.dps!.weights[Stat.StatSpellPower];
					result.dps!.weightsStdev[index] -= wdiff;
				}
			}
		});
	}

	private updateTable(iterations: number, result: StatWeightsResult) {
		this.tableHeader.remove();
		this.tableBody.innerHTML = '';
		this.tableBody.appendChild(this.tableHeader);

		const allStats = statOrder.filter(stat => ![Stat.StatMana, Stat.StatEnergy, Stat.StatRage].includes(stat));
		allStats.forEach(stat => {
			const row = this.makeTableRow(stat, iterations, result);
			if (!this.epStats.includes(stat)) {
				row.classList.add('non-ep-stat');
			}
			this.tableBody.appendChild(row);
		});
	}

	private makeTableRow(stat: Stat, iterations: number, result: StatWeightsResult): HTMLElement {
		const row = document.createElement('tr');
		row.innerHTML = `
			<td>${statNames[stat]}</td>
			<td class="stdev-cell type-weight"><span>${result.dps!.weights[stat].toFixed(2)}</span><span>${stDevToConf90(result.dps!.weightsStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell type-ep"><span>${result.dps!.epValues[stat].toFixed(2)}</span><span>${stDevToConf90(result.dps!.epValuesStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell threat-metrics type-weight"><span>${result.tps!.weights[stat].toFixed(2)}</span><span>${stDevToConf90(result.tps!.weightsStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell threat-metrics type-ep"><span>${result.tps!.epValues[stat].toFixed(2)}</span><span>${stDevToConf90(result.tps!.epValuesStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell threat-metrics type-weight"><span>${result.dtps!.weights[stat].toFixed(2)}</span><span>${stDevToConf90(result.dtps!.weightsStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell threat-metrics type-ep"><span>${result.dtps!.epValues[stat].toFixed(2)}</span><span>${stDevToConf90(result.dtps!.epValuesStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="current-ep"></td>
		`;

		const currentEpCell = row.querySelector('.current-ep') as HTMLElement;
		new NumberPicker(currentEpCell, this.simUI.player, {
			float: true,
			changedEvent: (player: Player<any>) => player.epWeightsChangeEmitter,
			getValue: (player: Player<any>) => player.getEpWeights().getStat(stat),
			setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
				const epWeights = player.getEpWeights().withStat(stat, newValue);
				player.setEpWeights(eventID, epWeights);
			},
		});

		return row;
	}

	private getPrevSimResult(): StatWeightsResult {
		return this.simUI.prevEpSimResult || StatWeightsResult.create({
			dps: {
				weights: new Stats().asArray(),
				weightsStdev: new Stats().asArray(),
				epValues: new Stats().asArray(),
				epValuesStdev: new Stats().asArray(),
			},
			tps: {
				weights: new Stats().asArray(),
				weightsStdev: new Stats().asArray(),
				epValues: new Stats().asArray(),
				epValuesStdev: new Stats().asArray(),
			},
			dtps: {
				weights: new Stats().asArray(),
				weightsStdev: new Stats().asArray(),
				epValues: new Stats().asArray(),
				epValuesStdev: new Stats().asArray(),
			},
		});
	}
}
