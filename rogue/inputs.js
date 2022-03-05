import { Rogue_Rotation_Builder as Builder, } from '/tbc/core/proto/rogue.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const RogueRotationConfig = {
    inputs: [
        {
            type: 'enum', cssClass: 'builder-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Builder',
                values: [
                    {
                        name: 'Auto', value: Builder.Auto,
                        tooltip: 'Automatically selects a builder based on weapons/talents.',
                    },
                    { name: 'Sinister Strike', value: Builder.SinisterStrike },
                    { name: 'Backstab', value: Builder.Backstab },
                    { name: 'Hemorrhage', value: Builder.Hemorrhage },
                    { name: 'Mutilate', value: Builder.Mutilate },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().builder,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.builder = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean', cssClass: 'maintain-expose-armor-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Maintain EA',
                labelTooltip: 'Keeps Expose Armor active on the primary target.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().maintainExposeArmor,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.maintainExposeArmor = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean', cssClass: 'use-rupture-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Rupture',
                labelTooltip: 'Uses Rupture over Eviscerate when appropriate.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useRupture,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useRupture = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'number', cssClass: 'min-combo-points-for-dps-finisher-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Min CPs for Damage Finisher',
                labelTooltip: 'Will not use Eviscerate or Rupture unless the Rogue has at least this many Combo Points.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().minComboPointsForDamageFinisher,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.minComboPointsForDamageFinisher = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
    ],
};
function makeBooleanRogueBuffInput(id, optionsFieldName) {
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
