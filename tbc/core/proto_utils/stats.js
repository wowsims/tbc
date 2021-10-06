import { Stat } from '/tbc/core/proto/common.js';
import { getEnumValues } from '/tbc/core/utils.js';
const STATS_LEN = getEnumValues(Stat).length;
/**
 * Represents values for all character stats (stam, agi, spell power, hit raiting, etc).
 *
 * This is an immutable type.
 */
export class Stats {
    constructor(stats) {
        this.stats = stats?.slice(0, STATS_LEN) || [];
        if (this.stats.length < STATS_LEN) {
            this.stats = this.stats.concat(new Array(STATS_LEN - (stats?.length || 0)).fill(0));
        }
        for (let i = 0; i < STATS_LEN; i++) {
            if (this.stats[i] == null)
                this.stats[i] = 0;
        }
    }
    equals(other) {
        return this.stats.every((newStat, statIdx) => newStat == other.getStat(statIdx));
    }
    getStat(stat) {
        return this.stats[stat];
    }
    withStat(stat, value) {
        const newStats = this.stats.slice();
        newStats[stat] = value;
        return new Stats(newStats);
    }
    computeEP(epWeights) {
        let total = 0;
        this.stats.forEach((stat, idx) => {
            total += stat * epWeights.stats[idx];
        });
        return total;
    }
    asArray() {
        return this.stats.slice();
    }
    toJson() {
        return this.asArray();
    }
    static fromJson(obj) {
        return new Stats(obj);
    }
    static fromMap(statsMap) {
        const statsArr = new Array(STATS_LEN).fill(0);
        Object.entries(statsMap).forEach(entry => {
            const [statStr, value] = entry;
            statsArr[Number(statStr)] = value;
        });
        return new Stats(statsArr);
    }
}
