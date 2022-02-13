import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { Potions } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js';
import { StrengthOfEarthType } from '/tbc/core/proto/common.js';
export function makeShow1hWeaponsSelector(parent, sim) {
    return new BooleanPicker(parent, sim, {
        extraCssClasses: [
            'show-1h-weapons-selector',
        ],
        label: '1H',
        changedEvent: (sim) => sim.show1hWeaponsChangeEmitter,
        getValue: (sim) => sim.getShow1hWeapons(),
        setValue: (eventID, sim, newValue) => {
            sim.setShow1hWeapons(eventID, newValue);
        },
    });
}
export function makeShow2hWeaponsSelector(parent, sim) {
    return new BooleanPicker(parent, sim, {
        extraCssClasses: [
            'show-2h-weapons-selector',
        ],
        label: '2H',
        changedEvent: (sim) => sim.show2hWeaponsChangeEmitter,
        getValue: (sim) => sim.getShow2hWeapons(),
        setValue: (eventID, sim, newValue) => {
            sim.setShow2hWeapons(eventID, newValue);
        },
    });
}
export function makeShowMatchingGemsSelector(parent, sim) {
    return new BooleanPicker(parent, sim, {
        extraCssClasses: [
            'show-matching-gems-selector',
        ],
        label: 'Match Socket',
        changedEvent: (sim) => sim.showMatchingGemsChangeEmitter,
        getValue: (sim) => sim.getShowMatchingGems(),
        setValue: (eventID, sim, newValue) => {
            sim.setShowMatchingGems(eventID, newValue);
        },
    });
}
export function makePhaseSelector(parent, sim) {
    return new EnumPicker(parent, sim, {
        extraCssClasses: [
            'phase-selector',
        ],
        values: [
            { name: 'Phase 1', value: 1 },
            { name: 'Phase 2', value: 2 },
            { name: 'Phase 3', value: 3 },
            { name: 'Phase 4', value: 4 },
            { name: 'Phase 5', value: 5 },
        ],
        changedEvent: (sim) => sim.phaseChangeEmitter,
        getValue: (sim) => sim.getPhase(),
        setValue: (eventID, sim, newValue) => {
            sim.setPhase(eventID, newValue);
        },
    });
}
export const StartingPotion = {
    type: 'enum',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'starting-potion-picker',
        ],
        label: 'Starting Potion',
        labelTooltip: 'If set, this potion will be used instead of the default potion for the first few uses.',
        values: [
            { name: 'None', value: Potions.UnknownPotion },
            { name: 'Destruction', value: Potions.DestructionPotion },
            { name: 'Haste', value: Potions.HastePotion },
            { name: 'Super Mana', value: Potions.SuperManaPotion },
        ],
        changedEvent: (player) => player.consumesChangeEmitter,
        getValue: (player) => player.getConsumes().startingPotion,
        setValue: (eventID, player, newValue) => {
            const newConsumes = player.getConsumes();
            newConsumes.startingPotion = newValue;
            player.setConsumes(eventID, newConsumes);
        },
    },
};
export const NumStartingPotions = {
    type: 'number',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'num-starting-potions-picker',
        ],
        label: '# to use',
        labelTooltip: 'The number of starting potions to use before going back to the default potion.',
        changedEvent: (player) => player.consumesChangeEmitter,
        getValue: (player) => player.getConsumes().numStartingPotions,
        setValue: (eventID, player, newValue) => {
            const newConsumes = player.getConsumes();
            newConsumes.numStartingPotions = newValue;
            player.setConsumes(eventID, newConsumes);
        },
        enableWhen: (player) => player.getConsumes().startingPotion != Potions.UnknownPotion,
    },
};
export const ShadowPriestDPS = {
    type: 'number',
    cssClass: 'shadow-priest-dps-picker',
    getModObject: (simUI) => simUI.player,
    config: {
        extraCssClasses: [
            'shadow-priest-dps-picker',
            'within-raid-sim-hide',
        ],
        label: 'Shadow Priest DPS',
        changedEvent: (player) => player.buffsChangeEmitter,
        getValue: (player) => player.getBuffs().shadowPriestDps,
        setValue: (eventID, player, newValue) => {
            const buffs = player.getBuffs();
            buffs.shadowPriestDps = newValue;
            player.setBuffs(eventID, buffs);
        },
    },
};
export const ISBUptime = {
    type: 'number',
    getModObject: (simUI) => simUI.sim.encounter.primaryTarget,
    config: {
        extraCssClasses: [
            'isb-uptime-picker',
            'within-raid-sim-hide',
        ],
        label: 'ISB Uptime %',
        changedEvent: (target) => target.debuffsChangeEmitter,
        getValue: (target) => Math.round(target.getDebuffs().isbUptime * 100),
        setValue: (eventID, target, newValue) => {
            const newDebuffs = target.getDebuffs();
            newDebuffs.isbUptime = newValue / 100;
            target.setDebuffs(eventID, newDebuffs);
        },
    },
};
export const ExposeWeaknessUptime = {
    type: 'number',
    getModObject: (simUI) => simUI.sim.encounter.primaryTarget,
    config: {
        extraCssClasses: [
            'expose-weakness-uptime-picker',
            'within-raid-sim-hide',
        ],
        label: 'Expose Weakness Uptime %',
        labelTooltip: 'Uptime for the Expose Weakness debuff, applied by 1 or more Survival hunters in your raid.',
        changedEvent: (target) => target.debuffsChangeEmitter,
        getValue: (target) => Math.round(target.getDebuffs().exposeWeaknessUptime * 100),
        setValue: (eventID, target, newValue) => {
            const newDebuffs = target.getDebuffs();
            newDebuffs.exposeWeaknessUptime = newValue / 100;
            target.setDebuffs(eventID, newDebuffs);
        },
    },
};
export const ExposeWeaknessHunterAgility = {
    type: 'number',
    getModObject: (simUI) => simUI.sim.encounter.primaryTarget,
    config: {
        extraCssClasses: [
            'expose-weakness-hunter-agility-picker',
            'within-raid-sim-hide',
        ],
        label: 'EW Hunter Agility',
        labelTooltip: 'The amount of agility on the Expose Weakness hunter.',
        changedEvent: (target) => target.debuffsChangeEmitter,
        getValue: (target) => Math.round(target.getDebuffs().exposeWeaknessHunterAgility),
        setValue: (eventID, target, newValue) => {
            const newDebuffs = target.getDebuffs();
            newDebuffs.exposeWeaknessHunterAgility = newValue;
            target.setDebuffs(eventID, newDebuffs);
        },
    },
};
export const SnapshotImprovedStrengthOfEarthTotem = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player.getParty(),
    config: {
        extraCssClasses: [
            'snapshot-improved-strength-of-earth-totem-picker',
            'within-raid-sim-hide',
        ],
        label: 'Snapshot Imp Strength of Earth',
        labelTooltip: 'An enhancement shaman in your party is snapshotting their improved Strength of Earth totem bonus from T4 2pc (+12 Strength) for the first 1:50s of the fight.',
        changedEvent: (party) => party.buffsChangeEmitter,
        getValue: (party) => party.getBuffs().snapshotImprovedStrengthOfEarthTotem,
        setValue: (eventID, party, newValue) => {
            const buffs = party.getBuffs();
            buffs.snapshotImprovedStrengthOfEarthTotem = newValue;
            party.setBuffs(eventID, buffs);
        },
        enableWhen: (party) => party.getBuffs().strengthOfEarthTotem == StrengthOfEarthType.Basic || party.getBuffs().strengthOfEarthTotem == StrengthOfEarthType.EnhancingTotems,
    },
};
export const SnapshotImprovedWrathOfAirTotem = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player.getParty(),
    config: {
        extraCssClasses: [
            'snapshot-improved-wrath-of-air-totem-picker',
            'within-raid-sim-hide',
        ],
        label: 'Snapshot Imp Wrath of Air',
        labelTooltip: 'An elemental shaman in your party is snapshotting their improved wrath of air totem bonus from T4 2pc (+20 spell power) for the first 1:50s of the fight.',
        changedEvent: (party) => party.buffsChangeEmitter,
        getValue: (party) => party.getBuffs().snapshotImprovedWrathOfAirTotem,
        setValue: (eventID, party, newValue) => {
            const buffs = party.getBuffs();
            buffs.snapshotImprovedWrathOfAirTotem = newValue;
            party.setBuffs(eventID, buffs);
        },
        enableWhen: (party) => party.getBuffs().wrathOfAirTotem == TristateEffect.TristateEffectRegular,
    },
};
export const SnapshotBsSolarianSapphire = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player.getParty(),
    config: {
        extraCssClasses: [
            'snapshot-bs-solarian-sapphire-picker',
            'within-raid-sim-hide',
        ],
        label: 'Snapshot BS Solarian\'s Sapphire',
        labelTooltip: 'A Warrior in your party is snapshotting their Battle Shout before combat, using the bonus from Solarian\'s Sapphire (+70 attack power) for the first 1:50s of the fight.',
        changedEvent: (party) => party.buffsChangeEmitter,
        getValue: (party) => party.getBuffs().snapshotBsSolarianSapphire,
        setValue: (eventID, party, newValue) => {
            const buffs = party.getBuffs();
            buffs.snapshotBsSolarianSapphire = newValue;
            party.setBuffs(eventID, buffs);
        },
        enableWhen: (party) => party.getBuffs().battleShout > 0 && !party.getBuffs().bsSolarianSapphire,
    },
};
export const SnapshotBsT2 = {
    type: 'boolean',
    getModObject: (simUI) => simUI.player.getParty(),
    config: {
        extraCssClasses: [
            'snapshot-bs-t2-picker',
            'within-raid-sim-hide',
        ],
        label: 'Snapshot BS T2',
        labelTooltip: 'A Warrior in your party is snapshotting their Battle Shout before combat, using the bonus from T2 3pc (+30 attack power) for the first 1:50s of the fight.',
        changedEvent: (party) => party.buffsChangeEmitter,
        getValue: (party) => party.getBuffs().snapshotBsT2,
        setValue: (eventID, party, newValue) => {
            const buffs = party.getBuffs();
            buffs.snapshotBsT2 = newValue;
            party.setBuffs(eventID, buffs);
        },
        enableWhen: (party) => party.getBuffs().battleShout > 0,
    },
};