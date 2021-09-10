package main

const tilesetWidth = 96
const tilesetHeight = 88
const tilesetFlags = 1 // BLIT_2BPP
var tileset = [2112]byte{0x54, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x69, 0x20, 0x08, 0x0a, 0xa0, 0x00, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6a, 0x40, 0x01, 0xa9, 0x08, 0x20, 0x20, 0x08, 0x03, 0x55, 0x55, 0xc0, 0xf0, 0x00, 0x00, 0xf0, 0xff, 0xff, 0xff, 0xf0, 0xf0, 0x00, 0x00, 0xf0, 0x1a, 0x90, 0x06, 0xa4, 0x02, 0x80, 0x20, 0x08, 0x0d, 0x55, 0x55, 0x70, 0xf0, 0x00, 0x00, 0xf0, 0xff, 0xff, 0xff, 0xf0, 0xf0, 0x00, 0x00, 0xf0, 0x06, 0xa4, 0x1a, 0x90, 0x02, 0x80, 0x20, 0x08, 0x35, 0x55, 0x55, 0x5c, 0xff, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x0f, 0x00, 0x01, 0xa9, 0x6a, 0x40, 0x08, 0x20, 0x20, 0x08, 0x35, 0x55, 0x55, 0x5c, 0xff, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x0f, 0x00, 0x00, 0x6a, 0xa9, 0x00, 0x20, 0x08, 0x0a, 0xa0, 0xd5, 0x55, 0x55, 0x57, 0xf0, 0xf0, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x1a, 0xa4, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd5, 0x55, 0x55, 0x57, 0xf0, 0xf0, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x1a, 0xa4, 0x00, 0x0f, 0xff, 0x00, 0x00, 0x35, 0x55, 0x55, 0x5c, 0xf0, 0x0f, 0x00, 0xf0, 0xff, 0xff, 0xff, 0xf0, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x6a, 0xa9, 0x00, 0x0f, 0xff, 0x00, 0x00, 0x0d, 0x55, 0x55, 0x70, 0xf0, 0x0f, 0x00, 0xf0, 0xff, 0xff, 0xff, 0xf0, 0x00, 0x0f, 0x00, 0x00, 0x01, 0xa9, 0x6a, 0x40, 0x00, 0x00, 0x0f, 0xf0, 0x03, 0x55, 0x55, 0xc0, 0xf0, 0x00, 0xf0, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x06, 0xa4, 0x1a, 0x90, 0x00, 0x00, 0x0f, 0xf0, 0x03, 0x55, 0x55, 0xc0, 0xf0, 0x00, 0xf0, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x1a, 0x90, 0x06, 0xa4, 0x00, 0x00, 0x0f, 0xf0, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x0f, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x0f, 0x00, 0x6a, 0x40, 0x01, 0xa9, 0x00, 0x00, 0x0f, 0xf0, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x0f, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x0f, 0x00, 0x69, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0xff, 0xff, 0xff, 0xf0, 0xf0, 0x00, 0x00, 0xf0, 0x54, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0xff, 0xff, 0xff, 0xf0, 0xf0, 0x00, 0x00, 0xf0, 0x00, 0x05, 0x50, 0x00, 0xac, 0x00, 0x00, 0x35, 0x00, 0xd5, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5a, 0xa5, 0x00, 0xab, 0x00, 0x00, 0xd5, 0x00, 0xd5, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xaa, 0xaa, 0x40, 0xea, 0xc0, 0x03, 0x57, 0x00, 0xd5, 0x57, 0x00, 0x0f, 0xff, 0xff, 0x00, 0x0f, 0xff, 0xff, 0x00, 0xff, 0xff, 0xff, 0xf0, 0x06, 0xaa, 0xaa, 0x90, 0x3a, 0xb0, 0x0d, 0x5c, 0x00, 0xd5, 0x57, 0x00, 0x0f, 0xff, 0xff, 0x00, 0x0f, 0xff, 0xff, 0x00, 0xff, 0xff, 0xff, 0xf0, 0x1a, 0xa5, 0x5a, 0xa4, 0x0e, 0xac, 0x35, 0x70, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x1a, 0x90, 0x06, 0xa4, 0x03, 0xab, 0xd5, 0xc0, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x6a, 0x40, 0x01, 0xa9, 0x00, 0xea, 0x57, 0x00, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x6a, 0x40, 0x01, 0xa9, 0x00, 0x3a, 0x5c, 0x00, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x6a, 0x40, 0x01, 0xa9, 0x00, 0x3a, 0x5c, 0x00, 0x00, 0xd5, 0x57, 0x00, 0xff, 0xff, 0xff, 0xf0, 0x0f, 0xff, 0xff, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x6a, 0x40, 0x01, 0xa9, 0x00, 0xea, 0x57, 0x00, 0x00, 0xd5, 0x57, 0x00, 0xff, 0xff, 0xff, 0xf0, 0x0f, 0xff, 0xff, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x1a, 0x90, 0x06, 0xa4, 0x03, 0xab, 0xd5, 0xc0, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x0f, 0x00, 0x00, 0x1a, 0xa5, 0x5a, 0xa4, 0x0e, 0xac, 0x35, 0x70, 0x00, 0xd5, 0x57, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x0f, 0x00, 0x00, 0x06, 0xaa, 0xaa, 0x90, 0x3a, 0xb0, 0x0d, 0x5c, 0x03, 0x55, 0x55, 0xc0, 0xf0, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x0f, 0x00, 0x00, 0x01, 0xaa, 0xaa, 0x40, 0xea, 0xc0, 0x03, 0x57, 0x03, 0x55, 0x55, 0xc0, 0xf0, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x5a, 0xa5, 0x00, 0xab, 0x00, 0x00, 0xd5, 0x0d, 0x55, 0x55, 0x70, 0xf0, 0x00, 0x00, 0xf0, 0x0f, 0xff, 0xff, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x05, 0x50, 0x00, 0xac, 0x00, 0x00, 0x35, 0x35, 0x55, 0x55, 0x5c, 0xf0, 0x00, 0x00, 0xf0, 0x0f, 0xff, 0xff, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd5, 0x55, 0x55, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x00, 0x00, 0xf5, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x03, 0x55, 0x55, 0x55, 0x55, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0x5f, 0x00, 0x00, 0x00, 0x03, 0x55, 0x70, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x55, 0x5f, 0xf5, 0x55, 0x70, 0x00, 0x00, 0xfd, 0x7f, 0x00, 0x00, 0x0d, 0x55, 0xc0, 0x00, 0x00, 0x0d, 0x55, 0x5f, 0x00, 0x00, 0x00, 0x00, 0xf5, 0x55, 0xf0, 0x0f, 0x55, 0x5f, 0x00, 0xff, 0x55, 0x55, 0xff, 0x00, 0xf5, 0x55, 0x70, 0x00, 0x00, 0x35, 0x55, 0x55, 0xff, 0xff, 0xff, 0xff, 0x55, 0x57, 0x00, 0x00, 0xd5, 0x55, 0xff, 0x55, 0x55, 0x55, 0x55, 0xff, 0x55, 0x55, 0x5c, 0x00, 0x00, 0x35, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x57, 0x00, 0x00, 0xd5, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x5c, 0x00, 0x00, 0xd5, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x5c, 0x00, 0x00, 0x35, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x57, 0x00, 0x00, 0xd5, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x5c, 0x00, 0x00, 0x35, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x57, 0x00, 0x00, 0xd5, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x5c, 0x00, 0x00, 0x35, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x57, 0x00, 0x00, 0xd5, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x5c, 0x00, 0x00, 0x35, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x57, 0x00, 0x00, 0x35, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x57, 0x00, 0x00, 0xd5, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x5c, 0x00, 0x00, 0x35, 0x55, 0x55, 0xff, 0xff, 0xff, 0xff, 0x55, 0x57, 0x00, 0x00, 0xd5, 0x55, 0xff, 0x55, 0x55, 0x55, 0x55, 0xff, 0x55, 0x55, 0x5c, 0x00, 0x00, 0x0d, 0x55, 0x5f, 0x00, 0x00, 0x00, 0x00, 0xf5, 0x55, 0xf0, 0x0f, 0x55, 0x5f, 0x00, 0xff, 0x55, 0x55, 0xff, 0x00, 0xf5, 0x55, 0x70, 0x00, 0x00, 0x03, 0x55, 0x70, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x55, 0x5f, 0xf5, 0x55, 0x70, 0x00, 0x00, 0xfd, 0x7f, 0x00, 0x00, 0x0d, 0x55, 0xc0, 0x00, 0x00, 0x00, 0xf5, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x03, 0x55, 0x55, 0x55, 0x55, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0x5f, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd5, 0x55, 0x55, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x00, 0xfc, 0x00, 0x00, 0x3f, 0x00, 0x3a, 0xac, 0x00, 0x35, 0x55, 0x55, 0x5c, 0xfc, 0x00, 0x00, 0x3f, 0x00, 0x35, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x00, 0xeb, 0x00, 0x00, 0xeb, 0x00, 0xea, 0xab, 0x00, 0x0d, 0x55, 0x55, 0x70, 0xd7, 0x00, 0x00, 0xd7, 0x00, 0xd5, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0xea, 0xc0, 0x03, 0xab, 0x03, 0xab, 0xea, 0xc0, 0x03, 0x55, 0x55, 0xc0, 0xd5, 0xc0, 0x03, 0x57, 0x03, 0x57, 0xd5, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x3a, 0xb0, 0x0e, 0xac, 0x0e, 0xac, 0x3a, 0xb0, 0x03, 0x55, 0x55, 0xc0, 0x35, 0x70, 0x0d, 0x5c, 0x0d, 0x5c, 0x35, 0x70, 0x00, 0x00, 0x00, 0x00, 0x0e, 0xac, 0x3a, 0xb0, 0x3a, 0xb0, 0x0e, 0xac, 0x00, 0xd5, 0x57, 0x00, 0x0d, 0x5c, 0x35, 0x70, 0x35, 0x70, 0x0d, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x03, 0xab, 0xea, 0xc0, 0xea, 0xc0, 0x03, 0xab, 0x00, 0xd5, 0x57, 0x00, 0x03, 0x57, 0xd5, 0xc0, 0xd5, 0xc0, 0x03, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0xea, 0xab, 0x00, 0xeb, 0x00, 0x00, 0xeb, 0x00, 0xd5, 0x57, 0x00, 0x00, 0xd5, 0x57, 0x00, 0xd7, 0x00, 0x00, 0xd7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3a, 0xac, 0x00, 0xfc, 0x00, 0x00, 0x3f, 0x00, 0xd5, 0x57, 0x00, 0x00, 0x35, 0x5c, 0x00, 0xfc, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0xfc, 0x00, 0x00, 0x00, 0x03, 0x55, 0x55, 0xc0, 0x00, 0x00, 0x00, 0x3f, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xea, 0xab, 0xf0, 0x00, 0x00, 0x03, 0x55, 0x55, 0xc0, 0x00, 0x00, 0x0f, 0xd5, 0x57, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfa, 0xaa, 0xaa, 0xaf, 0x00, 0x00, 0x03, 0x55, 0x55, 0xc0, 0x00, 0x00, 0xf5, 0x55, 0x55, 0x5f, 0x00, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0x00, 0x03, 0xaa, 0xaa, 0xaa, 0xaa, 0xc0, 0x00, 0x03, 0x55, 0x55, 0xc0, 0x00, 0x03, 0x55, 0x55, 0x55, 0x55, 0xc0, 0x00, 0xf0, 0x00, 0x00, 0xf0, 0x00, 0x0e, 0xaa, 0xbf, 0xfe, 0xaa, 0xb0, 0x00, 0x0d, 0x55, 0x55, 0x70, 0x00, 0x0d, 0x55, 0x7f, 0xfd, 0x55, 0x70, 0x00, 0xf0, 0x00, 0x0f, 0xf0, 0x00, 0x3a, 0xab, 0xc0, 0x03, 0xea, 0xac, 0x00, 0x0d, 0x55, 0x55, 0x70, 0x00, 0x35, 0x57, 0xc0, 0x03, 0xd5, 0x5c, 0x00, 0xf0, 0x00, 0x0f, 0xf0, 0x00, 0xea, 0xbc, 0x00, 0x00, 0x3e, 0xab, 0x00, 0x0d, 0x55, 0x55, 0x70, 0x00, 0xd5, 0x7c, 0x00, 0x00, 0x3d, 0x57, 0x00, 0xf0, 0x00, 0xf0, 0xf0, 0x03, 0xaa, 0xc0, 0x00, 0x00, 0x03, 0xaa, 0xc0, 0x35, 0x55, 0x55, 0x5c, 0x03, 0x55, 0xc0, 0x00, 0x00, 0x03, 0x55, 0xc0, 0xf0, 0x00, 0xf0, 0xf0, 0x0e, 0xab, 0x00, 0x00, 0x00, 0x00, 0xea, 0xb0, 0x35, 0x55, 0x55, 0x5c, 0x0d, 0x57, 0x00, 0x00, 0x00, 0x00, 0xd5, 0x70, 0xf0, 0x0f, 0x00, 0xf0, 0x0e, 0xac, 0x00, 0x00, 0x00, 0x00, 0x3a, 0xb0, 0x0d, 0x55, 0x55, 0x70, 0x0d, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x35, 0x70, 0xf0, 0x0f, 0x00, 0xf0, 0x3a, 0xac, 0x00, 0x00, 0x00, 0x00, 0x3a, 0xac, 0x0d, 0x55, 0x55, 0x70, 0x35, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x35, 0x5c, 0xf0, 0xf0, 0x00, 0xf0, 0x3a, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x0e, 0xac, 0x0d, 0x55, 0x55, 0x70, 0x35, 0x70, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x5c, 0xf0, 0xf0, 0x00, 0xf0, 0x3a, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x0e, 0xac, 0x03, 0x55, 0x55, 0xc0, 0x35, 0x70, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x5c, 0xff, 0x00, 0x00, 0xf0, 0xea, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x03, 0xab, 0x03, 0x55, 0x55, 0xc0, 0xd5, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x03, 0x57, 0xff, 0x00, 0x00, 0xf0, 0xea, 0xc0, 0x00, 0x03, 0xc0, 0x00, 0x03, 0xab, 0x03, 0x55, 0x55, 0xc0, 0xd5, 0xc0, 0x00, 0x00, 0xf0, 0x00, 0x03, 0x57, 0xf0, 0x00, 0x00, 0xf0, 0xea, 0xc0, 0x00, 0x0e, 0xb0, 0x00, 0x03, 0xab, 0x03, 0x55, 0x55, 0xc0, 0xd5, 0xc0, 0x00, 0x03, 0x5c, 0x00, 0x03, 0x57, 0xf0, 0x00, 0x00, 0xf0, 0xea, 0xc0, 0x00, 0x0e, 0xb0, 0x00, 0x03, 0xab, 0x00, 0xd5, 0x57, 0x00, 0xd5, 0xc0, 0x00, 0x03, 0x5c, 0x00, 0x03, 0x57, 0x00, 0x00, 0x00, 0x00, 0xea, 0xc0, 0x00, 0x03, 0xc0, 0x00, 0x03, 0xab, 0x00, 0xd5, 0x57, 0x00, 0xd5, 0xc0, 0x00, 0x00, 0xf0, 0x00, 0x03, 0x57, 0x00, 0x00, 0x00, 0x00, 0xea, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x03, 0xab, 0x00, 0xd5, 0x57, 0x00, 0xd5, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x03, 0x57, 0xf0, 0x00, 0x00, 0xf0, 0x3a, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x0e, 0xac, 0x00, 0xd5, 0x57, 0x00, 0x35, 0x70, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x5c, 0xf0, 0x00, 0x00, 0xf0, 0x3a, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x0e, 0xac, 0x03, 0x55, 0x55, 0xc0, 0x35, 0x70, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x5c, 0xf0, 0x00, 0x00, 0xf0, 0x3a, 0xac, 0x00, 0x00, 0x00, 0x00, 0x3a, 0xac, 0x03, 0x55, 0x55, 0xc0, 0x35, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x35, 0x5c, 0xf0, 0x00, 0x00, 0xf0, 0x0e, 0xac, 0x00, 0x00, 0x00, 0x00, 0x3a, 0xb0, 0x0d, 0x55, 0x55, 0x70, 0x0d, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x35, 0x70, 0xf0, 0x00, 0x00, 0xf0, 0x0e, 0xab, 0x00, 0x00, 0x00, 0x00, 0xea, 0xb0, 0x35, 0x55, 0x55, 0x5c, 0x0d, 0x57, 0x00, 0x00, 0x00, 0x00, 0xd5, 0x70, 0xf0, 0x00, 0x00, 0xf0, 0x03, 0xaa, 0xc0, 0x00, 0x00, 0x03, 0xaa, 0xc0, 0xd5, 0x55, 0x55, 0x57, 0x03, 0x55, 0xc0, 0x00, 0x00, 0x03, 0x55, 0xc0, 0xff, 0xff, 0xff, 0xf0, 0x00, 0xea, 0xbc, 0x00, 0x00, 0x3e, 0xab, 0x00, 0xd5, 0x55, 0x55, 0x57, 0x00, 0xd5, 0x7c, 0x00, 0x00, 0x3d, 0x57, 0x00, 0xff, 0xff, 0xff, 0xf0, 0x00, 0x3a, 0xab, 0xc0, 0x03, 0xea, 0xac, 0x00, 0x35, 0x55, 0x55, 0x5c, 0x00, 0x35, 0x57, 0xc0, 0x03, 0xd5, 0x5c, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x0e, 0xaa, 0xbf, 0xfe, 0xaa, 0xb0, 0x00, 0x35, 0x55, 0x55, 0x5c, 0x00, 0x0d, 0x55, 0x7f, 0xfd, 0x55, 0x70, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x03, 0xaa, 0xaa, 0xaa, 0xaa, 0xc0, 0x00, 0x0d, 0x55, 0x55, 0x70, 0x00, 0x03, 0x55, 0x55, 0x55, 0x55, 0xc0, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x00, 0xfa, 0xaa, 0xaa, 0xaf, 0x00, 0x00, 0x03, 0x55, 0x55, 0xc0, 0x00, 0x00, 0xf5, 0x55, 0x55, 0x5f, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x0f, 0xea, 0xab, 0xf0, 0x00, 0x00, 0x00, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0x0f, 0xd5, 0x57, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x00, 0x3f, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x3f, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0}