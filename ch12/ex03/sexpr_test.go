// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara

package sexpr

import (
	"math"
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		// Sequel はsexprでマーシャリングされません。
		Sequel *string `sexpr:"-"`

		// Award はisgetawardという名前でマーシャリングされます。
		Award      bool `sexpr:"isgetaward"`
		ComplexNum complex64
		ZeroValue  int
		ZeroStruct struct{}

		Description interface{}
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},

		Award:      true,
		ComplexNum: 1 + 2i,

		//Description: MovieDescription{"This is Description.", "This is Author"},
		Description: "This is Description.",
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestInterface(t *testing.T) {
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

	// Encode it
	data, err := Marshal(testdata)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var ifTest IfTest
	if err := Unmarshal(data, &ifTest); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal()  = %+v\n", ifTest)
	t.Logf("OriginalData = %+v\n", testdata)

	// Check equality.
	if !reflect.DeepEqual(ifTest, testdata) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(testdata)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}
