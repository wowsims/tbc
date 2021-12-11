import { ShadowPriest_Rotation_RotationType as RotationType } from '/tbc/core/proto/priest.js';
import { Race } from '/tbc/core/proto/common.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const ShadowPriestRotationConfig = {
    inputs: [
        {
            type: 'enum', cssClass: 'rotation-enum-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Rotation Type',
                labelTooltip: 'Choose how to clip your mindflay. Basic will never clip. Clipping will clip for other spells and use a 2xMF2 when there is time for 4 ticks. Ideal will evaluate the DPS gain of every action to determine MF actions.',
                values: [
                    {
                        name: 'Basic', value: RotationType.Basic,
                    },
                    {
                        name: 'Clipping', value: RotationType.Clipping,
                    },
                    {
                        name: 'Ideal', value: RotationType.Ideal,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().rotationType,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.rotationType = newValue;
                    player.setRotation(newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'shadowfiend-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Shadowfiend',
                labelTooltip: 'Use Shadowfiend when low mana and off CD.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getSpecOptions().useShadowfiend,
                setValue: (player, newValue) => {
                    const newOptions = player.getSpecOptions();
                    newOptions.useShadowfiend = newValue;
                    player.setSpecOptions(newOptions);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'precastvt-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Precast Vampiric Touch',
                labelTooltip: 'Start fight with VT landing at time 0',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().precastVt,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.precastVt = newValue;
                    player.setRotation(newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'devplague-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Devouring Plague',
                labelTooltip: 'Use Devouring Plague whenever off CD.',
                changedEvent: (player) => player.raceChangeEmitter,
                getValue: (player) => player.getRotation().useDevPlague,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useDevPlague = newValue;
                    player.setRotation(newRotation);
                },
                enableWhen: (player) => player.getRace() == Race.RaceUndead,
            },
        },
        {
            type: 'number',
            cssClass: 'latency-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Channeling Latency (ms)',
                labelTooltip: 'Latency after a channel that lasts longer than GCD. 0 to disable. Has a minimum value of 100ms if set.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().latency,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.latency = newValue;
                    player.setRotation(newRotation);
                },
            },
        },
    ],
};
