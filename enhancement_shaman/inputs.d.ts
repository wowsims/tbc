import { IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { EnhancementShaman_Rotation_RotationType as RotationType } from '/tbc/core/proto/shaman.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare const IconBloodlust: IconPickerConfig<Player<any>, boolean>;
export declare const IconWaterShield: IconPickerConfig<Player<any>, boolean>;
export declare const EnhancementShamanRotationConfig: {
    inputs: {
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
            changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => import("/tbc/core/typed_event").TypedEvent<void>;
            getValue: (player: Player<Spec.SpecEnhancementShaman>) => RotationType;
            setValue: (player: Player<Spec.SpecEnhancementShaman>, newValue: number) => void;
        };
    }[];
};
