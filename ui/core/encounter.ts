import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Target, TargetConfig } from '/tbc/core/target.js';

import { Listener } from './typed_event.js';
import { Sim } from './sim.js';
import { TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
import { wait } from './utils.js';

export interface EncounterConfig {
	primaryTarget: TargetConfig,
}

// Manages all the settings for an Encounter.
export class Encounter {
  readonly durationChangeEmitter = new TypedEvent<void>();
  readonly numTargetsChangeEmitter = new TypedEvent<void>();

  // Emits when any of the above emitters emit.
  readonly changeEmitter = new TypedEvent<void>();

  private duration: number;
  private numTargets: number;

	readonly primaryTarget: Target;

	private readonly sim: Sim;

  constructor(config: EncounterConfig, sim: Sim) {
		this.sim = sim;

		this.duration = 300;
		this.numTargets = 1;
		this.primaryTarget = new Target(config.primaryTarget, sim);

    [
      this.durationChangeEmitter,
      this.numTargetsChangeEmitter,
      this.primaryTarget.changeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
  }
  
  getDuration(): number {
    return this.duration;
  }
  setDuration(newDuration: number) {
    if (newDuration == this.duration)
			return;

		this.duration = newDuration;
		this.durationChangeEmitter.emit();
  }
  
  getNumTargets(): number {
    return this.numTargets;
  }
  setNumTargets(newNumTargets: number) {
    if (newNumTargets == this.numTargets)
			return;

		this.numTargets = newNumTargets;
		this.numTargetsChangeEmitter.emit();
  }

	toProto(): EncounterProto {
		const numTargets = Math.max(1, this.numTargets);
		const targetProtos = [];
		for (let i = 0; i < numTargets; i++) {
			targetProtos.push(this.primaryTarget.toProto());
		}

		return EncounterProto.create({
			duration: this.duration,
			targets: targetProtos,
		});
	}

	fromProto(proto: EncounterProto) {
		this.setDuration(proto.duration);
		this.setNumTargets(proto.targets.length);

		if (proto.targets.length > 0) {
			this.primaryTarget.fromProto(proto.targets[0]);
		}
	}

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
			'duration': this.getDuration(),
			'numTargets': this.getNumTargets(),
			'primaryTarget': this.primaryTarget.toJson(),
    };
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		const parsedDuration = parseInt(obj['duration']);
		if (!isNaN(parsedDuration) && parsedDuration != 0) {
			this.setDuration(parsedDuration);
		}

		const parsedNumTargets = parseInt(obj['numTargets']);
		if (!isNaN(parsedNumTargets) && parsedNumTargets != 0) {
			this.setNumTargets(parsedNumTargets);
		}

		try {
			this.primaryTarget.fromJson(obj['primaryTarget']);
		} catch (e) {
			console.warn('Failed to parse debuffs: ' + e);
		}
  }
}
