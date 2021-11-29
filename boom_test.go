package bloom

import (
	"fmt"
	"math"
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

type bit float64

const (
	b   bit = 8
	kib     = 1 << 10 * b
	mib     = 1 << 10 * kib
	gib     = 1 << 10 * mib
)

func TestKMP(t *testing.T) {
	var n = 1e10   // 样本量
	var p = 0.0001 // 预期失误率

	var k float64 // hash 函数个数

	// 根据样本量和预期失误率，求 k、m
	{
		// mb bit 大小
		// mk byte 大小
		var m float64
		m = math.Ceil(-1 * n * math.Log(p) / (math.Ln2 * math.Ln2))
		k = math.Ceil(math.Ln2 * m / n)
		fmt.Printf("样本量: %.f\n", n)
		fmt.Println("哈希函数:", k)
		fmt.Printf("内存占用: %.f GiB\n", math.Ceil(m/float64(gib)))
		fmt.Println("预期失误率:", p)
	}
}
