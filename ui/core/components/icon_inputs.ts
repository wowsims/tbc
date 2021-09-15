import { Buffs } from '../api/newapi'
import { Consumes } from '../api/newapi'
import { ItemOrSpellId } from '../resources'
import { Sim } from '../sim'

import { ExclusivityTag } from '../themes/theme';
import { IconInput } from './icon_picker'

// Keep each section in alphabetical order.
// Buffs
export const ArcaneBrilliance = makeBooleanBuffInput({spellId:27127}, 'arcaneBrilliance');
export const BlessingOfKings = makeBooleanBuffInput({spellId:25898}, 'blessingOfKings');
export const BlessingOfWisdom = makeTristateBuffInput({spellId:27143}, {spellId:20245}, 'blessingOfWisdom');
export const Bloodlust = makeMultistateBuffInput({spellId:2825}, 11, 'bloodlust');
export const ChainOfTheTwilightOwl = makeBooleanBuffInput({spellId:31035}, 'chainOfTheTwilightOwl');
export const DivineSpirit = makeTristateBuffInput({spellId:25312}, {spellId:33182}, 'divineSpirit');
export const EyeOfTheNight = makeBooleanBuffInput({spellId:31033}, 'eyeOfTheNight');
export const GiftOfTheWild = makeTristateBuffInput({spellId:26991}, {spellId:17055}, 'giftOfTheWild');
export const JadePendantOfBlasting = makeBooleanBuffInput({spellId:25607}, 'jadePendantOfBlasting');
export const MoonkinAura = makeTristateBuffInput({spellId:24907}, {itemId:32387}, 'moonkinAura');

// Debuffs
export const ImprovedSealOfTheCrusader = makeBooleanBuffInput({spellId:20337}, 'improvedSealOfTheCrusader');
export const JudgementOfWisdom = makeBooleanBuffInput({spellId:27164}, 'judgementOfWisdom');
export const Misery = makeBooleanBuffInput({spellId:33195}, 'misery');

// Consumes
export const AdeptsElixir = makeBooleanConsumeInput({itemId:28103}, 'adeptsElixir', ['Battle Elixir']);
export const BlackenedBasilisk = makeBooleanConsumeInput({itemId:27657}, 'blackenedBasilisk', ['Food']);
export const BrilliantWizardOil = makeBooleanConsumeInput({itemId:20749}, 'brilliantWizardOil', ['Weapon Imbue']);
export const DarkRune = makeBooleanConsumeInput({itemId:12662}, 'darkRune', ['Rune']);
export const DestructionPotion = makeBooleanConsumeInput({itemId:22839}, 'destructionPotion', ['Potion']);
export const DrumsOfBattle = makeBooleanConsumeInput({spellId:35476}, 'drumsOfBattle', ['Drums']);
export const DrumsOfRestoration = makeBooleanConsumeInput({spellId:35478}, 'drumsOfRestoration', ['Drums']);
export const ElixirOfDraenicWisdom = makeBooleanConsumeInput({itemId:32067}, 'elixirOfDraenicWisdom', ['Guardian Elixir']);
export const ElixirOfMajorFirePower = makeBooleanConsumeInput({itemId:22833}, 'elixirOfMajorFirePower', ['Battle Elixir']);
export const ElixirOfMajorFrostPower = makeBooleanConsumeInput({itemId:22827}, 'elixirOfMajorFrostPower', ['Battle Elixir']);
export const ElixirOfMajorMageblood = makeBooleanConsumeInput({itemId:22840}, 'elixirOfMajorMageblood', ['Guardian Elixir']);
export const ElixirOfMajorShadowPower = makeBooleanConsumeInput({itemId:22835}, 'elixirOfMajorShadowPower', ['Battle Elixir']);
export const FlaskOfBlindingLight = makeBooleanConsumeInput({itemId:22861}, 'flaskOfBlindingLight', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfMightyRestoration = makeBooleanConsumeInput({itemId:22853}, 'flaskOfMightyRestoration', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfPureDeath = makeBooleanConsumeInput({itemId:22866}, 'flaskOfPureDeath', ['Battle Elixir', 'Guardian Elixir']);
export const FlaskOfSupremePower = makeBooleanConsumeInput({itemId:13512}, 'flaskOfSupremePower', ['Battle Elixir', 'Guardian Elixir']);
export const SkullfishSoup = makeBooleanConsumeInput({itemId:33825}, 'skullfishSoup', ['Food']);
export const SuperManaPotion = makeBooleanConsumeInput({itemId:22832}, 'superManaPotion', ['Potion']);
export const SuperiorWizardOil = makeBooleanConsumeInput({itemId:22522}, 'superiorWizardOil', ['Weapon Imbue']);

function makeBooleanBuffInput(id: ItemOrSpellId, buffsFieldName: keyof Buffs): IconInput {
  return {
    id: id,
    states: 2,
    changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
    getValue: (sim: Sim) => sim.getBuffs()[buffsFieldName],
    setBooleanValue: (sim: Sim, newValue: boolean) => {
      const newBuffs = sim.getBuffs();
      (newBuffs[buffsFieldName] as boolean) = newValue;
      sim.setBuffs(newBuffs);
    },
  }
}

function makeTristateBuffInput(id: ItemOrSpellId, impId: ItemOrSpellId, buffsFieldName: keyof Buffs): IconInput {
  return {
    id: id,
    states: 3,
    improvedId: impId,
    changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
    getValue: (sim: Sim) => sim.getBuffs()[buffsFieldName],
    setNumberValue: (sim: Sim, newValue: number) => {
      const newBuffs = sim.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      sim.setBuffs(newBuffs);
    },
  }
}

function makeMultistateBuffInput(id: ItemOrSpellId, numStates: number, buffsFieldName: keyof Buffs): IconInput {
  return {
    id: id,
    states: numStates,
    changedEvent: (sim: Sim) => sim.buffsChangeEmitter,
    getValue: (sim: Sim) => sim.getBuffs()[buffsFieldName],
    setNumberValue: (sim: Sim, newValue: number) => {
      const newBuffs = sim.getBuffs();
      (newBuffs[buffsFieldName] as number) = newValue;
      sim.setBuffs(newBuffs);
    },
  }
}

function makeBooleanConsumeInput(id: ItemOrSpellId, consumesFieldName: keyof Consumes, exclusivityTags?: Array<ExclusivityTag>): IconInput {
  return {
    id: id,
    states: 2,
    exclusivityTags: exclusivityTags,
    changedEvent: (sim: Sim) => sim.consumesChangeEmitter,
    getValue: (sim: Sim) => sim.getConsumes()[consumesFieldName],
    setBooleanValue: (sim: Sim, newValue: boolean) => {
      const newBuffs = sim.getConsumes();
      (newBuffs[consumesFieldName] as boolean) = newValue;
      sim.setConsumes(newBuffs);
    },
  }
}
