package chain

import "errors"

var (
	ErrVariableNotFound    = errors.New("variable not found in context")
	ErrInvalidVariableType = errors.New("invalid type of variable")
)
