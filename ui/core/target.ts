import { Debuffs } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';

import { Listener } from './typed_event.js';
import { Sim } from './sim.js';
import { TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
import { wait } from './utils.js';

export interface TargetConfig {
  defaults: {
		armor: number,
		debuffs: Debuffs,
  },
}

// Manages all the settings for a single Target.
export class Target {
  readonly armorChangeEmitter = new TypedEvent<void>();
  readonly debuffsChangeEmitter = new TypedEvent<void>();

  // Emits when any of the above emitters emit.
  readonly changeEmitter = new TypedEvent<void>();

  // Current values
	private armor: number;
  private debuffs: Debuffs;

	private readonly sim: Sim;

  constructor(config: TargetConfig, sim: Sim) {
		this.sim = sim;

    this.armor = config.defaults.armor;
    this.debuffs = config.defaults.debuffs;

    [
      this.armorChangeEmitter,
      this.debuffsChangeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
  }

  getArmor(): number {
    return this.armor;
  }

  setArmor(newArmor: number) {
    if (newArmor == this.armor)
      return;

		this.armor = newArmor;
    this.armorChangeEmitter.emit();
  }

  getDebuffs(): Debuffs {
    // Make a defensive copy
    return Debuffs.clone(this.debuffs);
  }

  setDebuffs(newDebuffs: Debuffs) {
    if (Debuffs.equals(this.debuffs, newDebuffs))
      return;

    // Make a defensive copy
    this.debuffs = Debuffs.clone(newDebuffs);
    this.debuffsChangeEmitter.emit();
  }

	toProto(): TargetProto {
		return TargetProto.create({
			armor: this.armor,
			debuffs: this.debuffs,
		});
	}

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'armor': this.armor,
      'debuffs': Debuffs.toJson(this.debuffs),
    };
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		const parsedArmor = parseInt(obj['armor']);
		if (!isNaN(parsedArmor) && parsedArmor != 0) {
			this.armor = parsedArmor;
		}

		try {
			this.setDebuffs(Debuffs.fromJson(obj['debuffs']));
		} catch (e) {
			console.warn('Failed to parse debuffs: ' + e);
		}
  }
}
