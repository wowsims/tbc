import { Spec } from '/tbc/core/proto/common.js';
import { DruidTalents as DruidTalents } from '/tbc/core/proto/druid.js';
import { DruidSpecs } from '/tbc/core/proto_utils/utils.js';
import { Sim } from '/tbc/core/sim.js';

import { TalentsPicker } from './talents_picker.js';

// Talents are the same for all Druid specs, so its ok to just use BalanceDruid here
export class DruidTalentsPicker extends TalentsPicker<Spec.SpecBalanceDruid> {
  constructor(parent: HTMLElement, sim: Sim<Spec.SpecBalanceDruid>) {
    super(parent, sim, [
      {
        name: 'Balance',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/283.jpg',
        talents: [
        ],
      },
      {
        name: 'Feral Combat',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/281.jpg',
        talents: [
        ],
      },
      {
        name: 'Restoration',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/282.jpg',
        talents: [
        ],
      },
    ]);
  }
}
