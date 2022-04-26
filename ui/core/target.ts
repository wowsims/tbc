import { Debuffs } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';

import * as Mechanics from '/tbc/core/constants/mechanics.js';

import { Listener } from './typed_event.js';
import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
import { wait } from './utils.js';

// Manages all the settings for a single Target.
export class Target {
	private readonly sim: Sim;

	private level: number = Mechanics.BOSS_LEVEL;
	private armor: number = 7684;
	private mobType: MobType = MobType.MobTypeDemon;
	private debuffs: Debuffs = Debuffs.create();

	readonly levelChangeEmitter = new TypedEvent<void>();
	readonly armorChangeEmitter = new TypedEvent<void>();
	readonly mobTypeChangeEmitter = new TypedEvent<void>();
	readonly debuffsChangeEmitter = new TypedEvent<void>();

	// Emits when any of the above emitters emit.
	readonly changeEmitter = new TypedEvent<void>();

	constructor(sim: Sim) {
		this.sim = sim;

		[
			this.levelChangeEmitter,
			this.armorChangeEmitter,
			this.mobTypeChangeEmitter,
			this.debuffsChangeEmitter,
		].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
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

	getArmor(): number {
		return this.armor;
	}

	setArmor(eventID: EventID, newArmor: number) {
		if (newArmor == this.armor)
			return;

		this.armor = newArmor;
		this.armorChangeEmitter.emit(eventID);
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
			armor: this.armor,
			mobType: this.mobType,
			debuffs: this.debuffs,
		});
	}

	fromProto(eventID: EventID, proto: TargetProto) {
		TypedEvent.freezeAllAndDo(() => {
			this.setLevel(eventID, proto.level);
			this.setArmor(eventID, proto.armor);
			this.setMobType(eventID, proto.mobType);
			this.setDebuffs(eventID, proto.debuffs || Debuffs.create());
		});
	}
}
