import { Spec } from '../proto/common.js';
import { Sim } from '../sim.js';
import { TalentsPicker } from './talents_picker.js';
export declare class ShamanTalentsPicker extends TalentsPicker<Spec.SpecElementalShaman> {
    constructor(parent: HTMLElement, sim: Sim<Spec.SpecElementalShaman>);
}
