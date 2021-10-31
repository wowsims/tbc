import { Potions } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { SimUI } from '/tbc/core/sim_ui.js';

export const StartingPotion = {
	type: 'enum' as const,
	cssClass: 'starting-potion-picker',
	getModObject: (simUI: SimUI<any>) => simUI.player,
	config: {
		label: 'Starting Potion',
		labelTooltip: 'If set, this potion will be used instead of the default potion for the first few uses.',
		values: [
			{ name: 'None', value: Potions.UnknownPotion },
			{ name: 'Destruction', value: Potions.DestructionPotion },
			{ name: 'Super Mana', value: Potions.SuperManaPotion },
		],
		changedEvent: (player: Player<any>) => player.consumesChangeEmitter,
		getValue: (player: Player<any>) => player.getConsumes().startingPotion,
		setValue: (player: Player<any>, newValue: number) => {
			const newConsumes = player.getConsumes();
			newConsumes.startingPotion = newValue;
			player.setConsumes(newConsumes);
		},
	},
};

export const NumStartingPotions = {
	type: 'number' as const,
	cssClass: 'num-starting-potions-picker',
	getModObject: (simUI: SimUI<any>) => simUI.player,
	config: {
		label: '# to use',
		labelTooltip: 'The number of starting potions to use before going back to the default potion.',
		changedEvent: (player: Player<any>) => player.consumesChangeEmitter,
		getValue: (player: Player<any>) => player.getConsumes().numStartingPotions,
		setValue: (player: Player<any>, newValue: number) => {
			const newConsumes = player.getConsumes();
			newConsumes.numStartingPotions = newValue;
			player.setConsumes(newConsumes);
		},
		enableWhen: (player: Player<any>) => player.getConsumes().startingPotion != Potions.UnknownPotion,
	},
};

export const ShadowPriestDPS = {
  type: 'number' as const,
  cssClass: 'shadow-priest-dps-picker',
	getModObject: (simUI: SimUI<any>) => simUI.sim,
  config: {
    label: 'Shadow Priest DPS',
    changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
    getValue: (sim: Sim) => sim.getBuffs().shadowPriestDps,
    setValue: (sim: Sim, newValue: number) => {
      const buffs = sim.getBuffs();
      buffs.shadowPriestDps = newValue;
      sim.setBuffs(buffs);
    },
  },
};
