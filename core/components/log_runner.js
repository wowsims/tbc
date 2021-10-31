import { Component } from './component.js';
export class LogRunner extends Component {
    constructor(parent, simUI, results, detailedResults) {
        super(parent, 'log-runner-root');
        const controlBar = document.createElement('div');
        controlBar.classList.add('log-runner-control-bar');
        this.rootElem.appendChild(controlBar);
        const simButton = document.createElement('button');
        simButton.classList.add('log-runner-button');
        simButton.textContent = 'Sim once with logs';
        controlBar.appendChild(simButton);
        const logsDiv = document.createElement('div');
        logsDiv.classList.add('log-runner-logs');
        this.rootElem.appendChild(logsDiv);
        simButton.addEventListener('click', async () => {
            const simRequest = simUI.makeCurrentIndividualSimRequest(1, true);
            results.setPending();
            detailedResults.setPending();
            const result = await simUI.sim.individualSim(simRequest);
            results.setSimResult(simRequest, result);
            detailedResults.setSimResult(simRequest, result);
            const lines = result.logs.split('\n');
            logsDiv.textContent = '';
            lines.forEach(line => {
                const lineElem = document.createElement('span');
                lineElem.textContent = line;
                logsDiv.appendChild(lineElem);
            });
        });
    }
}
