import { Player } from '/tbc/core/proto/api.js';
import { PlayerOptions } from '/tbc/core/proto/api.js';
import { SimOptions } from '/tbc/core/proto/api.js';
import { ComputeStatsRequest } from '/tbc/core/proto/api.js';
import { IndividualSimRequest } from '/tbc/core/proto/api.js';
import { withSpecProto } from './utils.js';
export function makeComputeStatsRequest(raidBuffs, partyBuffs, individualBuffs, consumes, customStats, encounter, gear, race, rotation, talents, classOptions) {
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
export function makeIndividualSimRequest(raidBuffs, partyBuffs, individualBuffs, consumes, customStats, encounter, gear, race, rotation, talents, classOptions, iterations, debug) {
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
