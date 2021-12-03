import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { specIconsLarge } from '/tbc/core/proto_utils/utils.js';

import * as ElementalShamanPresets from '/tbc/elemental_shaman/presets.js';

import { RaidSimUI } from './raid_sim_ui.js';

const ui = new RaidSimUI(document.body, {
	specs: [
		{
			spec: Spec.SpecElementalShaman,
			rotation: ElementalShamanPresets.DefaultRotation,
			talents: ElementalShamanPresets.StandardTalents.data,
			specOptions: ElementalShamanPresets.DefaultOptions,
			consumes: ElementalShamanPresets.DefaultConsumes,
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
			iconUrl: specIconsLarge[Spec.SpecElementalShaman],
		},
	],
});
ui.init();
