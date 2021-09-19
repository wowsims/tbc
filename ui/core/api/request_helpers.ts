import { Gear } from './gear';
import { Buffs } from './newapi';
import { Consumes } from './newapi';
import { Class } from './newapi';
import { Encounter } from './newapi';
import { Player } from './newapi';
import { PlayerOptions } from './newapi';
import { Race } from './newapi';
import { Spec } from './newapi';
import { Stats } from './stats';
import { SpecAgent } from './utils';
import { SpecTalents } from './utils';
import { SpecOptions } from './utils';

import { BalanceDruid, BalanceDruid_BalanceDruidAgent as BalanceDruidAgent, DruidTalents, BalanceDruid_BalanceDruidOptions as BalanceDruidOptions} from './newapi';
import { ElementalShaman, ElementalShaman_ElementalShamanAgent as ElementalShamanAgent, ShamanTalents, ElementalShaman_ElementalShamanOptions as ElementalShamanOptions } from './newapi';

import { IndividualSimRequest, IndividualSimResult } from './newapi';

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
