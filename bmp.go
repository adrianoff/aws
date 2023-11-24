package main

import (
	"encoding/binary"
	"fmt"
	"os"
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

func main() {
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

	fmt.Println("Bitmap File Header:")
	fmt.Printf("Type: %c%c\n", fileHeader.Type[0], fileHeader.Type[1])
	fmt.Printf("File Size: %d bytes\n", fileHeader.FileSize)
	fmt.Printf("Pixel Data Offset: %d bytes\n", fileHeader.PixelDataOffset)
	fmt.Println()

	fmt.Println("Bitmap Info Header:")
	fmt.Printf("Size: %d bytes\n", infoHeader.Size)
	fmt.Printf("Width: %d pixels\n", infoHeader.Width)
	fmt.Printf("Height: %d pixels\n", infoHeader.Height)
	fmt.Printf("Planes: %d\n", infoHeader.Planes)
	fmt.Printf("Bits Per Pixel: %d\n", infoHeader.BitsPerPixel)
	fmt.Printf("Compression: %d\n", infoHeader.Compression)
	fmt.Printf("Image Size: %d bytes\n", infoHeader.ImageSize)
	fmt.Printf("Horizontal Resolution: %d pixels per meter\n", infoHeader.HorizontalResolution)
	fmt.Printf("Vertical Resolution: %d pixels per meter\n", infoHeader.VerticalResolution)
	fmt.Printf("Colors Used: %d\n", infoHeader.ColorsUsed)
	fmt.Printf("Colors Important: %d\n", infoHeader.ColorsImportant)

	pixelDataSize := infoHeader.ImageSize
	file.Seek(int64(fileHeader.PixelDataOffset), 0)
	pixelData := make([]byte, pixelDataSize)
	_, err = file.Read(pixelData)
	if err != nil {
		fmt.Println("Error reading pixel data:", err)
		return
	}

	// Print a sample of pixel data (first 100 bytes)
	fmt.Println("Pixel Data (First 100 bytes):")
	s := ""
	for i := 0; i < 48000 && i < len(pixelData); i++ {
		s = s + fmt.Sprintf("%02X", pixelData[i])
		
		// if (i+1)%16 == 0 {
		// 	fmt.Println()
		// }
	}
	fmt.Println(len(s))
}