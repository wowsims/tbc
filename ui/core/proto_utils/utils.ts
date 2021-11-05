import { intersection } from '/tbc/core/utils.js';

import { Player } from '/tbc/core/proto/api.js';
import { PlayerOptions } from '/tbc/core/proto/api.js';
import { Class } from '/tbc/core/proto/common.js';
import { Enchant } from '/tbc/core/proto/common.js';
import { Gem } from '/tbc/core/proto/common.js';
import { GemColor } from '/tbc/core/proto/common.js';
import { HandType } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemType } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { RangedWeaponType } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { WeaponType } from '/tbc/core/proto/common.js';

import { BalanceDruid, BalanceDruid_Rotation as BalanceDruidRotation, DruidTalents, BalanceDruid_Options as BalanceDruidOptions} from '/tbc/core/proto/druid.js';
import { ElementalShaman, ElementalShaman_Rotation as ElementalShamanRotation, ShamanTalents, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';
import { Hunter, Hunter_Rotation as HunterRotation, HunterTalents, Hunter_Options as HunterOptions } from '/tbc/core/proto/hunter.js';
import { Mage, Mage_Rotation as MageRotation, MageTalents, Mage_Options as MageOptions } from '/tbc/core/proto/mage.js';
import { Rogue, Rogue_Rotation as RogueRotation, RogueTalents, Rogue_Options as RogueOptions } from '/tbc/core/proto/rogue.js';
import { RetributionPaladin, RetributionPaladin_Rotation as RetributionPaladinRotation, PaladinTalents, RetributionPaladin_Options as RetributionPaladinOptions } from '/tbc/core/proto/paladin.js';
import { ShadowPriest, ShadowPriest_Rotation as ShadowPriestRotation, PriestTalents, ShadowPriest_Options as ShadowPriestOptions } from '/tbc/core/proto/priest.js';
import { Warlock, Warlock_Rotation as WarlockRotation, WarlockTalents, Warlock_Options as WarlockOptions } from '/tbc/core/proto/warlock.js';
import { Warrior, Warrior_Rotation as WarriorRotation, WarriorTalents, Warrior_Options as WarriorOptions } from '/tbc/core/proto/warrior.js';

export type DruidSpecs = Spec.SpecBalanceDruid;
export type HunterSpecs = Spec.SpecHunter;
export type MageSpecs = Spec.SpecMage;
export type RogueSpecs = Spec.SpecRogue;
export type PaladinSpecs = Spec.SpecRetributionPaladin;
export type PriestSpecs = Spec.SpecShadowPriest;
export type ShamanSpecs = Spec.SpecElementalShaman;
export type WarlockSpecs = Spec.SpecWarlock;
export type WarriorSpecs = Spec.SpecWarrior;

export type RotationUnion =
		BalanceDruidRotation |
		ElementalShamanRotation |
		HunterRotation |
		MageRotation |
		RogueRotation |
		RetributionPaladinRotation |
		ShadowPriestRotation |
		WarlockRotation |
		WarriorRotation;
export type SpecRotation<T extends Spec> =
		T extends Spec.SpecBalanceDruid ? BalanceDruidRotation :
		T extends Spec.SpecElementalShaman ? ElementalShamanRotation :
		T extends Spec.SpecHunter ? HunterRotation :
		T extends Spec.SpecMage ? MageRotation :
		T extends Spec.SpecRogue ? RogueRotation :
		T extends Spec.SpecRetributionPaladin ? RetributionPaladinRotation :
		T extends Spec.SpecShadowPriest ? ShadowPriestRotation :
		T extends Spec.SpecWarlock ? WarlockRotation :
		T extends Spec.SpecWarrior ? WarriorRotation :
		ElementalShamanRotation; // Should never reach this case

export type TalentsUnion =
		DruidTalents |
		HunterTalents |
		MageTalents |
		RogueTalents |
		PaladinTalents |
		PriestTalents |
		ShamanTalents |
		WarlockTalents |
		WarriorTalents;
export type SpecTalents<T extends Spec> =
		T extends Spec.SpecBalanceDruid ? DruidTalents :
		T extends Spec.SpecElementalShaman ? ShamanTalents :
		T extends Spec.SpecHunter ? HunterTalents :
		T extends Spec.SpecMage ? MageTalents :
		T extends Spec.SpecRogue ? RogueTalents :
		T extends Spec.SpecRetributionPaladin ? PaladinTalents :
		T extends Spec.SpecShadowPriest ? PriestTalents :
		T extends Spec.SpecWarlock ? WarlockTalents :
		T extends Spec.SpecWarrior ? WarriorTalents :
		ShamanTalents; // Should never reach this case

export type SpecOptionsUnion =
		BalanceDruidOptions |
		ElementalShamanOptions |
		HunterOptions |
		MageOptions |
		RogueOptions |
		RetributionPaladinOptions |
		ShadowPriestOptions |
		WarlockOptions |
		WarriorOptions;
export type SpecOptions<T extends Spec> =
		T extends Spec.SpecBalanceDruid ? BalanceDruidOptions :
		T extends Spec.SpecElementalShaman ? ElementalShamanOptions :
		T extends Spec.SpecHunter ? HunterOptions :
		T extends Spec.SpecMage ? MageOptions :
		T extends Spec.SpecRogue ? RogueOptions :
		T extends Spec.SpecRetributionPaladin ? RetributionPaladinOptions :
		T extends Spec.SpecShadowPriest ? ShadowPriestOptions :
		T extends Spec.SpecWarlock ? WarlockOptions :
		T extends Spec.SpecWarrior ? WarriorOptions :
		ElementalShamanOptions; // Should never reach this case

export type SpecProtoUnion =
		BalanceDruid |
		ElementalShaman |
		Hunter |
		Mage |
		Rogue |
		RetributionPaladin |
		ShadowPriest |
		Warlock |
		Warrior;
export type SpecProto<T extends Spec> =
		T extends Spec.SpecBalanceDruid ? BalanceDruid :
		T extends Spec.SpecElementalShaman ? ElementalShaman :
		T extends Spec.SpecHunter ? Hunter :
		T extends Spec.SpecMage ? Mage :
		T extends Spec.SpecRogue ? Rogue :
		T extends Spec.SpecRetributionPaladin ? RetributionPaladin :
		T extends Spec.SpecShadowPriest ? ShadowPriest :
		T extends Spec.SpecWarlock ? Warlock :
		T extends Spec.SpecWarrior ? Warrior :
		ElementalShaman; // Should never reach this case

export type SpecTypeFunctions<SpecType extends Spec> = {
  rotationCreate: () => SpecRotation<SpecType>;
  rotationEquals: (a: SpecRotation<SpecType>, b: SpecRotation<SpecType>) => boolean;
  rotationCopy: (a: SpecRotation<SpecType>) => SpecRotation<SpecType>;
  rotationToJson: (a: SpecRotation<SpecType>) => any;
  rotationFromJson: (obj: any) => SpecRotation<SpecType>;

  talentsCreate: () => SpecTalents<SpecType>;
  talentsEquals: (a: SpecTalents<SpecType>, b: SpecTalents<SpecType>) => boolean;
  talentsCopy: (a: SpecTalents<SpecType>) => SpecTalents<SpecType>;
  talentsToJson: (a: SpecTalents<SpecType>) => any;
  talentsFromJson: (obj: any) => SpecTalents<SpecType>;

  optionsCreate: () => SpecOptions<SpecType>;
  optionsEquals: (a: SpecOptions<SpecType>, b: SpecOptions<SpecType>) => boolean;
  optionsCopy: (a: SpecOptions<SpecType>) => SpecOptions<SpecType>;
  optionsToJson: (a: SpecOptions<SpecType>) => any;
  optionsFromJson: (obj: any) => SpecOptions<SpecType>;
};

export const specTypeFunctions: Partial<Record<Spec, SpecTypeFunctions<any>>> = {
  [Spec.SpecBalanceDruid]: {
    rotationCreate: () => BalanceDruidRotation.create(),
    rotationEquals: (a, b) => BalanceDruidRotation.equals(a as BalanceDruidRotation, b as BalanceDruidRotation),
    rotationCopy: (a) => BalanceDruidRotation.clone(a as BalanceDruidRotation),
    rotationToJson: (a) => BalanceDruidRotation.toJson(a as BalanceDruidRotation),
    rotationFromJson: (obj) => BalanceDruidRotation.fromJson(obj),

    talentsCreate: () => DruidTalents.create(),
    talentsEquals: (a, b) => DruidTalents.equals(a as DruidTalents, b as DruidTalents),
    talentsCopy: (a) => DruidTalents.clone(a as DruidTalents),
    talentsToJson: (a) => DruidTalents.toJson(a as DruidTalents),
    talentsFromJson: (obj) => DruidTalents.fromJson(obj),

    optionsCreate: () => BalanceDruidOptions.create(),
    optionsEquals: (a, b) => BalanceDruidOptions.equals(a as BalanceDruidOptions, b as BalanceDruidOptions),
    optionsCopy: (a) => BalanceDruidOptions.clone(a as BalanceDruidOptions),
    optionsToJson: (a) => BalanceDruidOptions.toJson(a as BalanceDruidOptions),
    optionsFromJson: (obj) => BalanceDruidOptions.fromJson(obj),
  },
  [Spec.SpecElementalShaman]: {
    rotationCreate: () => ElementalShamanRotation.create(),
    rotationEquals: (a, b) => ElementalShamanRotation.equals(a as ElementalShamanRotation, b as ElementalShamanRotation),
    rotationCopy: (a) => ElementalShamanRotation.clone(a as ElementalShamanRotation),
    rotationToJson: (a) => ElementalShamanRotation.toJson(a as ElementalShamanRotation),
    rotationFromJson: (obj) => ElementalShamanRotation.fromJson(obj),

    talentsCreate: () => ShamanTalents.create(),
    talentsEquals: (a, b) => ShamanTalents.equals(a as ShamanTalents, b as ShamanTalents),
    talentsCopy: (a) => ShamanTalents.clone(a as ShamanTalents),
    talentsToJson: (a) => ShamanTalents.toJson(a as ShamanTalents),
    talentsFromJson: (obj) => ShamanTalents.fromJson(obj),

    optionsCreate: () => ElementalShamanOptions.create(),
    optionsEquals: (a, b) => ElementalShamanOptions.equals(a as ElementalShamanOptions, b as ElementalShamanOptions),
    optionsCopy: (a) => ElementalShamanOptions.clone(a as ElementalShamanOptions),
    optionsToJson: (a) => ElementalShamanOptions.toJson(a as ElementalShamanOptions),
    optionsFromJson: (obj) => ElementalShamanOptions.fromJson(obj),
  },
  [Spec.SpecHunter]: {
    rotationCreate: () => HunterRotation.create(),
    rotationEquals: (a, b) => HunterRotation.equals(a as HunterRotation, b as HunterRotation),
    rotationCopy: (a) => HunterRotation.clone(a as HunterRotation),
    rotationToJson: (a) => HunterRotation.toJson(a as HunterRotation),
    rotationFromJson: (obj) => HunterRotation.fromJson(obj),

    talentsCreate: () => HunterTalents.create(),
    talentsEquals: (a, b) => HunterTalents.equals(a as HunterTalents, b as HunterTalents),
    talentsCopy: (a) => HunterTalents.clone(a as HunterTalents),
    talentsToJson: (a) => HunterTalents.toJson(a as HunterTalents),
    talentsFromJson: (obj) => HunterTalents.fromJson(obj),

    optionsCreate: () => HunterOptions.create(),
    optionsEquals: (a, b) => HunterOptions.equals(a as HunterOptions, b as HunterOptions),
    optionsCopy: (a) => HunterOptions.clone(a as HunterOptions),
    optionsToJson: (a) => HunterOptions.toJson(a as HunterOptions),
    optionsFromJson: (obj) => HunterOptions.fromJson(obj),
  },
  [Spec.SpecMage]: {
    rotationCreate: () => MageRotation.create(),
    rotationEquals: (a, b) => MageRotation.equals(a as MageRotation, b as MageRotation),
    rotationCopy: (a) => MageRotation.clone(a as MageRotation),
    rotationToJson: (a) => MageRotation.toJson(a as MageRotation),
    rotationFromJson: (obj) => MageRotation.fromJson(obj),

    talentsCreate: () => MageTalents.create(),
    talentsEquals: (a, b) => MageTalents.equals(a as MageTalents, b as MageTalents),
    talentsCopy: (a) => MageTalents.clone(a as MageTalents),
    talentsToJson: (a) => MageTalents.toJson(a as MageTalents),
    talentsFromJson: (obj) => MageTalents.fromJson(obj),

    optionsCreate: () => MageOptions.create(),
    optionsEquals: (a, b) => MageOptions.equals(a as MageOptions, b as MageOptions),
    optionsCopy: (a) => MageOptions.clone(a as MageOptions),
    optionsToJson: (a) => MageOptions.toJson(a as MageOptions),
    optionsFromJson: (obj) => MageOptions.fromJson(obj),
  },
  [Spec.SpecRetributionPaladin]: {
    rotationCreate: () => RetributionPaladinRotation.create(),
    rotationEquals: (a, b) => RetributionPaladinRotation.equals(a as RetributionPaladinRotation, b as RetributionPaladinRotation),
    rotationCopy: (a) => RetributionPaladinRotation.clone(a as RetributionPaladinRotation),
    rotationToJson: (a) => RetributionPaladinRotation.toJson(a as RetributionPaladinRotation),
    rotationFromJson: (obj) => RetributionPaladinRotation.fromJson(obj),

    talentsCreate: () => PaladinTalents.create(),
    talentsEquals: (a, b) => PaladinTalents.equals(a as PaladinTalents, b as PaladinTalents),
    talentsCopy: (a) => PaladinTalents.clone(a as PaladinTalents),
    talentsToJson: (a) => PaladinTalents.toJson(a as PaladinTalents),
    talentsFromJson: (obj) => PaladinTalents.fromJson(obj),

    optionsCreate: () => RetributionPaladinOptions.create(),
    optionsEquals: (a, b) => RetributionPaladinOptions.equals(a as RetributionPaladinOptions, b as RetributionPaladinOptions),
    optionsCopy: (a) => RetributionPaladinOptions.clone(a as RetributionPaladinOptions),
    optionsToJson: (a) => RetributionPaladinOptions.toJson(a as RetributionPaladinOptions),
    optionsFromJson: (obj) => RetributionPaladinOptions.fromJson(obj),
  },
  [Spec.SpecRogue]: {
    rotationCreate: () => RogueRotation.create(),
    rotationEquals: (a, b) => RogueRotation.equals(a as RogueRotation, b as RogueRotation),
    rotationCopy: (a) => RogueRotation.clone(a as RogueRotation),
    rotationToJson: (a) => RogueRotation.toJson(a as RogueRotation),
    rotationFromJson: (obj) => RogueRotation.fromJson(obj),

    talentsCreate: () => RogueTalents.create(),
    talentsEquals: (a, b) => RogueTalents.equals(a as RogueTalents, b as RogueTalents),
    talentsCopy: (a) => RogueTalents.clone(a as RogueTalents),
    talentsToJson: (a) => RogueTalents.toJson(a as RogueTalents),
    talentsFromJson: (obj) => RogueTalents.fromJson(obj),

    optionsCreate: () => RogueOptions.create(),
    optionsEquals: (a, b) => RogueOptions.equals(a as RogueOptions, b as RogueOptions),
    optionsCopy: (a) => RogueOptions.clone(a as RogueOptions),
    optionsToJson: (a) => RogueOptions.toJson(a as RogueOptions),
    optionsFromJson: (obj) => RogueOptions.fromJson(obj),
  },
  [Spec.SpecShadowPriest]: {
    rotationCreate: () => ShadowPriestRotation.create(),
    rotationEquals: (a, b) => ShadowPriestRotation.equals(a as ShadowPriestRotation, b as ShadowPriestRotation),
    rotationCopy: (a) => ShadowPriestRotation.clone(a as ShadowPriestRotation),
    rotationToJson: (a) => ShadowPriestRotation.toJson(a as ShadowPriestRotation),
    rotationFromJson: (obj) => ShadowPriestRotation.fromJson(obj),

    talentsCreate: () => PriestTalents.create(),
    talentsEquals: (a, b) => PriestTalents.equals(a as PriestTalents, b as PriestTalents),
    talentsCopy: (a) => PriestTalents.clone(a as PriestTalents),
    talentsToJson: (a) => PriestTalents.toJson(a as PriestTalents),
    talentsFromJson: (obj) => PriestTalents.fromJson(obj),

    optionsCreate: () => ShadowPriestOptions.create(),
    optionsEquals: (a, b) => ShadowPriestOptions.equals(a as ShadowPriestOptions, b as ShadowPriestOptions),
    optionsCopy: (a) => ShadowPriestOptions.clone(a as ShadowPriestOptions),
    optionsToJson: (a) => ShadowPriestOptions.toJson(a as ShadowPriestOptions),
    optionsFromJson: (obj) => ShadowPriestOptions.fromJson(obj),
  },
  [Spec.SpecWarlock]: {
    rotationCreate: () => WarlockRotation.create(),
    rotationEquals: (a, b) => WarlockRotation.equals(a as WarlockRotation, b as WarlockRotation),
    rotationCopy: (a) => WarlockRotation.clone(a as WarlockRotation),
    rotationToJson: (a) => WarlockRotation.toJson(a as WarlockRotation),
    rotationFromJson: (obj) => WarlockRotation.fromJson(obj),

    talentsCreate: () => WarlockTalents.create(),
    talentsEquals: (a, b) => WarlockTalents.equals(a as WarlockTalents, b as WarlockTalents),
    talentsCopy: (a) => WarlockTalents.clone(a as WarlockTalents),
    talentsToJson: (a) => WarlockTalents.toJson(a as WarlockTalents),
    talentsFromJson: (obj) => WarlockTalents.fromJson(obj),

    optionsCreate: () => WarlockOptions.create(),
    optionsEquals: (a, b) => WarlockOptions.equals(a as WarlockOptions, b as WarlockOptions),
    optionsCopy: (a) => WarlockOptions.clone(a as WarlockOptions),
    optionsToJson: (a) => WarlockOptions.toJson(a as WarlockOptions),
    optionsFromJson: (obj) => WarlockOptions.fromJson(obj),
  },
  [Spec.SpecWarrior]: {
    rotationCreate: () => WarriorRotation.create(),
    rotationEquals: (a, b) => WarriorRotation.equals(a as WarriorRotation, b as WarriorRotation),
    rotationCopy: (a) => WarriorRotation.clone(a as WarriorRotation),
    rotationToJson: (a) => WarriorRotation.toJson(a as WarriorRotation),
    rotationFromJson: (obj) => WarriorRotation.fromJson(obj),

    talentsCreate: () => WarriorTalents.create(),
    talentsEquals: (a, b) => WarriorTalents.equals(a as WarriorTalents, b as WarriorTalents),
    talentsCopy: (a) => WarriorTalents.clone(a as WarriorTalents),
    talentsToJson: (a) => WarriorTalents.toJson(a as WarriorTalents),
    talentsFromJson: (obj) => WarriorTalents.fromJson(obj),

    optionsCreate: () => WarriorOptions.create(),
    optionsEquals: (a, b) => WarriorOptions.equals(a as WarriorOptions, b as WarriorOptions),
    optionsCopy: (a) => WarriorOptions.clone(a as WarriorOptions),
    optionsToJson: (a) => WarriorOptions.toJson(a as WarriorOptions),
    optionsFromJson: (obj) => WarriorOptions.fromJson(obj),
  },
};

export const specToClass: Record<Spec, Class> = {
  [Spec.SpecBalanceDruid]: Class.ClassDruid,
  [Spec.SpecElementalShaman]: Class.ClassShaman,
  [Spec.SpecHunter]: Class.ClassHunter,
  [Spec.SpecMage]: Class.ClassMage,
  [Spec.SpecRogue]: Class.ClassRogue,
  [Spec.SpecRetributionPaladin]: Class.ClassPaladin,
  [Spec.SpecShadowPriest]: Class.ClassPriest,
  [Spec.SpecWarlock]: Class.ClassWarlock,
  [Spec.SpecWarrior]: Class.ClassWarrior,
};

const druidRaces = [
    Race.RaceNightElf,
    Race.RaceTauren,
];
const hunterRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceDwarf,
    Race.RaceNightElf,
    Race.RaceOrc,
    Race.RaceTauren,
    Race.RaceTroll10,
    Race.RaceTroll30,
];
const mageRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];
const paladinRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceDwarf,
    Race.RaceHuman,
];
const priestRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceDwarf,
    Race.RaceHuman,
    Race.RaceNightElf,
    Race.RaceOrc,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];
