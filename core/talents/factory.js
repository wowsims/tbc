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
export function newTalentsPicker(spec, parent, player) {
    switch (spec) {
        case Spec.SpecBalanceDruid:
            return new DruidTalentsPicker(parent, player);
            break;
        case Spec.SpecElementalShaman:
            return new ShamanTalentsPicker(parent, player);
            break;
        case Spec.SpecEnhancementShaman:
            return new ShamanTalentsPicker(parent, player);
            break;
        case Spec.SpecHunter:
            return new HunterTalentsPicker(parent, player);
            break;
        case Spec.SpecMage:
            return new MageTalentsPicker(parent, player);
            break;
        case Spec.SpecRetributionPaladin:
            return new PaladinTalentsPicker(parent, player);
            break;
        case Spec.SpecRogue:
            return new RogueTalentsPicker(parent, player);
            break;
        case Spec.SpecShadowPriest:
            return new PriestTalentsPicker(parent, player);
            break;
        case Spec.SpecWarlock:
            return new WarlockTalentsPicker(parent, player);
            break;
        case Spec.SpecWarrior:
            return new WarriorTalentsPicker(parent, player);
            break;
        default:
            const playerClass = specToClass[spec];
            throw new Error('Unimplemented class talents: ' + playerClass);
    }
}
