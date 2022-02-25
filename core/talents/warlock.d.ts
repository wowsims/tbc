import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { TalentsConfig, TalentsPicker } from './talents_picker.js';
export declare class WarlockTalentsPicker extends TalentsPicker<Spec.SpecWarlock> {
    constructor(parent: HTMLElement, player: Player<Spec.SpecWarlock>);
}
export declare const warlockTalentsConfig: TalentsConfig<Spec.SpecWarlock>;
