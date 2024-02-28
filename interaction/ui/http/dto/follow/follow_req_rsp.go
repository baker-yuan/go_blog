package dto

// AddFollowReq 新增关注关系
type AddFollowReq struct {
	UID       uint32 `json:"uid,omitempty"`
	FollowUID uint32 `json:"follow_uid,omitempty"`
}

// AddFollowRsp 新增关注关系
type AddFollowRsp struct {
}
