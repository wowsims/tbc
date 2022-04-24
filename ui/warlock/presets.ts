import { Consumes } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { Warlock, Warlock_Rotation as WarlockRotation, WarlockTalents as WarlockTalents, Warlock_Options as WarlockOptions, Warlock_Rotation_PrimarySpell, Warlock_Rotation_Curse, Warlock_Options_Armor as Armor, Warlock_Options_Summon as Summon } from '/tbc/core/proto/warlock.js';

import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
import { FelArmor, Sacrifice } from './inputs';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const DestructionTalents = {
	name: 'Destruction',
	data: '-20501301332001-50500051220051053105',
};

export const AfflicationTalents = {
	name: 'Afflication',
	data: '55022000102351055103--50500051220001',
}

export const DemonologistTalents = {
	name: 'Demonologist',
	data: '01-2050030133250101501351-5050005112',
}


export const DefaultRotation = WarlockRotation.create({
	primarySpell: Warlock_Rotation_PrimarySpell.Shadowbolt,
	immolate: true,
	curse: Warlock_Rotation_Curse.Elements,
});

export const DefaultOptions = WarlockOptions.create({
	armor: Armor.FelArmor,
	sacrificeSummon: true,
	summon: Summon.Succubus,
});

export const DefaultConsumes = Consumes.create({
	flask: Flask.FlaskOfPureDeath,
	food: Food.FoodBlackenedBasilisk,
	mainHandImbue: WeaponImbue.WeaponImbueBrilliantWizardOil,
	defaultPotion: Potions.DestructionPotion,
});

export const P3_DESTRO = {
	name: 'P3 Destro',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(`{"items": [
        {
          "id": 31051,
          "enchant": 29191,
          "gems": [
            32218,
            34220
          ]
        },
        {
          "id": 30015
        },
        {
          "id": 31054,
          "enchant": 28886,
          "gems": [
            32215,
            32218
          ]
        },
        {
          "id": 32331,
          "enchant": 33150
        },
        {
          "id": 31052,
          "enchant": 24003,
          "gems": [
            32196,
            32196,
            32196
          ]
        },
        {
          "id": 32586,
          "enchant": 22534
        },
        {
          "id": 31050,
          "enchant": 28272,
          "gems": [
            32196
          ]
        },
        {
          "id": 30038,
          "gems": [
            32215,
            32218
          ]
        },
        {
          "id": 30916,
          "enchant": 24274,
          "gems": [
            32196,
            32196,
            32196
          ]
        },
        {
          "id": 32239,
          "enchant": 35297,
          "gems": [
            32218,
            32215
          ]
        },
        {
          "id": 32527,
          "enchant": 22536
        },
        {
          "id": 29305,
          "enchant": 22536
        },
        {
          "id": 32483
        },
        {
          "id": 29370
        },
        {
          "id": 32374,
          "enchant": 22560
        },
        {},
        {
          "id": 29982
        }
	]}`),
};

export const P4_DESTRO = {
	name: 'P4 Destro',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(`{"items": [
		{
		  "id": 31051,
		  "enchant": 29191,
		  "gems": [
			32218,
			34220
		  ]
		},
		{
		  "id": 33281
		},
		{
		  "id": 31054,
		  "enchant": 28886,
		  "gems": [
			32215,
			32218
		  ]
		},
		{
		  "id": 32524,
		  "enchant": 33150
		},
		{
		  "id": 30107,
		  "enchant": 24003,
		  "gems": [
			32196,
			32196,
			32196
		  ]
		},
		{
		  "id": 32586,
		  "enchant": 22534
		},
		{
		  "id": 31050,
		  "enchant": 28272,
		  "gems": [
			32196
		  ]
		},
		{
		  "id": 30888,
		  "gems": [
			32196,
			32196
		  ]
		},
		{
		  "id": 31053,
		  "enchant": 24274,
		  "gems": [
			32196
		  ]
		},
		{
		  "id": 32239,
		  "enchant": 35297,
		  "gems": [
			32218,
			32215
		  ]
		},
		{
		  "id": 32527,
		  "enchant": 22536
		},
		{
		  "id": 33497,
		  "enchant": 22536
		},
		{
		  "id": 32483
		},
		{
		  "id": 33829
		},
		{
		  "id": 32374,
		  "enchant": 22560
		},
		{},
		{
		  "id": 33192,
		  "gems": [
			32215
		  ]
		}
	]}`),
};
