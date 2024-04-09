package chain

type Action[T any] func(ctx *Context[T])

type Actions[T any] []Action[T]
