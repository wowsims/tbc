import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { TypedEvent } from './typed_event.js';
// Manages all the settings for an Encounter.
export class Encounter {
    constructor(sim) {
        this.duration = 300;
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
            executeProportion: this.executeProportion,
            targets: targetProtos,
        });
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            this.setDuration(eventID, proto.duration);
            this.setExecuteProportion(eventID, proto.executeProportion);
            this.setNumTargets(eventID, proto.targets.length);
            if (proto.targets.length > 0) {
                this.primaryTarget.fromProto(eventID, proto.targets[0]);
            }
        });
    }
    // Returns JSON representing all the current values.
    toJson() {
        return {
            'duration': this.getDuration(),
            'executeProportion': this.getExecuteProportion(),
            'numTargets': this.getNumTargets(),
            'primaryTarget': this.primaryTarget.toJson(),
        };
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(eventID, obj) {
        TypedEvent.freezeAllAndDo(() => {
            const parsedDuration = parseInt(obj['duration']);
            if (!isNaN(parsedDuration) && parsedDuration != 0) {
                this.setDuration(eventID, parsedDuration);
            }
            const parsedExecuteProportion = parseInt(obj['executeProportion']);
            if (!isNaN(parsedExecuteProportion) && parsedExecuteProportion != 0) {
                this.setExecuteProportion(eventID, parsedExecuteProportion);
            }
            const parsedNumTargets = parseInt(obj['numTargets']);
            if (!isNaN(parsedNumTargets) && parsedNumTargets != 0) {
                this.setNumTargets(eventID, parsedNumTargets);
            }
            try {
                this.primaryTarget.fromJson(eventID, obj['primaryTarget']);
            }
            catch (e) {
                console.warn('Failed to parse debuffs: ' + e);
            }
        });
    }
}
