import { IconInput } from '../core/components/icon_picker.js';
import { ElementalShaman_Agent_AgentType as AgentType } from '../core/proto/shaman.js';
import { ElementalShaman_Options as ShamanOptions } from '../core/proto/shaman.js';
import { Spec } from '../core/proto/common.js';
import { ItemOrSpellId } from '../core/resources.js';
import { Sim } from '../core/sim.js';

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
				names: ['Adaptive', 'CL On Clearcast', 'Fixed LB+CL'],
				values: [AgentType.Adaptive, AgentType.CLOnClearcast, AgentType.FixedLBCL],
				changedEvent: (sim: Sim<Spec.SpecElementalShaman>) => sim.agentChangeEmitter,
				getValue: (sim: Sim<Spec.SpecElementalShaman>) => sim.getAgent().type,
				setValue: (sim: Sim<Spec.SpecElementalShaman>, newValue: number) => {
					const newAgent = sim.getAgent();
					newAgent.type = newValue;
					sim.setAgent(newAgent);
				},
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
