package chain

type Processor[T any] struct {
	Actions *Actions[T]
}

func New[T any](actions *Actions[T]) *Processor[T] {
	return &Processor[T]{
		Actions: actions,
	}
}

func (p *Processor[T]) Run(extra *T) {
	ctx := &Context[T]{
		Err:        nil,
		Additional: make(map[string]interface{}),
		ExtraContext: ExtraContext[T]{
			Extra: *extra,
		},
	}

	for _, handler := range *p.Actions {
		handler(ctx)
	}
}
