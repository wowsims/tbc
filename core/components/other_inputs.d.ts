import { Sim } from '../sim.js';
export declare const ShadowPriestDPS: {
    type: "number";
    cssClass: string;
    config: {
        label: string;
        changedEvent: (sim: Sim<any>) => import("../typed_event.js").TypedEvent<void>;
        getValue: (sim: Sim<any>) => number;
        setValue: (sim: Sim<any>, newValue: number) => void;
    };
};
