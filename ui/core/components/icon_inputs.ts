import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Alchohol} from '/tbc/core/proto/common.js';
import { BattleElixir } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { GuardianElixir } from '/tbc/core/proto/common.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Conjured } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { PetFood } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { IndividualSimIconPickerConfig } from '/tbc/core/individual_sim_ui.js';
import { Party } from '/tbc/core/party.js';
import { Player } from '/tbc/core/player.js';
import { Raid } from '/tbc/core/raid.js';
import { Sim } from '/tbc/core/sim.js';
import { Target } from '/tbc/core/target.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

import { ExclusivityTag } from '/tbc/core/individual_sim_ui.js';
import { IconPickerConfig } from './icon_picker.js';

// Keep each section in alphabetical order.

// Raid Buffs
export const ArcaneBrilliance = makeBooleanRaidBuffInput(ActionId.fromSpellId(27127), 'arcaneBrilliance');
export const DivineSpirit = makeTristateRaidBuffInput(ActionId.fromSpellId(25312), ActionId.fromSpellId(33182), 'divineSpirit', ['Spirit']);
export const GiftOfTheWild = makeTristateRaidBuffInput(ActionId.fromSpellId(26991), ActionId.fromSpellId(17055), 'giftOfTheWild');

// Party Buffs
export const AtieshMage = makeMultistatePartyBuffInput(ActionId.fromSpellId(28142), 5, 'atieshMage');
export const AtieshWarlock = makeMultistatePartyBuffInput(ActionId.fromSpellId(28143), 5, 'atieshWarlock');
export const BattleChickens = makeMultistatePartyBuffInput(ActionId.fromItemId(10725), 5, 'battleChickens');
export const Bloodlust = makeMultistatePartyBuffInput(ActionId.fromSpellId(2825), 11, 'bloodlust');
export const BraidedEterniumChain = makeBooleanPartyBuffInput(ActionId.fromSpellId(31025), 'braidedEterniumChain');
export const ChainOfTheTwilightOwl = makeBooleanPartyBuffInput(ActionId.fromSpellId(31035), 'chainOfTheTwilightOwl');
export const DraeneiRacialCaster = makeBooleanPartyBuffInput(ActionId.fromSpellId(28878), 'draeneiRacialCaster');
export const DraeneiRacialMelee = makeBooleanPartyBuffInput(ActionId.fromSpellId(6562), 'draeneiRacialMelee');
export const EyeOfTheNight = makeBooleanPartyBuffInput(ActionId.fromSpellId(31033), 'eyeOfTheNight');
export const FerociousInspiration = makeMultistatePartyBuffInput(ActionId.fromSpellId(34460), 5, 'ferociousInspiration');
export const JadePendantOfBlasting = makeBooleanPartyBuffInput(ActionId.fromSpellId(25607), 'jadePendantOfBlasting');
export const LeaderOfThePack = makeTristatePartyBuffInput(ActionId.fromSpellId(17007), ActionId.fromItemId(32387), 'leaderOfThePack');
export const ManaSpringTotem = makeTristatePartyBuffInput(ActionId.fromSpellId(25570), ActionId.fromSpellId(16208), 'manaSpringTotem');
export const ManaTideTotem = makeMultistatePartyBuffInput(ActionId.fromSpellId(16190), 5, 'manaTideTotems');
export const MoonkinAura = makeTristatePartyBuffInput(ActionId.fromSpellId(24907), ActionId.fromItemId(32387), 'moonkinAura');
export const SanctityAura = makeTristatePartyBuffInput(ActionId.fromSpellId(20218), ActionId.fromSpellId(31870), 'sanctityAura');
export const TotemOfWrath = makeMultistatePartyBuffInput(ActionId.fromSpellId(30706), 5, 'totemOfWrath');
export const TrueshotAura = makeBooleanPartyBuffInput(ActionId.fromSpellId(27066), 'trueshotAura');
export const WrathOfAirTotem = makeTristatePartyBuffInput(ActionId.fromSpellId(3738), ActionId.fromSpellId(37212), 'wrathOfAirTotem');

