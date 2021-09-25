import { IndividualSimResult } from '../core/api/api';
import { TypedEvent } from '../core/typed_event';

import { DpsHistogram } from './dps_histogram';

const layoutHTML = `
<div class="dr-root">
	<div class="dps-histogram">
	</div>
</div>
`;

const resultsEmitter = new TypedEvent<IndividualSimResult | null>();
window.addEventListener('message', event => {
	// Null indicates pending results
	const data: IndividualSimResult | null = event.data;

	resultsEmitter.emit(event.data);
});

document.body.innerHTML = layoutHTML;
const dpsHistogram = new DpsHistogram(document.body.getElementsByClassName('dps-histogram')[0] as HTMLElement, resultsEmitter);
