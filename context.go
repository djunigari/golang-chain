package chain

type Context[T any] struct {
	ExtraContext[T]
	Err        error
	ActionErr  Action[T]
	Additional map[string]interface{}
}
