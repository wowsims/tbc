import { MobType } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/encounter.js';
import { Target } from '/tbc/core/target.js';
import { EnumPicker, EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';

import { Component } from './component.js';

export interface EncounterPickerConfig {
  showTargetArmor: boolean;
  showNumTargets: boolean;
}

export class EncounterPicker extends Component {
  constructor(parent: HTMLElement, modEncounter: Encounter, config: EncounterPickerConfig) {
    super(parent, 'encounter-picker-root');

    new NumberPicker(this.rootElem, modEncounter, {
      label: 'Duration',
      changedEvent: (encounter: Encounter) => encounter.durationChangeEmitter,
      getValue: (encounter: Encounter) => encounter.getDuration(),
      setValue: (encounter: Encounter, newValue: number) => {
				encounter.setDuration(newValue);
      },
    });

    if (config.showTargetArmor) {
      new NumberPicker(this.rootElem, modEncounter.primaryTarget, {
        label: 'Target Armor',
        changedEvent: (target: Target) => target.armorChangeEmitter,
        getValue: (target: Target) => target.getArmor(),
        setValue: (target: Target, newValue: number) => {
					target.setArmor(newValue);
        },
      });
    }

		new EnumPicker(this.rootElem, modEncounter.primaryTarget, MobTypePickerConfig);

    if (config.showNumTargets) {
      new NumberPicker(this.rootElem, modEncounter, {
        label: '# of Targets',
        changedEvent: (encounter: Encounter) => encounter.numTargetsChangeEmitter,
        getValue: (encounter: Encounter) => encounter.getNumTargets(),
        setValue: (encounter: Encounter, newValue: number) => {
					encounter.setNumTargets(newValue);
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
	setValue: (target: Target, newValue: number) => {
		target.setMobType(newValue);
	},
};
