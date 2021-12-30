import { EnumPicker, EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { Target } from '/tbc/core/target.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

export function makePhaseSelector(parent: HTMLElement, sim: Sim): EnumPicker<Sim> {
	return new EnumPicker<Sim>(parent, sim, {
		extraCssClasses: [
			'phase-selector',
		],
		values: [
			{ name: 'Phase 1', value: 1 },
			{ name: 'Phase 2', value: 2 },
			{ name: 'Phase 3', value: 3 },
			{ name: 'Phase 4', value: 4 },
			{ name: 'Phase 5', value: 5 },
		],
		changedEvent: (sim: Sim) => sim.phaseChangeEmitter,
		getValue: (sim: Sim) => sim.getPhase(),
		setValue: (eventID: EventID, sim: Sim, newValue: number) => {
			sim.setPhase(eventID, newValue);
		},
	});
}

export const StartingPotion = {
	type: 'enum' as const,
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		extraCssClasses: [
			'starting-potion-picker',
		],
		label: 'Starting Potion',
		labelTooltip: 'If set, this potion will be used instead of the default potion for the first few uses.',
		values: [
			{ name: 'None', value: Potions.UnknownPotion },
			{ name: 'Destruction', value: Potions.DestructionPotion },
			{ name: 'Haste', value: Potions.HastePotion },
			{ name: 'Super Mana', value: Potions.SuperManaPotion },
		],
		changedEvent: (player: Player<any>) => player.consumesChangeEmitter,
		getValue: (player: Player<any>) => player.getConsumes().startingPotion,
		setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
			const newConsumes = player.getConsumes();
			newConsumes.startingPotion = newValue;
			player.setConsumes(eventID, newConsumes);
		},
	},
};

export const NumStartingPotions = {
	type: 'number' as const,
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
	config: {
		extraCssClasses: [
			'num-starting-potions-picker',
		],
		label: '# to use',
		labelTooltip: 'The number of starting potions to use before going back to the default potion.',
		changedEvent: (player: Player<any>) => player.consumesChangeEmitter,
		getValue: (player: Player<any>) => player.getConsumes().numStartingPotions,
		setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
			const newConsumes = player.getConsumes();
			newConsumes.numStartingPotions = newValue;
			player.setConsumes(eventID, newConsumes);
		},
		enableWhen: (player: Player<any>) => player.getConsumes().startingPotion != Potions.UnknownPotion,
	},
};

export const ShadowPriestDPS = {
  type: 'number' as const,
  cssClass: 'shadow-priest-dps-picker',
	getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
  config: {
		extraCssClasses: [
			'shadow-priest-dps-picker',
			'within-raid-sim-hide',
		],
    label: 'Shadow Priest DPS',
    changedEvent: (player: Player<any>) => player.buffsChangeEmitter,
    getValue: (player: Player<any>) => player.getBuffs().shadowPriestDps,
    setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
      const buffs = player.getBuffs();
      buffs.shadowPriestDps = newValue;
      player.setBuffs(eventID, buffs);
    },
  },
};

export const ISBUptime = {
	type: 'number' as const,
	getModObject: (simUI: IndividualSimUI<any>) => simUI.sim.encounter.primaryTarget,
	config: {
		extraCssClasses: [
			'isb-uptime-picker',
			'within-raid-sim-hide',
		],
		label: 'ISB Uptime %',
		changedEvent: (target: Target) => target.debuffsChangeEmitter,
		getValue: (target: Target) => Math.round(target.getDebuffs().isbUptime*100),
		setValue: (eventID: EventID, target: Target, newValue: number) => {
			const newDebuffs = target.getDebuffs();
			newDebuffs.isbUptime = newValue/100;
			target.setDebuffs(eventID, newDebuffs);
		},
	},
};
