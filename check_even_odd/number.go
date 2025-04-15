package main

import "fmt"

type number []int64

func numberRange(r int) number {
	newNum := number{}
	for i := 1; i <= r; i++ {
		newNum = append(newNum, int64(i))
	}
	return newNum
}

func (n number) checkNumber(){
	var  even, odd int64 = 0, 0
	for _,num := range n{
		if num % 2 == 0 {
			fmt.Printf("The number %d is Even\n", num)
			even++
			}else{
			fmt.Printf("The number %d is Odd\n", num)
			odd ++
		}
		fmt.Printf("Total Even: %d\nTotal Odd: %d\n", even, odd)
	}
}