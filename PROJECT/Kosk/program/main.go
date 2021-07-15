package main

import (
	"fmt"
)

func double(a interface{}) interface{} {
	x, ok := a.(int)
	if ok {
		return 2 * x
	}

	y, ok := a.(string)
	if ok {
		return fmt.Sprintf("%s%s", y, y)
	}

	z, ok := a.(bool)
	if ok {
		z = !z
		return z
	}

	return nil
}

/*
func double(a interface{}) interface{} {

    switch a.(type) {
    case int:
        return 2*a.(int)
    case string:
        return fmt.Sprintf("%s%s",a,a)
    case bool:
        return !a.(bool)
    }

    return nil
}
*/
func main() {
	a := []interface{}{42, "Hello", false}

	for _, v := range a {
		fmt.Println(v)
		fmt.Println(double(v))
	}
}
