package slice

import (
	"reflect"
	"testing"
)

// Person 结构体表示一个人的信息。
type Person struct {
	Name string
	Age  int
}

// People 是 Person 结构体切片的类型别名。
type People []Person

// TestExtractField 测试 FieldExtractor 的 ExtractField 方法。
func TestExtractField(t *testing.T) {

	// 创建一个示例切片
	people := People{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	// 创建 FieldExtractor 实例
	extractor := FieldExtractor[Person, string]{}

	// 使用 FieldExtractor 实例的 ExtractField 方法提取 Name 字段
	names := extractor.ExtractField(people, func(p Person) string {
		return p.Name
	})

	// 预期的结果
	expected := []string{"Alice", "Bob", "Charlie"}

	// 检查结果是否符合预期
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("ExtractField() = %v, want %v", names, expected)
	}
}
