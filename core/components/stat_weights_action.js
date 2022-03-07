import { Stat } from '/tbc/core/proto/common.js';
import { statNames } from '/tbc/core/proto_utils/names.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { stDevToConf90 } from '/tbc/core/utils.js';
export function addStatWeightsAction(simUI, epStats, epReferenceStat) {
    const resultsManager = new StatWeightsResultsManager(simUI);
    simUI.addAction('EP WEIGHTS', 'ep-weights-action', async () => {
        simUI.setResultsPending();
        const iterations = simUI.sim.getIterations();
        const result = await simUI.player.computeStatWeights(TypedEvent.nextEventID(), epStats, epReferenceStat, (progress) => {
            resultsManager.setSimProgress(progress);
        });
        resultsManager.setSimResult(iterations, epStats, epReferenceStat, result);
    });
}
class StatWeightsResultsManager {
    constructor(simUI) {
        this.simUI = simUI;
        this.statsType = 'ep';
    }
    setSimProgress(progress) {
        this.simUI.setResultsContent(`
  <div class="results-sim">
  			<div class=""> ${progress.completedSims} / ${progress.totalSims}<br>simulations complete</div>
  			<div class="">
				${progress.completedIterations} / ${progress.totalIterations}<br>iterations complete
			</div>
  </div>
`);
    }
    setSimResult(iterations, epStats, epReferenceStat, result) {
        if (epReferenceStat == Stat.StatSpellPower) {
            result.epValues.forEach((value, index) => {
                if (index == Stat.StatArcaneSpellPower ||
                    index == Stat.StatFireSpellPower ||
                    index == Stat.StatFrostSpellPower ||
                    index == Stat.StatHolySpellPower ||
                    index == Stat.StatNatureSpellPower ||
                    index == Stat.StatShadowSpellPower) {
                    if (value > result.epValues[epReferenceStat]) {
                        const diff = value - result.epValues[epReferenceStat];
                        result.epValues[index] = result.epValues[epReferenceStat];
                        result.epValuesStdev[index] -= diff;
                        const wdiff = result.weights[index] - result.weights[epReferenceStat];
                        result.weights[index] = result.weights[epReferenceStat];
                        result.weightsStdev[index] -= wdiff;
                    }
                }
            });
        }
        this.simUI.setResultsContent(`
			<div class="results-ep">
				<select class="results-ep-type-select">
					<option value="ep">EP</option>
					<option value="weight">DPS</option>
				</select>
				<table class="results-ep-table">
				` + epStats.map(stat => `<tr>
							<td>${statNames[stat]}:</td>
							<td>${result.weights[stat].toFixed(2)}</td>
							<td>${stDevToConf90(result.weightsStdev[stat], iterations).toFixed(2)}</td>
							<td>${result.epValues[stat].toFixed(2)}</td>
							<td>${stDevToConf90(result.epValuesStdev[stat], iterations).toFixed(2)}</td>
							</tr>`)
            .join('')
            + '</table></div>');
        const epElem = this.simUI.resultsContentElem.getElementsByClassName('results-ep')[0];
        const setType = (type) => {
            if (type == 'ep') {
                epElem.classList.remove('stats-type-dps');
                epElem.classList.add('stats-type-ep');
            }
            else {
                epElem.classList.add('stats-type-dps');
                epElem.classList.remove('stats-type-ep');
            }
            this.statsType = type;
        };
        const selectElem = epElem.getElementsByClassName('results-ep-type-select')[0];
        selectElem.addEventListener('input', event => {
            setType(selectElem.value);
        });
        selectElem.value = this.statsType;
        setType(this.statsType);
    }
}
