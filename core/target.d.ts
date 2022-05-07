import { Debuffs } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';
export declare class Target {
    private readonly sim;
    private level;
    private mobType;
    private stats;
    private debuffs;
    readonly levelChangeEmitter: TypedEvent<void>;
    readonly statsChangeEmitter: TypedEvent<void>;
    readonly mobTypeChangeEmitter: TypedEvent<void>;
    readonly debuffsChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    constructor(sim: Sim);
    getLevel(): number;
    setLevel(eventID: EventID, newLevel: number): void;
    getStats(): Stats;
    setStats(eventID: EventID, newStats: Stats): void;
    getMobType(): MobType;
    setMobType(eventID: EventID, newMobType: MobType): void;
    getDebuffs(): Debuffs;
    setDebuffs(eventID: EventID, newDebuffs: Debuffs): void;
    toProto(): TargetProto;
    fromProto(eventID: EventID, proto: TargetProto): void;
}