export const DrumsOfBattleBuff = makeEnumValuePartyBuffInput(ActionId.fromSpellId(35476), 'drums', Drums.DrumsOfBattle, ['Drums']);
export const DrumsOfRestorationBuff = makeEnumValuePartyBuffInput(ActionId.fromSpellId(35478), 'drums', Drums.DrumsOfRestoration, ['Drums']);

// Individual Buffs
export const BlessingOfKings = makeBooleanIndividualBuffInput(ActionId.fromSpellId(25898), 'blessingOfKings');
export const BlessingOfWisdom = makeTristateIndividualBuffInput(ActionId.fromSpellId(27143), ActionId.fromSpellId(20245), 'blessingOfWisdom');
export const BlessingOfMight = makeTristateIndividualBuffInput(ActionId.fromSpellId(27140), ActionId.fromSpellId(20048), 'blessingOfMight');
export const Innervate = makeMultistateIndividualBuffInput(ActionId.fromSpellId(29166), 11, 'innervates');
export const PowerInfusion = makeMultistateIndividualBuffInput(ActionId.fromSpellId(10060), 11, 'powerInfusions');
export const UnleashedRage = makeBooleanIndividualBuffInput(ActionId.fromSpellId(30811), 'unleashedRage');

// Debuffs
export const BloodFrenzy = makeBooleanDebuffInput(ActionId.fromSpellId(29859), 'bloodFrenzy');
export const HuntersMark = makeTristateDebuffInput(ActionId.fromSpellId(14325), ActionId.fromSpellId(19425), 'huntersMark');
export const ImprovedScorch = makeBooleanDebuffInput(ActionId.fromSpellId(12873), 'improvedScorch');
export const ImprovedSealOfTheCrusader = makeBooleanDebuffInput(ActionId.fromSpellId(20337), 'improvedSealOfTheCrusader');
export const JudgementOfWisdom = makeBooleanDebuffInput(ActionId.fromSpellId(27164), 'judgementOfWisdom');
export const Misery = makeBooleanDebuffInput(ActionId.fromSpellId(33195), 'misery');
export const CurseOfElements = makeTristateDebuffInput(ActionId.fromSpellId(27228), ActionId.fromSpellId(32484), 'curseOfElements');
export const CurseOfRecklessness = makeBooleanDebuffInput(ActionId.fromSpellId(27226), 'curseOfRecklessness');
export const FaerieFire = makeTristateDebuffInput(ActionId.fromSpellId(26993), ActionId.fromSpellId(33602), 'faerieFire');
export const ExposeArmor = makeTristateDebuffInput(ActionId.fromSpellId(26866), ActionId.fromSpellId(14169), 'exposeArmor');
export const SunderArmor = makeBooleanDebuffInput(ActionId.fromSpellId(25225), 'sunderArmor');
export const WintersChill = makeBooleanDebuffInput(ActionId.fromSpellId(28595), 'wintersChill');

// Consumes
export const BattleChicken = makeBooleanConsumeInput(ActionId.fromItemId(10725), 'battleChicken');

