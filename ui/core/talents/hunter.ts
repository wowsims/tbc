import { Spec } from '/tbc/core/proto/common.js';
import { HunterTalents as HunterTalents } from '/tbc/core/proto/hunter.js';
import { HunterSpecs } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { TalentsPicker } from './talents_picker.js';

export class HunterTalentsPicker extends TalentsPicker<Spec.SpecHunter> {
  constructor(parent: HTMLElement, player: Player<Spec.SpecHunter>) {
    super(parent, player, [
      {
        name: 'Beast Mastery',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/361.jpg',
        talents: [
          {
            fieldName: 'improvedAspectOfTheHawk',
            location: {
              rowIdx: 0,
              colIdx: 1,
            },
            spellIds: [19552],
            maxPoints: 5,
          },
          {
            fieldName: 'enduranceTraining',
            location: {
              rowIdx: 0,
              colIdx: 2,
            },
            spellIds: [19583],
            maxPoints: 5,
          },
          {
            fieldName: 'focusFire',
            location: {
              rowIdx: 1,
              colIdx: 0,
            },
            spellIds: [35029],
            maxPoints: 2,
          },
          {
            //fieldName: 'improvedAspectOfTheMonkey',
            location: {
              rowIdx: 1,
              colIdx: 1,
            },
            spellIds: [19549],
            maxPoints: 3,
          },
          {
            //fieldName: 'thickHide',
            location: {
              rowIdx: 1,
              colIdx: 2,
            },
            spellIds: [19609, 19610, 19612],
            maxPoints: 3,
          },
          {
            //fieldName: 'improvedRevivePet',
            location: {
              rowIdx: 1,
              colIdx: 3,
            },
            spellIds: [24443, 19575],
            maxPoints: 2,
          },
          {
            //fieldName: 'pathfinding',
            location: {
              rowIdx: 2,
              colIdx: 0,
            },
            spellIds: [19559],
            maxPoints: 2,
          },
          {
            //fieldName: 'Bestial Swiftness',
            location: {
              rowIdx: 2,
              colIdx: 1,
            },
            spellIds: [19596],
            maxPoints: 1,
          },
          {
            fieldName: 'unleashedFury',
            location: {
              rowIdx: 2,
              colIdx: 2,
            },
            spellIds: [19616],
            maxPoints: 5,
          },
          {
            //fieldName: 'improvedMendPet',
            location: {
              rowIdx: 3,
              colIdx: 1,
            },
            spellIds: [19572],
            maxPoints: 2,
          },
          {
            fieldName: 'ferocity',
            location: {
              rowIdx: 3,
              colIdx: 2,
            },
            spellIds: [19598],
            maxPoints: 5,
          },
          {
            //fieldName: 'Spirit Bond',
            location: {
              rowIdx: 4,
              colIdx: 0,
            },
            spellIds: [19578, 20895],
            maxPoints: 2,
          },
          {
            //fieldName: 'Intimidation',
            location: {
              rowIdx: 4,
              colIdx: 1,
            },
            spellIds: [19577],
            maxPoints: 1,
          },
          {
            fieldName: 'Bestial Discipline',
            location: {
              rowIdx: 4,
              colIdx: 3,
            },
            spellIds: [19590],
            maxPoints: 2,
          },
          {
            //fieldName: 'animalHandler',
            location: {
              rowIdx: 5,
              colIdx: 0,
            },
            spellIds: [34453],
            maxPoints: 2,
          },
          {
            fieldName: 'frenzy',
            location: {
              rowIdx: 5,
              colIdx: 2,
            },
            prereqLocation: {
              rowIdx: 3,
              colIdx: 2,
            },
            spellIds: [19621],
            maxPoints: 5,
          },
          {
            fieldName: 'ferociousInspiration',
            location: {
              rowIdx: 6,
              colIdx: 0,
            },
            spellIds: [34455, 34459],
            maxPoints: 3,
          },
          {
            fieldName: 'Bestial Wrath',
            location: {
              rowIdx: 6,
              colIdx: 1,
            },
            prereqLocation: {
              rowIdx: 4,
              colIdx: 1,
            },
            spellIds: [19574],
            maxPoints: 1,
          },
          {
            //fieldName: 'catlikeReflexes',
            location: {
              rowIdx: 6,
              colIdx: 2,
            },
            spellIds: [34462, 34464],
            maxPoints: 3,
          },
          {
            fieldName: 'serpentsSwiftness',
            location: {
              rowIdx: 7,
              colIdx: 2,
            },
            spellIds: [34466],
            maxPoints: 5,
          },
          {
            fieldName: 'theBeastWithin',
            location: {
              rowIdx: 8,
              colIdx: 1,
            },
            prereqLocation: {
              rowIdx: 6,
              colIdx: 1,
            },
            spellIds: [34692],
            maxPoints: 1,
          },
        ],
      },
      {
        name: 'Marksmanship',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/363.jpg',
        talents: [
        ],
      },
      {
        name: 'Survival',
        backgroundUrl: 'https://wow.zamimg.com/images/wow/talents/backgrounds/tbc/362.jpg',
        talents: [
        ],
      },
    ]);
  }
}
