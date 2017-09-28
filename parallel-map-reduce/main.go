package main

import (
	"fmt"
)

func main() {
	// 计算 1~N 范围内整数平方和
	// 伪代码： 1..N.map(x => x * x).sum()
	N := uint(100)

	fmt.Printf("normal_version_sum: %d\n", normal_version_sum(N))
	fmt.Printf("parallel_version_sum: %d\n", parallel_version_sum(N))
}

func normal_version_sum(N uint) (sum uint) {
	for i := uint(1); i <= N; i++ {
		sum += i * i
	}

	return
}

func parallel_version_sum(N uint) uint {
	ch := make(chan uint, N)

	for i := uint(1); i <= N; i++ {
		i := i
		go func() {
			ch <- i * i
		}()
	}

	sum := uint(0)
	i := uint(1)
	for x := range ch {
		// fmt.Println(i, x, math.Sqrt(float64(x)))

		sum += x

		i++
		if i > N {
			break
		}
	}

	return sum
}

func parallel_version_sum2(N uint) uint {
	ch := make(chan uint, N)

	for i := uint(1); i <= N; i++ {
		i := i
		go func() {
			ch <- i * i
		}()
	}

	sum := uint(0)
	for i := uint(1); i <= N; i++ {
		sum += <-ch
	}

	return sum
}
