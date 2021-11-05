import { Player } from '/tbc/core/player.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';

import { DruidTalentsPicker } from './druid.js';
import { HunterTalentsPicker } from './hunter.js';
import { MageTalentsPicker } from './mage.js';
import { PaladinTalentsPicker } from './paladin.js';
import { PriestTalentsPicker } from './priest.js';
import { RogueTalentsPicker } from './rogue.js';
import { ShamanTalentsPicker } from './shaman.js';
import { WarlockTalentsPicker } from './warlock.js';
import { WarriorTalentsPicker } from './warrior.js';
import { TalentsPicker } from './talents_picker.js';

export function newTalentsPicker<SpecType extends Spec>(spec: Spec, parent: HTMLElement, player: Player<SpecType>): TalentsPicker<SpecType> {
  switch (spec) {
    case Spec.SpecBalanceDruid:
      return new DruidTalentsPicker(parent, player as Player<Spec.SpecBalanceDruid>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecElementalShaman:
      return new ShamanTalentsPicker(parent, player as Player<Spec.SpecElementalShaman>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecHunter:
      return new HunterTalentsPicker(parent, player as Player<Spec.SpecHunter>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecMage:
      return new MageTalentsPicker(parent, player as Player<Spec.SpecMage>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecRetributionPaladin:
      return new PaladinTalentsPicker(parent, player as Player<Spec.SpecRetributionPaladin>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecRogue:
      return new RogueTalentsPicker(parent, player as Player<Spec.SpecRogue>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecShadowPriest:
      return new PriestTalentsPicker(parent, player as Player<Spec.SpecShadowPriest>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecWarlock:
      return new WarlockTalentsPicker(parent, player as Player<Spec.SpecWarlock>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecWarrior:
      return new WarriorTalentsPicker(parent, player as Player<Spec.SpecWarrior>) as TalentsPicker<SpecType>;
      break;
    default:
      const playerClass = specToClass[spec];
      throw new Error('Unimplemented class talents: ' + playerClass);
  }
}
