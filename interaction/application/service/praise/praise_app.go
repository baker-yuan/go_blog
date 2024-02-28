// Package service 应用服务
package service

import (
	"context"
	"errors"
	"time"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	"github.com/baker-yuan/go-blog/interaction/application/dto/praise/command"
	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/praise"
	event "github.com/baker-yuan/go-blog/interaction/domain/event/praise"
	port "github.com/baker-yuan/go-blog/interaction/domain/port/praise"
)

// AppService Application 服务
type AppService struct {
	PraisePort port.PraisePort
}

// ObjectPraise 新增点赞
func (a *AppService) ObjectPraise(ctx context.Context, cmd *command.ObjectPraiseCMD) (err error) {
	// 打印日志
	log.Debug(ctx, "application - service - ObjectPraise cmd: %+v", cmd)

	// 业务编排：取聚合的实体
	praiseEntity, err := a.getPraise(ctx, cmd)
	log.Debug(ctx, "application - service - ObjectPraise praiseEntity: %+v", praiseEntity)
	// 判断聚合的实体是否存在
	if err != nil {
		log.Error(ctx, "application - service - ObjectPraise a.getPraise err: %+v", err)
		return err
	}

	// 业务编排：变更状态
	err = praiseEntity.Praise(ctx)
	if err != nil {
		log.Error(ctx, "application - service - ObjectPraise ChangePraiseStateToTrue err: %+v", err)
		return err
	}

	// 业务编排：保存聚合实体
	fn := a.PraisePort.Save(ctx, praiseEntity)

	// 业务编排：执行 transaction
	err = a.PraisePort.TxEnd(ctx, fn)
	if err != nil {
		log.Error(ctx, "application - service - ObjectPraise a.PraisePort.TxEnd err: %+v", err)
		return err
	}

	// 事件逻辑：点赞事件
	domainEvent := &event.PraiseCreatedEvent{
		Praise: praiseEntity,
	}
	event.Publisher.Publish(ctx, domainEvent)
	return nil

}

// CancelObjectPraise 取消点赞
func (a *AppService) CancelObjectPraise(ctx context.Context, cmd *command.CancelObjectPraiseCMD) (err error) {
	// 打印日志
	log.Debug(ctx, "application - service - CancelObjectPraise cmd: %+v", cmd)

	// 业务编排：取聚合的实体
	praiseEntity, err := a.PraisePort.FindByUnique(ctx, cmd.ModuleCode, cmd.ObjectId, cmd.Uid)
	if err != nil {
		log.Error(ctx, "application - service - DeleteFollow a.PraisePort.FindByUnique err: %+v", err)
		return err
	}
	if praiseEntity == nil {
		log.Error(ctx, "application - service - CancelObjectPraise a.PraisePort.FindByUnique Praise is nil, cmd: %+v", cmd)
		return errors.New("praise is not exists")
	}

	// 业务编排：变更状态
	err = praiseEntity.CancelPraise(ctx)
	if err != nil {
		log.Error(ctx, "application - service - ObjectPraise ChangePraiseStateToTrue err: %+v", err)
		return err
	}

	// 业务编排：保存聚合实体
	fn := a.PraisePort.Save(ctx, praiseEntity)

	// 业务编排：执行 transaction
	err = a.PraisePort.TxEnd(ctx, fn)
	if err != nil {
		log.Error(ctx, "application - service - a.PraisePort.TxEnd err: %+v", err)
		return err
	}

	// 事件逻辑：发送关注关系删除的事件
	domainEvent := &event.PraiseDeletedEvent{
		Praise: praiseEntity,
	}
	event.Publisher.Publish(ctx, domainEvent)
	return nil
}

// getPraiseEntity get 或者 build 一个 praise
func (a *AppService) getPraise(ctx context.Context, cmd *command.ObjectPraiseCMD) (*entity.Praise, error) {
	var (
		praiseEntity *entity.Praise
		err          error
	)
	// 业务编排：取聚合的实体
	praiseEntity, err = a.PraisePort.FindByUnique(ctx, cmd.ModuleCode, cmd.ObjectId, cmd.Uid)
	log.Debug(ctx, "application - service - getPraise praiseEntity: %+v", praiseEntity)

	// 判断聚合的实体是否存在
	if err != nil {
		log.Error(ctx, "application - service - getPraise a.PraisePort.FindByUnique err: %+v", err)
		return nil, err
	}
	if praiseEntity != nil {
		return praiseEntity, nil
	}

	// 业务编排：生成 ObjectPraise 的实体
	praiseBuilder := &entity.PraiseBuilder{}
	praiseBuilder.UID(cmd.Uid)
	praiseBuilder.ModuleCode(cmd.ModuleCode)
	praiseBuilder.ObjectId(cmd.ObjectId)
	praiseBuilder.Status(entity.PraiseFalse)
	praiseBuilder.CreateTime(uint32(time.Now().Second()))
	praiseBuilder.UpdateTime(uint32(time.Now().Second()))
	praiseEntity, err = praiseBuilder.Build(ctx)
	if err != nil {
		log.Error(ctx, "application - service - getPraise build praise err: %+v", err)
		return nil, err
	}
	return praiseEntity, nil
}
