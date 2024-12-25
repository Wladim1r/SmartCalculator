package calculator

import (
	"calc/internal/brain"
	"calc/internal/model"
	"calc/internal/parser"
	"calc/internal/printer"
)

type Calculator struct {
	parser  parser.Parser
	brain   brain.Brain
	printer printer.Printer
}

func NewCalculator(
	calcParser parser.Parser,
	calcBrain brain.Brain,
	calcPrinter printer.Printer,
) *Calculator {
	o := new(Calculator)
	o.parser = calcParser
	o.printer = calcPrinter
	o.brain = calcBrain

	return o
}

func (c *Calculator) Run(rawData []model.UserInput, number int) {
	operations := c.parser.PrepareData()
	c.brain.Calculate(operations)
	c.printer.Print(operations, number)
}
