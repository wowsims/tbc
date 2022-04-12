import { FeralDruid_Rotation_FinishingMove as FinishingMove } from '/tbc/core/proto/druid.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare const SelfInnervate: {
    id: ActionId;
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecFeralDruid>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecFeralDruid>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecFeralDruid>, newValue: boolean) => void;
};
export declare const LatencyMs: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecFeralDruid>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecFeralDruid>) => number;
        setValue: (eventID: EventID, player: Player<Spec.SpecFeralDruid>, newValue: number) => void;
    };
};
export declare const FeralDruidRotationConfig: {
    inputs: ({
        type: "enum";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: FinishingMove;
            }[];
            changedEvent: (player: Player<Spec.SpecFeralDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecFeralDruid>) => FinishingMove;
            setValue: (eventID: EventID, player: Player<Spec.SpecFeralDruid>, newValue: number) => void;
            enableWhen?: undefined;
        };
    } | {
        type: "boolean";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecFeralDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecFeralDruid>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecFeralDruid>, newValue: boolean) => void;
            enableWhen: (player: Player<Spec.SpecFeralDruid>) => boolean;
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
                value: number;
            }[];
            changedEvent: (player: Player<Spec.SpecFeralDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecFeralDruid>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecFeralDruid>, newValue: number) => void;
            enableWhen: (player: Player<Spec.SpecFeralDruid>) => boolean;
        };
    } | {
        type: "boolean";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecFeralDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecFeralDruid>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecFeralDruid>, newValue: boolean) => void;
            values?: undefined;
            enableWhen?: undefined;
        };
    })[];
};