export const FlaskOfBlindingLight = makeEnumValueConsumeInput(ActionId.fromItemId(22861), 'flask', Flask.FlaskOfBlindingLight, ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfMightyRestoration = makeEnumValueConsumeInput(ActionId.fromItemId(22853), 'flask', Flask.FlaskOfMightyRestoration, ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfPureDeath = makeEnumValueConsumeInput(ActionId.fromItemId(22866), 'flask', Flask.FlaskOfPureDeath, ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfRelentlessAssault = makeEnumValueConsumeInput(ActionId.fromItemId(22854), 'flask', Flask.FlaskOfRelentlessAssault, ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfSupremePower = makeEnumValueConsumeInput(ActionId.fromItemId(13512), 'flask', Flask.FlaskOfSupremePower, ['Battle Elixir', 'Guardian Elixir']);

export const AdeptsElixir = makeEnumValueConsumeInput(ActionId.fromItemId(28103), 'battleElixir', BattleElixir.AdeptsElixir, ['Battle Elixir']);
export const ElixirOfDemonslaying = makeEnumValueConsumeInput(ActionId.fromItemId(9224), 'battleElixir', BattleElixir.ElixirOfDemonslaying, ['Battle Elixir']);
export const ElixirOfMajorAgility = makeEnumValueConsumeInput(ActionId.fromItemId(22831), 'battleElixir', BattleElixir.ElixirOfMajorAgility, ['Battle Elixir']);
export const ElixirOfMajorFirePower = makeEnumValueConsumeInput(ActionId.fromItemId(22833), 'battleElixir', BattleElixir.ElixirOfMajorFirePower, ['Battle Elixir']);
export const ElixirOfMajorFrostPower = makeEnumValueConsumeInput(ActionId.fromItemId(22827), 'battleElixir', BattleElixir.ElixirOfMajorFrostPower, ['Battle Elixir']);
export const ElixirOfMajorShadowPower = makeEnumValueConsumeInput(ActionId.fromItemId(22835), 'battleElixir', BattleElixir.ElixirOfMajorShadowPower, ['Battle Elixir']);
export const ElixirOfMajorStrength = makeEnumValueConsumeInput(ActionId.fromItemId(22824), 'battleElixir', BattleElixir.ElixirOfMajorStrength, ['Battle Elixir']);
export const ElixirOfTheMongoose = makeEnumValueConsumeInput(ActionId.fromItemId(13452), 'battleElixir', BattleElixir.ElixirOfTheMongoose, ['Battle Elixir']);

export const ElixirOfDraenicWisdom = makeEnumValueConsumeInput(ActionId.fromItemId(32067), 'guardianElixir', GuardianElixir.ElixirOfDraenicWisdom, ['Guardian Elixir']);
export const ElixirOfMajorMageblood = makeEnumValueConsumeInput(ActionId.fromItemId(22840), 'guardianElixir', GuardianElixir.ElixirOfMajorMageblood, ['Guardian Elixir']);

export const MainHandElementalSharpeningStone = makeEnumValueConsumeInput(ActionId.fromItemId(18262), 'mainHandImbue', WeaponImbue.WeaponImbueElementalSharpeningStone, ['MH Weapon Imbue']);
export const MainHandBrilliantWizardOil = makeEnumValueConsumeInput(ActionId.fromItemId(20749), 'mainHandImbue', WeaponImbue.WeaponImbueBrilliantWizardOil, ['MH Weapon Imbue']);
export const MainHandSuperiorWizardOil = makeEnumValueConsumeInput(ActionId.fromItemId(22522), 'mainHandImbue', WeaponImbue.WeaponImbueSuperiorWizardOil, ['MH Weapon Imbue']);
export const OffHandElementalSharpeningStone = makeEnumValueConsumeInput(ActionId.fromItemId(18262), 'offHandImbue', WeaponImbue.WeaponImbueElementalSharpeningStone, ['OH Weapon Imbue']);
export const MainHandAdamantiteSharpeningStone = makeAdamantiteStoneInput(ActionId.fromItemId(23529), true, true, WeaponImbue.WeaponImbueAdamantiteSharpeningStone);
export const MainHandAdamantiteWeightstone = makeAdamantiteStoneInput(ActionId.fromItemId(28421), false, true, WeaponImbue.WeaponImbueAdamantiteWeightstone);
export const OffHandAdamantiteSharpeningStone = makeAdamantiteStoneInput(ActionId.fromItemId(23529), true, false, WeaponImbue.WeaponImbueAdamantiteSharpeningStone);
export const OffHandAdamantiteWeightstone = makeAdamantiteStoneInput(ActionId.fromItemId(28421), false, false, WeaponImbue.WeaponImbueAdamantiteWeightstone);

export const BlackenedBasilisk = makeEnumValueConsumeInput(ActionId.fromItemId(27657), 'food', Food.FoodBlackenedBasilisk, ['Food']);
export const GrilledMudfish = makeEnumValueConsumeInput(ActionId.fromItemId(27664), 'food', Food.FoodGrilledMudfish, ['Food']);
export const RavagerDog = makeEnumValueConsumeInput(ActionId.fromItemId(27655), 'food', Food.FoodRavagerDog, ['Food']);
export const RoastedClefthoof = makeEnumValueConsumeInput(ActionId.fromItemId(27658), 'food', Food.FoodRoastedClefthoof, ['Food']);
export const SpicyHotTalbuk = makeEnumValueConsumeInput(ActionId.fromItemId(33872), 'food', Food.FoodSpicyHotTalbuk, ['Food']);
export const SkullfishSoup = makeEnumValueConsumeInput(ActionId.fromItemId(33825), 'food', Food.FoodSkullfishSoup, ['Food']);

export const KiblersBits = makeEnumValueConsumeInput(ActionId.fromItemId(33874), 'petFood', PetFood.PetFoodKiblersBits, ['Pet Food']);

export const KreegsStoutBeatdown = makeEnumValueConsumeInput(ActionId.fromItemId(18284), 'alchohol', Alchohol.AlchoholKreegsStoutBeatdown, ['Alchohol']);

export const DefaultDestructionPotion = makeEnumValueConsumeInput(ActionId.fromItemId(22839), 'defaultPotion', Potions.DestructionPotion, ['Potion']);
export const DefaultHastePotion = makeEnumValueConsumeInput(ActionId.fromItemId(22838), 'defaultPotion', Potions.HastePotion, ['Potion']);
export const DefaultMightyRagePotion = makeEnumValueConsumeInput(ActionId.fromItemId(13442), 'defaultPotion', Potions.MightyRagePotion, ['Potion']);
export const DefaultSuperManaPotion = makeEnumValueConsumeInput(ActionId.fromItemId(22832), 'defaultPotion', Potions.SuperManaPotion, ['Potion']);
export const DefaultFelManaPotion = makeEnumValueConsumeInput(ActionId.fromItemId(31677), 'defaultPotion', Potions.FelManaPotion, ['Potion']);

export const DefaultDarkRune = makeEnumValueConsumeInput(ActionId.fromItemId(12662), 'defaultConjured', Conjured.ConjuredDarkRune, ['Conjured']);
export const DefaultFlameCap = makeEnumValueConsumeInput(ActionId.fromItemId(22788), 'defaultConjured', Conjured.ConjuredFlameCap, ['Conjured']);
export const DefaultMageManaEmerald = makeEnumValueConsumeInput(ActionId.fromItemId(22044), 'defaultConjured', Conjured.ConjuredMageManaEmerald, ['Conjured']);

export const ScrollOfAgilityV = makeEnumValueConsumeInput(ActionId.fromItemId(27498), 'scrollOfAgility', 5);
export const ScrollOfSpiritV = makeEnumValueConsumeInput(ActionId.fromItemId(27501), 'scrollOfSpirit', 5, ['Spirit']);
export const ScrollOfStrengthV = makeEnumValueConsumeInput(ActionId.fromItemId(27503), 'scrollOfStrength', 5);

export const PetScrollOfAgilityV = makeEnumValueConsumeInput(ActionId.fromItemId(27498), 'petScrollOfAgility', 5);
export const PetScrollOfStrengthV = makeEnumValueConsumeInput(ActionId.fromItemId(27503), 'petScrollOfStrength', 5);

function removeOtherPartyMembersDrums(eventID: EventID, player: Player<any>, newValue: boolean) {
	if (newValue) {
		player.getOtherPartyMembers().forEach(otherPlayer => {
			const otherConsumes = otherPlayer.getConsumes();
			otherConsumes.drums = Drums.DrumsUnknown;
			otherPlayer.setConsumes(eventID, otherConsumes);
		});
	}
};
export const DrumsOfBattleConsume = makeEnumValueConsumeInput(ActionId.fromSpellId(35476), 'drums', Drums.DrumsOfBattle, ['Drums'], removeOtherPartyMembersDrums);
export const DrumsOfRestorationConsume = makeEnumValueConsumeInput(ActionId.fromSpellId(35478), 'drums', Drums.DrumsOfRestoration, ['Drums'], removeOtherPartyMembersDrums);

function makeBooleanRaidBuffInput(id: ActionId, buffsFieldName: keyof RaidBuffs, exclusivityTags?: Array<ExclusivityTag>): IndividualSimIconPickerConfig<Raid, boolean> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (raid: Raid) => raid.buffsChangeEmitter,
    getValue: (raid: Raid) => raid.getBuffs()[buffsFieldName] as boolean,
    setValue: (eventID: EventID, raid: Raid, newValue: boolean) => {
      const newBuffs = raid.getBuffs();
      (newBuffs[buffsFieldName] as boolean) = newValue;
      raid.setBuffs(eventID, newBuffs);
    },
  }
}

function makeTristateRaidBuffInput(id: ActionId, impId: ActionId, buffsFieldName: keyof RaidBuffs, exclusivityTags?: Array<ExclusivityTag>): IndividualSimIconPickerConfig<Raid, number> {
  return {
    id: id,
    states: 3,
    improvedId: impId,
    exclusivityTags: exclusivityTags,
    changedEvent: (raid: Raid) => raid.buffsChangeEmitter,
    getValue: (raid: Raid) => raid.getBuffs()[buffsFieldName] as number,
    setValue: (eventID: EventID, raid: Raid, newValue: number) => {
      const newBuffs = raid.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      raid.setBuffs(eventID, newBuffs);
    },
  }
}

function makeBooleanPartyBuffInput(id: ActionId, buffsFieldName: keyof PartyBuffs, exclusivityTags?: Array<ExclusivityTag>): IndividualSimIconPickerConfig<Party, boolean> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (party: Party) => party.buffsChangeEmitter,
    getValue: (party: Party) => party.getBuffs()[buffsFieldName] as boolean,
    setValue: (eventID: EventID, party: Party, newValue: boolean) => {
      const newBuffs = party.getBuffs();
      (newBuffs[buffsFieldName] as boolean) = newValue;
      party.setBuffs(eventID, newBuffs);
    },
  }
}

