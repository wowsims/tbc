import { BalanceDruid_Rotation_PrimarySpell as PrimarySpell } from '/tbc/core/proto/druid.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
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
export const BalanceDruidRotationConfig = {
    inputs: [
        {
            type: 'enum',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'primary-spell-enum-picker',
                ],
                label: 'Primary Spell',
                labelTooltip: 'If set to \'Adaptive\', will dynamically adjust rotation based on available mana.',
                values: [
                    {
                        name: 'Adaptive', value: PrimarySpell.Adaptive,
                    },
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
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.primarySpell = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'moonfire-picker',
                ],
                label: 'Use Moonfire',
                labelTooltip: 'Use Moonfire as the next cast after the dot expires.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().moonfire,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.moonfire = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => player.getRotation().primarySpell != PrimarySpell.Adaptive,
            },
        },
        {
            type: 'boolean',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'faerie-fire-picker',
                ],
                label: 'Use Faerie Fire',
                labelTooltip: 'Keep Faerie Fire active on the primary target.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().faerieFire,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.faerieFire = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'insect-swarm-picker',
                ],
                label: 'Use Insect Swarm',
                labelTooltip: 'Keep Insect Swarm active on the primary target.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().insectSwarm,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.insectSwarm = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => player.getTalents().insectSwarm,
            },
        },
        {
            type: 'boolean',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'hurricane-picker',
                ],
                label: 'Use Hurricane',
                labelTooltip: 'Casts Hurricane on cooldown.',
                changedEvent: (player) => player.specOptionsChangeEmitter,
                getValue: (player) => player.getRotation().hurricane,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.hurricane = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'battle-res-picker',
                ],
                label: 'Use Battle Res',
                labelTooltip: 'Cast Battle Res on an ally sometime during the encounter.',
                changedEvent: (player) => player.specOptionsChangeEmitter,
                getValue: (player) => player.getSpecOptions().battleRes,
                setValue: (eventID, player, newValue) => {
                    const newOptions = player.getSpecOptions();
                    newOptions.battleRes = newValue;
                    player.setSpecOptions(eventID, newOptions);
                },
            },
        },
    ],
};
