package bloom

import (
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"
)

func TestBoom(t *testing.T) {
	var loop int = 1e4
	boom := NewBoom(uint(loop), 0.0001)

	var count int
	start := time.Now()
	// 奇数加入黑名单
	for i := 2; i <= loop<<1; i += 2 {
		count++
		boom.Add(strconv.Itoa(i - 1))
	}
	fmt.Printf("%d条数据已经入黑名单, 耗时: %v\n", count, time.Since(start)) // 129.683855ms

	// 查询所有奇数
	for i := 2; i <= loop<<1; i += 2 {
		if !boom.Find(strconv.Itoa(i - 1)) {
			fmt.Println("失败 ---", i-1)
		}
	}

	// 查询所有偶数
	for i := 2; i <= loop<<1; i += 2 {
		if boom.Find(strconv.Itoa(i)) {
			fmt.Println("误判 ---", i)
		}
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

const (
	b   float64 = 8
	kib         = 1 << 10 * b
	mib         = 1 << 10 * kib
	gib         = 1 << 10 * mib
)

func TestKMP(t *testing.T) {
	var n = 1e1 // 样本量
	var p = 0.1 // 预期失误率

	// 根据样本量和预期失误率，求 k、m
	{
		// mb bit 大小
		// mk byte 大小
		var m, k float64
		m = math.Ceil(-1 * n * math.Log(p) / (math.Ln2 * math.Ln2))
		k = math.Ceil(math.Ln2 * m / n)
		fmt.Printf("样本量: %.f\n", n)
		fmt.Println("哈希函数:", k)
		fmt.Printf("内存占用: %.f\n", m)
		// fmt.Printf("内存占用: %.f GiB\n", math.Ceil(m/gib))
		fmt.Println("预期失误率:", p)
	}

	fmt.Println("+ + + + + + + + + +")

	// 根据实际内存大小和哈希函数个数，求真实失误率
	// p = (1 - e^(-1 * n * k / m))^k
	{
		var m = 32 * gib
		var k float64 = 14
		var p float64

		p = math.Pow((1 - math.Pow(math.E, (-1*n*k/m))), k)

		fmt.Printf("样本量: %.f\n", n)
		fmt.Println("哈希函数:", k)
		fmt.Printf("实际占用占用: %.f GiB\n", m/gib)
		fmt.Println("真实失误率:", p)
	}
}