const rogueRaces = [
    Race.RaceBloodElf,
    Race.RaceDwarf,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceNightElf,
    Race.RaceOrc,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];
const shamanRaces = [
    Race.RaceDraenei,
    Race.RaceOrc,
    Race.RaceTauren,
    Race.RaceTroll10,
    Race.RaceTroll30,
];
const warlockRaces = [
    Race.RaceBloodElf,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceOrc,
    Race.RaceUndead,
];
const warriorRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceDwarf,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceNightElf,
    Race.RaceOrc,
    Race.RaceTauren,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];

export const specToEligibleRaces: Record<Spec, Array<Race>> = {
  [Spec.SpecBalanceDruid]: druidRaces,
  [Spec.SpecElementalShaman]: shamanRaces,
  [Spec.SpecHunter]: hunterRaces,
  [Spec.SpecMage]: mageRaces,
  [Spec.SpecRetributionPaladin]: paladinRaces,
  [Spec.SpecRogue]: rogueRaces,
  [Spec.SpecShadowPriest]: priestRaces,
  [Spec.SpecWarlock]: warlockRaces,
  [Spec.SpecWarrior]: warriorRaces,
};

// Prefixes used for storing browser data for each site. Even if a Spec is
// renamed, DO NOT change these values or people will lose their saved data.
export const specToLocalStorageKey: Record<Spec, string> = {
  [Spec.SpecBalanceDruid]: '__balance_druid',
  [Spec.SpecElementalShaman]: '__elemental_shaman',
  [Spec.SpecHunter]: '__hunter',
  [Spec.SpecMage]: '__mage',
  [Spec.SpecRetributionPaladin]: '__retribution_paladin',
  [Spec.SpecRogue]: '__rogue',
  [Spec.SpecShadowPriest]: '__shadow_priest',
  [Spec.SpecWarlock]: '__warlock',
  [Spec.SpecWarrior]: '__warrior',
};

