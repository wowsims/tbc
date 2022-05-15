import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { WarriorShout, ProtectionWarrior_Rotation_DemoShout as DemoShout, ProtectionWarrior_Rotation_ShieldBlock as ShieldBlock, ProtectionWarrior_Rotation_ThunderClap as ThunderClap } from '/tbc/core/proto/warrior.js';
export declare const StartingRage: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecProtectionWarrior>) => number;
        setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => void;
    };
};
export declare const ShoutPicker: {
    type: "enum";
    cssClass: string;
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        label: string;
        labelTooltip: string;
        values: {
            name: string;
            value: WarriorShout;
        }[];
        changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecProtectionWarrior>) => WarriorShout;
        setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => void;
    };
};
export declare const PrecastShout: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecProtectionWarrior>) => boolean;
        setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: boolean) => void;
    };
};
export declare const PrecastShoutWithSapphire: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecProtectionWarrior>) => boolean;
        setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: boolean) => void;
        enableWhen: (player: Player<Spec.SpecProtectionWarrior>) => boolean;
    };
};
export declare const PrecastShoutWithT2: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecProtectionWarrior>) => boolean;
        setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: boolean) => void;
        enableWhen: (player: Player<Spec.SpecProtectionWarrior>) => boolean;
    };
};
export declare const ProtectionWarriorRotationConfig: {
    inputs: ({
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecProtectionWarrior>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: boolean) => void;
            values?: undefined;
        };
    } | {
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecProtectionWarrior>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => void;
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
                value: DemoShout;
            }[];
            changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecProtectionWarrior>) => DemoShout;
            setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => void;
            labelTooltip?: undefined;
        };
    } | {
        type: "enum";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            values: {
                name: string;
                value: ThunderClap;
            }[];
            changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecProtectionWarrior>) => ThunderClap;
            setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => void;
            labelTooltip?: undefined;
        };
    } | {
        type: "enum";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            values: {
                name: string;
                value: ShieldBlock;
            }[];
            changedEvent: (player: Player<Spec.SpecProtectionWarrior>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecProtectionWarrior>) => ShieldBlock;
            setValue: (eventID: EventID, player: Player<Spec.SpecProtectionWarrior>, newValue: number) => void;
        };
    })[];
};
