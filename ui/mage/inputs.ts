import { BalanceDruid_Rotation_PrimarySpell as PrimarySpell } from '/tbc/core/proto/druid.js';
import { BalanceDruid_Options as DruidOptions } from '/tbc/core/proto/druid.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ItemOrSpellId } from '/tbc/core/resources.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Target } from '/tbc/core/target.js';

import { Mage, Mage_Rotation as MageRotation, MageTalents as MageTalents, Mage_Options as MageOptions } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_Type as RotationType, Mage_Rotation_ArcaneRotation as ArcaneRotation, Mage_Rotation_FireRotation as FireRotation, Mage_Rotation_FrostRotation as FrostRotation } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_FireRotation_PrimarySpell as PrimaryFireSpell } from '/tbc/core/proto/mage.js';
import { Mage_Options_ArmorType as ArmorType } from '/tbc/core/proto/mage.js';

import * as Presets from './presets.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const MageArmor = {
	id: { spellId: 27125 },
	states: 2,
	extraCssClasses: [
		'mage-armor-picker',
	],
	changedEvent: (player: Player<Spec.SpecMage>) => player.specOptionsChangeEmitter,
	getValue: (player: Player<Spec.SpecMage>) => player.getSpecOptions().armor == ArmorType.MageArmor,
	setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => {
		const newOptions = player.getSpecOptions();
		newOptions.armor = newValue ? ArmorType.MageArmor : ArmorType.NoArmor;
		player.setSpecOptions(eventID, newOptions);
	},
};

export const MoltenArmor = {
	id: { spellId: 30482 },
	states: 2,
	extraCssClasses: [
		'molten-armor-picker',
	],
	changedEvent: (player: Player<Spec.SpecMage>) => player.specOptionsChangeEmitter,
	getValue: (player: Player<Spec.SpecMage>) => player.getSpecOptions().armor == ArmorType.MoltenArmor,
	setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => {
		const newOptions = player.getSpecOptions();
		newOptions.armor = newValue ? ArmorType.MoltenArmor : ArmorType.NoArmor;
		player.setSpecOptions(eventID, newOptions);
	},
};

export const EvocationTicks = {
	type: 'number' as const,
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		extraCssClasses: [
			'evocation-ticks-picker',
		],
		label: '# Evocation Ticks',
		labelTooltip: 'The number of ticks of Evocation to use, or 0 to use the full duration.',
		changedEvent: (player: Player<Spec.SpecMage>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecMage>) => player.getSpecOptions().evocationTicks,
		setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			newOptions.evocationTicks = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
};

export const MageRotationConfig = {
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
						name: 'Arcane', value: RotationType.Arcane,
					},
					{
						name: 'Fire', value: RotationType.Fire,
					},
					{
						name: 'Frost', value: RotationType.Frost,
					},
				],
				changedEvent: (simUI: IndividualSimUI<Spec.SpecMage>) => simUI.player.rotationChangeEmitter,
				getValue: (simUI: IndividualSimUI<Spec.SpecMage>) => simUI.player.getRotation().type,
				setValue: (eventID: EventID, simUI: IndividualSimUI<Spec.SpecMage>, newValue: number) => {
					const newRotation = simUI.player.getRotation();
					newRotation.type = newValue;

					if (newRotation.type == RotationType.Arcane) {
						simUI.player.setTalentsString(eventID, Presets.ArcaneTalents.data);
						if (!newRotation.arcane) {
							newRotation.arcane = ArcaneRotation.create();
						}
					} else if (newRotation.type == RotationType.Fire) {
						simUI.player.setTalentsString(eventID, Presets.FireTalents.data);
						if (!newRotation.fire) {
							newRotation.fire = FireRotation.create();
						}
					} else {
						simUI.player.setTalentsString(eventID, Presets.FrostTalents.data);
						if (!newRotation.frost) {
							newRotation.frost = FrostRotation.create();
						}
					}

					simUI.player.setRotation(eventID, newRotation);
					simUI.recomputeSettingsLayout();
				},
			},
		},
		{
			type: 'enum' as const,
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				extraCssClasses: [
					'rotation-type-enum-picker',
				],
				label: 'Primary Spell',
				values: [
					{
						name: 'Fireball', value: PrimaryFireSpell.Fireball,
					},
					{
						name: 'Scorch', value: PrimaryFireSpell.Scorch,
					},
				],
				changedEvent: (player: Player<Spec.SpecMage>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecMage>) => player.getRotation().fire?.primarySpell || PrimaryFireSpell.Fireball,
				setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => {
					const newRotation = player.getRotation();
					if (!newRotation.fire) {
						newRotation.fire = FireRotation.create();
					}
					newRotation.fire.primarySpell = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Fire,
			},
		},
		{
			type: 'boolean' as const,
			cssClass: 'maintain-improved-scorch-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Maintain Imp. Scorch',
				labelTooltip: 'Always use Scorch when below 5 stacks, or < 5.5s remaining on debuff.',
				changedEvent: (player: Player<Spec.SpecMage>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecMage>) => player.getRotation().fire?.maintainImprovedScorch || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.fire) {
						newRotation.fire = FireRotation.create();
					}
					newRotation.fire.maintainImprovedScorch = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Fire,
			},
		},
		{
			type: 'boolean' as const,
			cssClass: 'weave-fire-blast-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Weave Fire Blast',
				labelTooltip: 'Use Fire Blast whenever its off CD.',
				changedEvent: (player: Player<Spec.SpecMage>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecMage>) => player.getRotation().fire?.weaveFireBlast || false,
				setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => {
					const newRotation = player.getRotation();
					if (!newRotation.fire) {
						newRotation.fire = FireRotation.create();
					}
					newRotation.fire.weaveFireBlast = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Fire,
			},
		},
	],
};
