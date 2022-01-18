import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { Warrior_Rotation as WarriorRotation, Warrior_Options as WarriorOptions } from '/tbc/core/proto/warrior.js';
export declare const ArmsSlamTalents: {
    name: string;
    data: string;
};
export declare const ArmsDWTalents: {
    name: string;
    data: string;
};
export declare const FuryTalents: {
    name: string;
    data: string;
};
export declare const DefaultFuryRotation: WarriorRotation;
export declare const DefaultFuryOptions: WarriorOptions;
export declare const DefaultFuryConsumes: Consumes;
export declare const DefaultArmsSlamRotation: WarriorRotation;
export declare const DefaultArmsSlamOptions: WarriorOptions;
export declare const DefaultArmsSlamConsumes: Consumes;
export declare const DefaultArmsDWRotation: WarriorRotation;
export declare const DefaultArmsDWOptions: WarriorOptions;
export declare const DefaultArmsDWConsumes: Consumes;
export declare const P1_FURY_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P1_ARMSSLAM_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P1_ARMSDW_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
