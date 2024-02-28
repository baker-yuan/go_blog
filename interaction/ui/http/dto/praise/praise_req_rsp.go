package dto

// ObjectPraiseReq 内容点赞
type ObjectPraiseReq struct {
	ModuleCode string `json:"moduleCode"` // 模块标识
	ObjectId   uint32 `json:"objectId"`   // 信息ID
	Uid        uint32 `json:"uid"`        // 用户ID
}

// ObjectRsp 内容点赞
type ObjectRsp struct {
}

// CancelObjectPraiseReq 取消内容点赞
type CancelObjectPraiseReq struct {
	ModuleCode string `json:"moduleCode"` // 模块标识
	ObjectId   uint32 `json:"objectId"`   // 信息ID
	Uid        uint32 `json:"uid"`        // 用户ID
}

// CancelObjectPraiseRsp 取消内容点赞
type CancelObjectPraiseRsp struct {
}
