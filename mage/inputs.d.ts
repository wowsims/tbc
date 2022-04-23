import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Mage_Rotation_Type as RotationType } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_FireRotation_PrimarySpell as PrimaryFireSpell } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_AoeRotation_Rotation as AoeRotationSpells } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_ArcaneRotation_Filler as ArcaneFiller } from '/tbc/core/proto/mage.js';
export declare const MageArmor: {
    id: ActionId;
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecMage>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => void;
};
export declare const MoltenArmor: {
    id: ActionId;
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecMage>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: boolean) => void;
};
export declare const EvocationTicks: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecMage>) => number;
        setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => void;
    };
};
export declare const MageRotationConfig: {
    inputs: ({
        type: "enum";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: RotationType;
            }[];
            changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecMage>) => RotationType;
            setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => void;
            showWhen?: undefined;
            enableWhen?: undefined;
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
            extraCssClasses?: undefined;
            values?: undefined;
            showWhen?: undefined;
            enableWhen?: undefined;
        };
    } | {
        type: "enum";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            values: {
                name: string;
                value: AoeRotationSpells;
            }[];
            changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecMage>) => AoeRotationSpells.Flamestrike | AoeRotationSpells.Blizzard | 0;
            setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => void;
            showWhen: (player: Player<Spec.SpecMage>) => boolean;
            labelTooltip?: undefined;
            enableWhen?: undefined;
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
            enableWhen?: undefined;
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
            enableWhen?: undefined;
        };
    } | {
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecMage>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => void;
            showWhen: (player: Player<Spec.SpecMage>) => boolean;
            enableWhen: (player: Player<Spec.SpecMage>) => boolean;
            extraCssClasses?: undefined;
            values?: undefined;
        };
    } | {
        type: "enum";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: ArcaneFiller;
            }[];
            changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecMage>) => ArcaneFiller;
            setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => void;
            showWhen: (player: Player<Spec.SpecMage>) => boolean;
            enableWhen?: undefined;
        };
        cssClass?: undefined;
    } | {
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecMage>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecMage>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecMage>, newValue: number) => void;
            showWhen: (player: Player<Spec.SpecMage>) => boolean;
            extraCssClasses?: undefined;
            values?: undefined;
            enableWhen?: undefined;
        };
    })[];
};
