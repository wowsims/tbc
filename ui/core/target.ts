import { Debuffs } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';

import { Listener } from './typed_event.js';
import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
import { wait } from './utils.js';

// Manages all the settings for a single Target.
export class Target {
	private readonly sim: Sim;

	private armor: number = 7684;
	private mobType: MobType = MobType.MobTypeDemon;
	private debuffs: Debuffs = Debuffs.create();

	readonly armorChangeEmitter = new TypedEvent<void>();
	readonly mobTypeChangeEmitter = new TypedEvent<void>();
	readonly debuffsChangeEmitter = new TypedEvent<void>();

	// Emits when any of the above emitters emit.
	readonly changeEmitter = new TypedEvent<void>();

	constructor(sim: Sim) {
		this.sim = sim;

		[
			this.armorChangeEmitter,
			this.mobTypeChangeEmitter,
			this.debuffsChangeEmitter,
		].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
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
			armor: this.armor,
			mobType: this.mobType,
			debuffs: this.debuffs,
		});
	}

	fromProto(eventID: EventID, proto: TargetProto) {
		TypedEvent.freezeAllAndDo(() => {
			this.setArmor(eventID, proto.armor);
			this.setMobType(eventID, proto.mobType);
			this.setDebuffs(eventID, proto.debuffs || Debuffs.create());
		});
	}
}