// Returns a copy of playerOptions, with the class field set.
export function withSpecProto<SpecType extends Spec>(
    playerOptions: PlayerOptions,
    rotation: SpecRotation<SpecType>,
    talents: SpecTalents<SpecType>,
    specOptions: SpecOptions<SpecType>): PlayerOptions {
  const copy = PlayerOptions.clone(playerOptions);
  if (BalanceDruidRotation.is(rotation)) {
		copy.class = Class.ClassDruid;
    copy.spec = {
      oneofKind: 'balanceDruid',
      balanceDruid: BalanceDruid.create({
        rotation: rotation,
        talents: talents as DruidTalents,
        options: specOptions as BalanceDruidOptions,
      }),
    };
  } else if (ElementalShamanRotation.is(rotation)) {
		copy.class = Class.ClassShaman;
    copy.spec = {
      oneofKind: 'elementalShaman',
      elementalShaman: ElementalShaman.create({
        rotation: rotation,
        talents: talents as ShamanTalents,
        options: specOptions as ElementalShamanOptions,
      }),
    };
  } else if (HunterRotation.is(rotation)) {
		copy.class = Class.ClassHunter;
    copy.spec = {
      oneofKind: 'hunter',
      hunter: Hunter.create({
        rotation: rotation,
        talents: talents as HunterTalents,
        options: specOptions as HunterOptions,
      }),
    };
  } else if (MageRotation.is(rotation)) {
		copy.class = Class.ClassMage;
    copy.spec = {
      oneofKind: 'mage',
      mage: Mage.create({
        rotation: rotation,
        talents: talents as MageTalents,
        options: specOptions as MageOptions,
      }),
    };
  } else if (RetributionPaladinRotation.is(rotation)) {
		copy.class = Class.ClassPaladin;
    copy.spec = {
      oneofKind: 'retributionPaladin',
      retributionPaladin: RetributionPaladin.create({
        rotation: rotation,
        talents: talents as PaladinTalents,
        options: specOptions as RetributionPaladinOptions,
      }),
    };
  } else if (RogueRotation.is(rotation)) {
		copy.class = Class.ClassRogue;
    copy.spec = {
      oneofKind: 'rogue',
      rogue: Rogue.create({
        rotation: rotation,
        talents: talents as RogueTalents,
        options: specOptions as RogueOptions,
      }),
    };
  } else if (ShadowPriestRotation.is(rotation)) {
		copy.class = Class.ClassPriest;
    copy.spec = {
      oneofKind: 'shadowPriest',
      shadowPriest: ShadowPriest.create({
        rotation: rotation,
        talents: talents as PriestTalents,
        options: specOptions as ShadowPriestOptions,
      }),
    };
  } else if (WarlockRotation.is(rotation)) {
		copy.class = Class.ClassWarlock;
    copy.spec = {
      oneofKind: 'warlock',
      warlock: Warlock.create({
        rotation: rotation,
        talents: talents as WarlockTalents,
        options: specOptions as WarlockOptions,
      }),
    };
  } else if (WarriorRotation.is(rotation)) {
		copy.class = Class.ClassWarrior;
    copy.spec = {
      oneofKind: 'warrior',
      warrior: Warrior.create({
        rotation: rotation,
        talents: talents as WarriorTalents,
        options: specOptions as WarriorOptions,
      }),
    };
  } else {
    throw new Error('Unrecognized talents with options: ' + PlayerOptions.toJsonString(playerOptions));
  }
  return copy;
}

