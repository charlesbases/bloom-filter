package bloom

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
	"sync"
)

// secretKeyLen 密钥长度, 高位补0
const secretKeyLen = 64

type encoder interface {
	// point bitset index
	point(m uint64, data []byte) uint64
	// encode 返回二进制
	encode(data []byte) []byte
	// encodeToString 返回十六进制
	encodeToString(data []byte) string
}

type hmacsha256 struct {
	key []byte
	h   hash.Hash

	lock sync.Mutex
}

// newHMAC .
func newHMAC(key int) encoder {
	h := new(hmacsha256)
	h.secretKey(key)
	h.h = hmac.New(sha256.New, h.key)
	return h
}

// secretKey number 转换为二进制，做为 hmac256 密钥
func (h *hmacsha256) secretKey(number int) {
	h.key = make([]byte, secretKeyLen)
	for idx := range h.key {
		h.key[idx] = '0'
	}

	if number > 0 {
		for i := secretKeyLen - 1; number > 0; number >>= 1 {
			h.key[i] = bin2Hex[uint8(number%2)]
			i--
		}
	}
}

func (h *hmacsha256) point(m uint64, data []byte) uint64 {
	var res bytes = h.encode(data)

	var surplus uint64
	// 二进制取余
	for _, b := range res {
		switch surplus {
		case 0:
			// 余数为 0 时，不需要处理，直接取余计算
			surplus = uint64(b) % m
		default:
			// 余数不为 0 时，需要拼接上当前 8 位二进制
			// 余数左移 8 位，或上当前二进制，再进行取余
			// eg: surplus = 11011, byte = 11110000
			// 1、11011 << 8 ==> 11011 00000000
			// 2、11011 00000000 | 11110000 ==> 11011 11110000
			// 3、(11011 11110000) % m
			surplus = ((surplus << 8) | (uint64(b))) % m
		}
	}
	return surplus
}

func (h *hmacsha256) encode(data []byte) []byte {
	h.lock.Lock()
	h.h.Write(data)
	var res bytes = h.h.Sum(nil)
	h.h.Reset()
	h.lock.Unlock()
	return res
}

func (h *hmacsha256) encodeToString(data []byte) string {
	var res bytes = h.encode(data)

	var length = res.len()
	var bs bytes = make([]byte, 0, length<<1)

	for i := 0; i < length; i++ {
		bs = append(bs, bin2Hex[(res)[i]>>4], bin2Hex[(res)[i]&0xf])
	}
	return bs.String()
}
