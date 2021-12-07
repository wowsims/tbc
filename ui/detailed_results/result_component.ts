import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { RaidSimData } from '/tbc/core/components/detailed_results.js';
import { Component } from '/tbc/core/components/component.js';
import { TypedEvent } from '/tbc/core/typed_event.js';

import { ColorSettings } from './color_settings.js';

export type ResultComponentConfig = {
	parent: HTMLElement,
	rootCssClass?: string,
	resultsEmitter: TypedEvent<RaidSimData | null>;
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

	abstract onSimResult(request: RaidSimRequest, result: RaidSimResult): void;
}
