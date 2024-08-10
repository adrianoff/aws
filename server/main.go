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

	httphandler.PixelData, err = functions.ReadPixelData("./images/image.bmp")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/pixels", httphandler.HandlePixelRequest)

	http.ListenAndServe(":3333", nil)
}
