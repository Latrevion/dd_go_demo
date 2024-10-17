package main

import "fmt"

func main() {

	// var (
	// 	username string
	// 	age      int
	// 	sex      string
	// )

	// username = "lily"
	// age = 12
	// sex = "man"

	// fmt.Println(username, age, sex)

	// username := "luke"
	// fmt.Printf("%T\n", username)
	// fmt.Println(username)

	// a, b, c := 1, 2, 3
	// fmt.Println(a, b, c)

	var username, _ = getUserinfo()
	fmt.Println(username)
}
func getUserinfo() (string, int) {
	return "luke", 10
}
