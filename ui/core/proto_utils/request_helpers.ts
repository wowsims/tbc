import { Buffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';

import { Player } from '/tbc/core/proto/api.js';
import { PlayerOptions } from '/tbc/core/proto/api.js';
import { BalanceDruid, BalanceDruid_Rotation as BalanceDruidRotation, DruidTalents, BalanceDruid_Options as BalanceDruidOptions} from '/tbc/core/proto/druid.js';
import { ElementalShaman, ElementalShaman_Rotation as ElementalShamanRotation, ShamanTalents, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';

import { ComputeStatsRequest, ComputeStatsResult } from '/tbc/core/proto/api.js';
import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';

import { Gear } from './gear.js';
import { Stats } from './stats.js';
import { SpecRotation } from './utils.js';
import { SpecTalents } from './utils.js';
import { SpecOptions } from './utils.js';

export function makeComputeStatsRequest<SpecType extends Spec>(
    buffs: Buffs,
    consumes: Consumes,
    customStats: Stats,
    encounter: Encounter,
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
    buffs: buffs,
    encounter: encounter,
    iterations: iterations,
    debug: debug,
  });
}

// Returns a copy of playerOptions, with the class field set.
function withSpecProto<SpecType extends Spec>(
    playerOptions: PlayerOptions,
    rotation: SpecRotation<SpecType>,
    talents: SpecTalents<SpecType>,
    specOptions: SpecOptions<SpecType>): PlayerOptions {
  const copy = PlayerOptions.clone(playerOptions);
  if (BalanceDruidRotation.is(rotation)) {
    copy.spec = {
      oneofKind: 'balanceDruid',
      balanceDruid: BalanceDruid.create({
        rotation: rotation,
        talents: talents as DruidTalents,
        options: specOptions as BalanceDruidOptions,
      }),
    };
  } else if (ElementalShamanRotation.is(rotation)) {
    copy.spec = {
      oneofKind: 'elementalShaman',
      elementalShaman: ElementalShaman.create({
        rotation: rotation,
        talents: talents as ShamanTalents,
        options: specOptions as ElementalShamanOptions,
      }),
    };
  } else {
    throw new Error('Unrecognized talents with options: ' + PlayerOptions.toJsonString(playerOptions));
  }
  return copy;
}
