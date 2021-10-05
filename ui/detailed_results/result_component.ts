import { IndividualSimRequest, IndividualSimResult } from '../core/proto/api.js';
import { IndividualSimData } from '../core/components/detailed_results.js';
import { Component } from '../core/components/component.js';
import { TypedEvent } from '../core/typed_event.js';

import { ColorSettings } from './color_settings.js';

export type ResultComponentConfig = {
	parent: HTMLElement,
	rootCssClass?: string,
	resultsEmitter: TypedEvent<IndividualSimData | null>;
	colorSettings: ColorSettings;
};

export abstract class ResultComponent extends Component {
	private readonly colorSettings: ColorSettings;

  constructor(config: ResultComponentConfig) {
    super(config.parent, config.rootCssClass || '');
		this.colorSettings = config.colorSettings;

		config.resultsEmitter.on(simData => {
			if (!simData)
				return;

			this.onSimResult(simData.request, simData.result);
		});
	}

	abstract onSimResult(request: IndividualSimRequest, result: IndividualSimResult): void;
}
