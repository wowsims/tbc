import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { ElementalShaman_Rotation_RotationType as RotationType } from '/tbc/core/proto/shaman.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
export declare const IconBloodlust: IconPickerConfig<Player<any>, boolean>;
export declare const IconWaterShield: IconPickerConfig<Player<any>, boolean>;
export declare const SnapshotT42Pc: {
    type: "boolean";
    getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
    config: {
        extraCssClasses: string[];
        label: string;
        labelTooltip: string;
        changedEvent: (player: Player<Spec.SpecElementalShaman>) => TypedEvent<void>;
        getValue: (player: Player<Spec.SpecElementalShaman>) => boolean;
        setValue: (eventID: EventID, player: Player<Spec.SpecElementalShaman>, newValue: boolean) => void;
        enableWhen: (player: Player<Spec.SpecElementalShaman>) => boolean;
    };
};
export declare const ElementalShamanRotationConfig: {
    inputs: ({
        type: "enum";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            values: {
                name: string;
                value: RotationType;
                tooltip: string;
            }[];
            changedEvent: (player: Player<Spec.SpecElementalShaman>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecElementalShaman>) => RotationType;
            setValue: (eventID: EventID, player: Player<Spec.SpecElementalShaman>, newValue: number) => void;
            labelTooltip?: undefined;
            enableWhen?: undefined;
        };
    } | {
        type: "number";
        cssClass: string;
        getModObject: (simUI: IndividualSimUI<any>) => Player<any>;
        config: {
            label: string;
            labelTooltip: string;
            changedEvent: (player: Player<Spec.SpecElementalShaman>) => TypedEvent<void>;
            getValue: (player: Player<Spec.SpecElementalShaman>) => number;
            setValue: (eventID: EventID, player: Player<Spec.SpecElementalShaman>, newValue: number) => void;
            enableWhen: (player: Player<Spec.SpecElementalShaman>) => boolean;
            values?: undefined;
        };
    })[];
};