function makeTristatePartyBuffInput(id: ActionId, impId: ActionId, buffsFieldName: keyof PartyBuffs): IndividualSimIconPickerConfig<Party, number> {
  return {
    id: id,
    states: 3,
    improvedId: impId,
    changedEvent: (party: Party) => party.buffsChangeEmitter,
    getValue: (party: Party) => party.getBuffs()[buffsFieldName] as number,
    setValue: (eventID: EventID, party: Party, newValue: number) => {
      const newBuffs = party.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      party.setBuffs(eventID, newBuffs);
    },
  }
}

function makeMultistatePartyBuffInput(id: ActionId, numStates: number, buffsFieldName: keyof PartyBuffs): IndividualSimIconPickerConfig<Party, number> {
  return {
    id: id,
    states: numStates,
    changedEvent: (party: Party) => party.buffsChangeEmitter,
    getValue: (party: Party) => party.getBuffs()[buffsFieldName] as number,
    setValue: (eventID: EventID, party: Party, newValue: number) => {
      const newBuffs = party.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      party.setBuffs(eventID, newBuffs);
    },
  }
}

function makeEnumValuePartyBuffInput(id: ActionId, buffsFieldName: keyof PartyBuffs, enumValue: number, exclusivityTags?: Array<ExclusivityTag>): IndividualSimIconPickerConfig<Party, boolean> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (party: Party) => party.buffsChangeEmitter,
    getValue: (party: Party) => party.getBuffs()[buffsFieldName] == enumValue,
    setValue: (eventID: EventID, party: Party, newValue: boolean) => {
			const newBuffs = party.getBuffs();
			(newBuffs[buffsFieldName] as number) = newValue ? enumValue : 0;
			party.setBuffs(eventID, newBuffs);
    },
  }
}

