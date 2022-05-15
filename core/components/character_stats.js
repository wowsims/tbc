import { Stat } from '/tbc/core/proto/common.js';
import { statNames, statOrder } from '/tbc/core/proto_utils/names.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import * as Mechanics from '/tbc/core/constants/mechanics.js';
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
    constructor(parent, player, stats, modifyDisplayStats, statBreakdowns) {
        super(parent, 'character-stats-root');
        this.stats = statOrder.filter(stat => stats.includes(stat));
        this.player = player;
        this.modifyDisplayStats = modifyDisplayStats;
        this.statBreakdowns = statBreakdowns;
        const table = document.createElement('table');
        table.classList.add('character-stats-table');
        this.rootElem.appendChild(table);
        this.valueElems = [];
        this.tooltipElems = [];
        this.stats.forEach(stat => {
            const row = document.createElement('tr');
            row.classList.add('character-stats-table-row');
            row.innerHTML = `
				<td class="character-stats-table-label">
					<span>${statNames[stat].toUpperCase()}<span>
					<span class="character-stats-table-tooltip fas fa-search"></span>
				</td>
				<td class="character-stats-table-value"></td>
			`;
            table.appendChild(row);
            const valueElem = row.getElementsByClassName('character-stats-table-value')[0];
            this.valueElems.push(valueElem);
            const tooltipElem = row.getElementsByClassName('character-stats-table-tooltip')[0];
            this.tooltipElems.push(tooltipElem);
        });
        this.updateStats(new Stats(player.getCurrentStats().finalStats));
        TypedEvent.onAny([player.currentStatsEmitter, player.sim.changeEmitter]).on(() => {
            this.updateStats(new Stats(player.getCurrentStats().finalStats));
        });
    }
    updateStats(newStats) {
        if (this.modifyDisplayStats) {
            newStats = this.modifyDisplayStats(this.player, newStats);
        }
        const breakdowns = this.statBreakdowns ? this.statBreakdowns(this.player, newStats) : null;
        this.stats.forEach((stat, idx) => {
            let rawValue = newStats.getStat(stat);
            if (spellPowerTypeStats.includes(stat)) {
                rawValue = rawValue + newStats.getStat(Stat.StatSpellPower);
            }
            const displayStr = CharacterStats.statDisplayString(stat, rawValue);
            this.valueElems[idx].textContent = displayStr;
            const breakdown = breakdowns ? breakdowns[stat] : null;
            if (breakdown) {
                tippy(this.tooltipElems[idx], {
                    'content': breakdown.map(item => `
						<div class="character-stats-tooltip-row">
							<span>${item.label}:</span>
							<span>${CharacterStats.statDisplayString(stat, item.value)}</span>
						</div>
					`).join(''),
                    'allowHTML': true,
                });
                this.tooltipElems[idx].classList.remove('hide');
            }
            else {
                this.tooltipElems[idx].classList.add('hide');
            }
        });
    }
    static statDisplayString(stat, rawValue) {
        let displayStr = String(Math.round(rawValue));
        if (stat == Stat.StatMeleeHit) {
            displayStr += ` (${(rawValue / Mechanics.MELEE_HIT_RATING_PER_HIT_CHANCE).toFixed(2)}%)`;
        }
        else if (stat == Stat.StatSpellHit) {
            displayStr += ` (${(rawValue / Mechanics.SPELL_HIT_RATING_PER_HIT_CHANCE).toFixed(2)}%)`;
        }
        else if (stat == Stat.StatMeleeCrit || stat == Stat.StatSpellCrit) {
            displayStr += ` (${(rawValue / Mechanics.SPELL_CRIT_RATING_PER_CRIT_CHANCE).toFixed(2)}%)`;
        }
        else if (stat == Stat.StatMeleeHaste || stat == Stat.StatSpellHaste) {
            displayStr += ` (${(rawValue / Mechanics.HASTE_RATING_PER_HASTE_PERCENT).toFixed(2)}%)`;
        }
        else if (stat == Stat.StatExpertise) {
            displayStr += ` (${(Math.floor(rawValue / Mechanics.EXPERTISE_PER_QUARTER_PERCENT_REDUCTION)).toFixed(0)})`;
        }
        return displayStr;
    }
}
