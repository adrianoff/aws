package httphandler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var PixelData []byte

func HandlePixelRequest(w http.ResponseWriter, r *http.Request) {
	response := ""
	for i := 0; i < len(PixelData); i++ {
		response = response + fmt.Sprintf("%02X", ^(PixelData[i]))
	}

	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if offset+limit > len(response) {
		offset = 0
		limit = len(response)
	}

	fmt.Println(offset, " ", offset+limit)
	response = response[offset : offset+limit]

	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	io.WriteString(w, response)
}