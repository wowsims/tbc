import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { TypedEvent } from './typed_event.js';
// Manages all the settings for an Encounter.
export class Encounter {
    constructor(sim) {
        this.duration = 180;
        this.durationVariation = 5;
        this.executeProportion = 0.2;
        this.targetsChangeEmitter = new TypedEvent();
        this.durationChangeEmitter = new TypedEvent();
        this.executeProportionChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        this.targets = [Target.fromDefaults(TypedEvent.nextEventID(), sim)];
        [
            this.targetsChangeEmitter,
            this.durationChangeEmitter,
            this.executeProportionChangeEmitter,
        ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
    }
    get primaryTarget() {
        return this.targets[0];
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
        return this.targets.length;
    }
    getTargets() {
        return this.targets.slice();
    }
    setTargets(eventID, newTargets) {
        TypedEvent.freezeAllAndDo(() => {
            if (newTargets.length == 0) {
                newTargets = [Target.fromDefaults(eventID, this.sim)];
            }
            if (newTargets.length == this.targets.length && newTargets.every((target, i) => TargetProto.equals(target.toProto(), this.targets[i].toProto()))) {
                return;
            }
            this.targets = newTargets;
            this.targetsChangeEmitter.emit(eventID);
        });
    }
    matchesPreset(preset) {
        return preset.targets.length == this.targets.length && this.targets.every((t, i) => t.matchesPreset(preset.targets[i]));
    }
    applyPreset(eventID, preset) {
        TypedEvent.freezeAllAndDo(() => {
            let newTargets = this.targets.slice(0, preset.targets.length);
            while (newTargets.length < preset.targets.length) {
                newTargets.push(new Target(this.sim));
            }
            newTargets.forEach((nt, i) => nt.applyPreset(eventID, preset.targets[i]));
            this.setTargets(eventID, newTargets);
        });
    }
    toProto() {
        return EncounterProto.create({
            duration: this.duration,
            durationVariation: this.durationVariation,
            executeProportion: this.executeProportion,
            targets: this.targets.map(target => target.toProto()),
        });
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            this.setDuration(eventID, proto.duration);
            this.setDurationVariation(eventID, proto.durationVariation);
            this.setExecuteProportion(eventID, proto.executeProportion);
            if (proto.targets.length > 0) {
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
