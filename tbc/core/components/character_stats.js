import { Stat } from '/tbc/core/proto/common.js';
import { statNames } from '/tbc/core/api/names.js';
import { Stats } from '/tbc/core/api/stats.js';
import { Component } from './component.js';
const spellPowerTypeStats = [
    Stat.StatArcaneSpellPower,
    Stat.StatFireSpellPower,
    Stat.StatFrostSpellPower,
    Stat.StatHolySpellPower,
    Stat.StatNatureSpellPower,
    Stat.StatShadowSpellPower,
];
export class CharacterStats extends Component {
    constructor(parent, stats, sim) {
        super(parent, 'character-stats-root');
        this.stats = stats;
        const table = document.createElement('table');
        table.classList.add('character-stats-table');
        this.rootElem.appendChild(table);
        this.valueElems = [];
        this.stats.forEach(stat => {
            const row = document.createElement('tr');
            row.classList.add('character-stats-table-row');
            table.appendChild(row);
            const label = document.createElement('td');
            label.classList.add('character-stats-table-label');
            label.textContent = statNames[stat];
            row.appendChild(label);
            const value = document.createElement('td');
            value.classList.add('character-stats-table-value');
            row.appendChild(value);
            this.valueElems.push(value);
        });
        this.updateStats(new Stats());
        sim.characterStatsEmitter.on(() => {
            this.updateStats(new Stats(sim.getCurrentStats().finalStats));
        });
    }
    updateStats(newStats) {
        this.stats.forEach((stat, idx) => {
            let rawValue = newStats.getStat(stat);
            if (spellPowerTypeStats.includes(stat)) {
                rawValue = rawValue + newStats.getStat(Stat.StatSpellPower);
            }
            let displayStr = String(Math.round(rawValue));
            if (stat == Stat.StatMeleeHit) {
                displayStr += ` (${(rawValue / 15.8).toFixed(2)}%)`;
            }
            else if (stat == Stat.StatSpellHit) {
                displayStr += ` (${(rawValue / 12.6).toFixed(2)}%)`;
            }
            else if (stat == Stat.StatMeleeCrit || stat == Stat.StatSpellCrit) {
                displayStr += ` (${(rawValue / 22.08).toFixed(2)}%)`;
            }
            else if (stat == Stat.StatMeleeHaste || stat == Stat.StatSpellHaste) {
                displayStr += ` (${(rawValue / 15.76).toFixed(2)}%)`;
            }
            this.valueElems[idx].textContent = displayStr;
        });
    }
}
