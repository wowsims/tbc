import { Debuffs } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import * as Mechanics from '/tbc/core/constants/mechanics.js';
import { TypedEvent } from './typed_event.js';
// Manages all the settings for a single Target.
export class Target {
    constructor(sim) {
        this.level = Mechanics.BOSS_LEVEL;
        this.mobType = MobType.MobTypeDemon;
        this.stats = new Stats();
        this.debuffs = Debuffs.create();
        this.levelChangeEmitter = new TypedEvent();
        this.statsChangeEmitter = new TypedEvent();
        this.mobTypeChangeEmitter = new TypedEvent();
        this.debuffsChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        [
            this.levelChangeEmitter,
            this.statsChangeEmitter,
            this.mobTypeChangeEmitter,
            this.debuffsChangeEmitter,
        ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
    }
    getLevel() {
        return this.level;
    }
    setLevel(eventID, newLevel) {
        if (newLevel == this.level)
            return;
        this.level = newLevel;
        this.levelChangeEmitter.emit(eventID);
    }
    getStats() {
        return this.stats;
    }
    setStats(eventID, newStats) {
        if (newStats.equals(this.stats))
            return;
        this.stats = newStats;
        this.statsChangeEmitter.emit(eventID);
    }
    getMobType() {
        return this.mobType;
    }
    setMobType(eventID, newMobType) {
        if (newMobType == this.mobType)
            return;
        this.mobType = newMobType;
        this.mobTypeChangeEmitter.emit(eventID);
    }
    getDebuffs() {
        // Make a defensive copy
        return Debuffs.clone(this.debuffs);
    }
    setDebuffs(eventID, newDebuffs) {
        if (Debuffs.equals(this.debuffs, newDebuffs))
            return;
        // Make a defensive copy
        this.debuffs = Debuffs.clone(newDebuffs);
        this.debuffsChangeEmitter.emit(eventID);
    }
    toProto() {
        return TargetProto.create({
            level: this.level,
            stats: this.stats.asArray(),
            mobType: this.mobType,
            debuffs: this.debuffs,
        });
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            let stats = new Stats(proto.stats);
            if (proto.armor) {
                stats = stats.withStat(Stat.StatArmor, proto.armor);
            }
            this.setLevel(eventID, proto.level);
            this.setStats(eventID, stats);
            this.setMobType(eventID, proto.mobType);
            this.setDebuffs(eventID, proto.debuffs || Debuffs.create());
        });
    }
}
