import { Spec } from '../proto/common.js';
import { SpecTalents } from '../api/utils.js';
import { Sim } from '../sim.js';
import { Component } from '../components/component.js';
export declare abstract class TalentsPicker<SpecType extends Spec> extends Component {
    private readonly sim;
    frozen: boolean;
    readonly trees: Array<TalentTreePicker<SpecType>>;
    constructor(parent: HTMLElement, sim: Sim<SpecType>, treeConfigs: Array<TalentTreeConfig<SpecType>>);
    get numPoints(): number;
    isFull(): boolean;
    update(): void;
    getTalents(): SpecTalents<SpecType>;
    getTalentsString(): string;
    setTalentsString(str: string): void;
    freeze(): void;
}
declare class TalentTreePicker<SpecType extends Spec> extends Component {
    private readonly config;
    private readonly title;
    readonly talents: Array<TalentPicker<SpecType>>;
    readonly picker: TalentsPicker<SpecType>;
    numPoints: number;
    constructor(parent: HTMLElement, sim: Sim<SpecType>, config: TalentTreeConfig<SpecType>, picker: TalentsPicker<SpecType>);
    update(): void;
    getTalent(location: TalentLocation): TalentPicker<SpecType>;
    getTalentsString(): string;
    setTalentsString(str: string): void;
}
declare class TalentPicker<SpecType extends Spec> extends Component {
    readonly config: TalentConfig<SpecType>;
    private readonly tree;
    private readonly pointsDisplay;
    constructor(parent: HTMLElement, sim: Sim<SpecType>, config: TalentConfig<SpecType>, tree: TalentTreePicker<SpecType>);
    getRow(): number;
    getCol(): number;
    getPoints(): number;
    isFull(): boolean;
    canSetPoints(newPoints: number): boolean;
    setPoints(newPoints: number, checkValidity: boolean): void;
    getSpellIdForPoints(numPoints: number): number;
    update(): void;
}
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
export {};
