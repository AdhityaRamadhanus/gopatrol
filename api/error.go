package api

const (
	ErrFailedToReadBody       = "Failed to read body of http request"
	ErrFailedToCloseBody      = "Failed to close body of http request"
	ErrFailedToMarshalJSON    = "Failed to marshal json"
	ErrFailedToUnmarshalJSON  = "Failed to unmarshal json"
	ErrFailedToValidateStruct = "Failed to validate an object/request"
	ErrInternalServerError    = "Internal Server Error"

	// Specific TYpe Errors
	ErrUnknownEndpointTYpe = "Unknown endpoint type"
)
