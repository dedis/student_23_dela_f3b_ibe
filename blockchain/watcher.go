package blockchain

import "sync"

// Observer is the interface to implement to watch events.
type Observer interface {
	NotifyCallback(event interface{})
}

// Observable provides primitives to add and remove observers and to notify
// them of new events.
type Observable interface {
	Add(observer Observer)
	Remove(observer Observer)
	Notify(event interface{})
}

// Watcher is an implementation of the Observable interface.
type Watcher struct {
	sync.Mutex
	observers map[Observer]struct{}
}

// NewWatcher creates a new empty watcher.
func NewWatcher() *Watcher {
	return &Watcher{
		observers: make(map[Observer]struct{}),
	}
}

// Add adds the observer to the list of observers that will be notified of
// new events.
func (w *Watcher) Add(observer Observer) {
	w.Lock()
	w.observers[observer] = struct{}{}
	w.Unlock()
}

// Remove removes the observer from the list thus stopping it from receiving
// new events.
func (w *Watcher) Remove(observer Observer) {
	w.Lock()
	delete(w.observers, observer)
	w.Unlock()
}

// Notify notifies the whole list of observers one after each other.
func (w *Watcher) Notify(event interface{}) {
	w.Lock()
	defer w.Unlock()

	for w := range w.observers {
		w.NotifyCallback(event)
	}
}