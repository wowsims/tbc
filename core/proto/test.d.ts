import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.DpsTestResult
 */
export interface DpsTestResult {
    /**
     * @generated from protobuf field: double dps = 2;
     */
    dps: number;
}
/**
 * @generated from protobuf message proto.TestSuiteResult
 */
export interface TestSuiteResult {
    /**
     * Maps test names to their results.
     *
     * @generated from protobuf field: map<string, proto.DpsTestResult> dps_results = 1;
     */
    dpsResults: {
        [key: string]: DpsTestResult;
    };
}
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
    private binaryReadMap1;
    internalBinaryWrite(message: TestSuiteResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.TestSuiteResult
 */
export declare const TestSuiteResult: TestSuiteResult$Type;
export {};
