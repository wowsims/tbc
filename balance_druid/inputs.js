import { BalanceDruid_Rotation_PrimarySpell as PrimarySpell } from '/tbc/core/proto/druid.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const BalanceDruidRotationConfig = {
    inputs: [
        {
            type: 'enum', cssClass: 'primary-spell-enum-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Primary Spell',
                values: [
                    {
                        name: 'Starfire', value: PrimarySpell.Starfire,
                    },
                    {
                        name: 'Starfire R6', value: PrimarySpell.Starfire6,
                    },
                    {
                        name: 'Wrath', value: PrimarySpell.Wrath,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().primarySpell,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.primarySpell = newValue;
                    player.setRotation(newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'moonfire-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Moonfire',
                labelTooltip: 'Use Moonfire as the next cast after the dot expires.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().moonfire,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.moonfire = newValue;
                    player.setRotation(newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'faerie-fire-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Faerie Fire',
                labelTooltip: 'Keep Faerie Fire active on the primary target.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().faerieFire,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.faerieFire = newValue;
                    player.setRotation(newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'insect-swarm-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Insect Swarm',
                labelTooltip: 'Keep Insect Swarm active on the primary target.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().insectSwarm,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.insectSwarm = newValue;
                    player.setRotation(newRotation);
                },
                enableWhen: (player) => player.getTalents().insectSwarm,
            },
        },
    ],
};