function makeBooleanIndividualBuffInput(id: ActionId, buffsFieldName: keyof IndividualBuffs, exclusivityTags?: Array<ExclusivityTag>): IndividualSimIconPickerConfig<Player<any>, boolean> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (player: Player<any>) => player.buffsChangeEmitter,
    getValue: (player: Player<any>) => player.getBuffs()[buffsFieldName] as boolean,
    setValue: (eventID: EventID, player: Player<any>, newValue: boolean) => {
      const newBuffs = player.getBuffs();
      (newBuffs[buffsFieldName] as boolean) = newValue;
      player.setBuffs(eventID, newBuffs);
    },
  }
}

function makeTristateIndividualBuffInput(id: ActionId, impId: ActionId, buffsFieldName: keyof IndividualBuffs): IndividualSimIconPickerConfig<Player<any>, number> {
  return {
    id: id,
    states: 3,
    improvedId: impId,
    changedEvent: (player: Player<any>) => player.buffsChangeEmitter,
    getValue: (player: Player<any>) => player.getBuffs()[buffsFieldName] as number,
    setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
      const newBuffs = player.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      player.setBuffs(eventID, newBuffs);
    },
  }
}

function makeMultistateIndividualBuffInput(id: ActionId, numStates: number, buffsFieldName: keyof IndividualBuffs): IndividualSimIconPickerConfig<Player<any>, number> {
  return {
    id: id,
    states: numStates,
    changedEvent: (player: Player<any>) => player.buffsChangeEmitter,
    getValue: (player: Player<any>) => player.getBuffs()[buffsFieldName] as number,
    setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
      const newBuffs = player.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      player.setBuffs(eventID, newBuffs);
    },
  }
}

