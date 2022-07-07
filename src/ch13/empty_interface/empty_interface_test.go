package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomeThing(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("integer:", i)
	// 	return
	// }
	// if i, ok := p.(string); ok {
	// 	fmt.Printf("string: %v\n", i)
	// 	return
	// }
	// fmt.Println("Unknown type")
	switch v := p.(type) {
	case int:
		fmt.Println("integer:", v)
	case string:
		fmt.Printf("string: %v\n", v)
	default:
		fmt.Println("Unknown type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomeThing(10)
	DoSomeThing("hello")
	// DoSomeThing()
}
