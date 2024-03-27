package openapi3

import (
	"regexp"

	"github.com/getkin/kin-openapi/openapi3"
)

type OpenAPISpecFileType string

type Loader struct {
	// MergeMethod indicates whether to merge routes when multiple HTTP methods are on the same path
	MergeMethod bool
	// TaskName indicates the name of current import/export task
	TaskName string
}

type PathValue struct {
	Method string
	Value  *openapi3.Operation
}

var (
	regURIVar = regexp.MustCompile(`{.*?}`)
)
