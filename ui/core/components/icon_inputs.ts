import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { ItemOrSpellId } from '/tbc/core/resources.js';
import { Party } from '/tbc/core/party.js';
import { Player } from '/tbc/core/player.js';
import { Raid } from '/tbc/core/raid.js';
import { Sim } from '/tbc/core/sim.js';
import { Target } from '/tbc/core/target.js';

import { ExclusivityTag } from '/tbc/core/individual_sim_ui.js';
import { IconInput } from './icon_picker.js';

// Keep each section in alphabetical order.

// Raid Buffs
export const ArcaneBrilliance = makeBooleanRaidBuffInput({spellId:27127}, 'arcaneBrilliance');
export const DivineSpirit = makeTristateRaidBuffInput({spellId:25312}, {spellId:33182}, 'divineSpirit');
export const GiftOfTheWild = makeTristateRaidBuffInput({spellId:26991}, {spellId:17055}, 'giftOfTheWild');

// Party Buffs
export const AtieshMage = makeMultistatePartyBuffInput({spellId:28142}, 5, 'atieshMage');
export const AtieshWarlock = makeMultistatePartyBuffInput({spellId:28143}, 5, 'atieshWarlock');
export const Bloodlust = makeMultistatePartyBuffInput({spellId:2825}, 11, 'bloodlust');
export const BraidedEterniumChain = makeBooleanPartyBuffInput({spellId:31025}, 'braidedEterniumChain');
export const ChainOfTheTwilightOwl = makeBooleanPartyBuffInput({spellId:31035}, 'chainOfTheTwilightOwl');
export const DraeneiRacialCaster = makeBooleanPartyBuffInput({spellId:28878}, 'draeneiRacialCaster');
export const DraeneiRacialMelee = makeBooleanPartyBuffInput({spellId:6562}, 'draeneiRacialMelee');
export const EyeOfTheNight = makeBooleanPartyBuffInput({spellId:31033}, 'eyeOfTheNight');
export const JadePendantOfBlasting = makeBooleanPartyBuffInput({spellId:25607}, 'jadePendantOfBlasting');
export const ManaSpringTotem = makeTristatePartyBuffInput({spellId:25570}, {spellId:16208}, 'manaSpringTotem');
export const MoonkinAura = makeTristatePartyBuffInput({spellId:24907}, {itemId:32387}, 'moonkinAura');
export const TotemOfWrath = makeMultistatePartyBuffInput({spellId:30706}, 5, 'totemOfWrath');
export const WrathOfAirTotem = makeTristatePartyBuffInput({spellId:3738}, {spellId:37212}, 'wrathOfAirTotem');

export const DrumsOfBattleBuff = makeEnumValuePartyBuffInput({spellId:35476}, 'drums', Drums.DrumsOfBattle, ['Drums']);
export const DrumsOfRestorationBuff = makeEnumValuePartyBuffInput({spellId:35478}, 'drums', Drums.DrumsOfRestoration, ['Drums']);

// Individual Buffs
export const BlessingOfKings = makeBooleanIndividualBuffInput({spellId:25898}, 'blessingOfKings');
export const BlessingOfWisdom = makeTristateIndividualBuffInput({spellId:27143}, {spellId:20245}, 'blessingOfWisdom');
export const ManaTideTotem = makeBooleanIndividualBuffInput({spellId:16190}, 'manaTideTotem');

// Debuffs
export const ImprovedSealOfTheCrusader = makeBooleanDebuffInput({spellId:20337}, 'improvedSealOfTheCrusader');
export const JudgementOfWisdom = makeBooleanDebuffInput({spellId:27164}, 'judgementOfWisdom');
export const Misery = makeBooleanDebuffInput({spellId:33195}, 'misery');
export const CurseOfElements = makeTristateDebuffInput({spellId:27228}, {spellId:32484}, 'curseOfElements');

