package main

import (
	"bufio"
	"calc/internal/brain"
	"calc/internal/calculator"
	"calc/internal/model"
	"calc/internal/parser"
	"calc/internal/printer"
	"fmt"
	"os"
)

const (
	done    = "done\r\n"
	newLine = '\n'
)

func main() {

	fmt.Println()
	fmt.Fprintln(os.Stdout, "Добро пожаловать в Калькулятор!")
	fmt.Fprintln(os.Stdout, "Убедительная просьба ознакомиться с инструкцией использования")
	fmt.Fprintln(os.Stdout, "Маска выражения: <число1><пробел><действие><пробел><число2>")
	fmt.Fprintln(os.Stdout, "Возможные действия: -, +, *, /, ^, %")
	fmt.Fprintln(os.Stdout, "Пример записи выражения: 2 - 2")
	fmt.Fprintln(os.Stdout, "Пример записи выражения: -4 + -3")
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout, "Выберите с какой точностью будут проводиться расчеты")
	fmt.Fprintln(os.Stdout, "0 - округление ответа до целого числа")
	fmt.Fprintln(os.Stdout, "1 - округление ответа до одного знака после запятой")
	fmt.Fprintln(os.Stdout, "2 - округление ответа до двух знаков после запятой")
	fmt.Fprintln(os.Stdout, "3 - округление ответа до трех знаков после запятой")
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout, "После выбора варианта точности, можете вводить выражения для вычислений или же стоп-слово \"done\"")
	fmt.Fprintln(os.Stdout)

	var userInputs []model.UserInput

	numberOption := printer.ChooseOption()

	var k int

	for {
		if k != 0 {
			fmt.Fprint(os.Stdout, "Поле для ввода: ")
		}
		k++

		in := bufio.NewReader(os.Stdin)
		str, err := in.ReadString(newLine)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода", err)
			continue
		}

		if str == done {
			break
		}

		if str == string(newLine) {
			continue
		}

		var n1, n2 float64
		var action string

		n, err := fmt.Sscanf(str, "%f %s %f", &n1, &action, &n2)
		if n != 3 || err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода", err)
			continue
		}

		ui := model.UserInput{
			N1:     n1,
			Action: action,
			N2:     n2,
		}
		userInputs = append(userInputs, ui)
	}

	if len(userInputs) == 0 {
		fmt.Fprintln(os.Stderr, "Нет выражений - нет вычислений")
		return
	}

	calcParser := parser.NewParser(
		userInputs,
		model.ActionPlus,
		model.ActionMinus,
		model.ActionMultiple,
		model.ActionDivision,
		model.ActionPow,
		model.ActionRemains,
	)
	calcBrain := brain.NewBrain(1)
	calcPrinter := printer.NewPrinter()

	appCalculator := calculator.NewCalculator(*calcParser, *calcBrain, *calcPrinter)
	appCalculator.Run(userInputs, numberOption)
}
