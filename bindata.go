package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func frontend_assets_application_coffee() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x94, 0x55,
		0x4d, 0x6f, 0xdb, 0x30, 0x0c, 0xbd, 0xeb, 0x57, 0x08, 0x69, 0x36, 0xdb,
		0x43, 0xe2, 0x7e, 0xec, 0x56, 0x20, 0xdd, 0x65, 0x97, 0x0e, 0xdd, 0x5a,
		0x2c, 0xbb, 0x0d, 0x3b, 0xa8, 0x36, 0xed, 0x08, 0x55, 0x6c, 0x43, 0x92,
		0x63, 0x14, 0x45, 0xfe, 0xfb, 0x48, 0x59, 0x8e, 0xed, 0xc6, 0x69, 0xb7,
		0x5e, 0x6a, 0x91, 0x8f, 0x8f, 0xe4, 0x23, 0xa5, 0xc8, 0x8c, 0x37, 0xb2,
		0x48, 0xcb, 0x26, 0x5e, 0xdb, 0x52, 0x8b, 0x1c, 0xb8, 0x28, 0xd2, 0xce,
		0xf4, 0x6d, 0x7d, 0xff, 0x83, 0xf1, 0xee, 0x34, 0x37, 0x1e, 0xb1, 0xe2,
		0xe1, 0x13, 0x3c, 0x47, 0x7c, 0x79, 0x83, 0x4e, 0xce, 0x0d, 0xd8, 0x6b,
		0x1e, 0xee, 0x84, 0xaa, 0xe1, 0x60, 0xe3, 0x5c, 0x95, 0x89, 0x50, 0x9e,
		0x33, 0x46, 0xc8, 0xad, 0x85, 0x2d, 0x85, 0x2d, 0x38, 0xb1, 0xc6, 0xc6,
		0x6a, 0x59, 0xe4, 0x32, 0x7b, 0xf6, 0x81, 0x91, 0x0b, 0xcb, 0x89, 0xea,
		0x40, 0x21, 0x31, 0x04, 0x93, 0x8d, 0x98, 0xf2, 0x9e, 0x29, 0xf2, 0x30,
		0xc7, 0x57, 0x09, 0x6d, 0x20, 0xa4, 0x88, 0x88, 0x63, 0x4f, 0xf4, 0xc1,
		0xd8, 0xbc, 0xe5, 0xf2, 0xf5, 0x2b, 0x61, 0xec, 0x5d, 0x29, 0x52, 0xa4,
		0x2c, 0xca, 0x26, 0x8c, 0x7a, 0x8f, 0x86, 0x4c, 0x83, 0xd9, 0x3c, 0x08,
		0x63, 0x9a, 0x52, 0x13, 0xe0, 0xf3, 0x05, 0xfe, 0xf5, 0x00, 0x2b, 0x93,
		0x27, 0xd0, 0xb7, 0x85, 0x05, 0x8d, 0xe5, 0xa2, 0xff, 0x0a, 0xbd, 0xaf,
		0xdd, 0x68, 0xa6, 0x3e, 0x3b, 0x10, 0x19, 0x17, 0xd3, 0x04, 0x18, 0x39,
		0x0f, 0x83, 0xb3, 0xac, 0x4c, 0x6a, 0x03, 0xe9, 0x6d, 0x51, 0xd5, 0x36,
		0x88, 0xe2, 0x47, 0x84, 0xf2, 0x20, 0x51, 0x08, 0x0d, 0x16, 0x1c, 0xc5,
		0xae, 0x7e, 0xb6, 0x75, 0xbd, 0x05, 0x7f, 0x54, 0xb5, 0x46, 0x34, 0xc2,
		0xac, 0xd0, 0x76, 0x1c, 0x50, 0x56, 0x56, 0x96, 0xc5, 0x5a, 0xec, 0xe0,
		0x98, 0x1d, 0x8d, 0xf7, 0xce, 0x6d, 0xa8, 0x0f, 0x85, 0xb2, 0xf8, 0xa3,
		0xd3, 0x85, 0xce, 0x9d, 0x1a, 0x68, 0x60, 0xa8, 0x17, 0x4d, 0xdd, 0x8f,
		0xd7, 0x29, 0x08, 0x0d, 0xff, 0x2a, 0x2c, 0x38, 0x78, 0x4a, 0x63, 0xf9,
		0x25, 0xb7, 0x74, 0x62, 0x83, 0xc2, 0x07, 0x31, 0x89, 0x02, 0x71, 0xe8,
		0x3f, 0x1c, 0xa9, 0x12, 0x4d, 0xf7, 0x67, 0x40, 0x41, 0x62, 0x1d, 0x7f,
		0x26, 0x94, 0x01, 0xc6, 0xc6, 0x4d, 0x0e, 0xc8, 0xdf, 0x1d, 0xf0, 0x7f,
		0x0f, 0x88, 0x21, 0xf0, 0x41, 0x97, 0x39, 0x26, 0x32, 0x94, 0xa8, 0x02,
		0x9d, 0xf8, 0x64, 0x58, 0x6a, 0x5c, 0x79, 0xd7, 0xf2, 0x51, 0x68, 0x2c,
		0x35, 0x31, 0x26, 0x0c, 0x1a, 0x99, 0xda, 0x0d, 0x4a, 0x3b, 0x3b, 0x7b,
		0x21, 0xf4, 0xfe, 0xc3, 0x0c, 0xc5, 0x18, 0x0a, 0x39, 0x28, 0xb8, 0x9d,
		0x8c, 0x71, 0xcb, 0x3d, 0x56, 0x7e, 0x4e, 0x52, 0xf2, 0xd9, 0xf9, 0xee,
		0xf2, 0x1c, 0x3f, 0xba, 0xd0, 0x2f, 0x0a, 0x8a, 0xdc, 0x6e, 0x56, 0x67,
		0x2f, 0x3e, 0x12, 0xf7, 0xbd, 0xf5, 0xdc, 0x39, 0xc7, 0xfe, 0xa3, 0xa9,
		0x20, 0x91, 0x42, 0x0d, 0x10, 0xa8, 0xe5, 0xba, 0x35, 0xee, 0x67, 0x0b,
		0x1e, 0xa6, 0xc2, 0x8a, 0xc3, 0xf5, 0x9c, 0x90, 0x9b, 0xc6, 0xe2, 0x30,
		0x0e, 0x70, 0x4a, 0x51, 0x36, 0x58, 0x9c, 0xa9, 0x7e, 0x5c, 0xf0, 0xb8,
		0xb6, 0x6b, 0x97, 0x6d, 0x6c, 0x6b, 0x19, 0x7c, 0xd6, 0x36, 0x63, 0x5f,
		0x6e, 0x1b, 0xd0, 0x9f, 0x3b, 0xf0, 0xef, 0x8b, 0x3f, 0x71, 0xb2, 0x01,
		0x9c, 0x53, 0xca, 0x8e, 0x5f, 0xa5, 0x30, 0x58, 0x43, 0x52, 0x6b, 0xe8,
		0x24, 0xf3, 0x45, 0xba, 0x45, 0xb2, 0xa1, 0x2f, 0xb0, 0x5b, 0x35, 0x34,
		0x59, 0x7c, 0x7e, 0xcc, 0xf7, 0x32, 0x15, 0x0a, 0x21, 0x5b, 0xfa, 0x1f,
		0x06, 0x1b, 0x99, 0xe2, 0x55, 0x99, 0xd8, 0xff, 0xc1, 0x94, 0x26, 0xa7,
		0xf8, 0xcf, 0xb5, 0xe0, 0x4c, 0x5d, 0xbf, 0xf8, 0x44, 0x1d, 0xa2, 0x57,
		0xbc, 0x2e, 0x52, 0xc8, 0x64, 0x81, 0x7d, 0x91, 0x12, 0x63, 0x31, 0x8f,
		0xe5, 0xbc, 0xba, 0xf0, 0x8e, 0xa1, 0x64, 0xed, 0x15, 0xe1, 0xef, 0x69,
		0x3d, 0xbd, 0x3e, 0x9d, 0x2c, 0x6f, 0x6a, 0x8e, 0x6d, 0x1e, 0xaf, 0x56,
		0xaf, 0x02, 0x63, 0x74, 0x85, 0x86, 0xaf, 0x84, 0xcc, 0xb2, 0x6e, 0x6f,
		0xf8, 0xf2, 0xf5, 0x46, 0x21, 0x80, 0x6e, 0x09, 0xe1, 0x4f, 0x3c, 0xc2,
		0x4b, 0xc7, 0x10, 0xf1, 0xf3, 0x53, 0xaf, 0xf4, 0x27, 0x7e, 0xe9, 0xde,
		0xe8, 0xe1, 0x55, 0x25, 0xce, 0x56, 0x5e, 0x97, 0xfe, 0x66, 0x75, 0x22,
		0x98, 0xb5, 0x3f, 0x51, 0xa3, 0x29, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
		0x19, 0x94, 0x55, 0x68, 0x0b, 0x07, 0x00, 0x00,
	},
		"frontend/assets/application.coffee",
	)
}

