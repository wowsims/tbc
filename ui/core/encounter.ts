import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';

import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';

// Manages all the settings for an Encounter.
export class Encounter {
	private readonly sim: Sim;

	private duration: number = 180;
	private durationVariation: number = 5;
	private numTargets: number = 1;
	private executeProportion: number = 0.2;
	readonly primaryTarget: Target;

	readonly durationChangeEmitter = new TypedEvent<void>();
	readonly numTargetsChangeEmitter = new TypedEvent<void>();
	readonly executeProportionChangeEmitter = new TypedEvent<void>();

	// Emits when any of the above emitters emit.
	readonly changeEmitter = new TypedEvent<void>();

	constructor(sim: Sim) {
		this.sim = sim;
		this.primaryTarget = new Target(sim);

		[
			this.durationChangeEmitter,
			this.numTargetsChangeEmitter,
			this.executeProportionChangeEmitter,
			this.primaryTarget.changeEmitter,
		].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
	}

	getDurationVariation(): number {
		return this.durationVariation;
	}
	setDurationVariation(eventID: EventID, newDuration: number) {
		if (newDuration == this.durationVariation)
			return;

		this.durationVariation = newDuration;
		this.durationChangeEmitter.emit(eventID);
	}

	getDuration(): number {
		return this.duration;
	}
	setDuration(eventID: EventID, newDuration: number) {
		if (newDuration == this.duration)
			return;

		this.duration = newDuration;
		this.durationChangeEmitter.emit(eventID);
	}

	getExecuteProportion(): number {
		return this.executeProportion;
	}
	setExecuteProportion(eventID: EventID, newExecuteProportion: number) {
		if (newExecuteProportion == this.executeProportion)
			return;

		this.executeProportion = newExecuteProportion;
		this.executeProportionChangeEmitter.emit(eventID);
	}

	getNumTargets(): number {
		return this.numTargets;
	}
	setNumTargets(eventID: EventID, newNumTargets: number) {
		if (newNumTargets == this.numTargets)
			return;

		this.numTargets = newNumTargets;
		this.numTargetsChangeEmitter.emit(eventID);
	}

	toProto(): EncounterProto {
		const numTargets = Math.max(1, this.numTargets);
		const targetProtos = [];
		for (let i = 0; i < numTargets; i++) {
			targetProtos.push(this.primaryTarget.toProto());
		}

		return EncounterProto.create({
			duration: this.duration,
			durationVariation: this.durationVariation,
			executeProportion: this.executeProportion,
			targets: targetProtos,
		});
	}

	fromProto(eventID: EventID, proto: EncounterProto) {
		TypedEvent.freezeAllAndDo(() => {
			this.setDuration(eventID, proto.duration);
			this.setDurationVariation(eventID, proto.durationVariation);
			this.setExecuteProportion(eventID, proto.executeProportion);
			this.setNumTargets(eventID, proto.targets.length);

			if (proto.targets.length > 0) {
				this.primaryTarget.fromProto(eventID, proto.targets[0]);
			}
		});
	}

	applyDefaults(eventID: EventID) {
		this.fromProto(eventID, EncounterProto.create({
			duration: 180,
			durationVariation: 5,
			executeProportion: 0.2,
			targets: [TargetProto.create({
				level: 73,
				stats: new Stats().withStat(Stat.StatArmor, 7684).asArray(),
				mobType: MobType.MobTypeDemon,
			})],
		}));
	}
}
