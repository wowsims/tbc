import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ItemOrSpellId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Target } from '/tbc/core/target.js';

import { Warrior, Warrior_Rotation as WarriorRotation, WarriorTalents as WarriorTalents, Warrior_Options as WarriorOptions } from '/tbc/core/proto/warrior.js';
import { Warrior_Rotation_Type as RotationType, Warrior_Rotation_ArmsSlamRotation as ArmsSlamRotation, Warrior_Rotation_ArmsDWRotation as ArmsDWRotation, Warrior_Rotation_FuryRotation as FuryRotation, Warrior_Rotation_GeneralRotation as GeneralRotation } from '/tbc/core/proto/warrior.js';
import { Warrior_Rotation_FuryRotation_PrimaryInstant as PrimaryInstant } from '/tbc/core/proto/warrior.js';

import * as Presets from './presets.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const Recklessness = {
	id: { spellId: 1719},
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
				label: 'BT on Execute Phase',
				labelTooltip: 'Use Bloodthirst on Execute Phase.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().fury?.useBtExec || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.fury) {
						newRotation.fury = FuryRotation.clone(Presets.DefaultFuryRotation.fury!);
					}
					newRotation.fury.useBtExec = newValue;
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
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().fury?.rampageCdTresh || 0,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
					const newRotation = player.getRotation();
					if (!newRotation.fury) {
						newRotation.fury = FuryRotation.clone(Presets.DefaultFuryRotation.fury!);
					}
					newRotation.fury.rampageCdTresh = newValue;
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
				label: 'MS on Execute Phase',
				labelTooltip: 'Use Mortal Strike on Execute Phase.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().armsSlam?.useMsExec || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.armsSlam) {
						newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam!);
					}
					newRotation.armsSlam.useMsExec = newValue;
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
				label: 'Slam on Execute Phase',
				labelTooltip: 'Use Slam on Execute Phase.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().armsSlam?.useSlamExec || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.armsSlam) {
						newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam!);
					}
					newRotation.armsSlam.useSlamExec = newValue;
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
				label: 'MS on Execute Phase',
				labelTooltip: 'Use Mortal Strike on Execute Phase.',
				changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().armsDw?.useMsExec || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.armsDw) {
						newRotation.armsDw = ArmsDWRotation.clone(Presets.DefaultArmsDWRotation.armsDw!);
					}
					newRotation.armsDw.useMsExec = newValue;
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
			label: 'WW on Execute Phase',
			labelTooltip: 'Use Whirlwind on Execute Phase.',
			changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().general?.useWwExec || false,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
				const newRotation = player.getRotation();
				if (!newRotation.general) {
					newRotation.general = GeneralRotation.clone(Presets.DefaultGeneralRotation.general!);
				}
				newRotation.general.useWwExec = newValue;
				player.setRotation(eventID, newRotation);
			},
		},
	},
	{
		type: 'boolean' as const,
		cssClass: 'hs-exec-picker',
		getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
		config: {
			label: 'HS on Execute Phase',
			labelTooltip: 'Use Heroic Strike on Execute Phase.',
			changedEvent: (player: Player<Spec.SpecWarrior>) => player.rotationChangeEmitter,
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().general?.useHsExec || false,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
				const newRotation = player.getRotation();
				if (!newRotation.general) {
					newRotation.general = GeneralRotation.clone(Presets.DefaultGeneralRotation.general!);
				}
				newRotation.general.useHsExec = newValue;
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
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().general?.hsRageThresh || 0,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
				const newRotation = player.getRotation();
				if (!newRotation.general) {
					newRotation.general = GeneralRotation.clone(Presets.DefaultGeneralRotation.general!);
				}
				newRotation.general.hsRageThresh = newValue;
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
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().general?.useOverpower || false,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
				const newRotation = player.getRotation();
				if (!newRotation.general) {
					newRotation.general = GeneralRotation.clone(Presets.DefaultGeneralRotation.general!);
				}
				newRotation.general.useOverpower = newValue;
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
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().general?.overpowerRageThresh || 0,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
				const newRotation = player.getRotation();
				if (!newRotation.general) {
					newRotation.general = GeneralRotation.clone(Presets.DefaultGeneralRotation.general!);
				}
				newRotation.general.overpowerRageThresh = newValue;
				player.setRotation(eventID, newRotation);
			},
			showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().general?.useOverpower == true,
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
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().general?.useHamstring || false,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
				const newRotation = player.getRotation();
				if (!newRotation.general) {
					newRotation.general = GeneralRotation.clone(Presets.DefaultGeneralRotation.general!);
				}
				newRotation.general.useHamstring = newValue;
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
			getValue: (player: Player<Spec.SpecWarrior>) => player.getRotation().general?.hamstringRageThresh || 0,
			setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => {
				const newRotation = player.getRotation();
				if (!newRotation.general) {
					newRotation.general = GeneralRotation.clone(Presets.DefaultGeneralRotation.general!);
				}
				newRotation.general.hamstringRageThresh = newValue;
				player.setRotation(eventID, newRotation);
			},
			showWhen: (player: Player<Spec.SpecWarrior>) => player.getRotation().general?.useHamstring == true,
		},
	},
	],
};