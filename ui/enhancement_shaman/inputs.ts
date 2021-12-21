import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { EarthTotem, EnhancementShaman_Rotation_RotationType as RotationType, ShamanTotems } from '/tbc/core/proto/shaman.js';
import { EnhancementShaman_Options as ShamanOptions } from '/tbc/core/proto/shaman.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ItemOrSpellId } from '/tbc/core/resources.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Target } from '/tbc/core/target.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const IconBloodlust = makeBooleanShamanBuffInput({ spellId: 2825 }, 'bloodlust');
// export const IconManaSpringTotem = makeBoolShamanTotem({ spellId: 25570 }, 'manaSpringTotem');
// export const IconTotemOfWrath = makeBoolShamanTotem({ spellId: 30706 }, 'totemOfWrath');
export const IconWaterShield = makeBooleanShamanBuffInput({ spellId: 33736 }, 'waterShield');
// export const IconWrathOfAirTotem = makeBoolShamanTotem({ spellId: 3738 }, 'wrathOfAirTotem');

export const EnhancementShamanRotationConfig = {
	inputs: [
		{
			type: 'enum' as const, cssClass: 'rotation-enum-picker',
			getModObject: (simUI: IndividualSimUI<any>) => simUI.player,
			config: {
				label: 'Type',
				values: [
					{
						name: 'Basic', value: RotationType.Basic,
						tooltip: 'does basic stuff',
					},
				],
				changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.rotationChangeEmitter,
				getValue: (player: Player<Spec.SpecEnhancementShaman>) => player.getRotation().type,
				setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: number) => {
					const newRotation = player.getRotation();
					newRotation.type = newValue;
					player.setRotation(eventID, newRotation);
				},
			},
		}
	],
};

function makeBooleanShamanBuffInput(id: ItemOrSpellId, optionsFieldName: keyof ShamanOptions): IconPickerConfig<Player<any>, boolean> {
	return {
	  id: id,
	  states: 2,
		  changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.specOptionsChangeEmitter,
		  getValue: (player: Player<Spec.SpecEnhancementShaman>) => player.getSpecOptions()[optionsFieldName] as boolean,
		  setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: boolean) => {
			  const newOptions = player.getSpecOptions();
		(newOptions[optionsFieldName] as boolean) = newValue;
			  player.setSpecOptions(eventID, newOptions);
		  },
	}
}

// function makeBoolShamanTotem(id: ItemOrSpellId, optionsFieldName: keyof totems?): IconPickerConfig<Player<any>, boolean> {
//   return {
//     id: id,
//     states: 2,
// 		changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.specOptionsChangeEmitter,
// 		getValue: (player: Player<Spec.SpecEnhancementShaman>) => {
// 			const totems = player.getSpecOptions().totems as ShamanTotems;
// 			return totems[optionsFieldName] as boolean;
// 		},
// 		setValue: (player: Player<Spec.SpecEnhancementShaman>, newValue: boolean) => {
// 			const newOptions = player.getSpecOptions();
// 			const totems = newOptions.totems as ShamanTotems;
//       		(totems[optionsFieldName] as boolean) = newValue;
// 			newOptions.totems = totems;
// 			player.setSpecOptions(newOptions);
// 		},
//   }
// }
