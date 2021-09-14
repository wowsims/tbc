import { Race } from '../api/newapi';
import { SpecToEligibleRaces } from '../api/utils';
import { RaceNames } from '../api/names';
import { Sim } from '../sim';

import { Component } from './component';

export class RacePicker extends Component {
  constructor(parent: HTMLElement, sim: Sim) {
    super(parent, 'race-picker-root');

    const raceSelector = document.createElement('select');
    raceSelector.classList.add('race-picker-selector');
    this.rootElem.appendChild(raceSelector);

    const races = SpecToEligibleRaces[sim.spec];
    races.forEach(race => {
      const option = document.createElement('option');
      option.value = String(race);
      option.textContent = RaceNames[race];
      raceSelector.appendChild(option);
    });

    raceSelector.value = String(sim.race);
    sim.raceChangeEmitter.on(newRace => {
      raceSelector.value = String(newRace);
    });

    raceSelector.addEventListener('change', event => {
      sim.race = parseInt(raceSelector.value) as Race;
    });
  }
}
