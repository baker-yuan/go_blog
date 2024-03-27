package migrate

import (
	"context"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/store"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
)

func isConflicted(ctx context.Context, new *DataSet) (bool, *DataSet) {
	isConflict := false
	conflictedData := newDataSet()
	store.RangeStore(func(key store.HubKey, s *store.GenericStore) bool {
		new.rangeData(key, func(i int, obj interface{}) bool {
			// Only check key of store conflict for now.
			// TODO: Maybe check name of some entiries.
			_, err := s.CreateCheck(obj)
			if err != nil {
				isConflict = true
				err = conflictedData.Add(obj)
				if err != nil {
					log.Errorf("Add obj to conflict list failed:%s", err)
					return true
				}
			}
			return true
		})
		return true
	})
	return isConflict, conflictedData
}
