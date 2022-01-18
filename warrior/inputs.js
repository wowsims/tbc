import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Warrior_Rotation_Type as RotationType, Warrior_Rotation_ArmsSlamRotation as ArmsSlamRotation, Warrior_Rotation_ArmsDWRotation as ArmsDWRotation, Warrior_Rotation_FuryRotation as FuryRotation } from '/tbc/core/proto/warrior.js';
import { Warrior_Rotation_FuryRotation_PrimaryInstant as PrimaryInstant } from '/tbc/core/proto/warrior.js';
import * as Presets from './presets.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const Recklessness = {
    id: ActionId.fromSpellId(1719),
    states: 2,
    extraCssClasses: [
        'warrior-recklessness-picker',
    ],
    changedEvent: (player) => player.specOptionsChangeEmitter,
    getValue: (player) => player.getSpecOptions().recklessness,
    setValue: (eventID, player, newValue) => {
        const newOptions = player.getSpecOptions();
        newOptions.recklessness = newValue;
        player.setSpecOptions(eventID, newOptions);
    },
};
export const WarriorRotationConfig = {
    inputs: [
        {
            type: 'enum',
            getModObject: (simUI) => simUI,
            config: {
                extraCssClasses: [
                    'rotation-type-enum-picker',
                ],
                label: 'Spec',
                labelTooltip: 'Switches between spec rotation settings. Will also update talents to defaults for the selected spec.',
                values: [
                    {
                        name: 'Arms Slam', value: RotationType.ArmsSlam,
                    },
                    {
                        name: 'Arms DW', value: RotationType.ArmsDW,
                    },
                    {
                        name: 'Fury', value: RotationType.Fury,
                    },
                ],
                changedEvent: (simUI) => simUI.player.rotationChangeEmitter,
                getValue: (simUI) => simUI.player.getRotation().type,
                setValue: (eventID, simUI, newValue) => {
                    const newRotation = simUI.player.getRotation();
                    newRotation.type = newValue;
                    TypedEvent.freezeAllAndDo(() => {
                        if (newRotation.type == RotationType.Fury) {
                            simUI.player.setTalentsString(eventID, Presets.FuryTalents.data);
                            if (!newRotation.fury) {
                                newRotation.fury = FuryRotation.clone(Presets.DefaultFuryRotation.fury);
                            }
                        }
                        else if (newRotation.type == RotationType.ArmsSlam) {
                            simUI.player.setTalentsString(eventID, Presets.ArmsSlamTalents.data);
                            if (!newRotation.armsSlam) {
                                newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam);
                            }
                        }
                        else {
                            simUI.player.setTalentsString(eventID, Presets.ArmsDWTalents.data);
                            if (!newRotation.armsDw) {
                                newRotation.armsDw = ArmsDWRotation.clone(Presets.DefaultArmsDWRotation.armsDw);
                            }
                        }
                        simUI.player.setRotation(eventID, newRotation);
                    });
                    simUI.recomputeSettingsLayout();
                },
            },
        },
        // ********************************************************
        //                       FURY INPUTS
        // ********************************************************
        {
            type: 'enum',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'rotation-type-enum-picker',
                ],
                label: 'Primary Instant',
                labelTooltip: 'Main instant ability that will be prioritized above everything else while it is off CD.',
                values: [
                    {
                        name: 'Bloodthirst', value: PrimaryInstant.Bloodthirst,
                    },
                    {
                        name: 'Whirlwind', value: PrimaryInstant.Whirlwind,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().fury?.primaryInstant || PrimaryInstant.Whirlwind,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.fury) {
                        newRotation.fury = FuryRotation.create();
                    }
                    newRotation.fury.primaryInstant = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Fury,
            },
        },
        {
            type: 'boolean',
            cssClass: 'bt-exec-picker-fury',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'BT during Execute Phase',
                labelTooltip: 'Use Bloodthirst during Execute Phase.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().fury?.useBtDuringExecute || false,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.fury) {
                        newRotation.fury = FuryRotation.clone(Presets.DefaultFuryRotation.fury);
                    }
                    newRotation.fury.useBtDuringExecute = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Fury,
            },
        },
        {
            type: 'number',
            cssClass: 'rampage-duration-threshold',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Rampage refresh timing (seconds)',
                labelTooltip: 'Refresh rampage when it has certain duration left on it.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().fury?.rampageCdThreshold || 0,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.fury) {
                        newRotation.fury = FuryRotation.clone(Presets.DefaultFuryRotation.fury);
                    }
                    newRotation.fury.rampageCdThreshold = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Fury,
            },
        },
        // ********************************************************
        //                      ARMS SLAM INPUTS
        // ********************************************************
        {
            type: 'boolean',
            cssClass: 'ms-exec-picker-arms-slam',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'MS during Execute Phase',
                labelTooltip: 'Use Mortal Strike during Execute Phase.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().armsSlam?.useMsDuringExecute || false,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.armsSlam) {
                        newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam);
                    }
                    newRotation.armsSlam.useMsDuringExecute = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.ArmsSlam,
            },
        },
        {
            type: 'boolean',
            cssClass: 'slam-exec-picker-arms-slam',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Slam during Execute Phase',
                labelTooltip: 'Use Slam during Execute Phase.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().armsSlam?.useSlamDuringExecute || false,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.armsSlam) {
                        newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam);
                    }
                    newRotation.armsSlam.useSlamDuringExecute = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.ArmsSlam,
            },
        },
        {
            type: 'number',
            cssClass: 'slam-latency',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Slam Latency (ms)',
                labelTooltip: 'Add delay to slam casting.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().armsSlam?.slamLatency || 0,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.armsSlam) {
                        newRotation.armsSlam = ArmsSlamRotation.clone(Presets.DefaultArmsSlamRotation.armsSlam);
                    }
                    newRotation.armsSlam.slamLatency = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.ArmsSlam,
            },
        },
        // ********************************************************
        //                      ARMS DW INPUTS
        // ********************************************************
        {
            type: 'boolean',
            cssClass: 'ms-exec-picker-arms-dw',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'MS during Execute Phase',
                labelTooltip: 'Use Mortal Strike during Execute Phase.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().armsDw?.useMsDuringExecute || false,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.armsDw) {
                        newRotation.armsDw = ArmsDWRotation.clone(Presets.DefaultArmsDWRotation.armsDw);
                    }
                    newRotation.armsDw.useMsDuringExecute = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.ArmsDW,
            },
        },
        // ********************************************************
        //                      GENERAL INPUTS
        // ********************************************************
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
            type: 'number',
            cssClass: 'hs-rage-threshold',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'HS rage threshold',
                labelTooltip: 'Queue HS when rage is above:',
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
            type: 'boolean',
            cssClass: 'overpower-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Overpower',
                labelTooltip: 'Use Overpower when it is possible.',
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
            type: 'boolean',
            cssClass: 'hamstring-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Hamstring',
                labelTooltip: 'Use Hamstring on free global.',
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
            type: 'number',
            cssClass: 'hamstring-rage-threshold',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Hamstring rage threshold',
                labelTooltip: 'Use Hamstring when rage is above a ',
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
    ],
};
