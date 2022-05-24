import { PaladinAura as PaladinAura, PaladinJudgement as PaladinJudgement, } from '/tbc/core/proto/paladin.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const ProtectionPaladinRotationConfig = {
    inputs: [
        {
            type: 'enum', cssClass: 'consecration-rank-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Consecration Rank',
                labelTooltip: 'Use specified rank of Consecration during filler spell windows.',
                values: [
                    {
                        name: 'None', value: 0,
                    },
                    {
                        name: 'Rank 1', value: 1,
                    },
                    {
                        name: 'Rank 4', value: 4,
                    },
                    {
                        name: 'Rank 6', value: 6,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().consecrationRank,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.consecrationRank = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean', cssClass: 'prioritize-holy-shield-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Prio Holy Shield',
                labelTooltip: 'Uses Holy Shield as the highest priority spell. This is usually done when tanking a boss that can crush.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().prioritizeHolyShield,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.prioritizeHolyShield = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean', cssClass: 'exorcism-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Exorcism',
                labelTooltip: 'Includes Exorcism in the rotation. Will only be used if the primary target is an Undead or Demon type.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useExorcism,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useExorcism = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'enum', cssClass: 'mantain-judgement-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Maintain Judgement',
                values: [
                    { name: 'None', value: PaladinJudgement.NoPaladinJudgement },
                    { name: 'Wisdom', value: PaladinJudgement.JudgementOfWisdom },
                    { name: 'Light', value: PaladinJudgement.JudgementOfLight },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().maintainJudgement,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.maintainJudgement = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
    ],
};
export const AuraSelection = {
    type: 'enum', cssClass: 'aura-picker',
    getModObject: (simUI) => simUI.player,
    config: {
        label: 'Aura',
        values: [
            { name: 'None', value: PaladinAura.NoPaladinAura },
            { name: 'Sanctity Aura', value: PaladinAura.SanctityAura },
            { name: 'Devotion Aura', value: PaladinAura.DevotionAura },
            { name: 'Retribution Aura', value: PaladinAura.RetributionAura },
        ],
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().aura,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.aura = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
export const UseAvengingWrath = {
    type: 'boolean', cssClass: 'use-avenging-wrath-picker',
    getModObject: (simUI) => simUI.player,
    config: {
        label: 'Use Avenging Wrath',
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().useAvengingWrath,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.useAvengingWrath = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
