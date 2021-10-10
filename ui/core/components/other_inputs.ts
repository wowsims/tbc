import { Potions } from '/tbc/core/proto/common.js';
import { Sim } from '/tbc/core/sim.js';

export const StartingPotion = {
	type: 'enum' as const,
	cssClass: 'starting-potion-picker',
	config: {
		label: 'Starting Potion',
		labelTooltip: 'If set, this potion will be used instead of the default potion for the first few uses.',
		values: [
			{ name: 'None', value: Potions.UnknownPotion },
			{ name: 'Destruction', value: Potions.DestructionPotion },
			{ name: 'Super Mana', value: Potions.SuperManaPotion },
		],
		changedEvent: (sim: Sim<any>) => sim.consumesChangeEmitter,
		getValue: (sim: Sim<any>) => sim.getConsumes().startingPotion,
		setValue: (sim: Sim<any>, newValue: number) => {
			const newConsumes = sim.getConsumes();
			newConsumes.startingPotion = newValue;
			sim.setConsumes(newConsumes);
		},
	},
};

export const NumStartingPotions = {
	type: 'number' as const,
	cssClass: 'num-starting-potions-picker',
	config: {
		label: '# to use',
		labelTooltip: 'The number of starting potions to use before going back to the default potion.',
		changedEvent: (sim: Sim<any>) => sim.consumesChangeEmitter,
		getValue: (sim: Sim<any>) => sim.getConsumes().numStartingPotions,
		setValue: (sim: Sim<any>, newValue: number) => {
			const newConsumes = sim.getConsumes();
			newConsumes.numStartingPotions = newValue;
			sim.setConsumes(newConsumes);
		},
	},
};

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
