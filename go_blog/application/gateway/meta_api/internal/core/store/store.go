package store

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/entity"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/storage"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/utils"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/utils/runtime"
	"github.com/shiningrush/droplet/data"
)

var (
	storeNeedReInit = make([]*GenericStore, 0)
)

// Pagination 分页查询
type Pagination struct {
	PageSize   int `json:"page_size" form:"page_size" auto_read:"page_size"` // 页码
	PageNumber int `json:"page" form:"page" auto_read:"page"`                // 页面大小
}

// Interface 存储接口
type Interface interface {
	Type() HubKey                                                                            //
	Get(ctx context.Context, key string) (interface{}, error)                                //
	List(ctx context.Context, input ListInput) (*ListOutput, error)                          // 列表查询
	Create(ctx context.Context, obj interface{}) (interface{}, error)                        // 创建
	Update(ctx context.Context, obj interface{}, createIfNotExist bool) (interface{}, error) // 批量删除
	BatchDelete(ctx context.Context, keys []string) error
}

// GenericStore 通用存储
type GenericStore struct {
	Stg      storage.Interface  // 底层存储层操作
	opt      GenericStoreOption // 选项
	cache    sync.Map           // 缓存底层数据到内存，后续直接操作内存
	cancel   context.CancelFunc // 用于取消数据监听(监听数据同步内存)，GenericStore#watch
	initLock sync.Mutex         // GenericStore#Init加锁
	closing  bool               // GenericStore#Close调用标识
}

// GenericStoreOption 选项
type GenericStoreOption struct {
	BasePath   string                                            // 资源目录
	ObjType    reflect.Type                                      // 模型对应的type
	HubKey     HubKey                                            // 资源类型
	KeyFunc    func(obj interface{}) string                      // 生成资源唯一key
	Validator  Validator                                         // 数据校验，基于json schema，配置位置api/conf/schema.json
	StockCheck func(obj interface{}, stockObj interface{}) error //
}

func NewGenericStore(opt GenericStoreOption) (*GenericStore, error) {
	if opt.BasePath == "" {
		log.Error("base path empty")
		return nil, fmt.Errorf("base path can not be empty")
	}
	if opt.ObjType == nil {
		log.Errorf("object type is nil")
		return nil, fmt.Errorf("object type can not be nil")
	}
	if opt.KeyFunc == nil {
		log.Error("key func is nil")
		return nil, fmt.Errorf("key func can not be nil")
	}

	if opt.ObjType.Kind() == reflect.Ptr {
		opt.ObjType = opt.ObjType.Elem()
	}
	if opt.ObjType.Kind() != reflect.Struct {
		log.Error("obj type is invalid")
		return nil, fmt.Errorf("obj type is invalid")
	}
	s := &GenericStore{
		opt: opt,
	}
	// 底层存储基于etcd
	s.Stg = storage.GenEtcdStorage()

	return s, nil
}

