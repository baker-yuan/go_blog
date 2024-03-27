package consts

import (
	"errors"

	"github.com/shiningrush/droplet/data"
)

const (
	ErrBadRequest = 20001
	ErrForbidden  = 20002
)

const (
	// IDNotFound is the string use for can't find the cache by the id
	IDNotFound = "%s id: %s not found"
)

var (
	// ErrorUsernamePassword is the error means username or password is not correct
	ErrUsernamePassword = errors.New("username or password error")
	// ErrorIDUsername is the error use for the input's id and username is different
	ErrIDUsername = errors.New("consumer's id and username must be a same value")
	// ErrorParameterID is the error use for parameter ID is empty
	ErrParameterID = errors.New("Parameter IDs cannot be empty")
	// ErrorRouteData is the error that the route data is empty
	ErrRouteData = errors.New("Route data is empty, cannot be exported")
	// ErrorImportFile is the error that use for import a empty file
	ErrImportFile = errors.New("empty or invalid imported file")
	// ErrorImportFile means the certificate is invalid
	ErrSSLCertificate = errors.New("invalid certificate")
	// ErrorSSLCertificateResolution means the SSL certificate decode failed
	ErrSSLCertificateResolution = errors.New("Certificate resolution failed")
	// ErrorSSLKeyAndCert means the SSL key and SSL certificate don't match
	ErrSSLKeyAndCert = errors.New("key and cert don't match")
)

var (
	// base error please refer to github.com/shiningrush/droplet/data, such as data.ErrNotFound, data.ErrConflicted
	ErrInvalidRequest       = data.BaseError{Code: ErrBadRequest, Message: "invalid request"}
	ErrSchemaValidateFailed = data.BaseError{Code: ErrBadRequest, Message: "JSONSchema validate failed"}
	ErrIPNotAllow           = data.BaseError{Code: ErrForbidden, Message: "IP address not allowed"}
)
