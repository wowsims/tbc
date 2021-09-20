import { Sim } from '../sim';
import { Class } from '../api/common';
import { Spec } from '../api/common';
import { specToClass } from '../api/utils';

import { ShamanTalentsPicker } from './shaman';
import { TalentsPicker } from './talents_picker';

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
