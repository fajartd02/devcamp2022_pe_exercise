package main

func startApp() error {
	router := newRoutes(bookHandler)
	return startServer(router)
}
