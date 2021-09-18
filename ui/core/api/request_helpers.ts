import { Gear } from './gear';
import { Buffs } from './newapi';
import { Consumes } from './newapi';
import { Class } from './newapi';
import { Encounter } from './newapi';
import { Race } from './newapi';
import { Player } from './newapi';
import { PlayerOptions } from './newapi';
import { Stats } from './stats';
import { ClassAgent } from './utils';
import { ClassTalents } from './utils';
import { ClassOptions } from './utils';

import { Druid, Druid_DruidAgent as DruidAgent, Druid_DruidTalents as DruidTalents, Druid_DruidOptions as DruidOptions} from './newapi';
import { Shaman, Shaman_ShamanAgent as ShamanAgent, Shaman_ShamanTalents as ShamanTalents, Shaman_ShamanOptions as ShamanOptions } from './newapi';

import { IndividualSimRequest, IndividualSimResult } from './newapi';

export function makeIndividualSimRequest<ClassType extends Class>(
    buffs: Buffs,
    consumes: Consumes,
    customStats: Stats,
    encounter: Encounter,
    gear: Gear,
    race: Race,
    agent: ClassAgent<ClassType>,
    talents: ClassTalents<ClassType>,
    classOptions: ClassOptions<ClassType>,
    iterations: number,
    debug: boolean): IndividualSimRequest {
  return IndividualSimRequest.create({
    player: Player.create({
      customStats: customStats.asArray(),
      equipment: gear.asSpec(),
      options: withClassProto(PlayerOptions.create({
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
function withClassProto<ClassType extends Class>(
    playerOptions: PlayerOptions,
    agent: ClassAgent<ClassType>,
    talents: ClassTalents<ClassType>,
    classOptions: ClassOptions<ClassType>): PlayerOptions {
  const copy = PlayerOptions.clone(playerOptions);
  if (DruidTalents.is(talents)) {
    copy.class = {
      oneofKind: 'druid',
      druid: Druid.create({
        agent: agent as DruidAgent,
        talents: talents,
        options: classOptions as DruidOptions,
      }),
    };
  } else if (ShamanTalents.is(talents)) {
    copy.class = {
      oneofKind: 'shaman',
      shaman: Shaman.create({
        agent: agent as ShamanAgent,
        talents: talents,
        options: classOptions as ShamanOptions,
      }),
    };
  } else {
    throw new Error('Unrecognized talents with options: ' + PlayerOptions.toJsonString(playerOptions));
  }
  return copy;
}
