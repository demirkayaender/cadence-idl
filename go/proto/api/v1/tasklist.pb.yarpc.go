// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/api/v1/tasklist.proto

package apiv1

var yarpcFileDescriptorClosure216fa006947e00a0 = [][]byte{
	// uber/cadence/api/v1/tasklist.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xeb, 0x4e, 0xe3, 0x46,
		0x14, 0x6e, 0x6e, 0x5c, 0x4e, 0x16, 0x30, 0x03, 0x2c, 0x49, 0xb6, 0xdb, 0xb2, 0xf9, 0x81, 0x28,
		0x6a, 0x1d, 0x41, 0x5b, 0xa9, 0x6a, 0xab, 0xed, 0x06, 0x82, 0x76, 0x2d, 0x2e, 0x8b, 0x1c, 0x2f,
		0x15, 0x95, 0x2a, 0x77, 0x62, 0x0f, 0x61, 0xea, 0xcb, 0x58, 0x9e, 0x71, 0x02, 0x4f, 0xd0, 0x37,
		0xe8, 0xc3, 0xf4, 0x1d, 0xfa, 0x4e, 0xd5, 0x8c, 0x9d, 0x90, 0x8b, 0x41, 0xdd, 0x1f, 0xfb, 0x2f,
		0x73, 0xce, 0x7c, 0xf3, 0x9d, 0xef, 0xdc, 0x62, 0x68, 0x26, 0x3d, 0x12, 0xb7, 0x1c, 0xec, 0x92,
		0xd0, 0x21, 0x2d, 0x1c, 0xd1, 0xd6, 0xe0, 0xa0, 0x25, 0x30, 0xf7, 0x7c, 0xca, 0x85, 0x1e, 0xc5,
		0x4c, 0x30, 0xb4, 0x21, 0xef, 0xe8, 0xd9, 0x1d, 0x1d, 0x47, 0x54, 0x1f, 0x1c, 0x34, 0xbe, 0xe8,
		0x33, 0xd6, 0xf7, 0x49, 0x4b, 0x5d, 0xe9, 0x25, 0x37, 0x2d, 0x37, 0x89, 0xb1, 0xa0, 0x2c, 0x4c,
		0x41, 0x8d, 0x2f, 0x67, 0xfd, 0x82, 0x06, 0x84, 0x0b, 0x1c, 0x44, 0xd9, 0x85, 0xb9, 0x07, 0x86,
		0x31, 0x8e, 0x22, 0x12, 0xf3, 0xd4, 0xdf, 0xfc, 0x00, 0x4b, 0x16, 0xe6, 0xde, 0x19, 0xe5, 0x02,
		0x21, 0x28, 0x87, 0x38, 0x20, 0xb5, 0xc2, 0x4e, 0x61, 0x6f, 0xd9, 0x54, 0xbf, 0xd1, 0xf7, 0x50,
		0xf6, 0x68, 0xe8, 0xd6, 0x8a, 0x3b, 0x85, 0xbd, 0xd5, 0xc3, 0x57, 0x7a, 0x4e, 0x90, 0xfa, 0xe8,
		0x81, 0x53, 0x1a, 0xba, 0xa6, 0xba, 0xde, 0xc4, 0xa0, 0x8d, 0xac, 0xe7, 0x44, 0x60, 0x17, 0x0b,
		0x8c, 0xce, 0x61, 0x33, 0xc0, 0x77, 0xb6, 0x94, 0xcd, 0xed, 0x88, 0xc4, 0x36, 0x27, 0x0e, 0x0b,
		0x5d, 0x45, 0x57, 0x3d, 0xfc, 0x5c, 0x4f, 0x23, 0xd5, 0x47, 0x91, 0xea, 0x1d, 0x96, 0xf4, 0x7c,
		0x72, 0x85, 0xfd, 0x84, 0x98, 0xeb, 0x01, 0xbe, 0x93, 0x0f, 0xf2, 0x4b, 0x12, 0x77, 0x15, 0xac,
		0xf9, 0x01, 0xea, 0x23, 0x8a, 0x4b, 0x1c, 0x0b, 0x2a, 0xb3, 0x32, 0xe6, 0xd2, 0xa0, 0xe4, 0x91,
		0xfb, 0x4c, 0x89, 0xfc, 0x89, 0x76, 0x61, 0x8d, 0x0d, 0x43, 0x12, 0xdb, 0xb7, 0x8c, 0x0b, 0x5b,
		0xe9, 0x2c, 0x2a, 0xef, 0x8a, 0x32, 0xbf, 0x63, 0x5c, 0x5c, 0xe0, 0x80, 0x34, 0x3d, 0xd8, 0x32,
		0x38, 0xf3, 0x55, 0x92, 0xdf, 0xc6, 0x2c, 0x89, 0xce, 0x89, 0x88, 0xa9, 0xc3, 0x51, 0x0b, 0x36,
		0x43, 0x32, 0xcc, 0x0f, 0xbf, 0x60, 0xae, 0x87, 0x64, 0x38, 0x1d, 0x20, 0x7a, 0x05, 0xcf, 0x22,
		0xe6, 0xfb, 0x24, 0xb6, 0x1d, 0x96, 0x84, 0x42, 0xd1, 0x95, 0xcc, 0x6a, 0x6a, 0x3b, 0x96, 0xa6,
		0xe6, 0x5f, 0x65, 0x58, 0x1d, 0x89, 0xe8, 0x0a, 0x2c, 0x12, 0x8e, 0xbe, 0x06, 0xd4, 0xc3, 0x8e,
		0xe7, 0xb3, 0x7e, 0x0a, 0xb3, 0x6f, 0x69, 0x28, 0x14, 0x49, 0xc9, 0xd4, 0x32, 0x8f, 0x02, 0xbf,
		0xa3, 0xa1, 0x40, 0x2f, 0x01, 0x62, 0x82, 0x5d, 0xdb, 0x27, 0x03, 0xe2, 0x67, 0x0c, 0xcb, 0xd2,
		0x72, 0x26, 0x0d, 0xe8, 0x05, 0x2c, 0x63, 0xc7, 0xcb, 0xbc, 0x25, 0xe5, 0x5d, 0xc2, 0x8e, 0x97,
		0x3a, 0x77, 0x61, 0x2d, 0xc6, 0x82, 0x4c, 0x6a, 0x29, 0x2b, 0x2d, 0x2b, 0xd2, 0xfc, 0xa0, 0xa3,
		0x03, 0x2b, 0x52, 0xb4, 0x4d, 0x5d, 0xbb, 0xe7, 0x33, 0xc7, 0xab, 0x55, 0x54, 0xc1, 0x76, 0x1e,
		0xed, 0x05, 0xa3, 0x73, 0x24, 0xef, 0x99, 0x55, 0x09, 0x33, 0x5c, 0x75, 0x40, 0x03, 0xd8, 0xa6,
		0xa3, 0xbc, 0xda, 0x7d, 0x99, 0x58, 0x3b, 0x48, 0x33, 0x5b, 0x5b, 0xd8, 0x29, 0xed, 0x55, 0x0f,
		0x5f, 0x3f, 0xd9, 0x5b, 0x69, 0x76, 0xf4, 0xdc, 0xd2, 0x9c, 0x84, 0x22, 0xbe, 0x37, 0xb7, 0xe8,
		0x47, 0x95, 0x6d, 0xf1, 0x91, 0xb2, 0x35, 0x04, 0x34, 0x1e, 0x67, 0xc9, 0x69, 0xac, 0x37, 0x50,
		0x19, 0xc8, 0x1e, 0x55, 0xd9, 0xaf, 0x1e, 0xee, 0xe7, 0xca, 0xc8, 0x7d, 0xd1, 0x4c, 0x81, 0x3f,
		0x16, 0x7f, 0x28, 0x34, 0x7f, 0x81, 0xea, 0x44, 0xea, 0x50, 0x1d, 0x96, 0xb8, 0xc0, 0xb1, 0xb0,
		0xa9, 0x9b, 0xd5, 0x7e, 0x51, 0x9d, 0x0d, 0x17, 0x6d, 0xc1, 0x02, 0x09, 0x5d, 0xe9, 0x48, 0xcb,
		0x5d, 0x21, 0xa1, 0x6b, 0xb8, 0xcd, 0xbf, 0x0b, 0x00, 0x97, 0xaa, 0xb5, 0x8c, 0xf0, 0x86, 0xa1,
		0x0e, 0x68, 0x3e, 0xe6, 0xc2, 0xc6, 0x8e, 0x43, 0x38, 0xb7, 0xe5, 0x5a, 0xc8, 0x06, 0xad, 0x31,
		0x37, 0x68, 0xd6, 0x68, 0x67, 0x98, 0xab, 0x12, 0xd3, 0x56, 0x10, 0x69, 0x44, 0x0d, 0x58, 0xa2,
		0x2e, 0x09, 0x05, 0x15, 0xf7, 0xd9, 0xb4, 0x8c, 0xcf, 0x79, 0xed, 0x53, 0xca, 0x69, 0x9f, 0xe6,
		0x3f, 0x05, 0xa8, 0x77, 0x05, 0x75, 0xbc, 0xfb, 0x93, 0x3b, 0xe2, 0x24, 0x32, 0x09, 0x6d, 0x21,
		0x62, 0xda, 0x4b, 0x04, 0xe1, 0xe8, 0x2d, 0x68, 0x43, 0x16, 0x7b, 0x24, 0x56, 0x15, 0xb2, 0xe5,
		0x3e, 0xcc, 0xe2, 0x7c, 0xf9, 0x64, 0x3f, 0x98, 0xab, 0x29, 0x6c, 0xbc, 0xbc, 0x2c, 0xa8, 0x73,
		0xe7, 0x96, 0xb8, 0x89, 0x4f, 0x6c, 0xc1, 0xec, 0x34, 0x7b, 0x52, 0x36, 0x4b, 0x44, 0x56, 0x9a,
		0xfa, 0xfc, 0x8a, 0xc9, 0xb6, 0xa9, 0xf9, 0x7c, 0x84, 0xb5, 0x58, 0x57, 0x22, 0xad, 0x14, 0xd8,
		0x7c, 0x0d, 0xeb, 0x73, 0x4b, 0x06, 0x7d, 0x05, 0xda, 0x4c, 0x2b, 0xf3, 0x5a, 0x61, 0xa7, 0xb4,
		0xb7, 0x6c, 0xae, 0x4d, 0xf7, 0x20, 0x6f, 0xfe, 0x5b, 0x86, 0xed, 0xb9, 0x07, 0x8e, 0x59, 0x78,
		0x43, 0xfb, 0xa8, 0x06, 0x8b, 0x03, 0x12, 0x73, 0xca, 0xc2, 0x51, 0x89, 0xb3, 0x23, 0x3a, 0x84,
		0x8d, 0x30, 0x09, 0x6c, 0x35, 0xd9, 0xd1, 0x08, 0xc5, 0x95, 0x8a, 0xca, 0x51, 0xb1, 0x26, 0xdb,
		0x36, 0x09, 0x4c, 0x82, 0xdd, 0xf1, 0x93, 0x1c, 0x7d, 0x07, 0x9b, 0x12, 0x33, 0x8c, 0xa9, 0xac,
		0xc9, 0x03, 0xa8, 0x34, 0x06, 0xa1, 0x30, 0x09, 0x7e, 0x95, 0xee, 0x09, 0x14, 0x85, 0xb5, 0x59,
		0x96, 0xb2, 0x9a, 0xc6, 0x37, 0x4f, 0x66, 0x7f, 0x46, 0x8a, 0x3e, 0x1d, 0x4b, 0x3a, 0x8f, 0xab,
		0xf1, 0x74, 0x80, 0x3e, 0x68, 0x73, 0xc1, 0x55, 0x14, 0x57, 0xfb, 0xa3, 0xb8, 0x66, 0x24, 0xa4,
		0x64, 0x6b, 0xc3, 0x69, 0x6b, 0x83, 0xc2, 0x46, 0x4e, 0x50, 0x93, 0xe3, 0x5b, 0x49, 0xc7, 0xf7,
		0xe7, 0xe9, 0xf1, 0xdd, 0xfd, 0x7f, 0xb1, 0x4c, 0x8c, 0x6e, 0xe3, 0x4f, 0xd8, 0xcc, 0x8b, 0xe9,
		0x53, 0x70, 0xed, 0xff, 0x01, 0xcf, 0x26, 0xff, 0x6d, 0x51, 0x03, 0x9e, 0x5b, 0xed, 0xee, 0xa9,
		0x7d, 0x66, 0x74, 0x2d, 0xfb, 0xd4, 0xb8, 0xe8, 0xd8, 0xc6, 0xc5, 0x55, 0xfb, 0xcc, 0xe8, 0x68,
		0x9f, 0xa1, 0x3a, 0x6c, 0xcd, 0xf8, 0x2e, 0xde, 0x9b, 0xe7, 0xed, 0x33, 0xad, 0x90, 0xe3, 0xea,
		0x5a, 0xc6, 0xf1, 0xe9, 0xb5, 0x56, 0xdc, 0x77, 0x1f, 0x18, 0xac, 0xfb, 0x88, 0x4c, 0x33, 0x58,
		0xd7, 0x97, 0x27, 0x13, 0x0c, 0x2f, 0x60, 0x7b, 0xc6, 0xd7, 0x39, 0x39, 0x36, 0xba, 0xc6, 0xfb,
		0x0b, 0xad, 0x90, 0xe3, 0x6c, 0x1f, 0x5b, 0xc6, 0x95, 0x61, 0x5d, 0x6b, 0xc5, 0xa3, 0xdf, 0x61,
		0xdb, 0x61, 0x41, 0x9e, 0xfe, 0xa3, 0x95, 0x71, 0x02, 0xe4, 0x94, 0x5e, 0x16, 0x7e, 0x3b, 0xe8,
		0x53, 0x71, 0x9b, 0xf4, 0x74, 0x87, 0x05, 0xad, 0xc9, 0xef, 0xa8, 0x6f, 0xa8, 0xeb, 0xb7, 0xfa,
		0x2c, 0xfd, 0xb4, 0xc9, 0x3e, 0xaa, 0x7e, 0xc2, 0x11, 0x1d, 0x1c, 0xf4, 0x16, 0x94, 0xed, 0xdb,
		0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xf6, 0xce, 0x1e, 0x78, 0x09, 0x00, 0x00,
	},
	// google/protobuf/duration.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
		0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
		0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
		0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
		0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xd4, 0xcc, 0xc8, 0x25, 0x9c,
		0x9c, 0x9f, 0xab, 0x87, 0x66, 0xa6, 0x13, 0x2f, 0xcc, 0xc4, 0x00, 0x90, 0x48, 0x00, 0x63, 0x94,
		0x21, 0x54, 0x45, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc2, 0x81, 0x25,
		0x95, 0x05, 0xa9, 0xc5, 0xfa, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x70, 0xc7, 0x16, 0x24, 0xfd, 0x60,
		0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0xa2, 0x39, 0x00, 0xaa,
		0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4, 0x35, 0x89, 0x0d, 0x6c, 0x94,
		0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xb4, 0xc3, 0xfb, 0x00, 0x00, 0x00,
	},
	// google/protobuf/timestamp.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d,
		0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x03, 0x0b, 0x09, 0xf1, 0x43, 0x14, 0xe8, 0xc1, 0x14, 0x28,
		0x59, 0x73, 0x71, 0x86, 0xc0, 0xd4, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5,
		0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x79, 0x89,
		0x79, 0xf9, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x53, 0x2b, 0x23, 0x97,
		0x70, 0x72, 0x7e, 0xae, 0x1e, 0x9a, 0xa1, 0x4e, 0x7c, 0x70, 0x23, 0x03, 0x40, 0x42, 0x01, 0x8c,
		0x51, 0x46, 0x50, 0x25, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0x48, 0x6e,
		0xac, 0x2c, 0x48, 0x2d, 0xd6, 0xcf, 0xce, 0xcb, 0x2f, 0xcf, 0x43, 0xb8, 0xb7, 0x20, 0xe9, 0x07,
		0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x00, 0xa7, 0x55, 0x4c, 0x72, 0xee, 0x10, 0xdd, 0x01, 0x50,
		0x2d, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x20, 0x0d, 0x21, 0x20, 0xbd, 0x49, 0x6c, 0x60, 0xb3,
		0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x65, 0xce, 0x7d, 0xff, 0x00, 0x00, 0x00,
	},
	// google/protobuf/wrappers.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2f, 0x4a, 0x2c,
		0x28, 0x48, 0x2d, 0x2a, 0xd6, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0xca,
		0x5c, 0xdc, 0x2e, 0xf9, 0xa5, 0x49, 0x39, 0xa9, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x42, 0x22, 0x5c,
		0xac, 0x65, 0x20, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x84, 0xa3, 0xa4, 0xc4, 0xc5,
		0xe5, 0x96, 0x93, 0x9f, 0x58, 0x82, 0x45, 0x0d, 0x13, 0x92, 0x1a, 0xcf, 0xbc, 0x12, 0x33, 0x13,
		0x2c, 0x6a, 0x98, 0x61, 0x6a, 0x94, 0xb9, 0xb8, 0x43, 0x71, 0x29, 0x62, 0x41, 0x35, 0xc8, 0xd8,
		0x08, 0x8b, 0x1a, 0x56, 0x34, 0x83, 0xb0, 0x2a, 0xe2, 0x85, 0x29, 0x52, 0xe4, 0xe2, 0x74, 0xca,
		0xcf, 0xcf, 0xc1, 0xa2, 0x84, 0x03, 0xc9, 0x9c, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x2c, 0x8a,
		0x38, 0x91, 0x1c, 0xe4, 0x54, 0x59, 0x92, 0x5a, 0x8c, 0x45, 0x0d, 0x0f, 0x54, 0x8d, 0x53, 0x33,
		0x23, 0x97, 0x70, 0x72, 0x7e, 0xae, 0x1e, 0x5a, 0xf0, 0x3a, 0xf1, 0x86, 0x43, 0xc3, 0x3f, 0x00,
		0x24, 0x12, 0xc0, 0x18, 0x65, 0x08, 0x55, 0x91, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f,
		0x94, 0x8e, 0x88, 0xab, 0x92, 0xca, 0x82, 0xd4, 0x62, 0xfd, 0xec, 0xbc, 0xfc, 0xf2, 0x3c, 0x78,
		0xbc, 0x15, 0x24, 0xfd, 0x60, 0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce,
		0x1d, 0xa2, 0x39, 0x00, 0xaa, 0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4,
		0x35, 0x89, 0x0d, 0x6c, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x92, 0x48, 0x30, 0x06,
		0x02, 0x00, 0x00,
	},
}
