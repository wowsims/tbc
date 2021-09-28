import { Gear } from '../api/gear.js';
import { Buffs } from '../api/common.js';
import { Consumes } from '../api/common.js';
import { Encounter } from '../api/common.js';
import { EquipmentSpec } from '../api/common.js';
import { Race } from '../api/common.js';
import { Spec } from '../api/common.js';
import { Stat } from '../api/common.js';
import { Stats } from '../api/stats.js';
import { EnumPickerConfig } from '../components/enum_picker.js';
import { IconInput } from '../components/icon_picker.js';
import { NumberPickerConfig } from '../components/number_picker.js';
import { SimUI, SimUIConfig } from '../sim_ui.js';
export interface DefaultThemeConfig<SpecType extends Spec> extends SimUIConfig<SpecType> {
    displayStats: Array<Stat>;
    iconSections: Record<string, Array<IconInput>>;
    otherSections: Record<string, Array<{
        type: 'number';
        cssClass: string;
        config: NumberPickerConfig;
    } | {
        type: 'enum';
        cssClass: string;
        config: EnumPickerConfig;
    }>>;
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
