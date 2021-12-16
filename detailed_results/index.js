import { TypedEvent } from '/tbc/core/typed_event.js';
import { SimResult } from '/tbc/core/proto_utils/sim_result.js';
import { ResultsFilter } from './results_filter.js';
import { CastMetrics } from './cast_metrics.js';
import { SpellMetrics } from './cast_metrics.js';
import { AuraMetrics } from './aura_metrics.js';
import { DpsHistogram } from './dps_histogram.js';
import { DpsResult } from './dps_result.js';
import { PercentOom } from './percent_oom.js';
const urlParams = new URLSearchParams(window.location.search);
if (urlParams.has('mainBgColor')) {
    document.body.style.setProperty('--main-bg-color', urlParams.get('mainBgColor'));
}
if (urlParams.has('mainTextColor')) {
    document.body.style.setProperty('--main-text-color', urlParams.get('mainTextColor'));
}
const isIndividualSim = urlParams.has('isIndividualSim');
if (isIndividualSim) {
    document.body.classList.add('individual-sim');
}
const colorSettings = {
    mainTextColor: document.body.style.getPropertyValue('--main-text-color'),
};
Chart.defaults.color = colorSettings.mainTextColor;
const layoutHTML = `
<div class="dr-root">
	<ul class="dr-toolbar nav nav-tabs">
		<li class="results-filter">
		</li>
		<li class="tabs-filler">
		</li>
		<li class="dr-tab-tab active"><a data-toggle="tab" href="#damageTab">Damage</a></li>
		<li class="dr-tab-tab"><a data-toggle="tab" href="#buffsTab">Buffs</a></li>
		<li class="dr-tab-tab"><a data-toggle="tab" href="#debuffsTab">Debuffs</a></li>
		<li class="dr-tab-tab"><a data-toggle="tab" href="#castsTab">Casts</a></li>
	</ul>
	<div class="tab-content">
		<div id="damageTab" class="tab-pane dr-tab-content damage-content fade active in">
			<div class="dr-row topline-results">
			</div>
			<div class="dr-row">
				<div class="spell-metrics">
				</div>
			</div>
			<div class="dr-row dps-histogram">
			</div>
		</div>
		<div id="buffsTab" class="tab-pane dr-tab-content buffs-content fade">
			<div class="dr-row">
				<div class="buff-aura-metrics">
				</div>
			</div>
		</div>
		<div id="debuffsTab" class="tab-pane dr-tab-content debuffs-content fade">
			<div class="dr-row">
				<div class="debuff-aura-metrics">
				</div>
			</div>
		</div>
		<div id="castsTab" class="tab-pane dr-tab-content casts-content fade">
			<div class="dr-row">
				<div class="cast-metrics">
				</div>
			</div>
		</div>
	</div>
</div>
`;
document.body.innerHTML = layoutHTML;
const resultsEmitter = new TypedEvent();
const resultsFilter = new ResultsFilter({
    parent: document.body.getElementsByClassName('results-filter')[0],
    resultsEmitter: resultsEmitter,
    colorSettings: colorSettings,
});
const toplineResultsDiv = document.body.getElementsByClassName('topline-results')[0];
const dpsResult = new DpsResult({ parent: toplineResultsDiv, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const percentOom = new PercentOom({ parent: toplineResultsDiv, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const castMetrics = new CastMetrics({ parent: document.body.getElementsByClassName('cast-metrics')[0], resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const spellMetrics = new SpellMetrics({ parent: document.body.getElementsByClassName('spell-metrics')[0], resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const buffAuraMetrics = new AuraMetrics({
    parent: document.body.getElementsByClassName('buff-aura-metrics')[0],
    resultsEmitter: resultsEmitter,
    colorSettings: colorSettings,
}, false);
const debuffAuraMetrics = new AuraMetrics({
    parent: document.body.getElementsByClassName('debuff-aura-metrics')[0],
    resultsEmitter: resultsEmitter,
    colorSettings: colorSettings,
}, true);
const dpsHistogram = new DpsHistogram({ parent: document.body.getElementsByClassName('dps-histogram')[0], resultsEmitter: resultsEmitter, colorSettings: colorSettings });
let currentSimResult = null;
function updateResults() {
    if (currentSimResult == null) {
        resultsEmitter.emit(null);
    }
    else {
        resultsEmitter.emit({
            result: currentSimResult,
            filter: resultsFilter.getFilter(),
        });
    }
}
window.addEventListener('message', async (event) => {
    // Null indicates pending results
    const data = event.data;
    if (data) {
        currentSimResult = await SimResult.fromJson(data);
    }
    else {
        currentSimResult = null;
    }
    updateResults();
});
resultsFilter.changeEmitter.on(() => updateResults());
