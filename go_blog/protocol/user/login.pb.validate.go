// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/login.proto

package user

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

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReqMultiError, or nil
// if none found.
func (m *LoginReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LoginReqMultiError(errors)
	}

	return nil
}

// LoginReqMultiError is an error wrapping multiple validation errors returned
// by LoginReq.ValidateAll() if the designated constraints aren't met.
type LoginReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReqMultiError) AllErrors() []error { return m }

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

// Validate checks the field values on LoginRsp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRspMultiError, or nil
// if none found.
func (m *LoginRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LoginRspMultiError(errors)
	}

	return nil
}

// LoginRspMultiError is an error wrapping multiple validation errors returned
// by LoginRsp.ValidateAll() if the designated constraints aren't met.
type LoginRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRspMultiError) AllErrors() []error { return m }

// LoginRspValidationError is the validation error returned by
// LoginRsp.Validate if the designated constraints aren't met.
type LoginRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRspValidationError) ErrorName() string { return "LoginRspValidationError" }

// Error satisfies the builtin error interface
func (e LoginRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRspValidationError{}

// Validate checks the field values on LogoutReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogoutReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogoutReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogoutReqMultiError, or nil
// if none found.
func (m *LogoutReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LogoutReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LogoutReqMultiError(errors)
	}

	return nil
}

// LogoutReqMultiError is an error wrapping multiple validation errors returned
// by LogoutReq.ValidateAll() if the designated constraints aren't met.
type LogoutReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogoutReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogoutReqMultiError) AllErrors() []error { return m }

// LogoutReqValidationError is the validation error returned by
// LogoutReq.Validate if the designated constraints aren't met.
type LogoutReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutReqValidationError) ErrorName() string { return "LogoutReqValidationError" }

// Error satisfies the builtin error interface
func (e LogoutReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutReqValidationError{}

// Validate checks the field values on LogoutRsp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogoutRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogoutRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogoutRspMultiError, or nil
// if none found.
func (m *LogoutRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *LogoutRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LogoutRspMultiError(errors)
	}

	return nil
}

// LogoutRspMultiError is an error wrapping multiple validation errors returned
// by LogoutRsp.ValidateAll() if the designated constraints aren't met.
type LogoutRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogoutRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogoutRspMultiError) AllErrors() []error { return m }

// LogoutRspValidationError is the validation error returned by
// LogoutRsp.Validate if the designated constraints aren't met.
type LogoutRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutRspValidationError) ErrorName() string { return "LogoutRspValidationError" }

// Error satisfies the builtin error interface
func (e LogoutRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutRspValidationError{}

// Validate checks the field values on RefreshReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RefreshReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RefreshReqMultiError, or
// nil if none found.
func (m *RefreshReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RefreshReqMultiError(errors)
	}

	return nil
}

// RefreshReqMultiError is an error wrapping multiple validation errors
// returned by RefreshReq.ValidateAll() if the designated constraints aren't met.
type RefreshReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshReqMultiError) AllErrors() []error { return m }

// RefreshReqValidationError is the validation error returned by
// RefreshReq.Validate if the designated constraints aren't met.
type RefreshReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshReqValidationError) ErrorName() string { return "RefreshReqValidationError" }

// Error satisfies the builtin error interface
func (e RefreshReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshReqValidationError{}

// Validate checks the field values on RefreshRsp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RefreshRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RefreshRspMultiError, or
// nil if none found.
func (m *RefreshRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RefreshRspMultiError(errors)
	}

	return nil
}

// RefreshRspMultiError is an error wrapping multiple validation errors
// returned by RefreshRsp.ValidateAll() if the designated constraints aren't met.
type RefreshRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshRspMultiError) AllErrors() []error { return m }

// RefreshRspValidationError is the validation error returned by
// RefreshRsp.Validate if the designated constraints aren't met.
type RefreshRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshRspValidationError) ErrorName() string { return "RefreshRspValidationError" }

// Error satisfies the builtin error interface
func (e RefreshRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshRspValidationError{}
