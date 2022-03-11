import { SmitePriest_Rotation_RotationType as RotationType } from '/tbc/core/proto/priest.js';
import { Race, RaidTarget } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const SelfPowerInfusion = {
    id: ActionId.fromSpellId(10060),
    states: 2,
    extraCssClasses: [
        'self-power-infusion-picker',
        'within-raid-sim-hide',
    ],
    changedEvent: (player) => player.specOptionsChangeEmitter,
    getValue: (player) => player.getSpecOptions().powerInfusionTarget?.targetIndex != NO_TARGET,
    setValue: (eventID, player, newValue) => {
        const newOptions = player.getSpecOptions();
        newOptions.powerInfusionTarget = RaidTarget.create({
            targetIndex: newValue ? 0 : NO_TARGET,
        });
        player.setSpecOptions(eventID, newOptions);
    },
};
export const SmitePriestRotationConfig = {
    inputs: [
        {
            type: 'enum', cssClass: 'rotation-enum-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Rotation Type',
                labelTooltip: 'Choose whether to weave optionally weave holy fire for increase Shadow Word: Pain uptime',
                values: [
                    {
                        name: 'Basic', value: RotationType.Basic,
                    },
                    {
                        name: 'HF Weave', value: RotationType.HolyFireWeave,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().rotationType,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.rotationType = newValue;
                    player.setRotation(eventID, newRotation);
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
                setValue: (eventID, player, newValue) => {
                    const newOptions = player.getSpecOptions();
                    newOptions.useShadowfiend = newValue;
                    player.setSpecOptions(eventID, newOptions);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'mindblast-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Mind Blast',
                labelTooltip: 'Use Mind Blast whenever off CD.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useMindBlast,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useMindBlast = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean',
            cssClass: 'swd-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Shadow Word: Death',
                labelTooltip: 'Use Shadow Word: Death whenever off CD.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useShadowWordDeath,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useShadowWordDeath = newValue;
                    player.setRotation(eventID, newRotation);
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
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useDevPlague = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => player.getRace() == Race.RaceUndead,
            },
        },
        {
            type: 'boolean',
            cssClass: 'starshards-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use starshards',
                labelTooltip: 'Use Starshards whenever off CD.',
                changedEvent: (player) => player.raceChangeEmitter,
                getValue: (player) => player.getRotation().useStarshards,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useStarshards = newValue;
                    player.setRotation(eventID, newRotation);
                },
                enableWhen: (player) => player.getRace() == Race.RaceNightElf,
            },
        },
    ],
};
