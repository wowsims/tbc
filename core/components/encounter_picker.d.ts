import { Stat } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/encounter.js';
import { Target } from '/tbc/core/target.js';
import { EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { Component } from './component.js';
export interface EncounterPickerConfig {
    simpleTargetStats?: Array<Stat>;
    showNumTargets: boolean;
    showExecuteProportion: boolean;
}
export declare class EncounterPicker extends Component {
    constructor(parent: HTMLElement, modEncounter: Encounter, config: EncounterPickerConfig);
}
export declare const MobTypePickerConfig: EnumPickerConfig<Target>;
