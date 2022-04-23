import { Class } from '/tbc/core/proto/common.js';
import { specToClass, specTypeFunctions, } from '/tbc/core/proto_utils/utils.js';
import { druidTalentsConfig, DruidTalentsPicker } from './druid.js';
import { hunterTalentsConfig, HunterTalentsPicker } from './hunter.js';
import { mageTalentsConfig, MageTalentsPicker } from './mage.js';
import { paladinTalentsConfig, PaladinTalentsPicker } from './paladin.js';
import { priestTalentsConfig, PriestTalentsPicker } from './priest.js';
import { rogueTalentsConfig, RogueTalentsPicker } from './rogue.js';
import { shamanTalentsConfig, ShamanTalentsPicker } from './shaman.js';
import { warlockTalentsConfig, WarlockTalentsPicker } from './warlock.js';
import { warriorTalentsConfig, WarriorTalentsPicker } from './warrior.js';
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
const classTalentsConfig = {
    [Class.ClassUnknown]: [],
    [Class.ClassDruid]: druidTalentsConfig,
    [Class.ClassShaman]: shamanTalentsConfig,
    [Class.ClassHunter]: hunterTalentsConfig,
    [Class.ClassMage]: mageTalentsConfig,
    [Class.ClassRogue]: rogueTalentsConfig,
    [Class.ClassPaladin]: paladinTalentsConfig,
    [Class.ClassPriest]: priestTalentsConfig,
    [Class.ClassWarlock]: warlockTalentsConfig,
    [Class.ClassWarrior]: warriorTalentsConfig,
};
export function talentSpellIdsToTalentString(playerClass, talentIds) {
    const talentsConfig = classTalentsConfig[playerClass];
    const talentsStr = talentsConfig.map(treeConfig => {
        const treeStr = treeConfig.talents.map(talentConfig => {
            const spellIdIndex = talentConfig.spellIds.findIndex(spellId => talentIds.includes(spellId));
            if (spellIdIndex == -1) {
                return '0';
            }
            else {
                return String(spellIdIndex + 1);
            }
        }).join('').replace(/0+$/g, '');
        return treeStr;
    }).join('-').replace(/-+$/g, '');
    return talentsStr;
}
export function talentStringToProto(spec, talentString) {
    const talentsConfig = classTalentsConfig[specToClass[spec]];
    const specFunctions = specTypeFunctions[spec];
    const proto = specFunctions.talentsCreate();
    talentString.split('-').forEach((treeString, treeIdx) => {
        const treeConfig = talentsConfig[treeIdx];
        [...treeString].forEach((talentString, i) => {
            const talentConfig = treeConfig.talents[i];
            const points = parseInt(talentString);
            if (talentConfig.fieldName) {
                if (talentConfig.maxPoints == 1) {
                    proto[talentConfig.fieldName] = points == 1;
                }
                else {
                    proto[talentConfig.fieldName] = points;
                }
            }
        });
    });
    return proto;
}
