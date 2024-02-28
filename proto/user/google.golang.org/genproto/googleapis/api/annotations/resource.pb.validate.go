// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: google/api/resource.proto

package annotations

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

// Validate checks the field values on ResourceDescriptor with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ResourceDescriptor) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResourceDescriptor with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResourceDescriptorMultiError, or nil if none found.
func (m *ResourceDescriptor) ValidateAll() error {
	return m.validate(true)
}

func (m *ResourceDescriptor) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for NameField

	// no validation rules for History

	if len(errors) > 0 {
		return ResourceDescriptorMultiError(errors)
	}

	return nil
}

// ResourceDescriptorMultiError is an error wrapping multiple validation errors
// returned by ResourceDescriptor.ValidateAll() if the designated constraints
// aren't met.
type ResourceDescriptorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResourceDescriptorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResourceDescriptorMultiError) AllErrors() []error { return m }

// ResourceDescriptorValidationError is the validation error returned by
// ResourceDescriptor.Validate if the designated constraints aren't met.
type ResourceDescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceDescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceDescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceDescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceDescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceDescriptorValidationError) ErrorName() string {
	return "ResourceDescriptorValidationError"
}

// Error satisfies the builtin error interface
func (e ResourceDescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourceDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceDescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceDescriptorValidationError{}

// Validate checks the field values on ResourceReference with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ResourceReference) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResourceReference with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResourceReferenceMultiError, or nil if none found.
func (m *ResourceReference) ValidateAll() error {
	return m.validate(true)
}

func (m *ResourceReference) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for ChildType

	if len(errors) > 0 {
		return ResourceReferenceMultiError(errors)
	}

	return nil
}

// ResourceReferenceMultiError is an error wrapping multiple validation errors
// returned by ResourceReference.ValidateAll() if the designated constraints
// aren't met.
type ResourceReferenceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResourceReferenceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResourceReferenceMultiError) AllErrors() []error { return m }

// ResourceReferenceValidationError is the validation error returned by
// ResourceReference.Validate if the designated constraints aren't met.
type ResourceReferenceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceReferenceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceReferenceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceReferenceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceReferenceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceReferenceValidationError) ErrorName() string {
	return "ResourceReferenceValidationError"
}

// Error satisfies the builtin error interface
func (e ResourceReferenceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourceReference.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceReferenceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceReferenceValidationError{}
