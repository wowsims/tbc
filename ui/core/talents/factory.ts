import { Sim } from '../sim';
import { Class } from '../api/newapi';
import { Spec } from '../api/newapi';
import { SpecToClass } from '../api/utils';

import { ShamanTalentsPicker } from './shaman';
import { TalentsPicker } from './talents_picker';

export function newTalentsPicker(spec: Spec, parent: HTMLElement, sim: Sim): TalentsPicker {
  const playerClass = SpecToClass[spec];

  switch (playerClass) {
    case Class.ClassShaman:
      return new ShamanTalentsPicker(parent, sim);
      break;
    default:
      throw new Error('Unimplemented class talents: ' + playerClass);
  }
}
