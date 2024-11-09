package auth

import "fmt"

func LoginWithCredentials(username, password string) { // if the function name starts with smallcase then we can only use the funciton in the same package means i cant use in main.go as it is in main package but if we keep the name in upper case then we can use it in main.go
	fmt.Println("Logging User", username, password)
}
