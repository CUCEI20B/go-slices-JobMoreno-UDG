package main

import "fmt"

func main()  {
	var i,valor,suma int
	var s[] int
	fmt.Scan(&i)
	suma = 0
	for x:=0;x<i;x++{
		fmt.Scan(&valor)
		s = append(s,valor)
		suma = suma+valor
	}
	fmt.Print(suma) 
}