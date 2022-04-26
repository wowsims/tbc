import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { EventID } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { RetributionPaladin_Rotation_ConsecrationRank as ConsecrationRank, RetributionPaladin_Options_Judgement as Judgement } from '/tbc/core/proto/paladin.js';
export declare const RetributionPaladinRotationConfig: {
    inputs: ({
        type: "enum";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: ConsecrationRank;
            }[];
            changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => import("/tbc/core/typed_event.js").TypedEvent<void>;
            getValue: (player: Player<Spec.SpecRetributionPaladin>) => ConsecrationRank;
            setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: number) => void;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => import("/tbc/core/typed_event.js").TypedEvent<void>;
            getValue: (player: Player<Spec.SpecRetributionPaladin>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: boolean) => void;
            values?: undefined;
        };
    })[];
};
export declare const JudgementSelection: {
    type: "enum";
    cssClass: string;
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        label: string;
        labelTooltip: string;
        values: {
            name: string;
            value: Judgement;
        }[];
        changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => import("/tbc/core/typed_event.js").TypedEvent<void>;
        getValue: (player: Player<Spec.SpecRetributionPaladin>) => Judgement;
        setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: number) => void;
    };
};
export declare const CrusaderStrikeDelayMS: {
    type: "number";
    cssClass: string;
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => import("/tbc/core/typed_event.js").TypedEvent<void>;
        getValue: (player: Player<Spec.SpecRetributionPaladin>) => number;
        setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: number) => void;
    };
};
/*** Leave this for now. We'll ignore HasteLeeway for initial release, but we might come back to it at some point  ***/
export declare const DamgeTakenPerSecond: {
    type: "number";
    cssClass: string;
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => import("/tbc/core/typed_event.js").TypedEvent<void>;
        getValue: (player: Player<Spec.SpecRetributionPaladin>) => number;
        setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: number) => void;
    };
};
