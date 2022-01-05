import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { IconEnumPicker } from '/tbc/core/components/icon_enum_picker.js';
import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { AirTotem, EarthTotem, FireTotem, WaterTotem, EnhancementShaman_Rotation_RotationType as RotationType, ShamanTotems } from '/tbc/core/proto/shaman.js';
import { EnhancementShaman_Options as ShamanOptions } from '/tbc/core/proto/shaman.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ItemOrSpellId } from '/tbc/core/proto_utils/action_id.js';
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

export function TotemsSection(simUI: IndividualSimUI<Spec.SpecEnhancementShaman>, parentElem: HTMLElement): string {
	const customSectionsContainer = parentElem.closest('.custom-sections-container') as HTMLElement;
	customSectionsContainer.style.zIndex = '1000';

	parentElem.innerHTML = `
		<div class="totem-dropdowns-container"></div>
		<div class="totem-inputs-container"></div>
	`;
	const totemDropdownsContainer = parentElem.getElementsByClassName('totem-dropdowns-container')[0] as HTMLElement;
	const totemInputsContainer = parentElem.getElementsByClassName('totem-inputs-container')[0] as HTMLElement;

	const earthTotemPicker = new IconEnumPicker(totemDropdownsContainer, simUI.player, {
		extraCssClasses: [
			'earth-totem-picker',
		],
		numColumns: 1,
		values: [
			{ color: '#ffdfba', value: EarthTotem.NoEarthTotem },
			{ actionId: { id: { spellId: 25528 }}, value: EarthTotem.StrengthOfEarthTotem },
			{ actionId: { id: { spellId: 8143 }}, value: EarthTotem.TremorTotem },
		],
		equals: (a: EarthTotem, b: EarthTotem) => a == b,
		zeroValue: EarthTotem.NoEarthTotem,
		changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.rotationChangeEmitter,
		getValue: (player: Player<Spec.SpecEnhancementShaman>) => player.getRotation().totems?.earth || EarthTotem.NoEarthTotem,
		setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: number) => {
			const newRotation = player.getRotation();
			if (!newRotation.totems)
				newRotation.totems = ShamanTotems.create();
			newRotation.totems!.earth = newValue;
			player.setRotation(eventID, newRotation);
		},
	});

	const airTotemPicker = new IconEnumPicker(totemDropdownsContainer, simUI.player, {
		extraCssClasses: [
			'air-totem-picker',
		],
		numColumns: 1,
		values: [
			{ color: '#baffc9', value: AirTotem.NoAirTotem },
			{ actionId: { id: { spellId: 25359 }}, value: AirTotem.GraceOfAirTotem },
			{ actionId: { id: { spellId: 25908 }}, value: AirTotem.TranquilAirTotem },
			{ actionId: { id: { spellId: 25587 }}, value: AirTotem.WindfuryTotem },
			{ actionId: { id: { spellId: 3738 }}, value: AirTotem.WrathOfAirTotem },
		],
		equals: (a: AirTotem, b: AirTotem) => a == b,
		zeroValue: AirTotem.NoAirTotem,
		changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.rotationChangeEmitter,
		getValue: (player: Player<Spec.SpecEnhancementShaman>) => player.getRotation().totems?.air || AirTotem.NoAirTotem,
		setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: number) => {
			const newRotation = player.getRotation();
			if (!newRotation.totems)
				newRotation.totems = ShamanTotems.create();
			newRotation.totems!.air = newValue;
			player.setRotation(eventID, newRotation);
		},
	});

	const fireTotemPicker = new IconEnumPicker(totemDropdownsContainer, simUI.player, {
		extraCssClasses: [
			'fire-totem-picker',
		],
		numColumns: 1,
		values: [
			{ color: '#ffb3ba', value: FireTotem.NoFireTotem },
			{ actionId: { id: { spellId: 25552 }}, value: FireTotem.MagmaTotem },
			{ actionId: { id: { spellId: 25533 }}, value: FireTotem.SearingTotem },
			{ actionId: { id: { spellId: 30706 }}, value: FireTotem.TotemOfWrath },
		],
		equals: (a: FireTotem, b: FireTotem) => a == b,
		zeroValue: FireTotem.NoFireTotem,
		changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.rotationChangeEmitter,
		getValue: (player: Player<Spec.SpecEnhancementShaman>) => player.getRotation().totems?.fire || FireTotem.NoFireTotem,
		setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: number) => {
			const newRotation = player.getRotation();
			if (!newRotation.totems)
				newRotation.totems = ShamanTotems.create();
			newRotation.totems!.fire = newValue;
			player.setRotation(eventID, newRotation);
		},
	});

	const waterTotemPicker = new IconEnumPicker(totemDropdownsContainer, simUI.player, {
		extraCssClasses: [
			'water-totem-picker',
		],
		numColumns: 1,
		values: [
			{ color: '#bae1ff', value: WaterTotem.NoWaterTotem },
			{ actionId: { id: { spellId: 25570 }}, value: WaterTotem.ManaSpringTotem },
		],
		equals: (a: WaterTotem, b: WaterTotem) => a == b,
		zeroValue: WaterTotem.NoWaterTotem,
		changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.rotationChangeEmitter,
		getValue: (player: Player<Spec.SpecEnhancementShaman>) => player.getRotation().totems?.water || WaterTotem.NoWaterTotem,
		setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: number) => {
			const newRotation = player.getRotation();
			if (!newRotation.totems)
				newRotation.totems = ShamanTotems.create();
			newRotation.totems!.water = newValue;
			player.setRotation(eventID, newRotation);
		},
	});

	const twistWindfuryPicker = new BooleanPicker(totemInputsContainer, simUI.player, {
		extraCssClasses: [
			'twist-windfury-picker',
		],
		label: 'Twist Windfury',
		labelTooltip: 'Twist Windfury Totem with whichever air totem is selected.',
		changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.rotationChangeEmitter,
		getValue: (player: Player<Spec.SpecEnhancementShaman>) => player.getRotation().totems?.twistWindfury || false,
		setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: boolean) => {
			const newRotation = player.getRotation();
			if (!newRotation.totems)
				newRotation.totems = ShamanTotems.create();
			newRotation.totems!.twistWindfury = newValue;
			player.setRotation(eventID, newRotation);
		},
	});

	const twistFireNovaPicker = new BooleanPicker(totemInputsContainer, simUI.player, {
		extraCssClasses: [
			'twist-fire-nova-picker',
		],
		label: 'Twist Fire Nova',
		labelTooltip: 'Twist Fire Nova Totem with whichever fire totem is selected.',
		changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.rotationChangeEmitter,
		getValue: (player: Player<Spec.SpecEnhancementShaman>) => player.getRotation().totems?.twistFireNova || false,
		setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: boolean) => {
			const newRotation = player.getRotation();
			if (!newRotation.totems)
				newRotation.totems = ShamanTotems.create();
			newRotation.totems!.twistFireNova = newValue;
			player.setRotation(eventID, newRotation);
		},
	});

	return 'Totems';
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
