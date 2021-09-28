import { Sim, SimConfig } from './sim.js';
import { TypedEvent } from './typed_event.js';
import { Spec } from './api/common.js';
export interface SimUIConfig<SpecType extends Spec> extends SimConfig<SpecType> {
}
export declare abstract class SimUI<SpecType extends Spec> {
    readonly parentElem: HTMLElement;
    readonly sim: Sim<SpecType>;
    private readonly exclusivityMap;
    constructor(parentElem: HTMLElement, config: SimUIConfig<SpecType>);
    init(): Promise<void>;
    registerExclusiveEffect(effect: ExclusiveEffect): void;
}
export declare type ExclusivityTag = 'Battle Elixir' | 'Drums' | 'Food' | 'Guardian Elixir' | 'Potion' | 'Rune' | 'Weapon Imbue';
export interface ExclusiveEffect {
    tags: Array<ExclusivityTag>;
    changedEvent: TypedEvent<any>;
    isActive: () => boolean;
    deactivate: () => void;
}
