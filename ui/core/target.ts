import { Debuffs } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';

import * as Mechanics from '/tbc/core/constants/mechanics.js';

import { Listener } from './typed_event.js';
import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
import { wait } from './utils.js';

// Manages all the settings for a single Target.
export class Target {
	readonly sim: Sim;

	private level: number = Mechanics.BOSS_LEVEL;
	private mobType: MobType = MobType.MobTypeDemon;
	private stats: Stats = new Stats();
	private debuffs: Debuffs = Debuffs.create();

	readonly levelChangeEmitter = new TypedEvent<void>();
	readonly statsChangeEmitter = new TypedEvent<void>();
	readonly mobTypeChangeEmitter = new TypedEvent<void>();
	readonly debuffsChangeEmitter = new TypedEvent<void>();

	// Emits when any of the above emitters emit.
	readonly changeEmitter = new TypedEvent<void>();

	constructor(sim: Sim) {
		this.sim = sim;

		[
			this.levelChangeEmitter,
			this.statsChangeEmitter,
			this.mobTypeChangeEmitter,
			this.debuffsChangeEmitter,
		].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));

		this.changeEmitter.on(eventID => this.sim.encounter?.changeEmitter.emit(eventID));
	}

	getLevel(): number {
		return this.level;
	}

	setLevel(eventID: EventID, newLevel: number) {
		if (newLevel == this.level)
			return;

		this.level = newLevel;
		this.levelChangeEmitter.emit(eventID);
	}

	getStats(): Stats {
		return this.stats;
	}

	setStats(eventID: EventID, newStats: Stats) {
		if (newStats.equals(this.stats))
			return;

		this.stats = newStats;
		this.statsChangeEmitter.emit(eventID);
	}

	getMobType(): MobType {
		return this.mobType;
	}

	setMobType(eventID: EventID, newMobType: MobType) {
		if (newMobType == this.mobType)
			return;

		this.mobType = newMobType;
		this.mobTypeChangeEmitter.emit(eventID);
	}

	getDebuffs(): Debuffs {
		// Make a defensive copy
		return Debuffs.clone(this.debuffs);
	}

	setDebuffs(eventID: EventID, newDebuffs: Debuffs) {
		if (Debuffs.equals(this.debuffs, newDebuffs))
			return;

		// Make a defensive copy
		this.debuffs = Debuffs.clone(newDebuffs);
		this.debuffsChangeEmitter.emit(eventID);
	}

	toProto(): TargetProto {
		return TargetProto.create({
			level: this.level,
			stats: this.stats.asArray(),
			mobType: this.mobType,
			debuffs: this.debuffs,
		});
	}

	fromProto(eventID: EventID, proto: TargetProto) {
		TypedEvent.freezeAllAndDo(() => {
			let stats = new Stats(proto.stats);
			if (proto.armor) {
				stats = stats.withStat(Stat.StatArmor, proto.armor);
			}

			this.setLevel(eventID, proto.level);
			this.setStats(eventID, stats);
			this.setMobType(eventID, proto.mobType);
			this.setDebuffs(eventID, proto.debuffs || Debuffs.create());
		});
	}

	static defaultProto(): TargetProto {
		return TargetProto.create({
			level: Mechanics.BOSS_LEVEL,
			mobType: MobType.MobTypeDemon,
			stats: Stats.fromMap({
				[Stat.StatArmor]: 7683,
				[Stat.StatBlockValue]: 54,
				[Stat.StatAttackPower]: 320,
			}).asArray(),
		});
	}

	static fromDefaults(eventID: EventID, sim: Sim): Target {
		const target = new Target(sim);
		target.fromProto(eventID, Target.defaultProto());
		return target;
	}
}
