import { IndividualSimRequest } from '/tbc/core/proto/api.js';
import { RaidSimRequest } from '/tbc/core/proto/api.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { SpecOptions } from '/tbc/core/proto_utils/utils.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { Party } from './party.js';
import { Player } from './player.js';
import { Raid } from './raid.js';
import { Sim } from './sim.js';
import { Encounter } from './encounter.js';
import { TypedEvent } from './typed_event.js';
export declare type ReleaseStatus = 'Alpha' | 'Beta' | 'Live';
export interface SimUIConfig<SpecType extends Spec> {
    spec: Spec;
    releaseStatus: ReleaseStatus;
    knownIssues?: Array<string>;
    defaults: {
        gear: EquipmentSpec;
        epWeights: Stats;
        consumes: Consumes;
        rotation: SpecRotation<SpecType>;
        talents: string;
        specOptions: SpecOptions<SpecType>;
        raidBuffs: RaidBuffs;
        partyBuffs: PartyBuffs;
        individualBuffs: IndividualBuffs;
        debuffs: Debuffs;
    };
}
export declare abstract class SimUI<SpecType extends Spec> {
    readonly parentElem: HTMLElement;
    readonly sim: Sim;
    readonly raid: Raid;
    readonly party: Party;
    readonly player: Player<SpecType>;
    readonly encounter: Encounter;
    readonly simUiConfig: SimUIConfig<SpecType>;
    readonly changeEmitter: TypedEvent<void>;
    private readonly exclusivityMap;
    constructor(parentElem: HTMLElement, config: SimUIConfig<SpecType>);
    toJson(): Object;
    fromJson(obj: any): void;
    init(): Promise<void>;
    applyDefaults(): void;
    registerExclusiveEffect(effect: ExclusiveEffect): void;
    getSavedGearStorageKey(): string;
    getSavedEncounterStorageKey(): string;
    getSavedRotationStorageKey(): string;
    getSavedSettingsStorageKey(): string;
    getSavedTalentsStorageKey(): string;
    private getStorageKey;
    makeRaidSimRequest(iterations: number, debug: boolean): RaidSimRequest;
    makeCurrentIndividualSimRequest(iterations: number, debug: boolean): IndividualSimRequest;
}
export declare type ExclusivityTag = 'Battle Elixir' | 'Drums' | 'Food' | 'Alchohol' | 'Guardian Elixir' | 'Potion' | 'Rune' | 'Weapon Imbue';
export interface ExclusiveEffect {
    tags: Array<ExclusivityTag>;
    changedEvent: TypedEvent<any>;
    isActive: () => boolean;
    deactivate: () => void;
}
