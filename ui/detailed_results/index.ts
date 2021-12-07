import { RaidSimData } from '/tbc/core/components/detailed_results.js';
import { TypedEvent } from '/tbc/core/typed_event.js';

import { CastMetrics } from './cast_metrics.js';
import { SpellMetrics } from './cast_metrics.js';
import { BuffAuraMetrics } from './aura_metrics.js'
import { DebuffAuraMetrics } from './aura_metrics.js'
import { DpsHistogram } from './dps_histogram.js';
import { DpsResult } from './dps_result.js';
import { PercentOom } from './percent_oom.js';
import { SourceChart } from './source_chart.js';

declare var Chart: any;

const urlParams = new URLSearchParams(window.location.search);
if (urlParams.has('mainBgColor')) {
	document.body.style.setProperty('--main-bg-color', urlParams.get('mainBgColor')!);
}
if (urlParams.has('mainTextColor')) {
	document.body.style.setProperty('--main-text-color', urlParams.get('mainTextColor')!);
}

const colorSettings = {
	mainTextColor: document.body.style.getPropertyValue('--main-text-color'),
};

Chart.defaults.color = colorSettings.mainTextColor;

const layoutHTML = `
<div class="dr-root">
	<div class="dr-row topline-results">
	</div>
	<div class="dr-row">
		<div class="table-container">
			<div class="title-row">
				<span class="table-title">Spells</span>
			</div>
			<div class="spell-metrics scroll-table">
			</div>
		</div>
	</div>
	<div class="dr-row source-stats">
		<div class="source-chart">
		</div>
	</div>
	<div class="dr-row other-metrics">
		<div class="dr-col-3 start">
			<div class="table-container">
				<div class="title-row">
					<span class="table-title">Casts</span>
				</div>
				<div class="cast-metrics scroll-table">
				</div>
			</div>
		</div>
		<div class="dr-col-3">
			<div class="table-container">
				<div class="title-row">
					<span class="table-title">Buffs</span>
				</div>
				<div class="buff-aura-metrics scroll-table">
				</div>
			</div>
		</div>
		<div class="dr-col-3 end">
			<div class="table-container">
				<div class="title-row">
					<span class="table-title">Target Debuffs</span>
				</div>
				<div class="debuff-aura-metrics scroll-table">
				</div>
			</div>
		</div>
	</div>
	<div class="dr-row dps-histogram">
	</div>
</div>
`;

const resultsEmitter = new TypedEvent<RaidSimData | null>();
window.addEventListener('message', event => {
	// Null indicates pending results
	const data: RaidSimData | null = event.data;

	resultsEmitter.emit(event.data);
});

document.body.innerHTML = layoutHTML;

const toplineResultsDiv = document.body.getElementsByClassName('topline-results')[0] as HTMLElement;
const dpsResult = new DpsResult({ parent: toplineResultsDiv, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const percentOom = new PercentOom({ parent: toplineResultsDiv, resultsEmitter: resultsEmitter, colorSettings: colorSettings });

const castMetrics = new CastMetrics({ parent: document.body.getElementsByClassName('cast-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const spellMetrics = new SpellMetrics({ parent: document.body.getElementsByClassName('spell-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const buffAuraMetrics = new BuffAuraMetrics({ parent: document.body.getElementsByClassName('buff-aura-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const debuffAuraMetrics = new DebuffAuraMetrics({ parent: document.body.getElementsByClassName('debuff-aura-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const sourceChart = new SourceChart({ parent: document.body.getElementsByClassName('source-chart')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const dpsHistogram = new DpsHistogram({ parent: document.body.getElementsByClassName('dps-histogram')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