const itemTypeToSlotsMap: Partial<Record<ItemType, Array<ItemSlot>>> = {
  [ItemType.ItemTypeUnknown]: [],
  [ItemType.ItemTypeHead]: [ItemSlot.ItemSlotHead],
  [ItemType.ItemTypeNeck]: [ItemSlot.ItemSlotNeck],
  [ItemType.ItemTypeShoulder]: [ItemSlot.ItemSlotShoulder],
  [ItemType.ItemTypeBack]: [ItemSlot.ItemSlotBack],
  [ItemType.ItemTypeChest]: [ItemSlot.ItemSlotChest],
  [ItemType.ItemTypeWrist]: [ItemSlot.ItemSlotWrist],
  [ItemType.ItemTypeHands]: [ItemSlot.ItemSlotHands],
  [ItemType.ItemTypeWaist]: [ItemSlot.ItemSlotWaist],
  [ItemType.ItemTypeLegs]: [ItemSlot.ItemSlotLegs],
  [ItemType.ItemTypeFeet]: [ItemSlot.ItemSlotFeet],
  [ItemType.ItemTypeFinger]: [ItemSlot.ItemSlotFinger1, ItemSlot.ItemSlotFinger2],
  [ItemType.ItemTypeTrinket]: [ItemSlot.ItemSlotTrinket1, ItemSlot.ItemSlotTrinket2],
  [ItemType.ItemTypeRanged]: [ItemSlot.ItemSlotRanged],
};

