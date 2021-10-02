import { IndividualSimData } from '../core/components/detailed_results.js';
import { TypedEvent } from '../core/typed_event.js';

import { CastMetrics } from './cast_metrics.js';
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
const sourceChart = new SourceChart({ parent: document.body.getElementsByClassName('source-chart')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
const dpsHistogram = new DpsHistogram({ parent: document.body.getElementsByClassName('dps-histogram')[0] as HTMLElement, resultsEmitter: resultsEmitter, colorSettings: colorSettings });
