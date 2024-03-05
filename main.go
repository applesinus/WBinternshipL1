package main

import "fmt"

func main() {
	exercise := "none"
	for exercise != "exit" && exercise != "-1" {
		switch exercise {
		case "1":
			Ex1()
		}
		fmt.Println("Type here an exercise number ('exit or -1' to exit):")
		fmt.Scan(&exercise)
	}
}
