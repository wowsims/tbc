import { Debuffs } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Sim } from './sim.js';
import { TypedEvent } from './typed_event.js';
export interface TargetConfig {
    defaults: {
        armor: number;
        debuffs: Debuffs;
    };
}
export declare class Target {
    readonly armorChangeEmitter: TypedEvent<void>;
    readonly debuffsChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    private armor;
    private debuffs;
    private readonly sim;
    constructor(config: TargetConfig, sim: Sim);
    getArmor(): number;
    setArmor(newArmor: number): void;
    getDebuffs(): Debuffs;
    setDebuffs(newDebuffs: Debuffs): void;
    toProto(): TargetProto;
    toJson(): Object;
    fromJson(obj: any): void;
}
