import { FeralDruid_Rotation_FinishingMove as FinishingMove } from '/tbc/core/proto/druid.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
// Helper function for identifying whether 2pT6 is equipped, which impacts allowed rotation options
function numThunderheartPieces(player) {
    const gear = player.getGear();
    const itemIds = [31048, 31042, 31034, 31044, 31039];
    return gear.asArray().map(equippedItem => equippedItem?.item.id).filter(id => itemIds.includes(id)).length;
}
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const SelfInnervate = {
    id: ActionId.fromSpellId(29166),
    states: 2,
    extraCssClasses: [
        'self-innervate-picker',
        'within-raid-sim-hide',
    ],
    changedEvent: (player) => player.specOptionsChangeEmitter,
    getValue: (player) => player.getSpecOptions().innervateTarget?.targetIndex != NO_TARGET,
    setValue: (eventID, player, newValue) => {
        const newOptions = player.getSpecOptions();
        newOptions.innervateTarget = RaidTarget.create({
            targetIndex: newValue ? 0 : NO_TARGET,
        });
        player.setSpecOptions(eventID, newOptions);
    },
};
export const LatencyMs = {
    type: 'number',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'latency-ms-picker',
        ],
        label: 'Latency',
        labelTooltip: 'Player latency, in milliseconds. Adds a delay to actions that cannot be spell queued.',
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().latencyMs,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.latencyMs = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
export const FeralDruidRotationConfig = {
    inputs: [
        {
            type: 'enum',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'finishing-move-enum-picker',
                ],
                label: 'Finishing Move',
                labelTooltip: 'Specify whether Rip or Ferocious Bite should be used as the primary finisher in the DPS rotation.',
                values: [
                    {
                        name: 'Rip', value: FinishingMove.Rip,
                    },
                    {
                        name: 'Ferocious Bite', value: FinishingMove.Bite,
                    },
                    {
                        name: 'None', value: FinishingMove.None,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().finishingMove,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.finishingMove = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'biteweave-picker',
                ],
                label: 'Enable Bite-weaving',
                labelTooltip: 'Spend Combo Points on Ferocious Bite when Rip is already applied on the target.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().biteweave,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.biteweave = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => player.getRotation().finishingMove == FinishingMove.Rip,
            },
        },
        {
            type: 'boolean',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'ripweave-picker',
                ],
                label: 'Enable Rip-weaving',
                labelTooltip: 'Spend Combo Points on Rip when at 52 Energy or above.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().ripweave,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.ripweave = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => player.getRotation().finishingMove == FinishingMove.Bite,
            },
        },
        {
            type: 'enum',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'rip-cp-enum-picker',
                ],
                label: 'Rip CP Threshold',
                labelTooltip: 'Minimum Combo Points to accumulate before casting Rip as a finisher.',
                values: [
                    {
                        name: '4', value: 4,
                    },
                    {
                        name: '5', value: 5,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().ripMinComboPoints,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.ripMinComboPoints = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => (player.getRotation().finishingMove == FinishingMove.Rip) || (player.getRotation().ripweave && (player.getRotation().finishingMove != FinishingMove.None)),
            },
        },
        {
            type: 'enum',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'bite-cp-enum-picker',
                ],
                label: 'Bite CP Threshold',
                labelTooltip: 'Minimum Combo Points to accumulate before casting Ferocious Bite as a finisher.',
                values: [
                    {
                        name: '4', value: 4,
                    },
                    {
                        name: '5', value: 5,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().biteMinComboPoints,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.biteMinComboPoints = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => (player.getRotation().finishingMove == FinishingMove.Bite) || (player.getRotation().biteweave && (player.getRotation().finishingMove != FinishingMove.None)),
            },
        },
        {
            type: 'boolean',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'mangle-trick-picker',
                ],
                label: 'Use Mangle trick',
                labelTooltip: 'Cast Mangle rather than Shred when between 50-56 Energy with 2pT6, or 60-61 Energy without 2pT6, and with less than 1 second remaining until the next Energy tick.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().mangleTrick,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.mangleTrick = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'rake-trick-picker',
                ],
                label: 'Use Rake/Bite tricks',
                labelTooltip: 'Cast Rake or Ferocious Bite rather than powershifting when between 35-39 Energy without 2pT6, and with more than 1 second remaining until the next Energy tick.',
                changedEvent: (player) => player.changeEmitter,
                getValue: (player) => player.getRotation().rakeTrick,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.rakeTrick = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => numThunderheartPieces(player) < 2,
            },
        },
    ],
};
