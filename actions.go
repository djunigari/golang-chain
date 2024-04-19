package chain

type ActionFunc[T any] func(ctx *Context[T])

type Actions[T any] []Action[T]

type Action[T any] struct {
	Function    ActionFunc[T]
	IgnoreError bool
}

func NewAction[T any](function ActionFunc[T], ignoreError bool) *Action[T] {
	return &Action[T]{
		Function:    function,
		IgnoreError: ignoreError,
	}
}
