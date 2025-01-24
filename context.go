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

func (c *Context[T]) ErrMsg() string {
	return c.errMsg
}

func (c *Context[T]) RenameAdditionalKey(oldKey, newKey string) error {
	// Verifica se a chave antiga existe
	if _, exists := c.Additional[oldKey]; !exists {
		return ErrAttributeNotFound
	}

	// Se a nova chave já existir, retorna um erro
	if _, exists := c.Additional[newKey]; exists {
		return ErrKeyAlreadyExists
	}

	// Atribui o valor da chave antiga à nova chave
	c.Additional[newKey] = c.Additional[oldKey]

	// Remove a chave antiga
	delete(c.Additional, oldKey)

	return nil
}
