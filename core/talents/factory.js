import { Class } from '/tbc/core/proto/common.js';
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
            const talentSpellIds = [];
            let lastSeenId = -1;
            for (let i = 0; i < talentConfig.maxPoints; i++) {
                let curId = i < talentConfig.spellIds.length ? talentConfig.spellIds[i] : lastSeenId;
                if (curId == lastSeenId) {
                    curId++;
                }
                talentSpellIds.push(curId);
                lastSeenId = curId;
            }
            const spellIdIndex = talentSpellIds.findIndex(spellId => talentIds.includes(spellId));
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
