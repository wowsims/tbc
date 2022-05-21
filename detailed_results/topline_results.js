import { Spec } from '/tbc/core/proto/common.js';
import { ResultComponent } from './result_component.js';
import { RaidSimResultsManager } from '/tbc/core/components/raid_sim_action.js';
export class ToplineResults extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'topline-results-root';
        super(config);
    }
    onSimResult(resultData) {
        const players = resultData.result.getPlayers(resultData.filter);
        let content = RaidSimResultsManager.makeToplineResultsContent(resultData.result, players.length == 1);
        const noManaSpecs = [
            Spec.SpecFeralDruid,
            Spec.SpecFeralTankDruid,
            Spec.SpecRogue,
            Spec.SpecWarrior,
            Spec.SpecProtectionWarrior,
        ];
        if (players.length == 1 && !noManaSpecs.includes(players[0].spec)) {
            const player = players[0];
            const secondsOOM = player.secondsOomAvg;
            const percentOOM = secondsOOM / resultData.result.encounterMetrics.durationSeconds;
            const dangerLevel = percentOOM < 0.01 ? 'safe' : (percentOOM < 0.05 ? 'warning' : 'danger');
            content += `
				<div class="percent-oom ${dangerLevel}">
					<span class="topline-result-avg">${secondsOOM.toFixed(1)}s</span>
					<span class="topline-result-label"> spent OOM</span>
				</div>
			`;
        }
        this.rootElem.innerHTML = content;
    }
}
