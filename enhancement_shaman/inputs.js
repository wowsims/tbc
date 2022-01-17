import { EnhancementShaman_Rotation_PrimaryShock as PrimaryShock, ShamanWeaponImbue, } from '/tbc/core/proto/shaman.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const IconBloodlust = makeBooleanShamanBuffInput(ActionId.fromSpellId(2825), 'bloodlust');
export const IconWaterShield = makeBooleanShamanBuffInput(ActionId.fromSpellId(33736), 'waterShield');
export const MainHandImbue = makeShamanWeaponImbueInput(false);
export const OffHandImbue = makeShamanWeaponImbueInput(true);
function makeShamanWeaponImbueInput(isOffHand) {
    return {
        extraCssClasses: [
            'shaman-weapon-imbue-picker',
        ],
        numColumns: 1,
        values: [
            { color: 'grey', value: ShamanWeaponImbue.ImbueNone },
            { actionId: ActionId.fromSpellId(25505), value: ShamanWeaponImbue.ImbueWindfury },
            { actionId: ActionId.fromSpellId(25489), value: ShamanWeaponImbue.ImbueFlametongue },
            { actionId: ActionId.fromSpellId(25500), value: ShamanWeaponImbue.ImbueFrostbrand },
            { actionId: ActionId.fromSpellId(25485), value: ShamanWeaponImbue.ImbueRockbiter },
        ],
        equals: (a, b) => a == b,
        zeroValue: ShamanWeaponImbue.ImbueNone,
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => (!isOffHand ? player.getSpecOptions().mainHandImbue : player.getSpecOptions().offHandImbue) || ShamanWeaponImbue.ImbueNone,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            if (!isOffHand) {
                newOptions.mainHandImbue = newValue;
            }
            else {
                newOptions.offHandImbue = newValue;
            }
            player.setSpecOptions(eventID, newOptions);
        },
    };
}
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
