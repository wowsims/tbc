import { ItemOrSpellId } from '/tbc/core/resources.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { ExclusivityTag } from '/tbc/core/sim_ui.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { Component } from './component.js';
export declare class IconPicker<ModObject> extends Component {
    private readonly _input;
    private readonly _modObject;
    private readonly _rootAnchor;
    private readonly _improvedAnchor;
    private readonly _counterElem;
    private readonly _clickedEmitter;
    constructor(parent: HTMLElement, modObj: ModObject, input: IconInput<ModObject>, simUI: SimUI<any>);
    private getValue;
    private setValue;
    private updateIcon;
}
export declare type IconInput<ModObject> = {
    id: ItemOrSpellId;
    states: number;
    improvedId?: ItemOrSpellId;
    exclusivityTags?: Array<ExclusivityTag>;
    changedEvent: (obj: ModObject) => TypedEvent<any>;
    getValue: (obj: ModObject) => boolean | number;
    setBooleanValue?: (obj: ModObject, newValue: boolean) => void;
    setNumberValue?: (obj: ModObject, newValue: number) => void;
};
