import { Warlock_Rotation_PrimarySpell as PrimarySpell, Warlock_Rotation_Curse as Curse, Warlock_Options_Summon as Summon } from '/tbc/core/proto/warlock.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare const FelArmor: {
    id: ActionId;
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecWarlock>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecWarlock>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecWarlock>, newValue: boolean) => void;
};
export declare const DemonArmor: {
    id: ActionId;
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecWarlock>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecWarlock>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecWarlock>, newValue: boolean) => void;
};
export declare const Sacrifice: {
    id: ActionId;
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecWarlock>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecWarlock>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecWarlock>, newValue: boolean) => void;
};
export declare const DemonSummon: {
    extraCssClasses: string[];
    numColumns: number;
    values: ({
        color: string;
        value: Summon;
        actionId?: undefined;
    } | {
        actionId: ActionId;
        value: Summon;
        color?: undefined;
    })[];
    equals: (a: Summon, b: Summon) => boolean;
    zeroValue: Summon;
    changedEvent: (player: Player<Spec.SpecWarlock>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecWarlock>) => Summon;
    setValue: (eventID: EventID, player: Player<Spec.SpecWarlock>, newValue: number) => void;
};
export declare const WarlockRotationConfig: {
    inputs: ({
        type: "enum";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: PrimarySpell;
            }[];
            changedEvent: (player: Player<Spec.SpecWarlock>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarlock>) => PrimarySpell;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarlock>, newValue: number) => void;
            enableWhen?: undefined;
        };
    } | {
        type: "boolean";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecWarlock>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarlock>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarlock>, newValue: boolean) => void;
            values?: undefined;
            enableWhen?: undefined;
        };
    } | {
        type: "boolean";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecWarlock>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarlock>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarlock>, newValue: boolean) => void;
            enableWhen: (player: Player<Spec.SpecWarlock>) => boolean;
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
                value: Curse;
            }[];
            changedEvent: (player: Player<Spec.SpecWarlock>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarlock>) => Curse;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarlock>, newValue: number) => void;
            enableWhen?: undefined;
        };
    })[];
};
