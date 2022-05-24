import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Alchohol } from '/tbc/core/proto/common.js';
import { BattleElixir } from '/tbc/core/proto/common.js';
import { Explosive } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { GuardianElixir } from '/tbc/core/proto/common.js';
import { Conjured } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { PetFood } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Hunter_Rotation_WeaveType as WeaveType, } from '/tbc/core/proto/hunter.js';
// Keep each section in alphabetical order.
// Raid Buffs
export const ArcaneBrilliance = makeBooleanRaidBuffInput(ActionId.fromSpellId(27127), 'arcaneBrilliance');
export const DivineSpirit = makeTristateRaidBuffInput(ActionId.fromSpellId(25312), ActionId.fromSpellId(33182), 'divineSpirit', ['Spirit']);
export const GiftOfTheWild = makeTristateRaidBuffInput(ActionId.fromSpellId(26991), ActionId.fromSpellId(17055), 'giftOfTheWild');
export const Thorns = makeTristateRaidBuffInput(ActionId.fromSpellId(26992), ActionId.fromSpellId(16840), 'thorns');
export const PowerWordFortitude = makeTristateRaidBuffInput(ActionId.fromSpellId(25389), ActionId.fromSpellId(14767), 'powerWordFortitude');
export const ShadowProtection = makeBooleanRaidBuffInput(ActionId.fromSpellId(39374), 'shadowProtection');
// Party Buffs
export const AtieshMage = makeMultistatePartyBuffInput(ActionId.fromSpellId(28142), 5, 'atieshMage');
export const AtieshWarlock = makeMultistatePartyBuffInput(ActionId.fromSpellId(28143), 5, 'atieshWarlock');
export const Bloodlust = makeMultistatePartyBuffInput(ActionId.fromSpellId(2825), 11, 'bloodlust');
export const BraidedEterniumChain = makeBooleanPartyBuffInput(ActionId.fromSpellId(31025), 'braidedEterniumChain');
export const ChainOfTheTwilightOwl = makeBooleanPartyBuffInput(ActionId.fromSpellId(31035), 'chainOfTheTwilightOwl');
export const CommandingShout = makeTristatePartyBuffInput(ActionId.fromSpellId(469), ActionId.fromSpellId(12861), 'commandingShout');
export const DevotionAura = makeTristatePartyBuffInput(ActionId.fromSpellId(27149), ActionId.fromSpellId(20142), 'devotionAura');
export const DraeneiRacialCaster = makeBooleanPartyBuffInput(ActionId.fromSpellId(28878), 'draeneiRacialCaster');
export const DraeneiRacialMelee = makeBooleanPartyBuffInput(ActionId.fromSpellId(6562), 'draeneiRacialMelee');
export const EyeOfTheNight = makeBooleanPartyBuffInput(ActionId.fromSpellId(31033), 'eyeOfTheNight');
export const FerociousInspiration = makeMultistatePartyBuffInput(ActionId.fromSpellId(34460), 5, 'ferociousInspiration');
export const JadePendantOfBlasting = makeBooleanPartyBuffInput(ActionId.fromSpellId(25607), 'jadePendantOfBlasting');
export const LeaderOfThePack = makeTristatePartyBuffInput(ActionId.fromSpellId(17007), ActionId.fromItemId(32387), 'leaderOfThePack');
export const ManaSpringTotem = makeTristatePartyBuffInput(ActionId.fromSpellId(25570), ActionId.fromSpellId(16208), 'manaSpringTotem');
export const ManaTideTotem = makeMultistatePartyBuffInput(ActionId.fromSpellId(16190), 5, 'manaTideTotems');
export const MoonkinAura = makeTristatePartyBuffInput(ActionId.fromSpellId(24907), ActionId.fromItemId(32387), 'moonkinAura');
export const RetributionAura = makeTristatePartyBuffInput(ActionId.fromSpellId(27150), ActionId.fromSpellId(20092), 'retributionAura');
export const SanctityAura = makeTristatePartyBuffInput(ActionId.fromSpellId(20218), ActionId.fromSpellId(31870), 'sanctityAura');
export const TotemOfWrath = makeMultistatePartyBuffInput(ActionId.fromSpellId(30706), 5, 'totemOfWrath');
export const TrueshotAura = makeBooleanPartyBuffInput(ActionId.fromSpellId(27066), 'trueshotAura');
export const WrathOfAirTotem = makeTristatePartyBuffInput(ActionId.fromSpellId(3738), ActionId.fromSpellId(37212), 'wrathOfAirTotem');
export const BloodPact = makeTristatePartyBuffInput(ActionId.fromSpellId(27268), ActionId.fromSpellId(18696), 'bloodPact');
export const DrumsOfBattleBuff = makeEnumValuePartyBuffInput(ActionId.fromItemId(185848), 'drums', Drums.DrumsOfBattle, ['Drums']);
export const DrumsOfRestorationBuff = makeEnumValuePartyBuffInput(ActionId.fromItemId(185850), 'drums', Drums.DrumsOfRestoration, ['Drums']);
// Individual Buffs
export const BlessingOfKings = makeBooleanIndividualBuffInput(ActionId.fromSpellId(25898), 'blessingOfKings');
export const BlessingOfMight = makeTristateIndividualBuffInput(ActionId.fromSpellId(27140), ActionId.fromSpellId(20048), 'blessingOfMight');
export const BlessingOfSalvation = makeBooleanIndividualBuffInput(ActionId.fromSpellId(25895), 'blessingOfSalvation');
export const BlessingOfSanctuary = makeBooleanIndividualBuffInput(ActionId.fromSpellId(27169), 'blessingOfSanctuary');
export const BlessingOfWisdom = makeTristateIndividualBuffInput(ActionId.fromSpellId(27143), ActionId.fromSpellId(20245), 'blessingOfWisdom');
export const Innervate = makeMultistateIndividualBuffInput(ActionId.fromSpellId(29166), 11, 'innervates');
export const PowerInfusion = makeMultistateIndividualBuffInput(ActionId.fromSpellId(10060), 11, 'powerInfusions');
export const UnleashedRage = makeBooleanIndividualBuffInput(ActionId.fromSpellId(30811), 'unleashedRage');
// Debuffs
export const BloodFrenzy = makeBooleanDebuffInput(ActionId.fromSpellId(29859), 'bloodFrenzy');
export const HuntersMark = makeTristateDebuffInput(ActionId.fromSpellId(14325), ActionId.fromSpellId(19425), 'huntersMark');
export const ImprovedScorch = makeBooleanDebuffInput(ActionId.fromSpellId(12873), 'improvedScorch');
export const ImprovedSealOfTheCrusader = makeBooleanDebuffInput(ActionId.fromSpellId(20337), 'improvedSealOfTheCrusader');
export const JudgementOfWisdom = makeBooleanDebuffInput(ActionId.fromSpellId(27164), 'judgementOfWisdom');
export const JudgementOfLight = makeBooleanDebuffInput(ActionId.fromSpellId(27163), 'judgementOfLight');
export const Mangle = makeBooleanDebuffInput(ActionId.fromSpellId(33876), 'mangle');
export const Misery = makeBooleanDebuffInput(ActionId.fromSpellId(33195), 'misery');
export const ShadowWeaving = makeBooleanDebuffInput(ActionId.fromSpellId(15334), 'shadowWeaving');
export const CurseOfElements = makeTristateDebuffInput(ActionId.fromSpellId(27228), ActionId.fromSpellId(32484), 'curseOfElements');
export const CurseOfRecklessness = makeBooleanDebuffInput(ActionId.fromSpellId(27226), 'curseOfRecklessness');
export const FaerieFire = makeTristateDebuffInput(ActionId.fromSpellId(26993), ActionId.fromSpellId(33602), 'faerieFire');
export const ExposeArmor = makeTristateDebuffInput(ActionId.fromSpellId(26866), ActionId.fromSpellId(14169), 'exposeArmor');
export const SunderArmor = makeBooleanDebuffInput(ActionId.fromSpellId(25225), 'sunderArmor');
export const WintersChill = makeBooleanDebuffInput(ActionId.fromSpellId(28595), 'wintersChill');
export const GiftOfArthas = makeBooleanDebuffInput(ActionId.fromSpellId(11374), 'giftOfArthas');
export const DemoralizingRoar = makeTristateDebuffInput(ActionId.fromSpellId(26998), ActionId.fromSpellId(16862), 'demoralizingRoar');
export const DemoralizingShout = makeTristateDebuffInput(ActionId.fromSpellId(25203), ActionId.fromSpellId(12879), 'demoralizingShout');
export const ThunderClap = makeTristateDebuffInput(ActionId.fromSpellId(25264), ActionId.fromSpellId(12666), 'thunderClap');
export const InsectSwarm = makeBooleanDebuffInput(ActionId.fromSpellId(27013), 'insectSwarm');
export const ScorpidSting = makeBooleanDebuffInput(ActionId.fromSpellId(3043), 'scorpidSting');
export const ShadowEmbrace = makeBooleanDebuffInput(ActionId.fromSpellId(32394), 'shadowEmbrace');
// Consumes
export const SuperSapper = makeBooleanConsumeInput(ActionId.fromItemId(23827), 'superSapper', [], onSetExplosives);
export const GoblinSapper = makeBooleanConsumeInput(ActionId.fromItemId(10646), 'goblinSapper', [], onSetExplosives);
export const KiblersBits = makeEnumValueConsumeInput(ActionId.fromItemId(33874), 'petFood', PetFood.PetFoodKiblersBits, ['Pet Food']);
export const ScrollOfAgilityV = makeEnumValueConsumeInput(ActionId.fromItemId(27498), 'scrollOfAgility', 5);
export const ScrollOfSpiritV = makeEnumValueConsumeInput(ActionId.fromItemId(27501), 'scrollOfSpirit', 5, ['Spirit']);
export const ScrollOfStrengthV = makeEnumValueConsumeInput(ActionId.fromItemId(27503), 'scrollOfStrength', 5);
export const ScrollOfProtectionV = makeEnumValueConsumeInput(ActionId.fromItemId(27500), 'scrollOfProtection', 5);
export const PetScrollOfAgilityV = makeEnumValueConsumeInput(ActionId.fromItemId(27498), 'petScrollOfAgility', 5);
export const PetScrollOfStrengthV = makeEnumValueConsumeInput(ActionId.fromItemId(27503), 'petScrollOfStrength', 5);
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
function makeTristateRaidBuffInput(id, impId, buffsFieldName, exclusivityTags) {
    return {
        id: id,
        states: 3,
        improvedId: impId,
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
        extraCssClasses: buffsFieldName == 'blessingOfSalvation' ? ['threat-metrics'] : [],
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
        changedEvent: (raid) => raid.debuffsChangeEmitter,
        getValue: (raid) => raid.getDebuffs()[debuffsFieldName],
        setValue: (eventID, raid, newValue) => {
            const newDebuffs = raid.getDebuffs();
            newDebuffs[debuffsFieldName] = newValue;
            raid.setDebuffs(eventID, newDebuffs);
        },
    };
}
function makeTristateDebuffInput(id, impId, debuffsFieldName) {
    return {
        id: id,
        states: 3,
        improvedId: impId,
        changedEvent: (raid) => raid.debuffsChangeEmitter,
        getValue: (raid) => raid.getDebuffs()[debuffsFieldName],
        setValue: (eventID, raid, newValue) => {
            const newDebuffs = raid.getDebuffs();
            newDebuffs[debuffsFieldName] = newValue;
            raid.setDebuffs(eventID, newDebuffs);
        },
    };
}
function makeBooleanConsumeInput(id, consumesFieldName, exclusivityTags, onSet) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (player) => player.consumesChangeEmitter,
        getValue: (player) => player.getConsumes()[consumesFieldName],
        setValue: (eventID, player, newValue) => {
            const newConsumes = player.getConsumes();
            newConsumes[consumesFieldName] = newValue;
            TypedEvent.freezeAllAndDo(() => {
                player.setConsumes(eventID, newConsumes);
                if (onSet) {
                    onSet(eventID, player, newValue);
                }
            });
        },
    };
}
function makeEnumValueConsumeInput(id, consumesFieldName, enumValue, exclusivityTags, onSet, showWhen) {
    return {
        id: id,
        states: 2,
        exclusivityTags: exclusivityTags,
        changedEvent: (player) => player.consumesChangeEmitter,
        getValue: (player) => player.getConsumes()[consumesFieldName] == enumValue,
        setValue: (eventID, player, newValue) => {
            const newConsumes = player.getConsumes();
            newConsumes[consumesFieldName] = newValue ? enumValue : 0;
            TypedEvent.freezeAllAndDo(() => {
                player.setConsumes(eventID, newConsumes);
                if (onSet) {
                    onSet(eventID, player, newValue);
                }
            });
        },
        showWhen: showWhen,
    };
}
//////////////////////////////////////////////////////////////////////
// Custom buffs that don't fit into any of the helper functions above.
//////////////////////////////////////////////////////////////////////
export const GraceOfAirTotem = {
    id: ActionId.fromSpellId(25359),
    states: 3,
    improvedId: ActionId.fromSpellId(16295),
    changedEvent: (party) => party.buffsChangeEmitter,
    getValue: (party) => party.getBuffs().graceOfAirTotem,
    setValue: (eventID, party, newValue) => {
        const newBuffs = party.getBuffs();
        newBuffs.graceOfAirTotem = newValue;
        party.setBuffs(eventID, newBuffs);
    },
};
export const StrengthOfEarthTotem = {
    id: ActionId.fromSpellId(25528),
    states: 4,
    improvedId: ActionId.fromSpellId(16295),
    improvedId2: ActionId.fromSpellId(37223),
    changedEvent: (party) => party.buffsChangeEmitter,
    getValue: (party) => party.getBuffs().strengthOfEarthTotem > 2 ? party.getBuffs().strengthOfEarthTotem - 1 : party.getBuffs().strengthOfEarthTotem,
    setValue: (eventID, party, newValue) => {
        const newBuffs = party.getBuffs();
        // Skip cyclone-only value.
        newBuffs.strengthOfEarthTotem = newValue > 1 ? newValue + 1 : newValue;
        party.setBuffs(eventID, newBuffs);
    },
};
export const WindfuryTotem = {
    id: ActionId.fromSpellId(25587),
    states: 3,
    improvedId: ActionId.fromSpellId(29193),
    changedEvent: (party) => party.buffsChangeEmitter,
    getValue: (party) => {
        const buffs = party.getBuffs();
        if (buffs.windfuryTotemRank == 0) {
            return 0;
        }
        if (buffs.windfuryTotemIwt > 0) {
            return 2;
        }
        else {
            return 1;
        }
    },
    setValue: (eventID, party, newValue) => {
        const newBuffs = party.getBuffs();
        if (newValue == 0) {
            newBuffs.windfuryTotemRank = 0;
            newBuffs.windfuryTotemIwt = 0;
        }
        else {
            newBuffs.windfuryTotemRank = 5;
            if (newValue == 2) {
                newBuffs.windfuryTotemIwt = 2;
            }
            else {
                newBuffs.windfuryTotemIwt = 0;
            }
        }
        party.setBuffs(eventID, newBuffs);
    },
};
export const BattleShout = {
    id: ActionId.fromSpellId(2048),
    states: 4,
    improvedId: ActionId.fromSpellId(12861),
    improvedId2: ActionId.fromItemId(30446),
    changedEvent: (party) => party.buffsChangeEmitter,
    getValue: (party) => {
        const buffs = party.getBuffs();
        if (buffs.battleShout == TristateEffect.TristateEffectImproved) {
            return buffs.battleShout + Number(buffs.bsSolarianSapphire);
        }
        else {
            return buffs.battleShout;
        }
    },
    setValue: (eventID, party, newValue) => {
        const newBuffs = party.getBuffs();
        newBuffs.battleShout = Math.min(2, newValue);
        newBuffs.bsSolarianSapphire = newValue == 3;
        party.setBuffs(eventID, newBuffs);
    },
};
export const makePotionsInput = makeConsumeInputFactory('defaultPotion', [
    { actionId: ActionId.fromItemId(22839), value: Potions.DestructionPotion },
    { actionId: ActionId.fromItemId(22838), value: Potions.HastePotion },
    { actionId: ActionId.fromItemId(13442), value: Potions.MightyRagePotion },
    { actionId: ActionId.fromItemId(22832), value: Potions.SuperManaPotion },
    { actionId: ActionId.fromItemId(31677), value: Potions.FelManaPotion },
    { actionId: ActionId.fromItemId(22828), value: Potions.InsaneStrengthPotion },
    { actionId: ActionId.fromItemId(22849), value: Potions.IronshieldPotion },
    { actionId: ActionId.fromItemId(22837), value: Potions.HeroicPotion },
]);
export const makeConjuredInput = makeConsumeInputFactory('defaultConjured', [
    { actionId: ActionId.fromItemId(12662), value: Conjured.ConjuredDarkRune },
    { actionId: ActionId.fromItemId(22788), value: Conjured.ConjuredFlameCap },
    { actionId: ActionId.fromItemId(22044), value: Conjured.ConjuredMageManaEmerald },
    { actionId: ActionId.fromItemId(7676), value: Conjured.ConjuredRogueThistleTea },
]);
export const makeFlasksInput = makeConsumeInputFactory('flask', [
    { actionId: ActionId.fromItemId(22861), value: Flask.FlaskOfBlindingLight },
    { actionId: ActionId.fromItemId(22853), value: Flask.FlaskOfMightyRestoration },
    { actionId: ActionId.fromItemId(22866), value: Flask.FlaskOfPureDeath },
    { actionId: ActionId.fromItemId(22854), value: Flask.FlaskOfRelentlessAssault },
    { actionId: ActionId.fromItemId(13512), value: Flask.FlaskOfSupremePower },
    { actionId: ActionId.fromItemId(22851), value: Flask.FlaskOfFortification },
], (eventID, player, newValue) => {
    if (newValue) {
        const newConsumes = player.getConsumes();
        newConsumes.battleElixir = BattleElixir.BattleElixirUnknown;
        newConsumes.guardianElixir = GuardianElixir.GuardianElixirUnknown;
        player.setConsumes(eventID, newConsumes);
    }
});
export const makeBattleElixirsInput = makeConsumeInputFactory('battleElixir', [
    { actionId: ActionId.fromItemId(28103), value: BattleElixir.AdeptsElixir },
    { actionId: ActionId.fromItemId(9224), value: BattleElixir.ElixirOfDemonslaying },
    { actionId: ActionId.fromItemId(22831), value: BattleElixir.ElixirOfMajorAgility },
    { actionId: ActionId.fromItemId(22833), value: BattleElixir.ElixirOfMajorFirePower },
    { actionId: ActionId.fromItemId(22827), value: BattleElixir.ElixirOfMajorFrostPower },
    { actionId: ActionId.fromItemId(22835), value: BattleElixir.ElixirOfMajorShadowPower },
    { actionId: ActionId.fromItemId(22824), value: BattleElixir.ElixirOfMajorStrength },
    { actionId: ActionId.fromItemId(28104), value: BattleElixir.ElixirOfMastery },
    { actionId: ActionId.fromItemId(13452), value: BattleElixir.ElixirOfTheMongoose },
    { actionId: ActionId.fromItemId(31679), value: BattleElixir.FelStrengthElixir },
], (eventID, player, newValue) => {
    if (newValue) {
        const newConsumes = player.getConsumes();
        newConsumes.flask = Flask.FlaskUnknown;
        player.setConsumes(eventID, newConsumes);
    }
});
export const makeGuardianElixirsInput = makeConsumeInputFactory('guardianElixir', [
    { actionId: ActionId.fromItemId(9088), value: GuardianElixir.GiftOfArthas },
    { actionId: ActionId.fromItemId(32067), value: GuardianElixir.ElixirOfDraenicWisdom },
    { actionId: ActionId.fromItemId(32068), value: GuardianElixir.ElixirOfIronskin },
    { actionId: ActionId.fromItemId(22834), value: GuardianElixir.ElixirOfMajorDefense },
    { actionId: ActionId.fromItemId(32062), value: GuardianElixir.ElixirOfMajorFortitude },
    { actionId: ActionId.fromItemId(22840), value: GuardianElixir.ElixirOfMajorMageblood },
], (eventID, player, newValue) => {
    if (newValue) {
        const newConsumes = player.getConsumes();
        newConsumes.flask = Flask.FlaskUnknown;
        player.setConsumes(eventID, newConsumes);
    }
});
export const makeFoodInput = makeConsumeInputFactory('food', [
    { actionId: ActionId.fromItemId(27657), value: Food.FoodBlackenedBasilisk },
    { actionId: ActionId.fromItemId(27664), value: Food.FoodGrilledMudfish },
    { actionId: ActionId.fromItemId(27655), value: Food.FoodRavagerDog },
    { actionId: ActionId.fromItemId(27658), value: Food.FoodRoastedClefthoof },
    { actionId: ActionId.fromItemId(33872), value: Food.FoodSpicyHotTalbuk },
    { actionId: ActionId.fromItemId(33825), value: Food.FoodSkullfishSoup },
    { actionId: ActionId.fromItemId(33052), value: Food.FoodFishermansFeast },
]);
export const makeAlcoholInput = makeConsumeInputFactory('alchohol', [
    { actionId: ActionId.fromItemId(18284), value: Alchohol.AlchoholKreegsStoutBeatdown },
]);
export const makePetFoodInput = makeConsumeInputFactory('petFood', [
    { actionId: ActionId.fromItemId(33874), value: PetFood.PetFoodKiblersBits },
]);
function onSetDrums(eventID, player, newValue) {
    if (newValue) {
        const playerConsumes = player.getConsumes();
        playerConsumes.superSapper = false;
        playerConsumes.goblinSapper = false;
        playerConsumes.fillerExplosive = Explosive.ExplosiveUnknown;
        player.setConsumes(eventID, playerConsumes);
        player.getOtherPartyMembers().forEach(otherPlayer => {
            const otherConsumes = otherPlayer.getConsumes();
            otherConsumes.drums = Drums.DrumsUnknown;
            otherPlayer.setConsumes(eventID, otherConsumes);
        });
    }
}
;
export const DrumsInput = makeConsumeInput('drums', [
    { actionId: ActionId.fromItemId(185848), value: Drums.DrumsOfBattle },
    { actionId: ActionId.fromItemId(185850), value: Drums.DrumsOfRestoration },
    { actionId: ActionId.fromItemId(185852), value: Drums.DrumsOfWar },
], onSetDrums);
function onSetExplosives(eventID, player, newValue) {
    if (newValue) {
        const playerConsumes = player.getConsumes();
        playerConsumes.drums = Drums.DrumsUnknown;
        player.setConsumes(eventID, playerConsumes);
    }
}
;
export const FillerExplosiveInput = makeConsumeInput('fillerExplosive', [
    { actionId: ActionId.fromItemId(23736), value: Explosive.ExplosiveFelIronBomb },
    { actionId: ActionId.fromItemId(23737), value: Explosive.ExplosiveAdamantiteGrenade },
    { actionId: ActionId.fromItemId(23841), value: Explosive.ExplosiveGnomishFlameTurret },
    { actionId: ActionId.fromItemId(13180), value: Explosive.ExplosiveHolyWater },
], onSetExplosives);
export function makeWeaponImbueInput(isMainHand, options) {
    const allOptions = [
        { actionId: ActionId.fromItemId(18262), value: WeaponImbue.WeaponImbueElementalSharpeningStone },
        { actionId: ActionId.fromItemId(20749), value: WeaponImbue.WeaponImbueBrilliantWizardOil },
        { actionId: ActionId.fromItemId(22522), value: WeaponImbue.WeaponImbueSuperiorWizardOil },
        { actionId: ActionId.fromItemId(34539), value: WeaponImbue.WeaponImbueRighteousWeaponCoating },
        {
            actionId: ActionId.fromItemId(23529), value: WeaponImbue.WeaponImbueAdamantiteSharpeningStone,
            showWhen: (player) => !(isMainHand ? player.getGear().hasBluntMHWeapon() : player.getGear().hasBluntOHWeapon()),
        },
        {
            actionId: ActionId.fromItemId(28421), value: WeaponImbue.WeaponImbueAdamantiteWeightstone,
            showWhen: (player) => (isMainHand ? player.getGear().hasBluntMHWeapon() : player.getGear().hasBluntOHWeapon()),
        },
        { actionId: ActionId.fromSpellId(27186), value: WeaponImbue.WeaponImbueRogueDeadlyPoison },
        { actionId: ActionId.fromSpellId(26891), value: WeaponImbue.WeaponImbueRogueInstantPoison },
        { actionId: ActionId.fromSpellId(25505), value: WeaponImbue.WeaponImbueShamanWindfury },
        { actionId: ActionId.fromSpellId(25489), value: WeaponImbue.WeaponImbueShamanFlametongue },
        { actionId: ActionId.fromSpellId(25500), value: WeaponImbue.WeaponImbueShamanFrostbrand },
        { actionId: ActionId.fromSpellId(25485), value: WeaponImbue.WeaponImbueShamanRockbiter },
    ];
    if (isMainHand) {
        const config = makeConsumeInputFactory('mainHandImbue', allOptions)(options);
        config.enableWhen = (player) => !player.getParty()
            || player.getParty().getBuffs().windfuryTotemRank == 0
            || (player.spec == Spec.SpecHunter && player.getRotation().weave == WeaveType.WeaveNone);
        config.changedEvent = (player) => TypedEvent.onAny([player.getRaid()?.changeEmitter || player.consumesChangeEmitter]);
        return config;
    }
    else {
        return makeConsumeInputFactory('offHandImbue', allOptions)(options);
    }
}
function makeConsumeInputFactory(consumesFieldName, allOptions, onSet) {
    return (options) => {
        return {
            numColumns: 1,
            values: [
                { color: 'grey', value: 0 },
            ].concat(options.map(option => allOptions.find(allOption => allOption.value == option))),
            equals: (a, b) => a == b,
            zeroValue: 0,
            changedEvent: (player) => player.consumesChangeEmitter,
            getValue: (player) => player.getConsumes()[consumesFieldName],
            setValue: (eventID, player, newValue) => {
                const newConsumes = player.getConsumes();
                newConsumes[consumesFieldName] = newValue;
                TypedEvent.freezeAllAndDo(() => {
                    player.setConsumes(eventID, newConsumes);
                    if (onSet) {
                        onSet(eventID, player, newValue);
                    }
                });
            },
        };
    };
}
function makeConsumeInput(consumesFieldName, allOptions, onSet) {
    const factory = makeConsumeInputFactory(consumesFieldName, allOptions, onSet);
    return factory(allOptions.map(option => option.value));
}
