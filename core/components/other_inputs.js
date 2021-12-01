import { Potions } from '/tbc/core/proto/common.js';
export const StartingPotion = {
    type: 'enum',
    cssClass: 'starting-potion-picker',
    getModObject: (simUI) => simUI.player,
    config: {
        label: 'Starting Potion',
        labelTooltip: 'If set, this potion will be used instead of the default potion for the first few uses.',
        values: [
            { name: 'None', value: Potions.UnknownPotion },
            { name: 'Destruction', value: Potions.DestructionPotion },
            { name: 'Super Mana', value: Potions.SuperManaPotion },
        ],
        changedEvent: (player) => player.consumesChangeEmitter,
        getValue: (player) => player.getConsumes().startingPotion,
        setValue: (player, newValue) => {
            const newConsumes = player.getConsumes();
            newConsumes.startingPotion = newValue;
            player.setConsumes(newConsumes);
        },
    },
};
export const NumStartingPotions = {
    type: 'number',
    cssClass: 'num-starting-potions-picker',
    getModObject: (simUI) => simUI.player,
    config: {
        label: '# to use',
        labelTooltip: 'The number of starting potions to use before going back to the default potion.',
        changedEvent: (player) => player.consumesChangeEmitter,
        getValue: (player) => player.getConsumes().numStartingPotions,
        setValue: (player, newValue) => {
            const newConsumes = player.getConsumes();
            newConsumes.numStartingPotions = newValue;
            player.setConsumes(newConsumes);
        },
        enableWhen: (player) => player.getConsumes().startingPotion != Potions.UnknownPotion,
    },
};
export const ShadowPriestDPS = {
    type: 'number',
    cssClass: 'shadow-priest-dps-picker',
    getModObject: (simUI) => simUI.sim,
    config: {
        label: 'Shadow Priest DPS',
        changedEvent: (sim) => sim.individualBuffsChangeEmitter,
        getValue: (sim) => sim.getIndividualBuffs().shadowPriestDps,
        setValue: (sim, newValue) => {
            const individualBuffs = sim.getIndividualBuffs();
            individualBuffs.shadowPriestDps = newValue;
            sim.setIndividualBuffs(individualBuffs);
        },
    },
};
