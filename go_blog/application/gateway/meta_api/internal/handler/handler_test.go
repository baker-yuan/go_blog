package handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/store"
	"github.com/shiningrush/droplet"
	"github.com/shiningrush/droplet/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSpecCodeResponse(t *testing.T) {
	err := errors.New("schema validate failed: remote_addr: Must validate at least one schema (anyOf)")
	resp := SpecCodeResponse(err)
	assert.Equal(t, &data.SpecCodeResponse{StatusCode: http.StatusBadRequest}, resp)

	err = errors.New("data not found")
	resp = SpecCodeResponse(err)
	assert.Equal(t, &data.SpecCodeResponse{StatusCode: http.StatusNotFound}, resp)

	err = errors.New("system error")
	resp = SpecCodeResponse(err)
	assert.Equal(t, &data.SpecCodeResponse{StatusCode: http.StatusInternalServerError}, resp)
}

func TestIDCompare(t *testing.T) {
	// init
	cases := []struct {
		idOnPath, desc string
		idOnBody       interface{}
		wantError      error
	}{
		{
			desc:     "ID on body is int, and it could be considered the same as ID on path",
			idOnPath: "1",
			idOnBody: 1,
		},
		{
			desc:      "ID on body is int, and it is different from ID on path",
			idOnPath:  "1",
			idOnBody:  2,
			wantError: errors.New("ID on path (1) doesn't match ID on body (2)"),
		},
		{
			desc:     "ID on body is same as ID on path",
			idOnPath: "1",
			idOnBody: "1",
		},
		{
			desc:      "ID on body is different from ID on path",
			idOnPath:  "a",
			idOnBody:  "b",
			wantError: errors.New("ID on path (a) doesn't match ID on body (b)"),
		},
		{
			desc:     "No ID on body",
			idOnPath: "1",
		},
		{
			desc:     "No ID on path",
			idOnBody: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			err := IDCompare(c.idOnPath, c.idOnBody)
			assert.Equal(t, c.wantError, err)
		})
	}
}

func TestNameExistCheck(t *testing.T) {
	tests := []struct {
		resource string
		name     string
		id       interface{}
		mockErr  error
		mockRet  []interface{}
		wantErr  error
		wantRet  interface{}
		caseDesc string
	}{
		{
			caseDesc: "normal, name not exists",
			resource: "service",
			name:     "test",
		},
		{
			caseDesc: "get list error",
			resource: "route",
			name:     "test",
			mockErr:  errors.New("test error"),
			wantErr:  errors.New("test error"),
			wantRet:  &data.SpecCodeResponse{StatusCode: http.StatusInternalServerError},
		},
		{
			caseDesc: "name exists",
			resource: "upstream",
			name:     "test",
			mockRet:  []interface{}{"test"},
			wantErr:  errors.New("upstream name exists"),
			wantRet:  &data.SpecCodeResponse{StatusCode: http.StatusBadRequest},
		},
	}
	for _, tc := range tests {
		t.Run(tc.caseDesc, func(t *testing.T) {
			mStore := &store.MockInterface{}
			mStore.On("List", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
			}).Return(func(input store.ListInput) *store.ListOutput {
				return &store.ListOutput{
					Rows:      tc.mockRet,
					TotalSize: len(tc.mockRet),
				}
			}, tc.mockErr)

			ctx := droplet.NewContext()
			res, err := NameExistCheck(ctx.Context(), mStore, tc.resource, tc.name, tc.id)

			if res != nil {
				assert.Equal(t, tc.wantRet.(*data.SpecCodeResponse).StatusCode, res.(*data.SpecCodeResponse).StatusCode)
			}
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
