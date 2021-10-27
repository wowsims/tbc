import { Spec } from '/tbc/core/proto/common.js';
import { WarlockTalents as WarlockTalents } from '/tbc/core/proto/warlock.js';
import { WarlockSpecs } from '/tbc/core/proto_utils/utils.js';
import { Sim } from '/tbc/core/sim.js';

import { TalentsPicker } from './talents_picker.js';

export class WarlockTalentsPicker extends TalentsPicker<Spec.SpecWarlock> {
  constructor(parent: HTMLElement, sim: Sim<Spec.SpecWarlock>) {
    super(parent, sim, [
      {
        name: 'Affliction',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/302.jpg',
        talents: [
        ],
      },
      {
        name: 'Demonology',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/303.jpg',
        talents: [
        ],
      },
      {
        name: 'Destruction',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/301.jpg',
        talents: [
        ],
      },
    ]);
  }
}
