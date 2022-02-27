import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { IconEnumPicker, IconEnumPickerConfig } from '/tbc/core/components/icon_enum_picker.js';
import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Target } from '/tbc/core/target.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

import {
	Rogue,
	Rogue_Rotation as RogueRotation,
	Rogue_Options as RogueOptions,
} from '/tbc/core/proto/rogue.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const ThistleTea = makeBooleanRogueBuffInput(ActionId.fromItemId(7676), 'useThistleTea');

export const RogueRotationConfig = {
	inputs: [
		{
			type: 'boolean' as const, cssClass: 'maintain-expose-armor-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Maintain EA',
				labelTooltip: 'Keeps Expose Armor active on the primary target.',
				changedEvent: (player: Player<Spec.SpecRogue>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecRogue>) => player.getRotation().maintainExposeArmor,
				setValue: (eventID: EventID, player: Player<Spec.SpecRogue>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.maintainExposeArmor = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		},
		{
			type: 'boolean' as const, cssClass: 'use-rupture-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Use Rupture',
				labelTooltip: 'Uses Rupture over Eviscerate when appropriate.',
				changedEvent: (player: Player<Spec.SpecRogue>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecRogue>) => player.getRotation().useRupture,
				setValue: (eventID: EventID, player: Player<Spec.SpecRogue>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.useRupture = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		},
	],
};

function makeBooleanRogueBuffInput(id: ActionId, optionsFieldName: keyof RogueOptions): IconPickerConfig<Player<any>, boolean> {
	return {
	  id: id,
	  states: 2,
		changedEvent: (player: Player<Spec.SpecRogue>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecRogue>) => player.getSpecOptions()[optionsFieldName] as boolean,
		setValue: (eventID: EventID, player: Player<Spec.SpecRogue>, newValue: boolean) => {
			const newOptions = player.getSpecOptions();
			(newOptions[optionsFieldName] as boolean) = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	};
}
