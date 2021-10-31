import { Debuffs } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { TypedEvent } from './typed_event.js';
// Manages all the settings for a single Target.
export class Target {
    constructor(config, sim) {
        this.armorChangeEmitter = new TypedEvent();
        this.debuffsChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        this.armor = config.defaults.armor;
        this.debuffs = config.defaults.debuffs;
        [
            this.armorChangeEmitter,
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
            debuffs: this.debuffs,
        });
    }
    // Returns JSON representing all the current values.
    toJson() {
        return {
            'armor': this.armor,
            'debuffs': Debuffs.toJson(this.debuffs),
        };
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(obj) {
        const parsedArmor = parseInt(obj['armor']);
        if (!isNaN(parsedArmor) && parsedArmor != 0) {
            this.armor = parsedArmor;
        }
        try {
            this.setDebuffs(Debuffs.fromJson(obj['debuffs']));
        }
        catch (e) {
            console.warn('Failed to parse debuffs: ' + e);
        }
    }
}
