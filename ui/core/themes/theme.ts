import { Sim, SimConfig } from '../sim';
import { TypedEvent } from '../typed_event';
import { Class } from '../api/newapi';
import { Spec } from '../api/newapi';

export interface ThemeConfig<SpecType extends Spec> extends SimConfig<SpecType> {
}

export abstract class Theme<SpecType extends Spec> {
  readonly parentElem: HTMLElement;
  readonly sim: Sim<SpecType>;

  private readonly exclusivityMap: Record<ExclusivityTag, Array<ExclusiveEffect>>;

  constructor(parentElem: HTMLElement, config: ThemeConfig<SpecType>) {
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
