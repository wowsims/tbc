import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { WarriorShout, Warrior_Rotation_SunderArmor as SunderArmor, } from '/tbc/core/proto/warrior.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const Recklessness = {
    id: ActionId.fromSpellId(1719),
    states: 2,
    extraCssClasses: [
        'warrior-recklessness-picker',
    ],
    changedEvent: (player) => player.specOptionsChangeEmitter,
    getValue: (player) => player.getSpecOptions().useRecklessness,
    setValue: (eventID, player, newValue) => {
        const newOptions = player.getSpecOptions();
        newOptions.useRecklessness = newValue;
        player.setSpecOptions(eventID, newOptions);
    },
};
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
export const ShoutPicker = {
    extraCssClasses: [
        'shout-picker',
    ],
    numColumns: 1,
    values: [
        { color: 'c79c6e', value: WarriorShout.WarriorShoutNone },
        { actionId: ActionId.fromSpellId(2048), value: WarriorShout.WarriorShoutBattle },
        { actionId: ActionId.fromSpellId(469), value: WarriorShout.WarriorShoutCommanding },
    ],
    equals: (a, b) => a == b,
    zeroValue: WarriorShout.WarriorShoutNone,
    changedEvent: (player) => player.specOptionsChangeEmitter,
    getValue: (player) => player.getSpecOptions().shout,
    setValue: (eventID, player, newValue) => {
        const newOptions = player.getSpecOptions();
        newOptions.shout = newValue;
        player.setSpecOptions(eventID, newOptions);
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
export const WarriorRotationConfig = {
    inputs: [
        {
            type: 'boolean',
            cssClass: 'cleave-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Cleave',
                labelTooltip: 'Use Cleave instead of Heroic Strike.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useCleave,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useCleave = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'overpower-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Overpower',
                labelTooltip: 'Use Overpower when available.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useOverpower,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useOverpower = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'hamstring-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Hamstring',
                labelTooltip: 'Use Hamstring on free globals.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useHamstring,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useHamstring = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'slam-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Slam',
                labelTooltip: 'Use Slam whenever possible.',
                changedEvent: (player) => TypedEvent.onAny([player.rotationChangeEmitter, player.talentsChangeEmitter]),
                getValue: (player) => player.getRotation().useSlam,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useSlam = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getTalents().improvedSlam == 2,
            },
        },
        {
            type: 'boolean',
            cssClass: 'prioritize-ww-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Prioritize WW',
                labelTooltip: 'Prioritize Whirlwind over Bloodthirst or Mortal Strike.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().prioritizeWw,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.prioritizeWw = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'number',
            cssClass: 'hs-rage-threshold',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'HS rage threshold',
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
        {
            type: 'number',
            cssClass: 'overpower-rage-threshold',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Overpower rage threshold',
                labelTooltip: 'Use Overpower when rage is below a point.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().overpowerRageThreshold,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.overpowerRageThreshold = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().useOverpower,
            },
        },
        {
            type: 'number',
            cssClass: 'hamstring-rage-threshold',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Hamstring rage threshold',
                labelTooltip: 'Hamstring will only be used when rage is larger than this value.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().hamstringRageThreshold,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.hamstringRageThreshold = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().useHamstring,
            },
        },
        {
            type: 'number',
            cssClass: 'slam-latency',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Slam Latency',
                labelTooltip: 'Time between MH swing and start of the Slam cast, in milliseconds.',
                changedEvent: (player) => TypedEvent.onAny([player.rotationChangeEmitter, player.talentsChangeEmitter]),
                getValue: (player) => player.getRotation().slamLatency,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.slamLatency = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().useSlam && player.getTalents().improvedSlam == 2,
            },
        },
        {
            type: 'number',
            cssClass: 'slam-gcd-delay',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'experimental',
                ],
                label: 'Slam GCD Delay',
                labelTooltip: 'Amount of time Slam may delay the GCD, in milliseconds.',
                changedEvent: (player) => TypedEvent.onAny([player.rotationChangeEmitter, player.talentsChangeEmitter]),
                getValue: (player) => player.getRotation().slamGcdDelay,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.slamGcdDelay = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().useSlam && player.getTalents().improvedSlam == 2,
            },
        },
        {
            type: 'number',
            cssClass: 'slam-ms-ww-delay',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'experimental',
                ],
                label: 'Slam MS+WW Delay',
                labelTooltip: 'Amount of time Slam may delay MS+WW, in milliseconds.',
                changedEvent: (player) => TypedEvent.onAny([player.rotationChangeEmitter, player.talentsChangeEmitter]),
                getValue: (player) => player.getRotation().slamMsWwDelay,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.slamMsWwDelay = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().useSlam && player.getTalents().improvedSlam == 2,
            },
        },
        {
            type: 'number',
            cssClass: 'rampage-duration-threshold',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Rampage Refresh Time',
                labelTooltip: 'Refresh Rampage when the remaining duration is less than this amount of time (seconds).',
                changedEvent: (player) => TypedEvent.onAny([player.rotationChangeEmitter, player.talentsChangeEmitter]),
                getValue: (player) => player.getRotation().rampageCdThreshold,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.rampageCdThreshold = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getTalents().rampage,
            },
        },
        {
            type: 'boolean',
            cssClass: 'hs-exec-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'HS during Execute Phase',
                labelTooltip: 'Use Heroic Strike during Execute Phase.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useHsDuringExecute,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useHsDuringExecute = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'bt-exec-picker-fury',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'BT during Execute Phase',
                labelTooltip: 'Use Bloodthirst during Execute Phase.',
                changedEvent: (player) => TypedEvent.onAny([player.rotationChangeEmitter, player.talentsChangeEmitter]),
                getValue: (player) => player.getRotation().useBtDuringExecute,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useBtDuringExecute = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getTalents().bloodthirst,
            },
        },
        {
            type: 'boolean',
            cssClass: 'ms-exec-picker-fury',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'MS during Execute Phase',
                labelTooltip: 'Use Mortal Strike during Execute Phase.',
                changedEvent: (player) => TypedEvent.onAny([player.rotationChangeEmitter, player.talentsChangeEmitter]),
                getValue: (player) => player.getRotation().useMsDuringExecute,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useMsDuringExecute = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getTalents().mortalStrike,
            },
        },
        {
            type: 'boolean',
            cssClass: 'ww-exec-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'WW during Execute Phase',
                labelTooltip: 'Use Whirlwind during Execute Phase.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useWwDuringExecute,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useWwDuringExecute = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'slam-exec-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Slam during Execute Phase',
                labelTooltip: 'Use Slam during Execute Phase.',
                changedEvent: (player) => TypedEvent.onAny([player.rotationChangeEmitter, player.talentsChangeEmitter]),
                getValue: (player) => player.getRotation().useSlamDuringExecute,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useSlamDuringExecute = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().useSlam && player.getTalents().improvedSlam == 2,
            },
        },
        {
            type: 'enum', cssClass: 'sunder-armor-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Sunder Armor',
                values: [
                    { name: 'Never', value: SunderArmor.SunderArmorNone },
                    { name: 'Help Stack', value: SunderArmor.SunderArmorHelpStack },
                    { name: 'Maintain Debuff', value: SunderArmor.SunderArmorMaintain },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().sunderArmor,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.sunderArmor = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'maintain-demo-shout-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Maintain Demo Shout',
                labelTooltip: 'Keep Demo Shout active on the primary target.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().maintainDemoShout,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.maintainDemoShout = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'maintain-thunder-clap-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Maintain Thunder Clap',
                labelTooltip: 'Keep Thunder Clap active on the primary target.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().maintainThunderClap,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.maintainThunderClap = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
    ],
};
