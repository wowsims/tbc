import { ProtectionWarrior_Rotation_DemoShout as DemoShout, ProtectionWarrior_Rotation_ThunderClap as ThunderClap } from '/tbc/core/proto/warrior.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const ProtectionWarriorRotationConfig = {
    inputs: [
        {
            type: 'enum', cssClass: 'demo-shout-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Demo Shout',
                values: [
                    { name: 'None', value: DemoShout.DemoShoutNone },
                    { name: 'Maintain Debuff', value: DemoShout.DemoShoutMaintain },
                    { name: 'Filler', value: DemoShout.DemoShoutFiller },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().demoShout,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.demoShout = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'enum', cssClass: 'thunder-clap-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Thunder Clap',
                values: [
                    { name: 'None', value: ThunderClap.ThunderClapNone },
                    { name: 'Maintain Debuff', value: ThunderClap.ThunderClapMaintain },
                    { name: 'On CD', value: ThunderClap.ThunderClapOnCD },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().thunderClap,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.thunderClap = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
    ],
};
