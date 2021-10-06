import { Gear } from '../api/gear.js';
import { Buffs } from '../proto/common.js';
import { Consumes } from '../proto/common.js';
import { Encounter } from '../proto/common.js';
import { EquipmentSpec } from '../proto/common.js';
import { Race } from '../proto/common.js';
import { Spec } from '../proto/common.js';
import { Stat } from '../proto/common.js';
import { Stats } from '../api/stats.js';
import { EnumPickerConfig } from '../components/enum_picker.js';
import { IconInput } from '../components/icon_picker.js';
import { NumberPickerConfig } from '../components/number_picker.js';
import { SimUI, SimUIConfig } from '../sim_ui.js';
export interface DefaultThemeConfig<SpecType extends Spec> extends SimUIConfig<SpecType> {
    displayStats: Array<Stat>;
    iconSections: Record<string, {
        tooltip?: string;
        icons: Array<IconInput>;
    }>;
    otherSections: Record<string, {
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
    }>;
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
