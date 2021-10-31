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
            type: 'enum', cssClass: 'rotation-enum-picker',
            getModObject: (simUI) => simUI.player,
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
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().type,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.type = newValue;
                    player.setRotation(newRotation);
                },
            },
        },
        {
            type: 'number',
            cssClass: 'num-lbs-per-cl-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'LBs per CL',
                labelTooltip: 'The number of Lightning Bolts to cast between each Chain Lightning. Only used if Rotation is set to \'Fixed LB+CL\'.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().lbsPerCl,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.lbsPerCl = newValue;
                    player.setRotation(newRotation);
                },
                enableWhen: (player) => player.getRotation().type == RotationType.FixedLBCL,
            },
        },
    ],
};
function makeBooleanShamanBuffInput(id, optionsFieldName) {
    return {
        id: id,
        states: 2,
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions()[optionsFieldName],
        setBooleanValue: (player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions[optionsFieldName] = newValue;
            player.setSpecOptions(newOptions);
        },
    };
}
