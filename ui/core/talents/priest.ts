import { Spec } from '/tbc/core/proto/common.js';
import { PriestTalents as PriestTalents } from '/tbc/core/proto/priest.js';
import { PriestSpecs } from '/tbc/core/proto_utils/utils.js';
import { Sim } from '/tbc/core/sim.js';

import { TalentsPicker } from './talents_picker.js';

// Talents are the same for all Priest specs, so its ok to just use ShadowPriest here
export class PriestTalentsPicker extends TalentsPicker<Spec.SpecShadowPriest> {
  constructor(parent: HTMLElement, sim: Sim<Spec.SpecShadowPriest>) {
    super(parent, sim, [
      {
        name: 'Discipline',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/201.jpg',
        talents: [
          {
            //fieldName: 'martyrdom',
            location: {
              rowIdx: 3,
              colIdx: 1,
            },
            spellIds: [14531, 14774],
            maxPoints: 2,
          },
          {
            fieldName: 'powerInfusion',
            location: {
              rowIdx: 6,
              colIdx: 1,
            },
            prereqLocation: {
              rowIdx: 4,
              colIdx: 1,
            },
            spellIds: [10060],
            maxPoints: 1,
          },
          {
            fieldName: 'mentalAgility',
            location: {
              rowIdx: 3,
              colIdx: 1,
            },
            spellIds: [14520, 14780],
            maxPoints: 5,
          },
          {
            fieldName: 'unbreakableWill',
            location: {
              rowIdx: 0,
              colIdx: 1,
            },
            spellIds: [14522, 14788],
            maxPoints: 5,
          },
        ],
      },
      {
        name: 'Holy',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/202.jpg',
        talents: [
        ],
      },
      {
        name: 'Shadow',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/203.jpg',
        talents: [
        ],
      },
    ]);
  }
}
