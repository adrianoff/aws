package main

import (
	"fmt"
	"net/http"

	"github.com/adrianoff/aws/server/functions"
	httphandler "github.com/adrianoff/aws/server/httpHandler"
)

func main() {
	var err error

	_, err = functions.GetOpenMeteoForecast()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/pixels", httphandler.HandlePixelRequest)
	http.HandleFunc("/start_session", httphandler.HandleStartSession)
	http.HandleFunc("/stop_session", httphandler.HandleStopSession)

	http.ListenAndServe(":3333", nil)
}
