/** Provides a type-safe event interface. */
export class TypedEvent {
    constructor() {
        this.listeners = [];
        // The IDs of events which have already been fired from this TypedEvent.
        this.firedEvents = [];
        // IDs for currently frozen events pending on this TypedEvent. See freezeAll()
        // for more details.
        this.frozenEvents = [];
    }
    // Registers a new listener to this event.
    on(listener) {
        this.listeners.push(listener);
        return {
            dispose: () => this.off(listener),
        };
    }
    // Removes a listener from this event.
    off(listener) {
        const idx = this.listeners.indexOf(listener);
        if (idx != -1) {
            this.listeners.splice(idx, 1);
        }
    }
    // Convenience for on() which calls off() autmatically after firing once.
    once(listener) {
        const onceListener = (eventID, event) => {
            this.off(onceListener);
            listener(eventID, event);
        };
        return this.on(onceListener);
    }
    emit(eventID, event) {
        if (this.firedEvents.includes(eventID)) {
            if (!thawing) {
                console.warn('EventID collision outside of thawing!');
            }
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
        }
        else {
            this.fireEventInternal(eventID, event);
        }
    }
    fireEventInternal(eventID, event) {
        this.listeners.forEach(listener => listener(eventID, event));
    }
    // Freezes all TypedEvent objects so that new calls to emit() do not fire the event.
    // Instead, the events will be held until the next call to unfreezeAll(), at which point
    // all events will fire all of the events that were frozen.
    //
    // This is used when a single user action activates multiple separate events, to ensure
    // none of them fire until all changes have been applied.
    //
    // All usages of this function should be followed by a call to unfreezeAll()
    // within the same scope. These two functions are very similar to a locking
    // mechanism.
    static freezeAll() {
        freezeCount++;
    }
    static unfreezeAll() {
        freezeCount--;
        if (freezeCount > 0) {
            // Don't do anything until things are fully unfrozen.
            return;
        }
        thawing = true;
        const typedEvents = frozenTypedEvents.slice();
        frozenTypedEvents = [];
        typedEvents.forEach(typedEvent => {
            const frozenEvents = typedEvent.frozenEvents.slice();
            typedEvent.frozenEvents = [];
            frozenEvents.forEach(frozenEvent => typedEvent.fireEventInternal(frozenEvent.eventID, frozenEvent.event));
        });
        thawing = false;
    }
    static nextEventID() {
        return nextEventID++;
    }
}
// If this is > 0 then events are frozen.
let freezeCount = 0;
// Indicates whether we are currently in the process of unfreezing. Just used to add a warning.
let thawing = false;
let frozenTypedEvents = [];
let nextEventID = 0;
