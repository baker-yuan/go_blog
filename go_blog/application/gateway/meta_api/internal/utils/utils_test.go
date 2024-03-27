package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFlakeUid(t *testing.T) {
	id := GetFlakeUid()
	assert.NotEqual(t, 0, id)
}

func TestGetFlakeUidStr(t *testing.T) {
	id := GetFlakeUidStr()
	assert.NotEqual(t, "", id)
	assert.Equal(t, 18, len(id))
}

func TestGetLocalIPs(t *testing.T) {
	_, err := getLocalIPs()
	assert.Equal(t, nil, err)
}

func TestSumIPs_with_nil(t *testing.T) {
	total := sumIPs(nil)
	assert.Equal(t, uint16(0), total)
}

func TestObjectClone(t *testing.T) {
	type test struct {
		Str string
		Num int
	}

	origin := &test{Str: "a", Num: 1}
	copy := &test{}
	err := ObjectClone(origin, copy)
	assert.Nil(t, err)
	assert.Equal(t, origin, copy)

	// change value of the copy, should not change value of origin
	copy.Num = 2
	assert.NotEqual(t, copy.Num, origin.Num)
	assert.Equal(t, 1, origin.Num)
}

func TestGenLabelMap(t *testing.T) {
	expectedErr := errors.New("malformed label")
	mp, err := GenLabelMap("l1")
	assert.Nil(t, err)
	assert.Equal(t, mp["l1"], struct{}{})

	mp, err = GenLabelMap("l1,l2:v2")
	assert.Nil(t, err)
	assert.Equal(t, mp["l1"], struct{}{})
	assert.Equal(t, mp["l2:v2"], struct{}{})

	mp, err = GenLabelMap("l1:v1,l1:v2")
	assert.Nil(t, err)
	assert.Equal(t, mp["l1:v1"], struct{}{})
	assert.Equal(t, mp["l1:v2"], struct{}{})

	mp, err = GenLabelMap(",")
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, mp)

	mp, err = GenLabelMap(",l2:,")
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, mp)
}

func TestLabelContains(t *testing.T) {
	reqMap, _ := GenLabelMap("l1,l2:v2")
	mp := map[string]string{
		"l1": "v1",
	}
	assert.True(t, LabelContains(mp, reqMap))

	mp = map[string]string{
		"l1": "v1",
		"l2": "v3",
	}
	assert.True(t, LabelContains(mp, reqMap))

	mp = map[string]string{
		"l2": "v3",
	}
	assert.False(t, LabelContains(mp, reqMap))

	reqMap, _ = GenLabelMap("l1:v1,l1:v2")
	mp = map[string]string{
		"l1": "v1",
	}
	assert.True(t, LabelContains(mp, reqMap))

	reqMap, _ = GenLabelMap("l1:v1,l1:v2")
	mp = map[string]string{
		"l1": "v2",
	}
	assert.True(t, LabelContains(mp, reqMap))
}

func TestValidateLuaCode(t *testing.T) {
	validLuaCode := "local _M = {} \n function _M.access(api_ctx) \n ngx.log(ngx.WARN,\"hit access phase\") \n end \nreturn _M"
	err := ValidateLuaCode(validLuaCode)
	assert.Nil(t, err)

	invalidLuaCode := "local _M = {} \n function _M.access(api_ctx) \n ngx.log(ngx.WARN,\"hit access phase\")"
	err = ValidateLuaCode(invalidLuaCode)
	assert.NotNil(t, err)
	assert.Equal(t, "<string> at EOF:   syntax error\n", err.Error())
}
