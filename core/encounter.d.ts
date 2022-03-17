import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';
export declare class Encounter {
    private readonly sim;
    private duration;
    private durationVariation;
    private numTargets;
    private executeProportion;
    readonly primaryTarget: Target;
    readonly durationChangeEmitter: TypedEvent<void>;
    readonly numTargetsChangeEmitter: TypedEvent<void>;
    readonly executeProportionChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    constructor(sim: Sim);
    getDurationVariation(): number;
    setDurationVariation(eventID: EventID, newDuration: number): void;
    getDuration(): number;
    setDuration(eventID: EventID, newDuration: number): void;
    getExecuteProportion(): number;
    setExecuteProportion(eventID: EventID, newExecuteProportion: number): void;
    getNumTargets(): number;
    setNumTargets(eventID: EventID, newNumTargets: number): void;
    toProto(): EncounterProto;
    fromProto(eventID: EventID, proto: EncounterProto): void;
    applyDefaults(eventID: EventID): void;
}
