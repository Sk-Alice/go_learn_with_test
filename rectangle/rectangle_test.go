package rectangle

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("want '%.2f' but got '%.2f'", want, got)
	}
}

// 表格驱动测试----要创建一系列相同测试方式的测试用例时很有用
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	// %#v 输出 Go 语言的语法表示形式，包括类型信息。
	// 它可以用于任何值，包括基本类型、复合类型（如结构体、数组、切片、映射）、指针、通道等
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f.", tt.shape, got, tt.hasArea)
			}
		})
	}

}
