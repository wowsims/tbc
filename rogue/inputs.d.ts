import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
export declare const RogueRotationConfig: {
    inputs: {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecRogue>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecRogue>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecRogue>, newValue: boolean) => void;
        };
    }[];
};
