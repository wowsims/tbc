import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { Sim } from './sim.js';
import { TypedEvent } from './typed_event.js';
export declare class Encounter {
    private readonly sim;
    private duration;
    private numTargets;
    readonly primaryTarget: Target;
    readonly durationChangeEmitter: TypedEvent<void>;
    readonly numTargetsChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    constructor(sim: Sim);
    getDuration(): number;
    setDuration(newDuration: number): void;
    getNumTargets(): number;
    setNumTargets(newNumTargets: number): void;
    toProto(): EncounterProto;
    fromProto(proto: EncounterProto): void;
    toJson(): Object;
    fromJson(obj: any): void;
}