func frontend_assets_application_js() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xa4, 0x56,
		0x4d, 0x6f, 0xdb, 0x38, 0x10, 0xbd, 0xfb, 0x57, 0x10, 0x5e, 0x6f, 0x24,
		0xed, 0x2a, 0xb2, 0x9d, 0x3d, 0x2c, 0x10, 0xc3, 0xbb, 0x87, 0x16, 0x28,
		0x52, 0xa4, 0x4d, 0x50, 0xf7, 0x56, 0xf4, 0x40, 0x4b, 0x94, 0x4d, 0x44,
		0x11, 0x05, 0x92, 0xb6, 0x11, 0x14, 0xfa, 0xef, 0x1d, 0x7e, 0x88, 0x22,
		0xe5, 0x8f, 0xf4, 0x23, 0x17, 0x8b, 0xe4, 0x70, 0xe6, 0xf1, 0xbd, 0xc7,
		0x61, 0xa6, 0x53, 0xf4, 0x8e, 0xd4, 0x84, 0x63, 0x49, 0x0a, 0xb4, 0x7e,
		0x41, 0x6f, 0x58, 0x59, 0x12, 0xb2, 0xca, 0x39, 0x6d, 0x24, 0x9a, 0x67,
		0xff, 0x66, 0xf3, 0x51, 0x5c, 0xee, 0xea, 0x5c, 0x52, 0x56, 0xc7, 0x09,
		0xfa, 0x36, 0x42, 0x68, 0x8f, 0x39, 0xaa, 0x18, 0x2e, 0x1e, 0x1a, 0x35,
		0x29, 0x52, 0x3d, 0x78, 0xc4, 0x42, 0x1c, 0x18, 0x2f, 0x52, 0x54, 0xb3,
		0x43, 0x8a, 0x38, 0x11, 0x12, 0x73, 0xf9, 0x89, 0x94, 0xf0, 0xb5, 0x4d,
		0x91, 0xc0, 0x7b, 0xe2, 0xe2, 0x05, 0x91, 0x8f, 0x9c, 0x6d, 0x60, 0x45,
		0x0d, 0x24, 0x6b, 0x5c, 0x98, 0xa4, 0xf9, 0xd3, 0x62, 0x04, 0x35, 0x68,
		0x89, 0xe2, 0x03, 0xad, 0x0b, 0x76, 0xc8, 0x56, 0x92, 0x71, 0xbc, 0x21,
		0xe8, 0xea, 0x0a, 0xd9, 0x99, 0xf7, 0xab, 0x87, 0x8f, 0x06, 0x0a, 0xea,
		0xa6, 0x26, 0xc2, 0x46, 0x2d, 0x91, 0x43, 0xfb, 0x44, 0x5e, 0xba, 0x28,
		0x04, 0x80, 0xe4, 0x8e, 0xd7, 0x6e, 0x88, 0x14, 0x88, 0xdb, 0x3e, 0x76,
		0x8f, 0xab, 0x1d, 0x49, 0xbc, 0x65, 0xb7, 0xa3, 0x62, 0x39, 0xae, 0x2c,
		0x86, 0x0c, 0x36, 0xdd, 0x49, 0xf2, 0xac, 0x52, 0xa7, 0x48, 0xc1, 0xc8,
		0x84, 0xe4, 0xb4, 0xde, 0xd0, 0xf2, 0xc5, 0xa6, 0x48, 0x16, 0x2e, 0x45,
		0x9b, 0xba, 0xcf, 0x4d, 0x50, 0x2c, 0xac, 0xa3, 0xe8, 0xa4, 0x90, 0x74,
		0xe1, 0xcd, 0xa9, 0x31, 0x1c, 0x25, 0xa8, 0xbd, 0xe9, 0x6b, 0x27, 0x41,
		0x2c, 0x50, 0xa5, 0xe2, 0xc3, 0xac, 0x0e, 0xbf, 0x46, 0xd9, 0x60, 0x2e,
		0x88, 0x89, 0xf2, 0xb7, 0xb6, 0xa3, 0xe1, 0x57, 0x6b, 0x96, 0xf5, 0x4f,
		0xab, 0x84, 0x98, 0x0c, 0xd5, 0x77, 0x94, 0x57, 0x58, 0xc8, 0x7b, 0x50,
		0x1e, 0x70, 0x82, 0xe4, 0xb1, 0x4d, 0x6c, 0x17, 0xb9, 0x51, 0xb4, 0x73,
		0x05, 0xc4, 0xfc, 0x33, 0x83, 0xbf, 0x20, 0x46, 0x89, 0x4d, 0xf8, 0x5d,
		0x2d, 0x09, 0x07, 0xee, 0x20, 0xe4, 0xe6, 0x64, 0x00, 0x2c, 0x28, 0xda,
		0x6d, 0x58, 0xac, 0x26, 0xd3, 0xd3, 0x29, 0x2c, 0x84, 0x49, 0x1c, 0xfd,
		0x51, 0xb2, 0x7c, 0x27, 0x48, 0x71, 0x57, 0x37, 0x3b, 0x19, 0x25, 0xd9,
		0x1a, 0xc2, 0xe3, 0x28, 0xaf, 0x20, 0x3c, 0x0a, 0x2c, 0xf7, 0xda, 0x96,
		0x75, 0xb5, 0xe3, 0xd1, 0xd0, 0xce, 0xde, 0x26, 0xa6, 0x3d, 0xbd, 0x02,
		0x77, 0x1f, 0x57, 0xe9, 0x2d, 0x6f, 0x37, 0x78, 0x97, 0xa6, 0xa3, 0xcb,
		0x99, 0xac, 0xbf, 0x41, 0x66, 0xa9, 0x4d, 0xf4, 0x3d, 0x00, 0x66, 0x7d,
		0x4b, 0x77, 0x12, 0x28, 0xcf, 0x14, 0x26, 0x83, 0xe6, 0x9f, 0x1c, 0xd0,
		0x5b, 0xb8, 0xc1, 0x83, 0xac, 0x85, 0xf2, 0xcc, 0x67, 0xfa, 0x6c, 0xe7,
		0x5b, 0x9d, 0xd1, 0x3b, 0xfc, 0xa9, 0xcc, 0x79, 0x45, 0xb0, 0x23, 0x34,
		0x0e, 0x68, 0x3e, 0xcf, 0x95, 0x20, 0x15, 0xc9, 0xe5, 0xa0, 0x7a, 0x89,
		0x2b, 0x41, 0xfa, 0xba, 0x21, 0x85, 0xa7, 0x4a, 0x5f, 0xf4, 0x95, 0x4d,
		0xfa, 0x3b, 0xce, 0xb0, 0xe7, 0xef, 0x9b, 0x8f, 0x0f, 0xa2, 0x21, 0x3c,
		0xef, 0x80, 0xd8, 0x5a, 0x70, 0xd0, 0xac, 0xb1, 0xa1, 0xd7, 0x6b, 0xcc,
		0xe1, 0xa0, 0xb9, 0x10, 0x71, 0x74, 0xa0, 0x85, 0xdc, 0x82, 0xc0, 0xe3,
		0x31, 0xfa, 0x1b, 0xa9, 0x7d, 0xf0, 0x33, 0xfe, 0x73, 0xec, 0xd5, 0xf0,
		0xe5, 0x3c, 0x27, 0x9f, 0x71, 0x8e, 0x30, 0xa7, 0xb3, 0x03, 0x7d, 0xe5,
		0xcf, 0x79, 0x64, 0xa2, 0xd4, 0x8c, 0xc7, 0xd3, 0xfd, 0x7c, 0x0a, 0x1f,
		0x5d, 0xfa, 0xff, 0x2b, 0x52, 0x6f, 0xe4, 0x76, 0xa9, 0xb0, 0xd8, 0x2c,
		0x70, 0xdb, 0xcd, 0xda, 0xbd, 0x5e, 0x52, 0xe8, 0xae, 0x44, 0x43, 0x72,
		0x8a, 0xab, 0x20, 0x0c, 0x14, 0x5c, 0x99, 0xe9, 0xb4, 0x87, 0x58, 0x60,
		0x89, 0xfb, 0x4e, 0x72, 0x42, 0x6b, 0xc5, 0xb4, 0x0e, 0x5a, 0x84, 0x9d,
		0xf5, 0xa2, 0x7a, 0xad, 0x2f, 0x40, 0x7f, 0x2f, 0x7e, 0x92, 0x9b, 0x0e,
		0x55, 0x78, 0xbe, 0x5b, 0x8d, 0x32, 0x9c, 0x33, 0xf9, 0x2d, 0xda, 0xa4,
		0xeb, 0xc3, 0xfd, 0x81, 0xcd, 0x9e, 0x7e, 0xdc, 0xc5, 0x7f, 0x99, 0x7d,
		0xcd, 0xf2, 0x2d, 0x01, 0xeb, 0x14, 0x23, 0xaf, 0x1d, 0x0e, 0x1e, 0x99,
		0x38, 0x5a, 0x91, 0x7c, 0xc7, 0x49, 0xa7, 0x81, 0x3d, 0x8d, 0xbe, 0x08,
		0x32, 0x66, 0xc1, 0x9d, 0x57, 0x75, 0x60, 0x56, 0xc2, 0x23, 0x21, 0x3e,
		0xb0, 0x02, 0x57, 0x10, 0xf5, 0xac, 0x7e, 0xe3, 0x68, 0x4b, 0x0b, 0xe8,
		0x1a, 0xaf, 0xf5, 0x01, 0x67, 0xa8, 0x5f, 0xe3, 0xec, 0x87, 0xa1, 0x2b,
		0x73, 0x59, 0x30, 0xea, 0x45, 0x71, 0x19, 0x96, 0x4b, 0xb4, 0x67, 0xb4,
		0x40, 0xb3, 0xde, 0x15, 0xc7, 0x8a, 0x1c, 0x6b, 0x72, 0x33, 0xeb, 0x5f,
		0x3f, 0x9f, 0x77, 0xdd, 0x15, 0x06, 0x8f, 0x8d, 0x23, 0xea, 0x82, 0x88,
		0xa7, 0xbd, 0xed, 0x91, 0x7c, 0x51, 0x4c, 0x80, 0x7a, 0xec, 0xfa, 0x80,
		0x79, 0x8f, 0x3e, 0x43, 0xb9, 0x6a, 0x20, 0x67, 0x5b, 0x2f, 0x2d, 0xcb,
		0x54, 0x5f, 0x7d, 0xdb, 0x83, 0x61, 0xdc, 0x19, 0x1e, 0x5d, 0x0f, 0xaf,
		0x82, 0x89, 0xd1, 0x8d, 0x62, 0xe9, 0xfe, 0xab, 0x19, 0x3e, 0x90, 0xd7,
		0x3a, 0x49, 0x82, 0xa6, 0xe7, 0x5e, 0xd0, 0xbf, 0xd0, 0xbc, 0x7b, 0x1e,
		0xbd, 0x16, 0x66, 0xfa, 0x56, 0x2f, 0x9b, 0x46, 0xf2, 0xdf, 0xf2, 0x4c,
		0x92, 0xa3, 0x7f, 0x87, 0x8e, 0xfd, 0x66, 0xe4, 0x50, 0x14, 0xb4, 0xd0,
		0xee, 0x70, 0x05, 0x6d, 0x75, 0x4b, 0x95, 0x99, 0xbf, 0x07, 0x00, 0x00,
		0xff, 0xff, 0x75, 0xeb, 0x78, 0x8c, 0x29, 0x0a, 0x00, 0x00,
	},
		"frontend/assets/application.js",
	)
}

