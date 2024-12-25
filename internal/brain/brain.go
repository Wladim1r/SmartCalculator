package brain

import (
	"calc/internal/model"
	"errors"
	"math"
	"time"
)

var (
	ErrUnknownAction = errors.New("unknown action")
	ErrorWithZero    = errors.New("operation is not supported")
)

type Brain struct {
	delay int
}

func NewBrain(delay int) *Brain {
	return &Brain{delay: delay}
}

func (b *Brain) Calculate(operations []*model.Operation) {
	for _, operation := range operations {
		time.Sleep(time.Duration(b.delay) * time.Second)

		var result float64
		var resultErr error
		switch operation.Action {
		case model.ActionPlus:
			result = operation.Number1 + operation.Number2

		case model.ActionMinus:
			result = operation.Number1 - operation.Number2

		case model.ActionMultiple:
			result = operation.Number1 * operation.Number2

		case model.ActionDivision:
			if operation.Number2 != 0 {
				result = operation.Number1 / operation.Number2
			} else {
				resultErr = ErrorWithZero
			}

		case model.ActionPow:
			if operation.Number1 == 0 && operation.Number2 == 0 {
				resultErr = ErrorWithZero
			} else {
				result = math.Pow(operation.Number1, operation.Number2)
			}

		case model.ActionRemains:
			if operation.Number2 != 0 {
				result = math.Remainder(operation.Number1, operation.Number2)
			} else {
				resultErr = ErrorWithZero
			}

		default:
			resultErr = ErrUnknownAction
		}

		operation.Result = struct {
			Data float64
			Err  error
		}{
			Data: result, Err: resultErr,
		}

	}
}
