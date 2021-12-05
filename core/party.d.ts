import { Player } from './player.js';
import { TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';
export interface PartyConfig {
}
export declare const MAX_PARTY_SIZE = 5;
export declare class Party {
    readonly compChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    private players;
    readonly sim: Sim;
    private readonly playerChangeListener;
    constructor(sim: Sim);
    size(): number;
    empty(): boolean;
    getPlayers(): Array<Player<any> | null>;
    getPlayer(playerIndex: number): Player<any> | null;
    setPlayer(playerIndex: number, newPlayer: Player<any> | null): void;
    toJson(): Object;
    fromJson(obj: any): void;
}
