package main

import (
	"fmt"
)

func f1() {

	fmt.Println("F1")

}

func f2(p1 string, p2 string) {

	fmt.Printf("F2: %s %s \n", p1, p2)

}

func f3() string {
	return "F3"
}

func f4(p1 string, p2 string) string {

	return fmt.Sprintf("F4: %s %s \n", p1, p2)

}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {

	f1()

	f2("Olá", "Kayan")

	fmt.Println(f3())

	fmt.Println(f4("Olá", "Natacha"))

	r1, r2 := f5()

	fmt.Println(r1)

	fmt.Println(r2)

}
