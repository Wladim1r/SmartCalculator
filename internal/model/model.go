package model

type UserInput struct {
	N1     float64
	Action string
	N2     float64
}

type Operation struct {
	Number1 float64
	Action  Action
	Number2 float64
	Result  struct {
		Data float64
		Err  error
	}
}

func NewOperation(n1 float64, op Action, n2 float64) Operation {
	return Operation{
		Number1: n1,
		Action:  op,
		Number2: n2,
	}
}

type Action string

const (
	ActionPlus     Action = "+"
	ActionMinus    Action = "-"
	ActionMultiple Action = "*"
	ActionDivision Action = "/"
	ActionPow      Action = "^"
	ActionRemains  Action = "%"
)

type Actions []Action

func (actions Actions) IsAction(rawAction string) bool {
	for _, action := range actions {
		if action == Action(rawAction) {
			return true
		}
	}
	return false
}
