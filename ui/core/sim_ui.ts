import { Sim, SimConfig } from './sim.js';
import { TypedEvent } from './typed_event.js';
import { Class } from './proto/common.js';
import { Spec } from './proto/common.js';

declare var tippy: any;

const CURRENT_SETTINGS_KEY = '__currentSettings__';

export type ReleaseStatus = 'Alpha' | 'Beta' | 'Live';

export interface SimUIConfig<SpecType extends Spec> extends SimConfig<SpecType> {
  releaseStatus: ReleaseStatus;
	knownIssues?: Array<string>;
}

// Core UI module.
export abstract class SimUI<SpecType extends Spec> {
  readonly parentElem: HTMLElement;
  readonly sim: Sim<SpecType>;
	readonly simUiConfig: SimUIConfig<SpecType>;

  private readonly exclusivityMap: Record<ExclusivityTag, Array<ExclusiveEffect>>;

  constructor(parentElem: HTMLElement, config: SimUIConfig<SpecType>) {
    this.parentElem = parentElem;
    this.sim = new Sim<SpecType>(config);
		this.simUiConfig = config;

    this.exclusivityMap = {
      'Battle Elixir': [],
      'Drums': [],
      'Food': [],
      'Guardian Elixir': [],
      'Potion': [],
      'Rune': [],
      'Weapon Imbue': [],
    };
  }

  async init(): Promise<void> {
    await this.sim.init();

    let loadedSettings = false;

    let hash = window.location.hash;
    if (hash.length > 1) {
      // Remove leading '#'
      hash = hash.substring(1);
      try {
        const simJsonStr = atob(hash);
        this.sim.fromJson(JSON.parse(simJsonStr));
        loadedSettings = true;
      } catch (e) {
        console.warn('Failed to parse settings from window hash: ' + e);
      }
    }
		window.location.hash = '';

    const savedSettings = window.localStorage.getItem(CURRENT_SETTINGS_KEY);
    if (!loadedSettings && savedSettings != null) {
      try {
        this.sim.fromJson(JSON.parse(savedSettings));
      } catch (e) {
        console.warn('Failed to parse saved settings: ' + e);
      }
    }

    this.sim.changeEmitter.on(() => {
      const simJsonStr = JSON.stringify(this.sim.toJson());
      window.localStorage.setItem(CURRENT_SETTINGS_KEY, simJsonStr);
    });

		
		Array.from(document.getElementsByClassName('share-link')).forEach(element => {
			tippy(element, {
				'content': 'Shareable link',
				'allowHTML': true,
			});

			element.addEventListener('click', event => {
				const linkUrl = new URL(window.location.href);
				const simJsonStr = JSON.stringify(this.sim.toJson());
				const simEncoded = btoa(simJsonStr);
				linkUrl.hash = simEncoded;

				navigator.clipboard.writeText(linkUrl.toString());
				alert('Current settings copied to clipboard!');
			});
		});

		Array.from(document.getElementsByClassName('known-issues')).forEach(element => {
			if (!this.simUiConfig.knownIssues?.length) {
				(element as HTMLElement).style.display = 'none';
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
