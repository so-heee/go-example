package calc

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.Fatal("sum(1,2) should be 3, but doesn't match")
	}
}

func TestSumAll(t *testing.T) {
	// セットアップコード
	t.Log("setup")
	t.Run("Sum1-2", func(t *testing.T) {
		if Sum(1, 2) != 3 {
			t.Fatal("error")
		}
	})
	t.Run("Sum2-3", func(t *testing.T) {
		if Sum(2, 3) != 5 {
			t.Fatal("error")
		}
	})
	t.Run("Sum3-4", func(t *testing.T) {
		if Sum(3, 4) != 7 {
			t.Fatal("error")
		}
	})
	// ティアダウンコード
	t.Log("tear-down")
}

// 直列
func TestSumSleep(t *testing.T) {
	// セットアップコード
	t.Log("setup")
	t.Run("group", func(t *testing.T) {
		t.Run("Sum1-2", func(t *testing.T) {
			if SumSleep(1, 2) != 3 {
				t.Fatal("error")
			}
		})
		t.Run("Sum2-3", func(t *testing.T) {
			if SumSleep(2, 3) != 5 {
				t.Fatal("error")
			}
		})
		t.Run("Sum3-4", func(t *testing.T) {
			if SumSleep(3, 4) != 7 {
				t.Fatal("error")
			}
		})
	})
	// ティアダウンコード
	t.Log("tear-down")
}

// 並列
func TestSumParallelSleep(t *testing.T) {
	// セットアップコード
	t.Log("setup")
	t.Run("group", func(t *testing.T) {
		t.Run("Sum1-2", func(t *testing.T) {
			t.Parallel()
			if SumSleep(1, 2) != 3 {
				t.Fatal("error")
			}
		})
		t.Run("Sum2-3", func(t *testing.T) {
			t.Parallel()
			if SumSleep(2, 3) != 5 {
				t.Fatal("error")
			}
		})
		t.Run("Sum3-4", func(t *testing.T) {
			t.Parallel()
			if SumSleep(3, 4) != 7 {
				t.Fatal("error")
			}
		})
	})
	// ティアダウンコード
	t.Log("tear-down")
}
