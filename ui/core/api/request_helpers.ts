import { Gear } from './gear';
import { Buffs } from './common';
import { Consumes } from './common';
import { Class } from './common';
import { Encounter } from './common';
import { Race } from './common';
import { Spec } from './common';
import { Stats } from './stats';
import { SpecAgent } from './utils';
import { SpecTalents } from './utils';
import { SpecOptions } from './utils';

import { Player } from './api';
import { PlayerOptions } from './api';
import { BalanceDruid, BalanceDruid_BalanceDruidAgent as BalanceDruidAgent, DruidTalents, BalanceDruid_BalanceDruidOptions as BalanceDruidOptions} from './druid';
import { ElementalShaman, ElementalShaman_ElementalShamanAgent as ElementalShamanAgent, ShamanTalents, ElementalShaman_ElementalShamanOptions as ElementalShamanOptions } from './shaman';

import { ComputeStatsRequest, ComputeStatsResult } from './api';
import { IndividualSimRequest, IndividualSimResult } from './api';

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
