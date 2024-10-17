package chain

import (
	"fmt"

	logger "github.com/djunigari/golang-logger"
)

type Processor[T any] struct {
	Name     string
	Actions  *Actions[T]
	PrintLog bool
}

func New[T any](name string, actions *Actions[T], printLog bool) *Processor[T] {
	return &Processor[T]{
		Name:     name,
		Actions:  actions,
		PrintLog: printLog,
	}
}

func (p *Processor[T]) AddActions(actions *Actions[T]) {
	if p.Actions != nil {
		*p.Actions = append(*p.Actions, *actions...)
		return
	}
	p.Actions = actions
}

func (p *Processor[T]) AddAction(action *Action[T]) {
	if p.Actions != nil {
		*p.Actions = append(*p.Actions, action)
		return
	}
	p.Actions = &Actions[T]{action}
}

func (p *Processor[T]) LogError(ctx *Context[T]) {
	if ctx.Err() == nil {
		return
	}

	var errorDetails string
	if ctx.ActionErr != nil {
		errorDetails = fmt.Sprintf("[%s] %s : %s", ctx.ActionErr.Name, ctx.Err(), ctx.ErrMsg())
	} else {
		errorDetails = fmt.Sprintf("%s : %s", ctx.Err(), ctx.ErrMsg())
	}
	logger.LogError("failed "+p.Name, errorDetails)
}

func (p *Processor[T]) Run(extra *T) {
	ctx := &Context[T]{
		err:                 nil,
		errMsg:              "",
		Additional:          make(map[string]interface{}),
		ActionFlowDirection: "",
		ExtraContext: ExtraContext[T]{
			Extra: *extra,
		},
		LoopAction: make(map[string]bool),
	}

	p.execActions(ctx, p.Actions)
	p.LogError(ctx)
}

func (p *Processor[T]) execAction(ctx *Context[T], action *Action[T]) {
	if p.PrintLog {
		fmt.Println(action.Name)
	}
	ctx.LastActionCalled = action
	(*action.ActionFunc)(ctx)
}

func (p *Processor[T]) execActions(ctx *Context[T], actions *Actions[T]) {
	for _, action := range *actions {
		if ctx.err != nil && !action.ActionIgnoreError {
			continue
		}

		switch {
		case action.ActionType == SingleAction:
			p.execAction(ctx, action)
		case action.ActionType == FlowAction:
			p.execAction(ctx, action)
			if ctx.ActionFlowDirection == "" {
				continue
			}
			if subs, ok := action.SubActions[ctx.ActionFlowDirection]; ok {
				p.execActions(ctx, &subs)
			}
		case action.ActionType == LoopAction:
			p.execAction(ctx, action)
			if p.PrintLog {
				fmt.Println("starting loop :", action.Name)
			}
			for ctx.LoopAction[action.Name] {
				if subs, ok := action.SubActions[action.Name]; ok {
					p.execActions(ctx, &subs)
				}
				if ctx.err != nil && !action.ActionIgnoreError {
					delete(ctx.LoopAction, action.Name)
					break
				}
			}
			if p.PrintLog {
				fmt.Println("end loop :", action.Name)
			}
		default:
			p.execAction(ctx, action)
		}
	}
}
