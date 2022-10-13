package main

func startApp() error {
	router := newRoutes()
	return startServer(router)
}
