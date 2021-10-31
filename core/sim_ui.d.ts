import { IndividualSimRequest } from '/tbc/core/proto/api.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player, PlayerConfig } from './player.js';
import { Sim, SimConfig } from './sim.js';
import { Target, TargetConfig } from './target.js';
import { TypedEvent } from './typed_event.js';
export declare type ReleaseStatus = 'Alpha' | 'Beta' | 'Live';
export interface SimUIConfig<SpecType extends Spec> {
    releaseStatus: ReleaseStatus;
    knownIssues?: Array<string>;
    sim: SimConfig;
    player: PlayerConfig<SpecType>;
    target: TargetConfig;
}
export declare abstract class SimUI<SpecType extends Spec> {
    readonly parentElem: HTMLElement;
    readonly sim: Sim;
    readonly player: Player<SpecType>;
    readonly target: Target;
    readonly simUiConfig: SimUIConfig<SpecType>;
    readonly changeEmitter: TypedEvent<void>;
    private readonly exclusivityMap;
    constructor(parentElem: HTMLElement, config: SimUIConfig<SpecType>);
    toJson(): Object;
    fromJson(obj: any): void;
    init(): Promise<void>;
    registerExclusiveEffect(effect: ExclusiveEffect): void;
    getSavedGearStorageKey(): string;
    getSavedEncounterStorageKey(): string;
    getSavedRotationStorageKey(): string;
    getSavedSettingsStorageKey(): string;
    getSavedTalentsStorageKey(): string;
    private getStorageKey;
    makeCurrentIndividualSimRequest(iterations: number, debug: boolean): IndividualSimRequest;
}
export declare type ExclusivityTag = 'Battle Elixir' | 'Drums' | 'Food' | 'Guardian Elixir' | 'Potion' | 'Rune' | 'Weapon Imbue';
export interface ExclusiveEffect {
    tags: Array<ExclusivityTag>;
    changedEvent: TypedEvent<any>;
    isActive: () => boolean;
    deactivate: () => void;
}
