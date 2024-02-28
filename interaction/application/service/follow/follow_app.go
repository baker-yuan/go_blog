// Package service 应用服务
package service

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	"github.com/baker-yuan/go-blog/interaction/application/dto/follow/command"
	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/follow"
	event "github.com/baker-yuan/go-blog/interaction/domain/event/follow"
	port "github.com/baker-yuan/go-blog/interaction/domain/port/follow"
)

// AppService Application 服务
type AppService struct {
	FollowPort port.FollowPort
}

// AddFollow 新增关注关系
func (a *AppService) AddFollow(ctx context.Context, cmd *command.AddFollowCMD) (err error) {
	// 打印日志
	log.Debug(ctx, "application - service - AddFollow cmd: %+v", cmd)

	// 业务编排：取聚合的实体
	followEntity, err := a.getFollow(ctx, cmd)
	log.Debug(ctx, "application - service - AddFollow followEntity: %+v", followEntity)
	// 判断聚合的实体是否存在
	if err != nil {
		log.Error(ctx, "application - service - AddFollow a.FollowPort.Save err: %+v", err)
		return err
	}

	// 业务编排：变更状态
	err = followEntity.ChangeFollowStateToTrue(ctx)
	if err != nil {
		log.Error(ctx, "application - service - AddFollow ChangeFollowStateToTrue err: %+v", err)
		return err
	}

	// 业务编排：保存聚合实体
	fn := a.FollowPort.Save(ctx, followEntity)

	// 业务编排：执行 transaction
	err = a.FollowPort.TxEnd(ctx, fn)
	if err != nil {
		log.Error(ctx, "application - service - AddFollow a.FollowPort.Save err: %+v", err)
		return err
	}

	// 事件逻辑：发送关注关系创建的事件
	domainEvent := &event.FollowCreatedEvent{
		Follow: followEntity,
	}
	event.Publisher.Publish(ctx, domainEvent)

	return nil

}

// getFollowEntity get 或者 build 一个 follow
func (a *AppService) getFollow(ctx context.Context, cmd *command.AddFollowCMD) (followEntity *entity.Follow, err error) {
	// 业务编排：取聚合的实体
	followEntity, err = a.FollowPort.FindByUnique(ctx, cmd.UID, cmd.FollowUID)
	log.Debug(ctx, "application - service - getFollow followEntity: %+v", followEntity)
	// 判断聚合的实体是否存在
	if err != nil {
		log.Error(ctx, "application - service - getFollow a.FollowPort.Save err: %+v", err)
		return nil, err
	}
	if followEntity != nil {
		return followEntity, nil
	}

	// 业务编排：通过 followerUin 获取聚合实体的数量
	followCount, err := a.FollowPort.GetFolloweeCount(ctx, cmd.UID, entity.FollowStateTrue)
	if err != nil {
		log.Error(ctx, "application - service - GetFolloweeCount  a.FollowPort.GetFolloweeCount err: %+v", err)
		return nil, err
	}
	// 打印日志
	log.Debug(ctx, "application - service - GetFolloweeCount followCount: %+v", followCount)

	// 业务编排：生成 Follow 的实体
	followBuilder := &entity.FollowBuilder{}
	followBuilder.UID(cmd.UID)
	followBuilder.FollowUID(cmd.FollowUID)
	followBuilder.State(entity.FollowStateFalse)
	followBuilder.CreateTime(uint32(time.Now().Second()))
	followBuilder.UpdateTime(uint32(time.Now().Second()))
	followBuilder.CountByFollower(followCount)
	followEntity, err = followBuilder.Build(ctx)
	if err != nil {
		log.Error(ctx, "application - service - getFollow build follow err: %+v", err)
		return nil, err
	}
	return followEntity, nil
}
