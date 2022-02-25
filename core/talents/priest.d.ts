import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { TalentsConfig, TalentsPicker } from './talents_picker.js';
export declare class PriestTalentsPicker extends TalentsPicker<Spec.SpecShadowPriest> {
    constructor(parent: HTMLElement, player: Player<Spec.SpecShadowPriest>);
}
export declare const priestTalentsConfig: TalentsConfig<Spec.SpecShadowPriest>;
