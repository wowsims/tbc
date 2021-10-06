import { Spec } from '/tbc/core/proto/common.js';
import { Sim } from '/tbc/core/sim.js';
import { TalentsPicker } from './talents_picker.js';
export declare class ShamanTalentsPicker extends TalentsPicker<Spec.SpecElementalShaman> {
    constructor(parent: HTMLElement, sim: Sim<Spec.SpecElementalShaman>);
}
