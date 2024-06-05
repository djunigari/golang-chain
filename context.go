package chain

type Context[T any] struct {
	ExtraContext[T]
	err                 error
	errMsg              string
	ActionErr           *Action[T]
	LastActionCalled    *Action[T]
	Additional          map[string]interface{}
	LoopAction          map[string]bool
	ActionFlowDirection string
}

func (c *Context[T]) SetErr(err error) *Context[T] {
	c.ActionErr = c.LastActionCalled
	c.err = err
	return c
}

func (c *Context[T]) SetErrMsg(msg string) *Context[T] {
	c.errMsg = msg
	return c
}

func (c *Context[T]) Err() error {
	return c.err
}
