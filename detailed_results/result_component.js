import { Component } from '/tbc/core/components/component.js';
;
;
export class ResultComponent extends Component {
    constructor(config) {
        super(config.parent, config.rootCssClass || '');
        this.colorSettings = config.colorSettings;
        this.lastSimResult = null;
        config.resultsEmitter.on((eventID, resultData) => {
            if (!resultData)
                return;
            this.lastSimResult = resultData;
            this.onSimResult(resultData);
        });
    }
    getLastSimResult() {
        if (this.lastSimResult) {
            return this.lastSimResult;
        }
        else {
            throw new Error('No last sim result!');
        }
    }
}
