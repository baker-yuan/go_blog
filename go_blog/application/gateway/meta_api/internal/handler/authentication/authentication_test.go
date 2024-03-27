package authentication

import (
	"encoding/json"
	"testing"

	"github.com/shiningrush/droplet"
	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	// init
	handler := &Handler{}
	assert.NotNil(t, handler)

	// login
	input := &LoginInput{}
	ctx := droplet.NewContext()
	reqBody := `{
	  "username": "admin",
	  "password": "admin"
	}`
	err := json.Unmarshal([]byte(reqBody), input)
	assert.Nil(t, err)
	ctx.SetInput(input)
	_, err = handler.userLogin(ctx)
	assert.Nil(t, err)

	// username error
	input2 := &LoginInput{}
	reqBody = `{
	  "username": "sdfasdf",
	  "password": "admin"
	}`
	err = json.Unmarshal([]byte(reqBody), input2)
	assert.Nil(t, err)
	ctx.SetInput(input2)
	_, err = handler.userLogin(ctx)
	assert.EqualError(t, err, "username or password error")

	// password error
	input3 := &LoginInput{}
	reqBody = `{
	  "username": "admin",
	  "password": "admin9384938"
	}`
	err = json.Unmarshal([]byte(reqBody), input3)
	assert.Nil(t, err)
	ctx.SetInput(input3)
	_, err = handler.userLogin(ctx)
	assert.EqualError(t, err, "username or password error")

}
