import { MobType } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/encounter.js';
import { Target } from '/tbc/core/target.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { EnumPicker, EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';

import { Component } from './component.js';

export interface EncounterPickerConfig {
	showTargetArmor: boolean;
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

		new EnumPicker<Target>(this.rootElem, modEncounter.primaryTarget, {
			label: 'Target Level',
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
		if (config.showTargetArmor) {
			new NumberPicker(this.rootElem, modEncounter.primaryTarget, {
				label: 'Target Armor',
				changedEvent: (target: Target) => target.statsChangeEmitter,
				getValue: (target: Target) => target.getStats().getStat(Stat.StatArmor),
				setValue: (eventID: EventID, target: Target, newValue: number) => {
					target.setStats(eventID, target.getStats().withStat(Stat.StatArmor, newValue));
				},
			});
		}

		new EnumPicker(this.rootElem, modEncounter.primaryTarget, MobTypePickerConfig);

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
	}
}

export const MobTypePickerConfig: EnumPickerConfig<Target> = {
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
	changedEvent: (target: Target) => target.mobTypeChangeEmitter,
	getValue: (target: Target) => target.getMobType(),
	setValue: (eventID: EventID, target: Target, newValue: number) => {
		target.setMobType(eventID, newValue);
	},
};
