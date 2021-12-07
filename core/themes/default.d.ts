import { Party } from '/tbc/core/party.js';
import { Player } from '/tbc/core/player.js';
import { Raid } from '/tbc/core/raid.js';
import { Target } from '/tbc/core/target.js';
import { BooleanPickerConfig } from '/tbc/core/components/boolean_picker.js';
import { EncounterPickerConfig } from '/tbc/core/components/encounter_picker.js';
import { EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { IconInput } from '/tbc/core/components/icon_picker.js';
import { NumberPickerConfig } from '/tbc/core/components/number_picker.js';
import { SavedDataConfig } from '/tbc/core/components/saved_data_manager.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { SimUI, SimUIConfig } from '/tbc/core/sim_ui.js';
export interface InputSection {
    tooltip?: string;
    inputs: Array<{
        type: 'boolean';
        cssClass: string;
        getModObject: (simUI: SimUI<any>) => any;
        config: BooleanPickerConfig<any>;
    } | {
        type: 'number';
        cssClass: string;
        getModObject: (simUI: SimUI<any>) => any;
        config: NumberPickerConfig<any>;
    } | {
        type: 'enum';
        cssClass: string;
        getModObject: (simUI: SimUI<any>) => any;
        config: EnumPickerConfig<any>;
    }>;
}
export interface DefaultThemeConfig<SpecType extends Spec> extends SimUIConfig<SpecType> {
    epStats: Array<Stat>;
    epReferenceStat: Stat;
    displayStats: Array<Stat>;
    selfBuffInputs: Array<IconInput<Player<any>>>;
    raidBuffInputs: Array<IconInput<Raid>>;
    partyBuffInputs: Array<IconInput<Party>>;
    playerBuffInputs: Array<IconInput<Player<any>>>;
    debuffInputs: Array<IconInput<Target>>;
    consumeInputs: Array<IconInput<Player<any>>>;
    rotationInputs: InputSection;
    otherInputs?: InputSection;
    additionalSections?: Record<string, InputSection>;
    encounterPicker: EncounterPickerConfig;
    freezeTalents?: boolean;
    presets: {
        gear: Array<PresetGear>;
        talents: Array<SavedDataConfig<Player<any>, string>>;
    };
}
export interface GearAndStats {
    gear: Gear;
    bonusStats?: Stats;
}
export interface PresetGear {
    name: string;
    gear: EquipmentSpec;
    tooltip?: string;
    enableWhen?: (obj: Player<any>) => boolean;
}
export interface Settings {
    raidBuffs: RaidBuffs;
    partyBuffs: PartyBuffs;
    individualBuffs: IndividualBuffs;
    consumes: Consumes;
    race: Race;
}
export declare class DefaultTheme<SpecType extends Spec> extends SimUI<SpecType> {
    private readonly _config;
    constructor(parentElem: HTMLElement, config: DefaultThemeConfig<SpecType>);
    init(): Promise<void>;
}
