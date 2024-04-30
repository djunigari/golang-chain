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
		err:                 nil,
		Additional:          make(map[string]interface{}),
		ActionFlowDirection: nil,
		ExtraContext: ExtraContext[T]{
			Extra: *extra,
		},
	}

	p.execActions(ctx, p.Actions)
}

func (p *Processor[T]) execActions(ctx *Context[T], actions *Actions[T]) {
	for _, action := range *actions {
		if p.PrintLog {
			fmt.Println(action.Name)
		}

		if ctx.err != nil && !action.ActionIgnoreError {
			continue
		}

		switch {
		case action.ActionType == SingleAction:
			ctx.LastActionCalled = action
			(*action.ActionFunc)(ctx)
		case action.ActionType == FlowAction:
			ctx.LastActionCalled = action
			(*action.ActionFunc)(ctx)
			if ctx.ActionFlowDirection == "" {
				continue
			}
			if subs, ok := action.SubActions[ctx.ActionFlowDirection]; ok {
				p.execActions(ctx, &subs)
			}
		default:
			ctx.LastActionCalled = action
			(*action.ActionFunc)(ctx)
		}
	}
}
