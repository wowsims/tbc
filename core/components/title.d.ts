import { Component } from '/tbc/core/components/component.js';
import { Spec } from '/tbc/core/proto/common.js';
export interface SimLinkOption {
    iconUrl: string;
    href: string;
    text: string;
    color: string;
}
export declare class Title extends Component {
    private readonly buttonElem;
    constructor(parent: HTMLElement, currentSpec: Spec | null);
    private makeOptionData;
    private makeOption;
    static makeOptionElem(data: SimLinkOption): HTMLElement;
}
