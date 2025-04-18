package main

import "fmt"

func merge(left, right []int) []int {
	array := make([]int, 0)
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(array, right...)
		} else if len(right) == 0 {
			return append(array, left...)
		} else if left[0] < right[0] {
			array = append(array, left[0])
			left = left[1:]
		} else {
			array = append(array, right[0])
			right = right[1:]
		}
	}
	return array
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	channel := make(chan bool)
	var left, right []int
	go func() {
		left = mergeSort(data[:mid])
		channel <- true
	}()
	<-channel
	right = mergeSort(data[mid:])

	return merge(left, right)
}

func main() {
	data := []int{2, 5, 4, 3, 6, 8, 9, 10, 15, 9}

	mergeSort := mergeSort(data)
	fmt.Println("Unsorted array is:", data)
	fmt.Println("Sorted array is:", mergeSort)

}

/*
Now we have to make sure that Merge(left,right) is executed only once we get the return values
from both the recursive calls, i.e. both the left and right have been updated before Merge(left,right)
has to execute. Hence, we introduce a channel of type bool on line 26 and send true on it as soon as
left = MergeSort(data[:mid]) is executed (line 32). The <-done operation blocks the code on line 35
before the statement Merge(left,right) so that it does not proceed until our goroutine has finished.
After the goroutine has finished and we receive true on the done channel, the code proceeds forward to
Merge(left,right) statement on line 36.
*/
