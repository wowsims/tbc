import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
export declare class ProtectionPaladinSimUI extends IndividualSimUI<Spec.SpecProtectionPaladin> {
    constructor(parentElem: HTMLElement, player: Player<Spec.SpecProtectionPaladin>);
}
