package main

import (
	"fmt"
)

var k = []uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

func rotateRight(x, n uint32) uint32 {
	return (x >> n) | (x << (32 - n))
}

func sha256Padding(msg []byte) []byte {
	length := uint64(len(msg) * 8)
	msg = append(msg, 0x80)
	for len(msg)%64 != 56 {
		msg = append(msg, 0x00)
	}
	for i := 7; i >= 0; i-- {
		msg = append(msg, byte(length>>(uint(i)*8)))
	}
	return msg
}

func sha256(msg []byte) [32]byte {
	h := [8]uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
		0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
	}

	msg = sha256Padding(msg)
	for i := 0; i < len(msg); i += 64 {
		w := make([]uint32, 64)
		for j := 0; j < 16; j++ {
			w[j] = uint32(msg[i+4*j])<<24 | uint32(msg[i+4*j+1])<<16 | uint32(msg[i+4*j+2])<<8 | uint32(msg[i+4*j+3])
		}
		for j := 16; j < 64; j++ {
			s0 := rotateRight(w[j-15], 7) ^ rotateRight(w[j-15], 18) ^ (w[j-15] >> 3)
			s1 := rotateRight(w[j-2], 17) ^ rotateRight(w[j-2], 19) ^ (w[j-2] >> 10)
			w[j] = w[j-16] + s0 + w[j-7] + s1
		}

		a, b, c, d, e, f, g, h := h[0], h[1], h[2], h[3], h[4], h[5], h[6], h[7]
		for j := 0; j < 64; j++ {
			s1 := rotateRight(e, 6) ^ rotateRight(e, 11) ^ rotateRight(e, 25)
			ch := (e & f) ^ (^e & g)
			temp1 := h + s1 + ch + k[j] + w[j]
			s0 := rotateRight(a, 2) ^ rotateRight(a, 13) ^ rotateRight(a, 22)
			maj := (a & b) ^ (a & c) ^ (b & c)
			temp2 := s0 + maj

			h = g
			g = f
			f = e
			e = d + temp1
			d = c
			c = b
			b = a
			a = temp1 + temp2
		}

		h[0] += a
		h[1] += b
		h[2] += c
		h[3] += d
		h[4] += e
		h[5] += f
		h[6] += g
		h[7] += h
	}

	var digest [32]byte
	for i, v := range h {
		digest[4*i] = byte(v >> 24)
		digest[4*i+1] = byte(v >> 16)
		digest[4*i+2] = byte(v >> 8)
		digest[4*i+3] = byte(v)
	}
	return digest
}

func main() {
	message := "Hello, World!"
	hash := sha256([]byte(message))
	fmt.Printf("SHA-256 hash: %x\n", hash)
}