function makeBooleanDebuffInput(id: ActionId, debuffsFieldName: keyof Debuffs, exclusivityTags?: Array<ExclusivityTag>): IndividualSimIconPickerConfig<Target, boolean> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (target: Target) => target.debuffsChangeEmitter,
    getValue: (target: Target) => target.getDebuffs()[debuffsFieldName] as boolean,
    setValue: (eventID: EventID, target: Target, newValue: boolean) => {
      const newDebuffs = target.getDebuffs();
      (newDebuffs[debuffsFieldName] as boolean) = newValue;
      target.setDebuffs(eventID, newDebuffs);
    },
  }
}

function makeTristateDebuffInput(id: ActionId, impId: ActionId, debuffsFieldName: keyof Debuffs): IndividualSimIconPickerConfig<Target, number> {
  return {
    id: id,
    states: 3,
    improvedId: impId,
    changedEvent: (target: Target) => target.debuffsChangeEmitter,
    getValue: (target: Target) => target.getDebuffs()[debuffsFieldName] as number,
    setValue: (eventID: EventID, target: Target, newValue: number) => {
      const newDebuffs = target.getDebuffs();
      (newDebuffs[debuffsFieldName] as number) = newValue;
      target.setDebuffs(eventID, newDebuffs);
    },
  }
}

function makeBooleanConsumeInput(id: ActionId, consumesFieldName: keyof Consumes, exclusivityTags?: Array<ExclusivityTag>): IndividualSimIconPickerConfig<Player<any>, boolean> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (player: Player<any>) => player.consumesChangeEmitter,
    getValue: (player: Player<any>) => player.getConsumes()[consumesFieldName] as boolean,
    setValue: (eventID: EventID, player: Player<any>, newValue: boolean) => {
      const newBuffs = player.getConsumes();
      (newBuffs[consumesFieldName] as boolean) = newValue;
      player.setConsumes(eventID, newBuffs);
    },
  }
}

