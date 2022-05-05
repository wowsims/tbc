import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { WarriorShout, Warrior_Rotation_SunderArmor as SunderArmor } from '/tbc/core/proto/warrior.js';
export declare const Recklessness: {
    id: ActionId;
    states: number;
    extraCssClasses: string[];
    changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecWarrior>) => boolean;
    setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => void;
};
export declare const StartingRage: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecWarrior>) => number;
        setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => void;
    };
};
export declare const ShoutPicker: {
    extraCssClasses: string[];
    numColumns: number;
    values: ({
        color: string;
        value: WarriorShout;
        actionId?: undefined;
    } | {
        actionId: ActionId;
        value: WarriorShout;
        color?: undefined;
    })[];
    equals: (a: WarriorShout, b: WarriorShout) => boolean;
    zeroValue: WarriorShout;
    changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
    getValue: (player: Player<Spec.SpecWarrior>) => WarriorShout;
    setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => void;
};
export declare const PrecastShout: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecWarrior>) => boolean;
        setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => void;
    };
};
export declare const PrecastShoutWithSapphire: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecWarrior>) => boolean;
        setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => void;
        enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    };
};
export declare const PrecastShoutWithT2: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecWarrior>) => boolean;
        setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => void;
        enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    };
};
export declare const WarriorRotationConfig: {
    inputs: ({
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarrior>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: boolean) => void;
            showWhen?: undefined;
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
            showWhen?: undefined;
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
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            extraCssClasses: string[];
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarrior>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => void;
            showWhen: (player: Player<Spec.SpecWarrior>) => boolean;
            values?: undefined;
        };
    } | {
        type: "enum";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            values: {
                name: string;
                value: SunderArmor;
            }[];
            changedEvent: (player: Player<Spec.SpecWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecWarrior>) => SunderArmor;
            setValue: (eventID: EventID, player: Player<Spec.SpecWarrior>, newValue: number) => void;
            labelTooltip?: undefined;
            showWhen?: undefined;
            extraCssClasses?: undefined;
        };
    })[];
};
