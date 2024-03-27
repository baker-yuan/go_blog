package migrate

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/store"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
)

var (
	ErrConflict = errors.New("conflict")
)

func Export(ctx context.Context) ([]byte, error) {
	exportData := newDataSet()
	store.RangeStore(func(key store.HubKey, s *store.GenericStore) bool {
		s.Range(ctx, func(_ string, obj interface{}) bool {
			err := exportData.Add(obj)
			if err != nil {
				log.Errorf("Add obj to export list failed:%s", err)
				return true
			}
			return true
		})
		return true
	})

	data, err := json.Marshal(exportData)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ConflictMode int

const (
	ModeReturn ConflictMode = iota
	ModeOverwrite
	ModeSkip
)

func Import(ctx context.Context, data []byte, mode ConflictMode) (*DataSet, error) {
	importData := newDataSet()
	err := json.Unmarshal(data, &importData)
	if err != nil {
		return nil, err
	}
	conflict, conflictData := isConflicted(ctx, importData)
	if conflict && mode == ModeReturn {
		return conflictData, ErrConflict
	}
	store.RangeStore(func(key store.HubKey, s *store.GenericStore) bool {
		importData.rangeData(key, func(i int, obj interface{}) bool {
			_, e := s.CreateCheck(obj)
			if e != nil {
				switch mode {
				case ModeSkip:
					return true
				case ModeOverwrite:
					_, e := s.Update(ctx, obj, true)
					if e != nil {
						err = e
						return false
					}
				}
			} else {
				_, e := s.Create(ctx, obj)
				if err != nil {
					err = e
					return false
				}
			}
			return true
		})
		return true
	})
	return nil, err
}
