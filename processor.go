package chain

import (
	"fmt"
)

type Processor[T any] struct {
	Actions  *Actions[T]
	PrintLog bool
}

func New[T any](actions *Actions[T], printLog bool) *Processor[T] {
	return &Processor[T]{
		Actions:  actions,
		PrintLog: printLog,
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
		if p.PrintLog {
			fmt.Println(action.Name)
		}

		if ctx.err == nil || (action.ActionOptions != nil && action.ActionOptions.IgnoreError) {
			ctx.LastActionCalled = action
			(*action.ActionFunc)(ctx)
		}
	}
}
