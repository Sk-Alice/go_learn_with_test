package interation

// https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

// 基准测试,对一个特定的操作进行多次迭代，以确保测量结果的准确性
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 10)
	fmt.Println(repeat)
	// Output: bbbbbbbbbb
}
