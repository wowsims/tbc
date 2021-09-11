export interface Disposable {
  dispose(): void;
}

export interface Listener<T> {
  (event: T): any;
}

/** Provides a type-safe event interface. */
export class TypedEvent<T> {
  private _listeners: Listener<T>[] = [];

  on(listener: Listener<T>): Disposable {
    this._listeners.push(listener);
    return {
      dispose: () => this.off(listener),
    };
  }

  off(listener: Listener<T>) {
    const idx = this._listeners.indexOf(listener);
    if (idx != -1) {
      this._listeners.splice(idx, 1);
    }
  }

  emit(event: T) {
    this._listeners.forEach(listener => listener(event));
  }

  once(listener: Listener<T>): Disposable {
    const onceListener = (event: T) => {
      this.off(onceListener);
      listener(event);
    };

    return this.on(onceListener);
  }
}
