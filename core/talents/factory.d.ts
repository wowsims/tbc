import { Player } from '/tbc/core/player.js';
import { Spec } from '/tbc/core/proto/common.js';
import { TalentsPicker } from './talents_picker.js';
export declare function newTalentsPicker<SpecType extends Spec>(spec: Spec, parent: HTMLElement, player: Player<SpecType>): TalentsPicker<SpecType>;
