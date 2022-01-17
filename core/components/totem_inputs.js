import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { IconEnumPicker } from '/tbc/core/components/icon_enum_picker.js';
import { AirTotem, EarthTotem, FireTotem, WaterTotem, ShamanTotems, } from '/tbc/core/proto/shaman.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
export function TotemsSection(simUI, parentElem) {
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
        numColumns: 1,
        values: [
            { color: '#ffdfba', value: EarthTotem.NoEarthTotem },
            { actionId: ActionId.fromSpellId(25528), value: EarthTotem.StrengthOfEarthTotem },
            { actionId: ActionId.fromSpellId(8143), value: EarthTotem.TremorTotem },
        ],
        equals: (a, b) => a == b,
        zeroValue: EarthTotem.NoEarthTotem,
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
        numColumns: 1,
        values: [
            { color: '#baffc9', value: AirTotem.NoAirTotem },
            { actionId: ActionId.fromSpellId(25359), value: AirTotem.GraceOfAirTotem },
            { actionId: ActionId.fromSpellId(25908), value: AirTotem.TranquilAirTotem },
            { actionId: ActionId.fromSpellId(25587), value: AirTotem.WindfuryTotem },
            { actionId: ActionId.fromSpellId(3738), value: AirTotem.WrathOfAirTotem },
        ],
        equals: (a, b) => a == b,
        zeroValue: AirTotem.NoAirTotem,
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
        numColumns: 1,
        values: [
            { color: '#ffb3ba', value: FireTotem.NoFireTotem },
            { actionId: ActionId.fromSpellId(25552), value: FireTotem.MagmaTotem },
            { actionId: ActionId.fromSpellId(25533), value: FireTotem.SearingTotem },
            { actionId: ActionId.fromSpellId(30706), value: FireTotem.TotemOfWrath },
        ],
        equals: (a, b) => a == b,
        zeroValue: FireTotem.NoFireTotem,
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
        numColumns: 1,
        values: [
            { color: '#bae1ff', value: WaterTotem.NoWaterTotem },
            { actionId: ActionId.fromSpellId(25570), value: WaterTotem.ManaSpringTotem },
        ],
        equals: (a, b) => a == b,
        zeroValue: WaterTotem.NoWaterTotem,
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
    const windfuryRankPicker = new EnumPicker(totemInputsContainer, simUI.player, {
        extraCssClasses: [
            'windfury-rank-picker',
        ],
        values: [
            { name: 'No WF', value: 0 },
            { name: '1', value: 1 },
            { name: '2', value: 2 },
            { name: '3', value: 3 },
            { name: '4', value: 4 },
            { name: '5', value: 5 },
        ],
        label: 'WF Totem Rank',
        labelTooltip: 'Rank of Windfury Totem to use, if using Windfury Totem.',
        changedEvent: (player) => player.rotationChangeEmitter,
        getValue: (player) => player.getRotation().totems?.windfuryTotemRank || 0,
        setValue: (eventID, player, newValue) => {
            const newRotation = player.getRotation();
            if (!newRotation.totems)
                newRotation.totems = ShamanTotems.create();
            newRotation.totems.windfuryTotemRank = newValue;
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
