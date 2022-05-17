import { PaladinAura as PaladinAura, ProtectionPaladin_Options_PrimaryJudgement as PrimaryJudgement, } from '/tbc/core/proto/paladin.js';
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
            type: 'boolean', cssClass: 'exorcism-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Exorcism',
                labelTooltip: 'Use Exorcism during filler spell windows. Will only be used if the boss mob type is Undead or Demon.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useExorcism,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useExorcism = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        }
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
export const PrimaryJudgementSelection = {
    type: 'enum', cssClass: 'judgement-picker',
    getModObject: (simUI) => simUI.player,
    config: {
        label: 'Primary Judgement',
        values: [
            {
                name: 'Vengeance', value: PrimaryJudgement.Vengeance,
            },
            {
                name: 'Righteousness', value: PrimaryJudgement.Righteousness,
            },
            {
                name: 'Twist', value: PrimaryJudgement.Twist,
            },
        ],
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().primaryJudgement,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.primaryJudgement = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
