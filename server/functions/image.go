package functions

import (
	"encoding/binary"
	"os"

	"github.com/adrianoff/aws/server/structures"
)

func ReadPixelData(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fileHeader structures.BitmapFileHeader
	var infoHeader structures.BitmapInfoHeader

	err = binary.Read(file, binary.LittleEndian, &fileHeader)
	if err != nil {
		return nil, err
	}

	err = binary.Read(file, binary.LittleEndian, &infoHeader)
	if err != nil {
		return nil, err
	}

	pixelDataSize := infoHeader.ImageSize
	file.Seek(int64(fileHeader.PixelDataOffset), 0)
	pixelData := make([]byte, pixelDataSize)
	_, err = file.Read(pixelData)
	if err != nil {
		return nil, err
	}

	return pixelData, nil
}

func RemoveImage(filename string) {
	os.Remove(filename)
}
