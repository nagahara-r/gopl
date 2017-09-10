// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習12-3: float, complex, bool, chan, func, interface の実装をします。
// 練習12-7: S式のストリームデコーダを作成します。

// See page 344.

// Package sexpr provides a means for converting Go objects to and
// from S-expressions.
package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
	"unsafe"
)

var reflectTypes = map[string]reflect.Type{
	// "invalid":        reflect.TypeOf(nil),
	"bool":       reflect.TypeOf(bool(true)),
	"int":        reflect.TypeOf(int(0)),
	"int8":       reflect.TypeOf(int8(0)),
	"int16":      reflect.TypeOf(int16(0)),
	"int32":      reflect.TypeOf(int32(0)),
	"int64":      reflect.TypeOf(int64(0)),
	"uint":       reflect.TypeOf(uint(0)),
	"uint8":      reflect.TypeOf(uint8(0)),
	"uint16":     reflect.TypeOf(uint16(0)),
	"uint32":     reflect.TypeOf(uint32(0)),
	"uint64":     reflect.TypeOf(uint64(0)),
	"uintptr":    reflect.TypeOf(uintptr(0)),
	"float32":    reflect.TypeOf(float32(0)),
	"float64":    reflect.TypeOf(float64(0)),
	"complex64":  reflect.TypeOf(complex64(0)),
	"complex128": reflect.TypeOf(complex128(0)),
	"string":     reflect.TypeOf(""),
	"unsafe":     reflect.TypeOf(unsafe.Pointer(uintptr(0))),

	// Slice は型が基本形であればSliceにすることでサポートする。
	"[]bool":           reflect.TypeOf([]bool{}),
	"[]int":            reflect.TypeOf([]int{}),
	"[]int8":           reflect.TypeOf([]int8{}),
	"[]int16":          reflect.TypeOf([]int16{}),
	"[]int32":          reflect.TypeOf([]int32{}),
	"[]int64":          reflect.TypeOf([]int64{}),
	"[]uint":           reflect.TypeOf([]uint{}),
	"[]uint8":          reflect.TypeOf([]uint8{}),
	"[]uint16":         reflect.TypeOf([]uint16{}),
	"[]uint32":         reflect.TypeOf([]uint32{}),
	"[]uint64":         reflect.TypeOf([]uint64{}),
	"[]uintptr":        reflect.TypeOf([]uintptr{}),
	"[]float32":        reflect.TypeOf([]float32{}),
	"[]float64":        reflect.TypeOf([]float64{}),
	"[]complex64":      reflect.TypeOf([]complex64{}),
	"[]complex128":     reflect.TypeOf([]complex128{}),
	"[]string":         reflect.TypeOf([]string{}),
	"[]unsafe.Pointer": reflect.TypeOf([]unsafe.Pointer{}),

	// 以下は型へのマッピングが困難なため非サポート
	// "array"
	// "interface":      reflect.Interface,
	// "map":            reflect.Map,
	// "struct":         struct{},

	// 以下はマーシャリングできない
	// "chan":           reflect.Chan,
	// "func":           reflect.Func,
}

// A Decoder reads and decodes S-expression values from an input stream.
// 練習12-7
type Decoder struct {
	r   io.Reader
	buf []byte
}

// NewDecoder は新しいデコーダをリーダから作成します。
// 練習12-7
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Decode はリーダからS式をデコードします。
// 練習12-7
func (dec *Decoder) Decode(i interface{}) (err error) {
	data, err := ioutil.ReadAll(dec.r)
	if err != nil {
		return err
	}
	err = Unmarshal(data, i)
	if err != nil {
		return err
	}

	return nil
}

//!+Unmarshal
// Unmarshal parses S-expression data and populates the variable
// whose address is in the non-nil pointer out.
func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() // get the first token
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

//!-Unmarshal

//!+lexer
type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

//!-lexer

// The read function is a decoder for a small subset of well-formed
// S-expressions.  For brevity of our example, it takes many dubious
// shortcuts.
//
// The parser assumes
// - that the S-expression input is well-formed; it does no error checking.
// - that the S-expression input corresponds to the type of the variable.
// - that all numbers in the input are non-negative decimal integers.
// - that all keys in ((key value) ...) struct syntax are unquoted symbols.
// - that the input does not contain dotted lists such as (1 2 . 3).
// - that the input does not contain Lisp reader macros such 'x and #'x.
//
// The reflection logic assumes
// - that v is always a variable of the appropriate type for the
//   S-expression value.  For example, v must not be a boolean,
//   interface, channel, or function, and if v is an array, the input
//   must have the correct number of elements.
// - that v in the top-level call to read has the zero value of its
//   type and doesn't need clearing.
// - that if v is a numeric variable, it is a signed integer.

//!+read
func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names.
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		} else if lex.text() == "false" || lex.text() == "true" {
			b, _ := strconv.ParseBool(lex.text())
			v.SetBool(b)
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		v.SetString(s)
		lex.next()
		return
	case '-':
		lex.next()
		setNumber("-"+lex.text(), lex, v)
		lex.next()
		return
	case scanner.Int, scanner.Float:
		setNumber(lex.text(), lex, v)
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // consume ')'
		return
	case '#':
		lex.next() // C
		if lex.text() == "C" {
			lex.next() // (
			lex.next() // 1.00000
			real, _ := strconv.ParseFloat(lex.text(), 64)
			lex.next() // ,
			lex.next() // 2.00000
			imag, _ := strconv.ParseFloat(lex.text(), 64)
			v.SetComplex(complex(real, imag))

			lex.next()
		} else if lex.text() == "I" {
			lex.next() // (
			lex.next() // Type
			readList(lex, v)
		}
		lex.next() // consume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func setNumber(str string, lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Int:
		i, _ := strconv.ParseInt(str, 0, 64) // NOTE: ignoring errors
		//i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		v.SetInt(int64(i))
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64) // NOTE: ignoring errors
		v.SetFloat(f)
	}
}

//!-read

//!+readlist
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	case reflect.Interface:
		l := strings.Trim(lex.text(), "\"")
		iv := reflect.New(reflectTypes[l]).Elem()

		lex.next()
		read(lex, iv)
		v.Set(iv)

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

//!-readlist
