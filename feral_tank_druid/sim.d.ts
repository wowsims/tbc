import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare class FeralTankDruidSimUI extends IndividualSimUI<Spec.SpecFeralTankDruid> {
    constructor(parentElem: HTMLElement, player: Player<Spec.SpecFeralTankDruid>);
}
