package main

import (
	"flag"
	"fmt"
	// "unsafe"
)

const size = 2048

/*
参考: Intel AVX を使用して SIMD 演算を試してみる - kawa0810 のブログ http://kawa0810.hateblo.jp/entry/20120303/1330797281
*/
func main() {

	var (
		avx    bool
		output bool
		times  int
	)

	flag.BoolVar(&avx, "avx", false, "use avx")
	flag.BoolVar(&output, "output", false, "print result")
	flag.IntVar(&times, "times", 10, "benchmarking count")
	flag.Parse()

	for n := 0; n < times; n++ {
		x := mmMalloc(size)
		y := mmMalloc(size)
		z := mmMalloc(size)

		for i := 0; i < size; i++ {
			x[i] = float32(i) * 0.1
		}
		for i := 0; i < size; i++ {
			y[i] = float32(i+1) * 0.2
		}
		for i := 0; i < size; i++ {
			z[i] = 0.0
		}

		// check alignment
		// fmt.Println(uintptr(unsafe.Pointer(&x[0]))%32 == 0)
		// fmt.Println(uintptr(unsafe.Pointer(&y[0]))%32 == 0)
		// fmt.Println(uintptr(unsafe.Pointer(&z[0]))%32 == 0)

		if avx {
			fmt.Printf("Using AVX: %d\n", n)
			avxAdd(size, x, y, z)
		} else {
			fmt.Printf("Go loop: %d\n", n)
			avxAdd(size, x, y, z)
			for i := 0; i < size; i++ {
				z[i] = x[i] + y[i]
			}
		}
		if output {
			for i := 0; i < size; i++ {
				fmt.Printf("%v + %v = %v\n", x[i], y[i], z[i])
			}
		}
	}
}
