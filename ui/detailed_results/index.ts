import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';

import { SimResultData } from './result_component.js';
import { ResultsFilter } from './results_filter.js';
import { CastMetrics } from './cast_metrics.js';
import { MeleeMetrics } from './melee_metrics.js';
import { SpellMetrics } from './spell_metrics.js';
import { PlayerDamageMetrics } from './player_damage.js';
import { AuraMetrics } from './aura_metrics.js'
import { DpsHistogram } from './dps_histogram.js';
import { DpsResult } from './dps_result.js';
import { PercentOom } from './percent_oom.js';
import { SourceChart } from './source_chart.js';
import { Timeline } from './timeline.js';

declare var Chart: any;

const urlParams = new URLSearchParams(window.location.search);
if (urlParams.has('mainTextColor')) {
	document.body.style.setProperty('--main-text-color', urlParams.get('mainTextColor')!);
}
if (urlParams.has('themeColorPrimary')) {
	document.body.style.setProperty('--theme-color-primary', urlParams.get('themeColorPrimary')!);
}
if (urlParams.has('themeColorBackground')) {
	document.body.style.setProperty('--theme-color-background', urlParams.get('themeColorBackground')!);
}
if (urlParams.has('themeColorBackgroundRaw')) {
	document.body.style.setProperty('--theme-color-background-raw', urlParams.get('themeColorBackgroundRaw')!);
}
if (urlParams.has('themeBackgroundImage')) {
	document.body.style.setProperty('--theme-background-image', urlParams.get('themeBackgroundImage')!);
}

const isIndividualSim = urlParams.has('isIndividualSim');
if (isIndividualSim) {
	document.body.classList.add('individual-sim');
}

if (!window.frameElement) {
	// Means we're not inside an iframe, i.e. this is a new tab.
	document.body.classList.add('new-tab');
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
		<li class="dr-tab-tab active"><a data-toggle="tab" href="#damageTab">DAMAGE</a></li>
		<li class="dr-tab-tab"><a data-toggle="tab" href="#buffsTab">BUFFS</a></li>
		<li class="dr-tab-tab"><a data-toggle="tab" href="#debuffsTab">DEBUFFS</a></li>
		<li class="dr-tab-tab"><a data-toggle="tab" href="#castsTab">CASTS</a></li>
		<li class="dr-tab-tab"><a data-toggle="tab" href="#timelineTab" id="timelineTabTab">TIMELINE</a></li>
	</ul>
	<div class="tab-content">
		<div id="damageTab" class="tab-pane dr-tab-content damage-content fade active in">
			<div class="dr-row topline-results">
			</div>
			<div class="dr-row all-players-only">
				<div class="player-damage-metrics">
				</div>
			</div>
			<div class="dr-row single-player-only">
				<div class="melee-metrics">
				</div>
			</div>
			<div class="dr-row single-player-only">
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
		<div id="timelineTab" class="tab-pane dr-tab-content timeline-content fade">
			<div class="dr-row">
				<div class="timeline">
				</div>
			</div>
		</div>
	</div>
</div>
`;

document.body.innerHTML = layoutHTML;
const resultsEmitter = new TypedEvent<SimResultData | null>();

const resultsFilter = new ResultsFilter({
	parent: document.body.getElementsByClassName('results-filter')[0] as HTMLElement,
	resultsEmitter: resultsEmitter,
	colorSettings: colorSettings,
});

const toplineResultsDiv = document.body.getElementsByClassName('topline-results')[0] as HTMLElement;
const dpsResult = new DpsResult({ parent: toplineResultsDiv, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const percentOom = new PercentOom({ parent: toplineResultsDiv, resultsEmitter: resultsEmitter, colorSettings: colorSettings });

const castMetrics = new CastMetrics({ parent: document.body.getElementsByClassName('cast-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const meleeMetrics = new MeleeMetrics({ parent: document.body.getElementsByClassName('melee-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const spellMetrics = new SpellMetrics({ parent: document.body.getElementsByClassName('spell-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const playerDamageMetrics = new PlayerDamageMetrics({ parent: document.body.getElementsByClassName('player-damage-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings }, resultsFilter);
const buffAuraMetrics = new AuraMetrics({
	parent: document.body.getElementsByClassName('buff-aura-metrics')[0] as HTMLElement,
	resultsEmitter: resultsEmitter,
	colorSettings: colorSettings,
}, false);
const debuffAuraMetrics = new AuraMetrics({
	parent: document.body.getElementsByClassName('debuff-aura-metrics')[0] as HTMLElement,
	resultsEmitter: resultsEmitter,
	colorSettings: colorSettings,
}, true);
const dpsHistogram = new DpsHistogram({ parent: document.body.getElementsByClassName('dps-histogram')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });

const timeline = new Timeline({
	parent: document.body.getElementsByClassName('timeline')[0] as HTMLElement,
	resultsEmitter: resultsEmitter,
	colorSettings: colorSettings,
});
(document.getElementById('timelineTabTab') as HTMLElement).addEventListener('click', event => timeline.render());

let currentSimResult: SimResult | null = null;
function updateResults() {
	const eventID = TypedEvent.nextEventID();
	if (currentSimResult == null) {
		resultsEmitter.emit(eventID, null);
	} else {
		resultsEmitter.emit(eventID, {
			eventID: eventID,
			result: currentSimResult,
			filter: resultsFilter.getFilter(),
		});
	}
}

window.addEventListener('message', async event => {
	// Null indicates pending results
	const data: Object | null = event.data;

	if (data) {
		currentSimResult = await SimResult.fromJson(data);
	} else {
		currentSimResult = null;
	}
	updateResults();
});

resultsFilter.changeEmitter.on(() => updateResults());

const rootDiv = document.body.getElementsByClassName('dr-root')[0] as HTMLElement;
resultsEmitter.on((eventID, resultData) => {
	if (resultData?.filter.player || resultData?.filter.player === 0) {
		rootDiv.classList.remove('all-players');
		rootDiv.classList.add('single-player');
	} else {
		rootDiv.classList.add('all-players');
		rootDiv.classList.remove('single-player');
	}
});
