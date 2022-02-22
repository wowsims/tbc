import { ElementalShaman_Rotation_RotationType as RotationType } from '/tbc/core/proto/shaman.js';
import { AirTotem } from '/tbc/core/proto/shaman.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const IconBloodlust = makeBooleanShamanBuffInput(ActionId.fromSpellId(2825), 'bloodlust');
export const IconWaterShield = makeBooleanShamanBuffInput(ActionId.fromSpellId(33736), 'waterShield');
export const SnapshotT42Pc = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'snapshot-t4-2pc-picker',
        ],
        label: 'Snapshot T4 2pc',
        labelTooltip: 'Snapshots the improved wrath of air totem bonus from T4 2pc (+20 spell power) for the first 1:50s of the fight. Only works if the selected air totem is Wrath of Air Totem.',
        changedEvent: (player) => player.changeEmitter,
        getValue: (player) => player.getSpecOptions().snapshotT42Pc,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.snapshotT42Pc = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
        enableWhen: (player) => player.getRotation().totems?.air == AirTotem.WrathOfAirTotem,
    },
};
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
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.type = newValue;
                    player.setRotation(eventID, newRotation);
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
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.lbsPerCl = newValue;
                    player.setRotation(eventID, newRotation);
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
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions[optionsFieldName] = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    };
}
