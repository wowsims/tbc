import { Class } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Sim } from '/tbc/core/sim.js';
import { Player } from '/tbc/core/player.js';
import { classNames, nameToClass, nameToRace, statNames } from '/tbc/core/proto_utils/names.js';
import { talentSpellIdsToTalentString } from '/tbc/core/talents/factory.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { getEnumValues } from '/tbc/core/utils.js';

import { Popup } from './popup.js';

declare var $: any;
declare var tippy: any;

export class SettingsMenu<SpecType extends Spec> extends Popup {
	private readonly simUI: IndividualSimUI<SpecType>;

  constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>) {
    super(parent);
		this.rootElem.classList.add('settings-menu');
		this.simUI = simUI;

		this.rootElem.innerHTML = `
			<div class="settings-menu-title">
				<span>SETTINGS</span>
			</div>
			<div class="settings-menu-content">
				<div class="settings-menu-content-left">
					<button class="restore-defaults-button sim-button">RESTORE DEFAULTS</button>
					<div class="settings-menu-section">
						<div class="fixed-rng-seed">
						</div>
						<div class="last-used-rng-seed-container">
							<span>Last used RNG seed:</span><span class="last-used-rng-seed">0</span>
						</div>
					</div>
				</div>
				<div class="settings-menu-content-right">
					<div class="settings-menu-section settings-menu-ep-weights">
					</div>
				</div>
			</div>
		`;
		this.addCloseButton();

    const restoreDefaultsButton = this.rootElem.getElementsByClassName('restore-defaults-button')[0] as HTMLElement;
		restoreDefaultsButton.addEventListener('click', event => {
			this.simUI.applyDefaults(TypedEvent.nextEventID());
		});
		tippy(restoreDefaultsButton, {
			'content': 'Restores all default settings (gear, consumes, buffs, talents, EP weights, etc).',
			'allowHTML': true,
		});

    const fixedRngSeed = this.rootElem.getElementsByClassName('fixed-rng-seed')[0] as HTMLElement;
    new NumberPicker(fixedRngSeed, this.simUI.sim, {
      label: 'Fixed RNG Seed',
			labelTooltip: 'Seed value for the random number generator used during sims, or 0 to use different randomness each run. Use this to share exact sim results or for debugging.',
      changedEvent: (sim: Sim) => sim.fixedRngSeedChangeEmitter,
      getValue: (sim: Sim) => sim.getFixedRngSeed(),
      setValue: (eventID: EventID, sim: Sim, newValue: number) => {
				sim.setFixedRngSeed(eventID, newValue);
      },
    });

    const lastUsedRngSeed = this.rootElem.getElementsByClassName('last-used-rng-seed')[0] as HTMLElement;
		lastUsedRngSeed.textContent = String(this.simUI.sim.getLastUsedRngSeed());
		this.simUI.sim.lastUsedRngSeedChangeEmitter.on(() => lastUsedRngSeed.textContent = String(this.simUI.sim.getLastUsedRngSeed()));

		this.setupEpWeightsSettings();
  }

	private setupEpWeightsSettings() {
    const sectionRoot = this.rootElem.getElementsByClassName('settings-menu-ep-weights')[0] as HTMLElement;

    const label = document.createElement('span');
    label.classList.add('ep-weights-label');
    label.textContent = 'EP Weights';
		tippy(label, {
			'content': 'EP Weights for sorting the item selector.',
			'allowHTML': true,
		});
    sectionRoot.appendChild(label);

		//const epStats = this.simUI.individualConfig.epStats;
		const epStats = (getEnumValues(Stat) as Array<Stat>).filter(stat => ![Stat.StatMana, Stat.StatEnergy, Stat.StatRage].includes(stat));
    const weightPickers = epStats.map(stat => new NumberPicker(sectionRoot, this.simUI.player, {
			float: true,
      label: statNames[stat],
      changedEvent: (player: Player<any>) => player.epWeightsChangeEmitter,
      getValue: (player: Player<any>) => player.getEpWeights().getStat(stat),
      setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
        const epWeights = player.getEpWeights().withStat(stat, newValue);
        player.setEpWeights(eventID, epWeights);
      },
    }));
	}
}
