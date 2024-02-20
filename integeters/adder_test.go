package integeters

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// 示例函数：用于展示如何使用包中的函数或方法
// Output: 注释来指定预期的输出,如果删除注释,示例函数将不会执行。虽然函数会被编译，但是它不会执行
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
