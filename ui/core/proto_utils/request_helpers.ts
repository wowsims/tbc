import { Buffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';

import { Player } from '/tbc/core/proto/api.js';
import { PlayerOptions } from '/tbc/core/proto/api.js';
import { BalanceDruid, BalanceDruid_Agent as BalanceDruidAgent, DruidTalents, BalanceDruid_Options as BalanceDruidOptions} from '/tbc/core/proto/druid.js';
import { ElementalShaman, ElementalShaman_Agent as ElementalShamanAgent, ShamanTalents, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';

import { ComputeStatsRequest, ComputeStatsResult } from '/tbc/core/proto/api.js';
import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';

import { Gear } from './gear.js';
import { Stats } from './stats.js';
import { SpecAgent } from './utils.js';
import { SpecTalents } from './utils.js';
import { SpecOptions } from './utils.js';

export function makeComputeStatsRequest<SpecType extends Spec>(
    buffs: Buffs,
    consumes: Consumes,
    customStats: Stats,
    encounter: Encounter,
    gear: Gear,
    race: Race,
    agent: SpecAgent<SpecType>,
    talents: SpecTalents<SpecType>,
    classOptions: SpecOptions<SpecType>): ComputeStatsRequest {
  return ComputeStatsRequest.create({
    player: Player.create({
      customStats: customStats.asArray(),
      equipment: gear.asSpec(),
      options: withSpecProto(PlayerOptions.create({
        consumes: consumes,
        race: race,
      }), agent, talents, classOptions),
    }),
    buffs: buffs,
  });
}

export function makeIndividualSimRequest<SpecType extends Spec>(
    buffs: Buffs,
    consumes: Consumes,
    customStats: Stats,
    encounter: Encounter,
    gear: Gear,
    race: Race,
    agent: SpecAgent<SpecType>,
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
      }), agent, talents, classOptions),
    }),
    buffs: buffs,
    encounter: encounter,
    iterations: iterations,
    gcdMin: 0.75,
    debug: debug,
  });
}

// Returns a copy of playerOptions, with the class field set.
function withSpecProto<SpecType extends Spec>(
    playerOptions: PlayerOptions,
    agent: SpecAgent<SpecType>,
    talents: SpecTalents<SpecType>,
    specOptions: SpecOptions<SpecType>): PlayerOptions {
  const copy = PlayerOptions.clone(playerOptions);
  if (BalanceDruidAgent.is(agent)) {
    copy.spec = {
      oneofKind: 'balanceDruid',
      balanceDruid: BalanceDruid.create({
        agent: agent,
        talents: talents as DruidTalents,
        options: specOptions as BalanceDruidOptions,
      }),
    };
  } else if (ElementalShamanAgent.is(agent)) {
    copy.spec = {
      oneofKind: 'elementalShaman',
      elementalShaman: ElementalShaman.create({
        agent: agent,
        talents: talents as ShamanTalents,
        options: specOptions as ElementalShamanOptions,
      }),
    };
  } else {
    throw new Error('Unrecognized talents with options: ' + PlayerOptions.toJsonString(playerOptions));
  }
  return copy;
}
