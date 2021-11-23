import { IndividualSimData } from '/tbc/core/components/detailed_results.js';
import { TypedEvent } from '/tbc/core/typed_event.js';

import { CastMetrics } from './cast_metrics.js';
import { OtherCastMetrics } from './cast_metrics.js';
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
	<div class="dr-row cast-metrics">
	</div>
	<div class="dr-row source-stats">
		<div class="source-chart">
		</div>
	</div>
	<div class="dr-row other-metrics">
		<div class="dr-col other-cast-metrics">
		</div>
		<div class="dr-col buff-aura-metrics">
		</div>
		<div class="dr-col debuff-aura-metrics">
		</div>
	</div>
	<div class="dr-row dps-histogram">
	</div>
</div>
`;

const resultsEmitter = new TypedEvent<IndividualSimData | null>();
window.addEventListener('message', event => {
	// Null indicates pending results
	const data: IndividualSimData | null = event.data;

	resultsEmitter.emit(event.data);
});

document.body.innerHTML = layoutHTML;

const toplineResultsDiv = document.body.getElementsByClassName('topline-results')[0] as HTMLElement;
const dpsResult = new DpsResult({ parent: toplineResultsDiv, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const percentOom = new PercentOom({ parent: toplineResultsDiv, resultsEmitter: resultsEmitter, colorSettings: colorSettings });

const castMetrics = new CastMetrics({ parent: document.body.getElementsByClassName('cast-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const otherCastMetrics = new OtherCastMetrics({ parent: document.body.getElementsByClassName('other-cast-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const buffAuraMetrics = new BuffAuraMetrics({ parent: document.body.getElementsByClassName('buff-aura-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const debuffAuraMetrics = new DebuffAuraMetrics({ parent: document.body.getElementsByClassName('debuff-aura-metrics')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const sourceChart = new SourceChart({ parent: document.body.getElementsByClassName('source-chart')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const dpsHistogram = new DpsHistogram({ parent: document.body.getElementsByClassName('dps-histogram')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });

// Need to implement this function to run after all the above
// function balanceTables() {
// 	var otherMetricsTables = document.querySelectorAll('.other-metrics table');
// 	var max = 0;
// 	for (var i = 0; i < otherMetricsTables.length; i++) {
// 		var rows = otherMetricsTables[i].querySelectorAll('tbody tr').length;
// 		if (rows > max) max = rows;
// 	}
// 	for (var i = 0; i < otherMetricsTables.length; i++) {
// 		var rows = otherMetricsTables[i].querySelectorAll('tbody tr').length;
// 		for (var j = rows; j < max; j++) {
// 			var body = otherMetricsTables[i].querySelector('tbody');
// 			var emptyRow = document.createElement('tr');
// 			emptyRow.classList.add('empty');
// 			body.appendChild(emptyRow);
// 		}
// 	}
// 	console.log("executed");
// }
// balanceTables();
