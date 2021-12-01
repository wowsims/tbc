import { Debuffs } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Sim } from './sim.js';
import { TypedEvent } from './typed_event.js';
export declare class Target {
    private readonly sim;
    private armor;
    private mobType;
    private debuffs;
    readonly armorChangeEmitter: TypedEvent<void>;
    readonly mobTypeChangeEmitter: TypedEvent<void>;
    readonly debuffsChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    constructor(sim: Sim);
    getArmor(): number;
    setArmor(newArmor: number): void;
    getMobType(): MobType;
    setMobType(newMobType: MobType): void;
    getDebuffs(): Debuffs;
    setDebuffs(newDebuffs: Debuffs): void;
    toProto(): TargetProto;
    fromProto(proto: TargetProto): void;
    toJson(): Object;
    fromJson(obj: any): void;
}
