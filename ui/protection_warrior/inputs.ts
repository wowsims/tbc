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
	WarriorShout,
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

export const ShoutPicker = {
	type: 'enum' as const, cssClass: 'shout-picker',
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		label: 'Shout',
		labelTooltip: 'Shout buff to maintain.',
		values: [
			{ name: 'None', value: WarriorShout.WarriorShoutNone },
			{ name: 'Battle Shout', value: WarriorShout.WarriorShoutBattle },
			{ name: 'Commanding Shout', value: WarriorShout.WarriorShoutCommanding },
		],
		changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => player.rotationChangeEmitter,
		getValue: (player: Player<Spec.SpecProtectionWarrior>) => player.getSpecOptions().shout,
		setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			newOptions.shout = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
};

export const PrecastShout = {
	type: 'boolean' as const,
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		extraCssClasses: [
			'precast-shout-picker',
		],
		label: 'Precast Shout',
		labelTooltip: 'Selected shout is cast 10 seconds before combat starts.',
		changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecProtectionWarrior>) => player.getSpecOptions().precastShout,
		setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: boolean) => {
			const newOptions = player.getSpecOptions();
			newOptions.precastShout = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
};

export const PrecastShoutWithSapphire = {
	type: 'boolean' as const,
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		extraCssClasses: [
			'precast-shout-with-sapphire-picker',
		],
		label: 'Precast with Sapphire',
		labelTooltip: 'Snapshot bonus from Solarian\'s Sapphire (+70 attack power) with precast shout.',
		changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent.onAny([player.specOptionsChangeEmitter, player.gearChangeEmitter]),
		getValue: (player: Player<Spec.SpecProtectionWarrior>) => player.getSpecOptions().precastShoutSapphire,
		setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: boolean) => {
			const newOptions = player.getSpecOptions();
			newOptions.precastShoutSapphire = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
		enableWhen: (player: Player<Spec.SpecProtectionWarrior>) => player.getSpecOptions().shout == WarriorShout.WarriorShoutBattle && player.getSpecOptions().precastShout && !player.getGear().hasTrinket(30446),
	},
};

export const PrecastShoutWithT2 = {
	type: 'boolean' as const,
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		extraCssClasses: [
			'precast-shout-with-t2-picker',
		],
		label: 'Precast with T2',
		labelTooltip: 'Snapshot T2 set bonus (+30 attack power) with precast shout.',
		changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecProtectionWarrior>) => player.getSpecOptions().precastShoutT2,
		setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: boolean) => {
			const newOptions = player.getSpecOptions();
			newOptions.precastShoutT2 = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
		enableWhen: (player: Player<Spec.SpecProtectionWarrior>) => player.getSpecOptions().shout == WarriorShout.WarriorShoutBattle && player.getSpecOptions().precastShout,
	},
};

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
		{
			type: 'number' as const, cssClass: 'heroic-strike-threshold-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'HS Threshold',
				labelTooltip: 'Minimum Rage requires to queue Heroic Strike or Cleave.',
				changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecProtectionWarrior>) => player.getRotation().hsRageThreshold,
				setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => {
					const newRotation = player.getRotation();
					newRotation.hsRageThreshold = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		},
	],
};

function makeBooleanBuffInput(id: ActionId, optionsFieldName: keyof ProtectionWarriorOptions): IconPickerConfig<Player<any>, boolean> {
	return {
		id: id,
		states: 2,
		changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecProtectionWarrior>) => player.getSpecOptions()[optionsFieldName] as boolean,
		setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: boolean) => {
			const newOptions = player.getSpecOptions();
			(newOptions[optionsFieldName] as boolean) = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	};
}
