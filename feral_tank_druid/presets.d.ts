import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { FeralTankDruid_Rotation as DruidRotation, FeralTankDruid_Options as DruidOptions } from '/tbc/core/proto/druid.js';
export declare const StandardTalents: {
    name: string;
    data: string;
};
export declare const DefaultRotation: DruidRotation;
export declare const DefaultOptions: DruidOptions;
export declare const DefaultConsumes: Consumes;
export declare const P4_PRESET: {
    name: string;
    tooltip: string;
    gear: EquipmentSpec;
};