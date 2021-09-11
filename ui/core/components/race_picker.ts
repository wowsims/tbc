import { RaceBonusType } from '../api/newapi';
import { SpecToEligibleRaces } from '../api/utils';
import { RaceNames } from '../api/utils';
import { Sim } from '../sim.js';

import { Component } from './component.js';

export class RacePicker extends Component {
  readonly rootElem: HTMLDivElement;

  constructor(sim: Sim) {
    super();

    this.rootElem = document.createElement('div');
    this.rootElem.classList.add('race-picker-root');

    const label = document.createElement('span');
    label.classList.add('race-picker-label');
    label.textContent = 'Race:';
    this.rootElem.appendChild(label);

    const raceSelector = document.createElement('select');
    raceSelector.classList.add('race-picker-selector');
    this.rootElem.appendChild(raceSelector);

    const races = SpecToEligibleRaces[sim.spec];
    races.forEach(race => {
      const option = document.createElement('option');
      option.value = String(race as number);
      option.textContent = RaceNames[race];
      raceSelector.appendChild(option);
    });

    raceSelector.value = String(sim.race as number);
    sim.raceChangeEmitter.on(newRace => {
      raceSelector.value = String(newRace as number);
    });

    raceSelector.addEventListener('change', event => {
      const newRace: RaceBonusType = parseInt(raceSelector.value) as RaceBonusType;
      sim.race = newRace;
    });
  }

  getRootElement() {
    return this.rootElem;
  }
}
