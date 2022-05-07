import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { TypedEvent } from './typed_event.js';
// Manages all the settings for an Encounter.
export class Encounter {
    constructor(sim) {
        this.duration = 180;
        this.durationVariation = 5;
        this.numTargets = 1;
        this.executeProportion = 0.2;
        this.durationChangeEmitter = new TypedEvent();
        this.numTargetsChangeEmitter = new TypedEvent();
        this.executeProportionChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        this.primaryTarget = new Target(sim);
        [
            this.durationChangeEmitter,
            this.numTargetsChangeEmitter,
            this.executeProportionChangeEmitter,
            this.primaryTarget.changeEmitter,
        ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
    }
    getDurationVariation() {
        return this.durationVariation;
    }
    setDurationVariation(eventID, newDuration) {
        if (newDuration == this.durationVariation)
            return;
        this.durationVariation = newDuration;
        this.durationChangeEmitter.emit(eventID);
    }
    getDuration() {
        return this.duration;
    }
    setDuration(eventID, newDuration) {
        if (newDuration == this.duration)
            return;
        this.duration = newDuration;
        this.durationChangeEmitter.emit(eventID);
    }
    getExecuteProportion() {
        return this.executeProportion;
    }
    setExecuteProportion(eventID, newExecuteProportion) {
        if (newExecuteProportion == this.executeProportion)
            return;
        this.executeProportion = newExecuteProportion;
        this.executeProportionChangeEmitter.emit(eventID);
    }
    getNumTargets() {
        return this.numTargets;
    }
    setNumTargets(eventID, newNumTargets) {
        if (newNumTargets == this.numTargets)
            return;
        this.numTargets = newNumTargets;
        this.numTargetsChangeEmitter.emit(eventID);
    }
    toProto() {
        const numTargets = Math.max(1, this.numTargets);
        const targetProtos = [];
        for (let i = 0; i < numTargets; i++) {
            targetProtos.push(this.primaryTarget.toProto());
        }
        return EncounterProto.create({
            duration: this.duration,
            durationVariation: this.durationVariation,
            executeProportion: this.executeProportion,
            targets: targetProtos,
        });
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            this.setDuration(eventID, proto.duration);
            this.setDurationVariation(eventID, proto.durationVariation);
            this.setExecuteProportion(eventID, proto.executeProportion);
            this.setNumTargets(eventID, proto.targets.length);
            if (proto.targets.length > 0) {
                this.primaryTarget.fromProto(eventID, proto.targets[0]);
            }
        });
    }
    applyDefaults(eventID) {
        this.fromProto(eventID, EncounterProto.create({
            duration: 180,
            durationVariation: 5,
            executeProportion: 0.2,
            targets: [TargetProto.create({
                    level: 73,
                    stats: new Stats().withStat(Stat.StatArmor, 7684).asArray(),
                    mobType: MobType.MobTypeDemon,
                })],
        }));
    }
}
