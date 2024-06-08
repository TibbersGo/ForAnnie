package main

import (
	"fmt"
	"math"
)

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a subprocess.
	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func backspaceCompare(s string, t string) bool {
	sl, tl := len(s)-1, len(t)-1
	skipS, skipT := 0, 0
	for sl >= 0 && tl >= 0 {
		if s[sl] == '#' {
			skipS++
			sl--
			continue
		}
		if skipS > 0 {
			sl--
			skipS--
			continue
		}
		if t[tl] == '#' {
			skipT++
			tl--
			continue
		}
		if skipT > 0 {
			tl--
			skipT--
			continue
		}
		if s[sl] != t[tl] {
			return false
		}
		sl--
		tl--
	}
	return true
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	sum := 0
	mNum := math.MaxInt32
	for fast < len(nums) {
		sum = sum + nums[fast]
		for sum >= target {
			mNum = min(fast-slow+1, target)
			sum = sum - nums[slow]
			slow++
		}
		fast++
	}
	return mNum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func generateMatrix(n int) [][]int {
	loop := n / 2
	res := make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, n)
	}
	if n%2 == 1 {
		res[n/2][n/2] = n * n
	}
	startX, startY := 0, 0
	offset := 1
	num := 1
	for loop > 0 {
		x, y := startX, startY
		for x = startX; x < n-offset; x++ {
			fmt.Printf("x: %d, y: %d, num: %d \n", x, y, num)
			res[y][x] = num
			num++
		}
		fmt.Println("x: ", x)
		fmt.Println("y: ", y)
		for y = startY; y < n-offset; y++ {
			fmt.Printf("x: %d, y: %d, num: %d \n", x, y, num)

			res[y][x] = num
			num++
		}
		fmt.Println("x: ", x)
		fmt.Println("y: ", y)
		for x = n - offset; x > startX; x-- {
			fmt.Printf("x: %d, y: %d, num: %d \n", x, y, num)

			res[y][x] = num
			num++
		}
		fmt.Println("x: ", x)
		fmt.Println("y: ", y)

		for y = n - offset; y > startY; y-- {

			fmt.Printf("x: %d, y: %d, num: %d \n", x, y, num)
			res[y][x] = num
			num++
		}
		fmt.Println("x: ", x)
		fmt.Println("y: ", y)

		startX++
		startY++
		offset++
		loop--
	}
	return res
}

func main() {
	res := generateMatrix(4)
	fmt.Println("res: ", res)
}
