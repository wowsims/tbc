import { IconInput } from '/tbc/core/components/icon_picker.js';
import { ElementalShaman_Rotation_RotationType as RotationType } from '/tbc/core/proto/shaman.js';
import { ElementalShaman_Options as ShamanOptions } from '/tbc/core/proto/shaman.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ItemOrSpellId } from '/tbc/core/resources.js';
import { Sim } from '/tbc/core/sim.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const IconBloodlust = makeBooleanShamanBuffInput({ spellId: 2825 }, 'bloodlust');
export const IconManaSpringTotem = makeBooleanShamanBuffInput({ spellId: 25570 }, 'manaSpringTotem');
export const IconTotemOfWrath = makeBooleanShamanBuffInput({ spellId: 30706 }, 'totemOfWrath');
export const IconWaterShield = makeBooleanShamanBuffInput({ spellId: 33736 }, 'waterShield');
export const IconWrathOfAirTotem = makeBooleanShamanBuffInput({ spellId: 3738 }, 'wrathOfAirTotem');

export const ElementalShamanRotationConfig = {
	inputs: [
		{
			type: 'enum' as const,
			cssClass: 'rotation-enum-picker',
			config: {
				label: 'Type',
				values: [
					{
						name: 'Adaptive', value: RotationType.Adaptive,
						tooltip: 'Dynamically adapts based on available mana to maximize CL casts without going OOM.',
					},
					{
						name: 'CL On Clearcast', value: RotationType.CLOnClearcast,
						tooltip: 'Casts CL only after Clearcast procs.',
					},
					{
						name: 'CL On CD', value: RotationType.CLOnCD,
						tooltip: 'Casts CL if it is ready, otherwise LB.',
					},
					{
						name: 'Fixed LB+CL', value: RotationType.FixedLBCL,
						tooltip: 'Casts a fixed number of LBs between each CL (specified below), even if that means waiting. While temporary haste effects are active (drums, lust, etc) will cast extra LBs instead of waiting.',
					},
					{
						name: 'LB Only', value: RotationType.LBOnly,
						tooltip: 'Only casts Lightning Bolt.',
					},
				],
				changedEvent: (sim: Sim<Spec.SpecElementalShaman>) => sim.rotationChangeEmitter,
				getValue: (sim: Sim<Spec.SpecElementalShaman>) => sim.getRotation().type,
				setValue: (sim: Sim<Spec.SpecElementalShaman>, newValue: number) => {
					const newRotation = sim.getRotation();
					newRotation.type = newValue;
					sim.setRotation(newRotation);
				},
			},
		},
		{
			type: 'number' as const,
			cssClass: 'num-lbs-per-cl-picker',
			config: {
				label: 'LBs per CL',
				labelTooltip: 'The number of Lightning Bolts to cast between each Chain Lightning. Only used if Rotation is set to \'Fixed LB+CL\'.',
				changedEvent: (sim: Sim<Spec.SpecElementalShaman>) => sim.rotationChangeEmitter,
				getValue: (sim: Sim<Spec.SpecElementalShaman>) => sim.getRotation().lbsPerCl,
				setValue: (sim: Sim<Spec.SpecElementalShaman>, newValue: number) => {
					const newRotation = sim.getRotation();
					newRotation.lbsPerCl = newValue;
					sim.setRotation(newRotation);
				},
				enableWhen: (sim: Sim<Spec.SpecElementalShaman>) => sim.getRotation().type == RotationType.FixedLBCL,
			},
		},
	],
};

function makeBooleanShamanBuffInput(id: ItemOrSpellId, optionsFieldName: keyof ShamanOptions): IconInput {
  return {
    id: id,
    states: 2,
		changedEvent: (sim: Sim<Spec.SpecElementalShaman>) => sim.specOptionsChangeEmitter,
		getValue: (sim: Sim<Spec.SpecElementalShaman>) => sim.getSpecOptions()[optionsFieldName],
		setBooleanValue: (sim: Sim<Spec.SpecElementalShaman>, newValue: boolean) => {
			const newOptions = sim.getSpecOptions();
      (newOptions[optionsFieldName] as boolean) = newValue;
			sim.setSpecOptions(newOptions);
		},
  }
}
