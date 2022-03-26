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

import {
	WarriorTalents as WarriorTalents,
	ProtectionWarrior,
	ProtectionWarrior_Rotation as ProtectionWarriorRotation,
	ProtectionWarrior_Rotation_DemoShout as DemoShout,
	ProtectionWarrior_Rotation_ThunderClap as ThunderClap,
	ProtectionWarrior_Options as ProtectionWarriorOptions
} from '/tbc/core/proto/warrior.js';

import * as Presets from './presets.js';
import { SimUI } from '../core/sim_ui.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const ProtectionWarriorRotationConfig = {
	inputs: [
		{
			type: 'enum' as const, cssClass: 'demo-shout-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Demo Shout',
				values: [
					{ name: 'None', value: DemoShout.DemoShoutNone },
					{ name: 'Maintain Debuff', value: DemoShout.DemoShoutMaintain },
					{ name: 'Filler', value: DemoShout.DemoShoutFiller },
				],
				changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecProtectionWarrior>) => player.getRotation().demoShout,
				setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => {
					const newRotation = player.getRotation();
					newRotation.demoShout = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		},
		{
			type: 'enum' as const, cssClass: 'thunder-clap-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Thunder Clap',
				values: [
					{ name: 'None', value: ThunderClap.ThunderClapNone },
					{ name: 'Maintain Debuff', value: ThunderClap.ThunderClapMaintain },
					{ name: 'On CD', value: ThunderClap.ThunderClapOnCD },
				],
				changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecProtectionWarrior>) => player.getRotation().thunderClap,
				setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => {
					const newRotation = player.getRotation();
					newRotation.thunderClap = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		},
	],
};
