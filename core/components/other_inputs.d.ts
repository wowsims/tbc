import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { SimUI } from '/tbc/core/sim_ui.js';
export declare function makePhaseSelector(parent: HTMLElement, sim: Sim): EnumPicker<Sim>;
export declare const StartingPotion: {
    type: "enum";
    cssClass: string;
    getModObject: (simUI: SimUI<any>) => Player<any>;
    config: {
        label: string;
        labelTooltip: string;
        values: {
            name: string;
            value: Potions;
        }[];
        changedEvent: (player: Player<any>) => import("../typed_event").TypedEvent<void>;
        getValue: (player: Player<any>) => Potions;
        setValue: (player: Player<any>, newValue: number) => void;
    };
};
export declare const NumStartingPotions: {
    type: "number";
    cssClass: string;
    getModObject: (simUI: SimUI<any>) => Player<any>;
    config: {
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<any>) => import("../typed_event").TypedEvent<void>;
        getValue: (player: Player<any>) => number;
        setValue: (player: Player<any>, newValue: number) => void;
        enableWhen: (player: Player<any>) => boolean;
    };
};
export declare const ShadowPriestDPS: {
    type: "number";
    cssClass: string;
    getModObject: (simUI: SimUI<any>) => Player<any>;
    config: {
        label: string;
        changedEvent: (player: Player<any>) => import("../typed_event").TypedEvent<void>;
        getValue: (player: Player<any>) => number;
        setValue: (player: Player<any>, newValue: number) => void;
    };
};
