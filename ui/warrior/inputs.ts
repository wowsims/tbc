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

import { Warrior, Warrior_Rotation as WarriorRotation, WarriorTalents as WarriorTalents, Warrior_Options as WarriorOptions } from '/tbc/core/proto/warrior.js';
import { Warrior_Rotation_Type as RotationType, Warrior_Rotation_ArmsSlamRotation as ArmsSlamRotation, Warrior_Rotation_ArmsDWRotation as ArmsDWRotation, Warrior_Rotation_FuryRotation as FuryRotation } from '/tbc/core/proto/warrior.js';
import { Warrior_Rotation_FuryRotation_PrimaryInstant as PrimaryInstant } from '/tbc/core/proto/warrior.js';

import * as Presets from './presets.js';
import { SimUI } from '../core/sim_ui.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const Recklessness = {
	id: ActionId.fromSpellId(1719),
	states: 2,
	extraCssClasses: [
		'warrior-recklessness-picker',
	],
	changedEvent: (player: Player<Spec.SpecWarrior>) => player.specOptionsChangeEmitter,
	getValue: (player: Player<Spec.SpecWarrior>) => player.getSpecOptions().recklessness,
	setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
		const newOptions = player.getSpecOptions();
		newOptions.recklessness = newValue
		player.setSpecOptions(eventID, newOptions);
	},
};

