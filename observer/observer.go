// Package observer implements the Observer design pattern.
// Observers are defined by the Observer interface, and publishers (subjects) implement the Observable interface.
// TvShow is a concrete implementation of Observable that can broadcast TV programs to its subscribers.
// TvStation is a concrete implementation of Observer that reacts to updates through the Update callback function while
// subscribed to one or more TvShows.
package observer

func Run() {
	// Creating the TV stations:
	cpb := NewTvStation("Central Public Broadcasting", 150.50)
	sng := NewTvStation("Sport: National & Global channel", 40.5)

	// Creating the TV shows:
	trillionaire := NewTvShow("Everyone Wants To Be A Trillionaire")
	baseball := NewTvShow("National Baseball Cup 2025")
	news := NewTvShow("Short News Announcers")

	// Subscription:
	trillionaire.Subscribe(cpb)
	baseball.Subscribe(sng)
	news.Subscribe(cpb)
	news.Subscribe(sng)

	// Start broadcasts:
	trillionaire.Broadcast("Welcome to Everyone Wants to Be a Trillionaire! Meet our first contestant, Mr. Norman Bates!")
	baseball.Broadcast("The first game of the season begins shortly — grab your beer and get ready!")
	news.Broadcast(
		"The price of PoorCoin has dropped 80% in the past two days. Analysts predict it will rebound sharply within a " +
			"day or two.",
	)

	// Sport channel does not want to break for the news anymore:
	news.Unsubscribe(sng)
	news.Broadcast(
		"Breaking news! PoorCoin has crashed 280% in the last two minutes, following the President’s speech where he " +
			"sneezed into a blanket with the PoorCoin logo.",
	)
	baseball.Broadcast("We’re kicking off! Let’s meet our teams!")
}
