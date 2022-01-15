import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Conjured } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
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
export const BattleShout = makeTristatePartyBuffInput(ActionId.fromSpellId(2048), ActionId.fromSpellId(12861), 'battleShout');

export const DrumsOfBattleBuff = makeEnumValuePartyBuffInput(ActionId.fromSpellId(35476), 'drums', Drums.DrumsOfBattle, ['Drums']);
export const DrumsOfRestorationBuff = makeEnumValuePartyBuffInput(ActionId.fromSpellId(35478), 'drums', Drums.DrumsOfRestoration, ['Drums']);

// Individual Buffs
export const BlessingOfKings = makeBooleanIndividualBuffInput(ActionId.fromSpellId(25898), 'blessingOfKings');
export const BlessingOfWisdom = makeTristateIndividualBuffInput(ActionId.fromSpellId(27143), ActionId.fromSpellId(20245), 'blessingOfWisdom');
export const BlessingOfMight = makeTristateIndividualBuffInput(ActionId.fromSpellId(27140), ActionId.fromSpellId(20048), 'blessingOfMight');
export const Innervate = makeMultistateIndividualBuffInput(ActionId.fromSpellId(29166), 11, 'innervates');
export const PowerInfusion = makeMultistateIndividualBuffInput(ActionId.fromSpellId(10060), 11, 'powerInfusions');

// Debuffs
export const BloodFrenzy = makeBooleanDebuffInput(ActionId.fromSpellId(29859), 'bloodFrenzy');
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
export const AdeptsElixir = makeBooleanConsumeInput(ActionId.fromItemId(28103), 'adeptsElixir', ['Battle Elixir']);
export const BlackenedBasilisk = makeBooleanConsumeInput(ActionId.fromItemId(27657), 'blackenedBasilisk', ['Food']);
export const BrilliantWizardOil = makeBooleanConsumeInput(ActionId.fromItemId(20749), 'brilliantWizardOil', ['Weapon Imbue']);
export const ElixirOfDemonslaying = makeBooleanConsumeInput(ActionId.fromItemId(9224), 'elixirOfDemonslaying', ['Battle Elixir']);
export const ElixirOfDraenicWisdom = makeBooleanConsumeInput(ActionId.fromItemId(32067), 'elixirOfDraenicWisdom', ['Guardian Elixir']);
export const ElixirOfMajorFirePower = makeBooleanConsumeInput(ActionId.fromItemId(22833), 'elixirOfMajorFirePower', ['Battle Elixir']);
export const ElixirOfMajorFrostPower = makeBooleanConsumeInput(ActionId.fromItemId(22827), 'elixirOfMajorFrostPower', ['Battle Elixir']);
export const ElixirOfMajorMageblood = makeBooleanConsumeInput(ActionId.fromItemId(22840), 'elixirOfMajorMageblood', ['Guardian Elixir']);
export const ElixirOfMajorShadowPower = makeBooleanConsumeInput(ActionId.fromItemId(22835), 'elixirOfMajorShadowPower', ['Battle Elixir']);
export const ElixirOfMajorAgility = makeBooleanConsumeInput(ActionId.fromItemId(22831), 'elixirOfMajorAgility', ['Battle Elixir']);
export const FlaskOfBlindingLight = makeBooleanConsumeInput(ActionId.fromItemId(22861), 'flaskOfBlindingLight', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfMightyRestoration = makeBooleanConsumeInput(ActionId.fromItemId(22853), 'flaskOfMightyRestoration', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfPureDeath = makeBooleanConsumeInput(ActionId.fromItemId(22866), 'flaskOfPureDeath', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfSupremePower = makeBooleanConsumeInput(ActionId.fromItemId(13512), 'flaskOfSupremePower', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfRelentlessAssault = makeBooleanConsumeInput(ActionId.fromItemId(22854), 'flaskOfRelentlessAssault', ['Battle Elixir', 'Guardian Elixir']);
export const KreegsStoutBeatdown = makeBooleanConsumeInput(ActionId.fromItemId(18284), 'kreegsStoutBeatdown', ['Alchohol']);
export const SkullfishSoup = makeBooleanConsumeInput(ActionId.fromItemId(33825), 'skullfishSoup', ['Food']);
export const SuperiorWizardOil = makeBooleanConsumeInput(ActionId.fromItemId(22522), 'superiorWizardOil', ['Weapon Imbue']);
export const RoastedClefthoof = makeBooleanConsumeInput(ActionId.fromItemId(27658), 'roastedClefthoof', ['Food']);
export const ScrollOfStrengthV = makeBooleanConsumeInput(ActionId.fromItemId(27503), 'scrollOfStrengthV');
export const ScrollOfAgilityV = makeBooleanConsumeInput(ActionId.fromItemId(27498), 'scrollOfAgilityV');
export const ScrollOfSpiritV = makeBooleanConsumeInput(ActionId.fromItemId(27501), 'scrollOfSpiritV', ['Spirit']);

export const DefaultDestructionPotion = makeEnumValueConsumeInput(ActionId.fromItemId(22839), 'defaultPotion', Potions.DestructionPotion, ['Potion']);
export const DefaultHastePotion = makeEnumValueConsumeInput(ActionId.fromItemId(22838), 'defaultPotion', Potions.HastePotion, ['Potion']);
export const DefaultSuperManaPotion = makeEnumValueConsumeInput(ActionId.fromItemId(22832), 'defaultPotion', Potions.SuperManaPotion, ['Potion']);

export const DefaultDarkRune = makeEnumValueConsumeInput(ActionId.fromItemId(12662), 'defaultConjured', Conjured.ConjuredDarkRune, ['Conjured']);
export const DefaultFlameCap = makeEnumValueConsumeInput(ActionId.fromItemId(22788), 'defaultConjured', Conjured.ConjuredFlameCap, ['Conjured']);

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

function makeEnumValueConsumeInput(id: ActionId, consumesFieldName: keyof Consumes, enumValue: number, exclusivityTags?: Array<ExclusivityTag>, onSet?: (eventID: EventID, player: Player<any>, newValue: boolean) => void): IndividualSimIconPickerConfig<Player<any>, boolean> {
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
  }
}
