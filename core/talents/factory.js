import { Class } from '/tbc/core/proto/common.js';
import { DruidTalentsPicker } from './druid.js';
import { HunterTalentsPicker } from './hunter.js';
import { MageTalentsPicker } from './mage.js';
import { PaladinTalentsPicker } from './paladin.js';
import { PriestTalentsPicker } from './priest.js';
import { RogueTalentsPicker } from './rogue.js';
import { ShamanTalentsPicker } from './shaman.js';
import { WarlockTalentsPicker } from './warlock.js';
import { WarriorTalentsPicker } from './warrior.js';
export function newTalentsPicker(parent, player) {
    switch (player.getClass()) {
        case Class.ClassDruid:
            return new DruidTalentsPicker(parent, player);
            break;
        case Class.ClassShaman:
            return new ShamanTalentsPicker(parent, player);
            break;
        case Class.ClassHunter:
            return new HunterTalentsPicker(parent, player);
            break;
        case Class.ClassMage:
            return new MageTalentsPicker(parent, player);
            break;
        case Class.ClassPaladin:
            return new PaladinTalentsPicker(parent, player);
            break;
        case Class.ClassRogue:
            return new RogueTalentsPicker(parent, player);
            break;
        case Class.ClassPriest:
            return new PriestTalentsPicker(parent, player);
            break;
        case Class.ClassWarlock:
            return new WarlockTalentsPicker(parent, player);
            break;
        case Class.ClassWarrior:
            return new WarriorTalentsPicker(parent, player);
            break;
        default:
            throw new Error('Unimplemented class talents: ' + player.getClass());
    }
}