// Consumes
export const AdeptsElixir = makeBooleanConsumeInput({itemId:28103}, 'adeptsElixir', ['Battle Elixir']);
export const BlackenedBasilisk = makeBooleanConsumeInput({itemId:27657}, 'blackenedBasilisk', ['Food']);
export const BrilliantWizardOil = makeBooleanConsumeInput({itemId:20749}, 'brilliantWizardOil', ['Weapon Imbue']);
export const DarkRune = makeBooleanConsumeInput({itemId:12662}, 'darkRune', ['Rune']);
export const ElixirOfDraenicWisdom = makeBooleanConsumeInput({itemId:32067}, 'elixirOfDraenicWisdom', ['Guardian Elixir']);
export const ElixirOfMajorFirePower = makeBooleanConsumeInput({itemId:22833}, 'elixirOfMajorFirePower', ['Battle Elixir']);
export const ElixirOfMajorFrostPower = makeBooleanConsumeInput({itemId:22827}, 'elixirOfMajorFrostPower', ['Battle Elixir']);
export const ElixirOfMajorMageblood = makeBooleanConsumeInput({itemId:22840}, 'elixirOfMajorMageblood', ['Guardian Elixir']);
export const ElixirOfMajorShadowPower = makeBooleanConsumeInput({itemId:22835}, 'elixirOfMajorShadowPower', ['Battle Elixir']);
export const FlaskOfBlindingLight = makeBooleanConsumeInput({itemId:22861}, 'flaskOfBlindingLight', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfMightyRestoration = makeBooleanConsumeInput({itemId:22853}, 'flaskOfMightyRestoration', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfPureDeath = makeBooleanConsumeInput({itemId:22866}, 'flaskOfPureDeath', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfSupremePower = makeBooleanConsumeInput({itemId:13512}, 'flaskOfSupremePower', ['Battle Elixir', 'Guardian Elixir']);
export const KreegsStoutBeatdown = makeBooleanConsumeInput({itemId:18284}, 'kreegsStoutBeatdown', ['Alchohol']);
export const SkullfishSoup = makeBooleanConsumeInput({itemId:33825}, 'skullfishSoup', ['Food']);
export const SuperiorWizardOil = makeBooleanConsumeInput({itemId:22522}, 'superiorWizardOil', ['Weapon Imbue']);

export const DefaultDestructionPotion = makeEnumValueConsumeInput({itemId:22839}, 'defaultPotion', Potions.DestructionPotion, ['Potion']);
export const DefaultSuperManaPotion = makeEnumValueConsumeInput({itemId:22832}, 'defaultPotion', Potions.SuperManaPotion, ['Potion']);

export const DrumsOfBattleConsume = makeEnumValueConsumeInput({spellId:35476}, 'drums', Drums.DrumsOfBattle, ['Drums']);
export const DrumsOfRestorationConsume = makeEnumValueConsumeInput({spellId:35478}, 'drums', Drums.DrumsOfRestoration, ['Drums']);

function makeBooleanRaidBuffInput(id: ItemOrSpellId, buffsFieldName: keyof RaidBuffs, exclusivityTags?: Array<ExclusivityTag>): IconInput<Raid> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (raid: Raid) => raid.buffsChangeEmitter,
    getValue: (raid: Raid) => raid.getBuffs()[buffsFieldName],
    setBooleanValue: (raid: Raid, newValue: boolean) => {
      const newBuffs = raid.getBuffs();
      (newBuffs[buffsFieldName] as boolean) = newValue;
      raid.setBuffs(newBuffs);
    },
  }
}

function makeTristateRaidBuffInput(id: ItemOrSpellId, impId: ItemOrSpellId, buffsFieldName: keyof RaidBuffs): IconInput<Raid> {
  return {
    id: id,
    states: 3,
    improvedId: impId,
    changedEvent: (raid: Raid) => raid.buffsChangeEmitter,
    getValue: (raid: Raid) => raid.getBuffs()[buffsFieldName],
    setNumberValue: (raid: Raid, newValue: number) => {
      const newBuffs = raid.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      raid.setBuffs(newBuffs);
    },
  }
}

function makeBooleanPartyBuffInput(id: ItemOrSpellId, buffsFieldName: keyof PartyBuffs, exclusivityTags?: Array<ExclusivityTag>): IconInput<Party> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (party: Party) => party.buffsChangeEmitter,
    getValue: (party: Party) => party.getBuffs()[buffsFieldName],
    setBooleanValue: (party: Party, newValue: boolean) => {
      const newBuffs = party.getBuffs();
      (newBuffs[buffsFieldName] as boolean) = newValue;
      party.setBuffs(newBuffs);
    },
  }
}

function makeTristatePartyBuffInput(id: ItemOrSpellId, impId: ItemOrSpellId, buffsFieldName: keyof PartyBuffs): IconInput<Party> {
  return {
    id: id,
    states: 3,
    improvedId: impId,
    changedEvent: (party: Party) => party.buffsChangeEmitter,
    getValue: (party: Party) => party.getBuffs()[buffsFieldName],
    setNumberValue: (party: Party, newValue: number) => {
      const newBuffs = party.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      party.setBuffs(newBuffs);
    },
  }
}

