import { Class } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Sim } from '/tbc/core/sim.js';
import { Player } from '/tbc/core/player.js';
import { classNames, nameToClass, nameToRace } from '/tbc/core/proto_utils/names.js';
import { talentSpellIdsToTalentString } from '/tbc/core/talents/factory.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';

import { CloseButton } from './close_button.js';
import { Component } from './component.js';

declare var $: any;
declare var tippy: any;

export class SettingsMenu<SpecType extends Spec> extends Component {
	private readonly simUI: IndividualSimUI<SpecType>;

  constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>) {
    super(parent, 'settings-menu');
		this.simUI = simUI;

		this.rootElem.id = 'settingsMenu';
		this.rootElem.innerHTML = `
			<div class="settings-menu-title">
				<span>SETTINGS</span>
			</div>
			<div class="settings-menu-content">
				<button class="restore-defaults-button sim-button">RESTORE DEFAULTS</button>
				<div class="settings-menu-section">
					<div class="fixed-rng-seed">
					</div>
					<div>
						<span>Last used RNG seed:</span><span class="last-used-rng-seed">0</span>
					</div>
				</div>
			</div>
		`;

		const computedStyles = window.getComputedStyle(parent);
		this.rootElem.style.setProperty('--main-text-color', computedStyles.getPropertyValue('--main-text-color').trim());
		this.rootElem.style.setProperty('--theme-color-primary', computedStyles.getPropertyValue('--theme-color-primary').trim());
		this.rootElem.style.setProperty('--theme-color-background', computedStyles.getPropertyValue('--theme-color-background').trim());
		this.rootElem.style.setProperty('--theme-color-background-raw', computedStyles.getPropertyValue('--theme-color-background-raw').trim());

		new CloseButton(this.rootElem, () => {
			$('#settingsMenu').bPopup().close();
			this.rootElem.remove();
		});

		$('#settingsMenu').bPopup({
			onClose: () => {
				this.rootElem.remove();
			},
		});

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
  }
}
