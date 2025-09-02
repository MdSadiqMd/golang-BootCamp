// panic - used to panic - only use in case of unrecoverable errors
// recover - mostly placed in defer functions - also used in cleaning up resources
// we can add mutiple recovers, like take we have a db request failes, we have a recover in that, also add a panic at last of that recover function
// through this the panic will propagate to upper level code of services and so on until you need to
package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	panic("Panic!!")
}
