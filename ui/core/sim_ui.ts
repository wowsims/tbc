import { SimOptions } from '/tbc/core/proto/api.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';

import { Raid } from './raid.js';
import { Sim } from './sim.js';
import { Encounter } from './encounter.js';
import { Target } from './target.js';
import { TypedEvent } from './typed_event.js';

declare var tippy: any;
declare var pako: any;

const CURRENT_SETTINGS_STORAGE_KEY = '__currentSettings__';
const SAVED_GEAR_STORAGE_KEY = '__savedGear__';
const SAVED_ENCOUNTER_STORAGE_KEY = '__savedEncounter__';
const SAVED_ROTATION_STORAGE_KEY = '__savedRotation__';
const SAVED_SETTINGS_STORAGE_KEY = '__savedSettings__';
const SAVED_TALENTS_STORAGE_KEY = '__savedTalents__';

// Core UI module.
export abstract class SimUI {
  readonly parentElem: HTMLElement;
  readonly sim: Sim;
  readonly raid: Raid;
  readonly encounter: Encounter;

  // Emits when anything from sim, player, or target changes.
  readonly changeEmitter = new TypedEvent<void>();

  private readonly exclusivityMap: Record<ExclusivityTag, Array<ExclusiveEffect>>;

  constructor(parentElem: HTMLElement) {
    this.parentElem = parentElem;
    this.sim = new Sim();

		this.raid = new Raid(this.sim);
    this.encounter = new Encounter(this.sim);

    [
      this.sim.changeEmitter,
      this.raid.changeEmitter,
      this.encounter.changeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
  }

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'raid': this.raid.toJson(),
      'encounter': this.encounter.toJson(),
    };
	}

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		// For legacy format. Do not remove this until 2022/01/05 (1 month).
		if (obj['sim']) {
			if (!obj['raid']) {
				obj['raid'] = {
					'parties': [
						{
							'players': [
								{
									'spec': this.player.spec,
									'player': obj['player'],
								},
							],
							'buffs': obj['sim']['partyBuffs'],
						},
					],
					'buffs': obj['sim']['raidBuffs'],
				};
				obj['raid']['parties'][0]['players'][0]['player']['buffs'] = obj['sim']['individualBuffs'];
			}
		}

		if (obj['raid']) {
			this.raid.fromJson(obj['raid']);
		}

		if (obj['encounter']) {
			this.encounter.fromJson(obj['encounter']);
		}
  }

  async init(): Promise<void> {
    await this.sim.init();

    this.changeEmitter.on(() => {
      const jsonStr = JSON.stringify(this.toJson());
      window.localStorage.setItem(this.getStorageKey(CURRENT_SETTINGS_STORAGE_KEY), jsonStr);
    });
  }

	getSavedGearStorageKey(): string {
		return this.getStorageKey(SAVED_GEAR_STORAGE_KEY);
	}

	getSavedEncounterStorageKey(): string {
		return this.getStorageKey(SAVED_ENCOUNTER_STORAGE_KEY);
	}

	getSavedRotationStorageKey(): string {
		return this.getStorageKey(SAVED_ROTATION_STORAGE_KEY);
	}

	getSavedSettingsStorageKey(): string {
		return this.getStorageKey(SAVED_SETTINGS_STORAGE_KEY);
	}

	getSavedTalentsStorageKey(): string {
		return this.getStorageKey(SAVED_TALENTS_STORAGE_KEY);
	}

	// Returns the actual key to use for local storage, based on the given key part and the site context.
	private getStorageKey(keyPart: string): string {
		// Local storage is shared by all sites under the same domain, so we need to use
		// different keys for each spec site.
		return specToLocalStorageKey[this.player.spec] + keyPart;
	}

  makeRaidSimRequest(iterations: number, debug: boolean): RaidSimRequest {
		return RaidSimRequest.create({
			raid: this.raid.toProto(),
			encounter: this.encounter.toProto(),
			simOptions: SimOptions.create({
				iterations: iterations,
				debug: debug,
			}),
		});
  }
}
