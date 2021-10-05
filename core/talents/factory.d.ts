import { Sim } from '../sim.js';
import { Spec } from '../proto/common.js';
import { TalentsPicker } from './talents_picker.js';
export declare function newTalentsPicker<SpecType extends Spec>(spec: Spec, parent: HTMLElement, sim: Sim<SpecType>): TalentsPicker<SpecType>;