func frontend_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xbc, 0x58,
		0xff, 0x72, 0xdc, 0xb6, 0xf1, 0xff, 0xdf, 0x4f, 0xb1, 0xa6, 0xe7, 0x3b,
		0x49, 0x3c, 0x26, 0x69, 0x49, 0xce, 0xb7, 0xae, 0x7c, 0xa7, 0x8e, 0xeb,
		0xb8, 0xb1, 0x32, 0x4e, 0x9c, 0x44, 0xce, 0xb4, 0x99, 0x4c, 0xfe, 0x00,
		0xc9, 0x25, 0x09, 0x09, 0x04, 0x60, 0x00, 0xbc, 0xd3, 0x35, 0xd3, 0xe7,
		0xe8, 0x03, 0xf5, 0xc5, 0xba, 0x00, 0x7f, 0x1c, 0xef, 0x8e, 0x67, 0xc9,
		0x6d, 0xa6, 0xe3, 0x39, 0x8b, 0x00, 0x17, 0x8b, 0xfd, 0xf9, 0xd9, 0x5d,
		0x2e, 0x1e, 0x7e, 0xf5, 0xee, 0xd5, 0xfb, 0x9f, 0xbf, 0x7f, 0x0d, 0xb5,
		0x6b, 0xc4, 0xc5, 0x83, 0x85, 0xff, 0x03, 0x82, 0xc9, 0x6a, 0x19, 0xa1,
		0x8c, 0x2e, 0x1e, 0x00, 0x2c, 0x6a, 0x64, 0x85, 0x7f, 0xa0, 0xc7, 0x06,
		0x1d, 0x83, 0xbc, 0x66, 0xc6, 0xa2, 0x5b, 0x46, 0xad, 0x2b, 0xe3, 0xe7,
		0xd1, 0xf4, 0x55, 0xed, 0x9c, 0x8e, 0xf1, 0x43, 0xcb, 0x57, 0xcb, 0xe8,
		0x6f, 0xf1, 0x4f, 0x2f, 0xe3, 0x57, 0xaa, 0xd1, 0xcc, 0xf1, 0x4c, 0x60,
		0x04, 0xb9, 0x92, 0x0e, 0x25, 0x9d, 0xbb, 0x7c, 0xbd, 0xc4, 0xa2, 0xc2,
		0x9d, 0x93, 0x92, 0x35, 0xb8, 0x8c, 0x56, 0x1c, 0xd7, 0x5a, 0x19, 0x37,
		0x21, 0x5e, 0xf3, 0xc2, 0xd5, 0xcb, 0x02, 0x57, 0x3c, 0xc7, 0x38, 0x2c,
		0x9e, 0x00, 0x97, 0xdc, 0x71, 0x26, 0x62, 0x9b, 0x33, 0x81, 0xcb, 0x93,
		0x27, 0xd0, 0xb0, 0x5b, 0xde, 0xb4, 0xcd, 0x76, 0xa3, 0xb5, 0x68, 0xc2,
		0x8a, 0xd1, 0xd5, 0x4b, 0xa9, 0x86, 0xbb, 0x1e, 0xc6, 0x31, 0xbc, 0xaf,
		0x11, 0x58, 0xa6, 0x56, 0x08, 0x67, 0x10, 0xee, 0x76, 0xac, 0xb2, 0xf0,
		0xb8, 0x69, 0xad, 0x7b, 0x4c, 0xf7, 0x36, 0x08, 0x25, 0x37, 0xd6, 0xd1,
		0x2d, 0xe0, 0x88, 0xd4, 0xab, 0xff, 0x02, 0x98, 0xdc, 0x80, 0xa2, 0xa5,
		0x09, 0xeb, 0x41, 0x3c, 0xf0, 0x87, 0xba, 0x33, 0x8f, 0x59, 0xe9, 0xd0,
		0x3c, 0xf6, 0x47, 0x2c, 0x76, 0x2c, 0xe3, 0xb8, 0xbf, 0xd5, 0x71, 0x27,
		0xf0, 0xe2, 0x0a, 0xf3, 0xd6, 0x20, 0x7c, 0xcf, 0xac, 0x5d, 0x2b, 0x53,
		0x2c, 0xd2, 0x6e, 0xfb, 0xc1, 0x56, 0xb2, 0x3f, 0x2b, 0xe5, 0xac, 0x33,
		0x4c, 0x6f, 0x8f, 0x0a, 0x2e, 0x6f, 0xc0, 0xa0, 0x58, 0x46, 0xd6, 0x6d,
		0x04, 0xda, 0x1a, 0x91, 0xac, 0x53, 0x1b, 0x2c, 0x97, 0x91, 0xb7, 0xb6,
		0x3d, 0x4f, 0x53, 0xd2, 0x3e, 0x2f, 0x64, 0x92, 0x0d, 0xa7, 0xfd, 0x82,
		0x64, 0x4a, 0xc7, 0x8d, 0xf4, 0x2c, 0x39, 0x4b, 0x9e, 0xa5, 0xb9, 0xb5,
		0xdb, 0xbd, 0xa4, 0xe1, 0x44, 0x65, 0x6d, 0xf4, 0x09, 0x17, 0x85, 0xc3,
		0x6b, 0xe6, 0xf2, 0x3a, 0xf0, 0xd7, 0x4c, 0xa3, 0x99, 0xe5, 0xb8, 0x55,
		0xe9, 0xcd, 0xfb, 0x6f, 0xdf, 0x7e, 0x09, 0xb6, 0xe6, 0x0d, 0x99, 0xb0,
		0x80, 0x1f, 0xd1, 0x6a, 0x25, 0x8b, 0xe4, 0xda, 0x42, 0xa9, 0x0c, 0x5c,
		0xbe, 0x7e, 0x0e, 0xb6, 0xd5, 0xde, 0xe5, 0xa0, 0xca, 0x9e, 0x18, 0x05,
		0x36, 0x64, 0x5b, 0x1b, 0x0e, 0x34, 0x58, 0x70, 0x06, 0x1f, 0x5a, 0x34,
		0x1c, 0x27, 0x16, 0xf5, 0xac, 0xff, 0xfa, 0xf2, 0xc7, 0xef, 0x2e, 0xbf,
		0xfb, 0xfa, 0x7c, 0xca, 0xb4, 0x50, 0x68, 0xe5, 0x67, 0x0e, 0xc8, 0xc0,
		0x37, 0xc0, 0x4b, 0xd8, 0xa8, 0x16, 0x7c, 0x50, 0x05, 0x4f, 0x6a, 0x56,
		0x21, 0xad, 0x18, 0xf9, 0x57, 0x20, 0xa9, 0xb3, 0xc3, 0xee, 0x17, 0xa2,
		0x16, 0x8e, 0x24, 0x82, 0x3f, 0xfe, 0xda, 0xed, 0xd2, 0xbe, 0xcd, 0x0d,
		0xd7, 0x0e, 0xac, 0xc9, 0xb7, 0x46, 0x50, 0xd6, 0x26, 0xbd, 0xc5, 0xbd,
		0x11, 0x7c, 0xc2, 0x7c, 0x49, 0xfa, 0xad, 0xc8, 0xc8, 0x7f, 0x48, 0x4e,
		0xb7, 0xeb, 0x60, 0x8e, 0x6b, 0xb2, 0xc6, 0x22, 0xed, 0xd8, 0x7c, 0x0a,
		0x57, 0xd3, 0xa9, 0x94, 0x9e, 0x24, 0xcf, 0x88, 0x67, 0xbf, 0x3a, 0xc2,
		0x71, 0xf1, 0xf0, 0x17, 0x94, 0x05, 0x2f, 0x7f, 0x1d, 0xd5, 0x09, 0x1e,
		0x1c, 0xae, 0xcb, 0x54, 0xb1, 0x81, 0xdf, 0x20, 0x63, 0xf9, 0x4d, 0x65,
		0x54, 0x2b, 0x8b, 0x38, 0x57, 0x42, 0x99, 0x73, 0x78, 0x84, 0xa7, 0xfe,
		0xdf, 0x0b, 0xf8, 0x47, 0x4f, 0x99, 0xac, 0x6b, 0xee, 0x70, 0x9e, 0xb6,
		0x2c, 0xcb, 0x09, 0x61, 0x4e, 0xfe, 0xa1, 0x54, 0xf8, 0x0d, 0x1c, 0xde,
		0xba, 0x98, 0x09, 0x5e, 0xc9, 0x73, 0xe8, 0x36, 0x27, 0x54, 0x25, 0xfb,
		0x40, 0x24, 0x0d, 0x33, 0x15, 0x97, 0xb1, 0x53, 0xfa, 0x1c, 0xce, 0xb0,
		0x99, 0xbc, 0xd7, 0x46, 0x55, 0xa4, 0x9a, 0xdd, 0x12, 0x65, 0xca, 0x39,
		0xd5, 0x9c, 0xc3, 0x53, 0x7d, 0x3b, 0xa1, 0xbb, 0x6e, 0x1b, 0x7a, 0x61,
		0x94, 0x04, 0x3d, 0x43, 0x3a, 0x10, 0x92, 0x51, 0x06, 0xb5, 0x17, 0xe9,
		0x80, 0x59, 0x0b, 0xaf, 0x7d, 0x6f, 0x15, 0xc9, 0x56, 0x90, 0x0b, 0xca,
		0xbe, 0x65, 0x44, 0x8f, 0x19, 0x33, 0xd0, 0xfd, 0x89, 0x0b, 0x2c, 0x59,
		0x2b, 0x5c, 0x34, 0x3a, 0xa8, 0xe0, 0x23, 0xa5, 0x4f, 0x74, 0xc6, 0x25,
		0xa1, 0x49, 0x29, 0x5a, 0x5e, 0x8c, 0x34, 0x43, 0xc6, 0x1a, 0x1f, 0xa4,
		0xfe, 0xe7, 0x54, 0x55, 0x09, 0x84, 0x0a, 0x1d, 0x78, 0xd3, 0x69, 0x2c,
		0x42, 0x84, 0x67, 0xe8, 0xbc, 0xa5, 0x1a, 0x95, 0x51, 0xcc, 0x41, 0xc1,
		0xad, 0x16, 0x6c, 0x33, 0x06, 0xde, 0xfe, 0x6d, 0xbd, 0x40, 0x5e, 0x7a,
		0x34, 0x93, 0xbb, 0xbc, 0x22, 0x2d, 0xe9, 0x4b, 0x88, 0xb4, 0xd1, 0x84,
		0x93, 0xdd, 0x22, 0xda, 0x3b, 0xd6, 0x89, 0x10, 0x41, 0xc1, 0x1c, 0xeb,
		0x17, 0xcb, 0xa8, 0x51, 0x05, 0x13, 0xc3, 0x1e, 0x99, 0xce, 0x03, 0xf7,
		0x23, 0x82, 0x6f, 0xc7, 0x65, 0x65, 0xbf, 0x0d, 0x2f, 0xa7, 0xf7, 0xf8,
		0xe8, 0xd1, 0x4c, 0x0e, 0x9c, 0x2b, 0xb1, 0xd1, 0x35, 0x27, 0x23, 0xc0,
		0xf8, 0x44, 0x21, 0x51, 0x45, 0xc0, 0x0c, 0x67, 0x71, 0xcd, 0x8b, 0x02,
		0xe5, 0x32, 0x72, 0xa6, 0xc5, 0x10, 0x95, 0x74, 0x72, 0x47, 0xe8, 0xb4,
		0x13, 0x74, 0x67, 0x8f, 0xed, 0x49, 0x9d, 0x79, 0x13, 0x0e, 0x48, 0xf3,
		0x28, 0x3a, 0x84, 0x49, 0x36, 0xb1, 0x55, 0x4a, 0xc6, 0x9a, 0x37, 0x1d,
		0x05, 0xaa, 0x60, 0x9a, 0xa0, 0xb7, 0x67, 0x3b, 0xac, 0x23, 0xe0, 0x05,
		0x19, 0xcc, 0xc6, 0x78, 0xcb, 0x1a, 0x2d, 0x30, 0xde, 0x7b, 0x1f, 0x9f,
		0xec, 0xda, 0xb9, 0x15, 0x13, 0xf9, 0x06, 0x66, 0x93, 0x47, 0xc3, 0xab,
		0xda, 0xed, 0x9b, 0x4c, 0xf0, 0xdd, 0x8d, 0x3b, 0xfc, 0x95, 0x39, 0x09,
		0xf4, 0x1b, 0x02, 0x6f, 0x60, 0x4d, 0x5b, 0xff, 0xad, 0xf3, 0x7e, 0x0f,
		0x07, 0xc2, 0x55, 0x7f, 0xc1, 0xbe, 0x4a, 0x33, 0xde, 0xf4, 0xbb, 0xbb,
		0xca, 0x2f, 0xd2, 0x56, 0x1c, 0x71, 0xd8, 0x64, 0xb1, 0x48, 0x49, 0xe7,
		0xe9, 0xb1, 0x07, 0x87, 0xfe, 0xec, 0x13, 0x6f, 0x36, 0x2d, 0x8d, 0x5a,
		0x47, 0x47, 0xe3, 0x20, 0xbe, 0xb5, 0xf1, 0xc9, 0x29, 0xf8, 0x27, 0xdb,
		0xc4, 0xcf, 0x87, 0x07, 0x55, 0x96, 0x64, 0xba, 0xb8, 0x7b, 0xd1, 0x14,
		0xf1, 0xff, 0x0f, 0x0f, 0xfd, 0x8b, 0xb3, 0xa1, 0x70, 0x1d, 0x32, 0x1d,
		0x11, 0x68, 0xdf, 0xf1, 0x7a, 0xbc, 0x36, 0x80, 0xdf, 0x81, 0x37, 0x2e,
		0x61, 0xcd, 0x85, 0x20, 0x54, 0x20, 0x45, 0x18, 0x61, 0x2b, 0x35, 0x3a,
		0x54, 0x8c, 0x6c, 0x17, 0xe2, 0xba, 0x0f, 0x71, 0xc0, 0x15, 0x9a, 0x0d,
		0x9c, 0x3d, 0xf5, 0x2f, 0x08, 0xe7, 0x6d, 0x02, 0xdf, 0x84, 0xbe, 0x42,
		0xf0, 0xfc, 0xa6, 0xaf, 0x5c, 0x3d, 0xa5, 0x53, 0x70, 0x83, 0xa8, 0xc3,
		0x26, 0xf1, 0x30, 0xbe, 0x07, 0x51, 0x12, 0x13, 0xd8, 0xf3, 0x89, 0xde,
		0x75, 0x49, 0xb0, 0xfb, 0x11, 0xdd, 0xc8, 0xe5, 0x28, 0x20, 0xfc, 0x7f,
		0x80, 0x84, 0x47, 0xa8, 0x03, 0x40, 0x51, 0x84, 0x1c, 0xa8, 0xfb, 0xb3,
		0x6a, 0xcd, 0x28, 0xec, 0x9e, 0x48, 0x3b, 0x99, 0x3b, 0xcf, 0xd6, 0x83,
		0xf5, 0x01, 0xcf, 0x29, 0x19, 0x21, 0x6a, 0x13, 0x07, 0x78, 0x9d, 0x8b,
		0x7b, 0x2e, 0x75, 0xeb, 0x76, 0x48, 0x7d, 0x10, 0x19, 0x25, 0xfa, 0xd2,
		0xd4, 0x41, 0x41, 0xa9, 0x72, 0xea, 0x0e, 0x8b, 0x4b, 0x4f, 0x1c, 0xf5,
		0xf9, 0xe9, 0xcb, 0x58, 0x04, 0x2b, 0x26, 0x5a, 0x5a, 0x24, 0x49, 0x72,
		0x28, 0xc4, 0x8c, 0xf8, 0xf7, 0xd1, 0xa8, 0xa4, 0xe6, 0x68, 0x26, 0x2c,
		0x76, 0x08, 0xfb, 0x2a, 0x38, 0xa7, 0xd1, 0x0c, 0x59, 0x4c, 0x48, 0x11,
		0x41, 0x28, 0x75, 0x7d, 0x73, 0x7c, 0x0e, 0x27, 0x4f, 0x9f, 0xfe, 0xdf,
		0x0b, 0x9f, 0xbf, 0x07, 0x12, 0xdd, 0x4f, 0xf2, 0xff, 0x5d, 0x80, 0xfc,
		0xe5, 0xe5, 0x0f, 0xbf, 0x57, 0x58, 0x4c, 0x61, 0x66, 0xdc, 0x24, 0x2c,
		0xa2, 0xee, 0x87, 0x12, 0xb5, 0xba, 0xf8, 0x4a, 0x85, 0xee, 0xcf, 0x3a,
		0x45, 0xa9, 0xd6, 0x6c, 0xc6, 0xb0, 0xfc, 0x93, 0x6f, 0x14, 0x02, 0xc5,
		0x22, 0x33, 0x87, 0x2c, 0x28, 0x84, 0xa1, 0xc0, 0x9c, 0x17, 0x08, 0xd4,
		0x0e, 0xe5, 0xb5, 0x4f, 0x2f, 0x9f, 0x76, 0x8e, 0xdd, 0x60, 0x28, 0xf4,
		0x6b, 0xec, 0x72, 0x5a, 0x2a, 0xd7, 0x33, 0x0f, 0xb3, 0x41, 0x19, 0x72,
		0x72, 0x48, 0xf4, 0x62, 0xbc, 0x8e, 0x92, 0xf9, 0x7d, 0xcd, 0xa8, 0x23,
		0xa5, 0xb2, 0xef, 0xe5, 0xc9, 0x09, 0x9b, 0x33, 0xa4, 0xae, 0x97, 0x0e,
		0x2a, 0x29, 0xba, 0xcd, 0x1b, 0xa9, 0xd6, 0xfd, 0x75, 0x63, 0xaa, 0x07,
		0xe2, 0x5a, 0x59, 0x4c, 0x0e, 0xd5, 0x4c, 0x0f, 0x0b, 0xce, 0xae, 0xf2,
		0x6f, 0xd4, 0x88, 0x32, 0x8c, 0x7e, 0xdd, 0x50, 0x32, 0x8a, 0x74, 0x87,
		0x09, 0x5e, 0x0a, 0xb1, 0xa5, 0x0d, 0xe7, 0xf3, 0x1a, 0xf3, 0x1b, 0x52,
		0x8a, 0x55, 0x04, 0xc9, 0x84, 0x4c, 0x8c, 0x98, 0x87, 0x8e, 0xdd, 0xb4,
		0x34, 0x2a, 0x78, 0xf3, 0xa0, 0x0c, 0x1a, 0xd1, 0x45, 0x9b, 0xd0, 0x84,
		0x07, 0xf3, 0xf4, 0x18, 0x4e, 0xcc, 0x7c, 0x13, 0x24, 0x6d, 0xaf, 0xe2,
		0x5a, 0xb5, 0x82, 0x3a, 0x7b, 0x6f, 0x50, 0xee, 0x00, 0x99, 0xdd, 0x84,
		0x56, 0x29, 0x44, 0x38, 0x6b, 0x02, 0xbb, 0xaa, 0xf5, 0x7d, 0xe1, 0x14,
		0xfa, 0xfe, 0x13, 0x2b, 0xbc, 0x22, 0x5b, 0x5f, 0xd2, 0x50, 0x43, 0x03,
		0x9b, 0x67, 0x65, 0x09, 0x9d, 0x72, 0x52, 0x46, 0x15, 0x78, 0x87, 0x05,
		0xde, 0x95, 0x44, 0xd5, 0xd2, 0x94, 0x3b, 0xb8, 0x2c, 0x81, 0x4b, 0xf7,
		0x99, 0x05, 0xdd, 0x66, 0x82, 0xd3, 0x68, 0x54, 0x90, 0xeb, 0x7c, 0x27,
		0xb3, 0x3b, 0x1e, 0x55, 0xdc, 0xd5, 0x6d, 0x16, 0xfa, 0xf7, 0xb7, 0xed,
		0xdf, 0x79, 0x49, 0xc3, 0xd1, 0x20, 0x7d, 0x74, 0xf1, 0x35, 0x77, 0x6f,
		0xda, 0xcc, 0x77, 0x33, 0xbe, 0x55, 0xc4, 0x30, 0x4a, 0xae, 0xe9, 0x00,
		0xd9, 0xd2, 0x21, 0x99, 0xd4, 0xb6, 0xbe, 0xf7, 0xde, 0xb1, 0xe4, 0x9e,
		0x13, 0x2c, 0x2b, 0xef, 0x19, 0x0b, 0xbb, 0x35, 0xb8, 0xdb, 0xb9, 0x23,
		0xe5, 0x8f, 0x17, 0xea, 0x2d, 0xdd, 0x34, 0x33, 0xfb, 0xce, 0xc4, 0x03,
		0xea, 0x7c, 0x3f, 0x72, 0x40, 0x1c, 0xd3, 0x20, 0x27, 0x54, 0x75, 0xa4,
		0x6c, 0x77, 0x24, 0xfd, 0x58, 0xbd, 0xdb, 0x91, 0x1d, 0x50, 0xcd, 0xb4,
		0xc7, 0x77, 0x34, 0x5c, 0xb9, 0x50, 0x76, 0x68, 0x8c, 0xa9, 0x07, 0x6f,
		0xf8, 0x44, 0x85, 0x99, 0x46, 0xe8, 0x5f, 0xff, 0x3c, 0xd2, 0xeb, 0xd4,
		0xcf, 0x76, 0x25, 0x09, 0x13, 0xbc, 0x6f, 0x59, 0x3b, 0x0b, 0xd0, 0xdc,
		0xf1, 0x6c, 0xc6, 0xc8, 0x1f, 0xd3, 0x65, 0x06, 0xdb, 0xee, 0x59, 0xf0,
		0x16, 0x82, 0x65, 0x84, 0xcb, 0x44, 0xe1, 0x71, 0xb2, 0x0b, 0x94, 0xb7,
		0x28, 0x2b, 0x57, 0xbf, 0xd3, 0x8e, 0x4f, 0x95, 0xef, 0x8a, 0x60, 0x1c,
		0xe8, 0xa3, 0x8b, 0xa1, 0xb1, 0x06, 0x11, 0x88, 0xcf, 0x29, 0x80, 0xfc,
		0x8b, 0x03, 0xf6, 0x5d, 0x35, 0xed, 0xcc, 0x29, 0xa9, 0xff, 0xf1, 0xf5,
		0x73, 0xa6, 0xb6, 0x8e, 0x65, 0xf3, 0xf4, 0x69, 0x17, 0x0e, 0xb3, 0xb2,
		0x7c, 0x52, 0xe9, 0xbc, 0x53, 0xe9, 0xa3, 0xb5, 0xbf, 0x93, 0x36, 0xa0,
		0x55, 0xa6, 0x6e, 0x3b, 0x79, 0xa8, 0xd8, 0x5f, 0x69, 0x02, 0x74, 0x26,
		0x06, 0x59, 0xe0, 0x27, 0x4a, 0x6c, 0xdb, 0xed, 0x85, 0xaf, 0x59, 0x2c,
		0x27, 0x78, 0x3a, 0x6c, 0x77, 0x67, 0xee, 0x3a, 0x96, 0x47, 0x1f, 0x73,
		0xf1, 0x6c, 0x0f, 0xf0, 0x29, 0xf3, 0xc1, 0x7c, 0xe0, 0x5e, 0xbc, 0xf2,
		0x41, 0x7d, 0x24, 0x50, 0xef, 0xc3, 0x5d, 0x1b, 0x4e, 0xb3, 0xf4, 0xa6,
		0xb3, 0x92, 0x0a, 0xb6, 0xb9, 0x62, 0x2b, 0x1f, 0xce, 0xf4, 0xbf, 0xb7,
		0x8b, 0xac, 0xd0, 0xce, 0x8e, 0x70, 0x7b, 0x53, 0xd8, 0x7d, 0xa0, 0xc3,
		0x0f, 0xcb, 0xd7, 0x3f, 0xb4, 0xbe, 0xcf, 0xfd, 0x5c, 0x62, 0x4e, 0xe8,
		0x4e, 0x57, 0x07, 0xd8, 0x1f, 0x3f, 0x7a, 0x11, 0xbe, 0x7e, 0xc3, 0x56,
		0xec, 0xaa, 0xfb, 0x3a, 0xa2, 0x45, 0x4b, 0x63, 0xbe, 0xfd, 0x62, 0xfb,
		0x91, 0x66, 0xee, 0xb3, 0x09, 0xbb, 0x66, 0xb7, 0x49, 0xa5, 0x14, 0x4d,
		0x4b, 0x4c, 0x73, 0x1b, 0xb0, 0xd7, 0xef, 0x11, 0x28, 0x66, 0x36, 0xbd,
		0xf6, 0xdf, 0x8d, 0x36, 0xe9, 0x49, 0x72, 0x72, 0x92, 0x9c, 0xf6, 0xab,
		0xa3, 0x1f, 0x51, 0x48, 0xc0, 0x4b, 0x99, 0x8b, 0x96, 0x0a, 0x3f, 0xa3,
		0x02, 0x48, 0x9c, 0x34, 0x8d, 0xec, 0xc5, 0x20, 0x08, 0x7c, 0x4e, 0xc1,
		0xa0, 0xd6, 0x5f, 0x3c, 0x01, 0x92, 0x99, 0xf7, 0x84, 0x5c, 0x92, 0x8a,
		0xbc, 0x68, 0x29, 0x8e, 0xfc, 0x37, 0x25, 0xc2, 0x69, 0x2a, 0x7c, 0x88,
		0x05, 0x1d, 0xfb, 0xa8, 0xd8, 0xf7, 0xfd, 0x62, 0x77, 0xbd, 0xff, 0xc1,
		0x6e, 0x46, 0xf0, 0x29, 0x7f, 0xf2, 0x2f, 0x3a, 0x9b, 0x32, 0xad, 0x69,
		0x74, 0x60, 0xde, 0xa3, 0x07, 0x27, 0xc8, 0xa1, 0xe1, 0x93, 0xc8, 0x22,
		0xed, 0xbe, 0xf8, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x93, 0x1c,
		0xab, 0x02, 0x16, 0x00, 0x00,
	},
		"frontend/index.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"frontend/assets/application.coffee": frontend_assets_application_coffee,
	"frontend/assets/application.js": frontend_assets_application_js,
	"frontend/index.html": frontend_index_html,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"frontend": &_bintree_t{nil, map[string]*_bintree_t{
		"assets": &_bintree_t{nil, map[string]*_bintree_t{
			"application.js": &_bintree_t{frontend_assets_application_js, map[string]*_bintree_t{
			}},
			"application.coffee": &_bintree_t{frontend_assets_application_coffee, map[string]*_bintree_t{
			}},
		}},
		"index.html": &_bintree_t{frontend_index_html, map[string]*_bintree_t{
		}},
	}},
}}
