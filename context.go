package chain

type Context[T any] struct {
	ExtraContext[T]
	Err        error
	Additional map[string]interface{}
}
