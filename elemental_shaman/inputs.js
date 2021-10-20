import { ElementalShaman_Rotation_RotationType as RotationType } from '/tbc/core/proto/shaman.js';
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
            type: 'enum',
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
                changedEvent: (sim) => sim.rotationChangeEmitter,
                getValue: (sim) => sim.getRotation().type,
                setValue: (sim, newValue) => {
                    const newRotation = sim.getRotation();
                    newRotation.type = newValue;
                    sim.setRotation(newRotation);
                },
            },
        },
        {
            type: 'number',
            cssClass: 'num-lbs-per-cl-picker',
            config: {
                label: 'LBs per CL',
                labelTooltip: 'The number of Lightning Bolts to cast between each Chain Lightning. Only used if Rotation is set to \'Fixed LB+CL\'.',
                changedEvent: (sim) => sim.rotationChangeEmitter,
                getValue: (sim) => sim.getRotation().lbsPerCl,
                setValue: (sim, newValue) => {
                    const newRotation = sim.getRotation();
                    newRotation.lbsPerCl = newValue;
                    sim.setRotation(newRotation);
                },
                enableWhen: (sim) => sim.getRotation().type == RotationType.FixedLBCL,
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
