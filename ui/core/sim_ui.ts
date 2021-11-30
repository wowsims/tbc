import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';
import { Spec } from '/tbc/core/proto/common.js';
import { makeIndividualSimRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';
import { Player, PlayerConfig } from './player.js';
import { Sim, SimConfig } from './sim.js';
import { Encounter, EncounterConfig } from './encounter.js';
import { Target, TargetConfig } from './target.js';
import { TypedEvent } from './typed_event.js';

declare var tippy: any;
declare var pako: any;

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
	encounter: EncounterConfig;
}

// Core UI module.
export abstract class SimUI<SpecType extends Spec> {
  readonly parentElem: HTMLElement;
  readonly sim: Sim;
  readonly player: Player<SpecType>;
  readonly encounter: Encounter;
	readonly simUiConfig: SimUIConfig<SpecType>;

  // Emits when anything from sim, player, or target changes.
  readonly changeEmitter = new TypedEvent<void>();

  private readonly exclusivityMap: Record<ExclusivityTag, Array<ExclusiveEffect>>;

  constructor(parentElem: HTMLElement, config: SimUIConfig<SpecType>) {
    this.parentElem = parentElem;
    this.sim = new Sim(config.sim);
		this.player = new Player<SpecType>(config.player, this.sim);
    this.encounter = new Encounter(config.encounter, this.sim);
		this.simUiConfig = config;

    [
      this.sim.changeEmitter,
      this.player.changeEmitter,
      this.encounter.changeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));

    this.exclusivityMap = {
      'Battle Elixir': [],
      'Drums': [],
      'Food': [],
      'Alchohol': [],
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
      'encounter': this.encounter.toJson(),
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

		if (obj['encounter']) {
			this.encounter.fromJson(obj['encounter']);
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
        let jsonData;
        if (new URLSearchParams(window.location.search).has('uncompressed')) {
          const jsonStr = atob(hash);
          jsonData = JSON.parse(jsonStr);
        } else {
          const binary = atob(hash);
          const bytes = new Uint8Array(binary.length);
          for (let i = 0; i < bytes.length; i++) {
              bytes[i] = binary.charCodeAt(i);
          }
          const jsonStr = pako.inflate(bytes, { to: 'string' });  
          jsonData = JSON.parse(jsonStr);
        }
        this.fromJson(jsonData);
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
				const jsonStr = JSON.stringify(this.toJson());
        const val = pako.deflate(jsonStr, { to: 'string' });
        const encoded = btoa(String.fromCharCode(...val));
				
        const linkUrl = new URL(window.location.href);
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
    return makeIndividualSimRequest(
      this.sim.getRaidBuffs(),
      this.sim.getPartyBuffs(),
      this.sim.getIndividualBuffs(),
      this.player.getConsumes(),
      this.player.getCustomStats(),
      this.encounter.toProto(),
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
    | 'Alchohol'
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
