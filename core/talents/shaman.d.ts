import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { TalentsPicker } from './talents_picker.js';
export declare class ShamanTalentsPicker extends TalentsPicker<Spec.SpecElementalShaman> {
    constructor(parent: HTMLElement, player: Player<Spec.SpecElementalShaman>);
}