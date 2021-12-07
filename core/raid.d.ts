import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { Party } from './party.js';
import { Player } from './player.js';
import { TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';
export declare const MAX_NUM_PARTIES = 5;
export declare class Raid {
    private buffs;
    readonly compChangeEmitter: TypedEvent<void>;
    readonly buffsChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    private parties;
    readonly sim: Sim;
    constructor(sim: Sim);
    size(): number;
    isEmpty(): boolean;
    getParties(): Array<Party>;
    getParty(index: number): Party;
    getPlayers(): Array<Player<any> | null>;
    getPlayer(index: number): Player<any> | null;
    setPlayer(index: number, newPlayer: Player<any> | null): void;
    getBuffs(): RaidBuffs;
    setBuffs(newBuffs: RaidBuffs): void;
    toProto(): RaidProto;
    toJson(): Object;
    fromJson(obj: any): void;
}
