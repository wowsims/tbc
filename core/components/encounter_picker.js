import { MobType } from '/tbc/core/proto/common.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { NumberPicker } from '/tbc/core/components/number_picker.js';
import { Component } from './component.js';
export class EncounterPicker extends Component {
    constructor(parent, modEncounter, config) {
        super(parent, 'encounter-picker-root');
        new NumberPicker(this.rootElem, modEncounter, {
            label: 'Duration',
            changedEvent: (encounter) => encounter.durationChangeEmitter,
            getValue: (encounter) => encounter.getDuration(),
            setValue: (eventID, encounter, newValue) => {
                encounter.setDuration(eventID, newValue);
            },
        });
        if (config.showTargetArmor) {
            new NumberPicker(this.rootElem, modEncounter.primaryTarget, {
                label: 'Target Armor',
                changedEvent: (target) => target.armorChangeEmitter,
                getValue: (target) => target.getArmor(),
                setValue: (eventID, target, newValue) => {
                    target.setArmor(eventID, newValue);
                },
            });
        }
        new EnumPicker(this.rootElem, modEncounter.primaryTarget, MobTypePickerConfig);
        if (config.showNumTargets) {
            new NumberPicker(this.rootElem, modEncounter, {
                label: '# of Targets',
                changedEvent: (encounter) => encounter.numTargetsChangeEmitter,
                getValue: (encounter) => encounter.getNumTargets(),
                setValue: (eventID, encounter, newValue) => {
                    encounter.setNumTargets(eventID, newValue);
                },
            });
        }
    }
}
export const MobTypePickerConfig = {
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
    changedEvent: (target) => target.mobTypeChangeEmitter,
    getValue: (target) => target.getMobType(),
    setValue: (eventID, target, newValue) => {
        target.setMobType(eventID, newValue);
    },
};
