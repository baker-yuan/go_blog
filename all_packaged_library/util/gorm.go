package util

import (
	"reflect"
	"strings"

	"gorm.io/gorm"
)

const (
	// =
	eq = "eq"
	// !=
	ne = "ne"
	// >
	gt = "gt"
	// >=
	gte = "gte"
	// <
	lt = "lt"
	// <=
	lte = "lte"
	// in
	in = "in"
	// not in
	notIn = "notIn"
	// like
	like = "like"
	// not like
	notLike = "notLike"

	// 分隔符
	separator = "_"
)

func endWith(v string) bool {
	if strings.HasSuffix(v, eq) ||
		strings.HasSuffix(v, ne) ||
		strings.HasSuffix(v, gt) ||
		strings.HasSuffix(v, gte) ||
		strings.HasSuffix(v, lt) ||
		strings.HasSuffix(v, lte) ||
		strings.HasSuffix(v, in) ||
		strings.HasSuffix(v, notIn) ||
		strings.HasSuffix(v, like) ||
		strings.HasSuffix(v, notLike) {
		return true
	}
	return false
}
func IsNil(v interface{}) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		return rv.IsNil()
	}
	return v == nil
}

func BuildSearch(tx *gorm.DB, search map[string]interface{}) *gorm.DB {
	for k, v := range search {
		if IsNil(v) {
			continue
		}
		var (
			field string
			op    string
		)
		if endWith(k) {
			idx := strings.LastIndex(k, separator)
			field, op = k[:idx], k[idx+1:]
		} else {
			// 没有分隔符认为相等
			field, op = k, eq
		}
		switch op {
		case eq:
			tx.Where(field+" = ? ", v)
		case like:
			tx.Where(field+" like ? ", v)
		case in:
			tx.Where(field+" in ? ", v)
		}
	}
	return tx
}
