import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';

import { Player } from '/tbc/core/proto/api.js';
import { PlayerOptions } from '/tbc/core/proto/api.js';
import { SimOptions } from '/tbc/core/proto/api.js';

import { ComputeStatsRequest, ComputeStatsResult } from '/tbc/core/proto/api.js';
import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';

import { Gear } from './gear.js';
import { Stats } from './stats.js';
import { SpecRotation } from './utils.js';
import { SpecTalents } from './utils.js';
import { SpecOptions } from './utils.js';
import { withSpecProto } from './utils.js';

export function makeComputeStatsRequest<SpecType extends Spec>(
    raidBuffs: RaidBuffs,
    partyBuffs: PartyBuffs,
    individualBuffs: IndividualBuffs,
    consumes: Consumes,
    customStats: Stats,
    gear: Gear,
    race: Race,
    rotation: SpecRotation<SpecType>,
    talents: SpecTalents<SpecType>,
    classOptions: SpecOptions<SpecType>): ComputeStatsRequest {
  return ComputeStatsRequest.create({
    player: Player.create({
      customStats: customStats.asArray(),
      equipment: gear.asSpec(),
      options: withSpecProto(PlayerOptions.create({
        consumes: consumes,
        race: race,
      }), rotation, talents, classOptions),
    }),
    raidBuffs: raidBuffs,
    partyBuffs: partyBuffs,
    individualBuffs: individualBuffs,
  });
}

export function makeIndividualSimRequest<SpecType extends Spec>(
    raidBuffs: RaidBuffs,
    partyBuffs: PartyBuffs,
    individualBuffs: IndividualBuffs,
    consumes: Consumes,
    customStats: Stats,
    encounter: Encounter,
    gear: Gear,
    race: Race,
    rotation: SpecRotation<SpecType>,
    talents: SpecTalents<SpecType>,
    classOptions: SpecOptions<SpecType>,
    iterations: number,
    debug: boolean): IndividualSimRequest {
  return IndividualSimRequest.create({
    player: Player.create({
      customStats: customStats.asArray(),
      equipment: gear.asSpec(),
      options: withSpecProto(PlayerOptions.create({
        consumes: consumes,
        race: race,
      }), rotation, talents, classOptions),
    }),
    raidBuffs: raidBuffs,
    partyBuffs: partyBuffs,
    individualBuffs: individualBuffs,
    encounter: encounter,
		simOptions: SimOptions.create({
			iterations: iterations,
			debug: debug,
		}),
  });
}
