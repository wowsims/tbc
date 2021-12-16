import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';
import { Component } from '/tbc/core/components/component.js';
import { TypedEvent } from '/tbc/core/typed_event.js';

import { ColorSettings } from './color_settings.js';

export interface SimResultData {
	result: SimResult,
	filter: SimResultFilter,
};

export interface ResultComponentConfig {
	parent: HTMLElement,
	rootCssClass?: string,
	resultsEmitter: TypedEvent<SimResultData | null>,
	colorSettings: ColorSettings,
};

export abstract class ResultComponent extends Component {
	private readonly colorSettings: ColorSettings;

  constructor(config: ResultComponentConfig) {
    super(config.parent, config.rootCssClass || '');
		this.colorSettings = config.colorSettings;

		config.resultsEmitter.on(resultData => {
			if (!resultData)
				return;

			this.onSimResult(resultData);
		});
	}

	abstract onSimResult(resultData: SimResultData): void;
}
