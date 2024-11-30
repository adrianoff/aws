package main

import (
	"net/http"
	httphandler "github.com/adrianoff/aws/server/httpHandler"
)

func main() {

	http.HandleFunc("/pixels", httphandler.HandlePixelRequest)
	http.HandleFunc("/start_session", httphandler.HandleStartSession)
	http.HandleFunc("/stop_session", httphandler.HandleStopSession)

	http.ListenAndServe(":3333", nil)
}