export function getEligibleItemSlots(item: Item): Array<ItemSlot> {
  if (itemTypeToSlotsMap[item.type]) {
    return itemTypeToSlotsMap[item.type]!;
  }

  if (item.type == ItemType.ItemTypeWeapon) {
    if ([HandType.HandTypeMainHand, HandType.HandTypeTwoHand].includes(item.handType)) {
      return [ItemSlot.ItemSlotMainHand];
    } else if (item.handType == HandType.HandTypeOffHand) {
      return [ItemSlot.ItemSlotOffHand];
    } else {
      return [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand];
    }
  }

  // Should never reach here
  throw new Error('Could not find item slots for item: ' + Item.toJsonString(item));
};

// Returns whether the given main-hand and off-hand items can be worn at the
// same time.
export function validWeaponCombo(mainHand: Item | null | undefined, offHand: Item | null | undefined): boolean {
	if (mainHand == null || offHand == null) {
		return true;
	}

	if (mainHand.handType == HandType.HandTypeTwoHand) {
		return false;
	}

	return true;
}

// Returns all item slots to which the enchant might be applied.
// 
// Note that this alone is not enough; some items have further restrictions,
// e.g. some weapon enchants may only be applied to 2H weapons.
export function getEligibleEnchantSlots(enchant: Enchant): Array<ItemSlot> {
  if (itemTypeToSlotsMap[enchant.type]) {
    return itemTypeToSlotsMap[enchant.type]!;
  }

  if (enchant.type == ItemType.ItemTypeWeapon) {
    return [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand];
  }

  // Should never reach here
  throw new Error('Could not find item slots for enchant: ' + Enchant.toJsonString(enchant));
};

