import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { specIconsLarge } from '/tbc/core/proto_utils/utils.js';
import { specNames } from '/tbc/core/proto_utils/utils.js';

import * as BalanceDruidPresets from '/tbc/balance_druid/presets.js';
import * as ElementalShamanPresets from '/tbc/elemental_shaman/presets.js';
import * as ShadowPriestPresets from '/tbc/shadow_priest/presets.js';

import { RaidSimUI } from './raid_sim_ui.js';

const ui = new RaidSimUI(document.body, {
	presets: [
		{
			spec: Spec.SpecBalanceDruid,
			rotation: BalanceDruidPresets.DefaultRotation,
			talents: BalanceDruidPresets.StandardTalents.data,
			specOptions: BalanceDruidPresets.DefaultOptions,
			consumes: BalanceDruidPresets.DefaultConsumes,
			defaultName: specNames[Spec.SpecBalanceDruid],
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceNightElf,
				[Faction.Horde]: Race.RaceTauren,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: BalanceDruidPresets.P1_ALLIANCE_BIS.gear,
					2: BalanceDruidPresets.P2_ALLIANCE_BIS.gear,
				},
				[Faction.Horde]: {
					1: BalanceDruidPresets.P1_HORDE_BIS.gear,
					2: BalanceDruidPresets.P2_HORDE_BIS.gear,
				},
			},
			tooltip: specNames[Spec.SpecBalanceDruid],
			iconUrl: specIconsLarge[Spec.SpecBalanceDruid],
		},
		{
			spec: Spec.SpecElementalShaman,
			rotation: ElementalShamanPresets.DefaultRotation,
			talents: ElementalShamanPresets.StandardTalents.data,
			specOptions: ElementalShamanPresets.DefaultOptions,
			consumes: ElementalShamanPresets.DefaultConsumes,
			defaultName: specNames[Spec.SpecElementalShaman],
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceDraenei,
				[Faction.Horde]: Race.RaceOrc,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: ElementalShamanPresets.P1_BIS.gear,
					2: ElementalShamanPresets.P2_BIS.gear,
				},
				[Faction.Horde]: {
					1: ElementalShamanPresets.P1_BIS.gear,
					2: ElementalShamanPresets.P2_BIS.gear,
				},
			},
			tooltip: specNames[Spec.SpecElementalShaman],
			iconUrl: specIconsLarge[Spec.SpecElementalShaman],
		},
		{
			spec: Spec.SpecShadowPriest,
			rotation: ShadowPriestPresets.DefaultRotation,
			talents: ShadowPriestPresets.StandardTalents.data,
			specOptions: ShadowPriestPresets.DefaultOptions,
			consumes: ShadowPriestPresets.DefaultConsumes,
			defaultName: specNames[Spec.SpecShadowPriest],
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceDwarf,
				[Faction.Horde]: Race.RaceUndead,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: ShadowPriestPresets.P1_BIS.gear,
					2: ShadowPriestPresets.P2_BIS.gear,
				},
				[Faction.Horde]: {
					1: ShadowPriestPresets.P1_BIS.gear,
					2: ShadowPriestPresets.P2_BIS.gear,
				},
			},
			tooltip: specNames[Spec.SpecShadowPriest],
			iconUrl: specIconsLarge[Spec.SpecShadowPriest],
		},
	],
});
