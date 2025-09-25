package main

import (
	"fmt"
	"os"
	"strconv"
)

type Temperature struct {
    value float64
    unit  string
}
func celsiusToFahrenheit(t *Temperature) {
	t.value= t.value*9/5+32;
	t.unit="F"
	fmt.Println(t.value, t.unit)
	}
func fahrenheitToCelsius (t *Temperature) {
	t.value= (t.value-32)*5/9;
	t.unit="C"
	fmt.Println(t.value, t.unit)
	}

func main() {
	cmdArgs := os.Args
	var t1 Temperature
	inValue, err := strconv.ParseFloat(cmdArgs[1], 64)
    if err != nil {
        fmt.Println("Wrong input, first argument is float")
        return
    }

	if (cmdArgs[2] != "F" && cmdArgs[2] != "C") {
		fmt.Println("Wrong input, second argument is 'F' or 'C'")
		return
	} else {
		t1.unit = cmdArgs[2]
	}
	
	t1.value = inValue


	if (t1.unit == "F") {
		fahrenheitToCelsius(&t1)
	} else {
		celsiusToFahrenheit(&t1)
	}
}
