package common

import "errors"

var (
	ResponseInvalidID        = map[string]any{"error": "invalid id"}
	ResponseInternalError    = map[string]any{"error": "internal server error"}
	ResponseResourceExisted  = map[string]any{"error": "resource existed"}
	ResponseResourceNotFound = map[string]any{"error": "not found"}
	ResponseBadRequest       = map[string]any{"error": "bad request"}

	ErrResourceExisted  = errors.New("resource existed")
	ErrResourceNotFound = errors.New("resource not found")
)
