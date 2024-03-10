package db

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var defaultTx TransactionImpl

type TransactionImpl struct {
	DB *gorm.DB
}

func InitTransaction(db *gorm.DB) {
	defaultTx = TransactionImpl{
		DB: db,
	}
}

type contextTxKey struct{}

// PutDbToCtx 将tx放入到ctx中
func (t TransactionImpl) PutDbToCtx(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// GetDbToCtx 获取PutDbToCtx存入的tx
func (t TransactionImpl) GetDbToCtx(ctx context.Context) (*gorm.DB, error) {
	txKey := ctx.Value(contextTxKey{})
	tx, ok := txKey.(*gorm.DB)
	if ok {
		return tx, nil
	}
	return nil, errors.New("get gorm.DB fail from context")
}

// Transaction 将tx放入到ctx中
func Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return defaultTx.PutDbToCtx(ctx, fn)
}

// GetTransactionDB 获取PutDbToCtx存入的tx
func GetTransactionDB(ctx context.Context) (*gorm.DB, error) {
	return defaultTx.GetDbToCtx(ctx)
}
