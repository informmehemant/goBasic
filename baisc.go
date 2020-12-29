package main

import "fmt"

func main() {
	var val = []int{1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(val))
}
func findMaxConsecutiveOnes(nums []int) int  {   
   var  count int = 0
    temp := 0
    for _, value := range nums {
        if value == 1 {
            count++
            if temp < count {
                temp = count;
                
            }
        } else {
            count = 0
            
        }
        
    }
    return temp
}