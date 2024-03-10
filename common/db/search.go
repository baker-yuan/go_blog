package db

import (
	"context"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

// SearchValue 搜索值
type SearchValue struct {
	Value     string `json:"value"`     // 值
	Condition string `json:"condition"` // 条件
}

var (
	EQ       = "eq"      // =
	NE       = "ne"      // !=
	GT       = "gt"      // >
	GTE      = "gte"     // >=
	LT       = "lt"      // <
	LTE      = "lte"     // <=
	IN       = "in"      // in
	NOT_IN   = "notIn"   // not in
	LIKE     = "like"    // like
	NOT_LIKE = "notLike" // not like
)

// ChangeValue 扩展，改变值
type ChangeValue func(search map[string]*SearchValue)

func BuildSearch(ctx context.Context, strSearch string, gormDB *gorm.DB, changeValue ...ChangeValue) (*gorm.DB, error) {
	if len(strSearch) == 0 {
		return gormDB, nil
	}
	// 将 json string 转成map
	search := map[string]*SearchValue{}
	if err := json.Unmarshal([]byte(strSearch), &search); err != nil {
		return nil, err
	}
	// 扩展，允许对序列化对结果再次做改变
	if len(changeValue) != 0 {
		changeValue[0](search)
	}
	log.InfoContextf(ctx, "#buildSearch %+v", search)

	if len(search) == 0 {
		return gormDB, nil
	}

	// 构造条件
	for k, v := range search {
		switch v.Condition {
		case EQ:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` = ?", k), v.Value)
		case NE:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` != ?", k), v.Value)
		case GT:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` > ?", k), v.Value)
		case GTE:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` >= ?", k), v.Value)
		case LT:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` < ?", k), v.Value)
		case LTE:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` <= ?", k), v.Value)
		case IN:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` in ?", k), v.Value)
		case NOT_IN:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` not in ?", k), v.Value)
		case LIKE:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` like ?", k), "%"+v.Value+"%")
		case NOT_LIKE:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` not like ?", k), "%"+v.Value+"%")
		default:
			gormDB = gormDB.Where(fmt.Sprintf("`%s` = ?", k), v.Value)
		}
	}

	return gormDB, nil
}
