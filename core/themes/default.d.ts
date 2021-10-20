import { EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { IconInput } from '/tbc/core/components/icon_picker.js';
import { NumberPickerConfig } from '/tbc/core/components/number_picker.js';
import { Buffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { SimUI, SimUIConfig } from '/tbc/core/sim_ui.js';
export interface IconSection {
    tooltip?: string;
    icons: Array<IconInput>;
}
export interface InputSection {
    tooltip?: string;
    inputs: Array<{
        type: 'number';
        cssClass: string;
        config: NumberPickerConfig;
    } | {
        type: 'enum';
        cssClass: string;
        config: EnumPickerConfig;
    }>;
}
export interface DefaultThemeConfig<SpecType extends Spec> extends SimUIConfig<SpecType> {
    displayStats: Array<Stat>;
    selfBuffInputs: IconSection;
    buffInputs: IconSection;
    debuffInputs: IconSection;
    consumeInputs: IconSection;
    rotationInputs: InputSection;
    otherInputs?: InputSection;
    additionalSections?: Record<string, InputSection>;
    showTargetArmor: boolean;
    showNumTargets: boolean;
    freezeTalents: boolean;
    presets: {
        gear: Array<{
            name: string;
            tooltip?: string;
            equipment: EquipmentSpec;
        }>;
        encounters: Array<{
            name: string;
            tooltip?: string;
            encounter: Encounter;
        }>;
        talents: Array<{
            name: string;
            tooltip?: string;
            talents: string;
        }>;
    };
}
export interface GearAndStats {
    gear: Gear;
    customStats: Stats;
}
export interface Settings {
    buffs: Buffs;
    consumes: Consumes;
    race: Race;
}
export declare class DefaultTheme<SpecType extends Spec> extends SimUI<SpecType> {
    private readonly _config;
    constructor(parentElem: HTMLElement, config: DefaultThemeConfig<SpecType>);
    init(): Promise<void>;
}
