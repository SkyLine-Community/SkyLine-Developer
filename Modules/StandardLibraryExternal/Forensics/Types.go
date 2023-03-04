package SkyLine_Standard_External_Forensics

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// This path of (n) contains all digital forensics based functions for images, files, memory, binary files and much more among that list.
//
// n = StandardLibraryExternal/Forensics
//
// This type of module contains much more files and information for general digital forensics and steganography. This includes functions and
// types as well as parsers that can look up signatures in images, verify unifentified files, pull data from files and images, extract signatures
// and uses the syscall table to read and write data to images with an option that allows you to encrypt and then inject data into an image. Currently
// this library only works with mainly PNG images but also has support for JPG/JPEG images and as far as general injection goes it works with most images
// which are common that includes BMP, JPG, JPEG, PNG, GIF, WEBP, WebM and other various file formats.
//
//
// - Date of start : Wed 01 Mar 2023 02:18:32 PM EST
//
// - File name     : Types.go
//
// - File contains : All models, settings, codes and errors that can exist within or are used within this library.

const (
	PNG_SIG_1 = "PNG"
	PNG_SIG_2 = "89504E470D0A1A0A"
	PNG_PAT_3 = "89 50 4E 47 0D 0A 1A 0A"
	SIG_END   = "IEND"
)

var (
	// SIGNATURE FOR PNG IMAGE IN HEX
	PNG_PAT_V = []byte{
		0x89, 0x50, 0x4E,
		0x47, 0x0D, 0xA,
		0x1A, 0x0A,
	}
)

type General struct {
	Input  string // Input image
	Output string // Output image
}

type PNG_Meta struct {
	Offset int64
	Chunk  PNG_Chunk
}

type PNG_Header struct {
	HEAD uint64
}

type PNG_Chunk struct {
	CRC  uint32 // CRC v
	FD   []byte // File Data v
	Type uint32 // Type v
	Size uint32 // Sizeof v
}
