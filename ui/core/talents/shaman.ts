import { Sim } from '../sim';

import { TalentsPicker } from './talents_picker';

export class ShamanTalentsPicker extends TalentsPicker {
  constructor(parent: HTMLElement, sim: Sim) {
    super(parent, sim, [
      {
        name: 'Elemental',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/classic/261.jpg',
        talents: [
          {
            location: {
              rowIdx: 0,
              colIdx: 1,
            },
            spellIds: [16039, 16109],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 0,
              colIdx: 2,
            },
            spellIds: [16035, 16105],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 0,
            },
            spellIds: [16043, 16130],
            maxPoints: 2,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 1,
            },
            spellIds: [28996],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 2,
            },
            spellIds: [16038, 16160],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 0,
            },
            spellIds: [16164],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 1,
            },
            spellIds: [16040, 16113],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 2,
            },
            spellIds: [16041, 16117],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 3,
              colIdx: 0,
            },
            spellIds: [16086, 16544],
            maxPoints: 2,
          },
          {
            location: {
              rowIdx: 3,
              colIdx: 1,
            },
            spellIds: [29062, 29064],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 3,
              colIdx: 3,
            },
            spellIds: [30160, 29179],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 4,
              colIdx: 0,
            },
            spellIds: [28999],
            maxPoints: 2,
          },
          {
            location: {
              rowIdx: 4,
              colIdx: 1,
            },
            spellIds: [16089],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 4,
              colIdx: 3,
            },
            spellIds: [30664],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 5,
              colIdx: 0,
            },
            spellIds: [30672],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 5,
              colIdx: 2,
            },
            prereqLocation: {
              rowIdx: 2,
              colIdx: 2,
            },
            spellIds: [16578],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 6,
              colIdx: 1,
            },
            prereqLocation: {
              rowIdx: 4,
              colIdx: 1,
            },
            spellIds: [16166],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 6,
              colIdx: 2,
            },
            spellIds: [30669],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 7,
              colIdx: 1,
            },
            spellIds: [30675, 30678],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 8,
              colIdx: 1,
            },
            prereqLocation: {
              rowIdx: 7,
              colIdx: 1,
            },
            spellIds: [30706],
            maxPoints: 1,
          },
        ],
      },
      {
        name: 'Enchancement',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/classic/263.jpg',
        talents: [
          {
            location: {
              rowIdx: 0,
              colIdx: 1,
            },
            spellIds: [17485],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 0,
              colIdx: 2,
            },
            spellIds: [16253, 16298],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 0,
            },
            spellIds: [16258, 16293],
            maxPoints: 2,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 1,
            },
            spellIds: [16255, 16302],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 2,
            },
            spellIds: [16262, 16287],
            maxPoints: 2,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 3,
            },
            spellIds: [16261, 16290],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 0,
            },
            spellIds: [16259, 16295],
            maxPoints: 2,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 2,
            },
            spellIds: [43338],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 3,
            },
            spellIds: [16254, 16271],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 3,
              colIdx: 1,
            },
            prereqLocation: {
              rowIdx: 1,
              colIdx: 1,
            },
            spellIds: [16256, 16281],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 3,
              colIdx: 2,
            },
            spellIds: [16252, 16306],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 4,
              colIdx: 0,
            },
            spellIds: [29192],
            maxPoints: 2,
          },
          {
            location: {
              rowIdx: 4,
              colIdx: 1,
            },
            spellIds: [16268],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 4,
              colIdx: 2,
            },
            spellIds: [16266, 29079],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 5,
              colIdx: 0,
            },
            spellIds: [30812],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 5,
              colIdx: 3,
            },
            spellIds: [29082, 29084, 29086],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 6,
              colIdx: 0,
            },
            prereqLocation: {
              rowIdx: 6,
              colIdx: 1,
            },
            spellIds: [30816, 30818],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 6,
              colIdx: 1,
            },
            prereqLocation: {
              rowIdx: 4,
              colIdx: 1,
            },
            spellIds: [30798],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 6,
              colIdx: 2,
            },
            prereqLocation: {
              rowIdx: 4,
              colIdx: 2,
            },
            spellIds: [17364],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 7,
              colIdx: 1,
            },
            spellIds: [30802, 30808],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 8,
              colIdx: 1,
            },
            spellIds: [30823],
            maxPoints: 1,
          },
        ],
      },
      {
        name: 'Restoration',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/classic/262.jpg',
        talents: [
          {
            location: {
              rowIdx: 0,
              colIdx: 1,
            },
            spellIds: [16182, 16226],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 0,
              colIdx: 2,
            },
            spellIds: [16179, 16214],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 0,
            },
            spellIds: [16184, 16209],
            maxPoints: 2,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 1,
            },
            spellIds: [16176, 16235],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 1,
              colIdx: 2,
            },
            spellIds: [16173, 16222],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 0,
            },
            spellIds: [16180, 16196, 16198],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 1,
            },
            spellIds: [16181, 16230, 16232],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 2,
            },
            spellIds: [16189],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 2,
              colIdx: 3,
            },
            spellIds: [29187, 29189, 29191],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 3,
              colIdx: 1,
            },
            spellIds: [16187, 16205],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 3,
              colIdx: 2,
            },
            spellIds: [16194, 16218],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 4,
              colIdx: 0,
            },
            spellIds: [29206, 29205, 29202],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 4,
              colIdx: 2,
            },
            spellIds: [16188],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 4,
              colIdx: 3,
            },
            spellIds: [30864],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 5,
              colIdx: 2,
            },
            spellIds: [16178, 16210],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 6,
              colIdx: 1,
            },
            prereqLocation: {
              rowIdx: 3,
              colIdx: 1,
            },
            spellIds: [16190],
            maxPoints: 1,
          },
          {
            location: {
              rowIdx: 6,
              colIdx: 2,
            },
            spellIds: [30881, 30883],
            maxPoints: 5,
          },
          {
            location: {
              rowIdx: 7,
              colIdx: 1,
            },
            spellIds: [30867],
            maxPoints: 3,
          },
          {
            location: {
              rowIdx: 7,
              colIdx: 2,
            },
            spellIds: [30872],
            maxPoints: 2,
          },
          {
            location: {
              rowIdx: 8,
              colIdx: 1,
            },
            prereqLocation: {
              rowIdx: 7,
              colIdx: 1,
            },
            spellIds: [974],
            maxPoints: 1,
          },
        ],
      },
    ]);
  }
}
