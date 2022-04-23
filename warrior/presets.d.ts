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
export declare const DefaultRotation: WarriorRotation;
export declare const ArmsRotation: WarriorRotation;
export declare const DefaultOptions: WarriorOptions;
export declare const DefaultConsumes: Consumes;
export declare const P1_FURY_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P2_FURY_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P3_FURY_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P4_FURY_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P5_FURY_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P1_ARMS_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P2_ARMS_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P3_ARMS_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P4_ARMS_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
export declare const P5_ARMS_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecWarrior>) => boolean;
    gear: EquipmentSpec;
};
