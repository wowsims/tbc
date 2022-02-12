// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const RetributionPaladinRotationConfig = {
    inputs: [
        {
            type: 'boolean', cssClass: 'consecration-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Consecration',
                labelTooltip: 'Use consecration whenever the target does not already have the DoT.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().consecration,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.consecration = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        }
    ],
};
