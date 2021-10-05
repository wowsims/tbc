import { Sim } from '../sim.js';
import { Class } from '../proto/common.js';
import { Spec } from '../proto/common.js';
import { specToClass } from '../api/utils.js';

import { ShamanTalentsPicker } from './shaman.js';
import { TalentsPicker } from './talents_picker.js';

export function newTalentsPicker<SpecType extends Spec>(spec: Spec, parent: HTMLElement, sim: Sim<SpecType>): TalentsPicker<SpecType> {
  switch (spec) {
    case Spec.SpecElementalShaman:
      return new ShamanTalentsPicker(parent, sim as Sim<Spec.SpecElementalShaman>) as TalentsPicker<SpecType>;
      break;
    default:
      const playerClass = specToClass[spec];
      throw new Error('Unimplemented class talents: ' + playerClass);
  }
}
