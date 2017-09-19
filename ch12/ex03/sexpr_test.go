// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara

package sexpr

import (
	"math"
	"reflect"
	"strings"
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
	if err = Unmarshal(data, &movie); err != nil {
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

// 練習12-3, 12-10(Float)
func TestFloat(t *testing.T) {
	tests := []float64{
		0, 1, 1.5, -100.2, math.MaxFloat64, math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64 * -1,
	}

	for _, test := range tests {
		b, err := Marshal(test)
		if err != nil {
			t.Errorf("%v", err)
		}
		result := float64(0)
		err = Unmarshal(b, &result)
		if err != nil {
			t.Errorf("%v", err)
		}
		if test != result {
			t.Errorf("Parse result = %v, but expected %v", result, test)
		}
	}
}

// 練習12-3, 12-10(Complex)
func TestComplex(t *testing.T) {
	tests := []complex128{
		complex(0, 0),
		complex(0, 1),
		complex(1.5, 2.3),
		complex(-100.2, 5),
		complex(math.MaxFloat64, math.MaxFloat64),
		complex(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64),
	}

	for _, test := range tests {
		b, err := Marshal(test)
		if err != nil {
			t.Errorf("%v", err)
		}

		result := complex(0, 0)
		err = Unmarshal(b, &result)
		if err != nil {
			t.Errorf("%v", err)
		}
		if test != result {
			t.Errorf("Parse result = %v, but expected %v", result, test)
		}
	}
}

// 練習12-3, 12-10(Bool)
func TestBool(t *testing.T) {
	tests := []bool{
		true, false,
	}

	for _, test := range tests {
		b, err := Marshal(test)
		if err != nil {
			t.Errorf("%v", err)
		}

		var result bool
		err = Unmarshal(b, &result)
		if err != nil {
			t.Errorf("%v", err)
		}
		if test != result {
			t.Errorf("Parse result = %v, but expected %v", result, test)
		}
	}
}

// 練習12-3, 12-10 (interface)
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
	if err = Unmarshal(data, &ifTest); err != nil {
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

// 練習12-6: ゼロ値をマーシャリングしない（アンマーシャリングも可能かどうかチェック）
func TestZeroValue(t *testing.T) {
	type teststruct struct {
		ZeroBool    bool
		ZeroInt     int
		ZeroFloat   float64
		ZeroComplex complex128
		ZeroString  string
		Non0Int     int
	}
	test := teststruct{}
	test.Non0Int = 1

	b, err := Marshal(test)
	if err != nil {
		t.Errorf("%v", err)
	}

	// "Zero" がつく要素がマーシャリングされていないか
	if strings.Contains(string(b), "Zero") {
		t.Errorf("ZeroValue Included: \n%v", string(b))
	}

	result := teststruct{}
	err = Unmarshal(b, &result)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(test, result) {
		t.Errorf("Parse result = %v, but expected %v", result, test)
	}
}

// 練習12-13: フィールドタグ実装を確認
func TestFieldTag(t *testing.T) {
	type teststruct struct {
		Field1 bool `sexpr:"thisisfield1"`
		Field2 int  `sexpr:"-"` // ignoreing
	}
	test := teststruct{true, 100}
	expected := teststruct{true, 0} // Field2 は無視してほしいのでUnmarshal後の期待はゼロ値

	b, err := Marshal(test)
	if err != nil {
		t.Errorf("%v", err)
	}

	// "thisisfield1" がタグとして採用されているか
	if !strings.Contains(string(b), "thisisfield1") {
		t.Errorf("Field Tag not Used: \n%v", string(b))
	}

	// "Field2" が正しく無視されているか
	if strings.Contains(string(b), "Field2") {
		t.Errorf("Field ignoreing failed: \n%v", string(b))
	}

	result := teststruct{}
	err = Unmarshal(b, &result)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result = %v, but expected %v", result, expected)
	}
}
