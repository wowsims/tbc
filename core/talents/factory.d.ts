import { Sim } from '/tbc/core/sim.js';
import { Spec } from '/tbc/core/proto/common.js';
import { TalentsPicker } from './talents_picker.js';
export declare function newTalentsPicker<SpecType extends Spec>(spec: Spec, parent: HTMLElement, sim: Sim<SpecType>): TalentsPicker<SpecType>;
