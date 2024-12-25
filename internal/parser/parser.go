package parser

import (
	"calc/internal/model"
	"fmt"
	"os"
)

type Parser struct {
	rawData []model.UserInput
	actions model.Actions
}

func NewParser(raw []model.UserInput, actions ...model.Action) *Parser {
	return &Parser{
		rawData: raw,
		actions: actions,
	}
}

func (p *Parser) PrepareData() []*model.Operation {
	var operations []*model.Operation
	for _, ui := range p.rawData {
		ok := p.actions.IsAction(ui.Action)
		if !ok {
			fmt.Fprintf(os.Stderr, "Оператор %s не поддерживается\n", ui.Action)
			continue
		}
		operation := model.NewOperation(ui.N1, model.Action(ui.Action), ui.N2)
		operations = append(operations, &operation)
	}

	return operations
}
