import { FeralTankDruid_Rotation_Swipe as Swipe } from '/tbc/core/proto/druid.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const StartingRage = {
    type: 'number',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'starting-rage-picker',
        ],
        label: 'Starting Rage',
        labelTooltip: 'Initial rage at the start of each iteration.',
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().startingRage,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.startingRage = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
export const FeralTankDruidRotationConfig = {
    inputs: [
        {
            type: 'number', cssClass: 'maul-threshold-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Maul Threshold',
                labelTooltip: 'Queue Maul when rage is above this value.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().maulRageThreshold,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.maulRageThreshold = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'enum', cssClass: 'swipe-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Swipe',
                values: [
                    { name: 'Never', value: Swipe.SwipeNone },
                    { name: 'With Enough AP', value: Swipe.SwipeWithEnoughAP },
                    { name: 'Spam', value: Swipe.SwipeSpam },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().swipe,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.swipe = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'number', cssClass: 'swipe-ap-threshold-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Swipe AP Threshold',
                labelTooltip: 'Use Swipe when Attack Power is larger than this amount.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().swipeApThreshold,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.swipeApThreshold = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => player.getRotation().swipe == Swipe.SwipeWithEnoughAP,
            },
        },
        {
            type: 'boolean',
            cssClass: 'maintain-demo-roar-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Maintain Demo Roar',
                labelTooltip: 'Keep Demoralizing Roar active on the primary target. If a stronger debuff is active, will not cast.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().maintainDemoralizingRoar,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.maintainDemoralizingRoar = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'maintain-faerie-fire-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Maintain Faerie Fire',
                labelTooltip: 'Keep Faerie Fire active on the primary target. If a stronger debuff is active, will not cast.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().maintainFaerieFire,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.maintainFaerieFire = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
    ],
};
