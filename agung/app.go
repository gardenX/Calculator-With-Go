package agung

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	add   = "+"
	sub   = "-"
	time  = "*"
	div   = "/"
	reset = "R"
	done  = "C"
)

func main() {
	var (
		total    int
		operator string
		value    int
		inputs   []string
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple calc")
	fmt.Println("--------------------")

	for {
		fmt.Println("Total :", total)
		fmt.Println("Input format <operator> <value>")
		fmt.Println("Supported operator (+ - * /) (C) close (R) reset")
		fmt.Print("Your input : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		inputs = strings.Split(input, " ")

		if strings.ToUpper(inputs[0]) == reset {
			total, value = resetStep()
			continue
		} else if strings.ToUpper(inputs[0]) == done {
			break
		}

		if !validOperator(inputs[0]) {
			fmt.Println("Invalid operator")
			fmt.Println("--------------------")
			continue
		}

		valueInt, err := strconv.Atoi(inputs[1])

		if err != nil {
			fmt.Println("Invalid value")
			fmt.Println("--------------------")
			continue
		}

		operator = inputs[0]
		value = valueInt

		if operator == add {
			total = Add(total, value)
		}

		if operator == sub {
			total = Sub(total, value)
		}

		if operator == time {
			total = Time(total, value)
		}

		if operator == div {
			total = Div(total, value)
		}
		fmt.Println("--------------------")
	}

	fmt.Println("Closed")
}

func resetStep() (int, int) {
	fmt.Println("--------------------")
	return 0, 0
}

func validOperator(input string) bool {
	operators := map[string]byte{add: 1, sub: 1, time: 1, div: 1}
	return operators[input] == 1
}
