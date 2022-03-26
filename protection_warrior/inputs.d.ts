import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { ProtectionWarrior_Rotation_DemoShout as DemoShout, ProtectionWarrior_Rotation_ThunderClap as ThunderClap } from '/tbc/core/proto/warrior.js';
export declare const ProtectionWarriorRotationConfig: {
    inputs: ({
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
        };
    })[];
};
