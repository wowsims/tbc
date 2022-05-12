import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EncounterType } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';

import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';

// Manages all the settings for an Encounter.
export class Encounter {
	readonly sim: Sim;

	private type: EncounterType = EncounterType.EncounterTypeSimple;
	private duration: number = 180;
	private durationVariation: number = 5;
	private numTargets: number = 1;
	private executeProportion: number = 0.2;
	private targets: Array<Target>;

	readonly targetsChangeEmitter = new TypedEvent<void>();
	readonly typeChangeEmitter = new TypedEvent<void>();
	readonly durationChangeEmitter = new TypedEvent<void>();
	readonly numTargetsChangeEmitter = new TypedEvent<void>();
	readonly executeProportionChangeEmitter = new TypedEvent<void>();

	// Emits when any of the above emitters emit.
	readonly changeEmitter = new TypedEvent<void>();

	constructor(sim: Sim) {
		this.sim = sim;
		this.targets = [Target.fromDefaults(TypedEvent.nextEventID(), sim)];

		[
			this.targetsChangeEmitter,
			this.typeChangeEmitter,
			this.durationChangeEmitter,
			this.numTargetsChangeEmitter,
			this.executeProportionChangeEmitter,
		].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
	}

	get primaryTarget(): Target {
		return this.targets[0];
	}

	getType(): EncounterType {
		return this.type;
	}
	setType(eventID: EventID, newType: EncounterType) {
		if (newType == this.type)
			return;

		this.type = newType;
		this.typeChangeEmitter.emit(eventID);
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

	getTargets(): Array<Target> {
		return this.targets.slice();
	}
	setTargets(eventID: EventID, newTargets: Array<Target>) {
		if (newTargets.length == 0) {
			newTargets = [Target.fromDefaults(eventID, this.sim)];
		}
		if (newTargets.length == this.targets.length && newTargets.every((target, i) => TargetProto.equals(target.toProto(), this.targets[i].toProto()))) {
			return;
		}

		TypedEvent.freezeAllAndDo(() => {
			if (newTargets.length != this.targets.length) {
				this.numTargets = newTargets.length;
				this.numTargetsChangeEmitter.emit(eventID);
			}
			this.targets = newTargets;
			this.targetsChangeEmitter.emit(eventID);
		});
	}

	toProto(): EncounterProto {
		const numTargets = Math.max(1, this.numTargets);
		let targetProtos = [];
		if (this.getType() == EncounterType.EncounterTypeSimple) {
			for (let i = 0; i < numTargets; i++) {
				targetProtos.push(this.primaryTarget.toProto());
			}
		} else {
			targetProtos = this.targets.map(target => target.toProto());
		}

		return EncounterProto.create({
			type: this.type,
			duration: this.duration,
			durationVariation: this.durationVariation,
			executeProportion: this.executeProportion,
			targets: targetProtos,
		});
	}

	fromProto(eventID: EventID, proto: EncounterProto) {
		TypedEvent.freezeAllAndDo(() => {
			this.setType(eventID, proto.type);
			this.setDuration(eventID, proto.duration);
			this.setDurationVariation(eventID, proto.durationVariation);
			this.setExecuteProportion(eventID, proto.executeProportion);
			this.setNumTargets(eventID, Math.max(1, proto.targets.length));

			if (proto.targets.length > 0) {
				this.primaryTarget.fromProto(eventID, proto.targets[0]);
				this.setTargets(eventID, proto.targets.map(targetProto => {
					const target = new Target(this.sim);
					target.fromProto(eventID, targetProto);
					return target;
				}));
			} else {
				this.setTargets(eventID, [ Target.fromDefaults(eventID, this.sim) ]);
			}
		});
	}

	applyDefaults(eventID: EventID) {
		this.fromProto(eventID, EncounterProto.create({
			duration: 180,
			durationVariation: 5,
			executeProportion: 0.2,
			targets: [ Target.defaultProto() ],
		}));
	}
}
