import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { Hunter_Rotation as HunterRotation, Hunter_Options as HunterOptions } from '/tbc/core/proto/hunter.js';
export declare const BeastMasteryTalents: {
    name: string;
    data: string;
};
export declare const MarksmanTalents: {
    name: string;
    data: string;
};
export declare const SurvivalTalents: {
    name: string;
    data: string;
};
export declare const DefaultRotation: HunterRotation;
export declare const DefaultOptions: HunterOptions;
export declare const DefaultConsumes: Consumes;
export declare const P1_BM_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P2_BM_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P3_BM_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P1_SV_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P2_SV_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P3_SV_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P4_BM_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P4_SV_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P5_BM_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
export declare const P5_SV_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<any>) => boolean;
    gear: EquipmentSpec;
};
