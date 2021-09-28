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
