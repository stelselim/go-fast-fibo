package go_fib_module

import "errors"

func FastFibo(num int) (int, error) {
	if num < 0 {
		return -1, errors.New("")
	}
	return calculateFib(num), nil
}

func calculateFib(num int) int {
	hashMap := make(map[int]int)

	hashMap[0] = 0
	hashMap[1] = 1

	for i := 2; i <= num; i++ {
		hashMap[i] = hashMap[i-1] + hashMap[i-2]

	}
	return hashMap[num]
}
