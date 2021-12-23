package bloom

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

var pool = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

var (
	// loop 测试数据数量
	loop uint = 1e6
	// 允许的失误率
	rate = 0.001

	// store 测试数据
	store = map[string]bool{}
	// strlen 字符串长度范围
	strlenMin = 10
	strlenMax = 20
)

func init() {
	rand.Seed(time.Now().Unix())
}

// _random 生成随机字符串
func _random() string {
	var (
		strlen  = rand.Intn(strlenMax-strlenMin) + strlenMin
		poollen = len(pool)
	)

	var builder strings.Builder
	builder.Grow(strlen)
	for i := 0; i < strlen; i++ {
		builder.WriteString(pool[rand.Intn(poollen)])
	}

	return builder.String()
}

// TestBoom
//  ==> 数据已添加:  1000000
//  ==> 数据校验成功:  1000000
//  ==> 随机数据校验中...
//    总数据量: 2000000
//    误判次数: 2018
//    既定失误率: 0.001
//    实际失误率: 0.00100900
func TestBoom(t *testing.T) {
	var boom = NewBoom(loop, rate)

	// 生成随机字符串并添加至布隆过滤器
	{
		for i := 0; i < int(loop); i++ {
			var str = _random()

			// store
			store[str] = true

			// boom.add
			boom.Add(str)
		}

		fmt.Println("  ==> 数据已添加: ", loop)
	}

	// 数据检验
	{
		for key := range store {
			if !boom.Find(key) {
				fmt.Println("  ==> 数据校验失败: ", key)
				os.Exit(1)
			}
		}

		fmt.Println("  ==> 数据校验成功: ", loop)
	}

	// 随机数据校验失误率
	{
		var sum = loop << 1
		var mistakes int
		for i := 0; i < int(sum); i++ {
			var str = _random()

			// 未添加至过滤器，但是检验通过
			if !store[str] && boom.Find(str) {
				mistakes++
			}
		}

		fmt.Println("  ==> 随机数据校验中...")
		fmt.Println("    总数据量:", sum)
		fmt.Println("    误判次数:", mistakes)
		fmt.Println("    既定失误率:", rate)
		fmt.Printf("    实际失误率: %.8f\n", float64(mistakes)/float64(sum))
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
