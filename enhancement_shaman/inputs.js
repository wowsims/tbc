import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { IconEnumPicker } from '/tbc/core/components/icon_enum_picker.js';
import { AirTotem, EarthTotem, FireTotem, WaterTotem, EnhancementShaman_Rotation_RotationType as RotationType, ShamanTotems } from '/tbc/core/proto/shaman.js';
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
            type: 'enum', cssClass: 'rotation-enum-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Type',
                values: [
                    {
                        name: 'Basic', value: RotationType.Basic,
                        tooltip: 'does basic stuff',
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().type,
                setValue: (eventID, player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.type = newValue;
                    player.setRotation(eventID, newRotation);
                },
            },
        }
    ],
};
function makeBooleanShamanBuffInput(id, optionsFieldName) {
    return {
        id: id,
        states: 2,
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions()[optionsFieldName],
        setValue: (eventID, player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions[optionsFieldName] = newValue;
            player.setSpecOptions(eventID, newOptions);
        },
    };
}
export function TotemsSection(simUI, parentElem) {
    const customSectionsContainer = parentElem.closest('.custom-sections-container');
    customSectionsContainer.style.zIndex = '1000';
    parentElem.innerHTML = `
		<div class="totem-dropdowns-container"></div>
		<div class="totem-inputs-container"></div>
	`;
    const totemDropdownsContainer = parentElem.getElementsByClassName('totem-dropdowns-container')[0];
    const totemInputsContainer = parentElem.getElementsByClassName('totem-inputs-container')[0];
    const earthTotemPicker = new IconEnumPicker(totemDropdownsContainer, simUI.player, {
        extraCssClasses: [
            'earth-totem-picker',
        ],
        values: [
            { color: '#ffdfba', value: EarthTotem.NoEarthTotem },
            { id: { spellId: 25528 }, value: EarthTotem.StrengthOfEarthTotem },
            { id: { spellId: 8143 }, value: EarthTotem.TremorTotem },
        ],
        changedEvent: (player) => player.rotationChangeEmitter,
        getValue: (player) => player.getRotation().totems?.earth || EarthTotem.NoEarthTotem,
        setValue: (eventID, player, newValue) => {
            const newRotation = player.getRotation();
            if (!newRotation.totems)
                newRotation.totems = ShamanTotems.create();
            newRotation.totems.earth = newValue;
            player.setRotation(eventID, newRotation);
        },
    });
    const airTotemPicker = new IconEnumPicker(totemDropdownsContainer, simUI.player, {
        extraCssClasses: [
            'air-totem-picker',
        ],
        values: [
            { color: '#baffc9', value: AirTotem.NoAirTotem },
            { id: { spellId: 25359 }, value: AirTotem.GraceOfAirTotem },
            { id: { spellId: 25908 }, value: AirTotem.TranquilAirTotem },
            { id: { spellId: 25587 }, value: AirTotem.WindfuryTotem },
            { id: { spellId: 3738 }, value: AirTotem.WrathOfAirTotem },
        ],
        changedEvent: (player) => player.rotationChangeEmitter,
        getValue: (player) => player.getRotation().totems?.air || AirTotem.NoAirTotem,
        setValue: (eventID, player, newValue) => {
            const newRotation = player.getRotation();
            if (!newRotation.totems)
                newRotation.totems = ShamanTotems.create();
            newRotation.totems.air = newValue;
            player.setRotation(eventID, newRotation);
        },
    });
    const fireTotemPicker = new IconEnumPicker(totemDropdownsContainer, simUI.player, {
        extraCssClasses: [
            'fire-totem-picker',
        ],
        values: [
            { color: '#ffb3ba', value: FireTotem.NoFireTotem },
            { id: { spellId: 25552 }, value: FireTotem.MagmaTotem },
            { id: { spellId: 25533 }, value: FireTotem.SearingTotem },
            { id: { spellId: 30706 }, value: FireTotem.TotemOfWrath },
        ],
        changedEvent: (player) => player.rotationChangeEmitter,
        getValue: (player) => player.getRotation().totems?.fire || FireTotem.NoFireTotem,
        setValue: (eventID, player, newValue) => {
            const newRotation = player.getRotation();
            if (!newRotation.totems)
                newRotation.totems = ShamanTotems.create();
            newRotation.totems.fire = newValue;
            player.setRotation(eventID, newRotation);
        },
    });
    const waterTotemPicker = new IconEnumPicker(totemDropdownsContainer, simUI.player, {
        extraCssClasses: [
            'water-totem-picker',
        ],
        values: [
            { color: '#bae1ff', value: WaterTotem.NoWaterTotem },
            { id: { spellId: 25570 }, value: WaterTotem.ManaSpringTotem },
        ],
        changedEvent: (player) => player.rotationChangeEmitter,
        getValue: (player) => player.getRotation().totems?.water || WaterTotem.NoWaterTotem,
        setValue: (eventID, player, newValue) => {
            const newRotation = player.getRotation();
            if (!newRotation.totems)
                newRotation.totems = ShamanTotems.create();
            newRotation.totems.water = newValue;
            player.setRotation(eventID, newRotation);
        },
    });
    const twistWindfuryPicker = new BooleanPicker(totemInputsContainer, simUI.player, {
        extraCssClasses: [
            'twist-windfury-picker',
        ],
        label: 'Twist Windfury',
        labelTooltip: 'Twist Windfury Totem with whichever air totem is selected.',
        changedEvent: (player) => player.rotationChangeEmitter,
        getValue: (player) => player.getRotation().totems?.twistWindfury || false,
        setValue: (eventID, player, newValue) => {
            const newRotation = player.getRotation();
            if (!newRotation.totems)
                newRotation.totems = ShamanTotems.create();
            newRotation.totems.twistWindfury = newValue;
            player.setRotation(eventID, newRotation);
        },
    });
    const twistFireNovaPicker = new BooleanPicker(totemInputsContainer, simUI.player, {
        extraCssClasses: [
            'twist-fire-nova-picker',
        ],
        label: 'Twist Fire Nova',
        labelTooltip: 'Twist Fire Nova Totem with whichever fire totem is selected.',
        changedEvent: (player) => player.rotationChangeEmitter,
        getValue: (player) => player.getRotation().totems?.twistFireNova || false,
        setValue: (eventID, player, newValue) => {
            const newRotation = player.getRotation();
            if (!newRotation.totems)
                newRotation.totems = ShamanTotems.create();
            newRotation.totems.twistFireNova = newValue;
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
