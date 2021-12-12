import { Spec } from '/tbc/core/proto/common.js';
import { Sim } from '/tbc/core/sim.js';
import { Player } from '/tbc/core/player.js';

import { ElementalShamanSimUI } from './sim.js';

const sim = new Sim();
const player = new Player<Spec.SpecElementalShaman>(Spec.SpecElementalShaman, sim);
sim.raid.setPlayer(0, player);

const simUI = new ElementalShamanSimUI(document.body, player);
