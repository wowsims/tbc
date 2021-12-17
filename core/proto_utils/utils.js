import { getEnumValues } from '/tbc/core/utils.js';
import { intersection } from '/tbc/core/utils.js';
import { maxIndex } from '/tbc/core/utils.js';
import { sum } from '/tbc/core/utils.js';
import { Player } from '/tbc/core/proto/api.js';
import { ArmorType } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { Enchant } from '/tbc/core/proto/common.js';
import { GemColor } from '/tbc/core/proto/common.js';
import { HandType } from '/tbc/core/proto/common.js';
import { ItemCategory } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemType } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { RangedWeaponType } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { WeaponType } from '/tbc/core/proto/common.js';
import * as Gems from '/tbc/core/constants/gems.js';
import { BalanceDruid, BalanceDruid_Rotation as BalanceDruidRotation, DruidTalents, BalanceDruid_Options as BalanceDruidOptions } from '/tbc/core/proto/druid.js';
import { ElementalShaman, ElementalShaman_Rotation as ElementalShamanRotation, ShamanTalents, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';
import { Hunter, Hunter_Rotation as HunterRotation, HunterTalents, Hunter_Options as HunterOptions } from '/tbc/core/proto/hunter.js';
import { Mage, Mage_Rotation as MageRotation, MageTalents, Mage_Options as MageOptions } from '/tbc/core/proto/mage.js';
import { Rogue, Rogue_Rotation as RogueRotation, RogueTalents, Rogue_Options as RogueOptions } from '/tbc/core/proto/rogue.js';
import { RetributionPaladin, RetributionPaladin_Rotation as RetributionPaladinRotation, PaladinTalents, RetributionPaladin_Options as RetributionPaladinOptions } from '/tbc/core/proto/paladin.js';
import { ShadowPriest, ShadowPriest_Rotation as ShadowPriestRotation, PriestTalents, ShadowPriest_Options as ShadowPriestOptions } from '/tbc/core/proto/priest.js';
import { Warlock, Warlock_Rotation as WarlockRotation, WarlockTalents, Warlock_Options as WarlockOptions } from '/tbc/core/proto/warlock.js';
import { Warrior, Warrior_Rotation as WarriorRotation, WarriorTalents, Warrior_Options as WarriorOptions } from '/tbc/core/proto/warrior.js';
export const specNames = {
    [Spec.SpecBalanceDruid]: 'Balance Druid',
    [Spec.SpecElementalShaman]: 'Elemental Shaman',
    [Spec.SpecHunter]: 'Hunter',
    [Spec.SpecMage]: 'Mage',
    [Spec.SpecRogue]: 'Rogue',
    [Spec.SpecRetributionPaladin]: 'Retribution Paladin',
    [Spec.SpecShadowPriest]: 'Shadow Priest',
    [Spec.SpecWarlock]: 'Warlock',
    [Spec.SpecWarrior]: 'Warrior',
};
export const classColors = {
    [Class.ClassUnknown]: '#fff',
    [Class.ClassDruid]: '#ff7d0a',
    [Class.ClassHunter]: '#abd473',
    [Class.ClassMage]: '#69ccf0',
    [Class.ClassPaladin]: '#f58cba',
    [Class.ClassPriest]: '#fff',
    [Class.ClassRogue]: '#fff569',
    [Class.ClassShaman]: '#2459ff',
    [Class.ClassWarlock]: '#9482c9',
    [Class.ClassWarrior]: '#c79c6e',
};
export const specIconsLarge = {
    [Spec.SpecBalanceDruid]: 'https://wow.zamimg.com/images/wow/icons/large/spell_nature_starfall.jpg',
    [Spec.SpecElementalShaman]: 'https://wow.zamimg.com/images/wow/icons/large/spell_nature_lightning.jpg',
    [Spec.SpecHunter]: 'https://wow.zamimg.com/images/wow/icons/large/ability_marksmanship.jpg',
    [Spec.SpecMage]: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_magicalsentry.jpg',
    [Spec.SpecRogue]: 'https://wow.zamimg.com/images/wow/icons/large/ability_rogue_eviscerate.jpg',
    [Spec.SpecRetributionPaladin]: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_auraoflight.jpg',
    [Spec.SpecShadowPriest]: 'https://wow.zamimg.com/images/wow/icons/large/spell_shadow_shadowwordpain.jpg',
    [Spec.SpecWarlock]: 'https://wow.zamimg.com/images/wow/icons/large/spell_shadow_metamorphosis.jpg',
    [Spec.SpecWarrior]: 'https://wow.zamimg.com/images/wow/icons/large/ability_warrior_innerrage.jpg',
};
export const talentTreeIcons = {
    [Class.ClassUnknown]: [],
    [Class.ClassDruid]: [
        'https://wow.zamimg.com/images/wow/icons/medium/spell_nature_starfall.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/ability_racial_bearform.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_nature_healingtouch.jpg',
    ],
    [Class.ClassHunter]: [
        'https://wow.zamimg.com/images/wow/icons/medium/ability_hunter_beasttaming.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/ability_marksmanship.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/ability_hunter_swiftstrike.jpg',
    ],
    [Class.ClassMage]: [
        'https://wow.zamimg.com/images/wow/icons/medium/spell_holy_magicalsentry.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_fire_firebolt02.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_frost_frostbolt02.jpg',
    ],
    [Class.ClassPaladin]: [
        'https://wow.zamimg.com/images/wow/icons/medium/spell_holy_holybolt.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_holy_devotionaura.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_holy_auraoflight.jpg',
    ],
    [Class.ClassPriest]: [
        'https://wow.zamimg.com/images/wow/icons/medium/spell_holy_powerinfusion.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_holy_holybolt.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_shadow_shadowwordpain.jpg',
    ],
    [Class.ClassRogue]: [
        'https://wow.zamimg.com/images/wow/icons/medium/ability_rogue_eviscerate.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/ability_backstab.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/ability_stealth.jpg',
    ],
    [Class.ClassShaman]: [
        'https://wow.zamimg.com/images/wow/icons/medium/spell_nature_lightning.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/ability_shaman_stormstrike.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_nature_magicimmunity.jpg',
    ],
    [Class.ClassWarlock]: [
        'https://wow.zamimg.com/images/wow/icons/medium/spell_shadow_deathcoil.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_shadow_metamorphosis.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/spell_shadow_rainoffire.jpg',
    ],
    [Class.ClassWarrior]: [
        'https://wow.zamimg.com/images/wow/icons/medium/ability_warrior_savageblow.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/ability_warrior_innerrage.jpg',
        'https://wow.zamimg.com/images/wow/icons/medium/inv_shield_06.jpg',
    ],
};
// Returns the index of the talent tree (0, 1, or 2) that has the most points.
export function getTalentTree(talentsString) {
    const trees = talentsString.split('-');
    const points = trees.map(tree => sum([...tree].map(char => parseInt(char))));
    return maxIndex(points) || 0;
}
// Returns the index of the talent tree (0, 1, or 2) that has the most points.
export function getTalentTreeIcon(spec, talentsString) {
    const talentTreeIdx = getTalentTree(talentsString);
    return talentTreeIcons[specToClass[spec]][talentTreeIdx];
}
export const specTypeFunctions = {
    [Spec.SpecBalanceDruid]: {
        rotationCreate: () => BalanceDruidRotation.create(),
        rotationEquals: (a, b) => BalanceDruidRotation.equals(a, b),
        rotationCopy: (a) => BalanceDruidRotation.clone(a),
        rotationToJson: (a) => BalanceDruidRotation.toJson(a),
        rotationFromJson: (obj) => BalanceDruidRotation.fromJson(obj),
        rotationFromPlayer: (player) => player.spec.oneofKind == 'balanceDruid'
            ? player.spec.balanceDruid.rotation || BalanceDruidRotation.create()
            : BalanceDruidRotation.create(),
        talentsCreate: () => DruidTalents.create(),
        talentsEquals: (a, b) => DruidTalents.equals(a, b),
        talentsCopy: (a) => DruidTalents.clone(a),
        talentsToJson: (a) => DruidTalents.toJson(a),
        talentsFromJson: (obj) => DruidTalents.fromJson(obj),
        talentsFromPlayer: (player) => player.spec.oneofKind == 'balanceDruid'
            ? player.spec.balanceDruid.talents || DruidTalents.create()
            : DruidTalents.create(),
        optionsCreate: () => BalanceDruidOptions.create(),
        optionsEquals: (a, b) => BalanceDruidOptions.equals(a, b),
        optionsCopy: (a) => BalanceDruidOptions.clone(a),
        optionsToJson: (a) => BalanceDruidOptions.toJson(a),
        optionsFromJson: (obj) => BalanceDruidOptions.fromJson(obj),
        optionsFromPlayer: (player) => player.spec.oneofKind == 'balanceDruid'
            ? player.spec.balanceDruid.options || BalanceDruidOptions.create()
            : BalanceDruidOptions.create(),
    },
    [Spec.SpecElementalShaman]: {
        rotationCreate: () => ElementalShamanRotation.create(),
        rotationEquals: (a, b) => ElementalShamanRotation.equals(a, b),
        rotationCopy: (a) => ElementalShamanRotation.clone(a),
        rotationToJson: (a) => ElementalShamanRotation.toJson(a),
        rotationFromJson: (obj) => ElementalShamanRotation.fromJson(obj),
        rotationFromPlayer: (player) => player.spec.oneofKind == 'elementalShaman'
            ? player.spec.elementalShaman.rotation || ElementalShamanRotation.create()
            : ElementalShamanRotation.create(),
        talentsCreate: () => ShamanTalents.create(),
        talentsEquals: (a, b) => ShamanTalents.equals(a, b),
        talentsCopy: (a) => ShamanTalents.clone(a),
        talentsToJson: (a) => ShamanTalents.toJson(a),
        talentsFromJson: (obj) => ShamanTalents.fromJson(obj),
        talentsFromPlayer: (player) => player.spec.oneofKind == 'elementalShaman'
            ? player.spec.elementalShaman.talents || ShamanTalents.create()
            : ShamanTalents.create(),
        optionsCreate: () => ElementalShamanOptions.create(),
        optionsEquals: (a, b) => ElementalShamanOptions.equals(a, b),
        optionsCopy: (a) => ElementalShamanOptions.clone(a),
        optionsToJson: (a) => ElementalShamanOptions.toJson(a),
        optionsFromJson: (obj) => ElementalShamanOptions.fromJson(obj),
        optionsFromPlayer: (player) => player.spec.oneofKind == 'elementalShaman'
            ? player.spec.elementalShaman.options || ElementalShamanOptions.create()
            : ElementalShamanOptions.create(),
    },
    [Spec.SpecHunter]: {
        rotationCreate: () => HunterRotation.create(),
        rotationEquals: (a, b) => HunterRotation.equals(a, b),
        rotationCopy: (a) => HunterRotation.clone(a),
        rotationToJson: (a) => HunterRotation.toJson(a),
        rotationFromJson: (obj) => HunterRotation.fromJson(obj),
        rotationFromPlayer: (player) => player.spec.oneofKind == 'hunter'
            ? player.spec.hunter.rotation || HunterRotation.create()
            : HunterRotation.create(),
        talentsCreate: () => HunterTalents.create(),
        talentsEquals: (a, b) => HunterTalents.equals(a, b),
        talentsCopy: (a) => HunterTalents.clone(a),
        talentsToJson: (a) => HunterTalents.toJson(a),
        talentsFromJson: (obj) => HunterTalents.fromJson(obj),
        talentsFromPlayer: (player) => player.spec.oneofKind == 'hunter'
            ? player.spec.hunter.talents || HunterTalents.create()
            : HunterTalents.create(),
        optionsCreate: () => HunterOptions.create(),
        optionsEquals: (a, b) => HunterOptions.equals(a, b),
        optionsCopy: (a) => HunterOptions.clone(a),
        optionsToJson: (a) => HunterOptions.toJson(a),
        optionsFromJson: (obj) => HunterOptions.fromJson(obj),
        optionsFromPlayer: (player) => player.spec.oneofKind == 'hunter'
            ? player.spec.hunter.options || HunterOptions.create()
            : HunterOptions.create(),
    },
    [Spec.SpecMage]: {
        rotationCreate: () => MageRotation.create(),
        rotationEquals: (a, b) => MageRotation.equals(a, b),
        rotationCopy: (a) => MageRotation.clone(a),
        rotationToJson: (a) => MageRotation.toJson(a),
        rotationFromJson: (obj) => MageRotation.fromJson(obj),
        rotationFromPlayer: (player) => player.spec.oneofKind == 'mage'
            ? player.spec.mage.rotation || MageRotation.create()
            : MageRotation.create(),
        talentsCreate: () => MageTalents.create(),
        talentsEquals: (a, b) => MageTalents.equals(a, b),
        talentsCopy: (a) => MageTalents.clone(a),
        talentsToJson: (a) => MageTalents.toJson(a),
        talentsFromJson: (obj) => MageTalents.fromJson(obj),
        talentsFromPlayer: (player) => player.spec.oneofKind == 'mage'
            ? player.spec.mage.talents || MageTalents.create()
            : MageTalents.create(),
        optionsCreate: () => MageOptions.create(),
        optionsEquals: (a, b) => MageOptions.equals(a, b),
        optionsCopy: (a) => MageOptions.clone(a),
        optionsToJson: (a) => MageOptions.toJson(a),
        optionsFromJson: (obj) => MageOptions.fromJson(obj),
        optionsFromPlayer: (player) => player.spec.oneofKind == 'mage'
            ? player.spec.mage.options || MageOptions.create()
            : MageOptions.create(),
    },
    [Spec.SpecRetributionPaladin]: {
        rotationCreate: () => RetributionPaladinRotation.create(),
        rotationEquals: (a, b) => RetributionPaladinRotation.equals(a, b),
        rotationCopy: (a) => RetributionPaladinRotation.clone(a),
        rotationToJson: (a) => RetributionPaladinRotation.toJson(a),
        rotationFromJson: (obj) => RetributionPaladinRotation.fromJson(obj),
        rotationFromPlayer: (player) => player.spec.oneofKind == 'retributionPaladin'
            ? player.spec.retributionPaladin.rotation || RetributionPaladinRotation.create()
            : RetributionPaladinRotation.create(),
        talentsCreate: () => PaladinTalents.create(),
        talentsEquals: (a, b) => PaladinTalents.equals(a, b),
        talentsCopy: (a) => PaladinTalents.clone(a),
        talentsToJson: (a) => PaladinTalents.toJson(a),
        talentsFromJson: (obj) => PaladinTalents.fromJson(obj),
        talentsFromPlayer: (player) => player.spec.oneofKind == 'retributionPaladin'
            ? player.spec.retributionPaladin.talents || PaladinTalents.create()
            : PaladinTalents.create(),
        optionsCreate: () => RetributionPaladinOptions.create(),
        optionsEquals: (a, b) => RetributionPaladinOptions.equals(a, b),
        optionsCopy: (a) => RetributionPaladinOptions.clone(a),
        optionsToJson: (a) => RetributionPaladinOptions.toJson(a),
        optionsFromJson: (obj) => RetributionPaladinOptions.fromJson(obj),
        optionsFromPlayer: (player) => player.spec.oneofKind == 'retributionPaladin'
            ? player.spec.retributionPaladin.options || RetributionPaladinOptions.create()
            : RetributionPaladinOptions.create(),
    },
    [Spec.SpecRogue]: {
        rotationCreate: () => RogueRotation.create(),
        rotationEquals: (a, b) => RogueRotation.equals(a, b),
        rotationCopy: (a) => RogueRotation.clone(a),
        rotationToJson: (a) => RogueRotation.toJson(a),
        rotationFromJson: (obj) => RogueRotation.fromJson(obj),
        rotationFromPlayer: (player) => player.spec.oneofKind == 'rogue'
            ? player.spec.rogue.rotation || RogueRotation.create()
            : RogueRotation.create(),
        talentsCreate: () => RogueTalents.create(),
        talentsEquals: (a, b) => RogueTalents.equals(a, b),
        talentsCopy: (a) => RogueTalents.clone(a),
        talentsToJson: (a) => RogueTalents.toJson(a),
        talentsFromJson: (obj) => RogueTalents.fromJson(obj),
        talentsFromPlayer: (player) => player.spec.oneofKind == 'rogue'
            ? player.spec.rogue.talents || RogueTalents.create()
            : RogueTalents.create(),
        optionsCreate: () => RogueOptions.create(),
        optionsEquals: (a, b) => RogueOptions.equals(a, b),
        optionsCopy: (a) => RogueOptions.clone(a),
        optionsToJson: (a) => RogueOptions.toJson(a),
        optionsFromJson: (obj) => RogueOptions.fromJson(obj),
        optionsFromPlayer: (player) => player.spec.oneofKind == 'rogue'
            ? player.spec.rogue.options || RogueOptions.create()
            : RogueOptions.create(),
    },
    [Spec.SpecShadowPriest]: {
        rotationCreate: () => ShadowPriestRotation.create(),
        rotationEquals: (a, b) => ShadowPriestRotation.equals(a, b),
        rotationCopy: (a) => ShadowPriestRotation.clone(a),
        rotationToJson: (a) => ShadowPriestRotation.toJson(a),
        rotationFromJson: (obj) => ShadowPriestRotation.fromJson(obj),
        rotationFromPlayer: (player) => player.spec.oneofKind == 'shadowPriest'
            ? player.spec.shadowPriest.rotation || ShadowPriestRotation.create()
            : ShadowPriestRotation.create(),
        talentsCreate: () => PriestTalents.create(),
        talentsEquals: (a, b) => PriestTalents.equals(a, b),
        talentsCopy: (a) => PriestTalents.clone(a),
        talentsToJson: (a) => PriestTalents.toJson(a),
        talentsFromJson: (obj) => PriestTalents.fromJson(obj),
        talentsFromPlayer: (player) => player.spec.oneofKind == 'shadowPriest'
            ? player.spec.shadowPriest.talents || PriestTalents.create()
            : PriestTalents.create(),
        optionsCreate: () => ShadowPriestOptions.create(),
        optionsEquals: (a, b) => ShadowPriestOptions.equals(a, b),
        optionsCopy: (a) => ShadowPriestOptions.clone(a),
        optionsToJson: (a) => ShadowPriestOptions.toJson(a),
        optionsFromJson: (obj) => ShadowPriestOptions.fromJson(obj),
        optionsFromPlayer: (player) => player.spec.oneofKind == 'shadowPriest'
            ? player.spec.shadowPriest.options || ShadowPriestOptions.create()
            : ShadowPriestOptions.create(),
    },
    [Spec.SpecWarlock]: {
        rotationCreate: () => WarlockRotation.create(),
        rotationEquals: (a, b) => WarlockRotation.equals(a, b),
        rotationCopy: (a) => WarlockRotation.clone(a),
        rotationToJson: (a) => WarlockRotation.toJson(a),
        rotationFromJson: (obj) => WarlockRotation.fromJson(obj),
        rotationFromPlayer: (player) => player.spec.oneofKind == 'warlock'
            ? player.spec.warlock.rotation || WarlockRotation.create()
            : WarlockRotation.create(),
        talentsCreate: () => WarlockTalents.create(),
        talentsEquals: (a, b) => WarlockTalents.equals(a, b),
        talentsCopy: (a) => WarlockTalents.clone(a),
        talentsToJson: (a) => WarlockTalents.toJson(a),
        talentsFromJson: (obj) => WarlockTalents.fromJson(obj),
        talentsFromPlayer: (player) => player.spec.oneofKind == 'warlock'
            ? player.spec.warlock.talents || WarlockTalents.create()
            : WarlockTalents.create(),
        optionsCreate: () => WarlockOptions.create(),
        optionsEquals: (a, b) => WarlockOptions.equals(a, b),
        optionsCopy: (a) => WarlockOptions.clone(a),
        optionsToJson: (a) => WarlockOptions.toJson(a),
        optionsFromJson: (obj) => WarlockOptions.fromJson(obj),
        optionsFromPlayer: (player) => player.spec.oneofKind == 'warlock'
            ? player.spec.warlock.options || WarlockOptions.create()
            : WarlockOptions.create(),
    },
    [Spec.SpecWarrior]: {
        rotationCreate: () => WarriorRotation.create(),
        rotationEquals: (a, b) => WarriorRotation.equals(a, b),
        rotationCopy: (a) => WarriorRotation.clone(a),
        rotationToJson: (a) => WarriorRotation.toJson(a),
        rotationFromJson: (obj) => WarriorRotation.fromJson(obj),
        rotationFromPlayer: (player) => player.spec.oneofKind == 'warrior'
            ? player.spec.warrior.rotation || WarriorRotation.create()
            : WarriorRotation.create(),
        talentsCreate: () => WarriorTalents.create(),
        talentsEquals: (a, b) => WarriorTalents.equals(a, b),
        talentsCopy: (a) => WarriorTalents.clone(a),
        talentsToJson: (a) => WarriorTalents.toJson(a),
        talentsFromJson: (obj) => WarriorTalents.fromJson(obj),
        talentsFromPlayer: (player) => player.spec.oneofKind == 'warrior'
            ? player.spec.warrior.talents || WarriorTalents.create()
            : WarriorTalents.create(),
        optionsCreate: () => WarriorOptions.create(),
        optionsEquals: (a, b) => WarriorOptions.equals(a, b),
        optionsCopy: (a) => WarriorOptions.clone(a),
        optionsToJson: (a) => WarriorOptions.toJson(a),
        optionsFromJson: (obj) => WarriorOptions.fromJson(obj),
        optionsFromPlayer: (player) => player.spec.oneofKind == 'warrior'
            ? player.spec.warrior.options || WarriorOptions.create()
            : WarriorOptions.create(),
    },
};
export var Faction;
(function (Faction) {
    Faction[Faction["Unknown"] = 0] = "Unknown";
    Faction[Faction["Alliance"] = 1] = "Alliance";
    Faction[Faction["Horde"] = 2] = "Horde";
})(Faction || (Faction = {}));
export const raceToFaction = {
    [Race.RaceUnknown]: Faction.Unknown,
    [Race.RaceBloodElf]: Faction.Horde,
    [Race.RaceDraenei]: Faction.Alliance,
    [Race.RaceDwarf]: Faction.Alliance,
    [Race.RaceGnome]: Faction.Alliance,
    [Race.RaceHuman]: Faction.Alliance,
    [Race.RaceNightElf]: Faction.Alliance,
    [Race.RaceOrc]: Faction.Horde,
    [Race.RaceTauren]: Faction.Horde,
    [Race.RaceTroll10]: Faction.Horde,
    [Race.RaceTroll30]: Faction.Horde,
    [Race.RaceUndead]: Faction.Horde,
};
export const specToClass = {
    [Spec.SpecBalanceDruid]: Class.ClassDruid,
    [Spec.SpecElementalShaman]: Class.ClassShaman,
    [Spec.SpecHunter]: Class.ClassHunter,
    [Spec.SpecMage]: Class.ClassMage,
    [Spec.SpecRogue]: Class.ClassRogue,
    [Spec.SpecRetributionPaladin]: Class.ClassPaladin,
    [Spec.SpecShadowPriest]: Class.ClassPriest,
    [Spec.SpecWarlock]: Class.ClassWarlock,
    [Spec.SpecWarrior]: Class.ClassWarrior,
};
const druidRaces = [
    Race.RaceNightElf,
    Race.RaceTauren,
];
const hunterRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceDwarf,
    Race.RaceNightElf,
    Race.RaceOrc,
    Race.RaceTauren,
    Race.RaceTroll10,
    Race.RaceTroll30,
];
const mageRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];
const paladinRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceDwarf,
    Race.RaceHuman,
];
const priestRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceDwarf,
    Race.RaceHuman,
    Race.RaceNightElf,
    Race.RaceOrc,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];
