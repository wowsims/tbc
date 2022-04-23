import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { Raid } from './raid.js';
import { Player } from './player.js';
import { EventID, TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';
export declare const MAX_PARTY_SIZE = 5;
export declare class Party {
    readonly sim: Sim;
    readonly raid: Raid;
    private buffs;
    readonly compChangeEmitter: TypedEvent<void>;
    readonly buffsChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    private players;
    private readonly playerChangeListener;
    constructor(raid: Raid, sim: Sim);
    size(): number;
    isEmpty(): boolean;
    clear(eventID: EventID): void;
    getIndex(): number;
    getPlayers(): Array<Player<any> | null>;
    getPlayer(playerIndex: number): Player<any> | null;
    setPlayer(eventID: EventID, playerIndex: number, newPlayer: Player<any> | null): void;
    getBuffs(): PartyBuffs;
    setBuffs(eventID: EventID, newBuffs: PartyBuffs): void;
    toProto(forExport?: boolean): PartyProto;
    fromProto(eventID: EventID, proto: PartyProto): void;
}
