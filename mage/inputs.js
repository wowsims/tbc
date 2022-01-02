import { Mage_Rotation_Type as RotationType, Mage_Rotation_ArcaneRotation as ArcaneRotation, Mage_Rotation_FireRotation as FireRotation, Mage_Rotation_FrostRotation as FrostRotation } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_FireRotation_PrimarySpell as PrimaryFireSpell } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_ArcaneRotation_Filler as ArcaneFiller } from '/tbc/core/proto/mage.js';
import { Mage_Options_ArmorType as ArmorType } from '/tbc/core/proto/mage.js';
import * as Presets from './presets.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const ManaEmerald = makeBooleanMageBuffInput({ itemId: 22044 }, 'useManaEmeralds');
export const MageArmor = {
    id: { spellId: 27125 },
    states: 2,
    extraCssClasses: [
        'mage-armor-picker',
    ],
    changedEvent: (player) => player.specOptionsChangeEmitter,
    getValue: (player) => player.getSpecOptions().armor == ArmorType.MageArmor,
    setValue: (eventID, player, newValue) => {
        const newOptions = player.getSpecOptions();
        newOptions.armor = newValue ? ArmorType.MageArmor : ArmorType.NoArmor;
        player.setSpecOptions(eventID, newOptions);
    },
};
export const MoltenArmor = {
    id: { spellId: 30482 },
    states: 2,
    extraCssClasses: [
        'molten-armor-picker',
    ],
    changedEvent: (player) => player.specOptionsChangeEmitter,
    getValue: (player) => player.getSpecOptions().armor == ArmorType.MoltenArmor,
    setValue: (eventID, player, newValue) => {
        const newOptions = player.getSpecOptions();
        newOptions.armor = newValue ? ArmorType.MoltenArmor : ArmorType.NoArmor;
        player.setSpecOptions(eventID, newOptions);
    },
};
export const EvocationTicks = {
    type: 'number',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'evocation-ticks-picker',
        ],
        label: '# Evocation Ticks',
        labelTooltip: 'The number of ticks of Evocation to use, or 0 to use the full duration.',
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions().evocationTicks,
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions.evocationTicks = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    },
};
export const MageRotationConfig = {
    inputs: [
        {
            type: 'enum',
            getModObject: (simUI) => simUI,
            config: {
                extraCssClasses: [
                    'rotation-type-enum-picker',
                ],
                label: 'Spec',
                labelTooltip: 'Switches between spec rotation settings. Will also update talents to defaults for the selected spec.',
                values: [
                    {
                        name: 'Arcane', value: RotationType.Arcane,
                    },
                    {
                        name: 'Fire', value: RotationType.Fire,
                    },
                    {
                        name: 'Frost', value: RotationType.Frost,
                    },
                ],
                changedEvent: (simUI) => simUI.player.rotationChangeEmitter,
                getValue: (simUI) => simUI.player.getRotation().type,
                setValue: (eventID, simUI, newValue) => {
                    const newRotation = simUI.player.getRotation();
                    newRotation.type = newValue;
                    if (newRotation.type == RotationType.Arcane) {
                        simUI.player.setTalentsString(eventID, Presets.ArcaneTalents.data);
                        if (!newRotation.arcane) {
                            newRotation.arcane = ArcaneRotation.create();
                        }
                    }
                    else if (newRotation.type == RotationType.Fire) {
                        simUI.player.setTalentsString(eventID, Presets.FireTalents.data);
                        if (!newRotation.fire) {
                            newRotation.fire = FireRotation.create();
                        }
                    }
                    else {
                        simUI.player.setTalentsString(eventID, Presets.FrostTalents.data);
                        if (!newRotation.frost) {
                            newRotation.frost = FrostRotation.create();
                        }
                    }
                    simUI.player.setRotation(eventID, newRotation);
                    simUI.recomputeSettingsLayout();
                },
            },
        },
        // ********************************************************
        //                       FIRE INPUTS
        // ********************************************************
        {
            type: 'enum',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'rotation-type-enum-picker',
                ],
                label: 'Primary Spell',
                values: [
                    {
                        name: 'Fireball', value: PrimaryFireSpell.Fireball,
                    },
                    {
                        name: 'Scorch', value: PrimaryFireSpell.Scorch,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().fire?.primarySpell || PrimaryFireSpell.Fireball,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.fire) {
                        newRotation.fire = FireRotation.create();
                    }
                    newRotation.fire.primarySpell = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Fire,
            },
        },
        {
            type: 'boolean',
            cssClass: 'maintain-improved-scorch-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Maintain Imp. Scorch',
                labelTooltip: 'Always use Scorch when below 5 stacks, or < 5.5s remaining on debuff.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().fire?.maintainImprovedScorch || false,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.fire) {
                        newRotation.fire = FireRotation.create();
                    }
                    newRotation.fire.maintainImprovedScorch = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Fire,
            },
        },
        {
            type: 'boolean',
            cssClass: 'weave-fire-blast-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Weave Fire Blast',
                labelTooltip: 'Use Fire Blast whenever its off CD.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().fire?.weaveFireBlast || false,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.fire) {
                        newRotation.fire = FireRotation.create();
                    }
                    newRotation.fire.weaveFireBlast = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Fire,
            },
        },
        // ********************************************************
        //                      ARCANE INPUTS
        // ********************************************************
        {
            type: 'enum',
            getModObject: (simUI) => simUI.player,
            config: {
                extraCssClasses: [
                    'filler-enum-picker',
                ],
                label: 'Filler',
                labelTooltip: 'Spells to cast while waiting for Arcane Blast stacks to drop.',
                values: [
                    {
                        name: 'Frostbolt', value: ArcaneFiller.Frostbolt,
                    },
                    {
                        name: 'Arcane Missiles', value: ArcaneFiller.ArcaneMissiles,
                    },
                    {
                        name: 'Scorch', value: ArcaneFiller.Fireball,
                    },
                    {
                        name: 'Fireball', value: ArcaneFiller.Fireball,
                    },
                    {
                        name: 'AM + FrB', value: ArcaneFiller.ArcaneMissilesFrostbolt,
                    },
                    {
                        name: 'AM + Scorch', value: ArcaneFiller.ArcaneMissilesScorch,
                    },
                    {
                        name: 'Scorch + 2xFiB', value: ArcaneFiller.ScorchTwoFireball,
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().arcane?.filler || ArcaneFiller.Frostbolt,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.arcane) {
                        newRotation.arcane = ArcaneRotation.create();
                    }
                    newRotation.arcane.filler = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Arcane,
            },
        },
        {
            type: 'number',
            cssClass: 'arcane-blasts-between-fillers-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: '# ABs between Fillers',
                labelTooltip: 'Number of Arcane Blasts to cast once the stacks drop.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().arcane?.arcaneBlastsBetweenFillers || 0,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.arcane) {
                        newRotation.arcane = ArcaneRotation.create();
                    }
                    newRotation.arcane.arcaneBlastsBetweenFillers = newValue;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Arcane,
            },
        },
        {
            type: 'number',
            cssClass: 'start-regen-rotation-percent-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Start regen rotation at mana %',
                labelTooltip: 'Percent of mana pool, below which the regen rotation should be used (alternate fillers and a few ABs).',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => (player.getRotation().arcane?.startRegenRotationPercent || 0) * 100,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.arcane) {
                        newRotation.arcane = ArcaneRotation.create();
                    }
                    newRotation.arcane.startRegenRotationPercent = newValue / 100;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Arcane,
            },
        },
        {
            type: 'number',
            cssClass: 'stop-regen-rotation-percent-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Stop regen rotation at mana %',
                labelTooltip: 'Percent of mana pool, above which will go back to AB spam.',
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => (player.getRotation().arcane?.stopRegenRotationPercent || 0) * 100,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    if (!newRotation.arcane) {
                        newRotation.arcane = ArcaneRotation.create();
                    }
                    newRotation.arcane.stopRegenRotationPercent = newValue / 100;
                    player.setRotation(eventID, newRotation);
                },
                showWhen: (player) => player.getRotation().type == RotationType.Arcane,
            },
        },
    ],
};
function makeBooleanMageBuffInput(id, optionsFieldName) {
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
