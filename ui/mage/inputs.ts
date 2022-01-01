import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
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
import { Mage_Rotation_ArcaneRotation_Filler as ArcaneFiller } from '/tbc/core/proto/mage.js';
import { Mage_Options_ArmorType as ArmorType } from '/tbc/core/proto/mage.js';

import * as Presets from './presets.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const ManaEmerald = makeBooleanMageBuffInput({ itemId: 22044 }, 'useManaEmeralds');

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
		// ********************************************************
		//                       FIRE INPUTS
		// ********************************************************
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
		// ********************************************************
		//                      ARCANE INPUTS
		// ********************************************************
		{
			type: 'enum' as const,
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				extraCssClasses: [
					'filler-enum-picker',
				],
				label: 'Filler',
				labelTooltip: 'Spells to cast while waiting for Arcane Blast stacks to drop.',
				values: [
					{
						name: 'Frostbolt', value: ArcaneFiller.Frostbolt,
					},
					{
						name: 'Arcane Missles', value: ArcaneFiller.ArcaneMissles,
					},
					{
						name: 'Scorch', value: ArcaneFiller.Fireball,
					},
					{
						name: 'Fireball', value: ArcaneFiller.Fireball,
					},
					{
						name: 'AM + FrB', value: ArcaneFiller.ArcaneMisslesFrostbolt,
					},
					{
						name: 'AM + Scorch', value: ArcaneFiller.ArcaneMisslesScorch,
					},
					{
						name: 'Scorch + 2xFiB', value: ArcaneFiller.ScorchTwoFireball,
					},
				],
				changedEvent: (player: Player<Spec.SpecMage>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecMage>) => player.getRotation().arcane?.filler || ArcaneFiller.Frostbolt,
				setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => {
					const newRotation = player.getRotation();
					if (!newRotation.arcane) {
						newRotation.arcane = ArcaneRotation.create();
					}
					newRotation.arcane.filler = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Arcane,
			},
		},
		{
			type: 'number' as const,
			cssClass: 'arcane-blasts-between-fillers-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: '# ABs between Fillers',
				labelTooltip: 'Number of Arcane Blasts to cast once the stacks drop.',
				changedEvent: (player: Player<Spec.SpecMage>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecMage>) => player.getRotation().arcane?.arcaneBlastsBetweenFillers || 3,
				setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => {
					const newRotation = player.getRotation();
					if (!newRotation.arcane) {
						newRotation.arcane = ArcaneRotation.create();
					}
					newRotation.arcane.arcaneBlastsBetweenFillers = newValue;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Arcane,
			},
		},
		{
			type: 'number' as const,
			cssClass: 'start-regen-rotation-percent-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Start regen rotation at mana %',
				labelTooltip: 'Percent of mana pool, below which the regen rotation should be used (alternate fillers and a few ABs).',
				changedEvent: (player: Player<Spec.SpecMage>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecMage>) => (player.getRotation().arcane?.startRegenRotationPercent || 0.2) * 100,
				setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => {
					const newRotation = player.getRotation();
					if (!newRotation.arcane) {
						newRotation.arcane = ArcaneRotation.create();
					}
					newRotation.arcane.startRegenRotationPercent = newValue / 100;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Arcane,
			},
		},
		{
			type: 'number' as const,
			cssClass: 'stop-regen-rotation-percent-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Stop regen rotation at mana %',
				labelTooltip: 'Percent of mana pool, above which will go back to AB spam.',
				changedEvent: (player: Player<Spec.SpecMage>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecMage>) => (player.getRotation().arcane?.stopRegenRotationPercent || 0.3) * 100,
				setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => {
					const newRotation = player.getRotation();
					if (!newRotation.arcane) {
						newRotation.arcane = ArcaneRotation.create();
					}
					newRotation.arcane.stopRegenRotationPercent = newValue / 100;
					player.setRotation(eventID, newRotation);
				},
				showWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Arcane,
			},
		},
	],
};

function makeBooleanMageBuffInput(id: ItemOrSpellId, optionsFieldName: keyof MageOptions): IconPickerConfig<Player<any>, boolean> {
  return {
    id: id,
    states: 2,
		changedEvent: (player: Player<Spec.SpecMage>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecMage>) => player.getSpecOptions()[optionsFieldName] as boolean,
		setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => {
			const newOptions = player.getSpecOptions();
      (newOptions[optionsFieldName] as boolean) = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
  }
}
