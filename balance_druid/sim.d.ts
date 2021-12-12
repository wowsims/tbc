import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare class BalanceDruidSimUI extends IndividualSimUI<Spec.SpecBalanceDruid> {
    constructor(parentElem: HTMLElement, player: Player<Spec.SpecBalanceDruid>);
}
