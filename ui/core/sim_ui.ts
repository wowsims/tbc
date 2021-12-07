import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';
import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { StatWeightsRequest } from '/tbc/core/proto/api.js';
import { SimOptions } from '/tbc/core/proto/api.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { SpecOptions } from '/tbc/core/proto_utils/utils.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';

import { Party } from './party.js';
import { Player } from './player.js';
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

export type ReleaseStatus = 'Alpha' | 'Beta' | 'Live';

export interface SimUIConfig<SpecType extends Spec> {
	spec: Spec,
  releaseStatus: ReleaseStatus;
	knownIssues?: Array<string>;

  defaults: {
		gear: EquipmentSpec,
		epWeights: Stats,
    consumes: Consumes,
    rotation: SpecRotation<SpecType>,
    talents: string,
    specOptions: SpecOptions<SpecType>,

    raidBuffs: RaidBuffs,
    partyBuffs: PartyBuffs,
    individualBuffs: IndividualBuffs,

    debuffs: Debuffs,
  },
}

// Core UI module.
export abstract class SimUI<SpecType extends Spec> {
  readonly parentElem: HTMLElement;
  readonly sim: Sim;
  readonly raid: Raid;
  readonly party: Party;
  readonly player: Player<SpecType>;
  readonly encounter: Encounter;
	readonly simUiConfig: SimUIConfig<SpecType>;

  // Emits when anything from sim, player, or target changes.
  readonly changeEmitter = new TypedEvent<void>();

  private readonly exclusivityMap: Record<ExclusivityTag, Array<ExclusiveEffect>>;

  constructor(parentElem: HTMLElement, config: SimUIConfig<SpecType>) {
    this.parentElem = parentElem;
    this.sim = new Sim();

		this.raid = new Raid(this.sim);
		this.party = this.raid.getParty(0);
		this.player = new Player<SpecType>(config.spec, this.sim);
		this.raid.setPlayer(0, this.player);

    this.encounter = new Encounter(this.sim);

		this.simUiConfig = config;

    [
      this.sim.changeEmitter,
      this.raid.changeEmitter,
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
			this.applyDefaults();
		}
		this.player.setEpWeights(this.simUiConfig.defaults.epWeights);

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
        if (navigator.clipboard == undefined) {
          alert(linkUrl.toString());
        } else {
          navigator.clipboard.writeText(linkUrl.toString());
        }
				alert('Current settings copied to clipboard!');
			});
		});
  }

	applyDefaults() {
		this.player.setGear(this.sim.lookupEquipmentSpec(this.simUiConfig.defaults.gear));
		this.player.setConsumes(this.simUiConfig.defaults.consumes);
		this.player.setRotation(this.simUiConfig.defaults.rotation);
		this.player.setTalentsString(this.simUiConfig.defaults.talents);
		this.player.setSpecOptions(this.simUiConfig.defaults.specOptions);
		this.encounter.primaryTarget.setDebuffs(this.simUiConfig.defaults.debuffs);
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

  makeCurrentIndividualSimRequest(iterations: number, debug: boolean): IndividualSimRequest {
		return IndividualSimRequest.create({
			player: this.player.toProto(),
			raidBuffs: this.raid.getBuffs(),
			partyBuffs: this.party.getBuffs(),
			encounter: this.encounter.toProto(),
			simOptions: SimOptions.create({
				iterations: iterations,
				debug: debug,
			}),
		});
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
