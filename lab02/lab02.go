package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := Sum(n)
	fmt.Println(result)
}

func Sum(n int64) string {
	// TODO: Finish this function
	var sum int = 0
	var s string = ""
	for i:=1;int64(i)<=n;i++ {
		if (i%7!=0){
			sum+=i
			s+=strconv.Itoa(i)

			if (int64(i)!=n) {
				if(!(n==7&&i==6)){
					s+="+"
				}
			}
		}
	}
	s+="="
	s+=strconv.Itoa(sum)
	return s
}

