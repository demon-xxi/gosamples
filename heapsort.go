package main

import "fmt" 

func main(){
   array := []int{34, -6, 1, 45, 874, 565, 87435, 1, -34- 2, 0, 0, 123, -1, 4}
   fmt.Println(array)
   
   heapsort(array)
   
   fmt.Println(array)
}

func heapsort(array []int){
   
   buildMaxHeap(array)
   // fmt.Println(array)
   for len(array) > 1 {
      arrswap(array, 0, len(array)-1)
      array = array[:len(array)-1]
      siftDown(array, 0)
      // fmt.Println(array)
   }
   
}

func parent(idx int) int {
   return (idx-1) / 2
}

func leftChild(idx int) int {
   return idx*2 + 1
}

func rightChild(idx int) int {
   return idx*2 + 2
}

func siftDown(array []int, idx int){
  // fmt.Printf("siftDown(%d)", idx)
   for root := idx; leftChild(root) <= len(array)-1; {
      child := leftChild(root)
      swap := root
      if array[swap] < array[child] {
         swap = child
      }
      if child+1 < len(array)-1 && array[swap] < array[child+1] {
         swap = child + 1
      }
      if swap == root {
         break
      } else {
         arrswap(array, root, swap)
         root = swap
      }
   }
   
}

func arrswap(array []int, idx1 int, idx2 int){
   array[idx1], array[idx2] = array[idx2], array[idx1] 
}

func buildMaxHeap(array []int){
   start := parent(len(array) - 1)
   for ; start >= 0; start-- {
      siftDown(array, start)
   }
}