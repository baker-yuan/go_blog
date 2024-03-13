// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/user.proto

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

// Validate checks the field values on EmptyRsp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EmptyRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmptyRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EmptyRspMultiError, or nil
// if none found.
func (m *EmptyRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *EmptyRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyRspMultiError(errors)
	}

	return nil
}

// EmptyRspMultiError is an error wrapping multiple validation errors returned
// by EmptyRsp.ValidateAll() if the designated constraints aren't met.
type EmptyRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyRspMultiError) AllErrors() []error { return m }

// EmptyRspValidationError is the validation error returned by
// EmptyRsp.Validate if the designated constraints aren't met.
type EmptyRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyRspValidationError) ErrorName() string { return "EmptyRspValidationError" }

// Error satisfies the builtin error interface
func (e EmptyRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyRspValidationError{}

// Validate checks the field values on AddOrUpdateRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddOrUpdateRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddOrUpdateRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddOrUpdateRspMultiError,
// or nil if none found.
func (m *AddOrUpdateRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *AddOrUpdateRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AddOrUpdateRspMultiError(errors)
	}

	return nil
}

// AddOrUpdateRspMultiError is an error wrapping multiple validation errors
// returned by AddOrUpdateRsp.ValidateAll() if the designated constraints
// aren't met.
type AddOrUpdateRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddOrUpdateRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddOrUpdateRspMultiError) AllErrors() []error { return m }

// AddOrUpdateRspValidationError is the validation error returned by
// AddOrUpdateRsp.Validate if the designated constraints aren't met.
type AddOrUpdateRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddOrUpdateRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddOrUpdateRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddOrUpdateRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddOrUpdateRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddOrUpdateRspValidationError) ErrorName() string { return "AddOrUpdateRspValidationError" }

// Error satisfies the builtin error interface
func (e AddOrUpdateRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddOrUpdateRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddOrUpdateRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddOrUpdateRspValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Salt

	// no validation rules for UserType

	// no validation rules for Email

	// no validation rules for Nickname

	// no validation rules for Avatar

	// no validation rules for Intro

	// no validation rules for WebSite

	// no validation rules for LoginType

	// no validation rules for UnionId

	// no validation rules for Status

	// no validation rules for IsDeleted

	// no validation rules for CreateTime

	// no validation rules for UpdateTime

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on SearchUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SearchUserReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchUserReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SearchUserReqMultiError, or
// nil if none found.
func (m *SearchUserReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchUserReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPageNum() <= 0 {
		err := SearchUserReqValidationError{
			field:  "PageNum",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPageSize() <= 0 {
		err := SearchUserReqValidationError{
			field:  "PageSize",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SearchUserReqMultiError(errors)
	}

	return nil
}

// SearchUserReqMultiError is an error wrapping multiple validation errors
// returned by SearchUserReq.ValidateAll() if the designated constraints
// aren't met.
type SearchUserReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchUserReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchUserReqMultiError) AllErrors() []error { return m }

// SearchUserReqValidationError is the validation error returned by
// SearchUserReq.Validate if the designated constraints aren't met.
type SearchUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchUserReqValidationError) ErrorName() string { return "SearchUserReqValidationError" }

// Error satisfies the builtin error interface
func (e SearchUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchUserReqValidationError{}

// Validate checks the field values on SearchUserRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SearchUserRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchUserRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SearchUserRspMultiError, or
// nil if none found.
func (m *SearchUserRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchUserRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SearchUserRspValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SearchUserRspValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchUserRspValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SearchUserRspMultiError(errors)
	}

	return nil
}

// SearchUserRspMultiError is an error wrapping multiple validation errors
// returned by SearchUserRsp.ValidateAll() if the designated constraints
// aren't met.
type SearchUserRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchUserRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchUserRspMultiError) AllErrors() []error { return m }

// SearchUserRspValidationError is the validation error returned by
// SearchUserRsp.Validate if the designated constraints aren't met.
type SearchUserRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchUserRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchUserRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchUserRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchUserRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchUserRspValidationError) ErrorName() string { return "SearchUserRspValidationError" }

// Error satisfies the builtin error interface
func (e SearchUserRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchUserRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchUserRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchUserRspValidationError{}

// Validate checks the field values on UserDetailReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserDetailReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserDetailReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserDetailReqMultiError, or
// nil if none found.
func (m *UserDetailReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UserDetailReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UserDetailReqValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserDetailReqMultiError(errors)
	}

	return nil
}

// UserDetailReqMultiError is an error wrapping multiple validation errors
// returned by UserDetailReq.ValidateAll() if the designated constraints
// aren't met.
type UserDetailReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserDetailReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserDetailReqMultiError) AllErrors() []error { return m }

// UserDetailReqValidationError is the validation error returned by
// UserDetailReq.Validate if the designated constraints aren't met.
type UserDetailReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserDetailReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserDetailReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserDetailReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserDetailReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserDetailReqValidationError) ErrorName() string { return "UserDetailReqValidationError" }

// Error satisfies the builtin error interface
func (e UserDetailReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserDetailReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserDetailReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserDetailReqValidationError{}

// Validate checks the field values on AddOrUpdateUserReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddOrUpdateUserReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddOrUpdateUserReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddOrUpdateUserReqMultiError, or nil if none found.
func (m *AddOrUpdateUserReq) ValidateAll() error {
	return m.validate(true)
}

func (m *AddOrUpdateUserReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Salt

	// no validation rules for UserType

	// no validation rules for Email

	// no validation rules for Nickname

	// no validation rules for Avatar

	// no validation rules for Intro

	// no validation rules for WebSite

	// no validation rules for LoginType

	// no validation rules for UnionId

	// no validation rules for Status

	// no validation rules for IsDeleted

	if len(errors) > 0 {
		return AddOrUpdateUserReqMultiError(errors)
	}

	return nil
}

// AddOrUpdateUserReqMultiError is an error wrapping multiple validation errors
// returned by AddOrUpdateUserReq.ValidateAll() if the designated constraints
// aren't met.
type AddOrUpdateUserReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddOrUpdateUserReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddOrUpdateUserReqMultiError) AllErrors() []error { return m }

// AddOrUpdateUserReqValidationError is the validation error returned by
// AddOrUpdateUserReq.Validate if the designated constraints aren't met.
type AddOrUpdateUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddOrUpdateUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddOrUpdateUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddOrUpdateUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddOrUpdateUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddOrUpdateUserReqValidationError) ErrorName() string {
	return "AddOrUpdateUserReqValidationError"
}

// Error satisfies the builtin error interface
func (e AddOrUpdateUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddOrUpdateUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddOrUpdateUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddOrUpdateUserReqValidationError{}

// Validate checks the field values on DeleteUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteUserReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteUserReqMultiError, or
// nil if none found.
func (m *DeleteUserReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteUserReqValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteUserReqMultiError(errors)
	}

	return nil
}

// DeleteUserReqMultiError is an error wrapping multiple validation errors
// returned by DeleteUserReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserReqMultiError) AllErrors() []error { return m }

// DeleteUserReqValidationError is the validation error returned by
// DeleteUserReq.Validate if the designated constraints aren't met.
type DeleteUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserReqValidationError) ErrorName() string { return "DeleteUserReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserReqValidationError{}
