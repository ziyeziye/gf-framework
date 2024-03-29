package response

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST     = 10001
	ERROR_NOT_EXIST = 10002

	ERROR_AUTH_TOKEN_FAIL    = 20001
	ERROR_AUTH_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN_CREATE  = 20003
	ERROR_AUTH               = 20004
)
