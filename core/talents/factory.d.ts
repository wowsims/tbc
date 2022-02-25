import { Player } from '/tbc/core/player.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { TalentsPicker } from './talents_picker.js';
export declare function newTalentsPicker<SpecType extends Spec>(parent: HTMLElement, player: Player<SpecType>): TalentsPicker<SpecType>;
export declare function talentSpellIdsToTalentString(playerClass: Class, talentIds: Array<number>): string;
