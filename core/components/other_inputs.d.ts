import { Potions } from '/tbc/core/proto/common.js';
import { Sim } from '/tbc/core/sim.js';
export declare const StartingPotion: {
    type: "enum";
    cssClass: string;
    config: {
        label: string;
        labelTooltip: string;
        values: {
            name: string;
            value: Potions;
        }[];
        changedEvent: (sim: Sim<any>) => import("../typed_event").TypedEvent<void>;
        getValue: (sim: Sim<any>) => Potions;
        setValue: (sim: Sim<any>, newValue: number) => void;
    };
};
export declare const NumStartingPotions: {
    type: "number";
    cssClass: string;
    config: {
        label: string;
        labelTooltip: string;
        changedEvent: (sim: Sim<any>) => import("../typed_event").TypedEvent<void>;
        getValue: (sim: Sim<any>) => number;
        setValue: (sim: Sim<any>, newValue: number) => void;
        enableWhen: (sim: Sim<any>) => boolean;
    };
};
export declare const ShadowPriestDPS: {
    type: "number";
    cssClass: string;
    config: {
        label: string;
        changedEvent: (sim: Sim<any>) => import("../typed_event").TypedEvent<void>;
        getValue: (sim: Sim<any>) => number;
        setValue: (sim: Sim<any>, newValue: number) => void;
    };
};
