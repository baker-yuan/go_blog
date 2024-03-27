package store

import (
	"context"
	"sort"

	"github.com/stretchr/testify/mock"
)

type MockInterface struct {
	mock.Mock
	HubKey HubKey
}

func (m *MockInterface) Type() HubKey {
	return m.HubKey
}

func (m *MockInterface) Get(_ context.Context, key string) (interface{}, error) {
	ret := m.Mock.Called(key)
	return ret.Get(0), ret.Error(1)
}

func (m *MockInterface) List(_ context.Context, input ListInput) (*ListOutput, error) {
	ret := m.Called(input)

	var (
		r0 *ListOutput
		r1 error
	)

	if rf, ok := ret.Get(0).(func(ListInput) *ListOutput); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(*ListOutput)
	}

	if input.Less == nil {
		input.Less = defLessFunc
	}

	sort.Slice(r0.Rows, func(i, j int) bool {
		return input.Less(r0.Rows[i], r0.Rows[j])
	})

	r1 = ret.Error(1)

	return r0, r1
}

func (m *MockInterface) Create(ctx context.Context, obj interface{}) (interface{}, error) {
	ret := m.Mock.Called(ctx, obj)
	return ret.Get(0), ret.Error(1)
}

func (m *MockInterface) Update(ctx context.Context, obj interface{}, createOnFail bool) (interface{}, error) {
	ret := m.Mock.Called(ctx, obj, createOnFail)
	return ret.Get(0), ret.Error(1)
}

func (m *MockInterface) BatchDelete(ctx context.Context, keys []string) error {
	ret := m.Mock.Called(ctx, keys)
	return ret.Error(0)
}
