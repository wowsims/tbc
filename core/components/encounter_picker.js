import { MobType } from '/tbc/core/proto/common.js';
import { SpellSchool } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { ListPicker } from '/tbc/core/components/list_picker.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { isTankSpec } from '/tbc/core/proto_utils/utils.js';
import { statNames } from '/tbc/core/proto_utils/names.js';
import { Component } from './component.js';
import { Popup } from './popup.js';
import * as Mechanics from '/tbc/core/constants/mechanics.js';
export class EncounterPicker extends Component {
    constructor(parent, modEncounter, config, simUI) {
        super(parent, 'encounter-picker-root');
        addEncounterFieldPickers(this.rootElem, modEncounter, config.showExecuteProportion);
        // Need to wait so that the encounter and target presets will be loaded.
        modEncounter.sim.waitForInit().then(() => {
            const presetTargets = modEncounter.sim.getAllPresetTargets();
            new EnumPicker(this.rootElem, modEncounter, {
                extraCssClasses: ['npc-picker'],
                label: 'NPC',
                labelTooltip: 'Selects a preset NPC configuration.',
                values: [
                    { name: 'Custom', value: -1 },
                ].concat(presetTargets.map((pe, i) => {
                    return {
                        name: pe.path,
                        value: i,
                    };
                })),
                changedEvent: (encounter) => encounter.changeEmitter,
                getValue: (encounter) => presetTargets.findIndex(pe => encounter.primaryTarget.matchesPreset(pe)),
                setValue: (eventID, encounter, newValue) => {
                    if (newValue != -1) {
                        encounter.primaryTarget.applyPreset(eventID, presetTargets[newValue]);
                    }
                },
            });
            new EnumPicker(this.rootElem, modEncounter, {
                label: 'Target Level',
                values: [
                    { name: '73', value: 73 },
                    { name: '72', value: 72 },
                    { name: '71', value: 71 },
                    { name: '70', value: 70 },
                ],
                changedEvent: (encounter) => encounter.changeEmitter,
                getValue: (encounter) => encounter.primaryTarget.getLevel(),
                setValue: (eventID, encounter, newValue) => {
                    encounter.primaryTarget.setLevel(eventID, newValue);
                },
            });
            new EnumPicker(this.rootElem, modEncounter, {
                label: 'Mob Type',
                values: mobTypeEnumValues,
                changedEvent: (encounter) => encounter.changeEmitter,
                getValue: (encounter) => encounter.primaryTarget.getMobType(),
                setValue: (eventID, encounter, newValue) => {
                    encounter.primaryTarget.setMobType(eventID, newValue);
                },
            });
            if (config.simpleTargetStats) {
                config.simpleTargetStats.forEach(stat => {
                    new NumberPicker(this.rootElem, modEncounter, {
                        label: statNames[stat],
                        changedEvent: (encounter) => encounter.changeEmitter,
                        getValue: (encounter) => encounter.primaryTarget.getStats().getStat(stat),
                        setValue: (eventID, encounter, newValue) => {
                            encounter.primaryTarget.setStats(eventID, encounter.primaryTarget.getStats().withStat(stat, newValue));
                        },
                    });
                });
            }
            if (simUI.isIndividualSim() && isTankSpec(simUI.player.spec)) {
                new NumberPicker(this.rootElem, modEncounter, {
                    label: 'Min Base Damage',
                    labelTooltip: 'Base damage for auto attacks, i.e. lowest roll with 0 AP against a 0-armor Player.',
                    changedEvent: (encounter) => encounter.changeEmitter,
                    getValue: (encounter) => encounter.primaryTarget.getMinBaseDamage(),
                    setValue: (eventID, encounter, newValue) => {
                        encounter.primaryTarget.setMinBaseDamage(eventID, newValue);
                    },
                });
            }
            const advancedButton = document.createElement('button');
            advancedButton.classList.add('sim-button', 'advanced-button');
            advancedButton.textContent = 'ADVANCED';
            advancedButton.addEventListener('click', () => new AdvancedEncounterPicker(this.rootElem, modEncounter, simUI));
            this.rootElem.appendChild(advancedButton);
        });
    }
}
class AdvancedEncounterPicker extends Popup {
    constructor(parent, encounter, simUI) {
        super(parent);
        this.encounter = encounter;
        this.rootElem.classList.add('advanced-encounter-picker');
        this.rootElem.innerHTML = `
			<div class="encounter-type"></div>
			<div class="encounter-header">
			</div>
			<div class="encounter-targets">
			</div>
		`;
        this.addCloseButton();
        const presetEncounters = this.encounter.sim.getAllPresetEncounters();
        const encounterTypeContainer = this.rootElem.getElementsByClassName('encounter-type')[0];
        new EnumPicker(encounterTypeContainer, this.encounter, {
            label: 'ENCOUNTER',
            values: [
                { name: 'Custom', value: -1 },
            ].concat(presetEncounters.map((pe, i) => {
                return {
                    name: pe.path,
                    value: i,
                };
            })),
            changedEvent: (encounter) => encounter.changeEmitter,
            getValue: (encounter) => presetEncounters.findIndex(pe => encounter.matchesPreset(pe)),
            setValue: (eventID, encounter, newValue) => {
                if (newValue != -1) {
                    encounter.applyPreset(eventID, presetEncounters[newValue]);
                }
            },
        });
        const header = this.rootElem.getElementsByClassName('encounter-header')[0];
        const targetsElem = this.rootElem.getElementsByClassName('encounter-targets')[0];
        addEncounterFieldPickers(header, this.encounter, true);
        if (!simUI.isIndividualSim()) {
            new BooleanPicker(header, encounter, {
                label: 'Use Health',
                labelTooltip: 'Uses a damage limit in place of a duration limit. Damage limit is equal to sum of all targets health.',
                changedEvent: (encounter) => encounter.changeEmitter,
                getValue: (encounter) => encounter.getUseHealth(),
                setValue: (eventID, encounter, newValue) => {
                    encounter.setUseHealth(eventID, newValue);
                },
            });
        }
        new ListPicker(targetsElem, this.encounter, {
            extraCssClasses: [
                'targets-picker',
            ],
            itemLabel: 'Target',
            changedEvent: (encounter) => encounter.targetsChangeEmitter,
            getValue: (encounter) => encounter.getTargets(),
            setValue: (eventID, encounter, newValue) => {
                encounter.setTargets(eventID, newValue);
            },
            newItem: () => Target.fromDefaults(TypedEvent.nextEventID(), this.encounter.sim),
            copyItem: (oldItem) => oldItem.clone(TypedEvent.nextEventID()),
            newItemPicker: (parent, target) => new TargetPicker(parent, target),
        });
    }
}
class TargetPicker extends Component {
    constructor(parent, modTarget) {
        super(parent, 'target-picker-root');
        this.rootElem.innerHTML = `
			<div class="target-picker-section target-picker-section1"></div>
			<div class="target-picker-section target-picker-section2"></div>
			<div class="target-picker-section target-picker-section3 threat-metrics"></div>
		`;
        const encounter = modTarget.sim.encounter;
        const section1 = this.rootElem.getElementsByClassName('target-picker-section1')[0];
        const section2 = this.rootElem.getElementsByClassName('target-picker-section2')[0];
        const section3 = this.rootElem.getElementsByClassName('target-picker-section3')[0];
        const presetTargets = modTarget.sim.getAllPresetTargets();
        new EnumPicker(section1, modTarget, {
            extraCssClasses: ['npc-picker'],
            label: 'NPC',
            labelTooltip: 'Selects a preset NPC configuration.',
            values: [
                { name: 'Custom', value: -1 },
            ].concat(presetTargets.map((pe, i) => {
                return {
                    name: pe.path,
                    value: i,
                };
            })),
            changedEvent: (target) => target.changeEmitter,
            getValue: (target) => presetTargets.findIndex(pe => target.matchesPreset(pe)),
            setValue: (eventID, target, newValue) => {
                if (newValue != -1) {
                    target.applyPreset(eventID, presetTargets[newValue]);
                }
            },
        });
        new EnumPicker(section1, modTarget, {
            extraCssClasses: ['ai-picker'],
            label: 'AI',
            labelTooltip: `
				<p>Determines the target\'s ability rotation.</p>
				<p>Note that most rotations are not yet implemented.</p>
			`,
            values: [
                { name: 'None', value: 0 },
            ].concat(presetTargets.map(pe => {
                return {
                    name: pe.path,
                    value: pe.target.id,
                };
            })),
            changedEvent: (target) => target.changeEmitter,
            getValue: (target) => target.getId(),
            setValue: (eventID, target, newValue) => {
                target.setId(eventID, newValue);
            },
        });
        new EnumPicker(section1, modTarget, {
            label: 'Level',
            values: [
                { name: '73', value: 73 },
                { name: '72', value: 72 },
                { name: '71', value: 71 },
                { name: '70', value: 70 },
            ],
            changedEvent: (target) => target.levelChangeEmitter,
            getValue: (target) => target.getLevel(),
            setValue: (eventID, target, newValue) => {
                target.setLevel(eventID, newValue);
            },
        });
        new EnumPicker(section1, modTarget, {
            label: 'Mob Type',
            values: mobTypeEnumValues,
            changedEvent: (target) => target.mobTypeChangeEmitter,
            getValue: (target) => target.getMobType(),
            setValue: (eventID, target, newValue) => {
                target.setMobType(eventID, newValue);
            },
        });
        new EnumPicker(section1, modTarget, {
            extraCssClasses: ['threat-metrics'],
            label: 'Tanked By',
            labelTooltip: 'Determines which player in the raid this enemy will attack. If no player is assigned to the specified tank slot, this enemy will not attack.',
            values: [
                { name: 'None', value: -1 },
                { name: 'Main Tank', value: 0 },
                { name: 'Tank 2', value: 1 },
                { name: 'Tank 3', value: 2 },
                { name: 'Tank 4', value: 3 },
            ],
            changedEvent: (target) => target.propChangeEmitter,
            getValue: (target) => target.getTankIndex(),
            setValue: (eventID, target, newValue) => {
                target.setTankIndex(eventID, newValue);
            },
        });
        ALL_TARGET_STATS.forEach(statData => {
            const stat = statData.stat;
            new NumberPicker(section2, modTarget, {
                extraCssClasses: statData.extraCssClasses,
                label: statNames[stat],
                labelTooltip: statData.tooltip,
                changedEvent: (target) => target.statsChangeEmitter,
                getValue: (target) => target.getStats().getStat(stat),
                setValue: (eventID, target, newValue) => {
                    target.setStats(eventID, target.getStats().withStat(stat, newValue));
                },
            });
        });
        new NumberPicker(section3, modTarget, {
            label: 'Swing Speed',
            labelTooltip: 'Time in seconds between auto attacks. Set to 0 to disable auto attacks.',
            changedEvent: (target) => target.propChangeEmitter,
            getValue: (target) => target.getSwingSpeed(),
            setValue: (eventID, target, newValue) => {
                target.setSwingSpeed(eventID, newValue);
            },
        });
        new NumberPicker(section3, modTarget, {
            label: 'Min Base Damage',
            labelTooltip: 'Base damage for auto attacks, i.e. lowest roll with 0 AP against a 0-armor Player.',
            changedEvent: (target) => target.propChangeEmitter,
            getValue: (target) => target.getMinBaseDamage(),
            setValue: (eventID, target, newValue) => {
                target.setMinBaseDamage(eventID, newValue);
            },
        });
        new BooleanPicker(section3, modTarget, {
            label: 'Dual Wield',
            labelTooltip: 'Uses 2 separate weapons to attack.',
            changedEvent: (target) => target.propChangeEmitter,
            getValue: (target) => target.getDualWield(),
            setValue: (eventID, target, newValue) => {
                target.setDualWield(eventID, newValue);
            },
        });
        new BooleanPicker(section3, modTarget, {
            label: 'DW Miss Penalty',
            labelTooltip: 'Enables the Dual Wield Miss Penalty (+19% chance to miss) if dual wielding. Bosses in Hyjal/BT/SWP usually have this disabled to stop tanks from avoidance stacking.',
            changedEvent: (target) => target.changeEmitter,
            getValue: (target) => target.getDualWieldPenalty(),
            setValue: (eventID, target, newValue) => {
                target.setDualWieldPenalty(eventID, newValue);
            },
            enableWhen: (target) => target.getDualWield(),
        });
        new BooleanPicker(section3, modTarget, {
            label: 'Can Crush',
            labelTooltip: 'Whether crushing blows should be included in the attack table. Only applies to level 73 enemies.',
            changedEvent: (target) => target.changeEmitter,
            getValue: (target) => target.getCanCrush(),
            setValue: (eventID, target, newValue) => {
                target.setCanCrush(eventID, newValue);
            },
            enableWhen: (target) => target.getLevel() == Mechanics.BOSS_LEVEL,
        });
        new BooleanPicker(section3, modTarget, {
            label: 'Parry Haste',
            labelTooltip: 'Whether this enemy will gain parry haste when parrying attacks.',
            changedEvent: (target) => target.propChangeEmitter,
            getValue: (target) => target.getParryHaste(),
            setValue: (eventID, target, newValue) => {
                target.setParryHaste(eventID, newValue);
            },
        });
        new EnumPicker(section3, modTarget, {
            label: 'Spell School',
            labelTooltip: 'Type of damage caused by auto attacks. This is usually Physical, but some enemies have elemental attacks.',
            values: [
                { name: 'Physical', value: SpellSchool.SpellSchoolPhysical },
                { name: 'Arcane', value: SpellSchool.SpellSchoolArcane },
                { name: 'Fire', value: SpellSchool.SpellSchoolFire },
                { name: 'Frost', value: SpellSchool.SpellSchoolFrost },
                { name: 'Holy', value: SpellSchool.SpellSchoolHoly },
                { name: 'Nature', value: SpellSchool.SpellSchoolNature },
                { name: 'Shadow', value: SpellSchool.SpellSchoolShadow },
            ],
            changedEvent: (target) => target.propChangeEmitter,
            getValue: (target) => target.getSpellSchool(),
            setValue: (eventID, target, newValue) => {
                target.setSpellSchool(eventID, newValue);
            },
        });
        new BooleanPicker(section3, modTarget, {
            label: 'Sunwell Radiance',
            labelTooltip: 'Reduces the chance for this enemy\'s attacks to be dodged by 20% and be missed by 5%. All Sunwell Plateau bosses have this.',
            changedEvent: (target) => target.changeEmitter,
            getValue: (target) => target.getSuppressDodge(),
            setValue: (eventID, target, newValue) => {
                target.setSuppressDodge(eventID, newValue);
            },
            enableWhen: (target) => target.getLevel() == Mechanics.BOSS_LEVEL,
        });
    }
}
function addEncounterFieldPickers(rootElem, encounter, showExecuteProportion) {
    new NumberPicker(rootElem, encounter, {
        label: 'Duration',
        labelTooltip: 'The fight length for each sim iteration, in seconds.',
        changedEvent: (encounter) => encounter.changeEmitter,
        getValue: (encounter) => encounter.getDuration(),
        setValue: (eventID, encounter, newValue) => {
            encounter.setDuration(eventID, newValue);
        },
        enableWhen: (obj) => { return !encounter.getUseHealth(); },
    });
    new NumberPicker(rootElem, encounter, {
        label: 'Duration +/-',
        labelTooltip: 'Adds a random amount of time, in seconds, between [value, -1 * value] to each sim iteration. For example, setting Duration to 180 and Duration +/- to 10 will result in random durations between 170s and 190s.',
        changedEvent: (encounter) => encounter.changeEmitter,
        getValue: (encounter) => encounter.getDurationVariation(),
        setValue: (eventID, encounter, newValue) => {
            encounter.setDurationVariation(eventID, newValue);
        },
        enableWhen: (obj) => { return !encounter.getUseHealth(); },
    });
    if (showExecuteProportion) {
        new NumberPicker(rootElem, encounter, {
            label: 'Execute Duration (%)',
            labelTooltip: 'Percentage of the total encounter duration, for which the targets will be considered to be in execute range (< 20% HP) for the purpose of effects like Warrior Execute or Mage Molten Fury.',
            changedEvent: (encounter) => encounter.changeEmitter,
            getValue: (encounter) => encounter.getExecuteProportion() * 100,
            setValue: (eventID, encounter, newValue) => {
                encounter.setExecuteProportion(eventID, newValue / 100);
            },
            enableWhen: (obj) => { return !encounter.getUseHealth(); },
        });
    }
}
const ALL_TARGET_STATS = [
    { stat: Stat.StatHealth, tooltip: '', extraCssClasses: [] },
    { stat: Stat.StatArmor, tooltip: '', extraCssClasses: [] },
    { stat: Stat.StatArcaneResistance, tooltip: '', extraCssClasses: [] },
    { stat: Stat.StatFireResistance, tooltip: '', extraCssClasses: [] },
    { stat: Stat.StatFrostResistance, tooltip: '', extraCssClasses: [] },
    { stat: Stat.StatNatureResistance, tooltip: '', extraCssClasses: [] },
    { stat: Stat.StatShadowResistance, tooltip: '', extraCssClasses: [] },
    { stat: Stat.StatAttackPower, tooltip: '', extraCssClasses: ['threat-metrics'] },
    { stat: Stat.StatBlockValue, tooltip: '', extraCssClasses: ['threat-metrics'] },
];
const mobTypeEnumValues = [
    { name: 'None', value: MobType.MobTypeUnknown },
    { name: 'Beast', value: MobType.MobTypeBeast },
    { name: 'Demon', value: MobType.MobTypeDemon },
    { name: 'Dragonkin', value: MobType.MobTypeDragonkin },
    { name: 'Elemental', value: MobType.MobTypeElemental },
    { name: 'Giant', value: MobType.MobTypeGiant },
    { name: 'Humanoid', value: MobType.MobTypeHumanoid },
    { name: 'Mechanical', value: MobType.MobTypeMechanical },
    { name: 'Undead', value: MobType.MobTypeUndead },
];
