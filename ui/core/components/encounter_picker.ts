import { EncounterType } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/encounter.js';
import { Target } from '/tbc/core/target.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { EnumPicker, EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { ListPicker, ListPickerConfig } from '/tbc/core/components/list_picker.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { statNames } from '/tbc/core/proto_utils/names.js';
import { getEnumValues } from '/tbc/core/utils.js';

import { Component } from './component.js';
import { Popup } from './popup.js';

export interface EncounterPickerConfig {
	simpleTargetStats?: Array<Stat>;
	showNumTargets: boolean;
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

		if (config.showNumTargets) {
			new NumberPicker(this.rootElem, modEncounter, {
				label: '# of Targets',
				changedEvent: (encounter: Encounter) => encounter.changeEmitter,
				getValue: (encounter: Encounter) => encounter.getNumTargets(),
				setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
					encounter.setNumTargets(eventID, newValue);
				},
				enableWhen: (encounter: Encounter) => encounter.getType() == EncounterType.EncounterTypeSimple,
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

		const encounterTypeContainer = this.rootElem.getElementsByClassName('encounter-type')[0] as HTMLElement;
		new EnumPicker<Encounter>(encounterTypeContainer, this.encounter, {
			label: 'ENCOUNTER',
			values: [
				{ name: 'Simple', value: EncounterType.EncounterTypeSimple },
				{ name: 'Custom', value: EncounterType.EncounterTypeCustom },
			].concat((getEnumValues(EncounterType) as Array<EncounterType>).filter(val => ![EncounterType.EncounterTypeSimple, EncounterType.EncounterTypeCustom].includes(val)).map((val, i) => {
				return {
					name: '',
					value: val,
				};
			})),
			changedEvent: (encounter: Encounter) => encounter.typeChangeEmitter,
			getValue: (encounter: Encounter) => encounter.getType(),
			setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
				encounter.setType(eventID, newValue);
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
