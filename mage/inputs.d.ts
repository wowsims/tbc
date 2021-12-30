import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Mage_Rotation_Type as RotationType } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_FireRotation_PrimarySpell as PrimaryFireSpell } from '/tbc/core/proto/mage.js';
export declare const MageArmor: {
    id: {
        spellId: number;
    };
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecMage>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => void;
};
export declare const MoltenArmor: {
    id: {
        spellId: number;
    };
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecMage>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => void;
};
export declare const MageRotationConfig: {
    inputs: ({
        type: "enum";
        getModObject: (simUI: IndividualSimUI<any>) => IndividualSimUI<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: RotationType;
            }[];
            changedEvent: (simUI: IndividualSimUI<Spec.SpecMage>) => TypedEvent<void>;
            getValue: (simUI: IndividualSimUI<Spec.SpecMage>) => RotationType;
            setValue: (eventID: EventID, simUI: IndividualSimUI<Spec.SpecMage>, newValue: number) => void;
            showWhen?: undefined;
        };
        cssClass?: undefined;
    } | {
        type: "enum";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            values: {
                name: string;
                value: PrimaryFireSpell;
            }[];
            changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecMage>) => PrimaryFireSpell;
            setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => void;
            showWhen: (player: Player<Spec.SpecMage>) => boolean;
            labelTooltip?: undefined;
        };
        cssClass?: undefined;
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecMage>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => void;
            showWhen: (player: Player<Spec.SpecMage>) => boolean;
            extraCssClasses?: undefined;
            values?: undefined;
        };
    })[];
};
