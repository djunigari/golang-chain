package chain

type ActionFunc[T any] func(ctx *Context[T])

type Actions[T any] []*Action[T]

type Action[T any] struct {
	ActionFunc    *ActionFunc[T]
	ActionOptions *ActionOptions
}

type ActionOptions struct {
	IgnoreError bool
}

func NewAction[T any](name string) *Action[T] {
	return &Action[T]{
		ActionFunc:    nil,
		ActionOptions: nil,
	}
}

func (a *Action[T]) Function(function *ActionFunc[T]) *Action[T] {
	a.ActionFunc = function
	return a
}

func (a *Action[T]) Options(options *ActionOptions) *Action[T] {
	a.ActionOptions = options
	return a
}
