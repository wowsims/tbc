import { Debuffs } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { TypedEvent } from './typed_event.js';
// Manages all the settings for a single Target.
export class Target {
    constructor(sim) {
        this.armor = 0;
        this.mobType = MobType.MobTypeDemon;
        this.debuffs = Debuffs.create();
        this.armorChangeEmitter = new TypedEvent();
        this.mobTypeChangeEmitter = new TypedEvent();
        this.debuffsChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        [
            this.armorChangeEmitter,
            this.mobTypeChangeEmitter,
            this.debuffsChangeEmitter,
        ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
    }
    getArmor() {
        return this.armor;
    }
    setArmor(newArmor) {
        if (newArmor == this.armor)
            return;
        this.armor = newArmor;
        this.armorChangeEmitter.emit();
    }
    getMobType() {
        return this.mobType;
    }
    setMobType(newMobType) {
        if (newMobType == this.mobType)
            return;
        this.mobType = newMobType;
        this.mobTypeChangeEmitter.emit();
    }
    getDebuffs() {
        // Make a defensive copy
        return Debuffs.clone(this.debuffs);
    }
    setDebuffs(newDebuffs) {
        if (Debuffs.equals(this.debuffs, newDebuffs))
            return;
        // Make a defensive copy
        this.debuffs = Debuffs.clone(newDebuffs);
        this.debuffsChangeEmitter.emit();
    }
    toProto() {
        return TargetProto.create({
            armor: this.armor,
            mobType: this.mobType,
            debuffs: this.debuffs,
        });
    }
    fromProto(proto) {
        this.setArmor(proto.armor);
        this.setMobType(proto.mobType);
        this.setDebuffs(proto.debuffs || Debuffs.create());
    }
    toJson() {
        return TargetProto.toJson(this.toProto());
    }
    fromJson(obj) {
        this.fromProto(TargetProto.fromJson(obj));
    }
}
