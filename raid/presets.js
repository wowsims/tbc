import { Class } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { specIconsLarge } from '/tbc/core/proto_utils/utils.js';
import { specNames } from '/tbc/core/proto_utils/utils.js';
import { talentTreeIcons } from '/tbc/core/proto_utils/utils.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import * as BalanceDruidPresets from '/tbc/balance_druid/presets.js';
import * as ElementalShamanPresets from '/tbc/elemental_shaman/presets.js';
import * as EnhancementShamanPresets from '/tbc/enhancement_shaman/presets.js';
import * as HunterPresets from '/tbc/hunter/presets.js';
import * as MagePresets from '/tbc/mage/presets.js';
import * as ShadowPriestPresets from '/tbc/shadow_priest/presets.js';
import { BalanceDruidSimUI } from '/tbc/balance_druid/sim.js';
import { EnhancementShamanSimUI } from '/tbc/enhancement_shaman/sim.js';
import { ElementalShamanSimUI } from '/tbc/elemental_shaman/sim.js';
import { HunterSimUI } from '/tbc/hunter/sim.js';
import { MageSimUI } from '/tbc/mage/sim.js';
import { ShadowPriestSimUI } from '/tbc/shadow_priest/sim.js';
export const specSimFactories = {
    [Spec.SpecBalanceDruid]: (parentElem, player) => new BalanceDruidSimUI(parentElem, player),
    [Spec.SpecElementalShaman]: (parentElem, player) => new ElementalShamanSimUI(parentElem, player),
    [Spec.SpecEnhancementShaman]: (parentElem, player) => new EnhancementShamanSimUI(parentElem, player),
    [Spec.SpecHunter]: (parentElem, player) => new HunterSimUI(parentElem, player),
    [Spec.SpecMage]: (parentElem, player) => new MageSimUI(parentElem, player),
    [Spec.SpecShadowPriest]: (parentElem, player) => new ShadowPriestSimUI(parentElem, player),
};
export const playerPresets = [
    {
        spec: Spec.SpecBalanceDruid,
        rotation: BalanceDruidPresets.DefaultRotation,
        talents: BalanceDruidPresets.StandardTalents.data,
        specOptions: BalanceDruidPresets.DefaultOptions,
        consumes: BalanceDruidPresets.DefaultConsumes,
        defaultName: 'Balance Druid',
        defaultFactionRaces: {
            [Faction.Unknown]: Race.RaceUnknown,
            [Faction.Alliance]: Race.RaceNightElf,
            [Faction.Horde]: Race.RaceTauren,
        },
        defaultGear: {
            [Faction.Unknown]: {},
            [Faction.Alliance]: {
                1: BalanceDruidPresets.P1_ALLIANCE_PRESET.gear,
                2: BalanceDruidPresets.P2_ALLIANCE_PRESET.gear,
                3: BalanceDruidPresets.P3_PRESET.gear,
            },
            [Faction.Horde]: {
                1: BalanceDruidPresets.P1_HORDE_PRESET.gear,
                2: BalanceDruidPresets.P2_HORDE_PRESET.gear,
                3: BalanceDruidPresets.P3_PRESET.gear,
            },
        },
        tooltip: specNames[Spec.SpecBalanceDruid],
        iconUrl: specIconsLarge[Spec.SpecBalanceDruid],
    },
    {
        spec: Spec.SpecHunter,
        rotation: HunterPresets.DefaultRotation,
        talents: HunterPresets.BeastMasteryTalents.data,
        specOptions: HunterPresets.DefaultOptions,
        consumes: HunterPresets.DefaultConsumes,
        defaultName: 'BM Hunter',
        defaultFactionRaces: {
            [Faction.Unknown]: Race.RaceUnknown,
            [Faction.Alliance]: Race.RaceNightElf,
            [Faction.Horde]: Race.RaceOrc,
        },
        defaultGear: {
            [Faction.Unknown]: {},
            [Faction.Alliance]: {
                1: HunterPresets.P1_BM_PRESET.gear,
                2: HunterPresets.P2_BM_PRESET.gear,
                3: HunterPresets.P3_BM_PRESET.gear,
            },
            [Faction.Horde]: {
                1: HunterPresets.P1_BM_PRESET.gear,
                2: HunterPresets.P2_BM_PRESET.gear,
                3: HunterPresets.P3_BM_PRESET.gear,
            },
        },
        tooltip: 'BM Hunter',
        iconUrl: talentTreeIcons[Class.ClassHunter][0],
    },
    {
        spec: Spec.SpecHunter,
        rotation: HunterPresets.DefaultRotation,
        talents: HunterPresets.SurvivalTalents.data,
        specOptions: HunterPresets.DefaultOptions,
        consumes: HunterPresets.DefaultConsumes,
        defaultName: 'SV Hunter',
        defaultFactionRaces: {
            [Faction.Unknown]: Race.RaceUnknown,
            [Faction.Alliance]: Race.RaceNightElf,
            [Faction.Horde]: Race.RaceOrc,
        },
        defaultGear: {
            [Faction.Unknown]: {},
            [Faction.Alliance]: {
                1: HunterPresets.P1_SV_PRESET.gear,
                2: HunterPresets.P2_SV_PRESET.gear,
                3: HunterPresets.P3_SV_PRESET.gear,
            },
            [Faction.Horde]: {
                1: HunterPresets.P1_SV_PRESET.gear,
                2: HunterPresets.P2_SV_PRESET.gear,
                3: HunterPresets.P3_SV_PRESET.gear,
            },
        },
        tooltip: 'SV Hunter',
        iconUrl: talentTreeIcons[Class.ClassHunter][2],
    },
    {
        spec: Spec.SpecMage,
        rotation: MagePresets.DefaultArcaneRotation,
        talents: MagePresets.ArcaneTalents.data,
        specOptions: MagePresets.DefaultArcaneOptions,
        consumes: MagePresets.DefaultArcaneConsumes,
        defaultName: 'Arcane Mage',
        defaultFactionRaces: {
            [Faction.Unknown]: Race.RaceUnknown,
            [Faction.Alliance]: Race.RaceGnome,
            [Faction.Horde]: Race.RaceTroll10,
        },
        defaultGear: {
            [Faction.Unknown]: {},
            [Faction.Alliance]: {
                1: MagePresets.P1_ARCANE_PRESET.gear,
                2: MagePresets.P2_ARCANE_PRESET.gear,
                3: MagePresets.P3_ARCANE_PRESET.gear,
            },
            [Faction.Horde]: {
                1: MagePresets.P1_ARCANE_PRESET.gear,
                2: MagePresets.P2_ARCANE_PRESET.gear,
                3: MagePresets.P3_ARCANE_PRESET.gear,
            },
        },
        tooltip: 'Arcane Mage',
        iconUrl: talentTreeIcons[Class.ClassMage][0],
    },
    {
        spec: Spec.SpecMage,
        rotation: MagePresets.DefaultFireRotation,
        talents: MagePresets.FireTalents.data,
        specOptions: MagePresets.DefaultFireOptions,
        consumes: MagePresets.DefaultFireConsumes,
        defaultName: 'Fire Mage',
        defaultFactionRaces: {
            [Faction.Unknown]: Race.RaceUnknown,
            [Faction.Alliance]: Race.RaceGnome,
            [Faction.Horde]: Race.RaceTroll10,
        },
        defaultGear: {
            [Faction.Unknown]: {},
            [Faction.Alliance]: {
                1: MagePresets.P1_FIRE_PRESET.gear,
                2: MagePresets.P2_FIRE_PRESET.gear,
                3: MagePresets.P3_FIRE_PRESET.gear,
            },
            [Faction.Horde]: {
                1: MagePresets.P1_FIRE_PRESET.gear,
                2: MagePresets.P2_FIRE_PRESET.gear,
                3: MagePresets.P3_FIRE_PRESET.gear,
            },
        },
        tooltip: 'Fire Mage',
        iconUrl: talentTreeIcons[Class.ClassMage][1],
    },
    {
        spec: Spec.SpecMage,
        rotation: MagePresets.DefaultFrostRotation,
        talents: MagePresets.DeepFrostTalents.data,
        specOptions: MagePresets.DefaultFrostOptions,
        consumes: MagePresets.DefaultFrostConsumes,
        defaultName: 'Frost Mage',
        defaultFactionRaces: {
            [Faction.Unknown]: Race.RaceUnknown,
            [Faction.Alliance]: Race.RaceGnome,
            [Faction.Horde]: Race.RaceTroll10,
        },
        defaultGear: {
            [Faction.Unknown]: {},
            [Faction.Alliance]: {
                1: MagePresets.P1_FROST_PRESET.gear,
                2: MagePresets.P2_FROST_PRESET.gear,
                3: MagePresets.P3_FROST_PRESET.gear,
            },
            [Faction.Horde]: {
                1: MagePresets.P1_FROST_PRESET.gear,
                2: MagePresets.P2_FROST_PRESET.gear,
                3: MagePresets.P3_FROST_PRESET.gear,
            },
        },
        tooltip: 'Frost Mage',
        iconUrl: talentTreeIcons[Class.ClassMage][2],
    },
    {
        spec: Spec.SpecElementalShaman,
        rotation: ElementalShamanPresets.DefaultRotation,
        talents: ElementalShamanPresets.StandardTalents.data,
        specOptions: ElementalShamanPresets.DefaultOptions,
        consumes: ElementalShamanPresets.DefaultConsumes,
        defaultName: 'Ele Shaman',
        defaultFactionRaces: {
            [Faction.Unknown]: Race.RaceUnknown,
            [Faction.Alliance]: Race.RaceDraenei,
            [Faction.Horde]: Race.RaceOrc,
        },
        defaultGear: {
            [Faction.Unknown]: {},
            [Faction.Alliance]: {
                1: ElementalShamanPresets.P1_PRESET.gear,
                2: ElementalShamanPresets.P2_PRESET.gear,
                3: ElementalShamanPresets.P3_PRESET.gear,
            },
            [Faction.Horde]: {
                1: ElementalShamanPresets.P1_PRESET.gear,
                2: ElementalShamanPresets.P2_PRESET.gear,
                3: ElementalShamanPresets.P3_PRESET.gear,
            },
        },
        tooltip: specNames[Spec.SpecElementalShaman],
        iconUrl: specIconsLarge[Spec.SpecElementalShaman],
    },
    {
        spec: Spec.SpecEnhancementShaman,
        rotation: EnhancementShamanPresets.DefaultRotation,
        talents: EnhancementShamanPresets.StandardTalents.data,
        specOptions: EnhancementShamanPresets.DefaultOptions,
        consumes: EnhancementShamanPresets.DefaultConsumes,
        defaultName: 'Enh Shaman',
        defaultFactionRaces: {
            [Faction.Unknown]: Race.RaceUnknown,
            [Faction.Alliance]: Race.RaceDraenei,
            [Faction.Horde]: Race.RaceOrc,
        },
        defaultGear: {
            [Faction.Unknown]: {},
            [Faction.Alliance]: {
                1: EnhancementShamanPresets.P1_PRESET.gear,
                2: EnhancementShamanPresets.P2_PRESET.gear,
                3: EnhancementShamanPresets.P3_PRESET.gear,
            },
            [Faction.Horde]: {
                1: EnhancementShamanPresets.P1_PRESET.gear,
                2: EnhancementShamanPresets.P2_PRESET.gear,
                3: EnhancementShamanPresets.P3_PRESET.gear,
            },
        },
        tooltip: specNames[Spec.SpecEnhancementShaman],
        iconUrl: specIconsLarge[Spec.SpecEnhancementShaman],
    },
    {
        spec: Spec.SpecShadowPriest,
        rotation: ShadowPriestPresets.DefaultRotation,
        talents: ShadowPriestPresets.StandardTalents.data,
        specOptions: ShadowPriestPresets.DefaultOptions,
        consumes: ShadowPriestPresets.DefaultConsumes,
        defaultName: 'Shadow Priest',
        defaultFactionRaces: {
            [Faction.Unknown]: Race.RaceUnknown,
            [Faction.Alliance]: Race.RaceDwarf,
            [Faction.Horde]: Race.RaceUndead,
        },
        defaultGear: {
            [Faction.Unknown]: {},
            [Faction.Alliance]: {
                1: ShadowPriestPresets.P1_PRESET.gear,
                2: ShadowPriestPresets.P2_PRESET.gear,
                3: ShadowPriestPresets.P3_PRESET.gear,
            },
            [Faction.Horde]: {
                1: ShadowPriestPresets.P1_PRESET.gear,
                2: ShadowPriestPresets.P2_PRESET.gear,
                3: ShadowPriestPresets.P3_PRESET.gear,
            },
        },
        tooltip: specNames[Spec.SpecShadowPriest],
        iconUrl: specIconsLarge[Spec.SpecShadowPriest],
    },
];
export const implementedSpecs = [...new Set(playerPresets.map(preset => preset.spec))];
export const buffBotPresets = [
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Bear',
        spec: Spec.SpecBalanceDruid,
        name: 'Bear',
        tooltip: 'Bear: Adds Gift of the Wild, an Innervate, and Leader of the Pack.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/ability_racial_bearform.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            raidProto.buffs.giftOfTheWild = Math.max(raidProto.buffs.giftOfTheWild, TristateEffect.TristateEffectRegular);
            partyProto.buffs.leaderOfThePack = Math.max(partyProto.buffs.leaderOfThePack, TristateEffect.TristateEffectRegular);
            const innervateIndex = buffBot.getInnervateAssignment().targetIndex;
            if (innervateIndex != NO_TARGET) {
                const partyIndex = Math.floor(innervateIndex / 5);
                const playerIndex = innervateIndex % 5;
                const playerProto = raidProto.parties[partyIndex].players[playerIndex];
                if (playerProto.buffs) {
                    playerProto.buffs.innervates++;
                }
            }
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Resto Druid',
        spec: Spec.SpecBalanceDruid,
        name: 'Resto Druid',
        tooltip: 'Resto Druid: Adds Improved Gift of the Wild, and an Innervate.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_nature_healingtouch.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            raidProto.buffs.giftOfTheWild = TristateEffect.TristateEffectImproved;
            const innervateIndex = buffBot.getInnervateAssignment().targetIndex;
            if (innervateIndex != NO_TARGET) {
                const partyIndex = Math.floor(innervateIndex / 5);
                const playerIndex = innervateIndex % 5;
                const playerProto = raidProto.parties[partyIndex].players[playerIndex];
                if (playerProto.buffs) {
                    playerProto.buffs.innervates++;
                }
            }
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Mage',
        deprecated: true,
        spec: Spec.SpecMage,
        name: 'Mage',
        tooltip: 'Mage: Adds Arcane Brilliance.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_arcaneintellect.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            raidProto.buffs.arcaneBrilliance = true;
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Paladin',
        spec: Spec.SpecRetributionPaladin,
        name: 'Holy Paladin',
        tooltip: 'Holy Paladin: Adds a set of blessings.',
        iconUrl: talentTreeIcons[Class.ClassPaladin][0],
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            // Do nothing, blessings are handled elswhere.
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'JoW Paladin',
        spec: Spec.SpecRetributionPaladin,
        name: 'JoW Paladin',
        tooltip: 'JoW Paladin: Adds a set of blessings and Judgement of Wisdom.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_righteousnessaura.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            // Do nothing, blessings are handled elswhere.
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
            encounterProto.targets[0].debuffs.judgementOfWisdom = true;
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'JoC Paladin',
        spec: Spec.SpecRetributionPaladin,
        name: 'JoC Paladin',
        tooltip: 'JoC Paladin: Adds a set of blessings and Improved Judgement of the Crusader (+3% crit).',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_holy_holysmite.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            // Do nothing, blessings are handled elswhere.
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
            encounterProto.targets[0].debuffs.improvedSealOfTheCrusader = true;
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Holy Priest',
        spec: Spec.SpecShadowPriest,
        name: 'Holy Priest',
        tooltip: 'Holy Priest: Doesn\'t contribute to DPS, just fills a raid slot.',
        iconUrl: talentTreeIcons[Class.ClassPriest][1],
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Divine Spirit Priest',
        spec: Spec.SpecShadowPriest,
        name: 'Disc Priest',
        tooltip: 'Disc Priest: Adds Improved Divine Spirit and a Power Infusion.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/medium/spell_holy_powerinfusion.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            raidProto.buffs.divineSpirit = TristateEffect.TristateEffectImproved;
            const powerInfusionIndex = buffBot.getPowerInfusionAssignment().targetIndex;
            if (powerInfusionIndex != NO_TARGET) {
                const partyIndex = Math.floor(powerInfusionIndex / 5);
                const playerIndex = powerInfusionIndex % 5;
                const playerProto = raidProto.parties[partyIndex].players[playerIndex];
                if (playerProto.buffs) {
                    playerProto.buffs.powerInfusions++;
                }
            }
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Resto Shaman',
        spec: Spec.SpecElementalShaman,
        name: 'Resto Shaman',
        tooltip: 'Resto Shaman: Adds Bloodlust, Mana Spring Totem, Wrath of Air Totem, Mana Tide Totem, and Drums of Battle.',
        iconUrl: talentTreeIcons[Class.ClassShaman][2],
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            partyProto.buffs.bloodlust++;
            partyProto.buffs.manaSpringTotem = TristateEffect.TristateEffectImproved;
            partyProto.buffs.wrathOfAirTotem = Math.max(partyProto.buffs.wrathOfAirTotem, TristateEffect.TristateEffectRegular);
            partyProto.buffs.manaTideTotems++;
            partyProto.buffs.drums = Drums.DrumsOfBattle;
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'CoE Warlock',
        spec: Spec.SpecWarlock,
        name: 'CoE Warlock',
        tooltip: 'CoE Warlock: Adds Curse of Elements (regular). Also adds +20% uptime to ISB.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_shadow_chilltouch.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
            const debuffs = encounterProto.targets[0].debuffs;
            debuffs.curseOfElements = Math.max(debuffs.curseOfElements, TristateEffect.TristateEffectRegular);
            debuffs.isbUptime = Math.min(1.0, debuffs.isbUptime + 0.2);
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Malediction Warlock',
        spec: Spec.SpecWarlock,
        name: 'Aff Warlock',
        tooltip: 'Afflication Warlock: Adds Curse of Elements (improved). Also adds +20% uptime to ISB.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/spell_shadow_curseofachimonde.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
            const debuffs = encounterProto.targets[0].debuffs;
            debuffs.curseOfElements = TristateEffect.TristateEffectImproved;
            debuffs.isbUptime = Math.min(1.0, debuffs.isbUptime + 0.2);
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Arms Warrior',
        spec: Spec.SpecWarrior,
        name: 'Arms Warrior',
        tooltip: 'Arms Warrior: Adds Sunder Armor, Blood Frenzy, and Improved Battle Shout.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/medium/ability_warrior_savageblow.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            partyProto.buffs.battleShout = TristateEffect.TristateEffectImproved;
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
            const debuffs = encounterProto.targets[0].debuffs;
            debuffs.sunderArmor = true;
            debuffs.bloodFrenzy = true;
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Fury Warrior',
        spec: Spec.SpecWarrior,
        name: 'Fury Warrior',
        tooltip: 'Fury Warrior: Adds Sunder Armor and Improved Battle Shout.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/medium/ability_warrior_innerrage.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
            partyProto.buffs.battleShout = TristateEffect.TristateEffectImproved;
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
            const debuffs = encounterProto.targets[0].debuffs;
            debuffs.sunderArmor = true;
        },
    },
    {
        // The value of this field must never change, to preserve local storage data.
        buffBotId: 'Prot Warrior',
        spec: Spec.SpecWarrior,
        name: 'Prot Warrior',
        tooltip: 'Prot Warrior: Adds Sunder Armor.',
        iconUrl: 'https://wow.zamimg.com/images/wow/icons/medium/inv_shield_06.jpg',
        modifyRaidProto: (buffBot, raidProto, partyProto) => {
        },
        modifyEncounterProto: (buffBot, encounterProto) => {
            const debuffs = encounterProto.targets[0].debuffs;
            debuffs.sunderArmor = true;
        },
    },
];