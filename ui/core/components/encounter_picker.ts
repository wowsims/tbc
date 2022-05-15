import { MobType } from '/tbc/core/proto/common.js';
import { SpellSchool } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/encounter.js';
import { Target } from '/tbc/core/target.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { BooleanPicker } from '/tbc/core/components/boolean_picker.js';
import { EnumPicker, EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { ListPicker, ListPickerConfig } from '/tbc/core/components/list_picker.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { statNames } from '/tbc/core/proto_utils/names.js';
import { getEnumValues } from '/tbc/core/utils.js';

import { Component } from './component.js';
import { Popup } from './popup.js';

import * as Mechanics from '/tbc/core/constants/mechanics.js';

export interface EncounterPickerConfig {
	simpleTargetStats?: Array<Stat>;
	showExecuteProportion: boolean;
}

export class EncounterPicker extends Component {
	constructor(parent: HTMLElement, modEncounter: Encounter, config: EncounterPickerConfig) {
		super(parent, 'encounter-picker-root');

		addEncounterFieldPickers(this.rootElem, modEncounter, config.showExecuteProportion);

		new EnumPicker<Encounter>(this.rootElem, modEncounter, {
			label: 'Target Level',
			values: [
				{ name: '73', value: 73 },
				{ name: '72', value: 72 },
				{ name: '71', value: 71 },
				{ name: '70', value: 70 },
			],
			changedEvent: (encounter: Encounter) => encounter.changeEmitter,
			getValue: (encounter: Encounter) => encounter.primaryTarget.getLevel(),
			setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
				encounter.primaryTarget.setLevel(eventID, newValue);
			},
		});

		new EnumPicker(this.rootElem, modEncounter, {
			label: 'Mob Type',
			values: mobTypeEnumValues,
			changedEvent: (encounter: Encounter) => encounter.changeEmitter,
			getValue: (encounter: Encounter) => encounter.primaryTarget.getMobType(),
			setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
				encounter.primaryTarget.setMobType(eventID, newValue);
			},
		});

		if (config.simpleTargetStats) {
			config.simpleTargetStats.forEach(stat => {
				new NumberPicker(this.rootElem, modEncounter, {
					label: statNames[stat],
					changedEvent: (encounter: Encounter) => encounter.changeEmitter,
					getValue: (encounter: Encounter) => encounter.primaryTarget.getStats().getStat(stat),
					setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
						encounter.primaryTarget.setStats(eventID, encounter.primaryTarget.getStats().withStat(stat, newValue));
					},
				});
			});
		}
		
		const advancedButton = document.createElement('button');
		advancedButton.classList.add('sim-button', 'advanced-button');
		advancedButton.textContent = 'ADVANCED';
		advancedButton.addEventListener('click', () => new AdvancedEncounterPicker(this.rootElem, modEncounter));
		this.rootElem.appendChild(advancedButton);
	}
}

class AdvancedEncounterPicker extends Popup {
	private readonly encounter: Encounter;

	constructor(parent: HTMLElement, encounter: Encounter) {
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

		const encounterTypeContainer = this.rootElem.getElementsByClassName('encounter-type')[0] as HTMLElement;
		new EnumPicker<Encounter>(encounterTypeContainer, this.encounter, {
			label: 'ENCOUNTER',
			values: [
				{ name: 'Custom', value: -1 },
			].concat(presetEncounters.map((pe, i) => {
				return {
					name: pe.path,
					value: i,
				};
			})),
			changedEvent: (encounter: Encounter) => encounter.changeEmitter,
			getValue: (encounter: Encounter) => presetEncounters.findIndex(pe => encounter.matchesPreset(pe)),
			setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
				if (newValue != -1) {
					encounter.applyPreset(eventID, presetEncounters[newValue]);
				}
			},
		});

		const header = this.rootElem.getElementsByClassName('encounter-header')[0] as HTMLElement;
		const targetsElem = this.rootElem.getElementsByClassName('encounter-targets')[0] as HTMLElement;

		addEncounterFieldPickers(header, this.encounter, true);

		new ListPicker<Encounter, Target>(targetsElem, this.encounter, {
			extraCssClasses: [
				'targets-picker',
			],
			itemLabel: 'Target',
			changedEvent: (encounter: Encounter) => encounter.targetsChangeEmitter,
			getValue: (encounter: Encounter) => encounter.getTargets(),
			setValue: (eventID: EventID, encounter: Encounter, newValue: Array<Target>) => {
				encounter.setTargets(eventID, newValue);
			},
			newItem: () => Target.fromDefaults(TypedEvent.nextEventID(), this.encounter.sim),
			newItemPicker: (parent: HTMLElement, target: Target) => new TargetPicker(parent, target),
		});
	}
}

