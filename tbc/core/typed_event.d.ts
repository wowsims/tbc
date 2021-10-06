export interface Disposable {
    dispose(): void;
}
export interface Listener<T> {
    (event: T): any;
}
/** Provides a type-safe event interface. */
export declare class TypedEvent<T> {
    private _listeners;
    on(listener: Listener<T>): Disposable;
    off(listener: Listener<T>): void;
    emit(event: T): void;
    once(listener: Listener<T>): Disposable;
}
