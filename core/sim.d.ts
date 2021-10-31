import { Buffs } from '/tbc/core/proto/common.js';
import { Enchant } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Gem } from '/tbc/core/proto/common.js';
import { GemColor } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { TypedEvent } from './typed_event.js';
import { WorkerPool } from './worker_pool.js';
export interface SimConfig {
    defaults: {
        phase: number;
        encounter: Encounter;
        buffs: Buffs;
    };
}
export declare class Sim extends WorkerPool {
    readonly phaseChangeEmitter: TypedEvent<void>;
    readonly buffsChangeEmitter: TypedEvent<void>;
    readonly encounterChangeEmitter: TypedEvent<void>;
    readonly numTargetsChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    private _items;
    private _enchants;
    private _gems;
    readonly gearListEmitter: TypedEvent<void>;
    private _phase;
    private _buffs;
    private _encounter;
    private _numTargets;
    private _init;
    constructor(config: SimConfig);
    init(spec: Spec): Promise<void>;
    getItems(slot: ItemSlot | undefined): Array<Item>;
    getEnchants(slot: ItemSlot | undefined): Array<Enchant>;
    getGems(socketColor: GemColor | undefined): Array<Gem>;
    getMatchingGems(socketColor: GemColor): Array<Gem>;
    getPhase(): number;
    setPhase(newPhase: number): void;
    getBuffs(): Buffs;
    setBuffs(newBuffs: Buffs): void;
    getEncounter(): Encounter;
    setEncounter(newEncounter: Encounter): void;
    getNumTargets(): number;
    setNumTargets(newNumTargets: number): void;
    lookupItemSpec(itemSpec: ItemSpec): EquippedItem | null;
    lookupEquipmentSpec(equipSpec: EquipmentSpec): Gear;
    toJson(): Object;
    fromJson(obj: any): void;
}