func ReInit() error {
	for _, store := range storeNeedReInit {
		if err := store.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (s *GenericStore) Init() error {
	s.initLock.Lock()
	defer s.initLock.Unlock()
	return s.listAndWatch()
}

func (s *GenericStore) Type() HubKey {
	return s.opt.HubKey
}

// Get 单个查询
func (s *GenericStore) Get(_ context.Context, key string) (interface{}, error) {
	ret, ok := s.cache.Load(key)
	if !ok {
		log.Warnf("data not found by key: %s", key)
		return nil, data.ErrNotFound
	}
	return ret, nil
}

// ListInput 列表查询条件
type ListInput struct {
	Predicate  func(obj interface{}) bool        // 条件
	Format     func(obj interface{}) interface{} //
	PageSize   int                               // 页码
	PageNumber int                               // 页面大小 start from 1
	Less       func(i, j interface{}) bool       //
}

// ListOutput 列表查询返回
type ListOutput struct {
	Rows      []interface{} `json:"rows"`
	TotalSize int           `json:"total_size"`
}

// NewListOutput returns JSON marshalling safe struct pointer for empty slice
func NewListOutput() *ListOutput {
	return &ListOutput{Rows: make([]interface{}, 0)}
}

// defLessFunc 列表排序
var defLessFunc = func(i, j interface{}) bool {
	iBase := i.(entity.GetBaseInfo).GetBaseInfo()
	jBase := j.(entity.GetBaseInfo).GetBaseInfo()
	if iBase.CreateTime != jBase.CreateTime {
		return iBase.CreateTime < jBase.CreateTime
	}
	if iBase.UpdateTime != jBase.UpdateTime {
		return iBase.UpdateTime < jBase.UpdateTime
	}
	iID := utils.InterfaceToString(iBase.ID)
	jID := utils.InterfaceToString(jBase.ID)
	return iID < jID
}

// List 列表查询
func (s *GenericStore) List(_ context.Context, input ListInput) (*ListOutput, error) {
	var ret []interface{}

	// 遍历缓存
	s.cache.Range(func(key, value interface{}) bool {
		if input.Predicate != nil && !input.Predicate(value) {
			return true
		}
		if input.Format != nil {
			value = input.Format(value)
		}
		ret = append(ret, value)
		return true
	})

	// 数据为空返回空数组
	// should return an empty array not a null for client
	if ret == nil {
		ret = []interface{}{}
	}

	// 组装返回数据
	output := &ListOutput{
		Rows:      ret,
		TotalSize: len(ret),
	}

	// 排序
	if input.Less == nil {
		input.Less = defLessFunc
	}
	sort.Slice(output.Rows, func(i, j int) bool {
		return input.Less(output.Rows[i], output.Rows[j])
	})

	// 分页
	if input.PageSize > 0 && input.PageNumber > 0 {
		skipCount := (input.PageNumber - 1) * input.PageSize
		if skipCount > output.TotalSize {
			output.Rows = []interface{}{}
			return output, nil
		}
		endIdx := skipCount + input.PageSize
		if endIdx >= output.TotalSize {
			output.Rows = ret[skipCount:]
			return output, nil
		}
		output.Rows = ret[skipCount:endIdx]
	}

	return output, nil
}

func (s *GenericStore) Range(_ context.Context, f func(key string, obj interface{}) bool) {
	s.cache.Range(func(key, value interface{}) bool {
		return f(key.(string), value)
	})
}

func (s *GenericStore) ingestValidate(obj interface{}) (err error) {
	// 校验 json schema
	if s.opt.Validator != nil {
		if err := s.opt.Validator.Validate(obj); err != nil {
			log.Errorf("data validate failed: %s, %v", err, obj)
			return err
		}
	}

	if s.opt.StockCheck != nil {
		s.cache.Range(func(key, value interface{}) bool {
			if err = s.opt.StockCheck(obj, value); err != nil {
				return false
			}
			return true
		})
	}
	return err
}

func (s *GenericStore) CreateCheck(obj interface{}) ([]byte, error) {

	if setter, ok := obj.(entity.GetBaseInfo); ok {
		info := setter.GetBaseInfo()
		info.Creating()
	}

	if err := s.ingestValidate(obj); err != nil {
		return nil, err
	}

	key := s.opt.KeyFunc(obj)
	if key == "" {
		return nil, fmt.Errorf("key is required")
	}
	_, ok := s.cache.Load(key)
	if ok {
		log.Warnf("key: %s is conflicted", key)
		return nil, fmt.Errorf("key: %s is conflicted", key)
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Errorf("json marshal failed: %s", err)
		return nil, fmt.Errorf("json marshal failed: %s", err)
	}

	return bytes, nil
}

func (s *GenericStore) Create(ctx context.Context, obj interface{}) (interface{}, error) {
	if setter, ok := obj.(entity.GetBaseInfo); ok {
		info := setter.GetBaseInfo()
		info.Creating()
	}

	// 检查
	bytes, err := s.CreateCheck(obj)
	if err != nil {
		return nil, err
	}

	// 调用底层存储保存数据
	if err := s.Stg.Create(ctx, s.GetObjStorageKey(obj), string(bytes)); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GenericStore) Update(ctx context.Context, obj interface{}, createIfNotExist bool) (interface{}, error) {
	if err := s.ingestValidate(obj); err != nil {
		return nil, err
	}

	key := s.opt.KeyFunc(obj)
	if key == "" {
		return nil, fmt.Errorf("key is required")
	}
	storedObj, ok := s.cache.Load(key)
	if !ok {
		if createIfNotExist {
			return s.Create(ctx, obj)
		}
		log.Warnf("key: %s is not found", key)
		return nil, fmt.Errorf("key: %s is not found", key)
	}

	if setter, ok := obj.(entity.GetBaseInfo); ok {
		storedGetter := storedObj.(entity.GetBaseInfo)
		storedInfo := storedGetter.GetBaseInfo()
		info := setter.GetBaseInfo()
		info.Updating(storedInfo)
	}

	bs, err := json.Marshal(obj)
	if err != nil {
		log.Errorf("json marshal failed: %s", err)
		return nil, fmt.Errorf("json marshal failed: %s", err)
	}
	if err := s.Stg.Update(ctx, s.GetObjStorageKey(obj), string(bs)); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GenericStore) BatchDelete(ctx context.Context, keys []string) error {
	var storageKeys []string
	for i := range keys {
		storageKeys = append(storageKeys, s.GetStorageKey(keys[i]))
	}

	return s.Stg.BatchDelete(ctx, storageKeys)
}

// listAndWatch 全量获取数据，并且开始监听数据变化，存储到本地缓存s.cache里面
func (s *GenericStore) listAndWatch() error {
	lc, lcancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer lcancel()

	// 全量获取数据缓存到本地
	ret, err := s.Stg.List(lc, s.opt.BasePath)
	if err != nil {
		return err
	}
	for i := range ret {
		key := ret[i].Key[len(s.opt.BasePath)+1:]
		objPtr, err := s.StringToObjPtr(ret[i].Value, key)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error occurred while initializing logical store: %s, err: %v", s.opt.BasePath, err)
			return err
		}
		s.cache.Store(s.opt.KeyFunc(objPtr), objPtr)
	}

	// 监听增量数据变化 start watch
	s.cancel = s.watch()

	return nil
}

// watch 监听底层数据变化
func (s *GenericStore) watch() context.CancelFunc {
	// 取消watch
	c, cancel := context.WithCancel(context.TODO())
	// watch底层数据
	ch := s.Stg.Watch(c, s.opt.BasePath)
	go func() {
		defer func() {
			if !s.closing {
				log.Errorf("etcd watch exception closed, restarting: resource: %s", s.Type())
				storeNeedReInit = append(storeNeedReInit, s)
			}
		}()
		defer runtime.HandlePanic()
		for event := range ch {
			if event.Canceled {
				log.Warnf("etcd watch failed: %s", event.Error)
				return
			}

			// 维护本地缓存
			for i := range event.Events {
				switch event.Events[i].Type {
				case storage.EventTypePut:
					key := event.Events[i].Key[len(s.opt.BasePath)+1:]
					objPtr, err := s.StringToObjPtr(event.Events[i].Value, key)
					if err != nil {
						log.Warnf("value convert to obj failed: %s", err)
						continue
					}
					s.cache.Store(key, objPtr)
				case storage.EventTypeDelete:
					s.cache.Delete(event.Events[i].Key[len(s.opt.BasePath)+1:])
				}
			}
		}
	}()
	return cancel
}

func (s *GenericStore) Close() error {
	s.closing = true
	s.cancel()
	return nil
}

func (s *GenericStore) StringToObjPtr(str, key string) (interface{}, error) {
	objPtr := reflect.New(s.opt.ObjType)
	ret := objPtr.Interface()
	err := json.Unmarshal([]byte(str), ret)
	if err != nil {
		log.Errorf("json unmarshal failed: %s", err)
		return nil, fmt.Errorf("json unmarshal failed\n\tRelated Key:\t\t%s\n\tError Description:\t%s", key, err)
	}

	if setter, ok := ret.(entity.GetBaseInfo); ok {
		info := setter.GetBaseInfo()
		info.KeyCompat(key)
	}

	return ret, nil
}

func (s *GenericStore) GetObjStorageKey(obj interface{}) string {
	return s.GetStorageKey(s.opt.KeyFunc(obj))
}

func (s *GenericStore) GetStorageKey(key string) string {
	return fmt.Sprintf("%s/%s", s.opt.BasePath, key)
}