function makeMultistatePartyBuffInput(id: ItemOrSpellId, numStates: number, buffsFieldName: keyof PartyBuffs): IconInput<Party> {
  return {
    id: id,
    states: numStates,
    changedEvent: (party: Party) => party.buffsChangeEmitter,
    getValue: (party: Party) => party.getBuffs()[buffsFieldName],
    setNumberValue: (party: Party, newValue: number) => {
      const newBuffs = party.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      party.setBuffs(newBuffs);
    },
  }
}

function makeEnumValuePartyBuffInput(id: ItemOrSpellId, buffsFieldName: keyof PartyBuffs, enumValue: number, exclusivityTags?: Array<ExclusivityTag>): IconInput<Party> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (party: Party) => party.buffsChangeEmitter,
    getValue: (party: Party) => party.getBuffs()[buffsFieldName] == enumValue,
    setBooleanValue: (party: Party, newValue: boolean) => {
			const newBuffs = party.getBuffs();
			(newBuffs[buffsFieldName] as number) = newValue ? enumValue : 0;
			party.setBuffs(newBuffs);
    },
  }
}

function makeBooleanIndividualBuffInput(id: ItemOrSpellId, buffsFieldName: keyof IndividualBuffs, exclusivityTags?: Array<ExclusivityTag>): IconInput<Player<any>> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (player: Player<any>) => player.buffsChangeEmitter,
    getValue: (player: Player<any>) => player.getBuffs()[buffsFieldName],
    setBooleanValue: (player: Player<any>, newValue: boolean) => {
      const newBuffs = player.getBuffs();
      (newBuffs[buffsFieldName] as boolean) = newValue;
      player.setBuffs(newBuffs);
    },
  }
}

function makeTristateIndividualBuffInput(id: ItemOrSpellId, impId: ItemOrSpellId, buffsFieldName: keyof IndividualBuffs): IconInput<Player<any>> {
  return {
    id: id,
    states: 3,
    improvedId: impId,
    changedEvent: (player: Player<any>) => player.buffsChangeEmitter,
    getValue: (player: Player<any>) => player.getBuffs()[buffsFieldName],
    setNumberValue: (player: Player<any>, newValue: number) => {
      const newBuffs = player.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      player.setBuffs(newBuffs);
    },
  }
}

function makeBooleanDebuffInput(id: ItemOrSpellId, debuffsFieldName: keyof Debuffs, exclusivityTags?: Array<ExclusivityTag>): IconInput<Target> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (target: Target) => target.debuffsChangeEmitter,
    getValue: (target: Target) => target.getDebuffs()[debuffsFieldName],
    setBooleanValue: (target: Target, newValue: boolean) => {
      const newDebuffs = target.getDebuffs();
      (newDebuffs[debuffsFieldName] as boolean) = newValue;
      target.setDebuffs(newDebuffs);
    },
  }
}

function makeTristateDebuffInput(id: ItemOrSpellId, impId: ItemOrSpellId, debuffsFieldName: keyof Debuffs): IconInput<Target> {
  return {
    id: id,
    states: 3,
    improvedId: impId,
    changedEvent: (target: Target) => target.debuffsChangeEmitter,
    getValue: (target: Target) => target.getDebuffs()[debuffsFieldName],
    setNumberValue: (target: Target, newValue: number) => {
      const newDebuffs = target.getDebuffs();
      (newDebuffs[debuffsFieldName] as number) = newValue;
      target.setDebuffs(newDebuffs);
    },
  }
}

function makeBooleanConsumeInput(id: ItemOrSpellId, consumesFieldName: keyof Consumes, exclusivityTags?: Array<ExclusivityTag>): IconInput<Player<any>> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (player: Player<any>) => player.consumesChangeEmitter,
    getValue: (player: Player<any>) => player.getConsumes()[consumesFieldName],
    setBooleanValue: (player: Player<any>, newValue: boolean) => {
      const newBuffs = player.getConsumes();
      (newBuffs[consumesFieldName] as boolean) = newValue;
      player.setConsumes(newBuffs);
    },
  }
}

function makeEnumValueConsumeInput(id: ItemOrSpellId, consumesFieldName: keyof Consumes, enumValue: number, exclusivityTags?: Array<ExclusivityTag>): IconInput<Player<any>> {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (player: Player<any>) => player.consumesChangeEmitter,
    getValue: (player: Player<any>) => player.getConsumes()[consumesFieldName] == enumValue,
    setBooleanValue: (player: Player<any>, newValue: boolean) => {
			const newConsumes = player.getConsumes();
			(newConsumes[consumesFieldName] as number) = newValue ? enumValue : 0;
			player.setConsumes(newConsumes);
    },
  }
}
