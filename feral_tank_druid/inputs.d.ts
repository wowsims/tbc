import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { FeralTankDruid_Rotation_Swipe as Swipe } from '/tbc/core/proto/druid.js';
export declare const StartingRage: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecFeralTankDruid>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecFeralTankDruid>) => number;
        setValue: (eventID: EventID, player: Player<Spec.SpecFeralTankDruid>, newValue: number) => void;
    };
};
export declare const FeralTankDruidRotationConfig: {
    inputs: ({
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecFeralTankDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecFeralTankDruid>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecFeralTankDruid>, newValue: number) => void;
            values?: undefined;
            enableWhen?: undefined;
        };
    } | {
        type: "enum";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            values: {
                name: string;
                value: Swipe;
            }[];
            changedEvent: (player: Player<Spec.SpecFeralTankDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecFeralTankDruid>) => Swipe;
            setValue: (eventID: EventID, player: Player<Spec.SpecFeralTankDruid>, newValue: number) => void;
            labelTooltip?: undefined;
            enableWhen?: undefined;
        };
    } | {
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecFeralTankDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecFeralTankDruid>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecFeralTankDruid>, newValue: number) => void;
            enableWhen: (player: Player<Spec.SpecFeralTankDruid>) => boolean;
            values?: undefined;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecFeralTankDruid>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecFeralTankDruid>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecFeralTankDruid>, newValue: boolean) => void;
            values?: undefined;
            enableWhen?: undefined;
        };
    })[];
};
