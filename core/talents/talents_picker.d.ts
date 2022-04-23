import { Component } from '/tbc/core/components/component.js';
import { Spec } from '/tbc/core/proto/common.js';
import { SpecTalents } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';
import { EventID } from '/tbc/core/typed_event.js';
export declare abstract class TalentsPicker<SpecType extends Spec> extends Component {
    private readonly player;
    frozen: boolean;
    readonly trees: Array<TalentTreePicker<SpecType>>;
    constructor(parent: HTMLElement, player: Player<SpecType>, treeConfigs: TalentsConfig<SpecType>);
    get numPoints(): number;
    isFull(): boolean;
    update(eventID: EventID): void;
    getTalentsString(): string;
    setTalentsString(eventID: EventID, str: string): void;
    freeze(): void;
}
declare class TalentTreePicker<SpecType extends Spec> extends Component {
    private readonly config;
    private readonly title;
    readonly talents: Array<TalentPicker<SpecType>>;
    readonly picker: TalentsPicker<SpecType>;
    numPoints: number;
    constructor(parent: HTMLElement, player: Player<SpecType>, config: TalentTreeConfig<SpecType>, picker: TalentsPicker<SpecType>);
    update(): void;
    getTalent(location: TalentLocation): TalentPicker<SpecType>;
    getTalentsString(): string;
    setTalentsString(str: string): void;
}
declare class TalentPicker<SpecType extends Spec> extends Component {
    readonly config: TalentConfig<SpecType>;
    private readonly tree;
    private readonly pointsDisplay;
    private longTouchTimer?;
    constructor(parent: HTMLElement, player: Player<SpecType>, config: TalentConfig<SpecType>, tree: TalentTreePicker<SpecType>);
    getRow(): number;
    getCol(): number;
    getPoints(): number;
    isFull(): boolean;
    canSetPoints(newPoints: number): boolean;
    setPoints(newPoints: number, checkValidity: boolean): void;
    getSpellIdForPoints(numPoints: number): number;
    update(): void;
}
export declare type TalentsConfig<SpecType extends Spec> = Array<TalentTreeConfig<SpecType>>;
export declare type TalentTreeConfig<SpecType extends Spec> = {
    name: string;
    backgroundUrl: string;
    talents: Array<TalentConfig<SpecType>>;
};
export declare type TalentLocation = {
    rowIdx: number;
    colIdx: number;
};
export declare type TalentConfig<SpecType extends Spec> = {
    fieldName?: keyof SpecTalents<SpecType>;
    location: TalentLocation;
    prereqLocation?: TalentLocation;
    prereqOfLocation?: TalentLocation;
    spellIds: Array<number>;
    maxPoints: number;
};
export declare function newTalentsConfig<SpecType extends Spec>(talents: TalentsConfig<SpecType>): TalentsConfig<SpecType>;
export {};
