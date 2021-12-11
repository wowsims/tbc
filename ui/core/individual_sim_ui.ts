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

export interface IndividualSimUIConfig<SpecType extends Spec> {
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
export abstract class IndividualSimUI<SpecType extends Spec> extends SimUI {
  readonly party: Party;
  readonly player: Player<SpecType>;
	readonly simUiConfig: SimUIConfig<SpecType>;

  private readonly exclusivityMap: Record<ExclusivityTag, Array<ExclusiveEffect>>;

  constructor(parentElem: HTMLElement, sim: Sim, config: SimUIConfig<SpecType>) {
		super(parentElem, sim);
		this.rootElem.classList.add('individual-sim-ui');

		this.party = this.raid.getParty(0);
		this.player = new Player<SpecType>(config.spec, this.sim);
		this.raid.setPlayer(0, this.player);

		this.simUiConfig = config;

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
  }

  async init(): Promise<void> {
    await super.init();

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

	private applyDefaults() {
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
