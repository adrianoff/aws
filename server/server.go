package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type BitmapFileHeader struct {
	Type            [2]byte // Signature "BM"
	FileSize        uint32  // Size of the file in bytes
	Reserved1       uint16  // Reserved field (unused)
	Reserved2       uint16  // Reserved field (unused)
	PixelDataOffset uint32  // Offset to the start of pixel data
}

type BitmapInfoHeader struct {
	Size                 uint32 // Size of the info header (40 bytes)
	Width                int32  // Image width in pixels
	Height               int32  // Image height in pixels
	Planes               uint16 // Number of color planes (must be 1)
	BitsPerPixel         uint16 // Bits per pixel (1, 4, 8, 16, 24, 32)
	Compression          uint32 // Compression method
	ImageSize            uint32 // Size of the raw bitmap data
	HorizontalResolution int32  // Horizontal resolution (pixels per meter)
	VerticalResolution   int32  // Vertical resolution (pixels per meter)
	ColorsUsed           uint32 // Number of colors in the palette
	ColorsImportant      uint32 // Number of important colors used
}

type Forecast struct {
	Current struct {
		WeatherCode string

		Temperature struct {
			Real     int
			Apparent int
			High     int
			Low      int
		}
		IsDay         bool
		Precipitation float32
		Wind          struct {
			Speed     int
			Direction int
		}
		Sunrise string
		Sunset  string
	}

	Hours [3]struct {
		WeatherCode string
		Visibility  int
		Humidity    int

		Temperature struct {
			Real     int
			Apparent int
		}
		Wind struct {
			Speed     int
			Direction int
		}
		Precipitation struct {
			Probability int
			Quantity    float32
		}
	}
}

var pixelData []byte
var forecast Forecast

func readPixelData() {
	filename := "weather_image.bmp"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var fileHeader BitmapFileHeader
	var infoHeader BitmapInfoHeader

	err = binary.Read(file, binary.LittleEndian, &fileHeader)
	if err != nil {
		fmt.Println("Error reading file header:", err)
	}

	err = binary.Read(file, binary.LittleEndian, &infoHeader)
	if err != nil {
		fmt.Println("Error reading info header:", err)
		return
	}

	pixelDataSize := infoHeader.ImageSize
	file.Seek(int64(fileHeader.PixelDataOffset), 0)
	pixelData = make([]byte, pixelDataSize)
	_, err = file.Read(pixelData)
	if err != nil {
		fmt.Println("Error reading pixel data:", err)
		return
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	response := ""
	for i := 0; i < len(pixelData); i++ {
		response = response + fmt.Sprintf("%02X", ^(pixelData[i]))
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

func getForecast() {
	
}

func main() {
	readPixelData()
	http.HandleFunc("/data", getData)

	http.ListenAndServe(":3333", nil)
}
