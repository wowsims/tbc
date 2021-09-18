import { Stat } from './newapi';

const STATS_LEN = Object.keys(Stat).length;

/**
 * Represents values for all character stats (stam, agi, spell power, hit raiting, etc).
 *
 * This is an immutable type.
 */
export class Stats {
  private readonly stats: Array<number>;

  constructor(stats?: Array<number>) {
    this.stats = (stats?.slice(0, STATS_LEN) || []).concat(new Array(STATS_LEN - (stats?.length || 0)).fill(0));
  }

  equals(other: Stats): boolean {
    return this.stats.every((newStat, statIdx) => newStat == other.getStat(statIdx));
  }

  getStat(stat: Stat): number {
    return this.stats[stat];
  }

  withStat(stat: Stat, value: number): Stats {
    const newStats = this.stats.slice();
    newStats[stat] = value;
    return new Stats(newStats);
  }

  asArray(): Array<number> {
    return this.stats.slice();
  }

  toJson(): Object {
    return this.asArray();
  }

  static fromJson(obj: any): Stats {
    return new Stats(obj as Array<number>);
  }
}