export const WarriorRotationConfig = {
	inputs: [
		{
			type: 'enum' as const,
			getModObject: (simUI: IndividualSimUI<any>) => simUI,
			config: {
				extraCssClasses: [
					'rotation-type-enum-picker',
				],
				label: 'Spec',
				labelTooltip: 'Switches between spec rotation settings. Will also update talents to defaults for the selected spec.',
				values: [
					{
						name: 'Arms Slam', value: RotationType.ArmsSlam,
					},
					{
						name: 'Arms DW', value: RotationType.ArmsDW,
					},
					{
						name: 'Fury', value: RotationType.Fury,
					},
				],
				changedEvent: (simUI: IndividualSimUI<Spec.SpecWarrior>) => simUI.player.rotationChangeEmitter,
				getValue: (simUI: IndividualSimUI<Spec.SpecWarrior>) => simUI.player.getRotation().type,
				setValue: (eventID: EventID, simUI: IndividualSimUI<Spec.SpecWarrior>, newValue: number) => {
					const newRotation = simUI.player.getRotation();
					newRotation.type = newValue;

					TypedEvent.freezeAllAndDo(() => {
						if (newRotation.type == RotationType.Fury) {
							simUI.player.setTalentsString(eventID, Presets.FuryTalents.data);
							if (!newRotation.fury) {
								newRotation.fury = FuryRotation.clone(Presets.DefaultFuryRotation.fury!);
							}
						} else if (newRotation.type == RotationType.ArmsSlam) {
							simUI.player.setTalentsString(eventID, Presets.ArmsSlamTalents.data);
							if (!newRotation.armsSlam) {
								newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam!);
							}
						} else {
							simUI.player.setTalentsString(eventID, Presets.ArmsDWTalents.data);
							if (!newRotation.armsDw) {
								newRotation.armsDw = ArmsDWRotation.clone(Presets.DefaultArmsDWRotation.armsDw!);
							}
						}

						simUI.player.setRotation(eventID, newRotation);
					});

					simUI.recomputeSettingsLayout();
				},
			},
		},
		// ********************************************************
		//                       FURY INPUTS
		// ********************************************************
		{
			type: 'enum' as const,
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				extraCssClasses: [
					'rotation-type-enum-picker',
				],
				label: 'Primary Instant',
				labelTooltip:'Main instant ability that will be prioritized above everything else while it is off CD.',
				values: [
					{
						name: 'Bloodthirst', value: PrimaryInstant.Bloodthirst,
					},
					{
						name: 'Whirlwind', value: PrimaryInstant.Whirlwind,
					},
				],
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().fury?.primaryInstant || PrimaryInstant.Whirlwind,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
					const newRotation = player.getRotation();
					if (!newRotation.fury) {
						newRotation.fury = FuryRotation.create();
					}
					newRotation.fury.primaryInstant = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().type == RotationType.Fury,
			},
		},
		{
			type: 'boolean' as const,
			cssClass: 'bt-exec-picker-fury',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'BT during Execute Phase',
				labelTooltip: 'Use Bloodthirst during Execute Phase.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().fury?.useBtDuringExecute || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.fury) {
						newRotation.fury = FuryRotation.clone(Presets.DefaultFuryRotation.fury!);
					}
					newRotation.fury.useBtDuringExecute = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().type == RotationType.Fury,
			},
		},
		
		{
			type: 'number' as const,
			cssClass: 'rampage-duration-threshold',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Rampage refresh timing (seconds)',
				labelTooltip: 'Refresh rampage when it has certain duration left on it.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().fury?.rampageCdThreshold || 0,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
					const newRotation = player.getRotation();
					if (!newRotation.fury) {
						newRotation.fury = FuryRotation.clone(Presets.DefaultFuryRotation.fury!);
					}
					newRotation.fury.rampageCdThreshold = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().type == RotationType.Fury,
			},
		},
		// ********************************************************
		//                      ARMS SLAM INPUTS
		// ********************************************************
		{
			type: 'boolean' as const,
			cssClass: 'ms-exec-picker-arms-slam',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'MS during Execute Phase',
				labelTooltip: 'Use Mortal Strike during Execute Phase.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().armsSlam?.useMsDuringExecute || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.armsSlam) {
						newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam!);
					}
					newRotation.armsSlam.useMsDuringExecute = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().type == RotationType.ArmsSlam,
			},
		},

		{
			type: 'boolean' as const,
			cssClass: 'slam-exec-picker-arms-slam',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Slam during Execute Phase',
				labelTooltip: 'Use Slam during Execute Phase.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().armsSlam?.useSlamDuringExecute || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.armsSlam) {
						newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam!);
					}
					newRotation.armsSlam.useSlamDuringExecute = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().type == RotationType.ArmsSlam,
			},
		},

		{
			type: 'number' as const,
			cssClass: 'slam-latency',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Slam Latency (ms)',
				labelTooltip: 'Add delay to slam casting.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().armsSlam?.slamLatency || 0,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
					const newRotation = player.getRotation();
					if (!newRotation.armsSlam) {
						newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam!);
					}
					newRotation.armsSlam.slamLatency = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().type == RotationType.ArmsSlam,
			},
		},

		// ********************************************************
		//                      ARMS DW INPUTS
		// ********************************************************
		{
			type: 'boolean' as const,
			cssClass: 'ms-exec-picker-arms-dw',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'MS during Execute Phase',
				labelTooltip: 'Use Mortal Strike during Execute Phase.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().armsDw?.useMsDuringExecute || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.armsDw) {
						newRotation.armsDw = ArmsDWRotation.clone(Presets.DefaultArmsDWRotation.armsDw!);
					}
					newRotation.armsDw.useMsDuringExecute = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().type == RotationType.ArmsDW,
			},
		},

		// ********************************************************
		//                      GENERAL INPUTS
		// ********************************************************
	{
		type: 'boolean' as const,
		cssClass: 'ww-exec-picker',
		getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
		config: {
			label: 'WW during Execute Phase',
			labelTooltip: 'Use Whirlwind during Execute Phase.',
			changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().useWwDuringExecute,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
				const newRotation = player.getRotation();
				newRotation.useWwDuringExecute = newValue;
				player.setRotation(eventID, newRotation);
			},
		},
	},
	{
		type: 'boolean' as const,
		cssClass: 'hs-exec-picker',
		getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
		config: {
			label: 'HS during Execute Phase',
			labelTooltip: 'Use Heroic Strike during Execute Phase.',
			changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().useHsDuringExecute,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
				const newRotation = player.getRotation();
				newRotation.useHsDuringExecute = newValue;
				player.setRotation(eventID, newRotation);
			},
		},
	},
	{
		type: 'number' as const,
		cssClass: 'hs-rage-threshold',
		getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
		config: {
			label: 'HS rage threshold',
			labelTooltip: 'Queue HS when rage is above:',
			changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().hsRageThreshold,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
				const newRotation = player.getRotation();
				newRotation.hsRageThreshold = newValue;
				player.setRotation(eventID, newRotation);
			},
		},
	},
	{
		type: 'boolean' as const,
		cssClass: 'overpower-picker',
		getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
		config: {
			label: 'Use Overpower',
			labelTooltip: 'Use Overpower when it is possible.',
			changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().useOverpower,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
				const newRotation = player.getRotation();
				newRotation.useOverpower = newValue;
				player.setRotation(eventID, newRotation);
			},
		},
	},
	{
		type: 'number' as const,
		cssClass: 'overpower-rage-threshold',
		getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
		config: {
			label: 'Overpower rage threshold',
			labelTooltip: 'Use Overpower when rage is below a point.',
			changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().overpowerRageThreshold,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
				const newRotation = player.getRotation();
				newRotation.overpowerRageThreshold = newValue;
				player.setRotation(eventID, newRotation);
			},
			showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().useOverpower,
		},
	},
	{
		type: 'boolean' as const,
		cssClass: 'hamstring-picker',
		getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
		config: {
			label: 'Use Hamstring',
			labelTooltip: 'Use Hamstring on free global.',
			changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().useHamstring,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
				const newRotation = player.getRotation();
				newRotation.useHamstring = newValue;
				player.setRotation(eventID, newRotation);
			},
		},
	},
	{
		type: 'number' as const,
		cssClass: 'hamstring-rage-threshold',
		getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
		config: {
			label: 'Hamstring rage threshold',
			labelTooltip: 'Use Hamstring when rage is above a ',
			changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().hamstringRageThreshold,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
				const newRotation = player.getRotation();
				newRotation.hamstringRageThreshold = newValue;
				player.setRotation(eventID, newRotation);
			},
			showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().useHamstring,
		},
	},
	],
};