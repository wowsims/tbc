import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Hunter_Rotation_StingType as StingType, Hunter_Options_Ammo as Ammo, Hunter_Options_QuiverBonus as QuiverBonus, Hunter_Options_PetType as PetType, } from '/tbc/core/proto/hunter.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const Quiver = {
    extraCssClasses: [
        'quiver-picker',
    ],
    numColumns: 1,
    values: [
        { color: '82e89d', value: QuiverBonus.QuiverNone },
        { actionId: ActionId.fromItemId(18714), value: QuiverBonus.Speed15 },
        { actionId: ActionId.fromItemId(2662), value: QuiverBonus.Speed14 },
        { actionId: ActionId.fromItemId(8217), value: QuiverBonus.Speed13 },
        { actionId: ActionId.fromItemId(7371), value: QuiverBonus.Speed12 },
        { actionId: ActionId.fromItemId(3605), value: QuiverBonus.Speed11 },
        { actionId: ActionId.fromItemId(3573), value: QuiverBonus.Speed10 },
    ],
    equals: (a, b) => a == b,
    zeroValue: QuiverBonus.QuiverNone,
    changedEvent: (player) => player.specOptionsChangeEmitter,
    getValue: (player) => player.getSpecOptions().quiverBonus,
    setValue: (eventID, player, newValue) => {
        const newOptions = player.getSpecOptions();
        newOptions.quiverBonus = newValue;
        player.setSpecOptions(eventID, newOptions);
    },
};
export const WeaponAmmo = {
    extraCssClasses: [
        'ammo-picker',
    ],
    numColumns: 1,
    values: [
        { color: 'grey', value: Ammo.AmmoNone },
        { actionId: ActionId.fromItemId(31737), value: Ammo.TimelessArrow },
        { actionId: ActionId.fromItemId(34581), value: Ammo.MysteriousArrow },
        { actionId: ActionId.fromItemId(33803), value: Ammo.AdamantiteStinger },
        { actionId: ActionId.fromItemId(31949), value: Ammo.WardensArrow },
        { actionId: ActionId.fromItemId(30611), value: Ammo.HalaaniRazorshaft },
        { actionId: ActionId.fromItemId(28056), value: Ammo.BlackflightArrow },
    ],
    equals: (a, b) => a == b,
    zeroValue: Ammo.AmmoNone,
    changedEvent: (player) => player.specOptionsChangeEmitter,
    getValue: (player) => player.getSpecOptions().ammo,
    setValue: (eventID, player, newValue) => {
        const newOptions = player.getSpecOptions();
        newOptions.ammo = newValue;
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
        labelTooltip: 'Player latency, used for TODO',
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().latencyMs,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.latencyMs = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
export const PetTypeInput = {
    type: 'enum',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'pet-type-picker',
        ],
        label: 'Pet',
        values: [
            { name: 'None', value: PetType.PetNone },
            { name: 'Ravager', value: PetType.Ravager },
            { name: 'Wind Serpent', value: PetType.WindSerpent },
            //{ name: 'Bat', value: PetType.Bat },
            { name: 'Cat', value: PetType.Cat },
            //{ name: 'Owl', value: PetType.Owl },
            { name: 'Raptor', value: PetType.Raptor },
        ],
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().petType,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.petType = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
export const HunterRotationConfig = {
    inputs: [
        {
            type: 'boolean', cssClass: 'use-multi-shot-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Multi Shot',
                labelTooltip: 'Includes Multi Shot in the rotation.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useMultiShot,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useMultiShot = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean', cssClass: 'use-arcane-shot-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Arcane Shot',
                labelTooltip: 'Includes Arcane Shot in the rotation.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useArcaneShot,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useArcaneShot = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'enum', cssClass: 'sting-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Sting',
                labelTooltip: 'Maintains the selected Sting on the primary target.',
                values: [
                    { name: 'None', value: StingType.NoSting },
                    { name: 'Scorpid Sting', value: StingType.ScorpidSting },
                    { name: 'Serpent Sting', value: StingType.SerpentSting },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().sting,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.sting = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean', cssClass: 'use-french-rotation-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use French Rotation',
                labelTooltip: 'Allows the use of French Rotation (SS+MS or SS+AS between two ranged autos) when no temporary haste effects are active.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useFrenchRotation,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useFrenchRotation = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'number', cssClass: 'viper-start-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Viper Start Mana %',
                labelTooltip: 'Switch to Aspect of the Viper when mana goes below this amount.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().viperStartManaPercent * 100,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.viperStartManaPercent = newValue / 100;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'number', cssClass: 'viper-stop-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Viper Stop Mana %',
                labelTooltip: 'Switch back to Aspect of the Hawk when mana goes above this amount.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().viperStopManaPercent * 100,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.viperStopManaPercent = newValue / 100;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean', cssClass: 'precast-aimed-shot-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Precast Aimed Shot',
                labelTooltip: 'Starts the encounter with an instant Aimed Shot.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().precastAimedShot,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.precastAimedShot = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        },
        {
            type: 'boolean', cssClass: 'melee-weave-picker',
            getModObject: (simUI) => simUI,
            config: {
                label: 'Melee Weave',
                labelTooltip: 'Uses melee weaving in the rotation.',
                changedEvent: (simUI) => simUI.player.rotationChangeEmitter,
                getValue: (simUI) => simUI.player.getRotation().meleeWeave,
                setValue: (eventID, simUI, newValue) => {
                    const newRotation = simUI.player.getRotation();
                    newRotation.meleeWeave = newValue;
                    simUI.player.setRotation(eventID, newRotation);
                    simUI.recomputeSettingsLayout();
                },
            },
        },
        {
            type: 'boolean', cssClass: 'use-raptor-strike-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Use Raptor Strike',
                labelTooltip: 'Uses Raptor Strike instead of regular melee hits while weaving.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().useRaptorStrike,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.useRaptorStrike = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().meleeWeave,
            },
        },
        {
            type: 'number', cssClass: 'time-to-weave-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Time To Weave (ms)',
                labelTooltip: 'Amount of time, in milliseconds, between when you start moving towards the boss and when you re-engage your ranged autos.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().timeToWeaveMs,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.timeToWeaveMs = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().meleeWeave,
            },
        },
        {
            type: 'number', cssClass: 'percent-weaved-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Time Weaved (%)',
                labelTooltip: 'Percentage of fight to use melee weaving.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().percentWeaved * 100,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.percentWeaved = newValue / 100;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().meleeWeave,
            },
        },
    ],
};
