import { Sim } from '../sim';
import { Class } from '../api/newapi';
import { Spec } from '../api/newapi';
import { specToClass } from '../api/utils';

import { ShamanTalentsPicker } from './shaman';
import { TalentsPicker } from './talents_picker';

export function newTalentsPicker<ClassType extends Class>(spec: Spec, parent: HTMLElement, sim: Sim<ClassType>): TalentsPicker<ClassType> {
  const playerClass = specToClass[spec];

  switch (playerClass) {
    case Class.ClassShaman:
      return new ShamanTalentsPicker(parent, sim as Sim<Class.ClassShaman>) as TalentsPicker<ClassType>;
      break;
    default:
      throw new Error('Unimplemented class talents: ' + playerClass);
  }
}
