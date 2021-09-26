import { Sim } from '../sim.js';

export const ShadowPriestDPS = {
  type: 'number' as const,
  cssClass: 'shadow-priest-dps-picker',
  config: {
    label: 'Shadow Priest DPS',
    changedEvent: (sim: Sim<any>) => sim.buffsChangeEmitter,
    getValue: (sim: Sim<any>) => sim.getBuffs().shadowPriestDps,
    setValue: (sim: Sim<any>, newValue: number) => {
      const buffs = sim.getBuffs();
      buffs.shadowPriestDps = newValue;
      sim.setBuffs(buffs);
    },
  },
};
