package bloom

import (
	"fmt"
	"testing"
)

var blacklist = []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期天"}

var testlist = []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期天", "中秋节", "国庆节", "元旦节"}

func TestBoom(t *testing.T) {
	boom := NewBoom(1e10, 0.0001)

	boom.Add(blacklist...)

	for _, val := range testlist {
		fmt.Printf("%s: %v\n", val, boom.Find(val))
	}
}

func TestBitset(t *testing.T) {
	var bitset bitset = []byte{0}

	// 标记 1 3 5 bit 位置
	coords := []*coord{
		{
			setIndex: 0,
			bitIndex: 1,
		},
		{
			setIndex: 0,
			bitIndex: 3,
		},
		{
			setIndex: 0,
			bitIndex: 5,
		},
		{
			setIndex: 0,
			bitIndex: 7,
		},
	}

	fmt.Println("- - - - - - - - 标记前 - - - - - - - -")
	// 查询
	for i := 0; i < 8; i++ {
		fmt.Printf("%d: %v\n", i, bitset.inquiry(&coord{setIndex: 0, bitIndex: i}))
	}

	// 标记
	for _, coord := range coords {
		bitset.mark(coord)
	}

	fmt.Println("- - - - - - - - 标记后 - - - - - - - -")
	// 查询
	for i := 0; i < 8; i++ {
		fmt.Printf("%d: %v\n", i, bitset.inquiry(&coord{setIndex: 0, bitIndex: i}))
	}
}
