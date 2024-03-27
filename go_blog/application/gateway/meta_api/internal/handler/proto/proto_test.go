package proto

import (
	"encoding/json"
	"testing"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestStructUnmarshal(t *testing.T) {
	// define and parse data
	jsonStr := `{
    "id": 1,
    "create_time": 1700000000,
    "update_time": 1700000000,
    "desc": "desc",
    "content": "content"
}`
	proto := entity.Proto{}
	err := json.Unmarshal([]byte(jsonStr), &proto)

	// asserts
	assert.Nil(t, err)
	assert.Equal(t, proto.ID, float64(1))
	assert.Equal(t, proto.CreateTime, int64(1700000000))
	assert.Equal(t, proto.UpdateTime, int64(1700000000))
	assert.Equal(t, proto.Desc, "desc")
	assert.Equal(t, proto.Content, "content")
}
