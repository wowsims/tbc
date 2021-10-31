import { Player } from '/tbc/core/proto/api.js';
import { PlayerOptions } from '/tbc/core/proto/api.js';
import { SimOptions } from '/tbc/core/proto/api.js';
import { ComputeStatsRequest } from '/tbc/core/proto/api.js';
import { IndividualSimRequest } from '/tbc/core/proto/api.js';
import { withSpecProto } from './utils.js';
export function makeComputeStatsRequest(buffs, consumes, customStats, encounter, gear, race, rotation, talents, classOptions) {
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
export function makeIndividualSimRequest(buffs, consumes, customStats, encounter, gear, race, rotation, talents, classOptions, iterations, debug) {
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
        simOptions: SimOptions.create({
            iterations: iterations,
            debug: debug,
        }),
    });
}