export function enchantAppliesToItem(enchant: Enchant, item: Item): boolean {
  const sharedSlots = intersection(getEligibleEnchantSlots(enchant), getEligibleItemSlots(item));
  if (sharedSlots.length == 0)
    return false;

  if (sharedSlots.includes(ItemSlot.ItemSlotMainHand)) {
    if (enchant.twoHandedOnly && item.handType != HandType.HandTypeTwoHand)
      return false;
  }

  if (sharedSlots.includes(ItemSlot.ItemSlotOffHand)) {
    if (enchant.shieldOnly && item.weaponType != WeaponType.WeaponTypeShield)
      return false;
  }

  if (sharedSlots.includes(ItemSlot.ItemSlotRanged)) {
    if (![
          RangedWeaponType.RangedWeaponTypeBow,
          RangedWeaponType.RangedWeaponTypeCrossbow,
          RangedWeaponType.RangedWeaponTypeGun,
        ].includes(item.rangedWeaponType))
      return false;
  }

  return true;
};

const socketToMatchingColors = new Map<GemColor, Array<GemColor>>();
socketToMatchingColors.set(GemColor.GemColorMeta,   [GemColor.GemColorMeta]);
socketToMatchingColors.set(GemColor.GemColorBlue,   [GemColor.GemColorBlue, GemColor.GemColorPurple, GemColor.GemColorGreen]);
socketToMatchingColors.set(GemColor.GemColorRed,    [GemColor.GemColorRed, GemColor.GemColorPurple, GemColor.GemColorOrange]);
socketToMatchingColors.set(GemColor.GemColorYellow, [GemColor.GemColorYellow, GemColor.GemColorOrange, GemColor.GemColorGreen]);

// Whether the gem matches the given socket color, for the purposes of gaining the socket bonuses.
export function gemMatchesSocket(gem: Gem, socketColor: GemColor) {
  return socketToMatchingColors.has(socketColor) && socketToMatchingColors.get(socketColor)!.includes(gem.color);
}

// Whether the gem is capable of slotting into a socket of the given color.
export function gemEligibleForSocket(gem: Gem, socketColor: GemColor) {
  return (gem.color == GemColor.GemColorMeta) == (socketColor == GemColor.GemColorMeta);
}
