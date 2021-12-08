import { ShadowPriest_Rotation_RotationType as RotationType } from '/tbc/core/proto/priest.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { SimUI } from '/tbc/core/sim_ui.js';
export declare const ShadowPriestRotationConfig: {
    inputs: ({
        type: "enum";
        cssClass: string;
        getModObject: (simUI: SimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: RotationType;
            }[];
            changedEvent: (player: Player<Spec.SpecShadowPriest>) => import("/tbc/core/typed_event").TypedEvent<void>;
            getValue: (player: Player<Spec.SpecShadowPriest>) => RotationType;
            setValue: (player: Player<Spec.SpecShadowPriest>, newValue: number) => void;
            enableWhen?: undefined;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: SimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecShadowPriest>) => import("/tbc/core/typed_event").TypedEvent<void>;
            getValue: (player: Player<Spec.SpecShadowPriest>) => boolean;
            setValue: (player: Player<Spec.SpecShadowPriest>, newValue: boolean) => void;
            values?: undefined;
            enableWhen?: undefined;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: SimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecShadowPriest>) => import("/tbc/core/typed_event").TypedEvent<void>;
            getValue: (player: Player<Spec.SpecShadowPriest>) => boolean;
            setValue: (player: Player<Spec.SpecShadowPriest>, newValue: boolean) => void;
            enableWhen: (player: Player<Spec.SpecShadowPriest>) => boolean;
            values?: undefined;
        };
    } | {
        type: "number";
        cssClass: string;
        getModObject: (simUI: SimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecShadowPriest>) => import("/tbc/core/typed_event").TypedEvent<void>;
            getValue: (player: Player<Spec.SpecShadowPriest>) => number;
            setValue: (player: Player<Spec.SpecShadowPriest>, newValue: number) => void;
            values?: undefined;
            enableWhen?: undefined;
        };
    })[];
};
