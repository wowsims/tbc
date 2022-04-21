import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { ElementalShaman_Rotation as ElementalShamanRotation, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';
export declare const StandardTalents: {
    name: string;
    data: string;
};
export declare const DefaultRotation: ElementalShamanRotation;
export declare const DefaultOptions: ElementalShamanOptions;
export declare const DefaultConsumes: Consumes;
export declare const P1_PRESET: {
    name: string;
    tooltip: string;
    gear: EquipmentSpec;
};
export declare const P2_PRESET: {
    name: string;
    tooltip: string;
    gear: EquipmentSpec;
};
export declare const P3_PRESET: {
    name: string;
    tooltip: string;
    gear: EquipmentSpec;
};
export declare const P4_PRESET: {
    name: string;
    tooltip: string;
    gear: EquipmentSpec;
};
export declare const P5_ALLIANCE_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P5_HORDE_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
