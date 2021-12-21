export declare type EventID = number;
export interface Disposable {
    dispose(): void;
}
export interface Listener<T> {
    (eventID: EventID, event: T): any;
}
/** Provides a type-safe event interface. */
export declare class TypedEvent<T> {
    private listeners;
    private firedEvents;
    private frozenEvents;
    on(listener: Listener<T>): Disposable;
    off(listener: Listener<T>): void;
    once(listener: Listener<T>): Disposable;
    emit(eventID: EventID, event: T): void;
    private fireEventInternal;
    static freezeAll(): void;
    static unfreezeAll(): void;
    static nextEventID(): EventID;
}
