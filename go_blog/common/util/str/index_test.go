package str

import "testing"

// TestRandomString 测试 RandomString 函数
func TestRandomString(t *testing.T) {
	length := 20
	str := StrUtils{}.RandomString(length)
	if len(str) != length {
		t.Errorf("RandomString() length = %d, want %d", len(str), length)
	}

	// 测试随机性，生成两个随机字符串并确保它们不相同
	str1 := StrUtils{}.RandomString(length)
	str2 := StrUtils{}.RandomString(length)
	if str1 == str2 {
		t.Errorf("RandomString() returned non-unique strings: %s and %s", str1, str2)
	}
}
