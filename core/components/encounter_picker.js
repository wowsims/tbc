import { EncounterType } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { ListPicker } from '/tbc/core/components/list_picker.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { statNames } from '/tbc/core/proto_utils/names.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { Component } from './component.js';
import { Popup } from './popup.js';
export class EncounterPicker extends Component {
    constructor(parent, modEncounter, config) {
        super(parent, 'encounter-picker-root');
        addEncounterFieldPickers(this.rootElem, modEncounter, config.showExecuteProportion);
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
        if (config.showNumTargets) {
            new NumberPicker(this.rootElem, modEncounter, {
                label: '# of Targets',
                changedEvent: (encounter) => encounter.changeEmitter,
                getValue: (encounter) => encounter.getNumTargets(),
                setValue: (eventID, encounter, newValue) => {
                    encounter.setNumTargets(eventID, newValue);
                },
                enableWhen: (encounter) => encounter.getType() == EncounterType.EncounterTypeSimple,
            });
        }
        // Simple/Custom/Preset [Edit Button]
        const advancedButton = document.createElement('button');
        advancedButton.classList.add('sim-button', 'advanced-button', 'experimental');
        advancedButton.textContent = 'ADVANCED';
        advancedButton.addEventListener('click', () => new AdvancedEncounterPicker(this.rootElem, modEncounter));
        this.rootElem.appendChild(advancedButton);
    }
}
class AdvancedEncounterPicker extends Popup {
    constructor(parent, encounter) {
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
        const encounterTypeContainer = this.rootElem.getElementsByClassName('encounter-type')[0];
        new EnumPicker(encounterTypeContainer, this.encounter, {
            label: 'ENCOUNTER',
            values: [
                { name: 'Simple', value: EncounterType.EncounterTypeSimple },
                { name: 'Custom', value: EncounterType.EncounterTypeCustom },
            ].concat(getEnumValues(EncounterType).filter(val => ![EncounterType.EncounterTypeSimple, EncounterType.EncounterTypeCustom].includes(val)).map((val, i) => {
                return {
                    name: '',
                    value: val,
                };
            })),
            changedEvent: (encounter) => encounter.typeChangeEmitter,
            getValue: (encounter) => encounter.getType(),
            setValue: (eventID, encounter, newValue) => {
                encounter.setType(eventID, newValue);
            },
        });
        const header = this.rootElem.getElementsByClassName('encounter-header')[0];
        const targetsElem = this.rootElem.getElementsByClassName('encounter-targets')[0];
        addEncounterFieldPickers(header, this.encounter, true);
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
			<div class="target-picker-section target-picker-section3"></div>
		`;
        const section1 = this.rootElem.getElementsByClassName('target-picker-section1')[0];
        const section2 = this.rootElem.getElementsByClassName('target-picker-section2')[0];
        const section3 = this.rootElem.getElementsByClassName('target-picker-section3')[0];
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
        ALL_TARGET_STATS.forEach(stat => {
            new NumberPicker(section2, modTarget, {
                label: statNames[stat],
                changedEvent: (target) => target.statsChangeEmitter,
                getValue: (target) => target.getStats().getStat(stat),
                setValue: (eventID, target, newValue) => {
                    target.setStats(eventID, target.getStats().withStat(stat, newValue));
                },
            });
        });
    }
}
function addEncounterFieldPickers(rootElem, encounter, showExecuteProportion) {
    new NumberPicker(rootElem, encounter, {
        label: 'Duration',
        changedEvent: (encounter) => encounter.durationChangeEmitter,
        getValue: (encounter) => encounter.getDuration(),
        setValue: (eventID, encounter, newValue) => {
            encounter.setDuration(eventID, newValue);
        },
    });
    new NumberPicker(rootElem, encounter, {
        label: 'Duration +/-',
        changedEvent: (encounter) => encounter.durationChangeEmitter,
        getValue: (encounter) => encounter.getDurationVariation(),
        setValue: (eventID, encounter, newValue) => {
            encounter.setDurationVariation(eventID, newValue);
        },
    });
    if (showExecuteProportion) {
        new NumberPicker(rootElem, encounter, {
            label: 'Execute Duration (%)',
            labelTooltip: 'Percentage of the total encounter duration, for which the targets will be considered to be in execute range (< 20% HP) for the purpose of effects like Warrior Execute or Mage Molten Fury.',
            changedEvent: (encounter) => encounter.executeProportionChangeEmitter,
            getValue: (encounter) => encounter.getExecuteProportion() * 100,
            setValue: (eventID, encounter, newValue) => {
                encounter.setExecuteProportion(eventID, newValue / 100);
            },
        });
    }
}
const ALL_TARGET_STATS = [
    Stat.StatArmor,
    Stat.StatArcaneResistance,
    Stat.StatFireResistance,
    Stat.StatFrostResistance,
    Stat.StatNatureResistance,
    Stat.StatShadowResistance,
    Stat.StatAttackPower,
    Stat.StatBlockValue,
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
