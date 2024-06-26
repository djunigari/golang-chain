package chain

type ActionFunc[T any] func(ctx *Context[T])

type Actions[T any] []*Action[T]

type ActionType string

const (
	SingleAction ActionType = "SingleAction"
	FlowAction   ActionType = "FlowAction"
	LoopAction   ActionType = "LoopAction"
)

type Action[T any] struct {
	Name              string
	ActionFunc        *ActionFunc[T]
	ActionIgnoreError bool
	ActionType        ActionType
	FlowDirection     string
	SubActions        map[string]Actions[T]
}

func NewAction[T any](name string) *Action[T] {
	return &Action[T]{
		Name:              name,
		ActionFunc:        nil,
		ActionIgnoreError: false,
		ActionType:        SingleAction,
		SubActions:        nil,
	}
}

func (a *Action[T]) Function(function ActionFunc[T]) *Action[T] {
	a.ActionFunc = &function
	return a
}

func (a *Action[T]) IgnoreError(ignoreError bool) *Action[T] {
	a.ActionIgnoreError = ignoreError
	return a
}

func (a *Action[T]) Type(actionType ActionType) *Action[T] {
	a.ActionType = actionType
	return a
}

func (a *Action[T]) Loop(actions ...*Action[T]) *Action[T] {
	if a.SubActions == nil {
		a.SubActions = make(map[string]Actions[T])
	}
	a.SubActions[a.Name] = actions
	return a
}

func (a *Action[T]) Flow(condition string, actions ...*Action[T]) *Action[T] {
	if a.SubActions == nil {
		a.SubActions = make(map[string]Actions[T])
	}
	a.SubActions[condition] = actions
	return a
}
