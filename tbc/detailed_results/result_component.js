import { Component } from '../core/components/component.js';
export class ResultComponent extends Component {
    constructor(config) {
        super(config.parent, config.rootCssClass || '');
        this.colorSettings = config.colorSettings;
        config.resultsEmitter.on(simData => {
            if (!simData)
                return;
            this.onSimResult(simData.request, simData.result);
        });
    }
}
