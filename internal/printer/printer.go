package printer

import (
	"calc/internal/model"
	"fmt"
	"os"
	"strconv"
)

type Printer struct {
}

func NewPrinter() *Printer {
	o := new(Printer)
	return o
}

func ChooseOption() int {
	var number int

	for {
		fmt.Fprint(os.Stdout, "Точность округления: ")

		var str string
		fmt.Scan(&str)

		number, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stdout, "Используйте только указанные цифры(0, 1, 2, 3)")
			continue
		}

		if number < 0 || number > 3 {
			fmt.Fprintln(os.Stderr, "Делайте выбор на основе предложенных вариантов")
			continue
		}
		break
	}
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout, "Калькулятор готов принимать выражения для счета")
	return number
}

func (p *Printer) Print(operations []*model.Operation, number int) {
	for _, o := range operations {
		if o.Result.Err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка вычисления!", o.Result.Err)
			continue
		}

		switch number {
		case 0:
			fmt.Fprintf(os.Stdout, "%v %s %v = %.0f\n", o.Number1, o.Action, o.Number2, o.Result.Data)
		case 1:
			fmt.Fprintf(os.Stdout, "%v %s %v = %.1f\n", o.Number1, o.Action, o.Number2, o.Result.Data)
		case 2:
			fmt.Fprintf(os.Stdout, "%v %s %v = %.2f\n", o.Number1, o.Action, o.Number2, o.Result.Data)
		case 3:
			fmt.Fprintf(os.Stdout, "%v %s %v = %.3f\n", o.Number1, o.Action, o.Number2, o.Result.Data)
		}
	}
}
