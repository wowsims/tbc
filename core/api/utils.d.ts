import { Class } from './common.js';
import { Enchant } from './common.js';
import { Gem } from './common.js';
import { GemColor } from './common.js';
import { ItemSlot } from './common.js';
import { Item } from './common.js';
import { Race } from './common.js';
import { Spec } from './common.js';
import { BalanceDruid, BalanceDruid_Agent as BalanceDruidAgent, DruidTalents, BalanceDruid_Options as BalanceDruidOptions } from './druid.js';
import { ElementalShaman, ElementalShaman_Agent as ElementalShamanAgent, ShamanTalents, ElementalShaman_Options as ElementalShamanOptions } from './shaman.js';
export declare type ShamanSpecs = Spec.SpecElementalShaman;
export declare type AgentUnion = BalanceDruidAgent | ElementalShamanAgent;
export declare type SpecAgent<T extends Spec> = T extends Spec.SpecBalanceDruid ? BalanceDruidAgent : ElementalShamanAgent;
export declare type TalentsUnion = DruidTalents | ShamanTalents;
export declare type SpecTalents<T extends Spec> = T extends Spec.SpecBalanceDruid ? DruidTalents : ShamanTalents;
export declare type SpecOptionsUnion = BalanceDruidOptions | ElementalShamanOptions;
export declare type SpecOptions<T extends Spec> = T extends Spec.SpecBalanceDruid ? BalanceDruidOptions : ElementalShamanOptions;
export declare type SpecProtoUnion = BalanceDruid | ElementalShaman;
export declare type SpecProto<T extends Spec> = T extends Spec.SpecBalanceDruid ? BalanceDruid : ElementalShaman;
export declare type SpecTypeFunctions<SpecType extends Spec> = {
    agentCreate: () => SpecAgent<SpecType>;
    agentEquals: (a: SpecAgent<SpecType>, b: SpecAgent<SpecType>) => boolean;
    agentCopy: (a: SpecAgent<SpecType>) => SpecAgent<SpecType>;
    agentToJson: (a: SpecAgent<SpecType>) => any;
    agentFromJson: (obj: any) => SpecAgent<SpecType>;
    talentsCreate: () => SpecTalents<SpecType>;
    talentsEquals: (a: SpecTalents<SpecType>, b: SpecTalents<SpecType>) => boolean;
    talentsCopy: (a: SpecTalents<SpecType>) => SpecTalents<SpecType>;
    talentsToJson: (a: SpecTalents<SpecType>) => any;
    talentsFromJson: (obj: any) => SpecTalents<SpecType>;
    optionsCreate: () => SpecOptions<SpecType>;
    optionsEquals: (a: SpecOptions<SpecType>, b: SpecOptions<SpecType>) => boolean;
    optionsCopy: (a: SpecOptions<SpecType>) => SpecOptions<SpecType>;
    optionsToJson: (a: SpecOptions<SpecType>) => any;
    optionsFromJson: (obj: any) => SpecOptions<SpecType>;
};
export declare const specTypeFunctions: Partial<Record<Spec, SpecTypeFunctions<any>>>;
export declare const specToClass: Record<Spec, Class>;
export declare const specToEligibleRaces: Record<Spec, Array<Race>>;
export declare function getEligibleItemSlots(item: Item): Array<ItemSlot>;
/**
 * Returns all item slots to which the enchant might be applied.
 *
 * Note that this alone is not enough; some items have further restrictions,
 * e.g. some weapon enchants may only be applied to 2H weapons.
 */
export declare function getEligibleEnchantSlots(enchant: Enchant): Array<ItemSlot>;
export declare function enchantAppliesToItem(enchant: Enchant, item: Item): boolean;
export declare function gemMatchesSocket(gem: Gem, socketColor: GemColor): boolean;
export declare function gemEligibleForSocket(gem: Gem, socketColor: GemColor): boolean;
