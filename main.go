package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)
func main(){
	if len(os.Args) != 2{
		fmt.Println("please enter file name")
		return
	}
	data,err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	number := strings.Fields(string(data))
	var numbers []int
	for _,v:= range number{
		value,_ :=strconv.Atoi(v)
		numbers = append(numbers,value)
		
	}
	
	fmt.Println("Average: ",Average(numbers))
	fmt.Println("Median: ",Median(numbers))
	fmt.Println("Variance: ",Variance(numbers))
	fmt.Println("Standard Deviation: ",Standard_Deviation(numbers))
}
func Average(d1 []int)float64{
	sum :=0
	for i:= 0;i<len(d1);i++{
		sum += d1[i]
	}
	return math.Round(float64(sum)/float64(len(d1)))

}
func Median(d2 []int)float64{
	for i := 0; i < len(d2); i++ {
		for j := 0; j < len(d2)-1-i; j++ {
			if d2[j] > d2[j+1] {
				d2[j], d2[j+1] = d2[j+1], d2[j]
		}
	}
}

	if len(d2)%2==0{
		return math.Round((float64((d2[len(d2)/2])) + float64(d2[(len(d2)/2)-1]))/2 )
	}
	return math.Round(float64((d2[(len(d2)/2)])))
}
func Variance(d3 []int)float64{
	mean := Average(d3)
	var sum float64
	for _, v := range d3 {
		diff := float64(v) - mean
		sum += diff * diff
}
	return math.Round(sum / float64(len(d3)))

		
}
func Standard_Deviation(d4 []int)float64{
	return math.Round(math.Sqrt(Variance(d4)))
}
