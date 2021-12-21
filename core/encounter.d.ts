import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';
export declare class Encounter {
    private readonly sim;
    private duration;
    private numTargets;
    readonly primaryTarget: Target;
    readonly durationChangeEmitter: TypedEvent<void>;
    readonly numTargetsChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    private modifyEncounterProto;
    constructor(sim: Sim);
    getDuration(): number;
    setDuration(eventID: EventID, newDuration: number): void;
    getNumTargets(): number;
    setNumTargets(eventID: EventID, newNumTargets: number): void;
    setModifyEncounterProto(newModFn: (encounterProto: EncounterProto) => void): void;
    toProto(): EncounterProto;
    fromProto(eventID: EventID, proto: EncounterProto): void;
    toJson(): Object;
    fromJson(eventID: EventID, obj: any): void;
}
