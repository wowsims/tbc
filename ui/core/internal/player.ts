import { ItemSlot } from '../api/newapi';
import { Item } from '../api/newapi';
import { RaceBonusType } from '../api/newapi';
import { Spec } from '../api/newapi';
import { SpecToEligibleRaces } from '../api/utils';

export class Player {
  readonly spec: Spec;
  race: RaceBonusType;
  gear: Partial<Record<ItemSlot, Item>>;

  constructor(spec: Spec) {
    this.spec = spec;
    this.race = SpecToEligibleRaces[this.spec][0];
    this.gear = {};
  }
}
