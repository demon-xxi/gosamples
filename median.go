package main

import "fmt" 

func main(){
   array :=  []int{1, 1, 3, 3, 2, 2}
   fmt.Println(array)
   
   sort(array)
   
   fmt.Println(array)
}

func mid(arr []int) int{
    sum := 0;
    for _, val := range arr{
        sum += val
    }
    return sum/len(arr)
}

func sort(arr []int){
   m := mid(arr)
   
   mpos := -1
   for i, val := range arr{
    if val > m {
        mpos = i
        break
    }
   }
   
   n := len(arr)
   for i, j := 1, mpos ; i < mpos && j <= n-1 ; i, j = i + 2, j + 2 {
       fmt.Printf("%d <=> %d\n", i ,j)
       tmp := arr[i]
       arr[i] = arr[j]
       arr[j] = tmp
   }
}