package main

import (
	"strings"
)

// https://www.compart.com/en/unicode/block/U+2580

const UpperLeftBit = 1
const UpperRightBit = 2
const LowerLeftBit = 4
const LowerRightBit = 8

var bitToBlockTable = []rune{
	' ', '\u2598', '\u259D', '\u2580',
	'\u2596', '\u258C', '\u259E', '\u259B',
	'\u2597', '\u259A', '\u2590', '\u259C',
	'\u2584', '\u2599', '\u259F', '\u2588'}

func BitToBlock(source []byte, outputFieldSeperator string, outputRecordSeperator string) string {
	bits := make([]byte, 0, 32)
	for {
		for i := 0; i < 8 && i < len(source); i++ {
			bits = append(bits, source[i]&0x3)
			bits = append(bits, (source[i]>>2)&0x3)
			bits = append(bits, (source[i]>>4)&0x3)
			bits = append(bits, (source[i]>>6)&0x3)
		}
		for i := 8; i < 16 && i < len(source); i++ {
			bits[len(bits)-32+(i-8)*4] |= (source[i] & 0x3) << 2
			bits[len(bits)-32+(i-8)*4+1] |= ((source[i] >> 2) & 0x3) << 2
			bits[len(bits)-32+(i-8)*4+2] |= ((source[i] >> 4) & 0x3) << 2
			bits[len(bits)-32+(i-8)*4+3] |= ((source[i] >> 6) & 0x3) << 2
		}
		if len(source) < 16 {
			break
		}
		source = source[16:]
	}
	var buffer strings.Builder
	ofs := ""
	ors := ""
	for i, c := range bits {
		if i%4 == 0 {
			if i%16 == 0 {
				buffer.WriteString(ors)
			} else {
				buffer.WriteString(ofs)
			}
			ors = outputRecordSeperator
			ofs = outputFieldSeperator
		}
		buffer.WriteRune(bitToBlockTable[c])
	}
	return buffer.String()
}

func main() {
	var bitPattern [256]byte
	for i := 0; i < 256; i++ {
		bitPattern[i] = byte(i)
	}
	println(BitToBlock(bitPattern[:], "|", "\n"))
}
