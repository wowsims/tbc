import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { IconEnumPicker, IconEnumPickerConfig } from '/tbc/core/components/icon_enum_picker.js';
import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { Hunter_Options as HunterOptions } from '/tbc/core/proto/shaman.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Target } from '/tbc/core/target.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

import {
	Hunter,
	Hunter_Rotation as HunterRotation,
	Hunter_Options as HunterOptions,
	Hunter_Options_Ammo as Ammo,
	Hunter_Options_QuiverBonus as QuiverBonus,
} from '/tbc/core/proto/shaman.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const IconBloodlust = makeBooleanHunterBuffInput(ActionId.fromSpellId(2825), 'bloodlust');
export const IconWaterShield = makeBooleanHunterBuffInput(ActionId.fromSpellId(33736), 'waterShield');
export const MainHandImbue = makeHunterWeaponImbueInput(false);
export const OffHandImbue = makeHunterWeaponImbueInput(true);

export const Quiver = {
	extraCssClasses: [
		'quiver-picker',
	],
	numColumns: 1,
	values: [
		{ color: '82e89d', value: QuiverBonus.QuiverNone },
		{ actionId: ActionId.fromItemId(18714), value: QuiverBonus.Speed15 },
		{ actionId: ActionId.fromItemId(2662), value: QuiverBonus.Speed14 },
		{ actionId: ActionId.fromItemId(8217), value: QuiverBonus.Speed13 },
		{ actionId: ActionId.fromItemId(7371), value: QuiverBonus.Speed12 },
		{ actionId: ActionId.fromItemId(3605), value: QuiverBonus.Speed11 },
		{ actionId: ActionId.fromItemId(3573), value: QuiverBonus.Speed10 },
	],
	equals: (a: QuiverBonus, b: QuiverBonus) => a == b,
	zeroValue: QuiverBonus.ImbueNone,
	changedEvent: (player: Player<Spec.SpecHunter>) => player.specOptionsChangeEmitter,
	getValue: (player: Player<Spec.SpecHunter>) => player.getSpecOptions().quiverBonus,
	setValue: (eventID: EventID, player: Player<Spec.SpecHunter>, newValue: number) => {
		const newOptions = player.getSpecOptions();
		newOptions.quiverBonus = newValue;
		player.setSpecOptions(eventID, newOptions);
	},
};

export const DelayOffhandSwings = {
	type: 'boolean' as const,
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		extraCssClasses: [
			'delay-offhand-swings-picker',
		],
		label: 'Delay Offhand Swings',
		labelTooltip: 'Uses the startattack macro to delay OH swings, so they always follow within 0.5s of a MH swing.',
		changedEvent: (player: Player<Spec.SpecHunter>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecHunter>) => player.getSpecOptions().delayOffhandSwings,
		setValue: (eventID: EventID, player: Player<Spec.SpecHunter>, newValue: boolean) => {
			const newOptions = player.getSpecOptions();
			newOptions.delayOffhandSwings = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
};

export const HunterRotationConfig = {
	inputs: [
		{
			type: 'enum' as const, cssClass: 'primary-shock-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Primary Shock',
				values: [
					{
						name: 'None', value: PrimaryShock.None,
					},
					{
						name: 'Earth Shock', value: PrimaryShock.Earth,
					},
					{
						name: 'Frost Shock', value: PrimaryShock.Frost,
					},
				],
				changedEvent: (player: Player<Spec.SpecHunter>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecHunter>) => player.getRotation().primaryShock,
				setValue: (eventID: EventID, player: Player<Spec.SpecHunter>, newValue: number) => {
					const newRotation = player.getRotation();
					newRotation.primaryShock = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		},
		{
			type: 'boolean' as const, cssClass: 'weave-flame-shock-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Weave Flame Shock',
				labelTooltip: 'Use Flame Shock whenever the target does not already have the DoT.',
				changedEvent: (player: Player<Spec.SpecHunter>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecHunter>) => player.getRotation().weaveFlameShock,
				setValue: (eventID: EventID, player: Player<Spec.SpecHunter>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.weaveFlameShock = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		},
	],
};

function makeBooleanHunterBuffInput(id: ActionId, optionsFieldName: keyof HunterOptions): IconPickerConfig<Player<any>, boolean> {
	return {
	  id: id,
	  states: 2,
		changedEvent: (player: Player<Spec.SpecHunter>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecHunter>) => player.getSpecOptions()[optionsFieldName] as boolean,
		setValue: (eventID: EventID, player: Player<Spec.SpecHunter>, newValue: boolean) => {
			const newOptions = player.getSpecOptions();
	(newOptions[optionsFieldName] as boolean) = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	};
}
