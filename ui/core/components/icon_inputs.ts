import { Buffs } from '../api/newapi'
import { Consumes } from '../api/newapi'
import { Sim } from '../sim'

import { ExclusivityTag } from '../themes/theme';
import { IconInput } from './icon_picker'

// Keep each section in alphabetical order.
// Buffs
export const ArcaneBrilliance = makeBooleanBuffSpellInput(27127, 'arcaneInt');
export const BlessingOfKings = makeBooleanBuffSpellInput(25898, 'blessingOfKings');
export const BlessingOfWisdom = makeTristateBuffSpellInput(27143, 20245, 'blessingOfWisdom');
export const Bloodlust = makeMultistateBuffSpellInput(2825, 11, 'bloodlust');
export const ChainOfTheTwilightOwl = makeBooleanBuffSpellInput(31035, 'chainOfTheTwilightOwl');
export const EyeOfTheNight = makeBooleanBuffSpellInput(31033, 'eyeOfTheNight');
export const GiftOfTheWild = makeBooleanBuffSpellInput(26991, 'giftOfTheWild');
export const ImprovedDivineSpirit = makeBooleanBuffSpellInput(33182, 'improvedDivineSpirit');
export const JadePendantOfBlasting = makeBooleanBuffSpellInput(25607, 'jadePendantOfBlasting');
export const MoonkinAura = makeTristateBuffSpellItemInput(24907, 32387, 'moonkinAura');

// Debuffs
export const ImprovedSealOfTheCrusader = makeBooleanBuffSpellInput(20337, 'improvedSealOfTheCrusader');
export const JudgementOfWisdom = makeBooleanBuffSpellInput(27164, 'judgementOfWisdom');
export const Misery = makeBooleanBuffSpellInput(33195, 'misery');

// Consumes
export const AdeptsElixir = makeBooleanConsumeItemInput(28103, 'adeptsElixir', ['Battle Elixir']);
export const BlackenedBasilisk = makeBooleanConsumeItemInput(27657, 'blackenedBasilisk', ['Food']);
export const BrilliantWizardOil = makeBooleanConsumeItemInput(20749, 'brilliantWizardOil', ['Weapon Imbue']);
export const DarkRune = makeBooleanConsumeItemInput(12662, 'darkRune', ['Rune']);
export const DestructionPotion = makeBooleanConsumeItemInput(22839, 'destructionPotion', ['Potion']);
export const DrumsOfBattle = makeBooleanConsumeSpellInput(35476, 'drumsOfBattle', ['Drums']);
export const DrumsOfRestoration = makeBooleanConsumeSpellInput(35478, 'drumsOfRestoration', ['Drums']);
export const ElixirOfDraenicWisdom = makeBooleanConsumeItemInput(32067, 'elixirOfDraenicWisdom', ['Guardian Elixir']);
export const ElixirOfMajorFirePower = makeBooleanConsumeItemInput(22833, 'elixirOfMajorFirePower', ['Battle Elixir']);
export const ElixirOfMajorFrostPower = makeBooleanConsumeItemInput(22827, 'elixirOfMajorFrostPower', ['Battle Elixir']);
export const ElixirOfMajorMageblood = makeBooleanConsumeItemInput(22840, 'elixirOfMajorMageblood', ['Guardian Elixir']);
export const ElixirOfMajorShadowPower = makeBooleanConsumeItemInput(22835, 'elixirOfMajorShadowPower', ['Battle Elixir']);
export const FlaskOfBlindingLight = makeBooleanConsumeItemInput(22861, 'flaskOfBlindingLight', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfMightyRestoration = makeBooleanConsumeItemInput(22853, 'flaskOfMightyRestoration', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfPureDeath = makeBooleanConsumeItemInput(22866, 'flaskOfPureDeath', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfSupremePower = makeBooleanConsumeItemInput(13512, 'flaskOfSupremePower', ['Battle Elixir', 'Guardian Elixir']);
export const SkullfishSoup = makeBooleanConsumeItemInput(33825, 'skullfishSoup', ['Food']);
export const SuperManaPotion = makeBooleanConsumeItemInput(22832, 'superManaPotion', ['Potion']);
export const SuperiorWizardOil = makeBooleanConsumeItemInput(22522, 'superiorWizardOil', ['Weapon Imbue']);

function makeBooleanBuffSpellInput(spellId: number, buffsFieldName: keyof Buffs): IconInput {
  return {
    spellId: spellId,
    states: 2,
    changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
    getValue: (sim: Sim) => sim.buffs[buffsFieldName],
    setBooleanValue: (sim: Sim, newValue: boolean) => {
      const newBuffs = sim.buffs;
      (newBuffs[buffsFieldName] as boolean) = newValue;
      sim.buffs = newBuffs;
    },
  }
}

function makeTristateBuffSpellInput(spellId: number, impSpellId: number, buffsFieldName: keyof Buffs): IconInput {
  return {
    spellId: spellId,
    states: 3,
    improvedSpellId: impSpellId,
    changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
    getValue: (sim: Sim) => sim.buffs[buffsFieldName],
    setNumberValue: (sim: Sim, newValue: number) => {
      const newBuffs = sim.buffs;
      (newBuffs[buffsFieldName] as number) = newValue;
      sim.buffs = newBuffs;
    },
  }
}

function makeTristateBuffSpellItemInput(spellId: number, impItemId: number, buffsFieldName: keyof Buffs): IconInput {
  return {
    spellId: spellId,
    states: 3,
    improvedItemId: impItemId,
    changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
    getValue: (sim: Sim) => sim.buffs[buffsFieldName],
    setNumberValue: (sim: Sim, newValue: number) => {
      const newBuffs = sim.buffs;
      (newBuffs[buffsFieldName] as number) = newValue;
      sim.buffs = newBuffs;
    },
  }
}

function makeMultistateBuffSpellInput(spellId: number, numStates: number, buffsFieldName: keyof Buffs): IconInput {
  return {
    spellId: spellId,
    states: numStates,
    changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
    getValue: (sim: Sim) => sim.buffs[buffsFieldName],
    setNumberValue: (sim: Sim, newValue: number) => {
      const newBuffs = sim.buffs;
      (newBuffs[buffsFieldName] as number) = newValue;
      sim.buffs = newBuffs;
    },
  }
}

function makeBooleanConsumeItemInput(itemId: number, consumesFieldName: keyof Consumes, exclusivityTags?: Array<ExclusivityTag>): IconInput {
  return {
    itemId: itemId,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (sim: Sim) => sim.consumesChangeEmitter,
    getValue: (sim: Sim) => sim.consumes[consumesFieldName],
    setBooleanValue: (sim: Sim, newValue: boolean) => {
      const newBuffs = sim.consumes;
      (newBuffs[consumesFieldName] as boolean) = newValue;
      sim.consumes = newBuffs;
    },
  }
}

function makeBooleanConsumeSpellInput(spellId: number, consumesFieldName: keyof Consumes, exclusivityTags?: Array<ExclusivityTag>): IconInput {
  return {
    spellId: spellId,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (sim: Sim) => sim.consumesChangeEmitter,
    getValue: (sim: Sim) => sim.consumes[consumesFieldName],
    setBooleanValue: (sim: Sim, newValue: boolean) => {
      const newBuffs = sim.consumes;
      (newBuffs[consumesFieldName] as boolean) = newValue;
      sim.consumes = newBuffs;
    },
  }
}
