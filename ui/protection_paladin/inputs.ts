import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { EventID } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';

import {
	PaladinAura as PaladinAura,
	PaladinJudgement as PaladinJudgement,
	ProtectionPaladin_Rotation as ProtectionPaladinRotation,
	ProtectionPaladin_Options as ProtectionPaladinOptions,
	ProtectionPaladin_Options_PrimaryJudgement as PrimaryJudgement,
} from '/tbc/core/proto/paladin.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const ProtectionPaladinRotationConfig = {
	inputs: [
		{
			type: 'enum' as const, cssClass: 'consecration-rank-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Consecration Rank',
				labelTooltip: 'Use specified rank of Consecration during filler spell windows.',
				values: [
					{
						name: 'None', value: 0,
					},
					{
						name: 'Rank 1', value: 1,
					},
					{
						name: 'Rank 4', value: 4,
					},
					{
						name: 'Rank 6', value: 6,
					},
				],
				changedEvent: (player: Player<Spec.SpecProtectionPaladin>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecProtectionPaladin>) => player.getRotation().consecrationRank,
				setValue: (eventID: EventID, player: Player<Spec.SpecProtectionPaladin>, newValue: number) => {
					const newRotation = player.getRotation();
					newRotation.consecrationRank = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		},
		{
			type: 'boolean' as const, cssClass: 'exorcism-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Use Exorcism',
				labelTooltip: 'Use Exorcism during filler spell windows. Will only be used if the boss mob type is Undead or Demon.',
				changedEvent: (player: Player<Spec.SpecProtectionPaladin>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecProtectionPaladin>) => player.getRotation().useExorcism,
				setValue: (eventID: EventID, player: Player<Spec.SpecProtectionPaladin>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.useExorcism = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		}
	],
}

export const AuraSelection = {
	type: 'enum' as const, cssClass: 'aura-picker',
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		label: 'Aura',
		values: [
			{ name: 'None', value: PaladinAura.NoPaladinAura },
			{ name: 'Sanctity Aura', value: PaladinAura.SanctityAura },
			{ name: 'Devotion Aura', value: PaladinAura.DevotionAura },
			{ name: 'Retribution Aura', value: PaladinAura.RetributionAura },
		],
		changedEvent: (player: Player<Spec.SpecProtectionPaladin>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecProtectionPaladin>) => player.getSpecOptions().aura,
		setValue: (eventID: EventID, player: Player<Spec.SpecProtectionPaladin>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			newOptions.aura = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
}

export const PrimaryJudgementSelection = {
	type: 'enum' as const, cssClass: 'judgement-picker',
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		label: 'Primary Judgement',
		values: [
			{
				name: 'Vengeance', value: PrimaryJudgement.Vengeance,
			},
			{
				name: 'Righteousness', value: PrimaryJudgement.Righteousness,
			},
			{
				name: 'Twist', value: PrimaryJudgement.Twist,
			},
		],
		changedEvent: (player: Player<Spec.SpecProtectionPaladin>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecProtectionPaladin>) => player.getSpecOptions().primaryJudgement,
		setValue: (eventID: EventID, player: Player<Spec.SpecProtectionPaladin>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			newOptions.primaryJudgement = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
}
