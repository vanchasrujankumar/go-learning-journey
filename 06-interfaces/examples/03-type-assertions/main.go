package main

import "fmt"

func describe(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("%d is an integer\n", val)
	case float64:
		fmt.Printf("%.2f is a float64\n", val)
	case string:
		fmt.Printf("%q is a string of length %d\n", val, len(val))
	case bool:
		fmt.Printf("%v is a boolean\n", val)
	case []int:
		fmt.Printf("%v is an int slice with length %d\n", val, len(val))
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	fmt.Println("=== Type Assertions ===")
	var i interface{} = "hello, go"

	s := i.(string)
	fmt.Println("Asserted string:", s)

	n, ok := i.(int)
	if ok {
		fmt.Println("Asserted int:", n)
	} else {
		fmt.Println("Not an int (no panic - ok is false)")
	}

	fmt.Println("\n=== Type Switch ===")
	values := []interface{}{42, 3.14, "hello", true, []int{1, 2, 3}, struct{}{}}
	for _, v := range values {
		describe(v)
	}
}
