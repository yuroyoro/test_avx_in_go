package main

import (
	"testing"
)

/*
参考: Intel AVX を使用して SIMD 演算を試してみる - kawa0810 のブログ http://kawa0810.hateblo.jp/entry/20120303/1330797281

$ system_profiler SPHardwareDataType
Hardware:

    Hardware Overview:

      Model Name: MacBook Pro
      Model Identifier: MacBookPro11,3
      Processor Name: Intel Core i7
      Processor Speed: 2.3 GHz
      Number of Processors: 1
      Total Number of Cores: 4
      L2 Cache (per Core): 256 KB
      L3 Cache: 6 MB
      Memory: 16 GB
      Boot ROM Version: MBP112.0138.B02
      SMC Version (system): 2.19f3
      Serial Number (system): C02LJG4RFD57
      Hardware UUID: 249473D1-1C18-536A-9B95-DCE616DE5749

$ go test -benchmem -bench=.
testing: warning: no tests to run
PASS
BenchmarkAvxAdd  3000000               415 ns/op               0 B/op          0 allocs/op
BenchmarkNonAlignedAvxAdd        1000000              1143 ns/op               0 B/op          0 allocs/op
BenchmarkGoAdd   1000000              2059 ns/op               0 B/op          0 allocs/op
ok      github.com/yuroyoro/test_avx_in_go      4.914s
*/

func BenchmarkAvxAdd(b *testing.B) {

	x := mmMalloc(size)
	y := mmMalloc(size)
	z := mmMalloc(size)

	defer mmFree(x)
	defer mmFree(y)
	defer mmFree(z)

	for i := 0; i < size; i++ {
		x[i] = float32(i) * 0.1
	}
	for i := 0; i < size; i++ {
		y[i] = float32(i+1) * 0.2
	}
	for i := 0; i < size; i++ {
		z[i] = 0.0
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		avxAdd(size, x, y, z)
	}

}
func BenchmarkNonAlignedAvxAdd(b *testing.B) {

	x := make([]float32, size)
	y := make([]float32, size)
	z := make([]float32, size)

	for i := 0; i < size; i++ {
		x[i] = float32(i) * 0.1
	}
	for i := 0; i < size; i++ {
		y[i] = float32(i+1) * 0.2
	}
	for i := 0; i < size; i++ {
		z[i] = 0.0
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		avxAddu(size, x, y, z)
	}

}

func BenchmarkGoAdd(b *testing.B) {

	x := make([]float32, size)
	y := make([]float32, size)
	z := make([]float32, size)

	for i := 0; i < size; i++ {
		x[i] = float32(i) * 0.1
	}
	for i := 0; i < size; i++ {
		y[i] = float32(i+1) * 0.2
	}
	for i := 0; i < size; i++ {
		z[i] = 0.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < size; i++ {
			z[i] = x[i] + y[i]
		}
	}

}
