import { Component } from '/tbc/core/components/component.js';
;
;
export class ResultComponent extends Component {
    constructor(config) {
        super(config.parent, config.rootCssClass || '');
        this.colorSettings = config.colorSettings;
        config.resultsEmitter.on(resultData => {
            if (!resultData)
                return;
            this.onSimResult(resultData);
        });
    }
}
