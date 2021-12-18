import { Spec } from '/tbc/core/proto/common.js';
import { Sim } from '/tbc/core/sim.js';
import { Player } from '/tbc/core/player.js';

import { EnhancementShamanSimUI } from './sim.js';

const sim = new Sim();
const player = new Player<Spec.SpecEnhancementShaman>(Spec.SpecEnhancementShaman, sim);
sim.raid.setPlayer(0, player);

const simUI = new EnhancementShamanSimUI(document.body, player);
