import { Encounter } from '/tbc/core/encounter.js';
import { Target } from '/tbc/core/target.js';
import { EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { Component } from './component.js';
export interface EncounterPickerConfig {
    showTargetArmor: boolean;
    showNumTargets: boolean;
    showExecuteProportion: boolean;
}
export declare class EncounterPicker extends Component {
    constructor(parent: HTMLElement, modEncounter: Encounter, config: EncounterPickerConfig);
}
export declare const MobTypePickerConfig: EnumPickerConfig<Target>;
