package bloom

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	var list = make([]encoder, 10)
	for idx := range list {
		list[idx] = newHMAC(idx)
	}

	var data []byte = []byte("爱我中华")
	for idx, val := range list {
		fmt.Println(idx, val.encode(data))
	}
}

func TestPoint(t *testing.T) {
	var (
		// dividend 被除数
		dividend uint64 = 200
		// decimal 十进制
		decimal uint64 = 4294967295
		// binary 二进制 11111111 11111111 11111111 11111111
		binary bytes = []byte{255, 255, 255, 255}
	)
	fmt.Println("1 ==>", decimal%dividend)

	// 二进制取余
	{
		var surplus uint64
		for _, b := range binary {
			switch surplus {
			case 0:
				surplus = uint64(b) % dividend
			default:
				surplus = ((surplus << 8) | (uint64(b))) % dividend
			}
		}
		fmt.Println("2 ==>", surplus)
	}
}
