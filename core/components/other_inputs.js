import { Potions } from '/tbc/core/proto/common.js';
export const StartingPotion = {
    type: 'enum',
    cssClass: 'starting-potion-picker',
    config: {
        label: 'Starting Potion',
        labelTooltip: 'If set, this potion will be used instead of the default potion for the first few uses.',
        values: [
            { name: 'None', value: Potions.UnknownPotion },
            { name: 'Destruction', value: Potions.DestructionPotion },
            { name: 'Super Mana', value: Potions.SuperManaPotion },
        ],
        changedEvent: (sim) => sim.consumesChangeEmitter,
        getValue: (sim) => sim.getConsumes().startingPotion,
        setValue: (sim, newValue) => {
            const newConsumes = sim.getConsumes();
            newConsumes.startingPotion = newValue;
            sim.setConsumes(newConsumes);
        },
    },
};
export const NumStartingPotions = {
    type: 'number',
    cssClass: 'num-starting-potions-picker',
    config: {
        label: '# to use',
        labelTooltip: 'The number of starting potions to use before going back to the default potion.',
        changedEvent: (sim) => sim.consumesChangeEmitter,
        getValue: (sim) => sim.getConsumes().numStartingPotions,
        setValue: (sim, newValue) => {
            const newConsumes = sim.getConsumes();
            newConsumes.numStartingPotions = newValue;
            sim.setConsumes(newConsumes);
        },
        enableWhen: (sim) => sim.getConsumes().startingPotion != Potions.UnknownPotion,
    },
};
export const ShadowPriestDPS = {
    type: 'number',
    cssClass: 'shadow-priest-dps-picker',
    config: {
        label: 'Shadow Priest DPS',
        changedEvent: (sim) => sim.buffsChangeEmitter,
        getValue: (sim) => sim.getBuffs().shadowPriestDps,
        setValue: (sim, newValue) => {
            const buffs = sim.getBuffs();
            buffs.shadowPriestDps = newValue;
            sim.setBuffs(buffs);
        },
    },
};
