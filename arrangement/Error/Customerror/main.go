package main

import "fmt"

type overHeatError float64

func (o overHeatError) Error() string {
	return fmt.Sprintf("over heat is : %0.2f", float64(o))
}

func checkTemperatur(actual float64, criteria float64) error {
	excess := actual - criteria
	if excess > 0 {
		return overHeatError(excess)
	}
	return nil
}

func main() {
	err := checkTemperatur(38.5, 37.5)
	if err != nil {
		fmt.Println(err)
	}
}
