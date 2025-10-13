package observer

// Observer defines the interface for all broadcasts observers.
type Observer interface {
	Update(data any)
}

// Observable defines the interface for broadcasts subscriptions.
type Observable interface {
	Subscribe(listener Observer)
	Unsubscribe(listener Observer)
	Notify(data any)
}

// Broadcasting defines the common interface for all broadcasting programs and shows.
type Broadcasting interface {
	Observable
	Broadcast(message string)
}
