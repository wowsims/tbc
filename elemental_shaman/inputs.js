import { ElementalShaman_Agent_AgentType as AgentType } from '../core/proto/shaman.js';
export const IconBloodlust = makeBooleanShamanBuffInput({ spellId: 2825 }, 'bloodlust');
export const IconManaSpringTotem = makeBooleanShamanBuffInput({ spellId: 25570 }, 'manaSpringTotem');
export const IconTotemOfWrath = makeBooleanShamanBuffInput({ spellId: 30706 }, 'totemOfWrath');
export const IconWaterShield = makeBooleanShamanBuffInput({ spellId: 33737 }, 'waterShield');
export const IconWrathOfAirTotem = makeBooleanShamanBuffInput({ spellId: 3738 }, 'wrathOfAirTotem');
export const ElementalShamanRotationConfig = {
    inputs: [
        {
            type: 'enum',
            cssClass: 'rotation-enum-picker',
            config: {
                names: ['Adaptive', 'CL On Clearcast', 'Fixed LB+CL'],
                values: [AgentType.Adaptive, AgentType.CLOnClearcast, AgentType.FixedLBCL],
                changedEvent: (sim) => sim.agentChangeEmitter,
                getValue: (sim) => sim.getAgent().type,
                setValue: (sim, newValue) => {
                    const newAgent = sim.getAgent();
                    newAgent.type = newValue;
                    sim.setAgent(newAgent);
                },
            },
        },
    ],
};
function makeBooleanShamanBuffInput(id, optionsFieldName) {
    return {
        id: id,
        states: 2,
        changedEvent: (sim) => sim.specOptionsChangeEmitter,
        getValue: (sim) => sim.getSpecOptions()[optionsFieldName],
        setBooleanValue: (sim, newValue) => {
            const newOptions = sim.getSpecOptions();
            newOptions[optionsFieldName] = newValue;
            sim.setSpecOptions(newOptions);
        },
    };
}
