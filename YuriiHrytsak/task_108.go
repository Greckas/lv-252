package main

import (
	"math"
	"fmt"
	"strconv"
)

func less_numb(a float64) {
	var b = 0
	for i := 0; math.Pow(2,float64(i)) <= a; i++ {
		if float64(a) >= math.Pow(2,float64(i)){
			b += 1
	}else{
		continue}

	}
	fmt.Println(b)
}

func main(){
	fmt.Print("Enter number: ")
	var input string
	fmt.Scanln(&input)
	a, _ := strconv.ParseFloat(input, 64);
	less_numb(float64(a))
}




