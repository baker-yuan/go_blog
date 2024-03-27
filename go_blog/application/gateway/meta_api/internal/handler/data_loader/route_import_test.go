package data_loader

import (
	"testing"

	"github.com/shiningrush/droplet"
	"github.com/stretchr/testify/assert"
)

func TestImport_invalid_loader(t *testing.T) {
	input := &ImportInput{}
	input.Type = "test"
	input.FileName = "file1.yaml"
	input.FileContent = []byte("hello")

	h := ImportHandler{}
	ctx := droplet.NewContext()
	ctx.SetInput(input)

	_, err := h.Import(ctx)
	assert.EqualError(t, err, "unsupported data loader type: test")
}

func TestImport_openapi3_invalid_file_type(t *testing.T) {
	input := &ImportInput{}
	input.FileName = "file1.txt"
	input.FileContent = []byte("hello")

	h := ImportHandler{}
	ctx := droplet.NewContext()
	ctx.SetInput(input)

	_, err := h.Import(ctx)
	assert.EqualError(t, err, "required file type is .yaml, .yml or .json but got: .txt")
}

func TestImport_openapi3_invalid_content(t *testing.T) {
	input := &ImportInput{}
	input.Type = "openapi3"
	input.FileName = "file1.json"
	input.FileContent = []byte(`{"test": "a"}`)

	h := ImportHandler{}
	ctx := droplet.NewContext()
	ctx.SetInput(input)

	_, err := h.Import(ctx)
	assert.EqualError(t, err, "empty or invalid imported file: OpenAPI documentation does not contain any paths")
}
