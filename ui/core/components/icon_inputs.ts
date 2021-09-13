import { Sim } from '../sim'

import { IconInput } from './icon_picker'

// Keep these in alphabetical order.

export const ArcaneBrilliance = {
  spellId: 27127,
  states: 2,
  changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
  getValue: (sim: Sim) => sim.buffs.arcaneInt,
  setBooleanValue: (sim: Sim, newValue: boolean) => {
    const newBuffs = sim.buffs;
    newBuffs.arcaneInt = newValue;
    sim.buffs = newBuffs;
  },
};

export const BlessingOfKings = {
  spellId: 25898,
  states: 2,
  changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
  getValue: (sim: Sim) => sim.buffs.blessingOfKings,
  setBooleanValue: (sim: Sim, newValue: boolean) => {
    const newBuffs = sim.buffs;
    newBuffs.blessingOfKings = newValue;
    sim.buffs = newBuffs;
  },
};

export const BlessingOfWisdom = {
  spellId: 27143,
  states: 3,
  improvedSpellId: 20245,
  changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
  getValue: (sim: Sim) => Number(sim.buffs.improvedBlessingOfWisdom) * 2,
  setBooleanValue: (sim: Sim, newValue: boolean) => {
    const newBuffs = sim.buffs;
    newBuffs.improvedBlessingOfWisdom = newValue;
    sim.buffs = newBuffs;
  },
};

export const Bloodlust = {
  spellId: 2825,
  states: 0,
  changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
  getValue: (sim: Sim) => sim.buffs.bloodlust,
  setNumberValue: (sim: Sim, newValue: number) => {
    const newBuffs = sim.buffs;
    newBuffs.bloodlust = newValue;
    sim.buffs = newBuffs;
  },
};

export const GiftOfTheWild = {
  spellId: 26991,
  states: 2,
  changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
  getValue: (sim: Sim) => sim.buffs.giftOfTheWild,
  setBooleanValue: (sim: Sim, newValue: boolean) => {
    const newBuffs = sim.buffs;
    newBuffs.giftOfTheWild = newValue;
    sim.buffs = newBuffs;
  },
};
