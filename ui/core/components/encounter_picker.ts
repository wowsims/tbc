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

import { Component } from './component.js';

export interface EncounterPickerConfig {
	simpleTargetStats?: Array<Stat>;
	showNumTargets: boolean;
	showExecuteProportion: boolean;
}

export class EncounterPicker extends Component {
	constructor(parent: HTMLElement, modEncounter: Encounter, config: EncounterPickerConfig) {
		super(parent, 'encounter-picker-root');

		new NumberPicker(this.rootElem, modEncounter, {
			label: 'Duration',
			changedEvent: (encounter: Encounter) => encounter.durationChangeEmitter,
			getValue: (encounter: Encounter) => encounter.getDuration(),
			setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
				encounter.setDuration(eventID, newValue);
			},
		});
		new NumberPicker(this.rootElem, modEncounter, {
			label: 'Duration +/-',
			changedEvent: (encounter: Encounter) => encounter.durationChangeEmitter,
			getValue: (encounter: Encounter) => encounter.getDurationVariation(),
			setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
				encounter.setDurationVariation(eventID, newValue);
			},
		});

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

		new EnumPicker(this.rootElem, modEncounter, MobTypePickerConfig);

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
				changedEvent: (encounter: Encounter) => encounter.numTargetsChangeEmitter,
				getValue: (encounter: Encounter) => encounter.getNumTargets(),
				setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
					encounter.setNumTargets(eventID, newValue);
				},
			});
		}

		if (config.showExecuteProportion) {
			new NumberPicker(this.rootElem, modEncounter, {
				label: 'Execute Duration (%)',
				labelTooltip: 'Percentage of the total encounter duration, for which the targets will be considered to be in execute range (< 20% HP) for the purpose of effects like Warrior Execute or Mage Molten Fury.',
				changedEvent: (encounter: Encounter) => encounter.executeProportionChangeEmitter,
				getValue: (encounter: Encounter) => encounter.getExecuteProportion() * 100,
				setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
					encounter.setExecuteProportion(eventID, newValue / 100);
				},
			});
		}
		
		// Simple/Custom/Preset [Edit Button]
		const advancedButton = document.createElement('button');
		advancedButton.classList.add('sim-button', 'advanced-button');
		advancedButton.textContent = 'ADVANCED';
		this.rootElem.appendChild(advancedButton);
	}
}

class TargetPicker extends Component {
	constructor(parent: HTMLElement, modTarget: Target) {
		super(parent, 'target-picker-root');

		new EnumPicker<Target>(this.rootElem, modTarget, {
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

		new EnumPicker(this.rootElem, modTarget, MobTypePickerConfig);

		ALL_TARGET_STATS.forEach(stat => {
			new NumberPicker(this.rootElem, modTarget, {
				label: statNames[stat],
				changedEvent: (target: Target) => target.statsChangeEmitter,
				getValue: (target: Target) => target.getStats().getStat(stat),
				setValue: (eventID: EventID, target: Target, newValue: number) => {
					target.setStats(eventID, target.getStats().withStat(stat, newValue));
				},
			});
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
];

export const MobTypePickerConfig: EnumPickerConfig<Encounter> = {
	label: 'Mob Type',
	values: [
		{ name: 'None', value: MobType.MobTypeUnknown },
		{ name: 'Beast', value: MobType.MobTypeBeast },
		{ name: 'Demon', value: MobType.MobTypeDemon },
		{ name: 'Dragonkin', value: MobType.MobTypeDragonkin },
		{ name: 'Elemental', value: MobType.MobTypeElemental },
		{ name: 'Giant', value: MobType.MobTypeGiant },
		{ name: 'Humanoid', value: MobType.MobTypeHumanoid },
		{ name: 'Mechanical', value: MobType.MobTypeMechanical },
		{ name: 'Undead', value: MobType.MobTypeUndead },
	],
	changedEvent: (encounter: Encounter) => encounter.changeEmitter,
	getValue: (encounter: Encounter) => encounter.primaryTarget.getMobType(),
	setValue: (eventID: EventID, encounter: Encounter, newValue: number) => {
		encounter.primaryTarget.setMobType(eventID, newValue);
	},
};
