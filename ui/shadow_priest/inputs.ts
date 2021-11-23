import { IconInput } from '/tbc/core/components/icon_picker.js';
import { ShadowPriest_Rotation_RotationType as RotationType } from '/tbc/core/proto/priest.js';
import { Race, RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ItemOrSpellId } from '/tbc/core/resources.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { Target } from '/tbc/core/target.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const ShadowPriestRotationConfig = {
	inputs: [
		{
			type: 'enum' as const, cssClass: 'rotation-enum-picker',
			getModObject: (simUI: SimUI<any>) => simUI.player,
			config: {
				label: 'Rotation Type',
				labelTooltip: 'Choose how to clip your mindflay',
				values: [
					{
						name: 'Basic', value: RotationType.Basic,
					},
					{
						name: 'Clip Always', value: RotationType.ClipAlways,
					},
					{
						name: 'Intelligent', value: RotationType.IntelligentClipping,
					},
				],
				changedEvent: (player: Player<Spec.SpecShadowPriest>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecShadowPriest>) => player.getRotation().rotationType,
				setValue: (player: Player<Spec.SpecShadowPriest>, newValue: number) => {
					const newRotation = player.getRotation();
					newRotation.rotationType = newValue;
					player.setRotation(newRotation);
				},
			},
		},
		{
			type: 'boolean' as const,
			cssClass: 'devplague-picker',
			getModObject: (simUI: SimUI<any>) => simUI.player,
			config: {
				label: 'Use Devouring Plague',
				labelTooltip: 'Use Devouring Plague whenever off CD.',
				changedEvent: (player: Player<Spec.SpecShadowPriest>) => player.raceChangeEmitter,
				getValue: (player: Player<Spec.SpecShadowPriest>) => player.getRotation().useDevPlague,
				setValue: (player: Player<Spec.SpecShadowPriest>, newValue: boolean) => {
					const newRotation = player.getRotation();
					newRotation.useDevPlague = newValue;
					player.setRotation(newRotation);
				},
				enableWhen: (player: Player<Spec.SpecShadowPriest>) => player.getRace() == Race.RaceUndead,
			},
		},
	],
};
