package benchmark

import "bytes"

// Cat += 演算時による文字列結合
func Cat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}
	return r
}

// Buf bytes.Bufferによる文字列結合
func Buf(ss ...string) string {
	var b bytes.Buffer
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}
