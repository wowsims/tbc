import { IconInput } from '/tbc/core/components/icon_picker.js';
import { BalanceDruid_Rotation_PrimarySpell as PrimarySpell } from '/tbc/core/proto/druid.js';
import { BalanceDruid_Options as DruidOptions } from '/tbc/core/proto/druid.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ItemOrSpellId } from '/tbc/core/resources.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { Target } from '/tbc/core/target.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const SelfInnervate = {
	id: { spellId: 29166 },
	states: 2,
	changedEvent: (player: Player<Spec.SpecBalanceDruid>) => player.specOptionsChangeEmitter,
	getValue: (player: Player<Spec.SpecBalanceDruid>) => player.getSpecOptions().innervateTarget?.targetIndex != NO_TARGET,
	setBooleanValue: (player: Player<Spec.SpecBalanceDruid>, newValue: boolean) => {
		const newOptions = player.getSpecOptions();
		newOptions.innervateTarget = RaidTarget.create({
			targetIndex: newValue ? 0 : NO_TARGET,
		});
		player.setSpecOptions(newOptions);
	},
};

export const BalanceDruidRotationConfig = {
	inputs: [
		{
			type: 'enum' as const, cssClass: 'primary-spell-enum-picker',
			getModObject: (simUI: SimUI<any>) => simUI.player,
			config: {
				label: 'Primary Spell',
				labelTooltip: 'If set to \'Adaptive\', will dynamically adjust rotation based on available mana.',
				values: [
					{
						name: 'Adaptive', value: PrimarySpell.Adaptive,
					},
					{
						name: 'Starfire', value: PrimarySpell.Starfire,
					},
					{
						name: 'Starfire R6', value: PrimarySpell.Starfire6,
					},
					{
						name: 'Wrath', value: PrimarySpell.Wrath,
					},
				],
				changedEvent: (player: Player<Spec.SpecBalanceDruid>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecBalanceDruid>) => player.getRotation().primarySpell,
				setValue: (player: Player<Spec.SpecBalanceDruid>, newValue: number) => {
					const newRotation = player.getRotation();
					newRotation.primarySpell = newValue;
					player.setRotation(newRotation);
				},
			},
		},
		{
			type: 'boolean' as const,
			cssClass: 'moonfire-picker',
			getModObject: (simUI: SimUI<any>) => simUI.player,
			config: {
				label: 'Use Moonfire',
				labelTooltip: 'Use Moonfire as the next cast after the dot expires.',
				changedEvent: (player: Player<Spec.SpecBalanceDruid>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecBalanceDruid>) => player.getRotation().moonfire,
				setValue: (player: Player<Spec.SpecBalanceDruid>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.moonfire = newValue;
					player.setRotation(newRotation);
				},
				enableWhen: (player: Player<Spec.SpecBalanceDruid>) => player.getRotation().primarySpell != PrimarySpell.Adaptive,
			},
		},
		{
			type: 'boolean' as const,
			cssClass: 'faerie-fire-picker',
			getModObject: (simUI: SimUI<any>) => simUI.player,
			config: {
				label: 'Use Faerie Fire',
				labelTooltip: 'Keep Faerie Fire active on the primary target.',
				changedEvent: (player: Player<Spec.SpecBalanceDruid>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecBalanceDruid>) => player.getRotation().faerieFire,
				setValue: (player: Player<Spec.SpecBalanceDruid>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.faerieFire = newValue;
					player.setRotation(newRotation);
				},
			},
		},
		{
			type: 'boolean' as const,
			cssClass: 'insect-swarm-picker',
			getModObject: (simUI: SimUI<any>) => simUI.player,
			config: {
				label: 'Use Insect Swarm',
				labelTooltip: 'Keep Insect Swarm active on the primary target.',
				changedEvent: (player: Player<Spec.SpecBalanceDruid>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecBalanceDruid>) => player.getRotation().insectSwarm,
				setValue: (player: Player<Spec.SpecBalanceDruid>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.insectSwarm = newValue;
					player.setRotation(newRotation);
				},
				enableWhen: (player: Player<Spec.SpecBalanceDruid>) => player.getTalents().insectSwarm,
			},
		},
	],
};
