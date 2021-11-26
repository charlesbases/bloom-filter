package bloom

import "unsafe"

// bin2Hex 二进制转为十六进制
var bin2Hex = map[uint8]byte{
	0x0: '0', 0x1: '1', 0x2: '2', 0x3: '3', 0x4: '4', 0x5: '5', 0x6: '6', 0x7: '7',
	0x8: '8', 0x9: '9', 0xa: 'a', 0xb: 'b', 0xc: 'c', 0xd: 'd', 0xe: 'e', 0xf: 'f',
}

type bytes []byte

// len .
func (bs *bytes) len() int {
	return len(*bs)
}

// String .
func (bs *bytes) String() string {
	return *(*string)(unsafe.Pointer(bs))
}
