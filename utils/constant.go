package utils

const (
	API_VERSION                 = "/api/v1"
	MESSAGE_SUCCESS             = "Success"
	MESSAGE_OK                  = "OK"
	MESSAGE_TOO_MANY_REQUESTS   = "Too many requests, rate limit exceeded!"
	MESSAGE_FORBIDDEN           = "Error forbidden while extracting identifier"
	MESSAGE_UNAUTHORIZED        = "Unauthorized"
	MESSAGE_VALIDATION_ERROR    = "Validation failed"
	MESSAGE_INTERNAL_SERVER     = "Internal server error"
	MESSAGE_NOT_FOUND           = "Not found"
	MESSAGE_BAD_REQUEST         = "Invalid argument"
	MESSAGE_METHOD_NOT_ALLOWED  = "Method not allowed"
	MESSAGE_REQUEST_TIMEOUT     = "Request timeout"
	MESSAGE_SERVICE_UNAVAILABLE = "Service unavailable"
	MESSAGE_GRCP_SERVER_ERROR   = "GRPC server error"
)

const (
	SCOPE_MODULE   = "persediaan"
	SCOPE_VW_TABLE = "auth.vw_be_url_scopes"
)
