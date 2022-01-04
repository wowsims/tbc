import { Component } from '/tbc/core/components/component.js';
import { Player } from '/tbc/core/player.js';
export declare class CooldownsPicker extends Component {
    readonly player: Player<any>;
    private cooldownPickers;
    constructor(parentElem: HTMLElement, player: Player<any>);
    private update;
    private makeActionPicker;
    private makeTimingsPicker;
}