function makeEnumValueConsumeInput(id: ActionId, consumesFieldName: keyof Consumes, enumValue: number, exclusivityTags?: Array<ExclusivityTag>, onSet?: (eventID: EventID, player: Player<any>, newValue: boolean) => void, showWhen?: (player: Player<any>) => boolean): IndividualSimIconPickerConfig<Player<any>, boolean> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (player: Player<any>) => player.consumesChangeEmitter,
    getValue: (player: Player<any>) => player.getConsumes()[consumesFieldName] == enumValue,
    setValue: (eventID: EventID, player: Player<any>, newValue: boolean) => {
			const newConsumes = player.getConsumes();
			(newConsumes[consumesFieldName] as number) = newValue ? enumValue : 0;
			TypedEvent.freezeAllAndDo(() => {
				player.setConsumes(eventID, newConsumes);
				if (onSet) {
					onSet(eventID, player, newValue);
				}
			});
    },
		showWhen: showWhen,
  }
}

//////////////////////////////////////////////////////////////////////
// Custom buffs that don't fit into any of the helper functions above.
//////////////////////////////////////////////////////////////////////

export const GraceOfAirTotem = {
	id: ActionId.fromSpellId(25359),
	states: 3,
	improvedId: ActionId.fromSpellId(16295),
	changedEvent: (party: Party) => party.buffsChangeEmitter,
	getValue: (party: Party) => party.getBuffs().graceOfAirTotem,
	setValue: (eventID: EventID, party: Party, newValue: number) => {
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
	changedEvent: (party: Party) => party.buffsChangeEmitter,
	getValue: (party: Party) => party.getBuffs().strengthOfEarthTotem > 2 ? party.getBuffs().strengthOfEarthTotem - 1 : party.getBuffs().strengthOfEarthTotem,
	setValue: (eventID: EventID, party: Party, newValue: number) => {
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
	changedEvent: (party: Party) => party.buffsChangeEmitter,
	getValue: (party: Party) => {
		const buffs = party.getBuffs();
		if (buffs.windfuryTotemRank == 0) {
			return 0;
		}

		if (buffs.windfuryTotemIwt > 0) {
			return 2;
		} else {
			return 1;
		}
	},
	setValue: (eventID: EventID, party: Party, newValue: number) => {
		const newBuffs = party.getBuffs();
		if (newValue == 0) {
			newBuffs.windfuryTotemRank = 0;
			newBuffs.windfuryTotemIwt = 0;
		} else {
			newBuffs.windfuryTotemRank = 5;
			if (newValue == 2) {
				newBuffs.windfuryTotemIwt = 2;
			} else {
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
	changedEvent: (party: Party) => party.buffsChangeEmitter,
	getValue: (party: Party) => {
		const buffs = party.getBuffs();
		if (buffs.battleShout == TristateEffect.TristateEffectImproved) {
			return buffs.battleShout + Number(buffs.bsSolarianSapphire);
		} else {
			return buffs.battleShout;
		}
	},
	setValue: (eventID: EventID, party: Party, newValue: number) => {
		const newBuffs = party.getBuffs();
		newBuffs.battleShout = Math.min(2, newValue);
		newBuffs.bsSolarianSapphire = newValue == 3;
		party.setBuffs(eventID, newBuffs);
	},
};

function makeAdamantiteStoneInput(id: ActionId, isSharp: boolean, isMH: boolean, enumValue: number): IndividualSimIconPickerConfig<Player<any>, boolean> {
	const exclusivityTags: Array<ExclusivityTag> = isMH ? ['MH Weapon Imbue'] : ['OH Weapon Imbue'];
	const consumesFieldName: keyof Consumes = isMH ? 'mainHandImbue' : 'offHandImbue';

  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (player: Player<any>) => TypedEvent.onAny([player.consumesChangeEmitter, player.gearChangeEmitter]),
    getValue: (player: Player<any>) => player.getConsumes()[consumesFieldName] == enumValue,
    setValue: (eventID: EventID, player: Player<any>, newValue: boolean) => {
			const newConsumes = player.getConsumes();
			(newConsumes[consumesFieldName] as number) = newValue ? enumValue : 0;
			player.setConsumes(eventID, newConsumes);
    },
		showWhen: (player: Player<any>) => (isMH ? player.getGear().hasBluntMHWeapon() : player.getGear().hasBluntOHWeapon()) != isSharp,
  }
}
