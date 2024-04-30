package chain

type ActionFunc[T any] func(ctx *Context[T])

type Actions[T any] []*Action[T]

type ActionType string

const (
	SingleAction ActionType = "SingleAction"
	FlowAction   ActionType = "FlowAction"
)

type Action[T any] struct {
	Name          string
	ActionFunc    *ActionFunc[T]
	ActionOptions *ActionOptions
	ActionType    ActionType
	FlowDirection string
	SubActions    map[string]Actions[T]
}

type ActionOptions struct {
	IgnoreError bool
}

func NewAction[T any](name string) *Action[T] {
	return &Action[T]{
		Name:          name,
		ActionFunc:    nil,
		ActionOptions: &ActionOptions{IgnoreError: false},
		ActionType:    SingleAction,
		FlowDirection: "none",
		SubActions:    nil,
	}
}

func (a *Action[T]) Function(function *ActionFunc[T]) *Action[T] {
	a.ActionFunc = function
	return a
}

func (a *Action[T]) Options(options *ActionOptions) *Action[T] {
	a.ActionOptions = options
	return a
}

func (a *Action[T]) Type(actionType ActionType) *Action[T] {
	a.ActionType = actionType
	return a
}

func (a *Action[T]) AddSubActions(condition string, actions Actions[T]) *Action[T] {
	if a.SubActions == nil {
		a.SubActions = make(map[string]Actions[T])
	}
	a.SubActions[condition] = actions
	return a
}
