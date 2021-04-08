package main

import (
	"bufio"
	"os"
	"strconv"
)

/******
### Problem statement and example
You have an array arr of length n where arr[i] = (2 * i) + 1 for all valid values of i (i.e. 0 <= i < n).
Given an integer n, the length of the array. Return the minimum number of operations needed to make all the elements of arr equal.

Input: n = 3
Output: 2
Explanation: arr = [1, 3, 5]
First operation choose x = 2 and y = 0, this leads arr to be [2, 3, 4]
In the second operation choose x = 2 and y = 0 again, thus arr = [3, 3, 3].

*******/

/**
	generated array will be in the series 1, 3, 5, 7, 9 .... since all the numbers are getting
	incremented by 2 the series can also be written as 2 (1, 2, 3, 4, ..). with this series all
	the numbers should equalate to mid number.

	eg.

	arr = [1, 3, 5, 7, 9]

	left = 1
	right = 9

	mid = left+right/2 = 5  //so after performing increments and decrements operations, the final array will be [5, 5, 5, 5, 5]

 */

func minOperations(n int) int {
	if n==1 {
		return 0
	}

	if n==2 {
		return 1
	}

	first := 1
	last := (2*(n-1)) + 1
	avg := (first + last)/2
	ans := last - avg
	return ans + minOperations(n-2)
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	input, _ := sc.ReadString('\n')
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	minOperations(num)
}