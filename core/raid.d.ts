import { Party } from './party.js';
import { TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';
export declare const MAX_NUM_PARTIES = 5;
export declare class Raid {
    readonly changeEmitter: TypedEvent<void>;
    private parties;
    private readonly sim;
    constructor(sim: Sim);
    size(): number;
    empty(): boolean;
    getParties(): Array<Party>;
    toJson(): Object;
    fromJson(obj: any): void;
}