const rogueRaces = [
    Race.RaceBloodElf,
    Race.RaceDwarf,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceNightElf,
    Race.RaceOrc,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];
const shamanRaces = [
    Race.RaceDraenei,
    Race.RaceOrc,
    Race.RaceTauren,
    Race.RaceTroll10,
    Race.RaceTroll30,
];
const warlockRaces = [
    Race.RaceBloodElf,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceOrc,
    Race.RaceUndead,
];
const warriorRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceDwarf,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceNightElf,
    Race.RaceOrc,
    Race.RaceTauren,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];
export const specToEligibleRaces = {
    [Spec.SpecBalanceDruid]: druidRaces,
    [Spec.SpecElementalShaman]: shamanRaces,
    [Spec.SpecHunter]: hunterRaces,
    [Spec.SpecMage]: mageRaces,
    [Spec.SpecRetributionPaladin]: paladinRaces,
    [Spec.SpecRogue]: rogueRaces,
    [Spec.SpecShadowPriest]: priestRaces,
    [Spec.SpecWarlock]: warlockRaces,
    [Spec.SpecWarrior]: warriorRaces,
};
export const specToEligibleItemCategories = {
    [Spec.SpecBalanceDruid]: [ItemCategory.ItemCategoryCaster],
    [Spec.SpecElementalShaman]: [ItemCategory.ItemCategoryCaster],
    [Spec.SpecHunter]: [ItemCategory.ItemCategoryMelee],
    [Spec.SpecMage]: [ItemCategory.ItemCategoryCaster],
    [Spec.SpecRetributionPaladin]: [ItemCategory.ItemCategoryMelee, ItemCategory.ItemCategoryHybrid],
    [Spec.SpecRogue]: [ItemCategory.ItemCategoryMelee],
    [Spec.SpecShadowPriest]: [ItemCategory.ItemCategoryCaster],
    [Spec.SpecWarlock]: [ItemCategory.ItemCategoryCaster],
    [Spec.SpecWarrior]: [ItemCategory.ItemCategoryMelee],
};
// Specs that can dual wield. This could be based on class, except that
// Enhancement Shaman learn dual wield from a talent.
const dualWieldSpecs = [
    Spec.SpecHunter,
    Spec.SpecRogue,
    Spec.SpecWarrior,
];
// Prefixes used for storing browser data for each site. Even if a Spec is
// renamed, DO NOT change these values or people will lose their saved data.
export const specToLocalStorageKey = {
    [Spec.SpecBalanceDruid]: '__balance_druid',
    [Spec.SpecElementalShaman]: '__elemental_shaman',
    [Spec.SpecHunter]: '__hunter',
    [Spec.SpecMage]: '__mage',
    [Spec.SpecRetributionPaladin]: '__retribution_paladin',
    [Spec.SpecRogue]: '__rogue',
    [Spec.SpecShadowPriest]: '__shadow_priest',
    [Spec.SpecWarlock]: '__warlock',
    [Spec.SpecWarrior]: '__warrior',
};
// Returns a copy of playerOptions, with the class field set.
export function withSpecProto(player, rotation, talents, specOptions) {
    const copy = Player.clone(player);
    if (BalanceDruidRotation.is(rotation)) {
        copy.class = Class.ClassDruid;
        copy.spec = {
            oneofKind: 'balanceDruid',
            balanceDruid: BalanceDruid.create({
                rotation: rotation,
                talents: talents,
                options: specOptions,
            }),
        };
    }
    else if (ElementalShamanRotation.is(rotation)) {
        copy.class = Class.ClassShaman;
        copy.spec = {
            oneofKind: 'elementalShaman',
            elementalShaman: ElementalShaman.create({
                rotation: rotation,
                talents: talents,
                options: specOptions,
            }),
        };
    }
    else if (HunterRotation.is(rotation)) {
        copy.class = Class.ClassHunter;
        copy.spec = {
            oneofKind: 'hunter',
            hunter: Hunter.create({
                rotation: rotation,
                talents: talents,
                options: specOptions,
            }),
        };
    }
    else if (MageRotation.is(rotation)) {
        copy.class = Class.ClassMage;
        copy.spec = {
            oneofKind: 'mage',
            mage: Mage.create({
                rotation: rotation,
                talents: talents,
                options: specOptions,
            }),
        };
    }
    else if (RetributionPaladinRotation.is(rotation)) {
        copy.class = Class.ClassPaladin;
        copy.spec = {
            oneofKind: 'retributionPaladin',
            retributionPaladin: RetributionPaladin.create({
                rotation: rotation,
                talents: talents,
                options: specOptions,
            }),
        };
    }
    else if (RogueRotation.is(rotation)) {
        copy.class = Class.ClassRogue;
        copy.spec = {
            oneofKind: 'rogue',
            rogue: Rogue.create({
                rotation: rotation,
                talents: talents,
                options: specOptions,
            }),
        };
    }
    else if (ShadowPriestRotation.is(rotation)) {
        copy.class = Class.ClassPriest;
        copy.spec = {
            oneofKind: 'shadowPriest',
            shadowPriest: ShadowPriest.create({
                rotation: rotation,
                talents: talents,
                options: specOptions,
            }),
        };
    }
    else if (WarlockRotation.is(rotation)) {
        copy.class = Class.ClassWarlock;
        copy.spec = {
            oneofKind: 'warlock',
            warlock: Warlock.create({
                rotation: rotation,
                talents: talents,
                options: specOptions,
            }),
        };
    }
    else if (WarriorRotation.is(rotation)) {
        copy.class = Class.ClassWarrior;
        copy.spec = {
            oneofKind: 'warrior',
            warrior: Warrior.create({
                rotation: rotation,
                talents: talents,
                options: specOptions,
            }),
        };
    }
    else {
        throw new Error('Unrecognized talents with options: ' + Player.toJsonString(player));
    }
    return copy;
}
export function playerToSpec(player) {
    const specValues = getEnumValues(Spec);
    for (let i = 0; i < specValues.length; i++) {
        const spec = specValues[i];
        let specString = Spec[spec]; // Returns 'SpecBalanceDruid' for BalanceDruid.
        specString = specString.substring('Spec'.length); // 'BalanceDruid'
        specString = specString.charAt(0).toLowerCase() + specString.slice(1); // 'balanceDruid'
        if (player.spec.oneofKind == specString) {
            return spec;
        }
    }
    throw new Error('Unable to parse spec from player proto: ' + player);
}
const classToMaxArmorType = {
    [Class.ClassUnknown]: ArmorType.ArmorTypeUnknown,
    [Class.ClassDruid]: ArmorType.ArmorTypeLeather,
    [Class.ClassHunter]: ArmorType.ArmorTypeMail,
    [Class.ClassMage]: ArmorType.ArmorTypeCloth,
    [Class.ClassPaladin]: ArmorType.ArmorTypePlate,
    [Class.ClassPriest]: ArmorType.ArmorTypeCloth,
    [Class.ClassRogue]: ArmorType.ArmorTypeLeather,
    [Class.ClassShaman]: ArmorType.ArmorTypeMail,
    [Class.ClassWarlock]: ArmorType.ArmorTypeCloth,
    [Class.ClassWarrior]: ArmorType.ArmorTypePlate,
};
const classToEligibleRangedWeaponTypes = {
    [Class.ClassUnknown]: [],
    [Class.ClassDruid]: [RangedWeaponType.RangedWeaponTypeIdol],
    [Class.ClassHunter]: [
        RangedWeaponType.RangedWeaponTypeBow,
        RangedWeaponType.RangedWeaponTypeCrossbow,
        RangedWeaponType.RangedWeaponTypeGun,
        RangedWeaponType.RangedWeaponTypeThrown,
    ],
    [Class.ClassMage]: [RangedWeaponType.RangedWeaponTypeWand],
    [Class.ClassPaladin]: [RangedWeaponType.RangedWeaponTypeLibram],
    [Class.ClassPriest]: [RangedWeaponType.RangedWeaponTypeWand],
    [Class.ClassRogue]: [
        RangedWeaponType.RangedWeaponTypeBow,
        RangedWeaponType.RangedWeaponTypeCrossbow,
        RangedWeaponType.RangedWeaponTypeGun,
        RangedWeaponType.RangedWeaponTypeThrown,
    ],
    [Class.ClassShaman]: [RangedWeaponType.RangedWeaponTypeTotem],
    [Class.ClassWarlock]: [RangedWeaponType.RangedWeaponTypeWand],
    [Class.ClassWarrior]: [
        RangedWeaponType.RangedWeaponTypeBow,
        RangedWeaponType.RangedWeaponTypeCrossbow,
        RangedWeaponType.RangedWeaponTypeGun,
        RangedWeaponType.RangedWeaponTypeThrown,
    ],
};
const classToEligibleWeaponTypes = {
    [Class.ClassUnknown]: [],
    [Class.ClassDruid]: [
        { weaponType: WeaponType.WeaponTypeDagger },
        { weaponType: WeaponType.WeaponTypeFist },
        { weaponType: WeaponType.WeaponTypeMace, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeOffHand },
        { weaponType: WeaponType.WeaponTypeStaff, canUseTwoHand: true },
    ],
    [Class.ClassHunter]: [
        { weaponType: WeaponType.WeaponTypeAxe, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeDagger },
        { weaponType: WeaponType.WeaponTypeFist },
        { weaponType: WeaponType.WeaponTypeOffHand },
        { weaponType: WeaponType.WeaponTypePolearm, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeSword, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeStaff, canUseTwoHand: true },
    ],
    [Class.ClassMage]: [
        { weaponType: WeaponType.WeaponTypeDagger },
        { weaponType: WeaponType.WeaponTypeOffHand },
        { weaponType: WeaponType.WeaponTypeStaff, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeSword },
    ],
    [Class.ClassPaladin]: [
        { weaponType: WeaponType.WeaponTypeAxe, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeMace, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeOffHand },
        { weaponType: WeaponType.WeaponTypePolearm, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeShield },
        { weaponType: WeaponType.WeaponTypeSword, canUseTwoHand: true },
    ],
    [Class.ClassPriest]: [
        { weaponType: WeaponType.WeaponTypeDagger },
        { weaponType: WeaponType.WeaponTypeMace },
        { weaponType: WeaponType.WeaponTypeOffHand },
        { weaponType: WeaponType.WeaponTypeStaff, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeSword },
    ],
    [Class.ClassRogue]: [
        { weaponType: WeaponType.WeaponTypeDagger },
        { weaponType: WeaponType.WeaponTypeFist },
        { weaponType: WeaponType.WeaponTypeMace },
        { weaponType: WeaponType.WeaponTypeOffHand },
        { weaponType: WeaponType.WeaponTypeSword },
    ],
    [Class.ClassShaman]: [
        { weaponType: WeaponType.WeaponTypeAxe, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeDagger },
        { weaponType: WeaponType.WeaponTypeFist },
        { weaponType: WeaponType.WeaponTypeMace, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeOffHand },
        { weaponType: WeaponType.WeaponTypeShield },
        { weaponType: WeaponType.WeaponTypeStaff, canUseTwoHand: true },
    ],
    [Class.ClassWarlock]: [
        { weaponType: WeaponType.WeaponTypeDagger },
        { weaponType: WeaponType.WeaponTypeOffHand },
        { weaponType: WeaponType.WeaponTypeStaff, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeSword },
    ],
    [Class.ClassWarrior]: [
        { weaponType: WeaponType.WeaponTypeAxe, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeDagger },
        { weaponType: WeaponType.WeaponTypeFist },
        { weaponType: WeaponType.WeaponTypeMace, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeOffHand },
        { weaponType: WeaponType.WeaponTypePolearm, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeShield },
        { weaponType: WeaponType.WeaponTypeStaff, canUseTwoHand: true },
        { weaponType: WeaponType.WeaponTypeSword, canUseTwoHand: true },
    ],
};
// Custom functions for determining the EP value of meta gem effects.
// Default meta effect EP value is 0, so just handle the ones relevant to your spec.
const metaGemEffectEPs = {
    [Spec.SpecBalanceDruid]: (gem, playerStats) => {
        if (gem.id == Gems.CHAOTIC_SKYFIRE_DIAMOND) {
            // TODO: Fix this
            return (((playerStats.getStat(Stat.StatSpellPower) * 0.795) + 603) * 2 * (playerStats.getStat(Stat.StatSpellCrit) / 2208) * 0.045) / 0.795;
        }
        return 0;
    },
    [Spec.SpecElementalShaman]: (gem, playerStats) => {
        if (gem.id == Gems.CHAOTIC_SKYFIRE_DIAMOND) {
            return (((playerStats.getStat(Stat.StatSpellPower) * 0.795) + 603) * 2 * (playerStats.getStat(Stat.StatSpellCrit) / 2208) * 0.045) / 0.795;
        }
        return 0;
    },
};
export function getMetaGemEffectEP(spec, gem, playerStats) {
    if (metaGemEffectEPs[spec]) {
        return metaGemEffectEPs[spec](gem, playerStats);
    }
    else {
        return 0;
    }
}
// Returns true if this item may be equipped in at least 1 slot for the given Spec.
export function canEquipItem(item, spec) {
    const playerClass = specToClass[spec];
    if (item.classAllowlist.length > 0 && !item.classAllowlist.includes(playerClass)) {
        return false;
    }
    if ([ItemType.ItemTypeFinger, ItemType.ItemTypeTrinket].includes(item.type)) {
        return true;
    }
    if (item.type == ItemType.ItemTypeWeapon) {
        const eligibleWeaponType = classToEligibleWeaponTypes[playerClass].find(wt => wt.weaponType == item.weaponType);
        if (!eligibleWeaponType) {
            return false;
        }
        if (item.handType == HandType.HandTypeOffHand
            && ![WeaponType.WeaponTypeShield, WeaponType.WeaponTypeOffHand].includes(item.weaponType)
            && !dualWieldSpecs.includes(spec)) {
            return false;
        }
        if (item.handType == HandType.HandTypeTwoHand && !eligibleWeaponType.canUseTwoHand) {
            return false;
        }
        return true;
    }
    if (item.type == ItemType.ItemTypeRanged) {
        return classToEligibleRangedWeaponTypes[playerClass].includes(item.rangedWeaponType);
    }
    // At this point, we know the item is an armor piece (feet, chest, legs, etc).
    return classToMaxArmorType[playerClass] >= item.armorType;
}
const itemTypeToSlotsMap = {
    [ItemType.ItemTypeUnknown]: [],
    [ItemType.ItemTypeHead]: [ItemSlot.ItemSlotHead],
    [ItemType.ItemTypeNeck]: [ItemSlot.ItemSlotNeck],
    [ItemType.ItemTypeShoulder]: [ItemSlot.ItemSlotShoulder],
    [ItemType.ItemTypeBack]: [ItemSlot.ItemSlotBack],
    [ItemType.ItemTypeChest]: [ItemSlot.ItemSlotChest],
    [ItemType.ItemTypeWrist]: [ItemSlot.ItemSlotWrist],
    [ItemType.ItemTypeHands]: [ItemSlot.ItemSlotHands],
    [ItemType.ItemTypeWaist]: [ItemSlot.ItemSlotWaist],
    [ItemType.ItemTypeLegs]: [ItemSlot.ItemSlotLegs],
    [ItemType.ItemTypeFeet]: [ItemSlot.ItemSlotFeet],
    [ItemType.ItemTypeFinger]: [ItemSlot.ItemSlotFinger1, ItemSlot.ItemSlotFinger2],
    [ItemType.ItemTypeTrinket]: [ItemSlot.ItemSlotTrinket1, ItemSlot.ItemSlotTrinket2],
    [ItemType.ItemTypeRanged]: [ItemSlot.ItemSlotRanged],
};
export function getEligibleItemSlots(item) {
    if (itemTypeToSlotsMap[item.type]) {
        return itemTypeToSlotsMap[item.type];
    }
    if (item.type == ItemType.ItemTypeWeapon) {
        if ([HandType.HandTypeMainHand, HandType.HandTypeTwoHand].includes(item.handType)) {
            return [ItemSlot.ItemSlotMainHand];
        }
        else if (item.handType == HandType.HandTypeOffHand) {
            return [ItemSlot.ItemSlotOffHand];
        }
        else {
            return [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand];
        }
    }
    // Should never reach here
    throw new Error('Could not find item slots for item: ' + Item.toJsonString(item));
}
;
// Returns whether the given main-hand and off-hand items can be worn at the
// same time.
export function validWeaponCombo(mainHand, offHand) {
    if (mainHand == null || offHand == null) {
        return true;
    }
    if (mainHand.handType == HandType.HandTypeTwoHand) {
        return false;
    }
    return true;
}
// Returns all item slots to which the enchant might be applied.
// 
// Note that this alone is not enough; some items have further restrictions,
// e.g. some weapon enchants may only be applied to 2H weapons.
export function getEligibleEnchantSlots(enchant) {
    if (itemTypeToSlotsMap[enchant.type]) {
        return itemTypeToSlotsMap[enchant.type];
    }
    if (enchant.type == ItemType.ItemTypeWeapon) {
        return [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand];
    }
    // Should never reach here
    throw new Error('Could not find item slots for enchant: ' + Enchant.toJsonString(enchant));
}
;
export function enchantAppliesToItem(enchant, item) {
    const sharedSlots = intersection(getEligibleEnchantSlots(enchant), getEligibleItemSlots(item));
    if (sharedSlots.length == 0)
        return false;
    if (sharedSlots.includes(ItemSlot.ItemSlotMainHand)) {
        if (enchant.twoHandedOnly && item.handType != HandType.HandTypeTwoHand)
            return false;
    }
    if (sharedSlots.includes(ItemSlot.ItemSlotOffHand)) {
        if (enchant.shieldOnly && item.weaponType != WeaponType.WeaponTypeShield)
            return false;
    }
    if (sharedSlots.includes(ItemSlot.ItemSlotRanged)) {
        if (![
            RangedWeaponType.RangedWeaponTypeBow,
            RangedWeaponType.RangedWeaponTypeCrossbow,
            RangedWeaponType.RangedWeaponTypeGun,
        ].includes(item.rangedWeaponType))
            return false;
    }
    return true;
}
;
const socketToMatchingColors = new Map();
socketToMatchingColors.set(GemColor.GemColorMeta, [GemColor.GemColorMeta]);
socketToMatchingColors.set(GemColor.GemColorBlue, [GemColor.GemColorBlue, GemColor.GemColorPurple, GemColor.GemColorGreen]);
socketToMatchingColors.set(GemColor.GemColorRed, [GemColor.GemColorRed, GemColor.GemColorPurple, GemColor.GemColorOrange]);
socketToMatchingColors.set(GemColor.GemColorYellow, [GemColor.GemColorYellow, GemColor.GemColorOrange, GemColor.GemColorGreen]);
// Whether the gem matches the given socket color, for the purposes of gaining the socket bonuses.
export function gemMatchesSocket(gem, socketColor) {
    return socketToMatchingColors.has(socketColor) && socketToMatchingColors.get(socketColor).includes(gem.color);
}
// Whether the gem is capable of slotting into a socket of the given color.
export function gemEligibleForSocket(gem, socketColor) {
    return (gem.color == GemColor.GemColorMeta) == (socketColor == GemColor.GemColorMeta);
}
export const NO_TARGET = -1;
