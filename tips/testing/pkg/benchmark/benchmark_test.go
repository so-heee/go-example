package benchmark

import (
	"fmt"
	"testing"
)

// "a"をn個並べたスライスを生成（ベンチマーク用のトークン）
func seed(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, "a")
	}
	return s
}

// 文字列長と文字列結合の関数名を受け取りベンチマークを実行します
func bench(b *testing.B, n int, f func(...string) string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(seed(n)...)
	}
}

func BenchmarkCat3(b *testing.B)     { bench(b, 3, Cat) }
func BenchmarkBuf3(b *testing.B)     { bench(b, 3, Buf) }
func BenchmarkCat100(b *testing.B)   { bench(b, 100, Cat) }
func BenchmarkBuf100(b *testing.B)   { bench(b, 100, Buf) }
func BenchmarkCat10000(b *testing.B) { bench(b, 10000, Cat) }
func BenchmarkBuf10000(b *testing.B) { bench(b, 10000, Buf) }

// BenchmarkXの中でベンチマークを3つ実行
func BenchmarkX(b *testing.B) {
	b.Run("n=3", func(b *testing.B) { bench(b, 3, Cat) })
	b.Run("n=10", func(b *testing.B) { bench(b, 10, Cat) })
	b.Run("n=100", func(b *testing.B) { bench(b, 100, Cat) })
}

// サブベンチマークを使って記述を簡略化
func BenchmarkConcatenate(b *testing.B) {
	benchCases := []struct {
		name string
		n    int
		f    func(...string) string
	}{
		{"Cat", 3, Cat},
		{"Buf", 3, Cat},
		{"Cat", 100, Cat},
		{"Buf", 100, Cat},
		{"Cat", 10000, Cat},
		{"Buf", 10000, Cat},
	}
	for _, c := range benchCases {
		b.Run(fmt.Sprintf("%s%d", c.name, c.n),
			func(b *testing.B) { bench(b, c.n, c.f) })
	}
}
