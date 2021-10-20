import { ItemOrSpellId } from '/tbc/core/resources.js';
import { Sim } from '/tbc/core/sim.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { ExclusivityTag } from '/tbc/core/sim_ui.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { Component } from './component.js';
export declare class IconPicker extends Component {
    private readonly _inputs;
    constructor(parent: HTMLElement, sim: Sim<any>, inputs: Array<IconInput>, simUI: SimUI<any>);
}
export declare type IconInput = {
    id: ItemOrSpellId;
    states: number;
    improvedId?: ItemOrSpellId;
    exclusivityTags?: Array<ExclusivityTag>;
    changedEvent: (sim: Sim<any>) => TypedEvent<any>;
    getValue: (sim: Sim<any>) => boolean | number;
    setBooleanValue?: (sim: Sim<any>, newValue: boolean) => void;
    setNumberValue?: (sim: Sim<any>, newValue: number) => void;
};
