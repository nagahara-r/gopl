// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習12-4: S式をプリティプリントしながらencodeを実施します。

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf *bytes.Buffer, v reflect.Value, depth int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), depth)

	case reflect.Array, reflect.Slice: // (value ...)
		depth++
		writeDepth(buf, depth)

		buf.WriteByte('(')

		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i), depth); err != nil {
				return err
			}

		}
		buf.WriteByte(')')
		depth--
		writeDepth(buf, depth)

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		depth++

		for i := 0; i < v.NumField(); i++ {
			writeDepth(buf, depth)
			if i > 0 {
				buf.WriteByte(' ')
			}

			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), depth); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		depth++

		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			writeDepth(buf, depth)
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')

			if err := encode(buf, key, depth); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), depth); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
		depth--
		writeDepth(buf, depth)

	case reflect.Bool:
		fmt.Fprintf(buf, "%v", v.Bool())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%f, %f)", real(v.Complex()), imag(v.Complex()))

	case reflect.Interface:
		fmt.Fprintf(buf, "#I(\"%s\" ", reflect.ValueOf(v.Interface()).Type())

		encode(buf, reflect.ValueOf(v.Interface()), depth)
		fmt.Fprintf(buf, ")")
	default: // chan, func
		// chan はシステムコールに関わるため、シリアライズできない。
		// func は関数であり、シリアライズできない。
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func writeDepth(buf *bytes.Buffer, depth int) {
	buf.WriteByte('\n')
	for i := 0; i < depth; i++ {
		buf.WriteByte('\t')
	}
}

//!-encode
