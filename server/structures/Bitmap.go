package structures

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