class TargetPicker extends Component {
	constructor(parent: HTMLElement, modTarget: Target) {
		super(parent, 'target-picker-root');
		this.rootElem.innerHTML = `
			<div class="target-picker-section target-picker-section1"></div>
			<div class="target-picker-section target-picker-section2"></div>
			<div class="target-picker-section target-picker-section3"></div>
		`;
		const section1 = this.rootElem.getElementsByClassName('target-picker-section1')[0] as HTMLElement;
		const section2 = this.rootElem.getElementsByClassName('target-picker-section2')[0] as HTMLElement;
		const section3 = this.rootElem.getElementsByClassName('target-picker-section3')[0] as HTMLElement;

		new EnumPicker<Target>(section1, modTarget, {
			label: 'Level',
			values: [
				{ name: '73', value: 73 },
				{ name: '72', value: 72 },
				{ name: '71', value: 71 },
				{ name: '70', value: 70 },
			],
			changedEvent: (target: Target) => target.levelChangeEmitter,
			getValue: (target: Target) => target.getLevel(),
			setValue: (eventID: EventID, target: Target, newValue: number) => {
				target.setLevel(eventID, newValue);
			},
		});
		new EnumPicker(section1, modTarget, {
			label: 'Mob Type',
			values: mobTypeEnumValues,
			changedEvent: (target: Target) => target.mobTypeChangeEmitter,
			getValue: (target: Target) => target.getMobType(),
			setValue: (eventID: EventID, target: Target, newValue: number) => {
				target.setMobType(eventID, newValue);
			},
		});
		new EnumPicker<Target>(section1, modTarget, {
			label: 'Tanked By',
			values: [
				{ name: 'None', value: -1 },
				{ name: 'Main Tank', value: 0 },
				{ name: 'Tank 2', value: 1 },
				{ name: 'Tank 3', value: 2 },
			],
			changedEvent: (target: Target) => target.propChangeEmitter,
			getValue: (target: Target) => target.getTankIndex(),
			setValue: (eventID: EventID, target: Target, newValue: number) => {
				target.setTankIndex(eventID, newValue);
			},
		});

		ALL_TARGET_STATS.forEach(stat => {
			new NumberPicker(section2, modTarget, {
				label: statNames[stat],
				changedEvent: (target: Target) => target.statsChangeEmitter,
				getValue: (target: Target) => target.getStats().getStat(stat),
				setValue: (eventID: EventID, target: Target, newValue: number) => {
					target.setStats(eventID, target.getStats().withStat(stat, newValue));
				},
			});
		});

		new NumberPicker(section3, modTarget, {
			label: 'Swing Speed',
			labelTooltip: 'Time in seconds between auto attacks. Set to 0 to disable auto attacks.',
			changedEvent: (target: Target) => target.propChangeEmitter,
			getValue: (target: Target) => target.getSwingSpeed(),
			setValue: (eventID: EventID, target: Target, newValue: number) => {
				target.setSwingSpeed(eventID, newValue);
			},
		});
		new NumberPicker(section3, modTarget, {
			label: 'Min Base Damage',
			labelTooltip: 'Base damage for auto attacks, i.e. lowest roll with 0 AP against a 0-armor Player.',
			changedEvent: (target: Target) => target.propChangeEmitter,
			getValue: (target: Target) => target.getMinBaseDamage(),
			setValue: (eventID: EventID, target: Target, newValue: number) => {
				target.setMinBaseDamage(eventID, newValue);
			},
		});
		new BooleanPicker(section3, modTarget, {
			label: 'Dual Wield',
			labelTooltip: 'Uses 2 separate weapons to attack.',
			changedEvent: (target: Target) => target.propChangeEmitter,
			getValue: (target: Target) => target.getDualWield(),
			setValue: (eventID: EventID, target: Target, newValue: boolean) => {
				target.setDualWield(eventID, newValue);
			},
		});
		new BooleanPicker(section3, modTarget, {
			label: 'Can Crush',
			labelTooltip: 'Whether crushing blows should be included in the attack table. Only applies to level 73 enemies.',
			changedEvent: (target: Target) => target.changeEmitter,
			getValue: (target: Target) => target.getCanCrush(),
			setValue: (eventID: EventID, target: Target, newValue: boolean) => {
				target.setCanCrush(eventID, newValue);
			},
			enableWhen: (target: Target) => target.getLevel() == Mechanics.BOSS_LEVEL,
		});
		new BooleanPicker(section3, modTarget, {
			label: 'Parry Haste',
			labelTooltip: 'Whether this enemy will gain parry haste when parrying attacks.',
			changedEvent: (target: Target) => target.propChangeEmitter,
			getValue: (target: Target) => target.getParryHaste(),
			setValue: (eventID: EventID, target: Target, newValue: boolean) => {
				target.setParryHaste(eventID, newValue);
			},
		});
		new EnumPicker<Target>(section3, modTarget, {
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
			changedEvent: (target: Target) => target.levelChangeEmitter,
			getValue: (target: Target) => target.getSpellSchool(),
			setValue: (eventID: EventID, target: Target, newValue: number) => {
				target.setSpellSchool(eventID, newValue);
			},
		});
	}
}

function addEncounterFieldPickers(rootElem: HTMLElement, encounter: Encounter, showExecuteProportion: boolean) {
	new NumberPicker(rootElem, encounter, {
		label: 'Duration',
		changedEvent: (encounter: Encounter) => encounter.durationChangeEmitter,
		getValue: (encounter: Encounter) => encounter.getDuration(),
		setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
			encounter.setDuration(eventID, newValue);
		},
	});
	new NumberPicker(rootElem, encounter, {
		label: 'Duration +/-',
		changedEvent: (encounter: Encounter) => encounter.durationChangeEmitter,
		getValue: (encounter: Encounter) => encounter.getDurationVariation(),
		setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
			encounter.setDurationVariation(eventID, newValue);
		},
	});

	if (showExecuteProportion) {
		new NumberPicker(rootElem, encounter, {
			label: 'Execute Duration (%)',
			labelTooltip: 'Percentage of the total encounter duration, for which the targets will be considered to be in execute range (< 20% HP) for the purpose of effects like Warrior Execute or Mage Molten Fury.',
			changedEvent: (encounter: Encounter) => encounter.executeProportionChangeEmitter,
			getValue: (encounter: Encounter) => encounter.getExecuteProportion() * 100,
			setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
				encounter.setExecuteProportion(eventID, newValue / 100);
			},
		});
	}
}

const ALL_TARGET_STATS: Array<Stat> = [
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
