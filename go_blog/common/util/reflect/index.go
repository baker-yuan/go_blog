package reflect

import (
	"fmt"
	"reflect"
)

// ReflectUtils 反射
type ReflectUtils struct {
}

func (c ReflectUtils) TypeNameOf(v interface{}) string {
	return c.TypeName(reflect.TypeOf(v))
}

func (c ReflectUtils) TypeName(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		return fmt.Sprint("*", c.TypeName(t.Elem()))
	}
	return fmt.Sprintf("%s.%s", t.PkgPath(), t.String())
}
