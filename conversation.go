package main

import "fmt"

func main() {
	fmt.Println("Enter 1 for fahreneit to kelvin")
	fmt.Println("Enter 2 for kelvin to fahreneit")
	fmt.Println("Enter 3 for celcius to fahreneit")
	fmt.Println("Enter 4 for fahreneit to celcius")
	fmt.Println("Enter 5 for celcius to kelvin")
	fmt.Println("Enter 6 for kelvin to celcius")
	var value int
	var degree float64
	fmt.Scan(&value)
	fmt.Println("Enter the value")
	fmt.Scan(&degree)

	// checks the input string or zero value.
	if value > 6 && 1 > value {
		fmt.Println(fmt.Print("please enter valid number."))
	} else if value == 1 {
		fmt.Println(fahreneit2Kelvin(degree))
	} else if value == 2 {
		fmt.Println(kelvin2Fahreneit(degree))
	} else if value == 3 {
		fmt.Println(celcius2Fahreneit(degree))
	} else if value == 4 {
		fmt.Println(fahreneit2Celcius(degree))
	} else if value == 5 {
		fmt.Println(celcius2Kelvin(degree))
	} else if value == 6 {
		fmt.Println(kelvin2Celcius(degree))
	}

}
func celcius2Kelvin(celcius float64) float64 {
	return (celcius + 273.0)
}

func celcius2Fahreneit(celcius float64) float64 {
	return (celcius * 1.8) + 32
}

func fahreneit2Celcius(fahreneit float64) float64 {
	return 5 * (fahreneit - 320) / 32
}
func kelvin2Celcius(kelvin float64) float64 {
	return kelvin - 273
}

func kelvin2Fahreneit(kelvin float64) float64 {
	return celcius2Fahreneit(kelvin2Celcius(kelvin))
}

func fahreneit2Kelvin(fahreneit float64) float64 {
	return celcius2Kelvin(fahreneit2Celcius(fahreneit))
}
