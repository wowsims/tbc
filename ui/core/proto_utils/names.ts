import { Class } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';

export const raceNames: Record<Race, string> = {
  [Race.RaceUnknown]: 'None',
  [Race.RaceBloodElf]: 'Blood Elf',
  [Race.RaceDraenei]: 'Draenei',
  [Race.RaceDwarf]: 'Dwarf',
  [Race.RaceGnome]: 'Gnome',
  [Race.RaceHuman]: 'Human',
  [Race.RaceNightElf]: 'Night Elf',
  [Race.RaceOrc]: 'Orc',
  [Race.RaceTauren]: 'Tauren',
  [Race.RaceTroll10]: 'Troll (+10% Haste)',
  [Race.RaceTroll30]: 'Troll (+30% Haste)',
  [Race.RaceUndead]: 'Undead',
};

export function nameToRace(name: string): Race {
	const lower = name.toLowerCase();
	if (lower.includes('troll')) {
		return Race.RaceTroll10;
	}

	for (const key in raceNames) {
		const race = parseInt(key) as Race;
		if (raceNames[race].toLowerCase() == lower) {
			return race;
		}
	}

	return Race.RaceUnknown;
}

export const classNames: Record<Class, string> = {
	[Class.ClassUnknown]: 'None',
	[Class.ClassDruid]: 'Druid',
	[Class.ClassHunter]: 'Hunter',
	[Class.ClassMage]: 'Mage',
	[Class.ClassPaladin]: 'Paladin',
	[Class.ClassPriest]: 'Priest',
	[Class.ClassRogue]: 'Rogue',
	[Class.ClassShaman]: 'Shaman',
	[Class.ClassWarlock]: 'Warlock',
	[Class.ClassWarrior]: 'Warrior',
}

export function nameToClass(name: string): Class {
	const lower = name.toLowerCase();
	for (const key in classNames) {
		const charClass = parseInt(key) as Class;
		if (classNames[charClass].toLowerCase() == lower) {
			return charClass;
		}
	}

	return Class.ClassUnknown;
}

export const statNames: Record<Stat, string> = {
  [Stat.StatStrength]: 'Strength',
  [Stat.StatAgility]: 'Agility',
  [Stat.StatStamina]: 'Stamina',
  [Stat.StatIntellect]: 'Intellect',
  [Stat.StatSpirit]: 'Spirit',
  [Stat.StatSpellPower]: 'Spell Dmg',
  [Stat.StatHealingPower]: 'Healing Power',
  [Stat.StatArcaneSpellPower]: 'Arcane Dmg',
  [Stat.StatFireSpellPower]: 'Fire Dmg',
  [Stat.StatFrostSpellPower]: 'Frost Dmg',
  [Stat.StatHolySpellPower]: 'Holy Dmg',
  [Stat.StatNatureSpellPower]: 'Nature Dmg',
  [Stat.StatShadowSpellPower]: 'Shadow Dmg',
  [Stat.StatMP5]: 'MP5',
  [Stat.StatSpellHit]: 'Spell Hit Rating',
  [Stat.StatSpellCrit]: 'Spell Crit Rating',
  [Stat.StatSpellHaste]: 'Spell Haste Rating',
  [Stat.StatSpellPenetration]: 'Spell Pen',
  [Stat.StatAttackPower]: 'Attack Power',
  [Stat.StatMeleeHit]: 'Hit Rating',
  [Stat.StatMeleeCrit]: 'Crit Rating',
  [Stat.StatMeleeHaste]: 'Haste Rating',
  [Stat.StatArmorPenetration]: 'Armor Pen',
  [Stat.StatExpertise]: 'Expertise Rating',
  [Stat.StatMana]: 'Mana',
  [Stat.StatEnergy]: 'Energy',
  [Stat.StatRage]: 'Rage',
  [Stat.StatArmor]: 'Armor',
  [Stat.StatRangedAttackPower]: 'Ranged AP',
};

export const slotNames: Record<ItemSlot, string> = {
  [ItemSlot.ItemSlotHead]: 'Head',
  [ItemSlot.ItemSlotNeck]: 'Neck',
  [ItemSlot.ItemSlotShoulder]: 'Shoulders',
  [ItemSlot.ItemSlotBack]: 'Back',
  [ItemSlot.ItemSlotChest]: 'Chest',
  [ItemSlot.ItemSlotWrist]: 'Wrist',
  [ItemSlot.ItemSlotHands]: 'Hands',
  [ItemSlot.ItemSlotWaist]: 'Waist',
  [ItemSlot.ItemSlotLegs]: 'Legs',
  [ItemSlot.ItemSlotFeet]: 'Feet',
  [ItemSlot.ItemSlotFinger1]: 'Finger 1',
  [ItemSlot.ItemSlotFinger2]: 'Finger 2',
  [ItemSlot.ItemSlotTrinket1]: 'Trinket 1',
  [ItemSlot.ItemSlotTrinket2]: 'Trinket 2',
  [ItemSlot.ItemSlotMainHand]: 'Main Hand',
  [ItemSlot.ItemSlotOffHand]: 'Off Hand',
  [ItemSlot.ItemSlotRanged]: 'Ranged',
};
