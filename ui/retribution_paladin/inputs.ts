import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Target } from '/tbc/core/target.js';

import { RetributionPaladin, RetributionPaladin_Rotation as RetributionPaladinRotation, PaladinTalents as PaladinTalents, RetributionPaladin_Options as RetributionPaladinOptions } from '/tbc/core/proto/paladin.js';

import * as Presets from './presets.js';
import { SimUI } from '../core/sim_ui.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const RetributionPaladinRotationConfig = {
	inputs: [
		{
			type: 'boolean' as const, cssClass: 'consecration-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Use Consecration',
				labelTooltip: 'Use consecration whenever the target does not already have the DoT.',
				changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecRetributionPaladin>) => player.getRotation().consecration,
				setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.consecration = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		}
	],
}