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
                labelTooltip: 'Choose how to clip your mindflay',
                values: [
                    {
                        name: 'Basic', value: RotationType.Basic,
                    },
                    {
                        name: 'Clip Always', value: RotationType.ClipAlways,
                    },
                    {
                        name: 'Intelligent', value: RotationType.IntelligentClipping,
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
            cssClass: 'devplague-picker',
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
    ],
};
