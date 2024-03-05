package main

import "fmt"

func main() {
	exercise := "none"
	for exercise != "exit" && exercise != "-1" {
		switch exercise {
		case "1":
			Ex1()
		case "2":
			Ex2()
		case "3":
			Ex3()
		case "4":
			Ex4()
		case "5":
			Ex5()
		case "6":
			Ex6()
		case "7":
			Ex7()
		case "8":
			Ex8()
		case "9":
			Ex9()
		case "10":
			Ex10()
		case "11":
			Ex11()
		case "12":
			Ex12()
		case "13":
			Ex13()
		case "14":
			Ex14()
		case "15":
			Ex15()
		case "16":
			Ex16()
		case "17":
			Ex17()
		case "18":
			Ex18()
		case "19":
			Ex19()
		case "20":
			Ex20()
		case "21":
			Ex21()
		case "22":
			Ex22()
		case "23":
			Ex23()
		case "24":
			Ex24()
		case "25":
			Ex25()
		case "26":
			Ex26()
		default:
			fmt.Println("Invalid exercise number or command")
		}
		fmt.Println("Type here an exercise number ('exit or -1' to exit):")
		fmt.Scan(&exercise)
	}
}
