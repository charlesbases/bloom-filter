package bloom

import (
	"math"
	"unsafe"
)

type BoomFilter interface {
	Add(data ...string)
	Find(data string) bool
}

// boom .
type boom struct {
	// n 样本数量
	// p 允许的失误率
	// m bit 大小. m = -(n * lnp) / ((ln2)^2)
	// k 哈希函数个数. k = ln2 * (m / n)
	n, p float64
	m    uint64
	k    int

	// l []byte 长度
	// l = m / 8 (向上取整)
	l uint64

	bitset   bitset
	hashList []encoder
}

// NewBoom new boom-filter
// n: 样本数量, p: 允许的失误率
func NewBoom(n uint, p float64) BoomFilter {
	boom := &boom{n: float64(n), p: p}
	boom.init()
	return boom
}

func (boom *boom) init() {
	// 计算 m k l
	m := math.Ceil(-1 * boom.n * math.Log(boom.p) / (math.Ln2 * math.Ln2))
	boom.k = int(math.Ceil(math.Ln2 * m / boom.n))
	boom.l = uint64(math.Ceil(m / 8))
	boom.m = uint64(m)

	// 初始化 bitset
	boom.bitset = make([]byte, boom.l)

	// 初始化 hash
	boom.hashList = make([]encoder, boom.k)
	for idx := range boom.hashList {
		boom.hashList[idx] = newHMAC(idx)
	}
}

func (boom *boom) Add(v ...string) {
	for _, data := range v {
		for _, coord := range boom.coords(data) {
			boom.bitset.mark(coord)
		}
	}
}

func (boom *boom) Find(data string) bool {
	for _, coord := range boom.coords(data) {
		if boom.bitset.inquiry(coord) == true {
			return true
		}
	}
	return false
}

func (boom *boom) coords(data string) []*coord {
	var coords = make([]*coord, boom.k)
	for idx, hash := range boom.hashList {
		// point := hash.point(boom.m, *boom.decodeString(&data))
		point := hash.point(boom.m, []byte(data))
		coords[idx] = &coord{
			setIndex: int(math.Ceil(float64(point) / 8)),
			bitIndex: int(point % 8),
		}
	}
	return coords
}

func (boom *boom) encodeString(data *[]byte) *string {
	return (*string)(unsafe.Pointer(data))
}

func (boom *boom) decodeString(data *string) *[]byte {
	strPointer := (*[2]uintptr)(unsafe.Pointer(data))
	bytesPointer := [3]uintptr{strPointer[0], strPointer[1], strPointer[1]}
	return (*[]byte)(unsafe.Pointer(&bytesPointer))
}
