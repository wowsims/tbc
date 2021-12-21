import { Drums } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
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
export const Innervate = makeMultistateIndividualBuffInput({ spellId: 29166 }, 6, 'innervates');
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
function removeOtherPartyMembersDrums(eventID, player, newValue) {
    if (newValue) {
        player.getOtherPartyMembers().forEach(otherPlayer => {
            const otherConsumes = otherPlayer.getConsumes();
            otherConsumes.drums = Drums.DrumsUnknown;
            otherPlayer.setConsumes(eventID, otherConsumes);
        });
    }
}
;
export const DrumsOfBattleConsume = makeEnumValueConsumeInput({ spellId: 35476 }, 'drums', Drums.DrumsOfBattle, ['Drums'], removeOtherPartyMembersDrums);
export const DrumsOfRestorationConsume = makeEnumValueConsumeInput({ spellId: 35478 }, 'drums', Drums.DrumsOfRestoration, ['Drums'], removeOtherPartyMembersDrums);
function makeBooleanRaidBuffInput(id, buffsFieldName, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (raid) => raid.buffsChangeEmitter,
        getValue: (raid) => raid.getBuffs()[buffsFieldName],
        setValue: (eventID, raid, newValue) => {
            const newBuffs = raid.getBuffs();
            newBuffs[buffsFieldName] = newValue;
            raid.setBuffs(eventID, newBuffs);
        },
    };
}
function makeTristateRaidBuffInput(id, impId, buffsFieldName) {
    return {
        id: id,
        states: 3,
        improvedId: impId,
        changedEvent: (raid) => raid.buffsChangeEmitter,
        getValue: (raid) => raid.getBuffs()[buffsFieldName],
        setValue: (eventID, raid, newValue) => {
            const newBuffs = raid.getBuffs();
            newBuffs[buffsFieldName] = newValue;
            raid.setBuffs(eventID, newBuffs);
        },
    };
}
function makeBooleanPartyBuffInput(id, buffsFieldName, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (party) => party.buffsChangeEmitter,
        getValue: (party) => party.getBuffs()[buffsFieldName],
        setValue: (eventID, party, newValue) => {
            const newBuffs = party.getBuffs();
            newBuffs[buffsFieldName] = newValue;
            party.setBuffs(eventID, newBuffs);
        },
    };
}
function makeTristatePartyBuffInput(id, impId, buffsFieldName) {
    return {
        id: id,
        states: 3,
        improvedId: impId,
        changedEvent: (party) => party.buffsChangeEmitter,
        getValue: (party) => party.getBuffs()[buffsFieldName],
        setValue: (eventID, party, newValue) => {
            const newBuffs = party.getBuffs();
            newBuffs[buffsFieldName] = newValue;
            party.setBuffs(eventID, newBuffs);
        },
    };
}
function makeMultistatePartyBuffInput(id, numStates, buffsFieldName) {
    return {
        id: id,
        states: numStates,
        changedEvent: (party) => party.buffsChangeEmitter,
        getValue: (party) => party.getBuffs()[buffsFieldName],
        setValue: (eventID, party, newValue) => {
            const newBuffs = party.getBuffs();
            newBuffs[buffsFieldName] = newValue;
            party.setBuffs(eventID, newBuffs);
        },
    };
}
function makeEnumValuePartyBuffInput(id, buffsFieldName, enumValue, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (party) => party.buffsChangeEmitter,
        getValue: (party) => party.getBuffs()[buffsFieldName] == enumValue,
        setValue: (eventID, party, newValue) => {
            const newBuffs = party.getBuffs();
            newBuffs[buffsFieldName] = newValue ? enumValue : 0;
            party.setBuffs(eventID, newBuffs);
        },
    };
}
function makeBooleanIndividualBuffInput(id, buffsFieldName, exclusivityTags) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (player) => player.buffsChangeEmitter,
        getValue: (player) => player.getBuffs()[buffsFieldName],
        setValue: (eventID, player, newValue) => {
            const newBuffs = player.getBuffs();
            newBuffs[buffsFieldName] = newValue;
            player.setBuffs(eventID, newBuffs);
        },
    };
}
function makeTristateIndividualBuffInput(id, impId, buffsFieldName) {
    return {
        id: id,
        states: 3,
        improvedId: impId,
        changedEvent: (player) => player.buffsChangeEmitter,
        getValue: (player) => player.getBuffs()[buffsFieldName],
        setValue: (eventID, player, newValue) => {
            const newBuffs = player.getBuffs();
            newBuffs[buffsFieldName] = newValue;
            player.setBuffs(eventID, newBuffs);
        },
    };
}
function makeMultistateIndividualBuffInput(id, numStates, buffsFieldName) {
    return {
        id: id,
        states: numStates,
        changedEvent: (player) => player.buffsChangeEmitter,
        getValue: (player) => player.getBuffs()[buffsFieldName],
        setValue: (eventID, player, newValue) => {
            const newBuffs = player.getBuffs();
            newBuffs[buffsFieldName] = newValue;
            player.setBuffs(eventID, newBuffs);
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
        setValue: (eventID, target, newValue) => {
            const newDebuffs = target.getDebuffs();
            newDebuffs[debuffsFieldName] = newValue;
            target.setDebuffs(eventID, newDebuffs);
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
        setValue: (eventID, target, newValue) => {
            const newDebuffs = target.getDebuffs();
            newDebuffs[debuffsFieldName] = newValue;
            target.setDebuffs(eventID, newDebuffs);
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
        setValue: (eventID, player, newValue) => {
            const newBuffs = player.getConsumes();
            newBuffs[consumesFieldName] = newValue;
            player.setConsumes(eventID, newBuffs);
        },
    };
}
function makeEnumValueConsumeInput(id, consumesFieldName, enumValue, exclusivityTags, onSet) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (player) => player.consumesChangeEmitter,
        getValue: (player) => player.getConsumes()[consumesFieldName] == enumValue,
        setValue: (eventID, player, newValue) => {
            const newConsumes = player.getConsumes();
            newConsumes[consumesFieldName] = newValue ? enumValue : 0;
            TypedEvent.freezeAll();
            player.setConsumes(eventID, newConsumes);
            if (onSet) {
                onSet(eventID, player, newValue);
            }
            TypedEvent.unfreezeAll();
        },
    };
}
