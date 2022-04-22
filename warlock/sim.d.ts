import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare class WarlockSimUI extends IndividualSimUI<Spec.SpecWarlock> {
    constructor(parentElem: HTMLElement, player: Player<Spec.SpecWarlock>);
}
