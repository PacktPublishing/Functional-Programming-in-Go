package main

import (
	"fmt"
)

func isAuthenticated() bool {
	fmt.Println("isAuthenticated()")
	return false
}

func isAdmin() bool {
	fmt.Println("isAdmin()")
	return true
}

func main() {
	if isAuthenticated() && isAdmin() {
		fmt.Println("is authenticated and admin")
	}

	if isAuthenticated() || isAdmin() {
		fmt.Println("is authenticated or admin")
	}

	if isAdmin() || isAuthenticated() {
		fmt.Println("admin or authenticated")
	}

}
