package main

func main() {
	nums := []int{1, 2, 3, 4}
	for i := range 10 {
		nums = append(nums, i)
	}
}
