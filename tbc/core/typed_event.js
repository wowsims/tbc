/** Provides a type-safe event interface. */
export class TypedEvent {
    constructor() {
        this._listeners = [];
    }
    on(listener) {
        this._listeners.push(listener);
        return {
            dispose: () => this.off(listener),
        };
    }
    off(listener) {
        const idx = this._listeners.indexOf(listener);
        if (idx != -1) {
            this._listeners.splice(idx, 1);
        }
    }
    emit(event) {
        this._listeners.forEach(listener => listener(event));
    }
    once(listener) {
        const onceListener = (event) => {
            this.off(onceListener);
            listener(event);
        };
        return this.on(onceListener);
    }
}
