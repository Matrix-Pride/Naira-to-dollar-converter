package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// understanding the concept of types creation and alias
/*
	Alias is clone of a datatype, already existing primitive datatype
	The syntax is:
		type aliasName = datatype
*/

type Naria int64
type Dollar float32

func (value Naria) toDollar() (dollar Dollar) {
	return Dollar(value) * 0.00013
}

func (value Dollar) toNaria() (naria Naria) {
	res := 0.00013 / value
	return Naria(res)
}

func main() {
	// scanner that read  data  from the console
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Meet your Naria ¬ Dollar Converter")
	fmt.Println("-------------------")
	fmt.Println("                    -------------------")
	fmt.Println("Enter 1 to select Naria to Dollar convertion Or Enter 2 to select Dollar to Naria convertion")
	scanner.Scan()
	userChoice := scanner.Text()
	switch userChoice {
	case "1":
		fmt.Println("Enter the naria currency value i.e 20, 40")
		scanner.Scan()
		userInCurrency := scanner.Text()
		userInCurrencyToInt, err := strconv.Atoi(userInCurrency)
		if err == nil {
			nariaToDollar := Naria.toDollar(Naria(userInCurrencyToInt))
			fmt.Printf("%s Naria equals %v Dollar", userInCurrency, nariaToDollar)
		}
	case "2":
		fmt.Println("Enter the Dollar currency value i.e 20, 40")
		scanner.Scan()
		userInCurrency := scanner.Text()
		userInCurrencyToFloat, err := strconv.ParseFloat(userInCurrency, 32)
		if err == nil {
			dollarToNaria := Dollar.toNaria(Dollar(userInCurrencyToFloat))
			fmt.Printf("%s Dollar equals %v Naria", userInCurrency, dollarToNaria)
		}
	}
}
