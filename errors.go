package chain

import (
	"errors"
	"fmt"

	logger "github.com/djunigari/golang-logger"
)

var (
	ErrVariableNotFound    = errors.New("variable not found in context")
	ErrInvalidVariableType = errors.New("invalid type of variable")
	ErrInvalidVariable     = errors.New("invalid variable")
	ErrKeyAlreadyExists    = errors.New("key already exists in context")
	ErrAttributeNotFound   = errors.New("attribute not found")
)

func LogError[T any](parentName string) *Action[T] {
	return NewAction[T]("chains.LogError").
		IgnoreError(true).
		Function(func(ctx *Context[T]) {
			if ctx.Err() != nil {
				var errorDetails string
				if ctx.ActionErr != nil {
					errorDetails = fmt.Sprintf("[%s] %s : %s", ctx.ActionErr.Name, ctx.Err(), ctx.ErrMsg())
				} else {
					errorDetails = fmt.Sprintf("%s : %s", ctx.Err(), ctx.ErrMsg())
				}
				logger.LogError("failed "+parentName, errorDetails)
			}
		})
}
