// Package query 查询模式入参
package query

// GetUserListQRY GetUserList QRY
type GetUserListQRY struct {
	LoginType uint32 // 登录类型
	Nickname  string // 昵称
	Current   uint32 // 当前页码
	Size      uint32 // 页码条数
}
