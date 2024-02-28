// Package dao DAO 层
package dao

import (
	"github.com/baker-yuan/go-blog/interaction/infrastructure/persistence/follow/mysql/po"
)

// FollowDAO Follow 的 DAO
type FollowDAO struct {
	FollowPO po.FollowPO
}
