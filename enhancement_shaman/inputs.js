import { makeWeaponImbueInput } from '/tbc/core/components/icon_inputs.js';
import { EarthTotem, EnhancementShaman_Rotation_PrimaryShock as PrimaryShock, } from '/tbc/core/proto/shaman.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const IconBloodlust = makeBooleanShamanBuffInput(ActionId.fromSpellId(2825), 'bloodlust');
export const IconWaterShield = makeBooleanShamanBuffInput(ActionId.fromSpellId(33736), 'waterShield');
const imbueOptions = [
    WeaponImbue.WeaponImbueShamanWindfury,
    WeaponImbue.WeaponImbueShamanFlametongue,
    WeaponImbue.WeaponImbueShamanFrostbrand,
    WeaponImbue.WeaponImbueShamanRockbiter,
];
export const MainHandImbue = makeWeaponImbueInput(true, imbueOptions);
export const OffHandImbue = makeWeaponImbueInput(false, imbueOptions);
export const DelayOffhandSwings = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'delay-offhand-swings-picker',
        ],
        label: 'Delay Offhand Swings',
        labelTooltip: 'Uses the startattack macro to delay OH swings, so they always follow within 0.5s of a MH swing.',
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().delayOffhandSwings,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.delayOffhandSwings = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
export const SnapshotT42Pc = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'snapshot-t4-2pc-picker',
        ],
        label: 'Snapshot T4 2pc',
        labelTooltip: 'Snapshots the improved Strength of Earth totem bonus from T4 2pc (+12 strength) for the first 1:50s of the fight. Only works if the selected Earth totem is Strength of Earth Totem.',
        changedEvent: (player) => player.changeEmitter,
        getValue: (player) => player.getSpecOptions().snapshotT42Pc,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.snapshotT42Pc = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
        enableWhen: (player) => player.getRotation().totems?.earth == EarthTotem.StrengthOfEarthTotem,
    },
};
export const EnhancementShamanRotationConfig = {
    inputs: [
        {
            type: 'enum', cssClass: 'primary-shock-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Primary Shock',
                values: [
                    {
                        name: 'None', value: PrimaryShock.None,
                    },
                    {
                        name: 'Earth Shock', value: PrimaryShock.Earth,
                    },
                    {
                        name: 'Frost Shock', value: PrimaryShock.Frost,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().primaryShock,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.primaryShock = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean', cssClass: 'weave-flame-shock-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Weave Flame Shock',
                labelTooltip: 'Use Flame Shock whenever the target does not already have the DoT.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().weaveFlameShock,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.weaveFlameShock = newValue;
                    player.setRotation(eventID, newRotation);
                },
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
