import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EncounterType } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { TypedEvent } from './typed_event.js';
// Manages all the settings for an Encounter.
export class Encounter {
    constructor(sim) {
        this.type = EncounterType.EncounterTypeSimple;
        this.duration = 180;
        this.durationVariation = 5;
        this.numTargets = 1;
        this.executeProportion = 0.2;
        this.targetsChangeEmitter = new TypedEvent();
        this.typeChangeEmitter = new TypedEvent();
        this.durationChangeEmitter = new TypedEvent();
        this.numTargetsChangeEmitter = new TypedEvent();
        this.executeProportionChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        this.targets = [Target.fromDefaults(TypedEvent.nextEventID(), sim)];
        [
            this.targetsChangeEmitter,
            this.typeChangeEmitter,
            this.durationChangeEmitter,
            this.numTargetsChangeEmitter,
            this.executeProportionChangeEmitter,
        ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
    }
    get primaryTarget() {
        return this.targets[0];
    }
    getType() {
        return this.type;
    }
    setType(eventID, newType) {
        if (newType == this.type)
            return;
        this.type = newType;
        this.typeChangeEmitter.emit(eventID);
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
    getTargets() {
        return this.targets.slice();
    }
    setTargets(eventID, newTargets) {
        if (newTargets.length == 0) {
            newTargets = [Target.fromDefaults(eventID, this.sim)];
        }
        if (newTargets.length == this.targets.length && newTargets.every((target, i) => TargetProto.equals(target.toProto(), this.targets[i].toProto()))) {
            return;
        }
        TypedEvent.freezeAllAndDo(() => {
            if (newTargets.length != this.targets.length) {
                this.numTargets = newTargets.length;
                this.numTargetsChangeEmitter.emit(eventID);
            }
            this.targets = newTargets;
            this.targetsChangeEmitter.emit(eventID);
        });
    }
    toProto() {
        const numTargets = Math.max(1, this.numTargets);
        let targetProtos = [];
        if (this.getType() == EncounterType.EncounterTypeSimple) {
            for (let i = 0; i < numTargets; i++) {
                targetProtos.push(this.primaryTarget.toProto());
            }
        }
        else {
            targetProtos = this.targets.map(target => target.toProto());
        }
        return EncounterProto.create({
            type: this.type,
            duration: this.duration,
            durationVariation: this.durationVariation,
            executeProportion: this.executeProportion,
            targets: targetProtos,
        });
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            this.setType(eventID, proto.type);
            this.setDuration(eventID, proto.duration);
            this.setDurationVariation(eventID, proto.durationVariation);
            this.setExecuteProportion(eventID, proto.executeProportion);
            this.setNumTargets(eventID, Math.max(1, proto.targets.length));
            if (proto.targets.length > 0) {
                this.primaryTarget.fromProto(eventID, proto.targets[0]);
                this.setTargets(eventID, proto.targets.map(targetProto => {
                    const target = new Target(this.sim);
                    target.fromProto(eventID, targetProto);
                    return target;
                }));
            }
            else {
                this.setTargets(eventID, [Target.fromDefaults(eventID, this.sim)]);
            }
        });
    }
    applyDefaults(eventID) {
        this.fromProto(eventID, EncounterProto.create({
            duration: 180,
            durationVariation: 5,
            executeProportion: 0.2,
            targets: [Target.defaultProto()],
        }));
    }
}
