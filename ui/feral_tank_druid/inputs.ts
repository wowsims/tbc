import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Player } from '/tbc/core/player.js';
import { Sim } from '/tbc/core/sim.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Target } from '/tbc/core/target.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { ItemSlot } from '/tbc/core/proto/common.js';

import {
	FeralTankDruid,
	FeralTankDruid_Rotation as DruidRotation,
	FeralTankDruid_Options as DruidOptions
} from '/tbc/core/proto/druid.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const FeralTankDruidRotationConfig = {
	inputs: [
	],
};
