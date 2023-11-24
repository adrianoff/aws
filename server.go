package main

import (
	"io"
	"net/http"
	"strconv"
	"fmt"
	"os"
	"encoding/binary"
)

type BitmapFileHeader struct {
	Type              [2]byte // Signature "BM"
	FileSize          uint32  // Size of the file in bytes
	Reserved1         uint16  // Reserved field (unused)
	Reserved2         uint16  // Reserved field (unused)
	PixelDataOffset   uint32  // Offset to the start of pixel data
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

func getTest(w http.ResponseWriter, r *http.Request) {
	filename := "test.bmp"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var fileHeader BitmapFileHeader
	var infoHeader BitmapInfoHeader

	err = binary.Read(file, binary.LittleEndian, &fileHeader)
	if err != nil {
		fmt.Println("Error reading file header:", err)
		return
	}

	err = binary.Read(file, binary.LittleEndian, &infoHeader)
	if err != nil {
		fmt.Println("Error reading info header:", err)
		return
	}

	pixelDataSize := infoHeader.ImageSize
	file.Seek(int64(fileHeader.PixelDataOffset), 0)
	pixelData := make([]byte, pixelDataSize)
	_, err = file.Read(pixelData)
	if err != nil {
		fmt.Println("Error reading pixel data:", err)
		return
	}

	response := ""
	for i := 0; i < 48000 && i < len(pixelData); i++ {
		response = response + fmt.Sprintf("%02X", pixelData[i])
	}

	response = Reverse(response)

	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if (limit-offset > len(response)) {
		offset = 0
		limit  = len(response)
	}

	response = response[offset:offset+limit]

	w.Header().Set("Content-Length", strconv.Itoa(len(response)));
	io.WriteString(w, response);
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
	http.HandleFunc("/test", getTest)

	http.ListenAndServe(":3333", nil)
}
