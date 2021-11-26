package bloom

var bits = [8]uint8{
	0x1 << 7 /* 10000000 */, 0x1 << 6 /* 01000000 */, 0x1 << 5 /* 00100000 */, 0x1 << 4, /* 00010000 */
	0x1 << 3 /* 00001000 */, 0x1 << 2 /* 00000100 */, 0x1 << 1 /* 00000010 */, 0x1, /* 00000001 */
}

// bitset 数据集合
type bitset []byte

// coord .
type coord struct {
	// setIndex bitset index
	setIndex int
	// bitIndex bit index in byte (0 ~ 7)
	bitIndex int
}

// mark 标记
func (bit *bitset) mark(c *coord) {
	(*bit)[c.setIndex] |= bits[c.bitIndex]
}

// inquiry 查询位置是否标记
func (bit *bitset) inquiry(c *coord) bool {
	return (*bit)[c.setIndex]&bits[c.bitIndex] == bits[c.bitIndex]
}
