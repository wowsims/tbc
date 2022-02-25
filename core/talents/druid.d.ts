import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { TalentsConfig, TalentsPicker } from './talents_picker.js';
export declare class DruidTalentsPicker extends TalentsPicker<Spec.SpecBalanceDruid> {
    constructor(parent: HTMLElement, player: Player<Spec.SpecBalanceDruid>);
}
export declare const druidTalentsConfig: TalentsConfig<Spec.SpecBalanceDruid>;
