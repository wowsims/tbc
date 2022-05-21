import { Component } from '/tbc/core/components/component.js';
import { RaidTargetPicker } from '/tbc/core/components/raid_target_picker.js';
import { emptyRaidTarget } from '/tbc/core/proto_utils/utils.js';
const MAX_TANKS = 4;
export class TanksPicker extends Component {
    constructor(parentElem, raidSimUI) {
        super(parentElem, 'tanks-picker-root');
        this.raidSimUI = raidSimUI;
        this.rootElem.innerHTML = `
			<fieldset class="tanks-picker-container settings-section">
				<legend>TANKS</legend>
			</fieldset>
		`;
        const tanksContainer = this.rootElem.getElementsByClassName('tanks-picker-container')[0];
        const raid = this.raidSimUI.sim.raid;
        for (let i = 0; i < MAX_TANKS; i++) {
            const row = document.createElement('div');
            row.classList.add('tank-picker-row');
            tanksContainer.appendChild(row);
            const sourceElem = document.createElement('span');
            sourceElem.textContent = i == 0 ? 'MAIN TANK' : `TANK ${i + 1}`;
            sourceElem.classList.add('tank-picker-label');
            row.appendChild(sourceElem);
            const arrow = document.createElement('span');
            arrow.classList.add('fa', 'fa-arrow-right');
            row.appendChild(arrow);
            const tankIndex = i;
            const raidTargetPicker = new RaidTargetPicker(row, raid, raid, {
                extraCssClasses: [
                    'tank-picker',
                ],
                noTargetLabel: 'Unassigned',
                compChangeEmitter: raid.compChangeEmitter,
                changedEvent: (raid) => raid.tanksChangeEmitter,
                getValue: (raid) => raid.getTanks()[tankIndex] || emptyRaidTarget(),
                setValue: (eventID, raid, newValue) => {
                    const tanks = raid.getTanks();
                    for (let i = 0; i < tankIndex; i++) {
                        if (!tanks[i]) {
                            tanks.push(emptyRaidTarget());
                        }
                    }
                    tanks[tankIndex] = newValue;
                    raid.setTanks(eventID, tanks);
                },
            });
        }
    }
}
