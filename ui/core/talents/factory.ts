import { Sim } from '/tbc/core/sim.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';

import { PriestTalentsPicker } from './priest.js';
import { ShamanTalentsPicker } from './shaman.js';
import { TalentsPicker } from './talents_picker.js';

export function newTalentsPicker<SpecType extends Spec>(spec: Spec, parent: HTMLElement, sim: Sim<SpecType>): TalentsPicker<SpecType> {
  switch (spec) {
    case Spec.SpecElementalShaman:
      return new ShamanTalentsPicker(parent, sim as Sim<Spec.SpecElementalShaman>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecShadowPriest:
      return new PriestTalentsPicker(parent, sim as Sim<Spec.SpecShadowPriest>) as TalentsPicker<SpecType>;
      break;
    default:
      const playerClass = specToClass[spec];
      throw new Error('Unimplemented class talents: ' + playerClass);
  }
}
