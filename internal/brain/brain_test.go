package brain

import (
	"calc/internal/model"
	"testing"
)

func TestCalculatePlus(t *testing.T) {
	const delay = 0
	b := NewBrain(delay)

	var opeations []*model.Operation

	expected := []float64{4, 0}

	opPlus := model.NewOperation(2, model.ActionPlus, 2)
	opeations = append(opeations, &opPlus)

	opMinus := model.NewOperation(2, model.ActionMinus, 2)
	opeations = append(opeations, &opMinus)

	b.Calculate(opeations)

	for i, o := range opeations {
		if o.Result.Err != nil {
			t.Log("result err is not nil")
			t.FailNow()
		}
		if o.Result.Data != expected[i] {
			t.Logf("want %v, but got %v", expected, o.Result.Data)
			t.FailNow()
		}
	}
}
