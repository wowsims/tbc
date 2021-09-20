import { Sim, SimConfig } from './sim';
import { TypedEvent } from './typed_event';
import { Class } from './api/common';
import { Spec } from './api/common';

const CURRENT_SETTINGS_KEY = '__currentSettings__';

export interface SimUIConfig<SpecType extends Spec> extends SimConfig<SpecType> {
}

// Core UI module.
export abstract class SimUI<SpecType extends Spec> {
  readonly parentElem: HTMLElement;
  readonly sim: Sim<SpecType>;

  private readonly exclusivityMap: Record<ExclusivityTag, Array<ExclusiveEffect>>;

  constructor(parentElem: HTMLElement, config: SimUIConfig<SpecType>) {
    this.parentElem = parentElem;
    this.sim = new Sim<SpecType>(config);

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

      const b64Str = btoa(simJsonStr);
      window.location.hash = b64Str;
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
