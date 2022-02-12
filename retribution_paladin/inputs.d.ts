import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare const RetributionPaladinRotationConfig: {
    inputs: {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecRetributionPaladin>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecRetributionPaladin>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecRetributionPaladin>, newValue: boolean) => void;
        };
    }[];
};
