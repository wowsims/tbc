import { Debuffs } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Sim } from './sim.js';
import { TypedEvent } from './typed_event.js';
export interface TargetConfig {
    defaults: {
        armor: number;
        mobType: MobType;
        debuffs: Debuffs;
    };
}
export declare class Target {
    readonly armorChangeEmitter: TypedEvent<void>;
    readonly mobTypeChangeEmitter: TypedEvent<void>;
    readonly debuffsChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    private armor;
    private mobType;
    private debuffs;
    private readonly sim;
    constructor(config: TargetConfig, sim: Sim);
    getArmor(): number;
    setArmor(newArmor: number): void;
    getMobType(): MobType;
    setMobType(newMobType: MobType): void;
    getDebuffs(): Debuffs;
    setDebuffs(newDebuffs: Debuffs): void;
    toProto(): TargetProto;
    toJson(): Object;
    fromJson(obj: any): void;
}
