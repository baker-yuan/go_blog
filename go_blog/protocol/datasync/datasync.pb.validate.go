// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: datasync/datasync.proto

package datasync

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on TableChange with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TableChange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TableChange with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TableChangeMultiError, or
// nil if none found.
func (m *TableChange) ValidateAll() error {
	return m.validate(true)
}

func (m *TableChange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DbName

	// no validation rules for TbName

	// no validation rules for ColumnMap

	// no validation rules for BeforeColumnMap

	// no validation rules for AfterColumnMap

	if len(errors) > 0 {
		return TableChangeMultiError(errors)
	}

	return nil
}

// TableChangeMultiError is an error wrapping multiple validation errors
// returned by TableChange.ValidateAll() if the designated constraints aren't met.
type TableChangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TableChangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TableChangeMultiError) AllErrors() []error { return m }

// TableChangeValidationError is the validation error returned by
// TableChange.Validate if the designated constraints aren't met.
type TableChangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TableChangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TableChangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TableChangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TableChangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TableChangeValidationError) ErrorName() string { return "TableChangeValidationError" }

// Error satisfies the builtin error interface
func (e TableChangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTableChange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TableChangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TableChangeValidationError{}

// Validate checks the field values on DataChangeRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DataChangeRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DataChangeRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DataChangeRspMultiError, or
// nil if none found.
func (m *DataChangeRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *DataChangeRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DataChangeRspMultiError(errors)
	}

	return nil
}

// DataChangeRspMultiError is an error wrapping multiple validation errors
// returned by DataChangeRsp.ValidateAll() if the designated constraints
// aren't met.
type DataChangeRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataChangeRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataChangeRspMultiError) AllErrors() []error { return m }

// DataChangeRspValidationError is the validation error returned by
// DataChangeRsp.Validate if the designated constraints aren't met.
type DataChangeRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataChangeRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataChangeRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataChangeRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataChangeRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataChangeRspValidationError) ErrorName() string { return "DataChangeRspValidationError" }

// Error satisfies the builtin error interface
func (e DataChangeRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataChangeRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataChangeRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataChangeRspValidationError{}