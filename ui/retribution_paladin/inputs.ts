import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { MobType, RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Target } from '/tbc/core/target.js';

import { RetributionPaladin_Rotation_ConsecrateRank as ConsecrateRank,  RetributionPaladin_Options_Judgement as Judgement } from '/tbc/core/proto/paladin.js';
import { RetributionPaladin, RetributionPaladin_Rotation as RetributionPaladinRotation, PaladinTalents as PaladinTalents, RetributionPaladin_Options as RetributionPaladinOptions } from '/tbc/core/proto/paladin.js';

import * as Presets from './presets.js';
import { SimUI } from '../core/sim_ui.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const RetributionPaladinRotationConfig = {
	inputs: [
		{
			type: 'enum' as const, cssClass: 'consecrate-rank-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Consecrate Rank',
				labelTooltip: 'Use specified rank of Consecrate during filler spell windows.',
				values: [
					{
						name: 'None', value: ConsecrateRank.None,
					},
					{
						name: 'Rank 1', value: ConsecrateRank.Rank1,
					},
					{
						name: 'Rank 4', value: ConsecrateRank.Rank4,
					},
					{
						name: 'Rank 6', value: ConsecrateRank.Rank6,
					},
				],
				changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecRetributionPaladin>) => player.getRotation().consecrateRank,
				setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: number) => {
					const newRotation = player.getRotation();
					newRotation.consecrateRank = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		},
		{
			type: 'boolean' as const, cssClass: 'exorcism-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Use Exorcism',
				labelTooltip: 'Use exorcism during filler spell windows.',
				changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecRetributionPaladin>) => player.getRotation().exorcism,
				setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.exorcism = newValue;
					player.setRotation(eventID, newRotation);
				},
				// enableWhen: (target: Target) => target.getMobType() == MobType.MobTypeDemon,
			},
		}
	],
}

export const JudgementSelection = {
	type: 'enum' as const, cssClass: 'judgement-picker',
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		label: 'Judgement',
		labelTooltip: 'Judgement debuff you will use on the target during the encounter.',
		values: [
			{
				name: 'None', value: Judgement.None,
			},
			{
				name: 'Judgement of Wisdom', value: Judgement.Wisdom,
			},
			{
				name: 'Judgement of the Crusader', value: Judgement.Crusader,
			},
		],
		changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecRetributionPaladin>) => player.getSpecOptions().judgement,
		setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			newOptions.judgement = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
}

export const CSDelay = {
	type: 'number' as const, cssClass: 'cs-delay-picker',
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		label: 'Crusader Strike Delay',
		labelTooltip: "Maximum of time (in miliseconds) we will delay Crusader Strike in order to seal twist. Experiment with values between 0 - 3000 miliseconds.",
		changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecRetributionPaladin>) => player.getSpecOptions().csDelay,
		setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			newOptions.csDelay = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
}

export const HasteLeeway = {
	type: 'number' as const, cssClass: 'haste-leeway-picker',
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		label: 'Haste Leeway',
		labelTooltip: "Arbitrary value used to account for haste procs preventing seal twists. Experiment with values between 100 - 200 miliseconds.\nDo not modify this value if you do not understand it's use.",
		changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecRetributionPaladin>) => player.getSpecOptions().hasteLeeway,
		setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			newOptions.hasteLeeway = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
}

export const DamgeTaken = {
	type: 'number' as const, cssClass: 'damage-taken-picker',
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		label: 'Damage Taken',
		labelTooltip: "Damage taken per second across the encounter. Used to model mana regeneration from Spiritual Attunement. This value should NOT include damage taken from Seal of Blood / Judgement of Blood. Leave at 0 if you do not take damage during the encounter.",
		changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<Spec.SpecRetributionPaladin>) => player.getSpecOptions().damageTaken,
		setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			newOptions.damageTaken = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	},
}