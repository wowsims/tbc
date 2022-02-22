import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { EnhancementShaman_Rotation_PrimaryShock as PrimaryShock } from '/tbc/core/proto/shaman.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
export declare const IconBloodlust: IconPickerConfig<Player<any>, boolean>;
export declare const IconWaterShield: IconPickerConfig<Player<any>, boolean>;
export declare const DelayOffhandSwings: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecEnhancementShaman>) => boolean;
        setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: boolean) => void;
    };
};
export declare const SnapshotT42Pc: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecEnhancementShaman>) => boolean;
        setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: boolean) => void;
        enableWhen: (player: Player<Spec.SpecEnhancementShaman>) => boolean;
    };
};
export declare const EnhancementShamanRotationConfig: {
    inputs: ({
        type: "enum";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            values: {
                name: string;
                value: PrimaryShock;
            }[];
            changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecEnhancementShaman>) => PrimaryShock;
            setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: number) => void;
            labelTooltip?: undefined;
        };
    } | {
        type: "boolean";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecEnhancementShaman>) => boolean;
            setValue: (eventID: EventID, player: Player<Spec.SpecEnhancementShaman>, newValue: boolean) => void;
            values?: undefined;
        };
    })[];
};
