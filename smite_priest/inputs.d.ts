import { SmitePriest_Rotation_RotationType as RotationType } from '/tbc/core/proto/priest.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
export declare const SelfPowerInfusion: {
    id: ActionId;
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecSmitePriest>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecSmitePriest>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecSmitePriest>, newValue: boolean) => void;
};
export declare const SmitePriestRotationConfig: {
    inputs: ({
        type: "enum";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: RotationType;
            }[];
            changedEvent: (player: Player<Spec.SpecSmitePriest>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecSmitePriest>) => RotationType;
            setValue: (eventID: EventID, player: Player<Spec.SpecSmitePriest>, newValue: number) => void;
            enableWhen?: undefined;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecSmitePriest>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecSmitePriest>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecSmitePriest>, newValue: boolean) => void;
            values?: undefined;
            enableWhen?: undefined;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecSmitePriest>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecSmitePriest>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecSmitePriest>, newValue: boolean) => void;
            enableWhen: (player: Player<Spec.SpecSmitePriest>) => boolean;
            values?: undefined;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecSmitePriest>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecSmitePriest>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecSmitePriest>, newValue: boolean) => void;
            enableWhen: (player: Player<Spec.SpecShadowPriest>) => boolean;
            values?: undefined;
        };
    })[];
};
