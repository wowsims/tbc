// An event ID uniquely identifies a single event that occurred, usually due to
// some user action like changing a piece of gear.
//
// Event IDs allow us to make sure that hierarchies of TypedEvents fire only once,
// for a given event. This is very important for certain features, like undo/redo.
export type EventID = number;

export interface Disposable {
  dispose(): void;
}

export interface Listener<T> {
  (eventID: EventID, event: T): any;
}

interface FrozenEventData<T> {
	eventID: EventID,
	event: T,
}

/** Provides a type-safe event interface. */
export class TypedEvent<T> {
  private listeners: Array<Listener<T>> = [];

	// The IDs of events which have already been fired from this TypedEvent.
	private firedEvents: Array<EventID> = [];

	// IDs for currently frozen events pending on this TypedEvent. See freezeAll()
	// for more details.
	private frozenEvents: Array<FrozenEventData<T>> = [];

	// Registers a new listener to this event.
  on(listener: Listener<T>): Disposable {
    this.listeners.push(listener);
    return {
      dispose: () => this.off(listener),
    };
  }

	// Removes a listener from this event.
  off(listener: Listener<T>) {
    const idx = this.listeners.indexOf(listener);
    if (idx != -1) {
      this.listeners.splice(idx, 1);
    }
  }

	// Convenience for on() which calls off() autmatically after firing once.
  once(listener: Listener<T>): Disposable {
    const onceListener = (eventID: EventID, event: T) => {
      this.off(onceListener);
      listener(eventID, event);
    };

    return this.on(onceListener);
  }

  emit(eventID: EventID, event: T) {
		if (this.firedEvents.includes(eventID)) {
			return;
		}
		this.firedEvents.push(eventID);

		if (freezeCount > 0) {
			if (this.frozenEvents.length == 0) {
				frozenTypedEvents.push(this);
			}
			this.frozenEvents.push({
				eventID: eventID,
				event: event,
			});
		} else {
			this.fireEventInternal(eventID, event);
		}
  }

	private fireEventInternal(eventID: EventID, event: T) {
		this.listeners.forEach(listener => listener(eventID, event));
	}

	// Freezes all TypedEvent objects so that new calls to emit() do not fire the event.
	// Instead, the events will be held until the next call to unfreezeAll(), at which point
	// all events will fire all of the events that were frozen.
	//
	// This is used when a single user action activates multiple separate events, to ensure
	// none of them fire until all changes have been applied.
	static freezeAll() {
		freezeCount++;
	}

	static unfreezeAll() {
		freezeCount--;
		if (freezeCount > 0) {
			// Don't do anything until things are fully unfrozen.
			return;
		}

		const typedEvents = frozenTypedEvents.slice();
		frozenTypedEvents = [];

		typedEvents.forEach(typedEvent => {
			const frozenEvents = typedEvent.frozenEvents.slice();
			typedEvent.frozenEvents = [];

			frozenEvents.forEach(frozenEvent => typedEvent.fireEventInternal(frozenEvent.eventID, frozenEvent.event));
		});
	}

	static nextEventID(): EventID {
		return nextEventID++;
	}
}

// If this is > 0 then events are frozen.
let freezeCount = 0;

let frozenTypedEvents: Array<TypedEvent<any>> = [];
let nextEventID: EventID = 0;
