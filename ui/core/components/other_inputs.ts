import { EnumPicker, EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { Target } from '/tbc/core/target.js';
import { SimUI } from '/tbc/core/sim_ui.js';

export function makePhaseSelector(parent: HTMLElement, sim: Sim): EnumPicker<Sim> {
	return new EnumPicker<Sim>(parent, sim, {
		values: [
			{ name: 'Phase 1', value: 1 },
			{ name: 'Phase 2', value: 2 },
			{ name: 'Phase 3', value: 3 },
			{ name: 'Phase 4', value: 4 },
			{ name: 'Phase 5', value: 5 },
		],
		changedEvent: (sim: Sim) => sim.phaseChangeEmitter,
		getValue: (sim: Sim) => sim.getPhase(),
		setValue: (sim: Sim, newValue: number) => {
			sim.setPhase(newValue);
		},
	});
}

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
	getModObject: (simUI: SimUI<any>) => simUI.player,
  config: {
    label: 'Shadow Priest DPS',
    changedEvent: (player: Player<any>) => player.buffsChangeEmitter,
    getValue: (player: Player<any>) => player.getBuffs().shadowPriestDps,
    setValue: (player: Player<any>, newValue: number) => {
      const buffs = player.getBuffs();
      buffs.shadowPriestDps = newValue;
      player.setBuffs(buffs);
    },
  },
};

export const ISBUptime = {
    type: 'number' as const,
    cssClass: 'isb-uptime-picker',
    getModObject: (simUI: SimUI<any>) => simUI.encounter.primaryTarget,
    config: {
      label: 'ISB Uptime %',
      changedEvent: (target: Target) => target.debuffsChangeEmitter,
      getValue: (target: Target) => Math.round(target.getDebuffs().isbUptime*100),
      setValue: (target: Target, newValue: number) => {
        const newDebuffs = target.getDebuffs();
        newDebuffs.isbUptime = newValue/100;
        target.setDebuffs(newDebuffs);
      },
    },
  };