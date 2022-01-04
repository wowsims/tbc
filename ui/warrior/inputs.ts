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
import { Warrior_Rotation_Type as RotationType, Warrior_Rotation_ArmsSlamRotation as ArmsSlamRotation, Warrior_Rotation_ArmsDWRotation as ArmsDWRotation, Warrior_Rotation_FuryRotation as FuryRotation } from '/tbc/core/proto/warrior.js';
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

					if (newRotation.type == RotationType.ArmsSlam) {
						simUI.player.setTalentsString(eventID, Presets.ArmsSlamTalents.data);
						if (!newRotation.armsSiam) {
							newRotation.armsSiam = ArmsSlamRotation.create();
						}
					} else if (newRotation.type == RotationType.ArmsDW) {
						simUI.player.setTalentsString(eventID, Presets.ArmsDWTalents.data);
						if (!newRotation.armsDw) {
							newRotation.armsDw = ArmsDWRotation.create();
						}
					} else {
						simUI.player.setTalentsString(eventID, Presets.FuryTalents.data);
						if (!newRotation.fury) {
							newRotation.fury = FuryRotation.create();
						}
					}

					simUI.player.setRotation(eventID, newRotation);
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
		// ********************************************************
		//                      ARMS SLAM INPUTS
		// ********************************************************

		// ********************************************************
		//                      ARMS DW INPUTS
		// ********************************************************
	],
};

function makeBooleanWarriorBuffInput(id: ItemOrSpellId, optionsFieldName: keyof WarriorOptions): IconPickerConfig<Player<any>, boolean> {
  return {
    id: id,
    states: 2,
		changedEvent: (player: Player<Spec.SpecWarrior>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecWarrior>) => player.getSpecOptions()[optionsFieldName] as boolean,
		setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => {
			const newOptions = player.getSpecOptions();
      (newOptions[optionsFieldName] as boolean) = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
  }
}
