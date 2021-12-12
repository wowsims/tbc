import { IconInput } from '/tbc/core/components/icon_picker.js';
import { ElementalShaman_Rotation_RotationType as RotationType } from '/tbc/core/proto/shaman.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare const IconBloodlust: IconInput<Player<any>>;
export declare const IconManaSpringTotem: IconInput<Player<any>>;
export declare const IconTotemOfWrath: IconInput<Player<any>>;
export declare const IconWaterShield: IconInput<Player<any>>;
export declare const IconWrathOfAirTotem: IconInput<Player<any>>;
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
            changedEvent: (player: Player<Spec.SpecElementalShaman>) => import("/tbc/core/typed_event").TypedEvent<void>;
            getValue: (player: Player<Spec.SpecElementalShaman>) => RotationType;
            setValue: (player: Player<Spec.SpecElementalShaman>, newValue: number) => void;
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
            changedEvent: (player: Player<Spec.SpecElementalShaman>) => import("/tbc/core/typed_event").TypedEvent<void>;
            getValue: (player: Player<Spec.SpecElementalShaman>) => number;
            setValue: (player: Player<Spec.SpecElementalShaman>, newValue: number) => void;
            enableWhen: (player: Player<Spec.SpecElementalShaman>) => boolean;
            values?: undefined;
        };
    })[];
};
