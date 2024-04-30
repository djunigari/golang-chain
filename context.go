package chain

type Context[T any] struct {
	ExtraContext[T]
	err                 error
	ActionErr           *Action[T]
	LastActionCalled    *Action[T]
	Additional          map[string]interface{}
	ActionFlowDirection *string
}

func (c *Context[T]) SetErr(err error) {
	c.ActionErr = c.LastActionCalled
	c.err = err
}

func (c *Context[T]) Err() error {
	return c.err
}
