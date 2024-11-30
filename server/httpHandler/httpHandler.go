package httphandler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/adrianoff/aws/server/functions"
	"github.com/xyproto/randomstring"
)

var Sessions = make(map[string][]byte)

func HandlePixelRequest(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	session_id := r.URL.Query().Get("session_id")

	response := ""
	pixelData, ok := Sessions[session_id]
	if !ok {
		response := "Session not found"
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Length", strconv.Itoa(len(response)))
		io.WriteString(w, response)
		return
	}

	for i := 0; i < len(pixelData); i++ {
		response = response + fmt.Sprintf("%02X", ^(pixelData[i]))
	}

	if offset+limit > len(response) {
		offset = 0
		limit = len(response)
	}

	fmt.Println(offset, " ", offset+limit)
	response = response[offset : offset+limit]

	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	io.WriteString(w, response)
}

func HandleStartSession(w http.ResponseWriter, r *http.Request) {
	forecast, _ := functions.GetOpenMeteoForecast()

	randomstring.Seed()
	session_id := randomstring.CookieFriendlyString(32)

	template := functions.ReadTemplate()
	html := functions.PrepareHtml(string(template), forecast)
	functions.ConvertToImage(html, session_id)

	image_filename := strings.Join([]string{"images/", session_id, ".bmp"}, "")
	pixelData, _ := functions.ReadPixelData(image_filename)
	functions.RemoveImage(image_filename)

	Sessions[session_id] = pixelData

	response := session_id
	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	io.WriteString(w, response)
}

func HandleStopSession(w http.ResponseWriter, r *http.Request) {
	session_id := r.URL.Query().Get("session_id")
	delete(Sessions, session_id)

	response := "Session removed"

	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	io.WriteString(w, response)
}
