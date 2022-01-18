import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Warrior_Rotation_Type as RotationType } from '/tbc/core/proto/warrior.js';
import { Warrior_Rotation_FuryRotation_PrimaryInstant as PrimaryInstant } from '/tbc/core/proto/warrior.js';
export declare const Recklessness: {
    id: ActionId;
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecWarrior>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => void;
};
export declare const WarriorRotationConfig: {
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
            changedEvent: (simUI: IndividualSimUI<Spec.SpecWarrior>) => TypedEvent<void>;
            getValue: (simUI: IndividualSimUI<Spec.SpecWarrior>) => RotationType;
            setValue: (eventID: EventID, simUI: IndividualSimUI<Spec.SpecWarrior>, newValue: number) => void;
            showWhen?: undefined;
        };
        cssClass?: undefined;
    } | {
        type: "enum";
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: PrimaryInstant;
            }[];
            changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarrior>) => PrimaryInstant.Whirlwind;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => void;
            showWhen: (player: Player<Spec.SpecWarrior>) => boolean;
        };
        cssClass?: undefined;
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarrior>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => void;
            showWhen: (player: Player<Spec.SpecWarrior>) => boolean;
            extraCssClasses?: undefined;
            values?: undefined;
        };
    } | {
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarrior>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => void;
            showWhen: (player: Player<Spec.SpecWarrior>) => boolean;
            extraCssClasses?: undefined;
            values?: undefined;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarrior>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => void;
            extraCssClasses?: undefined;
            values?: undefined;
            showWhen?: undefined;
        };
    } | {
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarrior>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => void;
            extraCssClasses?: undefined;
            values?: undefined;
            showWhen?: undefined;
        };
    })[];
};
