import { Player } from './player.js';
import { TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';
export interface PartyConfig {
}
export declare const MAX_PARTY_SIZE = 5;
export declare class Party {
    readonly changeEmitter: TypedEvent<void>;
    private players;
    private readonly sim;
    constructor(sim: Sim);
    size(): number;
    empty(): boolean;
    getPlayers(): Array<Player<any>>;
    addPlayer(player: Player<any>): void;
    removePlayer(playerToRemove: Player<any>): void;
    toJson(): Object;
    fromJson(obj: any): void;
}
