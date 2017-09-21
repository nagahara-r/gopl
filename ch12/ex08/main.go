// Copyright © 2017 Yuki Nagahara
// 練習12-8　Decode型にUnmarshalを適応させます。

package main

import (
	"log"
	"math"
	"os"
	"reflect"

	"github.com/naga718/golang-practice/ch12/ex08/sexpr"
)

func main() {
	unmarshalIf()
}

func unmarshalIf() {
	type IfTest struct {
		// Generics
		Booli       interface{}
		Inti        interface{}
		Int8i       interface{}
		Int16i      interface{}
		Int32i      interface{}
		Int64i      interface{}
		Float32i    interface{}
		Float64i    interface{}
		Complex64i  interface{}
		Complex128i interface{}
		Stringi     interface{}

		// Slices
		Boolsi       interface{}
		Intsi        interface{}
		Int8si       interface{}
		Int16si      interface{}
		Int32si      interface{}
		Int64si      interface{}
		Float32si    interface{}
		Float64si    interface{}
		Complex64si  interface{}
		Complex128si interface{}
		Stringsi     interface{}
	}
	testdata := IfTest{
		Booli:       bool(true),
		Inti:        int(0),
		Int8i:       int8(math.MaxInt8),
		Int16i:      int16(math.MaxInt16),
		Int32i:      int32(math.MaxInt32),
		Int64i:      int64(math.MaxInt64),
		Float32i:    float32(math.MaxFloat32),
		Float64i:    float64(math.MaxFloat64),
		Complex64i:  complex64(complex(math.MaxFloat32, math.MaxFloat32)),
		Complex128i: complex(math.MaxFloat64, math.MaxFloat64),
		Stringi:     "This is String",

		Boolsi:       []bool{bool(false), bool(true)},
		Intsi:        []int{int(0), int(1)},
		Int8si:       []int8{int8(math.MinInt8), int8(math.MaxInt8)},
		Int16si:      []int16{int16(math.MinInt16), int16(math.MaxInt16)},
		Int32si:      []int32{int32(math.MinInt32), int32(math.MaxInt32)},
		Int64si:      []int64{int64(math.MinInt64), int64(math.MaxInt64)},
		Float32si:    []float32{float32(0), float32(math.MaxFloat32)},
		Float64si:    []float64{float64(0), float64(math.MaxFloat64)},
		Complex64si:  []complex64{complex64(complex(0, 0)), complex64(complex(math.MaxFloat32, math.MaxFloat32))},
		Complex128si: []complex128{complex(0, 0), complex(math.MaxFloat64, math.MaxFloat64)},
		Stringsi:     []string{"This is String 1.", "文字列2です。"},
	}

	// Decode in Stdin
	var ifTest IfTest
	dec := sexpr.NewDecoder(os.Stdin)
	if err := dec.Unmarshal(&ifTest); err != nil {
		log.Fatalf("Unmarshal failed: %v", err)
	}
	log.Printf("OriginalData = %+v\n", testdata)
	log.Printf("Unmarshal()  = %+v\n", ifTest)

	// Check equality.
	log.Printf("Check Equality = %v", reflect.DeepEqual(ifTest, testdata))
}
