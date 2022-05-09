import { TypedEvent } from '/tbc/core/typed_event.js';
import { WarriorShout, ProtectionWarrior_Rotation_DemoShout as DemoShout, ProtectionWarrior_Rotation_ShieldBlock as ShieldBlock, ProtectionWarrior_Rotation_ThunderClap as ThunderClap } from '/tbc/core/proto/warrior.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const ShoutPicker = {
    type: 'enum', cssClass: 'shout-picker',
    getModObject: (simUI) => simUI.player,
    config: {
        label: 'Shout',
        labelTooltip: 'Shout buff to maintain.',
        values: [
            { name: 'None', value: WarriorShout.WarriorShoutNone },
            { name: 'Battle Shout', value: WarriorShout.WarriorShoutBattle },
            { name: 'Commanding Shout', value: WarriorShout.WarriorShoutCommanding },
        ],
        changedEvent: (player) => player.rotationChangeEmitter,
        getValue: (player) => player.getSpecOptions().shout,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.shout = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
export const PrecastShout = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'precast-shout-picker',
        ],
        label: 'Precast Shout',
        labelTooltip: 'Selected shout is cast 10 seconds before combat starts.',
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().precastShout,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.precastShout = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
export const PrecastShoutWithSapphire = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'precast-shout-with-sapphire-picker',
        ],
        label: 'Precast with Sapphire',
        labelTooltip: 'Snapshot bonus from Solarian\'s Sapphire (+70 attack power) with precast shout.',
        changedEvent: (player) => TypedEvent.onAny([player.specOptionsChangeEmitter, player.gearChangeEmitter]),
        getValue: (player) => player.getSpecOptions().precastShoutSapphire,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.precastShoutSapphire = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
        enableWhen: (player) => player.getSpecOptions().shout == WarriorShout.WarriorShoutBattle && player.getSpecOptions().precastShout && !player.getGear().hasTrinket(30446),
    },
};
export const PrecastShoutWithT2 = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'precast-shout-with-t2-picker',
        ],
        label: 'Precast with T2',
        labelTooltip: 'Snapshot T2 set bonus (+30 attack power) with precast shout.',
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().precastShoutT2,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.precastShoutT2 = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
        enableWhen: (player) => player.getSpecOptions().shout == WarriorShout.WarriorShoutBattle && player.getSpecOptions().precastShout,
    },
};
export const ProtectionWarriorRotationConfig = {
    inputs: [
        {
            type: 'enum', cssClass: 'demo-shout-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Demo Shout',
                values: [
                    { name: 'Never', value: DemoShout.DemoShoutNone },
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
                    { name: 'Never', value: ThunderClap.ThunderClapNone },
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
        {
            type: 'enum', cssClass: 'shield-block-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Shield Block',
                labelTooltip: 'When to use shield block.',
                values: [
                    { name: 'Never', value: ShieldBlock.ShieldBlockNone },
                    { name: 'To Proc Revenge', value: ShieldBlock.ShieldBlockToProcRevenge },
                    { name: 'On CD', value: ShieldBlock.ShieldBlockOnCD },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().shieldBlock,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.shieldBlock = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'number', cssClass: 'heroic-strike-threshold-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'HS Threshold',
                labelTooltip: 'Heroic Strike when rage is above:',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().hsRageThreshold,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.hsRageThreshold = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
    ],
};
function makeBooleanBuffInput(id, optionsFieldName) {
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
