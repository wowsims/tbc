import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { makeComputeStatsRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { makeIndividualSimRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';

import { Player, PlayerConfig } from './player.js';
import { Sim, SimConfig } from './sim.js';
import { Target, TargetConfig } from './target.js';
import { TypedEvent } from './typed_event.js';

declare var tippy: any;

const CURRENT_SETTINGS_STORAGE_KEY = '__currentSettings__';
const SAVED_GEAR_STORAGE_KEY = '__savedGear__';
const SAVED_ENCOUNTER_STORAGE_KEY = '__savedEncounter__';
const SAVED_ROTATION_STORAGE_KEY = '__savedRotation__';
const SAVED_SETTINGS_STORAGE_KEY = '__savedSettings__';
const SAVED_TALENTS_STORAGE_KEY = '__savedTalents__';

export type ReleaseStatus = 'Alpha' | 'Beta' | 'Live';

export interface SimUIConfig<SpecType extends Spec> {
  releaseStatus: ReleaseStatus;
	knownIssues?: Array<string>;
	sim: SimConfig;
	player: PlayerConfig<SpecType>;
	target: TargetConfig;
}

// Core UI module.
export abstract class SimUI<SpecType extends Spec> {
  readonly parentElem: HTMLElement;
  readonly sim: Sim;
  readonly player: Player<SpecType>;
  readonly target: Target;
	readonly simUiConfig: SimUIConfig<SpecType>;

  // Emits when anything from sim, player, or target changes.
  readonly changeEmitter = new TypedEvent<void>();

  private readonly exclusivityMap: Record<ExclusivityTag, Array<ExclusiveEffect>>;

  constructor(parentElem: HTMLElement, config: SimUIConfig<SpecType>) {
    this.parentElem = parentElem;
    this.sim = new Sim(config.sim);
		this.player = new Player<SpecType>(config.player, this.sim);
    this.target = new Target(config.target, this.sim);
		this.simUiConfig = config;

    [
      this.sim.changeEmitter,
      this.player.changeEmitter,
      this.target.changeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));

    this.exclusivityMap = {
      'Battle Elixir': [],
      'Drums': [],
      'Food': [],
      'Guardian Elixir': [],
      'Potion': [],
      'Rune': [],
      'Weapon Imbue': [],
    };

		Array.from(document.getElementsByClassName('known-issues')).forEach(element => {
			if (this.simUiConfig.knownIssues?.length) {
				(element as HTMLElement).style.display = 'initial';
			} else {
				return;
			}

			
			tippy(element, {
				'content': `
				<ul class="known-issues-tooltip">
					${this.simUiConfig.knownIssues.map(issue => '<li>' + issue + '</li>').join('')}
				</ul>
				`,
				'allowHTML': true,
			});
		});
  }

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'sim': this.sim.toJson(),
      'player': this.player.toJson(),
      'target': this.target.toJson(),
    };
	}

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		if (obj['sim']) {
			this.sim.fromJson(obj['sim']);
		}

		if (obj['player']) {
			this.player.fromJson(obj['player']);
		}

		if (obj['target']) {
			this.target.fromJson(obj['target']);
		}
  }

  async init(): Promise<void> {
    await this.sim.init(this.player.spec);

    let loadedSettings = false;

    let hash = window.location.hash;
    if (hash.length > 1) {
      // Remove leading '#'
      hash = hash.substring(1);
      try {
        const jsonStr = atob(hash);
        this.fromJson(JSON.parse(jsonStr));
        loadedSettings = true;
      } catch (e) {
        console.warn('Failed to parse settings from window hash: ' + e);
      }
    }
		window.location.hash = '';

    const savedSettings = window.localStorage.getItem(this.getStorageKey(CURRENT_SETTINGS_STORAGE_KEY));
    if (!loadedSettings && savedSettings != null) {
      try {
        this.fromJson(JSON.parse(savedSettings));
        loadedSettings = true;
      } catch (e) {
        console.warn('Failed to parse saved settings: ' + e);
      }
    }

		if (!loadedSettings) {
			this.player.setGear(this.sim.lookupEquipmentSpec(this.player.defaultGear));
		}

    this.changeEmitter.on(() => {
      const jsonStr = JSON.stringify(this.toJson());
      window.localStorage.setItem(this.getStorageKey(CURRENT_SETTINGS_STORAGE_KEY), jsonStr);
    });

		
		Array.from(document.getElementsByClassName('share-link')).forEach(element => {
			tippy(element, {
				'content': 'Shareable link',
				'allowHTML': true,
			});

			element.addEventListener('click', event => {
				const linkUrl = new URL(window.location.href);
				const jsonStr = JSON.stringify(this.toJson());
				const encoded = btoa(jsonStr);
				linkUrl.hash = encoded;

				navigator.clipboard.writeText(linkUrl.toString());
				alert('Current settings copied to clipboard!');
			});
		});
  }

  registerExclusiveEffect(effect: ExclusiveEffect) {
    effect.tags.forEach(tag => {
      this.exclusivityMap[tag].push(effect);

      effect.changedEvent.on(() => {
        if (!effect.isActive())
          return;

        this.exclusivityMap[tag].forEach(otherEffect => {
          if (otherEffect == effect || !otherEffect.isActive())
            return;

          otherEffect.deactivate();
        });
      });
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

  makeCurrentIndividualSimRequest(iterations: number, debug: boolean): IndividualSimRequest {
		const encounter = this.sim.getEncounter();
		const numTargets = Math.max(1, this.sim.getNumTargets());
		for (let i = 0; i < numTargets; i++) {
			encounter.targets.push(this.target.toProto());
		}

    return makeIndividualSimRequest(
      this.sim.getBuffs(),
      this.player.getConsumes(),
      this.player.getCustomStats(),
      encounter,
      this.player.getGear(),
      this.player.getRace(),
      this.player.getRotation(),
      this.player.getTalents(),
      this.player.getSpecOptions(),
      iterations,
      debug);
  }
}

export type ExclusivityTag =
    'Battle Elixir'
    | 'Drums'
    | 'Food'
    | 'Guardian Elixir'
    | 'Potion'
    | 'Rune'
    | 'Weapon Imbue';

export interface ExclusiveEffect {
  tags: Array<ExclusivityTag>;
  changedEvent: TypedEvent<any>;
  isActive: () => boolean;
  deactivate: () => void;
}
