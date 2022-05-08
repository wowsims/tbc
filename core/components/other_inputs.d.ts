import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { Conjured } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Party } from '/tbc/core/party.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { Target } from '/tbc/core/target.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
export declare function makeShow1hWeaponsSelector(parent: HTMLElement, sim: Sim): BooleanPicker<Sim>;
export declare function makeShow2hWeaponsSelector(parent: HTMLElement, sim: Sim): BooleanPicker<Sim>;
export declare function makeShowMatchingGemsSelector(parent: HTMLElement, sim: Sim): BooleanPicker<Sim>;
export declare function makePhaseSelector(parent: HTMLElement, sim: Sim): EnumPicker<Sim>;
export declare const StartingPotion: {
    type: "enum";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        values: {
            name: string;
            value: Potions;
        }[];
        changedEvent: (player: Player<any>) => TypedEvent<void>;
        getValue: (player: Player<any>) => Potions;
        setValue: (eventID: EventID, player: Player<any>, newValue: number) => void;
    };
};
export declare const NumStartingPotions: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<any>) => TypedEvent<void>;
        getValue: (player: Player<any>) => number;
        setValue: (eventID: EventID, player: Player<any>, newValue: number) => void;
        enableWhen: (player: Player<any>) => boolean;
    };
};
export declare const StartingConjured: {
    type: "enum";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        values: {
            name: string;
            value: Conjured;
        }[];
        changedEvent: (player: Player<any>) => TypedEvent<void>;
        getValue: (player: Player<any>) => Conjured;
        setValue: (eventID: EventID, player: Player<any>, newValue: number) => void;
    };
};
export declare const NumStartingConjured: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<any>) => TypedEvent<void>;
        getValue: (player: Player<any>) => number;
        setValue: (eventID: EventID, player: Player<any>, newValue: number) => void;
        enableWhen: (player: Player<any>) => boolean;
    };
};
export declare const ShadowPriestDPS: {
    type: "number";
    cssClass: string;
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        changedEvent: (player: Player<any>) => TypedEvent<void>;
        getValue: (player: Player<any>) => number;
        setValue: (eventID: EventID, player: Player<any>, newValue: number) => void;
    };
};
export declare const ISBUptime: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Target;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (target: Target) => TypedEvent<void>;
        getValue: (target: Target) => number;
        setValue: (eventID: EventID, target: Target, newValue: number) => void;
    };
};
export declare const ExposeWeaknessUptime: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Target;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (target: Target) => TypedEvent<void>;
        getValue: (target: Target) => number;
        setValue: (eventID: EventID, target: Target, newValue: number) => void;
    };
};
export declare const ExposeWeaknessHunterAgility: {
    type: "number";
    getModObject: (simUI: IndividualSimUI<any>) => Target;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (target: Target) => TypedEvent<void>;
        getValue: (target: Target) => number;
        setValue: (eventID: EventID, target: Target, newValue: number) => void;
    };
};
export declare const SnapshotImprovedStrengthOfEarthTotem: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Party;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (party: Party) => TypedEvent<void>;
        getValue: (party: Party) => boolean;
        setValue: (eventID: EventID, party: Party, newValue: boolean) => void;
        enableWhen: (party: Party) => boolean;
    };
};
export declare const SnapshotImprovedWrathOfAirTotem: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Party;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (party: Party) => TypedEvent<void>;
        getValue: (party: Party) => boolean;
        setValue: (eventID: EventID, party: Party, newValue: boolean) => void;
        enableWhen: (party: Party) => boolean;
    };
};
export declare const SnapshotBsSolarianSapphire: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Party;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (party: Party) => TypedEvent<void>;
        getValue: (party: Party) => boolean;
        setValue: (eventID: EventID, party: Party, newValue: boolean) => void;
        enableWhen: (party: Party) => boolean;
    };
};
export declare const SnapshotBsT2: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Party;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (party: Party) => TypedEvent<void>;
        getValue: (party: Party) => boolean;
        setValue: (eventID: EventID, party: Party, newValue: boolean) => void;
        enableWhen: (party: Party) => boolean;
    };
};
export declare const InFrontOfTarget: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<any>) => TypedEvent<void>;
        getValue: (player: Player<any>) => boolean;
        setValue: (eventID: EventID, player: Player<any>, newValue: boolean) => void;
    };
};
export declare const TankAssignment: {
    type: "enum";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        values: {
            name: string;
            value: number;
        }[];
        changedEvent: (player: Player<any>) => TypedEvent<void>;
        getValue: (player: Player<any>) => number;
        setValue: (eventID: EventID, player: Player<any>, newValue: number) => void;
    };
};
