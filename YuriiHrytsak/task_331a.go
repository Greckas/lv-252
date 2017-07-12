package main

import (
	"fmt"
	"math"
	"strconv"
)

func three_powers(n float64)(float64,float64,float64)  {
	var i,j,k float64 = 1,1,1
	var x = make([]float64,10000)
	for i =1;i<n;i++{
		for j =1; j < n;j++{
			for k =1;k<n;k++{
				if ((math.Pow(i,2)+math.Pow(j,2)+math.Pow(k,2)) == n){
					fmt.Println(i,j,k)
					x = append(x,i,j,k)
					return i,j,k
				}
			}
		}

	}
	var sum float64
	for _, num := range x {
		sum += num
	}
	if sum == 0 {
		fmt.Println("not found such numbers!")
	}
	return 0,0,0

}



func main(){
	fmt.Print("Enter number: ")
	var input string
	fmt.Scanln(&input)
	n, _ := strconv.ParseFloat(input, 64);
	three_powers(n)

}
