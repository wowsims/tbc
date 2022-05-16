import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { ProtectionPaladin_Rotation as ProtectionPaladinRotation, ProtectionPaladin_Options as ProtectionPaladinOptions } from '/tbc/core/proto/paladin.js';
export declare const SanctityTalents: {
    name: string;
    data: string;
};
export declare const AvengersShieldTalents: {
    name: string;
    data: string;
};
export declare const DefaultRotation: ProtectionPaladinRotation;
export declare const DefaultOptions: ProtectionPaladinOptions;
export declare const DefaultConsumes: Consumes;
export declare const P1_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecProtectionPaladin>) => boolean;
    gear: EquipmentSpec;
};
export declare const P2_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecProtectionPaladin>) => boolean;
    gear: EquipmentSpec;
};
export declare const P3_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecProtectionPaladin>) => boolean;
    gear: EquipmentSpec;
};
export declare const P4_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecProtectionPaladin>) => boolean;
    gear: EquipmentSpec;
};
export declare const P5_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecProtectionPaladin>) => boolean;
    gear: EquipmentSpec;
};
