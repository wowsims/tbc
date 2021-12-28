export declare type EventID = number;
export interface Disposable {
    dispose(): void;
}
export interface Listener<T> {
    (eventID: EventID, event: T): any;
}
/** Provides a type-safe event interface. */
export declare class TypedEvent<T> {
    private label;
    constructor(label?: string);
    private listeners;
    private firedEvents;
    private frozenEvents;
    on(listener: Listener<T>): Disposable;
    off(listener: Listener<T>): void;
    once(listener: Listener<T>): Disposable;
    emit(eventID: EventID, event: T): void;
    private fireEventInternal;
    static freezeAllAndDo(func: () => void): void;
    static nextEventID(): EventID;
    static onAny(events: Array<TypedEvent<any>>, label?: string): TypedEvent<void>;
}
