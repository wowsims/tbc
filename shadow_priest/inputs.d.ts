import { ShadowPriest_Rotation_RotationType as RotationType } from '/tbc/core/proto/priest.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
export declare const ShadowPriestRotationConfig: {
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
            changedEvent: (player: Player<Spec.SpecShadowPriest>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecShadowPriest>) => RotationType;
            setValue: (eventID: EventID, player: Player<Spec.SpecShadowPriest>, newValue: number) => void;
            enableWhen?: undefined;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecShadowPriest>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecShadowPriest>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecShadowPriest>, newValue: boolean) => void;
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
            changedEvent: (player: Player<Spec.SpecShadowPriest>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecShadowPriest>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecShadowPriest>, newValue: boolean) => void;
            enableWhen: (player: Player<Spec.SpecShadowPriest>) => boolean;
            values?: undefined;
        };
    } | {
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecShadowPriest>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecShadowPriest>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecShadowPriest>, newValue: number) => void;
            values?: undefined;
            enableWhen?: undefined;
        };
    })[];
};
