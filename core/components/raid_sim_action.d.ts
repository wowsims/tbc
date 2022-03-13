import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { ProgressMetrics } from '/tbc/core/proto/api.js';
import { SimResult } from '/tbc/core/proto_utils/sim_result.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
export declare function addRaidSimAction(simUI: SimUI): RaidSimResultsManager;
export declare type ReferenceData = {
    simResult: SimResult;
    settings: any;
    raidProto: RaidProto;
    encounterProto: EncounterProto;
};
export declare class RaidSimResultsManager {
    readonly currentChangeEmitter: TypedEvent<void>;
    readonly referenceChangeEmitter: TypedEvent<void>;
    readonly changeEmitter: TypedEvent<void>;
    private readonly simUI;
    private currentData;
    private referenceData;
    constructor(simUI: SimUI);
    setSimProgress(progress: ProgressMetrics): void;
    setSimResult(eventID: EventID, simResult: SimResult): void;
    private updateReference;
    getCurrentData(): ReferenceData | null;
    getReferenceData(): ReferenceData | null;
    static makeToplineResultsContent(simResult: SimResult, isIndividualSim: boolean): string;
}
