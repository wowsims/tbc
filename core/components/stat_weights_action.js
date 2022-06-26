import { StatWeightsResult } from '/tbc/core/proto/api.js';
import { GemColor } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { gemMatchesSocket, getMetaGemCondition } from '/tbc/core/proto_utils/gems.js';
import { statNames, statOrder } from '/tbc/core/proto_utils/names.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { stDevToConf90 } from '/tbc/core/utils.js';
import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { ResultsViewer } from '/tbc/core/components/results_viewer.js';
import { getEnumValues, sum } from '/tbc/core/utils.js';
import { Popup } from './popup.js';
export function addStatWeightsAction(simUI, epStats, epReferenceStat) {
    simUI.addAction('STAT WEIGHTS', 'ep-weights-action', () => {
        new EpWeightsMenu(simUI, epStats, epReferenceStat);
    });
}
class EpWeightsMenu extends Popup {
    constructor(simUI, epStats, epReferenceStat) {
        super(simUI.rootElem);
        this.simUI = simUI;
        this.statsType = 'ep';
        this.epStats = epStats;
        this.epReferenceStat = epReferenceStat;
        this.rootElem.classList.add('ep-weights-menu');
        this.rootElem.innerHTML = `
			<div class="ep-weights-header">
				<div class="ep-weights-actions">
					<button class="sim-button calc-weights">CALCULATE</button>
				</div>
				<div class="ep-weights-results">
				</div>
			</div>
			<div class="stats-controls-row">
				<div class="ep-weights-options">
					<select class="ep-type-select">
						<option value="ep">EP</option>
						<option value="weight">Weights</option>
					</select>
				</div>
				<div class="show-all-stats-container">
				</div>
				<button class="sim-button optimize-gems">OPTIMIZE GEMS</button>
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
        this.tableContainer = this.rootElem.getElementsByClassName('ep-weights-table')[0];
        this.tableBody = this.rootElem.querySelector('#ep-tbody');
        this.tableHeader = this.rootElem.querySelector('#ep-tbody > tr');
        const resultsViewerElem = this.rootElem.getElementsByClassName('ep-weights-results')[0];
        this.resultsViewer = new ResultsViewer(resultsViewerElem);
        const updateType = () => {
            if (this.statsType == 'ep') {
                this.tableContainer.classList.remove('stats-type-weight');
                this.tableContainer.classList.add('stats-type-ep');
            }
            else {
                this.tableContainer.classList.add('stats-type-weight');
                this.tableContainer.classList.remove('stats-type-ep');
            }
        };
        const selectElem = this.rootElem.getElementsByClassName('ep-type-select')[0];
        selectElem.addEventListener('input', event => {
            this.statsType = selectElem.value;
            updateType();
        });
        selectElem.value = this.statsType;
        updateType();
        const calcButton = this.rootElem.getElementsByClassName('calc-weights')[0];
        calcButton.addEventListener('click', async (event) => {
            this.resultsViewer.setPending();
            const iterations = this.simUI.sim.getIterations();
            const result = await this.simUI.player.computeStatWeights(TypedEvent.nextEventID(), this.epStats, this.epReferenceStat, (progress) => {
                this.setSimProgress(progress);
            });
            this.resultsViewer.hideAll();
            this.simUI.prevEpIterations = iterations;
            this.simUI.prevEpSimResult = result;
            this.preprocessResults(result);
            this.updateTable(iterations, result);
        });
        const colActionButtons = Array.from(this.rootElem.getElementsByClassName('col-action'));
        const makeUpdateWeights = (button, labelTooltip, tooltip, weightsFunc) => {
            tippy(button.previousSibling, {
                'content': labelTooltip,
                'allowHTML': true,
            });
            tippy(button, {
                'content': tooltip,
                'allowHTML': true,
            });
            button.addEventListener('click', event => {
                this.simUI.player.setEpWeights(TypedEvent.nextEventID(), new Stats(weightsFunc()));
            });
        };
        makeUpdateWeights(colActionButtons[0], 'Per-point increase in DPS (Damage Per Second) for each stat.', 'Copy to Current EP', () => this.getPrevSimResult().dps.weights);
        makeUpdateWeights(colActionButtons[1], `EP (Equivalency Points) for DPS (Damage Per Second) for each stat. Normalized by ${statNames[this.epReferenceStat]}.`, 'Copy to Current EP', () => this.getPrevSimResult().dps.epValues);
        makeUpdateWeights(colActionButtons[2], 'Per-point increase in TPS (Threat Per Second) for each stat.', 'Copy to Current EP', () => this.getPrevSimResult().tps.weights);
        makeUpdateWeights(colActionButtons[3], `EP (Equivalency Points) for TPS (Threat Per Second) for each stat. Normalized by ${statNames[this.epReferenceStat]}.`, 'Copy to Current EP', () => this.getPrevSimResult().tps.epValues);
        makeUpdateWeights(colActionButtons[4], 'Per-point increase in DTPS (Damage Taken Per Second) for each stat.', 'Copy to Current EP', () => this.getPrevSimResult().dtps.weights);
        makeUpdateWeights(colActionButtons[5], `EP (Equivalency Points) for DTPS (Damage Taken Per Second) for each stat. Normalized by ${statNames[Stat.StatArmor]}.`, 'Copy to Current EP', () => this.getPrevSimResult().dtps.epValues);
        makeUpdateWeights(colActionButtons[6], 'Current EP Weights. Used to sort the gear selector menus.', 'Restore Default EP', () => this.simUI.individualConfig.defaults.epWeights.asArray());
        const showAllStatsContainer = this.rootElem.getElementsByClassName('show-all-stats-container')[0];
        new BooleanPicker(showAllStatsContainer, this, {
            label: 'Show All Stats',
            changedEvent: () => new TypedEvent(),
            getValue: () => this.tableContainer.classList.contains('show-all-stats'),
            setValue: (eventID, menu, newValue) => {
                if (newValue) {
                    this.tableContainer.classList.add('show-all-stats');
                }
                else {
                    this.tableContainer.classList.remove('show-all-stats');
                }
                this.applyAlternatingColors();
            },
        });
        this.updateTable(this.simUI.prevEpIterations || 1, this.getPrevSimResult());
        const optimizeGemsButton = this.rootElem.getElementsByClassName('optimize-gems')[0];
        tippy(optimizeGemsButton, {
            'content': `
				<p>Optimizes equipped gems to maximize EP, based on the values in <b>Current EP</b>.</p>
				<p>WARNING: Ignores unique gems and does not pick the meta gem or ensure its condition is met.</p>
			`,
            'allowHTML': true,
        });
        optimizeGemsButton.addEventListener('click', event => this.optimizeGems(TypedEvent.nextEventID()));
        this.addCloseButton();
    }
    setSimProgress(progress) {
        this.resultsViewer.setContent(`
  <div class="results-sim">
  			<div class=""> ${progress.completedSims} / ${progress.totalSims}<br>simulations complete</div>
  			<div class="">
				${progress.completedIterations} / ${progress.totalIterations}<br>iterations complete
			</div>
  </div>
`);
    }
    preprocessResults(result) {
        // Values for a school's power should never exceed the value for regular spell power.
        result.dps.epValues.forEach((value, index) => {
            if (index == Stat.StatArcaneSpellPower ||
                index == Stat.StatFireSpellPower ||
                index == Stat.StatFrostSpellPower ||
                index == Stat.StatHolySpellPower ||
                index == Stat.StatNatureSpellPower ||
                index == Stat.StatShadowSpellPower) {
                if (value > result.dps.epValues[Stat.StatSpellPower]) {
                    const diff = value - result.dps.epValues[Stat.StatSpellPower];
                    result.dps.epValues[index] = result.dps.epValues[Stat.StatSpellPower];
                    result.dps.epValuesStdev[index] -= diff;
                    const wdiff = result.dps.weights[index] - result.dps.weights[Stat.StatSpellPower];
                    result.dps.weights[index] = result.dps.weights[Stat.StatSpellPower];
                    result.dps.weightsStdev[index] -= wdiff;
                }
            }
        });
    }
    updateTable(iterations, result) {
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
        this.applyAlternatingColors();
    }
    makeTableRow(stat, iterations, result) {
        const row = document.createElement('tr');
        row.innerHTML = `
			<td>${statNames[stat]}</td>
			<td class="stdev-cell type-weight"><span>${result.dps.weights[stat].toFixed(2)}</span><span>${stDevToConf90(result.dps.weightsStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell type-ep"><span>${result.dps.epValues[stat].toFixed(2)}</span><span>${stDevToConf90(result.dps.epValuesStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell threat-metrics type-weight"><span>${result.tps.weights[stat].toFixed(2)}</span><span>${stDevToConf90(result.tps.weightsStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell threat-metrics type-ep"><span>${result.tps.epValues[stat].toFixed(2)}</span><span>${stDevToConf90(result.tps.epValuesStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell threat-metrics type-weight"><span>${result.dtps.weights[stat].toFixed(2)}</span><span>${stDevToConf90(result.dtps.weightsStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="stdev-cell threat-metrics type-ep"><span>${result.dtps.epValues[stat].toFixed(2)}</span><span>${stDevToConf90(result.dtps.epValuesStdev[stat], iterations).toFixed(2)}</span></td>
			<td class="current-ep"></td>
		`;
        const currentEpCell = row.querySelector('.current-ep');
        new NumberPicker(currentEpCell, this.simUI.player, {
            float: true,
            changedEvent: (player) => player.epWeightsChangeEmitter,
            getValue: (player) => player.getEpWeights().getStat(stat),
            setValue: (eventID, player, newValue) => {
                const epWeights = player.getEpWeights().withStat(stat, newValue);
                player.setEpWeights(eventID, epWeights);
            },
        });
        return row;
    }
    applyAlternatingColors() {
        Array.from(this.tableBody.childNodes)
            .filter(row => window.getComputedStyle(row).getPropertyValue('display') != 'none')
            .forEach((row, i) => {
            if (i % 2 == 0) {
                row.classList.remove('odd');
            }
            else {
                row.classList.add('odd');
            }
        });
    }
    getPrevSimResult() {
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
    optimizeGems(eventID) {
        let epWeights = this.simUI.player.getEpWeights();
        // Replace 0 weights with a very tiny value, so we always prefer to take free stats even if the user gave a 0 weight.
        epWeights = new Stats(epWeights.asArray().map(w => w == 0 ? 1e-8 : w));
        const gemColors = getEnumValues(GemColor);
        const allGems = this.simUI.player.getGems().filter(gem => !gem.unique && gem.phase <= this.simUI.sim.getPhase());
        // Best gem when we need a gem of a specific color.
        const bestGemForColor = gemColors.map(color => null);
        const bestGemForColorEP = gemColors.map(color => 0);
        // Best gem when we need to match a socket to activate a bonus.
        const bestGemForSocket = bestGemForColor.slice();
        const bestGemForSocketEP = bestGemForColorEP.slice();
        // The single best gem, when color doesn't matter.
        let bestGem = allGems[0];
        let bestGemEP = 0;
        allGems.forEach(gem => {
            const gemEP = new Stats(gem.stats).computeEP(epWeights);
            if (gemEP > bestGemForColorEP[gem.color]) {
                bestGemForColorEP[gem.color] = gemEP;
                bestGemForColor[gem.color] = gem;
                if (gem.color != GemColor.GemColorMeta && gemEP > bestGemEP) {
                    bestGemEP = gemEP;
                    bestGem = gem;
                }
            }
            gemColors.forEach(socketColor => {
                if (gemMatchesSocket(gem, socketColor) && gemEP > bestGemForSocketEP[socketColor]) {
                    bestGemForSocketEP[socketColor] = gemEP;
                    bestGemForSocket[socketColor] = gem;
                }
            });
        });
        let gear = this.simUI.player.getGear();
        const items = gear.asMap();
        const socketBonusEPs = Object.values(items).map(item => item != null ? new Stats(item.item.socketBonus).computeEP(epWeights) : 0);
        // Start by optimally filling all items, ignoring meta condition.
        Object.entries(items).forEach(([itemSlot, equippedItem], i) => {
            if (equippedItem == null) {
                return;
            }
            const item = equippedItem.item;
            // Compare whether its better to match sockets + get socket bonus, or just use best gems.
            const bestGemEPNotMatchingSockets = sum(item.gemSockets.map(socketColor => socketColor == GemColor.GemColorMeta ? 0 : bestGemEP));
            const bestGemEPMatchingSockets = socketBonusEPs[i] + sum(item.gemSockets.map(socketColor => socketColor == GemColor.GemColorMeta ? 0 : bestGemForSocketEP[socketColor]));
            if (bestGemEPNotMatchingSockets > bestGemEPMatchingSockets) {
                item.gemSockets.forEach((socketColor, i) => {
                    if (socketColor != GemColor.GemColorMeta) {
                        equippedItem = equippedItem.withGem(bestGem, i);
                    }
                });
            }
            else {
                item.gemSockets.forEach((socketColor, i) => {
                    if (socketColor != GemColor.GemColorMeta) {
                        equippedItem = equippedItem.withGem(bestGemForSocket[socketColor], i);
                    }
                });
            }
            items[Number(itemSlot)] = equippedItem;
        });
        gear = new Gear(items);
        // Now make adjustments to satisfy meta condition.
        const metaGem = gear.getMetaGem();
        if (metaGem != null) {
            const condition = getMetaGemCondition(metaGem.id);
            // TODO: Satisfy condition. Not implementing this since we're about to move
            // to wrath which doesn't have meta conditions.
        }
        // Apply the new gear.
        this.simUI.player.setGear(eventID, gear);
    }
}
