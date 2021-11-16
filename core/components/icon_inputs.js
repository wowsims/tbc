import { Drums } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
// Keep each section in alphabetical order.
// Raid Buffs
export const ArcaneBrilliance = makeBooleanRaidBuffInput({ spellId: 27127 }, 'arcaneBrilliance');
export const DivineSpirit = makeTristateRaidBuffInput({ spellId: 25312 }, { spellId: 33182 }, 'divineSpirit');
export const GiftOfTheWild = makeTristateRaidBuffInput({ spellId: 26991 }, { spellId: 17055 }, 'giftOfTheWild');
// Party Buffs
export const AtieshMage = makeMultistatePartyBuffInput({ spellId: 28142 }, 5, 'atieshMage');
export const AtieshWarlock = makeMultistatePartyBuffInput({ spellId: 28143 }, 5, 'atieshWarlock');
export const Bloodlust = makeMultistatePartyBuffInput({ spellId: 2825 }, 11, 'bloodlust');
export const BraidedEterniumChain = makeBooleanPartyBuffInput({ spellId: 31025 }, 'braidedEterniumChain');
export const ChainOfTheTwilightOwl = makeBooleanPartyBuffInput({ spellId: 31035 }, 'chainOfTheTwilightOwl');
export const DraeneiRacialCaster = makeBooleanPartyBuffInput({ spellId: 28878 }, 'draeneiRacialCaster');
export const DraeneiRacialMelee = makeBooleanPartyBuffInput({ spellId: 6562 }, 'draeneiRacialMelee');
export const EyeOfTheNight = makeBooleanPartyBuffInput({ spellId: 31033 }, 'eyeOfTheNight');
export const JadePendantOfBlasting = makeBooleanPartyBuffInput({ spellId: 25607 }, 'jadePendantOfBlasting');
export const ManaSpringTotem = makeTristatePartyBuffInput({ spellId: 25570 }, { spellId: 16208 }, 'manaSpringTotem');
export const MoonkinAura = makeTristatePartyBuffInput({ spellId: 24907 }, { itemId: 32387 }, 'moonkinAura');
export const TotemOfWrath = makeMultistatePartyBuffInput({ spellId: 30706 }, 5, 'totemOfWrath');
export const WrathOfAirTotem = makeTristatePartyBuffInput({ spellId: 3738 }, { spellId: 37212 }, 'wrathOfAirTotem');
export const DrumsOfBattleBuff = makeEnumValuePartyBuffInput({ spellId: 35476 }, 'drums', Drums.DrumsOfBattle, ['Drums']);
export const DrumsOfRestorationBuff = makeEnumValuePartyBuffInput({ spellId: 35478 }, 'drums', Drums.DrumsOfRestoration, ['Drums']);
// Individual Buffs
export const BlessingOfKings = makeBooleanIndividualBuffInput({ spellId: 25898 }, 'blessingOfKings');
export const BlessingOfWisdom = makeTristateIndividualBuffInput({ spellId: 27143 }, { spellId: 20245 }, 'blessingOfWisdom');
export const ManaTideTotem = makeBooleanIndividualBuffInput({ spellId: 16190 }, 'manaTideTotem');
// Debuffs
export const ImprovedSealOfTheCrusader = makeBooleanDebuffInput({ spellId: 20337 }, 'improvedSealOfTheCrusader');
export const JudgementOfWisdom = makeBooleanDebuffInput({ spellId: 27164 }, 'judgementOfWisdom');
export const Misery = makeBooleanDebuffInput({ spellId: 33195 }, 'misery');
export const CurseOfElements = makeTristateDebuffInput({ spellId: 27228 }, { spellId: 32484 }, 'curseOfElements');
// Consumes
export const AdeptsElixir = makeBooleanConsumeInput({ itemId: 28103 }, 'adeptsElixir', ['Battle Elixir']);
export const BlackenedBasilisk = makeBooleanConsumeInput({ itemId: 27657 }, 'blackenedBasilisk', ['Food']);
export const BrilliantWizardOil = makeBooleanConsumeInput({ itemId: 20749 }, 'brilliantWizardOil', ['Weapon Imbue']);
export const DarkRune = makeBooleanConsumeInput({ itemId: 12662 }, 'darkRune', ['Rune']);
export const ElixirOfDraenicWisdom = makeBooleanConsumeInput({ itemId: 32067 }, 'elixirOfDraenicWisdom', ['Guardian Elixir']);
export const ElixirOfMajorFirePower = makeBooleanConsumeInput({ itemId: 22833 }, 'elixirOfMajorFirePower', ['Battle Elixir']);
export const ElixirOfMajorFrostPower = makeBooleanConsumeInput({ itemId: 22827 }, 'elixirOfMajorFrostPower', ['Battle Elixir']);
export const ElixirOfMajorMageblood = makeBooleanConsumeInput({ itemId: 22840 }, 'elixirOfMajorMageblood', ['Guardian Elixir']);
export const ElixirOfMajorShadowPower = makeBooleanConsumeInput({ itemId: 22835 }, 'elixirOfMajorShadowPower', ['Battle Elixir']);
export const FlaskOfBlindingLight = makeBooleanConsumeInput({ itemId: 22861 }, 'flaskOfBlindingLight', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfMightyRestoration = makeBooleanConsumeInput({ itemId: 22853 }, 'flaskOfMightyRestoration', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfPureDeath = makeBooleanConsumeInput({ itemId: 22866 }, 'flaskOfPureDeath', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfSupremePower = makeBooleanConsumeInput({ itemId: 13512 }, 'flaskOfSupremePower', ['Battle Elixir', 'Guardian Elixir']);
export const KreegsStoutBeatdown = makeBooleanConsumeInput({ itemId: 18284 }, 'kreegsStoutBeatdown', ['Alchohol']);
export const SkullfishSoup = makeBooleanConsumeInput({ itemId: 33825 }, 'skullfishSoup', ['Food']);
export const SuperiorWizardOil = makeBooleanConsumeInput({ itemId: 22522 }, 'superiorWizardOil', ['Weapon Imbue']);
export const DefaultDestructionPotion = makeEnumValueConsumeInput({ itemId: 22839 }, 'defaultPotion', Potions.DestructionPotion, ['Potion']);
export const DefaultSuperManaPotion = makeEnumValueConsumeInput({ itemId: 22832 }, 'defaultPotion', Potions.SuperManaPotion, ['Potion']);
export const DrumsOfBattleConsume = makeEnumValueConsumeInput({ spellId: 35476 }, 'drums', Drums.DrumsOfBattle, ['Drums']);
export const DrumsOfRestorationConsume = makeEnumValueConsumeInput({ spellId: 35478 }, 'drums', Drums.DrumsOfRestoration, ['Drums']);
function makeBooleanRaidBuffInput(id, raidBuffsFieldName, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (sim) => sim.raidBuffsChangeEmitter,
        getValue: (sim) => sim.getRaidBuffs()[raidBuffsFieldName],
        setBooleanValue: (sim, newValue) => {
            const newRaidBuffs = sim.getRaidBuffs();
            newRaidBuffs[raidBuffsFieldName] = newValue;
            sim.setRaidBuffs(newRaidBuffs);
        },
    };
}
function makeTristateRaidBuffInput(id, impId, raidBuffsFieldName) {
    return {
        id: id,
        states: 3,
        improvedId: impId,
        changedEvent: (sim) => sim.raidBuffsChangeEmitter,
        getValue: (sim) => sim.getRaidBuffs()[raidBuffsFieldName],
        setNumberValue: (sim, newValue) => {
            const newRaidBuffs = sim.getRaidBuffs();
            newRaidBuffs[raidBuffsFieldName] = newValue;
            sim.setRaidBuffs(newRaidBuffs);
        },
    };
}
function makeBooleanPartyBuffInput(id, partyBuffsFieldName, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (sim) => sim.partyBuffsChangeEmitter,
        getValue: (sim) => sim.getPartyBuffs()[partyBuffsFieldName],
        setBooleanValue: (sim, newValue) => {
            const newPartyBuffs = sim.getPartyBuffs();
            newPartyBuffs[partyBuffsFieldName] = newValue;
            sim.setPartyBuffs(newPartyBuffs);
        },
    };
}
function makeTristatePartyBuffInput(id, impId, partyBuffsFieldName) {
    return {
        id: id,
        states: 3,
        improvedId: impId,
        changedEvent: (sim) => sim.partyBuffsChangeEmitter,
        getValue: (sim) => sim.getPartyBuffs()[partyBuffsFieldName],
        setNumberValue: (sim, newValue) => {
            const newPartyBuffs = sim.getPartyBuffs();
            newPartyBuffs[partyBuffsFieldName] = newValue;
            sim.setPartyBuffs(newPartyBuffs);
        },
    };
}
function makeMultistatePartyBuffInput(id, numStates, partyBuffsFieldName) {
    return {
        id: id,
        states: numStates,
        changedEvent: (sim) => sim.partyBuffsChangeEmitter,
        getValue: (sim) => sim.getPartyBuffs()[partyBuffsFieldName],
        setNumberValue: (sim, newValue) => {
            const newPartyBuffs = sim.getPartyBuffs();
            newPartyBuffs[partyBuffsFieldName] = newValue;
            sim.setPartyBuffs(newPartyBuffs);
        },
    };
}
function makeEnumValuePartyBuffInput(id, partyBuffsFieldName, enumValue, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (sim) => sim.partyBuffsChangeEmitter,
        getValue: (sim) => sim.getPartyBuffs()[partyBuffsFieldName] == enumValue,
        setBooleanValue: (sim, newValue) => {
            const newPartyBuffs = sim.getPartyBuffs();
            newPartyBuffs[partyBuffsFieldName] = newValue ? enumValue : 0;
            sim.setPartyBuffs(newPartyBuffs);
        },
    };
}
function makeBooleanIndividualBuffInput(id, individualBuffsFieldName, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (sim) => sim.individualBuffsChangeEmitter,
        getValue: (sim) => sim.getIndividualBuffs()[individualBuffsFieldName],
        setBooleanValue: (sim, newValue) => {
            const newIndividualBuffs = sim.getIndividualBuffs();
            newIndividualBuffs[individualBuffsFieldName] = newValue;
            sim.setIndividualBuffs(newIndividualBuffs);
        },
    };
}
function makeTristateIndividualBuffInput(id, impId, individualBuffsFieldName) {
    return {
        id: id,
        states: 3,
        improvedId: impId,
        changedEvent: (sim) => sim.individualBuffsChangeEmitter,
        getValue: (sim) => sim.getIndividualBuffs()[individualBuffsFieldName],
        setNumberValue: (sim, newValue) => {
            const newIndividualBuffs = sim.getIndividualBuffs();
            newIndividualBuffs[individualBuffsFieldName] = newValue;
            sim.setIndividualBuffs(newIndividualBuffs);
        },
    };
}
function makeBooleanDebuffInput(id, debuffsFieldName, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (target) => target.debuffsChangeEmitter,
        getValue: (target) => target.getDebuffs()[debuffsFieldName],
        setBooleanValue: (target, newValue) => {
            const newDebuffs = target.getDebuffs();
            newDebuffs[debuffsFieldName] = newValue;
            target.setDebuffs(newDebuffs);
        },
    };
}
function makeTristateDebuffInput(id, impId, debuffsFieldName) {
    return {
        id: id,
        states: 3,
        improvedId: impId,
        changedEvent: (target) => target.debuffsChangeEmitter,
        getValue: (target) => target.getDebuffs()[debuffsFieldName],
        setNumberValue: (target, newValue) => {
            const newDebuffs = target.getDebuffs();
            newDebuffs[debuffsFieldName] = newValue;
            target.setDebuffs(newDebuffs);
        },
    };
}
function makeBooleanConsumeInput(id, consumesFieldName, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (player) => player.consumesChangeEmitter,
        getValue: (player) => player.getConsumes()[consumesFieldName],
        setBooleanValue: (player, newValue) => {
            const newBuffs = player.getConsumes();
            newBuffs[consumesFieldName] = newValue;
            player.setConsumes(newBuffs);
        },
    };
}
function makeEnumValueConsumeInput(id, consumesFieldName, enumValue, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (player) => player.consumesChangeEmitter,
        getValue: (player) => player.getConsumes()[consumesFieldName] == enumValue,
        setBooleanValue: (player, newValue) => {
            const newConsumes = player.getConsumes();
            newConsumes[consumesFieldName] = newValue ? enumValue : 0;
            player.setConsumes(newConsumes);
        },
    };
}
