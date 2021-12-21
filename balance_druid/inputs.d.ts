import { BalanceDruid_Rotation_PrimarySpell as PrimarySpell } from '/tbc/core/proto/druid.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare const SelfInnervate: {
    id: {
        spellId: number;
    };
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecBalanceDruid>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecBalanceDruid>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecBalanceDruid>, newValue: boolean) => void;
};
export declare const BalanceDruidRotationConfig: {
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
            changedEvent: (player: Player<Spec.SpecBalanceDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecBalanceDruid>) => PrimarySpell;
            setValue: (eventID: EventID, player: Player<Spec.SpecBalanceDruid>, newValue: number) => void;
            enableWhen?: undefined;
        };
    } | {
        type: "boolean";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecBalanceDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecBalanceDruid>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecBalanceDruid>, newValue: boolean) => void;
            enableWhen: (player: Player<Spec.SpecBalanceDruid>) => boolean;
            values?: undefined;
        };
    } | {
        type: "boolean";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecBalanceDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecBalanceDruid>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecBalanceDruid>, newValue: boolean) => void;
            values?: undefined;
            enableWhen?: undefined;
        };
    })[];
};
