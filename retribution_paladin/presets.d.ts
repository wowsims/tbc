import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { RetributionPaladin_Rotation as RetributionPaladinRotation, RetributionPaladin_Options as RetributionPaladinOptions } from '/tbc/core/proto/paladin.js';
export declare const RetributionPaladinTalents: {
    name: string;
    data: string;
};
export declare const DefaultRotation: RetributionPaladinRotation;
export declare const DefaultOptions: RetributionPaladinOptions;
export declare const DefaultConsumes: Consumes;
export declare const P2_PRESET: {
    name: string;
    tooltip: string;
    enableWhen: (player: Player<Spec.SpecRetributionPaladin>) => boolean;
    gear: EquipmentSpec;
};
