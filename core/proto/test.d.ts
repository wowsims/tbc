import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.CharacterStatsTestResult
 */
export interface CharacterStatsTestResult {
    /**
     * @generated from protobuf field: repeated double final_stats = 1;
     */
    finalStats: number[];
}
/**
 * @generated from protobuf message proto.StatWeightsTestResult
 */
export interface StatWeightsTestResult {
    /**
     * @generated from protobuf field: repeated double weights = 1;
     */
    weights: number[];
}
/**
 * @generated from protobuf message proto.DpsTestResult
 */
export interface DpsTestResult {
    /**
     * @generated from protobuf field: double dps = 1;
     */
    dps: number;
    /**
     * @generated from protobuf field: double tps = 2;
     */
    tps: number;
    /**
     * @generated from protobuf field: double dtps = 3;
     */
    dtps: number;
}
/**
 * @generated from protobuf message proto.TestSuiteResult
 */
export interface TestSuiteResult {
    /**
     * Maps test names to their results.
     *
     * @generated from protobuf field: map<string, proto.CharacterStatsTestResult> character_stats_results = 2;
     */
    characterStatsResults: {
        [key: string]: CharacterStatsTestResult;
    };
    /**
     * Maps test names to their results.
     *
     * @generated from protobuf field: map<string, proto.StatWeightsTestResult> stat_weights_results = 3;
     */
    statWeightsResults: {
        [key: string]: StatWeightsTestResult;
    };
    /**
     * Maps test names to their results.
     *
     * @generated from protobuf field: map<string, proto.DpsTestResult> dps_results = 1;
     */
    dpsResults: {
        [key: string]: DpsTestResult;
    };
}
declare class CharacterStatsTestResult$Type extends MessageType<CharacterStatsTestResult> {
    constructor();
    create(value?: PartialMessage<CharacterStatsTestResult>): CharacterStatsTestResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CharacterStatsTestResult): CharacterStatsTestResult;
    internalBinaryWrite(message: CharacterStatsTestResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.CharacterStatsTestResult
 */
export declare const CharacterStatsTestResult: CharacterStatsTestResult$Type;
declare class StatWeightsTestResult$Type extends MessageType<StatWeightsTestResult> {
    constructor();
    create(value?: PartialMessage<StatWeightsTestResult>): StatWeightsTestResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StatWeightsTestResult): StatWeightsTestResult;
    internalBinaryWrite(message: StatWeightsTestResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.StatWeightsTestResult
 */
export declare const StatWeightsTestResult: StatWeightsTestResult$Type;
declare class DpsTestResult$Type extends MessageType<DpsTestResult> {
    constructor();
    create(value?: PartialMessage<DpsTestResult>): DpsTestResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DpsTestResult): DpsTestResult;
    internalBinaryWrite(message: DpsTestResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.DpsTestResult
 */
export declare const DpsTestResult: DpsTestResult$Type;
declare class TestSuiteResult$Type extends MessageType<TestSuiteResult> {
    constructor();
    create(value?: PartialMessage<TestSuiteResult>): TestSuiteResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TestSuiteResult): TestSuiteResult;
    private binaryReadMap2;
    private binaryReadMap3;
    private binaryReadMap1;
    internalBinaryWrite(message: TestSuiteResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.TestSuiteResult
 */
export declare const TestSuiteResult: TestSuiteResult$Type;
export {};
