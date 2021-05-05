package errors

import "net/http"

// BadRequest new BadRequest error that is mapped to a 400 response.
func BadRequest(domain, reason, message string) *ErrorInfo {
	return Errorf(http.StatusBadRequest, domain, reason, message)
}

// IsBadRequest determines if err is an error which indicates a BadRequest error.
func IsBadRequest(err error) bool {
	return Code(err) == http.StatusBadRequest
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func Unauthorized(domain, reason, message string) *ErrorInfo {
	return Errorf(http.StatusUnauthorized, domain, reason, message)
}

// IsUnauthorized determines if err is an error which indicates a Unauthorized error.
func IsUnauthorized(err error) bool {
	return Code(err) == http.StatusUnauthorized
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func Forbidden(domain, reason, message string) *ErrorInfo {
	return Errorf(http.StatusForbidden, domain, reason, message)
}

// IsForbidden determines if err is an error which indicates a Forbidden error.
func IsForbidden(err error) bool {
	return Code(err) == http.StatusForbidden
}

// NotFound new NotFound error that is mapped to a 404 response.
func NotFound(domain, reason, message string) *ErrorInfo {
	return Errorf(http.StatusNotFound, domain, reason, message)
}

// IsNotFound determines if err is an error which indicates an NotFound error.
func IsNotFound(err error) bool {
	return Code(err) == http.StatusNotFound
}

// Conflict new Conflict error that is mapped to a 409 response.
func Conflict(domain, reason, message string) *ErrorInfo {
	return Errorf(http.StatusConflict, domain, reason, message)
}

// IsConflict determines if err is an error which indicates a Conflict error.
func IsConflict(err error) bool {
	return Code(err) == http.StatusConflict
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func InternalServer(domain, reason, message string) *ErrorInfo {
	return Errorf(http.StatusInternalServerError, domain, reason, message)
}

// IsInternalServer determines if err is an error which indicates an InternalServer error.
func IsInternalServer(err error) bool {
	return Code(err) == http.StatusInternalServerError
}

// ServiceUnavailable new ServiceUnavailable error that is mapped to a HTTP 503 response.
func ServiceUnavailable(domain, reason, message string) *ErrorInfo {
	return Errorf(http.StatusServiceUnavailable, domain, reason, message)
}

// IsServiceUnavailable determines if err is an error which indicates a ServiceUnavailable error.
func IsServiceUnavailable(err error) bool {
	return Code(err) == http.StatusServiceUnavailable
}

// InvalidArgument new InvalidArgument error that is mapped to a 400 response.
func InvalidArgument(domain, message string) *ErrorInfo {
	return BadRequest(domain, "invalid_argument", message)
}

// IsInvalidArgument determines if err is an error which indicates a InvalidArgument error.
func IsInvalidArgument(err error) bool {
	return Reason(err) == "invalid_argument"
}

// AlreadyExists new AlreadyExists error that is mapped to a 409 response.
func AlreadyExists(domain, message string) *ErrorInfo {
	return Conflict(domain, "already_exists", message)
}

// IsAlreadyExists determines if err is an error which indicates a AlreadyExists error.
func IsAlreadyExists(err error) bool {
	return Reason(err) == "already_exists"
}

// DatabaseError new DatabaseError error that is mapped to a 409 response.
func DatabaseError(domain, message string) *ErrorInfo {
	return InternalServer(domain, "database_error", message)
}

// IsDatabaseError determines if err is an error which indicates a DatabaseError error.
func IsDatabaseError(err error) bool {
	return Reason(err) == "database_error"
}
