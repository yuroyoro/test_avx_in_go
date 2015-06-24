package main

/*
#cgo CFLAGS: -mavx
#include <stdio.h>
#include <stdlib.h>
#include <immintrin.h>//AVX: -mavx

void avx_add(const size_t n, float *x, float *y, float *z)
{
	static const size_t single_size = 8; //単精度は8つずつ計算
	const size_t end = n / single_size;

  //AVX 専用の型にデータをロードする
  __m256 *vz = (__m256 *)z;
  __m256 *vx = (__m256 *)x;
  __m256 *vy = (__m256 *)y;

  for(size_t i=0; i<end; ++i) {
    vz[i] = _mm256_add_ps(vx[i], vy[i]); //AVX を用いる SIMD 演算
	}
}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

/*
参考: Intel AVX を使用して SIMD 演算を試してみる - kawa0810 のブログ http://kawa0810.hateblo.jp/entry/20120303/1330797281
*/

func mmMalloc(size int) []float32 {
	length := size * 32
	ptr := C._mm_malloc((C.size_t)(length), 32)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ptr)),
		Len:  length,
		Cap:  length,
	}
	goSlice := *(*[]float32)(unsafe.Pointer(&hdr))
	return goSlice
}

func avxAdd(size int, x, y, z []float32) {
	C.avx_add((C.size_t)(size), (*C.float)(&x[0]), (*C.float)(&y[0]), (*C.float)(&z[0]))
}
