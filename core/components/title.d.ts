import { Component } from '/tbc/core/components/component.js';
import { Spec } from '/tbc/core/proto/common.js';
export declare const titleIcons: Record<Spec, string>;
export declare const raidSimIcon: string;
export interface SimLinkOption {
    iconUrl: string;
    href: string;
    text: string;
    color: string;
}
export declare class Title extends Component {
    private readonly buttonElem;
    private readonly dropdownElem;
    constructor(parent: HTMLElement, currentSpec: Spec | null);
    private makeOptionData;
    private makeOption;
    static makeOptionElem(data: SimLinkOption): HTMLElement;
}
