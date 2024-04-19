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
		err:        nil,
		Additional: make(map[string]interface{}),
		ExtraContext: ExtraContext[T]{
			Extra: *extra,
		},
	}

	for _, action := range *p.Actions {
		if action.IgnoreError {
			ctx.LastActionCalled = &action
			action.Function(ctx)
		}
	}
}
