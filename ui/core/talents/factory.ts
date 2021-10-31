import { Player } from '/tbc/core/player.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';

import { ShamanTalentsPicker } from './shaman.js';
import { TalentsPicker } from './talents_picker.js';

export function newTalentsPicker<SpecType extends Spec>(spec: Spec, parent: HTMLElement, player: Player<SpecType>): TalentsPicker<SpecType> {
  switch (spec) {
    case Spec.SpecElementalShaman:
      return new ShamanTalentsPicker(parent, player as Player<Spec.SpecElementalShaman>) as TalentsPicker<SpecType>;
      break;
    default:
      const playerClass = specToClass[spec];
      throw new Error('Unimplemented class talents: ' + playerClass);
  }
}
